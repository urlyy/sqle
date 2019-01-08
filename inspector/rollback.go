package inspector

import (
	"database/sql"
	"fmt"
	"github.com/pingcap/tidb/ast"
	_model "github.com/pingcap/tidb/model"
	"sqle/errors"
	"sqle/model"
	"strconv"
	"strings"
)

func (i *Inspect) GenerateAllRollbackSql() ([]*model.RollbackSql, error) {
	i.Logger().Info("start generate rollback sql")

	rollbackSqls := []string{}
	for _, commitSql := range i.Task.CommitSqls {
		err := i.Add(&commitSql.Sql, func(sql *model.Sql) error {
			rollbackSql, err := i.GenerateRollbackSql(sql)
			if rollbackSql != "" {
				rollbackSqls = append(rollbackSqls, rollbackSql)
			}
			return err
		})
		if err != nil {
			i.Logger().Error("add rollback sql failed")
			return nil, err
		}
	}
	if err := i.Do(); err != nil {
		i.Logger().Errorf("generate rollback sql failed")
		return nil, err
	}
	i.Logger().Info("generate rollback sql finish")
	return i.GetAllRollbackSql(rollbackSqls), nil
}

func (i *Inspect) GetAllRollbackSql(sqls []string) []*model.RollbackSql {
	rollbackSqls := []*model.RollbackSql{}
	// Reverse order
	var number uint = 1
	for n := len(sqls) - 1; n >= 0; n-- {
		rollbackSqls = append(rollbackSqls, &model.RollbackSql{
			Sql: model.Sql{
				Number:  number,
				Content: sqls[n],
			},
		})
		number += 1
	}
	return rollbackSqls
}

func (i *Inspect) GenerateRollbackSql(sql *model.Sql) (string, error) {
	node := sql.Stmts[0]
	switch node.(type) {
	case ast.DDLNode:
		return i.GenerateDDLStmtRollbackSql(node)
	case ast.DMLNode:
		return i.GenerateDMLStmtRollbackSql(node)
	}
	return "", nil
}

func (i *Inspect) GenerateDDLStmtRollbackSql(node ast.Node) (rollbackSql string, err error) {
	switch stmt := node.(type) {
	case *ast.AlterTableStmt:
		rollbackSql, err = i.generateAlterTableRollbackSql(stmt)
	case *ast.CreateTableStmt:
		rollbackSql, err = i.generateCreateTableRollbackSql(stmt)
	case *ast.CreateDatabaseStmt:
		rollbackSql, err = i.generateCreateSchemaRollbackSql(stmt)
	case *ast.DropTableStmt:
		rollbackSql, err = i.generateDropTableRollbackSql(stmt)
	case *ast.CreateIndexStmt:
		rollbackSql, err = i.generateCreateIndexRollbackSql(stmt)
	case *ast.DropIndexStmt:
		rollbackSql, err = i.generateDropIndexRollbackSql(stmt)
	}
	return rollbackSql, err
}

func (i *Inspect) GenerateDMLStmtRollbackSql(node ast.Node) (rollbackSql string, err error) {
	if i.config.DMLRollbackMaxRows < 0 {
		return "", nil
	}
	switch stmt := node.(type) {
	case *ast.InsertStmt:
		rollbackSql, err = i.generateInsertRollbackSql(stmt)
	case *ast.DeleteStmt:
		rollbackSql, err = i.generateDeleteRollbackSql(stmt)
	case *ast.UpdateStmt:
		rollbackSql, err = i.generateUpdateRollbackSql(stmt)
	}
	return
}

