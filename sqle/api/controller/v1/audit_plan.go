package v1

import (
	"context"
	"fmt"
	"net/http"

	"actiontech.cloud/sqle/sqle/sqle/api/controller"
	"actiontech.cloud/sqle/sqle/sqle/driver"
	"actiontech.cloud/sqle/sqle/sqle/errors"
	"actiontech.cloud/sqle/sqle/sqle/log"
	"actiontech.cloud/sqle/sqle/sqle/model"

	"github.com/labstack/echo/v4"
	"github.com/ungerik/go-dry"
)

var (
	errAuditPlanNotExist         = errors.New(errors.DataNotExist, fmt.Errorf("audit plan is not exist"))
	errAuditPlanInstanceConflict = errors.New(errors.DataConflict, fmt.Errorf("instance_name can not be empty while instance_database is not empty"))
)

type CreateAuditPlanReqV1 struct {
	Name             string `json:"audit_plan_name" form:"audit_plan_name" example:"audit_plan_for_java_repo_1" valid:"required,name"`
	Cron             string `json:"audit_plan_cron" form:"audit_plan_cron" example:"0 */2 * * *" valid:"required,cron"`
	InstanceType     string `json:"audit_plan_instance_type" form:"audit_plan_instance_type" example:"mysql" valid:"required"`
	InstanceName     string `json:"audit_plan_instance_name" form:"audit_plan_instance_name" example:"test_mysql"`
	InstanceDatabase string `json:"audit_plan_instance_database" form:"audit_plan_instance_database" example:"app1"`
}

// @Summary 添加审核计划
// @Description create audit plan
// @Id createAuditPlanV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Accept json
// @Param audit_plan body v1.CreateAuditPlanReqV1 true "create audit plan"
// @Success 200 {object} controller.BaseRes
// @router /v1/audit_plans [post]
func CreateAuditPlan(c echo.Context) error {
	s := model.GetStorage()

	req := new(CreateAuditPlanReqV1)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	if !dry.StringInSlice(req.InstanceType, driver.AllDrivers()) {
		return controller.JSONBaseErrorReq(c, errors.New(errors.DriverNotExist, &driver.ErrDriverNotSupported{DriverTyp: req.InstanceType}))
	}

	if req.InstanceDatabase != "" && req.InstanceName == "" {
		return controller.JSONBaseErrorReq(c, errAuditPlanInstanceConflict)
	}

	_, exist, err := s.GetAuditPlanByName(req.Name)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if exist {
		return controller.JSONBaseErrorReq(c, errAuditPlanNotExist)
	}

	if req.InstanceName != "" {
		instance, exist, err := s.GetInstanceByName(req.InstanceName)
		if err != nil {
			return controller.JSONBaseErrorReq(c, err)
		}
		if !exist {
			return controller.JSONBaseErrorReq(c, errors.New(errors.DataNotExist, instanceNotExistError))
		}

		if req.InstanceDatabase != "" {
			d, err := driver.NewDriver(log.NewEntry(), instance, "")
			if err != nil {
				return controller.JSONBaseErrorReq(c, err)
			}
			schemas, err := d.Schemas(context.TODO())
			if err != nil {
				return controller.JSONBaseErrorReq(c, err)
			}
			if !dry.StringInSlice(req.InstanceDatabase, schemas) {
				return controller.JSONBaseErrorReq(c, errors.New(errors.DataNotExist, fmt.Errorf("database %v is not exist in instance", req.InstanceDatabase)))
			}
			d.Close(context.TODO())
		}
	}

	// todo trigger memory create
	// todo generate token before save audit plan

	return controller.JSONBaseErrorReq(c,
		s.Save(&model.AuditPlan{
			Name:             req.Name,
			Cron:             req.Cron,
			DBType:           req.InstanceType,
			InstanceName:     req.InstanceName,
			InstanceDatabase: req.InstanceDatabase,
		}))
}

