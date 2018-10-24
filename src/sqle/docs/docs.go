// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-24 14:30:23.198319954 +0800 CST m=+0.113385573

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server for dev.",
        "title": "Sqle API Docs",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/",
    "paths": {
        "/instances": {
            "get": {
                "description": "get all instances",
                "summary": "实例列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetAllInstReq"
                        }
                    }
                }
            },
            "post": {
                "description": "create a instance\u003cbr\u003ecreate a instance",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "summary": "添加实例",
                "parameters": [
                    {
                        "description": "add instance",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateInstReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateInstRes"
                        }
                    }
                }
            }
        },
        "/instances/{inst_id}/": {
            "put": {
                "description": "update instance db",
                "summary": "更新实例",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "inst_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update instance",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateInstReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete instance db",
                "summary": "删除实例",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "inst_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/instances/{instance_id}/connection": {
            "get": {
                "description": "test instance db connection",
                "summary": "实例连通性测试",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instance_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.PingInstRes"
                        }
                    }
                }
            }
        },
        "/instances/{instance_id}/schemas": {
            "get": {
                "description": "instance schema list",
                "summary": "实例 Schema 列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instance_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetSchemaRes"
                        }
                    }
                }
            }
        },
        "/rule_templates": {
            "get": {
                "description": "get all rule template",
                "summary": "规则模板列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetAllTplRes"
                        }
                    }
                }
            },
            "post": {
                "description": "create a rule template",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "summary": "添加规则模板",
                "parameters": [
                    {
                        "description": "add instance",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateTplReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/rule_templates/template_id/": {
            "put": {
                "description": "update rule template",
                "summary": "更新规则模板",
                "parameters": [
                    {
                        "description": "update rule template",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateTplReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete rule template",
                "summary": "删除规则模板",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/rule_templates/{template_id}/": {
            "get": {
                "description": "get rule template",
                "summary": "获取规则模板",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rule Template ID",
                        "name": "template_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetAllRuleRes"
                        }
                    }
                }
            }
        },
        "/rules": {
            "get": {
                "description": "get all rule template",
                "summary": "规则列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetAllRuleRes"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "get all tasks",
                "summary": "Sql审核列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetAllTaskRes"
                        }
                    }
                }
            },
            "post": {
                "description": "create a task",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "summary": "创建Sql审核单",
                "parameters": [
                    {
                        "description": "add task",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateTaskReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.CreateTaskRes"
                        }
                    }
                }
            }
        },
        "/tasks/{task_id}": {
            "get": {
                "description": "get task",
                "summary": "获取Sql审核单信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.GetTaskReq"
                        }
                    }
                }
            }
        },
        "/tasks/{task_id}/commit": {
            "post": {
                "description": "commit sql",
                "summary": "Sql提交上线",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/tasks/{task_id}/inspection": {
            "post": {
                "description": "inspect sql",
                "summary": "Sql提交审核",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/tasks/{task_id}/rollback": {
            "post": {
                "description": "rollback sql",
                "summary": "Sql提交回滚",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "manual execute rollback sql",
                        "name": "manual_execute",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.BaseRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.CreateInstReq": {
            "type": "object",
            "properties": {
                "db_type": {
                    "description": "mysql, mycat, sqlserver",
                    "type": "string",
                    "example": "mysql"
                },
                "desc": {
                    "type": "string",
                    "example": "this is a test instance"
                },
                "host": {
                    "type": "string",
                    "example": "10.10.10.10"
                },
                "name": {
                    "type": "string",
                    "example": "test"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "port": {
                    "type": "string",
                    "example": "3306"
                },
                "rule_templates": {
                    "description": "this a list for rule template name",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "template_1"
                    ]
                },
                "user": {
                    "type": "string",
                    "example": "root"
                }
            }
        },
        "controller.CreateInstRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/model.Instance"
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.CreateTaskReq": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string",
                    "example": "this is a test task"
                },
                "inst_name": {
                    "type": "string",
                    "example": "inst_1"
                },
                "name": {
                    "type": "string",
                    "example": "test"
                },
                "schema": {
                    "type": "string",
                    "example": "db1"
                },
                "sql": {
                    "type": "string",
                    "example": "alter table tb1 drop columns c1"
                }
            }
        },
        "controller.CreateTaskRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/model.Task"
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.CreateTplReq": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rule_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "ddl_create_table_not_exist"
                    ]
                }
            }
        },
        "controller.GetAllInstReq": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Instance"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.GetAllRuleRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Rule"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.GetAllTaskRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Task"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.GetAllTplRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RuleTemplate"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.GetSchemaRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "db1"
                    ]
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.GetTaskReq": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/model.Task"
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controller.PingInstRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "description": "true: ping success; false: ping failed",
                    "type": "boolean"
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "model.CommitSql": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string",
                    "example": "2018-10-21T16:40:23+08:00"
                },
                "exec_result": {
                    "type": "string"
                },
                "exec_status": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "inspect_level": {
                    "type": "string"
                },
                "inspect_result": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "sql": {
                    "type": "string"
                }
            }
        },
        "model.Instance": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string",
                    "example": "2018-10-21T16:40:23+08:00"
                },
                "db_type": {
                    "type": "string",
                    "example": "mysql"
                },
                "desc": {
                    "type": "string",
                    "example": "this is a instance"
                },
                "host": {
                    "type": "string",
                    "example": "10.10.10.10"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "string",
                    "example": "3306"
                },
                "user": {
                    "type": "string",
                    "example": "root"
                }
            }
        },
        "model.RollbackSql": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string",
                    "example": "2018-10-21T16:40:23+08:00"
                },
                "exec_result": {
                    "type": "string"
                },
                "exec_status": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "number": {
                    "type": "integer"
                },
                "sql": {
                    "type": "string"
                }
            }
        },
        "model.Rule": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string"
                },
                "level": {
                    "description": "notice, warn, error",
                    "type": "string",
                    "example": "error"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.RuleTemplate": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string",
                    "example": "2018-10-21T16:40:23+08:00"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string"
                },
                "rules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Rule"
                    }
                }
            }
        },
        "model.Sql": {
            "type": "object",
            "properties": {
                "commit_sqls": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CommitSql"
                    }
                },
                "create_at": {
                    "type": "string",
                    "example": "2018-10-21T16:40:23+08:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "progress": {
                    "type": "string"
                },
                "rollback_sqls": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RollbackSql"
                    }
                },
                "sql": {
                    "type": "string"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string",
                    "example": "2018-10-21T16:40:23+08:00"
                },
                "desc": {
                    "type": "string",
                    "example": "this is a task"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "instance": {
                    "type": "object",
                    "$ref": "#/definitions/model.Instance"
                },
                "name": {
                    "type": "string",
                    "example": "REQ201812578"
                },
                "schema": {
                    "type": "string",
                    "example": "db1"
                },
                "sql": {
                    "type": "object",
                    "$ref": "#/definitions/model.Sql"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