func (i *Inspect) generateAlterTableRollbackSql(stmt *ast.AlterTableStmt) (string, error) {
	schemaName := i.getSchemaName(stmt.Table)
	tableName := stmt.Table.Name.String()

	createTableStmt, exist, err := i.getCreateTableStmt(stmt.Table)
	if err != nil || !exist {
		return "", err
	}
	rollbackStmt := &ast.AlterTableStmt{
		Table: newTableName(schemaName, tableName),
		Specs: []*ast.AlterTableSpec{},
	}
	// rename table
	if specs := getAlterTableSpecByTp(stmt.Specs, ast.AlterTableRenameTable); len(specs) > 0 {
		spec := specs[len(specs)-1]
		rollbackStmt.Table = newTableName(schemaName, spec.NewTable.Name.String())
		rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
			Tp:       ast.AlterTableRenameTable,
			NewTable: newTableName(schemaName, tableName),
		})
	}
	// add columns need drop columns
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableAddColumns) {
		if spec.NewColumns == nil {
			continue
		}
		for _, col := range spec.NewColumns {
			rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
				Tp:            ast.AlterTableDropColumn,
				OldColumnName: &ast.ColumnName{Name: _model.NewCIStr(col.Name.String())},
			})
		}
	}
	// drop columns need add columns
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableDropColumn) {
		colName := spec.OldColumnName.String()
		for _, col := range createTableStmt.Cols {
			if col.Name.String() == colName {
				rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
					Tp:         ast.AlterTableAddColumns,
					NewColumns: []*ast.ColumnDef{col},
				})
			}
		}
	}
	// change column need change
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableChangeColumn) {
		if spec.NewColumns == nil {
			continue
		}
		for _, col := range createTableStmt.Cols {
			if col.Name.String() == spec.OldColumnName.String() {
				rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
					Tp:            ast.AlterTableChangeColumn,
					OldColumnName: spec.NewColumns[0].Name,
					NewColumns:    []*ast.ColumnDef{col},
				})
			}
		}
	}

	// modify column need modify
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableModifyColumn) {
		if spec.NewColumns == nil {
			continue
		}
		for _, col := range createTableStmt.Cols {
			if col.Name.String() == spec.NewColumns[0].Name.String() {
				rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
					Tp:         ast.AlterTableModifyColumn,
					NewColumns: []*ast.ColumnDef{col},
				})
			}
		}
	}

	/*
		+----------------------------------- alter column -----------------------------------+
		v1 varchar(20) NOT NULL  DEFAULT "test",
			1. alter column v1 set default "TEST" -> alter column v1 set default "test",
			2. alter column v1 drop default -> alter column v1 set default "test",

		v2 varchar(20) NOT NULL,
			1. alter column v1 set default "TEST", -> alter column v1 DROP DEFAULT,
			2. alter column v1 DROP DEFAULT, -> no nothing,
		+------------------------------------------------------------------------------------+
	*/
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableAlterColumn) {
		if spec.NewColumns == nil {
			continue
		}
		newColumn := spec.NewColumns[0]
		newSpec := &ast.AlterTableSpec{
			Tp: ast.AlterTableAlterColumn,
			NewColumns: []*ast.ColumnDef{
				&ast.ColumnDef{
					Name: newColumn.Name,
				},
			},
		}
		for _, col := range createTableStmt.Cols {
			if col.Name.String() == newColumn.Name.String() {
				if HasOneInOptions(col.Options, ast.ColumnOptionDefaultValue) {
					for _, op := range col.Options {
						if op.Tp == ast.ColumnOptionDefaultValue {
							newSpec.NewColumns[0].Options = []*ast.ColumnOption{
								&ast.ColumnOption{
									Expr: op.Expr,
								},
							}
							rollbackStmt.Specs = append(rollbackStmt.Specs, newSpec)
						}
					}
				} else {
					// if *ast.ColumnDef.Options is nil, it is "DROP DEFAULT",
					if newColumn.Options != nil {
						rollbackStmt.Specs = append(rollbackStmt.Specs, newSpec)
					} else {
						// do nothing
					}
				}
			}
		}
	}
	// drop index need add
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableDropIndex) {
		for _, constraint := range createTableStmt.Constraints {
			if constraint.Name == spec.Name {
				rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
					Tp:         ast.AlterTableAddConstraint,
					Constraint: constraint,
				})
			}
		}
	}
	// drop primary key need add
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableDropPrimaryKey) {
		_ = spec
		for _, constraint := range createTableStmt.Constraints {
			if constraint.Tp == ast.ConstraintPrimaryKey {
				rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
					Tp:         ast.AlterTableAddConstraint,
					Constraint: constraint,
				})
			}
		}
	}

	// drop foreign key need add
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableDropForeignKey) {
		for _, constraint := range createTableStmt.Constraints {
			if constraint.Name == spec.Name {
				rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
					Tp:         ast.AlterTableAddConstraint,
					Constraint: constraint,
				})
			}
		}
	}

	// rename index
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableRenameIndex) {
		spec.FromKey, spec.ToKey = spec.ToKey, spec.FromKey
		rollbackStmt.Specs = append(rollbackStmt.Specs, spec)
	}

	// add constraint (index key, primary key ...) need drop
	for _, spec := range getAlterTableSpecByTp(stmt.Specs, ast.AlterTableAddConstraint) {
		switch spec.Constraint.Tp {
		case ast.ConstraintIndex, ast.ConstraintUniq:
			// add index without index name, index name will be created by db
			if spec.Constraint.Name == "" {
				continue
			}
			rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
				Tp:   ast.AlterTableDropIndex,
				Name: spec.Constraint.Name,
			})
		case ast.ConstraintPrimaryKey:
			rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
				Tp: ast.AlterTableDropPrimaryKey,
			})
		case ast.ConstraintForeignKey:
			rollbackStmt.Specs = append(rollbackStmt.Specs, &ast.AlterTableSpec{
				Tp:   ast.AlterTableDropForeignKey,
				Name: spec.Constraint.Name,
			})
		}
	}

	rollbackSql := alterTableStmtFormat(rollbackStmt)
	return rollbackSql, nil
}

