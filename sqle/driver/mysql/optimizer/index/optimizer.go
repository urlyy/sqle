package index

import (
	"context"
	"fmt"
	"strings"

	"github.com/actiontech/sqle/sqle/driver/mysql/executor"
	"github.com/actiontech/sqle/sqle/driver/mysql/session"
	"github.com/actiontech/sqle/sqle/driver/mysql/util"
	indexoptimizer "github.com/actiontech/sqle/sqle/pkg/optimizer/index"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	defaultCalculateCardinalityMaxRow = 1000000
	defaultCompositeIndexMaxColumn    = 3
)

type Optimizer struct {
	*session.Context

	l *logrus.Entry

	// tables key is table name, use to match in execution plan.
	tables map[string]*tableInSelect

	// optimizer options:
	calculateCardinalityMaxRow int
	compositeIndexMaxColumn    int
	createIndexStatement       func(string, ...string) string
}

func NewOptimizer(log *logrus.Entry, ctx *session.Context, opts ...optimizerOption) *Optimizer {
	log = log.WithField("optimizer", "index")

	optimizer := &Optimizer{
		Context:                    ctx,
		l:                          log,
		tables:                     make(map[string]*tableInSelect),
		createIndexStatement:       defaultCreateIndexStatement,
		compositeIndexMaxColumn:    defaultCompositeIndexMaxColumn,
		calculateCardinalityMaxRow: defaultCalculateCardinalityMaxRow,
	}

	for _, opt := range opts {
		opt.apply(optimizer)
	}

	return optimizer
}

type OptimizeResult struct {
	TableName      string
	IndexedColumns []string

	Reason string
}

// tableInSelect store the information of a table in select statement for later optimize.
// 1. when we find a table in single table select statement, we will store the select statement.
// 2. when we find a table in join statement, we will store the join on condition.
type tableInSelect struct {
	joinOnColumn   string
	singleTableSel *ast.SelectStmt
}

// Optimize give index advice for the select statement.
func (o *Optimizer) Optimize(ctx context.Context, selectStmt *ast.SelectStmt) ([]*OptimizeResult, error) {
	// select 1; ...
	if selectStmt.From == nil {
		return nil, nil
	}

	o.parseSelectStmt(selectStmt)

	restoredSQL, err := restoreSelectStmt(selectStmt)
	if err != nil {
		return nil, err
	}

	executionPlan, err := o.GetExecutionPlan(restoredSQL)
	if err != nil {
		return nil, errors.Wrap(err, "get execution plan when optimize")
	}

	executionPlan = removeDrivingTable(executionPlan)

	var needOptimizedTables []string
	for _, record := range executionPlan {
		if o.needOptimize(record) {
			needOptimizedTables = append(needOptimizedTables, record.Table)
		}
	}

	if len(needOptimizedTables) == 0 {
		return nil, nil
	}

	o.l.Infof("need optimize tables: %v", needOptimizedTables)

	var results []*OptimizeResult
	for _, tbl := range needOptimizedTables {
		table, ok := o.tables[tbl]
		if !ok {
			return nil, errors.Errorf("table %s not found when index optimize", tbl)
		}

		var result *OptimizeResult
		if table.joinOnColumn == "" {
			result, err = o.optimizeSingleTable(ctx, tbl, table.singleTableSel)
		} else {
			result = o.optimizeJoinTable(tbl)
		}

		results = append(results, result)
	}

	return results, nil
}

