package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/service"

var (
	Api                 = new(api)
	serviceMsModel      = service.Service.MsModel
	serviceMsDataset    = service.Service.MsDataset
	serviceMsSpace      = service.Service.MsSpace
	serviceMsDiscussion = service.Service.MsDiscussion
)

type api struct {
	MsModel      msModel
	MsDataset    msDataset
	MsSpace      msSpace
	MsDiscussion msDiscussion
}
