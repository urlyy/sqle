package v1

import (
	"fmt"
	"net/http"

	parser "github.com/actiontech/mybatis-mapper-2-sql"
	"github.com/actiontech/sqle/sqle/api/controller"
	"github.com/actiontech/sqle/sqle/errors"
	"github.com/actiontech/sqle/sqle/log"
	"github.com/actiontech/sqle/sqle/model"
	"github.com/actiontech/sqle/sqle/server"

	"github.com/labstack/echo/v4"
)

type DirectAuditReqV1 struct {
	InstanceType string `json:"instance_type" form:"instance_type" example:"MySQL" valid:"required"`
	// 调用方不应该关心SQL是否被完美的拆分成独立的条目, 拆分SQL由SQLE实现
	SQLContent   string  `json:"sql_content" form:"sql_content" example:"select * from t1; select * from t2;" valid:"required"`
	SQLType      string  `json:"sql_type" form:"sql_type" example:"sql" enums:"sql,mybatis," valid:"omitempty,oneof=sql mybatis"`
	ProjectName  *string `json:"project_name" form:"project_name" example:"project1"`
	InstanceName *string `json:"instance_name" form:"instance_name" example:"instance1"`
	SchemaName   *string `json:"schema_name" form:"schema_name" example:"schema1"`
}

const (
	SQLTypeSQL     = "sql" // DirectAuditReqV1.SQLType 为空时默认为此
	SQLTypeMyBatis = "mybatis"
)

type DirectAuditResV1 struct {
	controller.BaseRes
	Data *AuditResDataV1 `json:"data"`
}

type AuditResDataV1 struct {
	AuditLevel string          `json:"audit_level" enums:"normal,notice,warn,error,"`
	Score      int32           `json:"score"`
	PassRate   float64         `json:"pass_rate"`
	SQLResults []AuditSQLResV1 `json:"sql_results"`
}

type AuditSQLResV1 struct {
	Number      uint   `json:"number"`
	ExecSQL     string `json:"exec_sql"`
	AuditResult string `json:"audit_result"`
	AuditLevel  string `json:"audit_level"`
}

var ErrDirectAudit = errors.New(errors.GenericError, fmt.Errorf("audit failed, please confirm whether the type of audit plugin supports static audit, please check the log for details"))

// @Summary 直接审核SQL
// @Description Direct audit sql
// @Id directAuditV1
// @Tags sql_audit
// @Security ApiKeyAuth
// @Param req body v1.DirectAuditReqV1 true "sqls that should be audited"
// @Success 200 {object} v1.DirectAuditResV1
// @router /v1/sql_audit [post]
func DirectAudit(c echo.Context) error {
	req := new(DirectAuditReqV1)
	err := controller.BindAndValidateReq(c, req)
	if err != nil {
		return err
	}

	sql := req.SQLContent
	if req.SQLType == SQLTypeMyBatis {
		sql, err = parser.ParseXML(req.SQLContent)
		if err != nil {
			return controller.JSONBaseErrorReq(c, err)
		}
	}

	l := log.NewEntry().WithField("/v1/sql_audit", "direct audit failed")

	s := model.GetStorage()
	var instance *model.Instance
	var exist bool
	if req.ProjectName != nil && req.InstanceName != nil {
		instance, exist, err = s.GetInstanceByNameAndProjectName(*req.InstanceName, *req.ProjectName)
		if err != nil {
			return controller.JSONBaseErrorReq(c, err)
		}
		if !exist {
			return controller.JSONBaseErrorReq(c, ErrInstanceNotExist)
		}
	}

	var schemaName string
	if req.SchemaName != nil {
		schemaName = *req.SchemaName
	}

	var task *model.Task
	if instance != nil && schemaName != "" {
		task, err = server.DirectAuditByInstance(l, sql, schemaName, instance)
	} else {
		task, err = server.AuditSQLByDBType(l, sql, req.InstanceType, nil, "")
	}
	if err != nil {
		l.Errorf("audit sqls failed: %v", err)
		return controller.JSONBaseErrorReq(c, ErrDirectAudit)
	}

	return c.JSON(http.StatusOK, DirectAuditResV1{
		BaseRes: controller.BaseRes{},
		Data:    convertTaskResultToAuditResV1(task),
	})
}

func convertTaskResultToAuditResV1(task *model.Task) *AuditResDataV1 {
	results := make([]AuditSQLResV1, len(task.ExecuteSQLs))
	for i, sql := range task.ExecuteSQLs {
		results[i] = AuditSQLResV1{
			Number:      sql.Number,
			ExecSQL:     sql.Content,
			AuditResult: sql.GetAuditResults(),
			AuditLevel:  sql.AuditLevel,
		}
	}
	return &AuditResDataV1{
		AuditLevel: task.AuditLevel,
		Score:      task.Score,
		PassRate:   task.PassRate,
		SQLResults: results,
	}
}

type GetSQLAnalysisReq struct {
	ProjectName  string `json:"project_name" query:"project_name" example:"default" valid:"required"`
	InstanceName string `json:"instance_name" query:"instance_name" example:"MySQL" valid:"required"`
	SchemaName   string `json:"schema_name" query:"schema_name" example:"test"`
	Sql          string `json:"sql" query:"sql" example:"select * from t1; select * from t2;"`
}

type DirectGetSQLAnalysisResV1 struct {
	controller.BaseRes
	Data []*SqlAnalysisResDataV1 `json:"data"`
}

type SqlAnalysisResDataV1 struct {
	SQLExplain SQLExplain  `json:"sql_explain"`
	TableMetas []TableMeta `json:"table_metas"`
}

// DirectGetSQLAnalysis
// @Summary 直接获取SQL分析结果
// @Description Direct get sql analysis result
// @Id directGetSQLAnalysisV1
// @Tags sql_analysis
// @Param project_name query string true "project name"
// @Param instance_name query string true "instance name"
// @Param schema_name query string false "schema name"
// @Param sql query string false "sql"
// @Security ApiKeyAuth
// @Success 200 {object} v1.DirectGetSQLAnalysisResV1
// @router /v1/sql_analysis [get]
func DirectGetSQLAnalysis(c echo.Context) error {
	return directGetSQLAnalysis(c)
}
