package context

import (
	"github.com/actiontech/sqle/sqle/driver/mysql/executor"
	"github.com/actiontech/sqle/sqle/driver/mysql/util"
	"github.com/pingcap/parser/ast"
)

type TableInfo struct {
	Size     float64
	sizeLoad bool

	// isLoad indicate whether TableInfo load from database or not.
	isLoad bool

	// OriginalTable save parser object from db by query "show create table ...";
	// using in inspect and generate rollback sql
	OriginalTable *ast.CreateTableStmt

	//
	MergedTable *ast.CreateTableStmt

	// save alter table parse object from input sql;
	AlterTables []*ast.AlterTableStmt
}

type SchemaInfo struct {
	DefaultEngine    string
	engineLoad       bool
	DefaultCharacter string
	characterLoad    bool
	DefaultCollation string
	collationLoad    bool
	Tables           map[string]*TableInfo
}

type Context struct {
	// currentSchema will change after sql "use database"
	currentSchema string

	schemas map[string]*SchemaInfo
	// if schemas info has collected, set true
	schemaHasLoad bool

	// executionPlan store batch SQLs' execution plan during one inspect context.
	executionPlan map[string][]*executor.ExplainRecord

	// sysVars keep some MySQL global system variables during one inspect context.
	sysVars map[string]string
}

func NewContext(parent *Context) *Context {
	ctx := &Context{
		schemas:       map[string]*SchemaInfo{},
		executionPlan: map[string][]*executor.ExplainRecord{},
		sysVars:       map[string]string{},
	}
	if parent == nil {
		return ctx
	}
	ctx.schemaHasLoad = parent.schemaHasLoad
	ctx.currentSchema = parent.currentSchema
	for schemaName, schema := range parent.schemas {
		newSchema := &SchemaInfo{
			Tables: map[string]*TableInfo{},
		}
		if schema == nil || schema.Tables == nil {
			continue
		}
		for tableName, table := range schema.Tables {
			newSchema.Tables[tableName] = &TableInfo{
				Size:          table.Size,
				sizeLoad:      table.sizeLoad,
				isLoad:        table.isLoad,
				OriginalTable: table.OriginalTable,
				MergedTable:   table.MergedTable,
				AlterTables:   table.AlterTables,
			}
		}
		ctx.schemas[schemaName] = newSchema
	}

	for k, v := range parent.sysVars {
		ctx.sysVars[k] = v
	}
	return ctx
}

func (c *Context) GetSysVar(name string) (string, bool) {
	v, exist := c.sysVars[name]
	return v, exist
}

func (c *Context) AddSysVar(name, value string) {
	c.sysVars[name] = value
	return
}

func (c *Context) HasLoadSchemas() bool {
	return c.schemaHasLoad
}

func (c *Context) SetSchemasLoad() {
	c.schemaHasLoad = true
}

func (c *Context) LoadSchemas(schemas []string) {
	if c.HasLoadSchemas() {
		return
	}
	for _, schema := range schemas {
		c.schemas[schema] = &SchemaInfo{}
	}
	c.SetSchemasLoad()
}

func (c *Context) GetSchema(schemaName string) (*SchemaInfo, bool) {
	schema, has := c.schemas[schemaName]
	return schema, has
}

func (c *Context) HasSchema(schemaName string) (has bool) {
	_, has = c.GetSchema(schemaName)
	return
}

func (c *Context) AddSchema(name string) {
	if c.HasSchema(name) {
		return
	}
	c.schemas[name] = &SchemaInfo{
		Tables: map[string]*TableInfo{},
	}
}

func (c *Context) DelSchema(name string) {
	delete(c.schemas, name)
}

func (c *Context) HasLoadTables(schemaName string) (hasLoad bool) {
	if schema, ok := c.GetSchema(schemaName); ok {
		if schema.Tables == nil {
			hasLoad = false
		} else {
			hasLoad = true
		}
	}
	return
}

func (c *Context) LoadTables(schemaName string, tablesName []string) {
	schema, ok := c.GetSchema(schemaName)
	if !ok {
		return
	}
	if c.HasLoadTables(schemaName) {
		return
	}
	schema.Tables = map[string]*TableInfo{}
	for _, name := range tablesName {
		schema.Tables[name] = &TableInfo{
			isLoad:      true,
			AlterTables: []*ast.AlterTableStmt{},
		}
	}
}