func (i *Inspect) generateCreateSchemaRollbackSql(stmt *ast.CreateDatabaseStmt) (string, error) {
	schemaName := stmt.Name
	schemaExist, err := i.isSchemaExist(schemaName)
	if err != nil {
		return "", err
	}
	if schemaExist {
		return "", err
	}
	rollbackSql := fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", schemaName)
	return rollbackSql, nil
}

func (i *Inspect) generateCreateTableRollbackSql(stmt *ast.CreateTableStmt) (string, error) {
	schemaExist, err := i.isSchemaExist(i.getSchemaName(stmt.Table))
	if err != nil {
		return "", err
	}
	// if schema not exist, create table will be failed. don't rollback
	if !schemaExist {
		return "", nil
	}

	tableExist, err := i.isTableExist(stmt.Table)
	if err != nil {
		return "", err
	}

	if tableExist {
		return "", nil
	}
	rollbackSql := fmt.Sprintf("DROP TABLE IF EXISTS %s", i.getTableNameWithQuote(stmt.Table))
	return rollbackSql, nil
}

func (i *Inspect) generateDropTableRollbackSql(stmt *ast.DropTableStmt) (string, error) {
	rollbackSql := ""
	for _, table := range stmt.Tables {
		stmt, tableExist, err := i.getCreateTableStmt(table)
		if err != nil {
			return "", err
		}
		// if table not exist, can not rollback it.
		if !tableExist {
			continue
		}
		rollbackSql += stmt.Text() + ";\n"
	}
	return rollbackSql, nil
}

func (i *Inspect) generateCreateIndexRollbackSql(stmt *ast.CreateIndexStmt) (string, error) {
	return fmt.Sprintf("DROP INDEX `%s` ON %s", stmt.IndexName, i.getTableNameWithQuote(stmt.Table)), nil
}

func (i *Inspect) generateDropIndexRollbackSql(stmt *ast.DropIndexStmt) (string, error) {
	indexName := stmt.IndexName
	createTableStmt, tableExist, err := i.getCreateTableStmt(stmt.Table)
	if err != nil {
		return "", err
	}
	// if table not exist, don't rollback
	if !tableExist {
		return "", nil
	}
	rollbackSql := ""
	for _, constraint := range createTableStmt.Constraints {
		if constraint.Name == indexName {
			sql := ""
			switch constraint.Tp {
			case ast.ConstraintIndex:
				sql = fmt.Sprintf("CREATE INDEX `%s` ON %s",
					indexName, i.getTableNameWithQuote(stmt.Table))
			case ast.ConstraintUniq:
				sql = fmt.Sprintf("CREATE UNIQUE INDEX `%s` ON %s",
					indexName, i.getTableNameWithQuote(stmt.Table))
			default:
				return "", nil
			}
			if constraint.Option != nil {
				sql = fmt.Sprintf("%s %s", sql, indexOptionFormat(constraint.Option))
			}
			rollbackSql = sql
		}
	}
	return rollbackSql, nil
}