// SelectStmt:
//   1. single select on single table
//   2. single select on multiple tables, such join
//   3. multi select on multiple tables, such subqueries
func (o *Optimizer) parseSelectStmt(ss *ast.SelectStmt) {
	visitor := util.SelectStmtExtractor{}
	ss.Accept(&visitor)

	for _, ss := range visitor.SelectStmts {
		if ss.From == nil {
			continue
		}

		if ss.From.TableRefs.Right == nil {
			leftTable, ok := ss.From.TableRefs.Left.(*ast.TableSource)
			if !ok {
				continue
			}

			if leftTable.AsName.L != "" {
				o.tables[leftTable.AsName.L] = &tableInSelect{singleTableSel: ss}
			} else {
				o.tables[leftTable.Source.(*ast.TableName).Name.L] = &tableInSelect{singleTableSel: ss}
			}

		} else {
			if ss.From.TableRefs.On != nil {
				boe, ok := ss.From.TableRefs.On.Expr.(*ast.BinaryOperationExpr)
				if !ok {
					continue
				}

				leftCNE := boe.L.(*ast.ColumnNameExpr)
				rightCNE := boe.R.(*ast.ColumnNameExpr)

				o.tables[leftCNE.Name.Table.L] = &tableInSelect{joinOnColumn: leftCNE.Name.Name.L}
				o.tables[rightCNE.Name.Table.L] = &tableInSelect{joinOnColumn: rightCNE.Name.Name.L}
			}
		}
	}
}

func (o *Optimizer) optimizeSingleTable(ctx context.Context, tbl string, ss *ast.SelectStmt) (*OptimizeResult, error) {
	var (
		err            error
		optimizeResult *OptimizeResult
	)

	optimizeResult, err = o.doGeneralOptimization(ctx, ss)
	if err != nil {
		return nil, err
	}

	if optimizeResult == nil {
		optimizeResult, err = o.doSpecifiedOptimization(ctx, ss)
		if err != nil {
			return nil, err
		}
	}

	if optimizeResult == nil {
		return nil, nil
	}

	if len(optimizeResult.IndexedColumns) > o.compositeIndexMaxColumn {
		optimizeResult.IndexedColumns = optimizeResult.IndexedColumns[:o.compositeIndexMaxColumn]
	}

	needIndex, err := o.needIndex(optimizeResult.TableName, optimizeResult.IndexedColumns...)
	if err != nil {
		return nil, err
	}

	if !needIndex {
		return nil, nil
	}

	o.l.Infof("table:%s, indexed columns:%v, reason:%s", optimizeResult.TableName, optimizeResult.IndexedColumns, optimizeResult.Reason)

	rowCount, err := o.GetTableRowCount(extractTableNameFromAST(ss, tbl))
	if err != nil {
		return nil, errors.Wrap(err, "get table row count when optimize")
	}
	if rowCount < o.calculateCardinalityMaxRow {
		optimizeResult.IndexedColumns, err = o.sortColumnsByCardinality(ctx, optimizeResult.IndexedColumns)
		if err != nil {
			return nil, err
		}
	}

	return optimizeResult, nil
}

func (o *Optimizer) optimizeJoinTable(tbl string) *OptimizeResult {
	return &OptimizeResult{
		TableName:      tbl,
		IndexedColumns: []string{o.tables[tbl].joinOnColumn},
		Reason:         fmt.Sprintf("字段 %s 为被驱动表 %s 上的关联字段", o.tables[tbl].joinOnColumn, tbl),
	}
}

// doSpecifiedOptimization optimize single table select.
func (o *Optimizer) doSpecifiedOptimization(ctx context.Context, selectStmt *ast.SelectStmt) (*OptimizeResult, error) {
	// todo
	return nil, nil
}

// doGeneralOptimization optimize single table select.
func (o *Optimizer) doGeneralOptimization(ctx context.Context, selectStmt *ast.SelectStmt) (*OptimizeResult, error) {
	generalOptimizer := indexoptimizer.NewOptimizer()

	restoredSQL, err := restoreSelectStmt(selectStmt)
	if err != nil {
		return nil, err
	}

	sa, err := newSelectAST(restoredSQL)
	if err != nil {
		return nil, err
	}

	indexedColumns, err := generalOptimizer.Optimize(sa)
	if err != nil {
		return nil, err
	}

	if len(indexedColumns) == 0 {
		return nil, nil
	}

	o.l.Infof("general optimize result: %v(index columns)", indexedColumns)

	return &OptimizeResult{
		TableName:      getTableNameFromSingleSelect(selectStmt),
		IndexedColumns: indexedColumns,
		Reason:         "三星索引建议",
	}, nil
}

