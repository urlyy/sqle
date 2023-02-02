package v2

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/actiontech/sqle/sqle/errors"
	"github.com/actiontech/sqle/sqle/notification"
	"github.com/actiontech/sqle/sqle/pkg/im"
	"github.com/actiontech/sqle/sqle/utils"

	"github.com/actiontech/sqle/sqle/model"

	v1 "github.com/actiontech/sqle/sqle/api/controller/v1"

	"github.com/actiontech/sqle/sqle/api/controller"
	"github.com/labstack/echo/v4"
)

var errTaskHasBeenUsed = errors.New(errors.DataConflict, fmt.Errorf("task has been used in other workflow"))

type WorkflowStepResV2 struct {
	Id            uint       `json:"workflow_step_id,omitempty"`
	Number        uint       `json:"number"`
	Type          string     `json:"type" enums:"create_workflow,update_workflow,sql_review,sql_execute"`
	Desc          string     `json:"desc,omitempty"`
	Users         []string   `json:"assignee_user_name_list,omitempty"`
	OperationUser string     `json:"operation_user_name,omitempty"`
	OperationTime *time.Time `json:"operation_time,omitempty"`
	State         string     `json:"state,omitempty" enums:"initialized,approved,rejected"`
	Reason        string     `json:"reason,omitempty"`
}

// @Summary 审批通过
// @Description approve workflow
// @Tags workflow
// @Id approveWorkflowV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param workflow_step_id path string true "workflow step id"
// @Param project_name path string true "project name"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/steps/{workflow_step_id}/approve [post]
func ApproveWorkflowV2(c echo.Context) error {
	return nil
}

type RejectWorkflowReqV2 struct {
	Reason string `json:"reason" form:"reason"`
}

// @Summary 审批驳回
// @Description reject workflow
// @Tags workflow
// @Id rejectWorkflowV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param project_name path string true "project name"
// @Param workflow_step_id path string true "workflow step id"
// @param workflow_approve body v2.RejectWorkflowReqV2 true "workflow approve request"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/steps/{workflow_step_id}/reject [post]
func RejectWorkflowV2(c echo.Context) error {
	return nil
}

// @Summary 审批关闭（中止）
// @Description cancel workflow
// @Tags workflow
// @Id cancelWorkflowV2
// @Security ApiKeyAuth
// @Param project_name path string true "project name"
// @Param workflow_id path string true "workflow id"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/cancel [post]
func CancelWorkflowV2(c echo.Context) error {
	return nil
}

type BatchCancelWorkflowsReqV2 struct {
	WorkflowIDList []string `json:"workflow_id_list" form:"workflow_id_list"`
}

// BatchCancelWorkflowsV2 batch cancel workflows.
// @Summary 批量取消工单
// @Description batch cancel workflows
// @Tags workflow
// @Id batchCancelWorkflowsV2
// @Security ApiKeyAuth
// @Param project_name path string true "project name"
// @Param BatchCancelWorkflowsReqV2 body v2.BatchCancelWorkflowsReqV2 true "batch cancel workflows request"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/cancel [post]
func BatchCancelWorkflowsV2(c echo.Context) error {
	return nil
}

type BatchCompleteWorkflowsReqV2 struct {
	WorkflowIDList []string `json:"workflow_id_list" form:"workflow_id_list"`
}

// BatchCompleteWorkflowsV2 complete workflows.
// @Summary 批量完成工单
// @Description this api will directly change the work order status to finished without real online operation
// @Tags workflow
// @Id batchCompleteWorkflowsV2
// @Security ApiKeyAuth
// @Param project_name path string true "project name"
// @Param data body v2.BatchCompleteWorkflowsReqV2 true "batch complete workflows request"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/complete [post]
func BatchCompleteWorkflowsV2(c echo.Context) error {
	return nil
}

// ExecuteOneTaskOnWorkflowV2
// @Summary 工单提交单个数据源上线
// @Description execute one task on workflow
// @Tags workflow
// @Id executeOneTaskOnWorkflowV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param project_name path string true "project name"
// @Param task_id path string true "task id"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/tasks/{task_id}/execute [post]
func ExecuteOneTaskOnWorkflowV2(c echo.Context) error {
	return nil
}

type GetWorkflowTasksResV2 struct {
	controller.BaseRes
	Data []*GetWorkflowTasksItemV2 `json:"data"`
}

type GetWorkflowTasksItemV2 struct {
	TaskId                   uint                       `json:"task_id"`
	InstanceName             string                     `json:"instance_name"`
	Status                   string                     `json:"status" enums:"wait_for_audit,wait_for_execution,exec_scheduled,exec_failed,exec_succeeded,executing,manually_executed"`
	ExecStartTime            *time.Time                 `json:"exec_start_time,omitempty"`
	ExecEndTime              *time.Time                 `json:"exec_end_time,omitempty"`
	ScheduleTime             *time.Time                 `json:"schedule_time,omitempty"`
	CurrentStepAssigneeUser  []string                   `json:"current_step_assignee_user_name_list,omitempty"`
	TaskPassRate             float64                    `json:"task_pass_rate"`
	TaskScore                int32                      `json:"task_score"`
	InstanceMaintenanceTimes []*v1.MaintenanceTimeResV1 `json:"instance_maintenance_times"`
	ExecutionUserName        string                     `json:"execution_user_name"`
}

