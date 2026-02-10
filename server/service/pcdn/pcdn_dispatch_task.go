package pcdn

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
)

type PcdnDispatchTaskService struct{}

func (s *PcdnDispatchTaskService) GetPcdnDispatchTaskList(_ context.Context, info pcdnReq.PcdnDispatchTaskSearch) (list []pcdn.PcdnDispatchTask, total int64, err error) {
	db := global.GVA_DB.Model(&pcdn.PcdnDispatchTask{})
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.TraceID != "" {
		db = db.Where("trace_id = ?", info.TraceID)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1))
	}
	err = db.Order("id desc").Find(&list).Error
	return
}