// @Summary 删除审核计划
// @Description delete audit plan
// @Id deleteAuditPlanV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Success 200 {object} controller.BaseRes
// @router /v1/audit_plans/{audit_plan_name}/ [delete]
func DeleteAuditPlan(c echo.Context) error {
	s := model.GetStorage()

	auditPlan, exist, err := s.GetAuditPlanByName(c.Param("audit_plan_name"))
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if !exist {
		return controller.JSONBaseErrorReq(c, errAuditPlanNotExist)
	}

	// todo trigger memory delete

	return controller.JSONBaseErrorReq(c, s.Delete(auditPlan))
}

type UpdateAuditPlanReqV1 struct {
	Cron             *string `json:"audit_plan_cron" form:"audit_plan_cron" example:"0 */2 * * *" valid:"omitempty,cron"`
	InstanceName     *string `json:"audit_plan_instance_name" form:"audit_plan_instance_name" example:"test_mysql"`
	InstanceDatabase *string `json:"audit_plan_instance_database" form:"audit_plan_instance_database" example:"app1"`
}

// @Summary 更新审核计划
// @Description update audit plan
// @Id updateAuditPlanV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @param audit_plan body v1.UpdateAuditPlanReqV1 true "update audit plan"
// @Success 200 {object} controller.BaseRes
// @router /v1/audit_plans/{audit_plan_name}/ [patch]
func UpdateAuditPlan(c echo.Context) error {
	s := model.GetStorage()

	req := new(UpdateAuditPlanReqV1)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return err
	}

	auditPlan, exist, err := s.GetAuditPlanByName(c.Param("audit_plan_name"))
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if !exist {
		return controller.JSONBaseErrorReq(c, errAuditPlanNotExist)
	}

	if req.Cron != nil {
		auditPlan.Cron = *req.Cron
	}
	if req.InstanceName != nil {
		if auditPlan.InstanceDatabase != "" && *req.InstanceName == "" {
			return controller.JSONBaseErrorReq(c, errAuditPlanInstanceConflict)
		}
		auditPlan.InstanceName = *req.InstanceName
	}
	if req.InstanceDatabase != nil {
		if auditPlan.InstanceName == "" && *req.InstanceDatabase != "" {
			return controller.JSONBaseErrorReq(c, errAuditPlanInstanceConflict)
		}
		auditPlan.InstanceDatabase = *req.InstanceDatabase
	}

	// todo trigger memory update

	return controller.JSONBaseErrorReq(c, s.Save(auditPlan))
}

type GetAuditPlansReqV1 struct {
	FilterAuditPlanDBType string `json:"filter_audit_plan_db_type" query:"filter_audit_plan_db_type"`
	PageIndex             uint32 `json:"page_index" query:"page_index" valid:"required"`
	PageSize              uint32 `json:"page_size" query:"page_size" valid:"required"`
}

type GetAuditPlansResV1 struct {
	controller.BaseRes
	Data      []AuditPlanResV1 `json:"data"`
	TotalNums uint64           `json:"total_nums"`
}

type AuditPlanResV1 struct {
	Name             string `json:"audit_plan_name" example:"audit_for_java_app1"`
	Cron             string `json:"audit_plan_cron" example:"0 */2 * * *"`
	DBType           string `json:"audit_plan_db_type" example:"mysql"`
	Token            string `json:"audit_plan_token" example:"it's a JWT Token for scanner"`
	InstanceName     string `json:"audit_plan_instance_name" example:"test_mysql"`
	InstanceDatabase string `json:"audit_plan_instance_database" example:"app1"`
}