func (o *Optimizer) sortColumnsByCardinality(ctx context.Context, indexedColumns []string) (sortedColumns []string, err error) {
	// todo
	return indexedColumns, nil
}

// needOptimize check table need optimize index of table or not.
//
// Optimize means that:
// 1. When SQL do not use index, we can create index for the select statement.
// 2. When SQL use index, but the index is not suitable, we should optimize the index.
//
// We do it by check MySQL execution plan's access_type field.
// ref: https://dev.mysql.com/doc/refman/5.7/en/explain-output.html#explain-join-types
func (o *Optimizer) needOptimize(record *executor.ExplainRecord) bool {

	// Full table scan: select * from t1 where common_column = 'a'
	// This SQL will scan all rows of table t1.
	if record.Type == executor.ExplainRecordAccessTypeAll {
		return true
	}

	// Index-only scan: select key_part2 from t1 where key_part3 = 'a'
	// This SQL will scan all rows of index idx_composite. It's a little better than previous case.
	if record.Type == executor.ExplainRecordAccessTypeIndex {
		return true
	}

	return false
}

// needIndex check need add index on tbl.columns or not.
func (o *Optimizer) needIndex(tbl string, columns ...string) (bool, error) {
	// todo
	return true, nil
}

type optimizerOption func(*Optimizer)

func (oo optimizerOption) apply(o *Optimizer) {
	oo(o)
}

func WithCalculateCardinalityMaxRow(row int) optimizerOption {
	return func(o *Optimizer) {
		o.calculateCardinalityMaxRow = row
	}
}

func WithCompositeIndexMaxColumn(column int) optimizerOption {
	return func(o *Optimizer) {
		o.compositeIndexMaxColumn = column
	}
}

func WithCreateIndexStatement(f func(tableName string, columns ...string) string) optimizerOption {
	return func(o *Optimizer) {
		o.createIndexStatement = f
	}
}

func defaultCreateIndexStatement(tableName string, columns ...string) string {
	indexName := fmt.Sprintf("idx_%s_%s", tableName, strings.Join(columns, "_"))

	return fmt.Sprintf("CREATE INDEX %s ON %s (%s)",
		indexName,
		tableName,
		strings.Join(columns, ", "))
}

func restoreSelectStmt(ss *ast.SelectStmt) (string, error) {
	var buf strings.Builder
	rc := format.NewRestoreCtx(0, &buf)

	if err := ss.Restore(rc); err != nil {
		return "", errors.Wrap(err, "restore select statement")
	}

	return buf.String(), nil
}

func extractTableNameFromAST(ss *ast.SelectStmt, tbl string) *ast.TableName {
	v := util.TableNameExtractor{TableNames: make(map[string]*ast.TableName)}
	ss.Accept(&v)

	for _, t := range v.TableNames {
		if t.Name.L == tbl {
			return t
		}
	}
	return nil
}

func getTableNameFromSingleSelect(ss *ast.SelectStmt) string {
	if ss.From.TableRefs.Left == nil {
		return ""
	}
	return ss.From.TableRefs.Left.(*ast.TableSource).Source.(*ast.TableName).Name.L
}

// removeDrivingTable remove driving table from execution plan.
//
// Index is not silver bullet, we only give advice on driven table.
// Such as : select * from t1, t2 where t1.id = t2.id;
// There are two records in execution plan, the first one is driving table, the second one is driven table.
func removeDrivingTable(records []*executor.ExplainRecord) []*executor.ExplainRecord {
	var result []*executor.ExplainRecord

	if len(records) == 0 || len(records) == 1 {
		return records
	}

	i, j := 0, 1
	for j < len(records) {
		if records[i].Id == records[j].Id {
			result = append(result, records[j])
		} else {
			i = j
		}
		j++
	}

	return result
}
