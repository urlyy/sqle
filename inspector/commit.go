package inspector

import (
	"database/sql/driver"
	"github.com/labstack/gommon/log"
	"github.com/pingcap/tidb/ast"
	"sqle/model"
)

func (i *Inspect) CommitAll() error {
	defer i.closeDbConn()
	for _, commitSql := range i.Task.CommitSqls {
		currentSql := commitSql
		err := i.Add(&currentSql.Sql, func(sql *model.Sql) error {
			err := i.Commit(sql)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return i.Do()
}

func (i *Inspect) RollbackAll(sql *model.RollbackSql) error {
	for _, rollbackSql := range i.Task.RollbackSqls {
		currentSql := rollbackSql
		err := i.Add(&currentSql.Sql, func(sql *model.Sql) error {
			err := i.Commit(sql)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return i.Do()
}

func (i *Inspect) Commit(sql *model.Sql) error {
	if i.isDMLStmt {
		return i.commitDML(sql)
	} else {
		return i.commitDDL(sql)
	}
}

func (i *Inspect) commitDDL(sql *model.Sql) error {
	conn, err := i.getDbConn()
	if err != nil {
		return err
	}
	if i.Task.Instance.DbType == model.DB_TYPE_MYCAT {
		return i.commitMycatDDL(sql)
	}
	//sql := i.GetCommitSql()
	_, err = conn.Db.Exec(sql.Content)
	if err != nil {
		sql.ExecStatus = model.TASK_ACTION_ERROR
		sql.ExecResult = err.Error()
	} else {
		sql.ExecStatus = model.TASK_ACTION_DONE
		sql.ExecResult = "ok"
	}
	return nil
}

func (i *Inspect) commitDML(sql *model.Sql) error {
	var err error
	var result driver.Result
	var rowAffect int64
	var qs []string

	conn, err := i.getDbConn()
	if err != nil {
		return err
	}

	sql.StartBinlogFile, sql.StartBinlogPos, err = conn.FetchMasterBinlogPos()
	if err != nil {
		goto ERROR
	}

	qs = make([]string, 0, len(sql.Stmts))
	for _, stmt := range sql.Stmts {
		qs = append(qs, stmt.Text())
	}

	if len(qs) > 1 && i.Task.Instance.DbType != model.DB_TYPE_MYCAT {
		err = conn.Db.Transact(qs...)
		if err != nil {
			goto ERROR
		}
	} else {
		for _, query := range qs {
			result, err = conn.Db.Exec(query)
			if err != nil {
				goto ERROR
			}
			rowAffect, err = result.RowsAffected()
			if err != nil {
				log.Warnf("get rows affect failed, error: %s", err)
			} else {
				sql.RowAffects += rowAffect
			}
		}
	}

	sql.ExecStatus = model.TASK_ACTION_DONE
	sql.ExecResult = "ok"
	// if sql has commit success, ignore error for check status.
	sql.EndBinlogFile, sql.EndBinlogPos, _ = conn.FetchMasterBinlogPos()
	return nil
ERROR:
	sql.ExecStatus = model.TASK_ACTION_ERROR
	sql.ExecResult = err.Error()
	return err
}

func (i *Inspect) commitMycatDDL(sql *model.Sql) error {
	conn, err := i.getDbConn()
	if err != nil {
		return err
	}
	var schema string
	var table string

	switch stmt := sql.Stmts[0].(type) {
	case *ast.CreateTableStmt:
		schema = i.getSchemaName(stmt.Table)
		table = stmt.Table.Name.String()
	case *ast.AlterTableStmt:
		schema = i.getSchemaName(stmt.Table)
		table = stmt.Table.Name.String()
	case *ast.CreateIndexStmt:
		schema = i.getSchemaName(stmt.Table)
		table = stmt.Table.Name.String()
	case *ast.UseStmt:
		goto DONE
	default:
	}
	err = conn.Db.ExecDDL(ReplaceTableName(sql.Stmts[0]), schema, table)

DONE:
	if err != nil {
		sql.ExecStatus = model.TASK_ACTION_ERROR
		sql.ExecResult = err.Error()
	} else {
		sql.ExecStatus = model.TASK_ACTION_DONE
		sql.ExecResult = "ok"
	}
	return nil
}