// @Summary 获取审核计划信息列表
// @Description get audit plan info list
// @Id getAuditPlansV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param filter_audit_plan_db_type query string false "filter audit plan db type"
// @Param page_index query uint32 false "page index"
// @Param page_size query uint32 false "size of per page"
// @Success 200 {object} v1.GetAuditPlansResV1
// @router /v1/audit_plans [get]
func GetAuditPlans(c echo.Context) error {
	s := model.GetStorage()

	req := new(GetAuditPlansReqV1)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return err
	}

	// todo refactor to api common utils
	var offset uint32
	if req.PageIndex >= 1 {
		offset = req.PageSize * (req.PageIndex - 1)
	}

	data := map[string]interface{}{
		"filter_audit_plan_db_type": req.FilterAuditPlanDBType,
		"limit":                     req.PageSize,
		"offset":                    offset,
	}
	auditPlans, count, err := s.GetAuditPlansByReq(data)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	var auditPlansResV1 []AuditPlanResV1
	for _, auditPlan := range auditPlans {
		auditPlansResV1 = append(auditPlansResV1, AuditPlanResV1{
			Name:             auditPlan.Name,
			Cron:             auditPlan.Cron,
			DBType:           auditPlan.DBType,
			InstanceName:     auditPlan.InstanceName.String,
			InstanceDatabase: auditPlan.InstanceDatabase.String,

			Token: auditPlan.Token,
		})
	}
	return c.JSON(http.StatusOK, &GetAuditPlansResV1{
		BaseRes:   controller.NewBaseReq(nil),
		Data:      auditPlansResV1,
		TotalNums: count,
	})
}

type GetAuditPlanReportsReqV1 struct {
	PageIndex uint32 `json:"page_index" query:"page_index" valid:"required"`
	PageSize  uint32 `json:"page_size" query:"page_size" valid:"required"`
}

type GetAuditPlanReportsResV1 struct {
	controller.BaseRes
	Data      []AuditPlanReportResV1 `json:"data"`
	TotalNums uint64                 `json:"total_nums"`
}

type AuditPlanReportResV1 struct {
	Id        string `json:"audit_plan_report_id" example:"1"`
	Timestamp string `json:"audit_plan_report_timestamp" example:"RFC3339"`
}

// @Summary 获取指定审核计划的报告列表
// @Description get audit plan report list
// @Id getAuditPlanReportsV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param page_index query uint32 false "page index"
// @Param page_size query uint32 false "size of per page"
// @Success 200 {object} v1.GetAuditPlanReportsResV1
// @router /v1/audit_plans/{audit_plan_name}/reports [get]
func GetAuditPlanReports(c echo.Context) error {
	s := model.GetStorage()

	req := new(GetAuditPlanReportsReqV1)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return err
	}

	// todo refactor to api common utils
	var offset uint32
	if req.PageIndex >= 1 {
		offset = req.PageSize * (req.PageIndex - 1)
	}

	data := map[string]interface{}{
		"audit_plan_name": c.Param("audit_plan_name"),
		"limit":           req.PageSize,
		"offset":          offset,
	}
	auditPlanReports, count, err := s.GetAuditPlanReportsByReq(data)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	var auditPlanReportsResV1 []AuditPlanReportResV1
	for _, auditPlanReport := range auditPlanReports {
		auditPlanReportsResV1 = append(auditPlanReportsResV1, AuditPlanReportResV1{
			Id:        auditPlanReport.ID,
			Timestamp: auditPlanReport.CreateAt,
		})
	}
	return c.JSON(http.StatusOK, &GetAuditPlanReportsResV1{
		BaseRes:   controller.NewBaseReq(nil),
		Data:      auditPlanReportsResV1,
		TotalNums: count,
	})
}

type GetAuditPlanReportReqV1 struct {
	PageIndex uint32 `json:"page_index" query:"page_index" valid:"required"`
	PageSize  uint32 `json:"page_size" query:"page_size" valid:"required"`
}

type GetAuditPlanReportResV1 struct {
	controller.BaseRes
	Data      []AuditPlanReportDetailResV1 `json:"data"`
	TotalNums uint64                       `json:"total_nums"`
}

type AuditPlanReportDetailResV1 struct {
	SQLFingerprint     string `json:"audit_plan_report_sql_fingerprint" example:"select * from t1 where id = ?"`
	SQLLastReceiveText string `json:"audit_plan_report_sql_last_receive_text" example:"select * from t1 where id = 1"`
	AuditResult        string `json:"audit_plan_report_audit_result" example:"same format as task audit result"`
	Timestamp          string `json:"audit_plan_report_timestamp" example:"RFC3339"`
}

