// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDeployModel = "deploy"

// DeployModel mapped from table <deploy>
type DeployModel struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name            string         `gorm:"column:name;comment:部署标题"`                            // 部署标题
	ProjectID       int            `gorm:"column:project_id;comment:项目ID"`                      // 项目ID
	Project         string         `gorm:"column:project;comment:项目"`                           // 项目
	Fingerprint     string         `gorm:"column:fingerprint;comment:模板指纹"`                     // 模板指纹
	TemplateName    string         `gorm:"column:template_name;comment:模板名"`                    // 模板名
	TemplateContent string         `gorm:"column:template_content;comment:模板原始值"`               // 模板原始值
	TemplateParse   string         `gorm:"column:template_parse;comment:模板解析值"`                 // 模板解析值
	Params          string         `gorm:"column:params;comment:模板变量"`                          // 模板变量
	Status          int            `gorm:"column:status;comment:状态 0：等待上线 1：上线中 2：上线成功 3：上线失败"` // 状态 0：等待上线 1：上线中 2：上线成功 3：上线失败
	CreateTime      time.Time      `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime      time.Time      `gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
	DeleteTime      gorm.DeletedAt `gorm:"column:delete_time"`
}

// TableName DeployModel's table name
func (*DeployModel) TableName() string {
	return TableNameDeployModel
}