// GetSummaryOfWorkflowTasksV2
// @Summary 获取工单数据源任务概览
// @Description get summary of workflow instance tasks
// @Tags workflow
// @Id getSummaryOfInstanceTasksV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param project_name path string true "project name"
// @Success 200 {object} v2.GetWorkflowTasksResV2
// @router /v2/projects/{project_name}/workflows/{workflow_id}/tasks [get]
func GetSummaryOfWorkflowTasksV2(c echo.Context) error {
	return nil
}

type CreateWorkflowReqV2 struct {
	Subject    string `json:"workflow_subject" form:"workflow_subject" valid:"required,name"`
	WorkflowId string `json:"workflow_id" form:"workflow_id" valid:"required"`
	Desc       string `json:"desc" form:"desc"`
	TaskIds    []uint `json:"task_ids" form:"task_ids" valid:"required"`
}

// CreateWorkflowV2
// @Summary 创建工单
// @Description create workflow
// @Accept json
// @Produce json
// @Tags workflow
// @Id createWorkflowV2
// @Security ApiKeyAuth
// @Param instance body v2.CreateWorkflowReqV2 true "create workflow request"
// @Param project_name path string true "project name"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows [post]
func CreateWorkflowV2(c echo.Context) error {
	req := new(CreateWorkflowReqV2)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	projectName := c.Param("project_name")

	s := model.GetStorage()
	project, exist, err := s.GetProjectByName(projectName)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if !exist {
		return controller.JSONBaseErrorReq(c, v1.ErrProjectNotExist(projectName))
	}

	user, err := controller.GetCurrentUser(c)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	if err := v1.CheckIsProjectMember(user.Name, project.Name); err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	_, exist, err = s.GetWorkflowByProjectNameAndWorkflowId(project.Name, req.WorkflowId)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if exist {
		return controller.JSONBaseErrorReq(c, errors.New(errors.DataExist, fmt.Errorf("workflow is exist")))
	}

	taskIds := utils.RemoveDuplicateUint(req.TaskIds)
	if len(taskIds) > v1.MaximumDataSourceNum {
		return controller.JSONBaseErrorReq(c, errors.New(errors.DataConflict, fmt.Errorf("the max task count of a workflow is %v", v1.MaximumDataSourceNum)))
	}
	tasks, foundAllTasks, err := s.GetTasksByIds(taskIds)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if !foundAllTasks {
		return controller.JSONBaseErrorReq(c, v1.ErrTaskNoAccess)
	}

	insIdtMap := make(map[uint] /* project instance id */ struct{}, len(project.Instances))
	for _, instance := range project.Instances {
		insIdtMap[instance.ID] = struct{}{}
	}

	workflowTemplateId := tasks[0].Instance.WorkflowTemplateId
	for _, task := range tasks {
		if task.Instance == nil {
			return controller.JSONBaseErrorReq(c, errors.New(errors.DataNotExist, fmt.Errorf("instance is not exist. taskId=%v", task.ID)))
		}

		if _, ok := insIdtMap[task.InstanceId]; !ok {
			return controller.JSONBaseErrorReq(c, errors.New(errors.DataNotExist, fmt.Errorf("instance is not in project. taskId=%v", task.ID)))
		}

		count, err := s.GetTaskSQLCountByTaskID(task.ID)
		if err != nil {
			return controller.JSONBaseErrorReq(c, err)
		}
		if count == 0 {
			return controller.JSONBaseErrorReq(c, errors.New(errors.DataInvalid, fmt.Errorf("workflow's execute sql is null. taskId=%v", task.ID)))
		}

		if task.CreateUserId != user.ID {
			return controller.JSONBaseErrorReq(c, errors.New(errors.DataConflict,
				fmt.Errorf("the task is not created by yourself. taskId=%v", task.ID)))
		}

		if task.SQLSource == model.TaskSQLSourceFromMyBatisXMLFile {
			return controller.JSONBaseErrorReq(c, v1.ErrForbidMyBatisXMLTask(task.ID))
		}

		// all instances must use the same workflow template
		if task.Instance.WorkflowTemplateId != workflowTemplateId {
			return controller.JSONBaseErrorReq(c, errors.New(errors.DataConflict,
				fmt.Errorf("all instances must use the same workflow template")))
		}
	}

	// check user role operations
	{
		err = v1.CheckCurrentUserCanCreateWorkflow(user, tasks, projectName)
		if err != nil {
			return controller.JSONBaseErrorReq(c, err)
		}
	}

	count, err := s.GetWorkflowRecordCountByTaskIds(taskIds)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if count > 0 {
		return controller.JSONBaseErrorReq(c, errTaskHasBeenUsed)
	}

	template, exist, err := s.GetWorkflowTemplateById(workflowTemplateId)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if !exist {
		return controller.JSONBaseErrorReq(c, errors.New(errors.DataNotExist,
			fmt.Errorf("the task instance is not bound workflow template")))
	}

	err = v1.CheckWorkflowCanCommit(template, tasks)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	stepTemplates, err := s.GetWorkflowStepsByTemplateId(template.ID)
	if err != nil {
		return err
	}
	err = s.CreateWorkflowV2(req.Subject, req.WorkflowId, req.Desc, user, tasks, stepTemplates, project.ID)
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}

	workflow, exist, err := s.GetLastWorkflow()
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if !exist {
		return controller.JSONBaseErrorReq(c, errors.New(errors.DataNotExist, fmt.Errorf("should exist at least one workflow after create workflow")))
	}

	workFlowId := strconv.Itoa(int(workflow.ID))
	go notification.NotifyWorkflow(workFlowId, notification.WorkflowNotifyTypeCreate)

	go im.CreateApprove(workFlowId)

	return c.JSON(http.StatusOK, controller.NewBaseReq(nil))
}