// @Summary 获取指定审核计划的报告详情
// @Description get audit plan report detail
// @Id getAuditPlanReportDetailV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param page_index query uint32 false "page index"
// @Param page_size query uint32 false "size of per page"
// @Success 200 {object} v1.GetAuditPlanReportResV1
// @router /v1/audit_plans/{audit_plan_name}/report/{audit_plan_report_id}/ [get]
func GetAuditPlanReportDetail(c echo.Context) error { return nil }

type FullSyncAuditPlanSQLsReqV1 struct {
	SQLs []AuditPlanSQLReqV1 `json:"audit_plan_sql_list" form:"audit_plan_sql_list"`
}

type AuditPlanSQLReqV1 struct {
	SQLFingerprint          string `json:"audit_plan_sql_fingerprint" form:"audit_plan_sql_fingerprint" example:"select * from t1 where id = ?" valid:"required"`
	SQLCounter              string `json:"audit_plan_sql_counter" form:"audit_plan_sql_counter" example:"6" valid:"required"`
	SQLLastReceiveText      string `json:"audit_plan_sql_last_receive_text" form:"audit_plan_sql_last_receive_text" example:"select * from t1 where id = 1"`
	SQLLastReceiveTimestamp string `json:"audit_plan_sql_last_receive_timestamp" form:"audit_plan_sql_last_receive_timestamp" example:"RFC3339"`
}

// @Summary 全量同步SQL到审核计划
// @Description full sync audit plan SQLs
// @Id fullSyncAuditPlanSQLsV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param sqls body v1.FullSyncAuditPlanSQLsReqV1 true "full sync audit plan SQLs request"
// @Success 200 {object} controller.BaseRes
// @router /v1/audit_plans/{audit_plan_name}/sqls/full [post]
func FullSyncAuditPlanSQLs(c echo.Context) error { return nil }

type PartialSyncAuditPlanSQLsReqV1 struct {
	SQLs []AuditPlanSQLReqV1 `json:"audit_plan_sql_list" form:"audit_plan_sql_list"`
}

// @Summary 增量同步SQL到审核计划
// @Description partial sync audit plan SQLs
// @Id partialSyncAuditPlanSQLsV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param sqls body v1.PartialSyncAuditPlanSQLsReqV1 true "partial sync audit plan SQLs request"
// @Success 200 {object} controller.BaseRes
// @router /v1/audit_plans/{audit_plan_name}/sqls/partial [post]
func PartialSyncAuditPlanSQLs(c echo.Context) error { return nil }

type GetAuditPlanSQLsReqV1 struct {
	PageIndex uint32 `json:"page_index" query:"page_index" valid:"required"`
	PageSize  uint32 `json:"page_size" query:"page_size" valid:"required"`
}

type GetAuditPlanSQLsResV1 struct {
	controller.BaseRes
	Data      []AuditPlanSQLResV1 `json:"data"`
	TotalNums uint64              `json:"total_nums"`
}

type AuditPlanSQLResV1 struct {
	SQLFingerprint          string `json:"audit_plan_sql_fingerprint" example:"select * from t1 where id = ?"`
	SQLCounter              string `json:"audit_plan_sql_counter" example:"6"`
	SQLLastReceiveText      string `json:"audit_plan_sql_last_receive_text" example:"select * from t1 where id = 1"`
	SQLLastReceiveTimestamp string `json:"audit_plan_sql_last_receive_timestamp" example:"RFC3339"`
}

// @Summary 获取指定审核计划的SQLs信息
// @Description get audit plan SQLs
// @Id getAuditPlanSQLsV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Param page_index query uint32 false "page index"
// @Param page_size query uint32 false "size of per page"
// @Success 200 {object} v1.GetAuditPlanSQLsResV1
// @router /v1/audit_plans/{audit_plan_name}/sqls [get]
func GetAuditPlanSQLs(c echo.Context) error { return nil }

// @Summary 触发审核计划
// @Description trigger audit plan
// @Id triggerAuditPlanV1
// @Tags audit_plan
// @Security ApiKeyAuth
// @Param audit_plan_name path string true "audit plan name"
// @Success 200 {object} controller.BaseRes
// @router /v1/audit_plans/{audit_plan_name}/trigger [post]
func TriggerAuditPlan(c echo.Context) error { return nil }
