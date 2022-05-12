package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type SqlQueryConfig struct {
	MaxPreQueryRows    int `json:"max_pre_query_rows"`
	QueryTimeoutSecond int `json:"query_timeout_second"`
}

// Scan impl sql.Scanner interface
func (c *SqlQueryConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal json value: %v", value)
	}
	if len(bytes) == 0 {
		return nil
	}
	result := SqlQueryConfig{}
	err := json.Unmarshal(bytes, &result)
	*c = result
	return err
}

// Value impl sql.driver.Valuer interface
func (c SqlQueryConfig) Value() (driver.Value, error) {
	v, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json value: %v", v)
	}
	return v, err
}

type SqlQueryHistory struct {
	Model
	CreateUserId uint   `json:"create_user_id" gorm:"not null"`
	InstanceId   uint   `json:"instance_id" gorm:"not null"`
	Database     string `json:"database"`
	RawSql       string `json:"raw_sql" gorm:"type:text;not null"`
}

type SqlQueryExecutionSql struct {
	Model
	SqlQueryHistoryId uint            `json:"sql_query_history_id" gorm:"not null"`
	Sql               string          `json:"sql" gorm:"type:text;not null"`
	ExecStartAt       *time.Time      `json:"exec_start_at"`
	ExecEndAt         *time.Time      `json:"exec_end_at"`
	ExecResult        string          `json:"exec_result" gorm:"type:text"`
	RawSqlInfo        SqlQueryHistory `json:"raw_sql_info" gorm:"foreignkey:SqlQueryHistoryId"`
}
