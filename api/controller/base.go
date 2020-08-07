package controller

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/url"
	"sqle/errors"
	"sqle/model"
	"sqle/utils"
)

var INSTANCE_NOT_EXIST_ERROR = NewBaseReq(errors.New(errors.INSTANCE_NOT_EXIST, fmt.Errorf("instance not exist")))
var INSTANCE_EXIST_ERROR = NewBaseReq(errors.New(errors.INSTANCE_EXIST, fmt.Errorf("inst is exist")))
var TASK_NOT_EXIST = NewBaseReq(errors.New(errors.TASK_NOT_EXIST, fmt.Errorf("task not exist")))

type BaseRes struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"ok"`
}

func NewBaseReq(err error) BaseRes {
	res := BaseRes{}
	switch e := err.(type) {
	case *errors.CodeError:
		res.Code = e.Code()
		res.Message = e.Error()
	default:
		if err == nil {
			res.Code = 0
			res.Message = "ok"
		} else {
			res.Code = -1
			res.Message = e.Error()
		}
	}
	return res
}

func readFileToByte(c echo.Context, name string) (fileName string, data []byte, err error) {
	file, err := c.FormFile(name)
	if err != nil {
		err = errors.New(errors.READ_UPLOAD_FILE_ERROR, err)
		return
	}
	src, err := file.Open()
	if err != nil {
		err = errors.New(errors.READ_UPLOAD_FILE_ERROR, err)
		return
	}
	defer src.Close()
	data, err = ioutil.ReadAll(src)
	if err != nil {
		err = errors.New(errors.READ_UPLOAD_FILE_ERROR, err)
		return
	}
	return
}

type CustomValidator struct {
}

func (cv *CustomValidator) Validate(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	return err
}

func unescapeParamString(params []*string) error {
	for i, p := range params {
		r, err := url.QueryUnescape(*p)
		if nil != err {
			return fmt.Errorf("unescape param [%v] failed: %v", params, err)
		}
		*params[i] = r
	}
	return nil
}

const ConfigPath = "/opt/sqle/etc/sqled.cnf"

// @Summary 加载数据库参数
// @Description reload base info
// @Accept x-www-form-urlencoded
// @Param mysql_user formData string true "mysql user"
// @Param mysql_password formData string true "mysql password"
// @Param mysql_host formData string true "mysql host"
// @Param mysql_port formData string true "mysql port"
// @Param mysql_schema formData string true "mysql schema"
// @Param config_path formData string false "confif path (Absolute Path)"
// @Success 200 {object} controller.BaseRes
// @router /base/reload [post]
func ReloadBaseInfo(c echo.Context) error {
	mysqlUser := c.FormValue("mysql_user")
	mysqlPassword := c.FormValue("mysql_password")
	mysqlHost := c.FormValue("mysql_host")
	mysqlPort := c.FormValue("mysql_port")
	mysqlSchema := c.FormValue("mysql_schema")
	configPath := c.FormValue("config_path")
	if configPath == "" {
		configPath = ConfigPath
	}
	conf, err := utils.LoadIniConf(configPath)
	if err != nil {
		return c.JSON(200, NewBaseReq(fmt.Errorf("load config path: %s failed", configPath)))
	}
	section ,err :=conf.GetSection("server")
	if err != nil {
		return c.JSON(200, NewBaseReq(err))
	}
	section.Key("mysql_user").SetValue(mysqlUser)
	section.Key("mysql_password").SetValue(mysqlPassword)
	section.Key("mysql_host").SetValue(mysqlHost)
	section.Key("mysql_port").SetValue(mysqlPort)
	section.Key("mysql_schema").SetValue(mysqlSchema)
	err = conf.Save(configPath)
	if err != nil {
		return c.JSON(200, NewBaseReq(err))
	}

	s, err := model.NewStorage(mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema, false)
	if err != nil {
		return c.JSON(200, NewBaseReq(err))
	}
	model.UpdateStorage(s)
	return c.JSON(200, NewBaseReq(nil))
}
