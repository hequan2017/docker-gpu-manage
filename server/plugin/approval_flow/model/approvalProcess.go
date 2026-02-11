
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ApprovalProcess 发版申请 结构体
type ApprovalProcess struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"comment:申请标题;column:title;size:100;" binding:"required"`  //申请标题
  Version  *string `json:"version" form:"version" gorm:"comment:版本号;column:version;size:50;" binding:"required"`  //版本号
  Content  *string `json:"content" form:"content" gorm:"comment:发版内容;column:content;size:1000;"`  //发版内容
  TargetServer  *string `json:"targetServer" form:"targetServer" gorm:"comment:目标服务器IP;column:target_server;size:50;" binding:"required"`  //目标服务器
  Command  *string `json:"command" form:"command" gorm:"comment:执行命令/脚本;column:command;size:500;" binding:"required"`  //执行命令
  Status  string `json:"status" form:"status" gorm:"default:pending;comment:审批状态;column:status;type:varchar(50);" binding:"required"`  //状态
  ApplicantId  *int64 `json:"applicantId" form:"applicantId" gorm:"comment:申请人ID;column:applicant_id;"`  //申请人
  ApproverId  *int64 `json:"approverId" form:"approverId" gorm:"comment:审批人ID;column:approver_id;"`  //审批人
  Logs  *string `json:"logs" form:"logs" gorm:"comment:执行日志;column:logs;size:5000;type:text;"`  //执行日志
}


// TableName 发版申请 ApprovalProcess自定义表名 af_approval_processes
func (ApprovalProcess) TableName() string {
    return "af_approval_processes"
}







