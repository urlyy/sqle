package v1

import (
	"net/http"

	"actiontech.cloud/sqle/sqle/sqle/api/controller"
	"actiontech.cloud/sqle/sqle/sqle/model"

	"github.com/labstack/echo/v4"
)

type UpdateSMTPConfigurationReqV1 struct {
	Host     *string `json:"smtp_host" form:"smtp_host" example:"smtp.exmail.qq.com"`
	Port     *string `json:"smtp_port" form:"smtp_port" example:"465" valid:"omitempty,port"`
	Username *string `json:"smtp_username" form:"smtp_username" example:"test@qq.com" valid:"omitempty,email"`
	Password *string `json:"smtp_password" form:"smtp_password" example:"123"`
}

// @Summary 添加 SMTP 配置
// @Description update SMTP configuration
// @Accept json
// @Id updateSMTPConfigurationV1
// @Tags configuration
// @Security ApiKeyAuth
// @Param instance body v1.UpdateSMTPConfigurationReqV1 true "update SMTP configuration req"
// @Success 200 {object} controller.BaseRes
// @router /v1/configurations/smtp [patch]
func UpdateSMTPConfiguration(c echo.Context) error {
	req := new(UpdateSMTPConfigurationReqV1)
	if err := controller.BindAndValidateReq(c, req); err != nil {
		return err
	}
	s := model.GetStorage()
	smtpC, _, err := s.GetSMTPConfiguration()
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	if req.Host != nil {
		smtpC.Host = *req.Host
	}
	if req.Port != nil {
		smtpC.Port = *req.Port
	}
	if req.Username != nil {
		smtpC.Username = *req.Username
	}
	if req.Password != nil {
		smtpC.Password = *req.Password
	}

	if err := s.Save(smtpC); err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	return controller.JSONBaseErrorReq(c, nil)
}

type GetSMTPConfigurationResV1 struct {
	controller.BaseRes
	Data SMTPConfigurationResV1 `json:"data"`
}

type SMTPConfigurationResV1 struct {
	Host     string `json:"smtp_host"`
	Port     string `json:"smtp_port"`
	Username string `json:"smtp_username"`
}

// @Summary 获取 SMTP 配置
// @Description get SMTP configuration
// @Id getSMTPConfigurationV1
// @Tags configuration
// @Security ApiKeyAuth
// @Success 200 {object} v1.GetSMTPConfigurationResV1
// @router /v1/configurations/smtp [get]
func GetSMTPConfiguration(c echo.Context) error {
	s := model.GetStorage()
	smtpC, _, err := s.GetSMTPConfiguration()
	if err != nil {
		return controller.JSONBaseErrorReq(c, err)
	}
	return c.JSON(http.StatusOK, &GetSMTPConfigurationResV1{
		BaseRes: controller.NewBaseReq(nil),
		Data: SMTPConfigurationResV1{
			Host:     smtpC.Host,
			Port:     smtpC.Port,
			Username: smtpC.Username,
		},
	})
}