func (i *Inspect) generateInsertRollbackSql(stmt *ast.InsertStmt) (string, error) {
	tables := getTables(stmt.Table.TableRefs)
	// table just has one in insert stmt.
	if len(tables) != 1 {
		return "", nil
	}
	if stmt.OnDuplicate != nil {
		return "", nil
	}
	table := tables[0]
	createTableStmt, exist, err := i.getCreateTableStmt(table)
	if err != nil {
		return "", err
	}
	// if table not exist, insert will failed.
	if !exist {
		return "", nil
	}
	pkColumnsName, hasPk, err := i.getPrimaryKey(createTableStmt)
	if err != nil || !hasPk {
		return "", nil
	}

	rollbackSql := ""

	// match "insert into table_name (column_name,...) value (v1,...)"
	// match "insert into table_name value (v1,...)"
	if stmt.Lists != nil {
		if int64(len(stmt.Lists)) > i.config.DMLRollbackMaxRows {
			return "", nil
		}
		columnsName := []string{}
		if stmt.Columns != nil {
			for _, col := range stmt.Columns {
				columnsName = append(columnsName, col.Name.String())
			}
		} else {
			for _, col := range createTableStmt.Cols {
				columnsName = append(columnsName, col.Name.String())
			}
		}
		for _, value := range stmt.Lists {
			where := []string{}
			// mysql will throw error: 1136 (21S01): Column count doesn't match value count
			if len(columnsName) != len(value) {
				return "", nil
			}
			for n, name := range columnsName {
				_, isPk := pkColumnsName[name]
				if isPk {
					where = append(where, fmt.Sprintf("%s = '%s'", name, exprFormat(value[n])))
				}
			}
			if len(where) != len(pkColumnsName) {
				return "", nil
			}
			rollbackSql += fmt.Sprintf("DELETE FROM %s WHERE %s;\n",
				i.getTableNameWithQuote(table), strings.Join(where, " AND "))
		}
		return rollbackSql, nil
	}

	// match "insert into table_name set col_name = value1, ..."
	if stmt.Setlist != nil {
		if 1 > i.config.DMLRollbackMaxRows {
			return "", nil
		}
		where := []string{}
		for _, setExpr := range stmt.Setlist {
			name := setExpr.Column.Name.String()
			_, isPk := pkColumnsName[name]
			if isPk {
				where = append(where, fmt.Sprintf("%s = '%s'", name, exprFormat(setExpr.Expr)))
			}
		}
		if len(where) != len(pkColumnsName) {
			return "", nil
		}
		rollbackSql = fmt.Sprintf("DELETE FROM %s WHERE %s;\n",
			i.getTableNameWithQuote(table), strings.Join(where, " AND "))
	}
	return rollbackSql, nil
}

func (i *Inspect) generateDeleteRollbackSql(stmt *ast.DeleteStmt) (string, error) {
	// not support multi-table syntax
	if stmt.IsMultiTable {
		return "", nil
	}
	var err error
	tables := getTables(stmt.TableRefs.TableRefs)
	table := tables[0]
	createTableStmt, exist, err := i.getCreateTableStmt(table)
	if err != nil || !exist {
		return "", err
	}
	_, hasPk, err := i.getPrimaryKey(createTableStmt)
	if err != nil || !hasPk {
		return "", nil
	}

	var max = i.config.DMLRollbackMaxRows
	limit, err := getLimitCount(stmt.Limit, max+1)
	if err != nil {
		return "", err
	}
	if limit > max {
		count, err := i.getRecordCount(table, stmt.Where, stmt.Order, limit)
		if err != nil {
			return "", err
		}
		if count > max {
			return "", nil
		}
	}
	records, err := i.getRecords(table, stmt.Where, stmt.Order, limit)

	values := []string{}

	columnsName := []string{}
	for _, col := range createTableStmt.Cols {
		columnsName = append(columnsName, col.Name.Name.String())
	}
	for _, record := range records {
		if len(record) != len(columnsName) {
			return "", nil
		}
		vs := []string{}
		for _, name := range columnsName {
			v := "NULL"
			if record[name].Valid {
				v = fmt.Sprintf("'%s'", record[name].String)
			}
			vs = append(vs, v)
		}
		values = append(values, fmt.Sprintf("(%s)", strings.Join(vs, ", ")))
	}
	rollbackSql := ""
	if len(values) > 0 {
		rollbackSql = fmt.Sprintf("INSERT INTO %s (`%s`) VALUES %s;",
			i.getTableNameWithQuote(table), strings.Join(columnsName, "`, `"),
			strings.Join(values, ", "))
	}
	return rollbackSql, nil
}

