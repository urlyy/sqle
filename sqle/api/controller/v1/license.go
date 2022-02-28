//go:build !enterprise
// +build !enterprise

package v1

import (
	e "errors"

	"github.com/actiontech/sqle/sqle/api/controller"
	"github.com/actiontech/sqle/sqle/errors"

	"github.com/labstack/echo/v4"
)

var ErrNoLicenseRequired = errors.New(errors.ErrAccessDeniedError, e.New("sqle-ce no license required"))

type LicenseItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Limit       int    `json:"limit"`
}

type GetLicenseResV1 struct {
	Content string        `json:"content"`
	License []LicenseItem `json:"license"`
}

// GetLicense get sqle license
// @Summary 获取 sqle license
// @Description get sqle license
// @Id getSQLELicenseV1
// @Tags configuration
// @Security ApiKeyAuth
// @Success 200 {object} v1.GetLicenseResV1
// @router /v1/configurations/license [get]
func GetLicense(c echo.Context) error {
	return controller.JSONBaseErrorReq(c, ErrNoLicenseRequired)
}

// GetSQLELicenseInfo get the information needed to generate the sqle license
// @Summary 获取生成 sqle license需要的的信息
// @Description get the information needed to generate the sqle license
// @Id GetSQLELicenseInfoV1
// @Tags configuration
// @Security ApiKeyAuth
// @Success 200 file 1 "server info"
// @router /v1/configurations/license/info [get]
func GetSQLELicenseInfo(c echo.Context) error {
	return controller.JSONBaseErrorReq(c, ErrNoLicenseRequired)
}

// SetLicense set sqle license
// @Summary 导入 sqle license
// @Description set sqle license
// @Id setSQLELicenseV1
// @Tags configuration
// @Security ApiKeyAuth
// @Param license_file formData file true "SQLE license file"
// @Success 200 {object} controller.BaseRes
// @router /v1/configurations/license [post]
func SetLicense(c echo.Context) error {
	return controller.JSONBaseErrorReq(c, ErrNoLicenseRequired)
}

type CheckLicenseResV1 struct {
	Content string        `json:"content"`
	License []LicenseItem `json:"license"`
}

// CheckLicense parse and check sqle license
// @Summary 解析和校验 sqle license
// @Description parse and check sqle license
// @Id checkSQLELicenseV1
// @Tags configuration
// @Security ApiKeyAuth
// @Param license_file formData file true "SQLE license file"
// @Success 200 {object} v1.CheckLicenseResV1
// @router /v1/configurations/license/check [post]
func CheckLicense(c echo.Context) error {
	return controller.JSONBaseErrorReq(c, ErrNoLicenseRequired)
}
