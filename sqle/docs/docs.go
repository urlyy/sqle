// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/login": {
            "post": {
                "description": "user login",
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "operationId": "loginV1",
                "parameters": [
                    {
                        "description": "user login request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserLoginReqV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetUserLoginResV1"
                        }
                    }
                }
            }
        },
        "/v1/role_tips": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get role tip list",
                "tags": [
                    "role"
                ],
                "summary": "获取角色提示列表",
                "operationId": "getRoleTipListV1",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetRoleTipsResV1"
                        }
                    }
                }
            }
        },
        "/v1/roles": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get role list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "获取角色列表",
                "operationId": "getRoleListV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filter role name",
                        "name": "filter_role_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter user name",
                        "name": "filter_user_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter instance name",
                        "name": "filter_instance_name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page index",
                        "name": "page_index",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size of per page",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetRolesResV1"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "创建角色",
                "operationId": "createRoleV1",
                "parameters": [
                    {
                        "description": "create role",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CreateRoleReqV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/v1/roles/{role_name}/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "删除角色",
                "operationId": "deleteRoleV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "更新角色信息",
                "operationId": "updateRoleV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update role request",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateRoleReqV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/v1/test": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "user login",
                "tags": [
                    "user"
                ],
                "summary": "test",
                "operationId": "testV1",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user_tips": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get user tip list",
                "tags": [
                    "user"
                ],
                "summary": "获取用户提示列表",
                "operationId": "getUserTipListV1",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetUserTipsResV1"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get user info list",
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息列表",
                "operationId": "getUserListV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filter user name",
                        "name": "filter_user_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter role name",
                        "name": "filter_role_name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page index",
                        "name": "page_index",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size of per page",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetUsersResV1"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户",
                "operationId": "createUserV1",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CreateUserReqV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            }
        },
        "/v1/users/{user_name}/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get user info",
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息",
                "operationId": "getUserV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "user_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.UserResV1"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete user",
                "tags": [
                    "user"
                ],
                "summary": "删除用户",
                "operationId": "deleteUserV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "user_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseRes"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "更新用户信息",
                "operationId": "updateUserV1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "user_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update user",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateUserReqV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
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
        "v1.CreateRoleReqV1": {
            "type": "object",
            "properties": {
                "instance_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role_desc": {
                    "type": "string"
                },
                "role_name": {
                    "type": "string"
                },
                "user_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.CreateUserReqV1": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@email.com"
                },
                "role_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_name": {
                    "type": "string",
                    "example": "test"
                },
                "user_password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "v1.GetRoleTipsResV1": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.RoleTipResV1"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "v1.GetRolesResV1": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.RoleResV1"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                },
                "total_nums": {
                    "type": "integer"
                }
            }
        },
        "v1.GetUserLoginResV1": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/v1.UserLoginResV1"
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "v1.GetUserTipsResV1": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.UserTipResV1"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "v1.GetUsersResV1": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.UserResV1"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                },
                "total_nums": {
                    "type": "integer"
                }
            }
        },
        "v1.RoleResV1": {
            "type": "object",
            "properties": {
                "instance_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role_desc": {
                    "type": "string"
                },
                "role_name": {
                    "type": "string"
                },
                "user_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.RoleTipResV1": {
            "type": "object",
            "properties": {
                "role_name": {
                    "type": "string"
                }
            }
        },
        "v1.UpdateRoleReqV1": {
            "type": "object",
            "properties": {
                "instance_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role_desc": {
                    "type": "string"
                },
                "user_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.UpdateUserReqV1": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "role_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.UserLoginReqV1": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "v1.UserLoginResV1": {
            "type": "object",
            "properties": {
                "is_admin": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string",
                    "example": "this is a jwt token string"
                }
            }
        },
        "v1.UserResV1": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "role_name_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "v1.UserTipResV1": {
            "type": "object",
            "properties": {
                "user_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Sqle API Docs",
	Description: "This is a sample server for dev.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
