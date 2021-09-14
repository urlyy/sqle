package api

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"actiontech.cloud/sqle/sqle/sqle/api/controller"
	v1 "actiontech.cloud/sqle/sqle/sqle/api/controller/v1"
	sqleMiddleware "actiontech.cloud/sqle/sqle/sqle/api/middleware"
	"actiontech.cloud/sqle/sqle/sqle/config"
	_ "actiontech.cloud/sqle/sqle/sqle/docs"
	"actiontech.cloud/sqle/sqle/sqle/errors"
	"actiontech.cloud/sqle/sqle/sqle/log"
	"actiontech.cloud/sqle/sqle/sqle/model"
	"actiontech.cloud/sqle/sqle/sqle/utils"

	"github.com/facebookgo/grace/gracenet"
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
func StartApi(net *gracenet.Net, exitChan chan struct{}, config config.SqleConfig) {
	defer close(exitChan)

	e := echo.New()
	output := log.NewRotateFile(config.LogPath, "/api.log", 1024 /*1GB*/)
	defer output.Close()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: output,
	}))
	e.HideBanner = true
	e.HidePort = true

	// custom handler http error
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if _, ok := err.(*errors.CodeError); ok {
			controller.JSONBaseErrorReq(c, err)
		} else {
			e.DefaultHTTPErrorHandler(err, c)
		}
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/v1/login", v1.Login)

	v1Router := e.Group("/v1")
	v1Router.Use(sqleMiddleware.JWTTokenAdapter(), middleware.JWT([]byte(utils.JWTSecret)))

	// v1 admin api, just admin user can access.
	{
		// user
		v1Router.GET("/users", v1.GetUsers, AdminUserAllowed())
		v1Router.POST("/users", v1.CreateUser, AdminUserAllowed())
		v1Router.GET("/users/:user_name/", v1.GetUser, AdminUserAllowed())
		v1Router.PATCH("/users/:user_name/", v1.UpdateUser, AdminUserAllowed())
		v1Router.DELETE("/users/:user_name/", v1.DeleteUser, AdminUserAllowed())
		v1Router.PATCH("/users/:user_name/password", v1.UpdateOtherUserPassword, AdminUserAllowed())

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
		v1Router.POST("/rule_templates/:rule_template_name/clone", v1.CloneRuleTemplate, AdminUserAllowed())
		v1Router.PATCH("/rule_templates/:rule_template_name/", v1.UpdateRuleTemplate, AdminUserAllowed())
		v1Router.DELETE("/rule_templates/:rule_template_name/", v1.DeleteRuleTemplate, AdminUserAllowed())

		// workflow template
		v1Router.GET("/workflow_templates", v1.GetWorkflowTemplates, AdminUserAllowed())
		v1Router.POST("/workflow_templates", v1.CreateWorkflowTemplate, AdminUserAllowed())
		v1Router.GET("/workflow_templates/:workflow_template_name/", v1.GetWorkflowTemplate, AdminUserAllowed())
		v1Router.PATCH("/workflow_templates/:workflow_template_name/", v1.UpdateWorkflowTemplate, AdminUserAllowed())
		v1Router.DELETE("/workflow_templates/:workflow_template_name/", v1.DeleteWorkflowTemplate, AdminUserAllowed())
		v1Router.GET("/workflow_template_tips", v1.GetWorkflowTemplateTips, AdminUserAllowed())

		// workflow
		v1Router.POST("/workflows/cancel", v1.BatchCancelWorkflows, AdminUserAllowed())

		// audit whitelist
		v1Router.GET("/audit_whitelist", v1.GetSqlWhitelist, AdminUserAllowed())
		v1Router.POST("/audit_whitelist", v1.CreateAuditWhitelist, AdminUserAllowed())
		v1Router.PATCH("/audit_whitelist/:audit_whitelist_id/", v1.UpdateAuditWhitelistById, AdminUserAllowed())
		v1Router.DELETE("/audit_whitelist/:audit_whitelist_id/", v1.DeleteAuditWhitelistById, AdminUserAllowed())

		// configurations
		v1Router.GET("/configurations/smtp", v1.GetSMTPConfiguration, AdminUserAllowed())
		v1Router.PATCH("/configurations/smtp", v1.UpdateSMTPConfiguration, AdminUserAllowed())
		v1Router.GET("/configurations/system_variables", v1.GetSystemVariables, AdminUserAllowed())
		v1Router.PATCH("/configurations/system_variables", v1.UpdateSystemVariables, AdminUserAllowed())
	}

	// user
	v1Router.GET("/user", v1.GetCurrentUser)
	v1Router.PATCH("/user", v1.UpdateCurrentUser)
	v1Router.GET("/user_tips", v1.GetUserTips)
	v1Router.PUT("/user/password", v1.UpdateCurrentUserPassword)

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

	// configurations
	v1Router.GET("/configurations/drivers", v1.GetDrivers)

	// audit plan
	v1Router.POST("/audit_plans", v1.CreateAuditPlan)
	v1Router.DELETE("/audit_plans/:audit_plan_name/", v1.DeleteAuditPlan)
	v1Router.PATCH("/audit_plans/:audit_plan_name/", v1.UpdateAuditPlan)
	v1Router.GET("/audit_plans/:audit_plan_name/", v1.GetAuditPlan)
	v1Router.GET("/audit_plans", v1.GetAuditPlans)
	v1Router.GET("/audit_plans/:audit_plan_name/reports", v1.GetAuditPlanReports)
	v1Router.GET("/audit_plans/:audit_plan_name/report/:audit_plan_report_id/", v1.GetAuditPlanReportSQLs)
	v1Router.GET("/audit_plans/:audit_plan_name/sqls", v1.GetAuditPlanSQLs)
	v1Router.POST("/audit_plans/:audit_plan_name/sqls/full", v1.FullSyncAuditPlanSQLs, sqleMiddleware.AuditPlanVerifyAdapter())
	v1Router.POST("/audit_plans/:audit_plan_name/sqls/partial", v1.PartialSyncAuditPlanSQLs, sqleMiddleware.AuditPlanVerifyAdapter())
	v1Router.POST("/audit_plans/:audit_plan_name/trigger", v1.TriggerAuditPlan)

	// UI
	e.File("/", "ui/index.html")
	e.Static("/static", "ui/static")
	e.File("/favicon.png", "ui/favicon.png")
	e.GET("/*", func(c echo.Context) error {
		return c.File("ui/index.html")
	})

	address := fmt.Sprintf(":%v", config.SqleServerPort)
	log.Logger().Infof("starting http server on %s", address)

	// start http server
	l, err := net.Listen("tcp4", address)
	if err != nil {
		log.Logger().Fatal(err)
		return
	}
	if config.EnableHttps {
		// Usually, it is easier to create an tls server using echo#StartTLS;
		// but I need create a graceful listener.
		if config.CertFilePath == "" || config.KeyFilePath == "" {
			log.Logger().Fatal("invalid tls configuration")
			return
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = make([]tls.Certificate, 1)
		tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(config.CertFilePath, config.KeyFilePath)
		e.TLSServer.TLSConfig = tlsConfig
		e.TLSListener = tls.NewListener(l, tlsConfig)

		log.Logger().Fatal(e.StartServer(e.TLSServer))
	} else {
		e.Listener = l
		log.Logger().Fatal(e.Start(""))
	}
	return
}

// AdminUserAllowed is a `echo` middleware, only allow admin user to access next.
func AdminUserAllowed() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if controller.GetUserName(c) == model.DefaultAdminUser {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden)
		}
	}
}
