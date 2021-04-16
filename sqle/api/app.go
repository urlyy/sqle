package api

import (
	"actiontech.cloud/universe/sqle/v4/sqle/api/controller"
	"actiontech.cloud/universe/sqle/v4/sqle/api/controller/v1"
	"net/http"
	"strings"

	"fmt"

	_ "actiontech.cloud/universe/sqle/v4/sqle/docs"
	"actiontech.cloud/universe/sqle/v4/sqle/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Sqle API Docs
// @version 1.0
// @description This is a sample server for dev.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func StartApi(port int, exitChan chan struct{}, logPath string) {
	e := echo.New()
	output := log.NewRotateFile(logPath, "/api.log", 1024 /*1GB*/)
	defer output.Close()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: output,
	}))
	e.HideBanner = true
	e.HidePort = true

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/v1/login", v1.Login)

	v1Router := e.Group("/v1")
	v1Router.Use(JWTTokenAdapter(), middleware.JWT([]byte(v1.JWTSecret)))

	// v1 admin api, just admin user can access.
	{
		// user
		v1Router.GET("/users", v1.GetUsers, AdminUserAllowed())
		v1Router.POST("/users", v1.CreateUser, AdminUserAllowed())
		v1Router.GET("/users/:user_name/", v1.GetUser, AdminUserAllowed())
		v1Router.PATCH("/users/:user_name/", v1.UpdateUser, AdminUserAllowed())
		v1Router.DELETE("/users/:user_name/", v1.DeleteUser, AdminUserAllowed())

		// role
		v1Router.GET("/roles", v1.GetRoles, AdminUserAllowed())
		v1Router.GET("/role_tips", v1.GetRoleTips, AdminUserAllowed())
		v1Router.POST("/roles", v1.CreateRole, AdminUserAllowed())
		v1Router.PATCH("/roles/:role_name/", v1.UpdateRole, AdminUserAllowed())
		v1Router.DELETE("/roles/:role_name/", v1.DeleteRole, AdminUserAllowed())

		// instance
		v1Router.POST("/instances", v1.CreateInstance, AdminUserAllowed())
		v1Router.DELETE("/instances/:instance_name/", v1.DeleteInstance, AdminUserAllowed())
		v1Router.PATCH("/instances/:instance_name/", v1.UpdateInstance, AdminUserAllowed())

		// rule template
		v1Router.POST("/rule_templates", v1.CreateRuleTemplate, AdminUserAllowed())
		v1Router.PATCH("/rule_templates/:rule_template_name/", v1.UpdateRuleTemplate, AdminUserAllowed())
		v1Router.DELETE("/rule_templates/:rule_template_name/", v1.DeleteRuleTemplate, AdminUserAllowed())

		// workflow template
		v1Router.GET("/workflow_templates", v1.GetWorkflowTemplates, AdminUserAllowed())
		v1Router.POST("/workflow_templates", v1.CreateWorkflowTemplate, AdminUserAllowed())
		v1Router.GET("/workflow_templates/:workflow_template_name/", v1.GetWorkflowTemplate, AdminUserAllowed())
		v1Router.PATCH("/workflow_templates/:workflow_template_name/", v1.UpdateWorkflowTemplate, AdminUserAllowed())
		v1Router.DELETE("/workflow_templates/:workflow_template_name/", v1.DeleteWorkflowTemplate, AdminUserAllowed())
		v1Router.GET("/workflow_template_tips", v1.GetWorkflowTemplateTips, AdminUserAllowed())

		// audit whitelist
		v1Router.GET("/audit_whitelist", v1.GetSqlWhitelist, AdminUserAllowed())
		v1Router.POST("/audit_whitelist", v1.CreateAuditWhitelist, AdminUserAllowed())
		v1Router.PATCH("/audit_whitelist/:audit_whitelist_id/", v1.UpdateAuditWhitelistById, AdminUserAllowed())
		v1Router.DELETE("/audit_whitelist/:audit_whitelist_id/", v1.DeleteAuditWhitelistById, AdminUserAllowed())
	}

	// user
	v1Router.GET("/user", v1.GetCurrentUser)
	v1Router.PATCH("/user", v1.UpdateCurrentUser)
	v1Router.GET("/user_tips", v1.GetUserTips)

	// instance
	v1Router.GET("/instances", v1.GetInstances)
	v1Router.GET("/instances/:instance_name/", v1.GetInstance)
	v1Router.GET("/instances/:instance_name/connection", v1.CheckInstanceIsConnectableByName)
	v1Router.POST("/instance_connection", v1.CheckInstanceIsConnectable)
	v1Router.GET("/instances/:instance_name/schemas", v1.GetInstanceSchemas)
	v1Router.GET("/instance_tips", v1.GetInstanceTips)
	v1Router.GET("/instances/:instance_name/rules", v1.GetInstanceRules)

	// rule template
	v1Router.GET("/rule_templates", v1.GetRuleTemplates)
	v1Router.GET("/rule_template_tips", v1.GetRuleTemplateTips)
	v1Router.GET("/rule_templates/:rule_template_name/", v1.GetRuleTemplate)

	//rule
	v1Router.GET("/rules", v1.GetRules)

	// workflow
	v1Router.POST("/workflows", v1.CreateWorkflow)
	v1Router.GET("/workflows/:workflow_id/", v1.GetWorkflow)
	v1Router.GET("/workflows", v1.GetWorkflows)
	v1Router.POST("/workflows/:workflow_id/steps/:workflow_step_id/approve", v1.ApproveWorkflow)
	v1Router.POST("/workflows/:workflow_id/steps/:workflow_step_id/reject", v1.RejectWorkflow)
	v1Router.POST("/workflows/:workflow_id/cancel", v1.CancelWorkflow)
	v1Router.PATCH("/workflows/:workflow_id/", v1.UpdateWorkflow)

	// task
	v1Router.POST("/tasks/audits", v1.CreateAndAuditTask)
	v1Router.GET("/tasks/audits/:task_id/", v1.GetTask)
	v1Router.GET("/tasks/audits/:task_id/sqls", v1.GetTaskSQLs)
	v1Router.GET("/tasks/audits/:task_id/sql_report", v1.DownloadTaskSQLReportFile)
	v1Router.GET("/tasks/audits/:task_id/sql_file", v1.DownloadTaskSQLFile)
	v1Router.GET("/tasks/audits/:task_id/sql_content", v1.GetAuditTaskSQLContent)

	// dashboard
	v1Router.GET("/dashboard", v1.Dashboard)

	// UI
	e.File("/", "ui/index.html")
	e.Static("/static", "ui/static")
	e.File("/favicon.png", "ui/favicon.png")
	e.GET("/*", func(c echo.Context) error {
		return c.File("ui/index.html")
	})

	address := fmt.Sprintf(":%v", port)
	log.Logger().Infof("starting http server on %s", address)
	log.Logger().Fatal(e.Start(address))
	close(exitChan)
}

// JWTTokenAdapter is a `echo` middleware,　by rewriting the header, the jwt token support header
// "Authorization: {token}" and "Authorization: Bearer {token}".
func JWTTokenAdapter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			if auth != "" && !strings.HasPrefix(auth, middleware.DefaultJWTConfig.AuthScheme) {
				c.Request().Header.Set(echo.HeaderAuthorization,
					fmt.Sprintf("%s %s", middleware.DefaultJWTConfig.AuthScheme, auth))
			}
			return next(c)
		}
	}
}

// AdminUserAllowed is a `echo` middleware, only allow admin user to access next.
func AdminUserAllowed() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if controller.GetUserName(c) == "admin" {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden)
		}
	}
}