func (c *Context) GetTable(schemaName, tableName string) (*TableInfo, bool) {
	schema, SchemaExist := c.GetSchema(schemaName)
	if !SchemaExist {
		return nil, false
	}
	if !c.HasLoadTables(schemaName) {
		return nil, false
	}
	table, tableExist := schema.Tables[tableName]
	return table, tableExist
}

func (c *Context) HasTable(schemaName, tableName string) (has bool) {
	_, has = c.GetTable(schemaName, tableName)
	return
}

func (c *Context) AddTable(schemaName, tableName string, table *TableInfo) {
	schema, exist := c.GetSchema(schemaName)
	if !exist {
		return
	}
	if !c.HasLoadTables(schemaName) {
		return
	}
	schema.Tables[tableName] = table
}

func (c *Context) DelTable(schemaName, tableName string) {
	schema, exist := c.GetSchema(schemaName)
	if !exist {
		return
	}
	delete(schema.Tables, tableName)
}

func (c *Context) UseSchema(schema string) {
	c.currentSchema = schema
}

func (c *Context) AddExecutionPlan(sql string, records []*executor.ExplainRecord) {
	c.executionPlan[sql] = records
}

func (c *Context) GetExecutionPlan(sql string) ([]*executor.ExplainRecord, bool) {
	records, ok := c.executionPlan[sql]
	return records, ok
}

func (c *Context) UpdateContext(node ast.Node) {
	switch s := node.(type) {
	case *ast.UseStmt:
		// change current schema
		if c.HasSchema(s.DBName) {
			c.UseSchema(s.DBName)
		}
	case *ast.CreateDatabaseStmt:
		if c.HasLoadSchemas() {
			c.AddSchema(s.Name)
		}
	case *ast.CreateTableStmt:
		schemaName := c.GetSchemaName(s.Table)
		tableName := s.Table.Name.L
		if c.HasTable(schemaName, tableName) {
			return
		}
		c.AddTable(schemaName, tableName,
			&TableInfo{
				Size:          0, // table is empty after create
				sizeLoad:      true,
				isLoad:        false,
				OriginalTable: s,
				AlterTables:   []*ast.AlterTableStmt{},
			})
	case *ast.DropDatabaseStmt:
		if c.HasLoadSchemas() {
			c.DelSchema(s.Name)
		}
	case *ast.DropTableStmt:
		if c.HasLoadSchemas() {
			for _, table := range s.Tables {
				schemaName := c.GetSchemaName(table)
				tableName := table.Name.L
				if c.HasTable(schemaName, tableName) {
					c.DelTable(schemaName, tableName)
				}
			}
		}

	case *ast.AlterTableStmt:
		info, exist := c.getTableInfo(s.Table)
		if exist {
			var oldTable *ast.CreateTableStmt
			var err error
			if info.MergedTable != nil {
				oldTable = info.MergedTable
			} else if info.OriginalTable != nil {
				oldTable, err = util.ParseCreateTableStmt(info.OriginalTable.Text())
				if err != nil {
					return
				}
			}
			info.MergedTable, _ = util.MergeAlterToTable(oldTable, s)
			info.AlterTables = append(info.AlterTables, s)
			// rename table
			if s.Table.Name.L != info.MergedTable.Table.Name.L {
				schemaName := c.GetSchemaName(s.Table)
				c.DelTable(schemaName, s.Table.Name.L)
				c.AddTable(schemaName, info.MergedTable.Table.Name.L, info)
			}
		}
	default:
	}
}

// GetSchemaName get schema name from AST or current schema.
func (c *Context) GetSchemaName(stmt *ast.TableName) string {
	if stmt.Schema.String() == "" {
		return c.currentSchema
	}

	return stmt.Schema.String()
}

// TODO: export and remove duplicated code
func (c *Context) getTableInfo(stmt *ast.TableName) (*TableInfo, bool) {
	schema := c.getSchemaName(stmt)
	table := stmt.Name.String()
	return c.GetTable(schema, table)
}