type UpdateWorkflowReqV2 struct {
	TaskIds []uint `json:"task_ids" form:"task_ids" valid:"required"`
}

// UpdateWorkflowV2
// @Summary 更新工单（驳回后才可更新）
// @Description update workflow when it is rejected to creator.
// @Tags workflow
// @Accept json
// @Produce json
// @Id updateWorkflowV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param project_name path string true "project name"
// @Param instance body v2.UpdateWorkflowReqV2 true "update workflow request"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/ [patch]
func UpdateWorkflowV2(c echo.Context) error {
	return nil
}

type UpdateWorkflowScheduleReqV2 struct {
	ScheduleTime *time.Time `json:"schedule_time"`
}

// UpdateWorkflowScheduleV2
// @Summary 设置工单数据源定时上线时间（设置为空则代表取消定时时间，需要SQL审核流程都通过后才可以设置）
// @Description update workflow schedule.
// @Tags workflow
// @Accept json
// @Produce json
// @Id updateWorkflowScheduleV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param task_id path string true "task id"
// @Param project_name path string true "project name"
// @Param instance body v2.UpdateWorkflowScheduleReqV2 true "update workflow schedule request"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/tasks/{task_id}/schedule [put]
func UpdateWorkflowScheduleV2(c echo.Context) error {
	return nil
}

// ExecuteTasksOnWorkflowV2
// @Summary 多数据源批量上线
// @Description execute tasks on workflow
// @Tags workflow
// @Id executeTasksOnWorkflowV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param project_name path string true "project name"
// @Success 200 {object} controller.BaseRes
// @router /v2/projects/{project_name}/workflows/{workflow_id}/tasks/execute [post]
func ExecuteTasksOnWorkflowV2(c echo.Context) error {
	return nil
}

type GetWorkflowResV2 struct {
	controller.BaseRes
	Data *WorkflowResV2 `json:"data"`
}

type WorkflowTaskItem struct {
	Id uint `json:"task_id"`
}

type WorkflowRecordResV2 struct {
	Tasks             []*WorkflowTaskItem  `json:"tasks"`
	CurrentStepNumber uint                 `json:"current_step_number,omitempty"`
	Status            string               `json:"status" enums:"wait_for_audit,wait_for_execution,rejected,canceled,exec_failed,executing,finished"`
	Steps             []*WorkflowStepResV2 `json:"workflow_step_list,omitempty"`
}

type WorkflowResV2 struct {
	Name          string                 `json:"workflow_name"`
	WorkflowID    string                 `json:"workflow_id"`
	Desc          string                 `json:"desc,omitempty"`
	Mode          string                 `json:"mode" enums:"same_sqls,different_sqls"`
	CreateUser    string                 `json:"create_user_name"`
	CreateTime    *time.Time             `json:"create_time"`
	Record        *WorkflowRecordResV2   `json:"record"`
	RecordHistory []*WorkflowRecordResV2 `json:"record_history_list,omitempty"`
}

// GetWorkflowV2
// @Summary 获取工单详情
// @Description get workflow detail
// @Tags workflow
// @Id getWorkflowV2
// @Security ApiKeyAuth
// @Param workflow_id path string true "workflow id"
// @Param project_name path string true "project name"
// @Success 200 {object} GetWorkflowResV2
// @router /v2/projects/{project_name}/workflows/{workflow_id}/ [get]
func GetWorkflowV2(c echo.Context) error {
	return nil
}
