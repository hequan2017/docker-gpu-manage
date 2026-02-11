
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/approval_flow/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/approval_flow/model/request"
)

var ApprovalProcess = new(approval)

type approval struct {}
// CreateApprovalProcess 创建发版申请记录
// Author [yourname](https://github.com/yourname)
func (s *approval) CreateApprovalProcess(ctx context.Context, approval *model.ApprovalProcess) (err error) {
	err = global.GVA_DB.Create(approval).Error
	return err
}

// DeleteApprovalProcess 删除发版申请记录
// Author [yourname](https://github.com/yourname)
func (s *approval) DeleteApprovalProcess(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.ApprovalProcess{},"id = ?",ID).Error
	return err
}

// DeleteApprovalProcessByIds 批量删除发版申请记录
// Author [yourname](https://github.com/yourname)
func (s *approval) DeleteApprovalProcessByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.ApprovalProcess{},"id in ?",IDs).Error
	return err
}

// UpdateApprovalProcess 更新发版申请记录
// Author [yourname](https://github.com/yourname)
func (s *approval) UpdateApprovalProcess(ctx context.Context, approval model.ApprovalProcess) (err error) {
	err = global.GVA_DB.Model(&model.ApprovalProcess{}).Where("id = ?",approval.ID).Updates(&approval).Error
	return err
}

// GetApprovalProcess 根据ID获取发版申请记录
// Author [yourname](https://github.com/yourname)
func (s *approval) GetApprovalProcess(ctx context.Context, ID string) (approval model.ApprovalProcess, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&approval).Error
	return
}
// GetApprovalProcessInfoList 分页获取发版申请记录
// Author [yourname](https://github.com/yourname)
func (s *approval) GetApprovalProcessInfoList(ctx context.Context, info request.ApprovalProcessSearch) (list []model.ApprovalProcess, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ApprovalProcess{})
    var approvals []model.ApprovalProcess
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title LIKE ?", "%"+ *info.Title+"%")
    }
    if info.Version != nil && *info.Version != "" {
        db = db.Where("version LIKE ?", "%"+ *info.Version+"%")
    }
    if info.Content != nil && *info.Content != "" {
        db = db.Where("content LIKE ?", "%"+ *info.Content+"%")
    }
    if info.TargetServer != nil && *info.TargetServer != "" {
        db = db.Where("target_server LIKE ?", "%"+ *info.TargetServer+"%")
    }
    if info.Command != nil && *info.Command != "" {
        db = db.Where("command LIKE ?", "%"+ *info.Command+"%")
    }
    if info.Status != "" {
        db = db.Where("status = ?", info.Status)
    }
    if info.ApplicantId != nil {
        db = db.Where("applicant_id = ?", *info.ApplicantId)
    }
    if info.ApproverId != nil {
        db = db.Where("approver_id = ?", *info.ApproverId)
    }
    if info.Logs != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&approvals).Error
	return  approvals, total, err
}

func (s *approval)GetApprovalProcessPublic(ctx context.Context) {

}

// ApproveRequest 批准发版申请
func (s *approval) ApproveRequest(ctx context.Context, approvalProcess *model.ApprovalProcess) (err error) {
    // status: 60 = Approved/Executing
    err = global.GVA_DB.WithContext(ctx).Model(&model.ApprovalProcess{}).Where("id = ?", approvalProcess.ID).Update("status", 60).Error
    return err
}

// RejectRequest 驳回发版申请
func (s *approval) RejectRequest(ctx context.Context, approvalProcess *model.ApprovalProcess) (err error) {
    // status: 70 = Rejected
    err = global.GVA_DB.WithContext(ctx).Model(&model.ApprovalProcess{}).Where("id = ?", approvalProcess.ID).Update("status", 70).Error
    return err
}

