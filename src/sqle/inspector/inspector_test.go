package inspector

import (
	"github.com/pingcap/parser"
	"sqle/model"
	"testing"
)

func DefaultMysqlInspect() *Inspector {
	return &Inspector{
		Config: map[string]*model.Rule{},
		Db: model.Instance{
			DbType: "mysql",
		},
		SqlArray:      []*model.CommitSql{},
		currentSchema: "exist_db",
		allSchema:     map[string]struct{}{"exist_db": struct{}{}},
		schemaHasLoad: true,
		allTable: map[string]map[string]struct{}{
			"exist_db": map[string]struct{}{
				"exist_tb_1": struct{}{},
				"exist_tb_2": struct{}{},
			}},
	}
}

func runInspectCase(t *testing.T, desc string, i *Inspector, sql string, results ...*InspectResults) {
	stmts, err := parseSql(i.Db.DbType, sql)
	if err != nil {
		t.Errorf("%s test failled, error: %v\n", desc, err)
		return
	}
	for n, stmt := range stmts {
		i.SqlArray = append(i.SqlArray, &model.CommitSql{
			Number: n + 1,
			Sql:    stmt.Text(),
		})
	}
	_, err = i.Inspect()
	if err != nil {
		t.Errorf("%s test failled, error: %v\n", desc, err)
		return
	}
	if len(i.SqlArray) != len(results) {
		t.Errorf("%s test failled, error: result is unknow\n", desc)
		return
	}
	for n, sql := range i.SqlArray {
		result := results[n]
		if sql.InspectLevel != result.level() || sql.InspectResult != result.message() {
			t.Errorf("%s test failled, \nsql: %s\nexpect level: %s\nexpect result:\n%s\nactual level: %s\nactual result:\n%s\n",
				desc, sql.Sql, result.level(), result.message(), sql.InspectLevel, sql.InspectResult)
		}
	}
}

func TestInspectMysql(t *testing.T) {
	runInspectCase(t, "use database exist", DefaultMysqlInspect(),
		"use exist_db",
		newInspectResults(),
	)
	runInspectCase(t, "use database not exist", DefaultMysqlInspect(),
		"use no_exist_db",
		newInspectResults(
			&Result{
				Level:   model.RULE_LEVEL_ERROR,
				Message: "database no_exist_db not exist",
			}),
	)
	baseCreateQuery := `ALTER TABLE t1.a1 ADD COLUMN v3 varchar(255) DEFAULT NULL;
`

	p := parser.New()
	_, err := p.Parse(baseCreateQuery, "", "")
	if err != nil {
		t.Error(err)
		return
	}
}
