package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/api"

var (
	Router          = new(router)
	apiMsModel      = api.Api.MsModel
	apiMsDataset    = api.Api.MsDataset
	apiMsSpace      = api.Api.MsSpace
	apiMsDiscussion = api.Api.MsDiscussion
)

type router struct {
	MsModel      msModel
	MsDataset    msDataset
	MsSpace      msSpace
	MsDiscussion msDiscussion
}