func (i *Inspect) generateUpdateRollbackSql(stmt *ast.UpdateStmt) (string, error) {
	tables := getTables(stmt.TableRefs.TableRefs)
	// multi table syntax
	if len(tables) != 1 {
		return "", nil
	}
	table := tables[0]
	createTableStmt, exist, err := i.getCreateTableStmt(table)
	if err != nil || !exist {
		return "", err
	}
	pkColumnsName, hasPk, err := i.getPrimaryKey(createTableStmt)
	if err != nil || !hasPk {
		return "", nil
	}

	var max = i.config.DMLRollbackMaxRows
	limit, err := getLimitCount(stmt.Limit, max+1)
	if err != nil {
		return "", err
	}
	if limit > max {
		count, err := i.getRecordCount(table, stmt.Where, stmt.Order, limit)
		if err != nil {
			return "", err
		}
		if count > max {
			return "", nil
		}
	}
	records, err := i.getRecords(table, stmt.Where, stmt.Order, limit)

	columnsName := []string{}
	rollbackSql := ""
	for _, col := range createTableStmt.Cols {
		columnsName = append(columnsName, col.Name.Name.String())
	}
	for _, record := range records {
		if len(record) != len(columnsName) {
			return "", nil
		}
		where := []string{}
		value := []string{}
		for _, col := range createTableStmt.Cols {
			colChanged := false
			_, isPk := pkColumnsName[col.Name.Name.L]
			isPkChanged := false
			pkValue := ""

			for _, l := range stmt.List {
				if col.Name.Name.L == l.Column.Name.L {
					colChanged = true
					if isPk {
						isPkChanged = true
						pkValue = exprFormat(l.Expr)
					}
				}
			}
			name := col.Name.String()
			v := "NULL"
			if record[name].Valid {
				v = fmt.Sprintf("'%s'", record[name].String)
			}

			if colChanged {
				value = append(value, fmt.Sprintf("%s = %s", name, v))
			}
			if isPk {
				if isPkChanged {
					where = append(where, fmt.Sprintf("%s = '%s'", name, pkValue))
				} else {
					where = append(where, fmt.Sprintf("%s = %s", name, v))

				}
			}
		}
		rollbackSql += fmt.Sprintf("UPDATE %s SET %s WHERE %s;", i.getTableNameWithQuote(table),
			strings.Join(value, ", "), strings.Join(where, " AND "))
	}
	return rollbackSql, nil
}

func (i *Inspect) getRecords(tableName *ast.TableName, where ast.ExprNode,
	order *ast.OrderByClause, limit int64) ([]map[string]sql.NullString, error) {
	conn, err := i.getDbConn()
	if err != nil {
		return nil, err
	}
	sql := i.generateGetRecordsSql("*", tableName, where, order, limit)
	return conn.Db.Query(sql)
}

func (i *Inspect) getRecordCount(tableName *ast.TableName, where ast.ExprNode,
	order *ast.OrderByClause, limit int64) (int64, error) {
	conn, err := i.getDbConn()
	if err != nil {
		return 0, err
	}
	sql := i.generateGetRecordsSql("count(*) as count", tableName, where, order, limit)

	var count int64
	var ok bool
	records, err := conn.Db.Query(sql)
	if err != nil {
		return 0, err
	}
	if len(records) != 1 {
		goto ERROR
	}
	_, ok = records[0]["count"]
	if !ok {
		goto ERROR
	}
	count, err = strconv.ParseInt(records[0]["count"].String, 10, 64)
	if err != nil {
		goto ERROR
	}
	return count, nil

ERROR:
	return 0, errors.New(errors.CONNECT_REMOTE_DB_ERROR, fmt.Errorf("do not match records for select count(*)"))
}

func (i *Inspect) generateGetRecordsSql(expr string, tableName *ast.TableName, where ast.ExprNode,
	order *ast.OrderByClause, limit int64) string {
	recordSql := fmt.Sprintf("SELECT %s FROM %s", expr, getTableNameWithQuote(tableName))
	if where != nil {
		recordSql = fmt.Sprintf("%s WHERE %s", recordSql, exprFormat(where))
	}
	if order != nil {
		recordSql = fmt.Sprintf("%s ORDER BY", recordSql)
		for _, item := range order.Items {
			recordSql = fmt.Sprintf("%s %s", recordSql, exprFormat(item.Expr))
			if item.Desc {
				recordSql = fmt.Sprintf("%s DESC", recordSql)
			}
		}
	}
	if limit > 0 {
		recordSql = fmt.Sprintf("%s LIMIT %d", recordSql, limit)
	}
	recordSql += ";"
	return recordSql
}
