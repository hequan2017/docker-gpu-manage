package system

import (
	"context"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCasbin = initOrderApiIgnore + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i *initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "888", V1: "/user/admin_register", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/deleteApisByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/api/syncApi", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/api/getApiGroups", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/api/enterSyncApi", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/ignoreApi", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/resetPassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setSelfSetting", V2: "PUT"},

		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/findFile", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/breakpointContinueFinish", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/breakpointContinue", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/removeChunk", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/importURL", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/getServerInfo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/getJumpboxConfig", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/customer/customerList", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getMeta", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/preview", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getColumn", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/rollback", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/delSysHistory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getSysHistory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/createPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getTemplates", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/delPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/createPlug", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/installPlugin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/pubPlug", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/addFunc", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/mcp", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/mcpTest", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/mcpList", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getDictionaryTreeList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getDictionaryTreeListByType", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getDictionaryDetailsByParent", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getDictionaryPath", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/email/emailTest", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/email/sendEmail", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/simpleUploader/upload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/simpleUploader/checkFileMd5", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/simpleUploader/mergeFileMd5", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/authorityBtn/setAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/getAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/canRemoveAuthorityBtn", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/createSysExportTemplate", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/deleteSysExportTemplate", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/deleteSysExportTemplateByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/updateSysExportTemplate", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/findSysExportTemplate", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/getSysExportTemplateList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/exportExcel", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/exportTemplate", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysExportTemplate/importExcel", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/info/createInfo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/info/deleteInfo", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/info/deleteInfoByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/info/updateInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/info/findInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/info/getInfoList", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/sysParams/createSysParams", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysParams/deleteSysParams", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysParams/deleteSysParamsByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysParams/updateSysParams", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysParams/findSysParams", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysParams/getSysParamsList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysParams/getSysParam", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/attachmentCategory/getCategoryList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/attachmentCategory/addCategory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/attachmentCategory/deleteCategory", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/sysVersion/findSysVersion", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysVersion/getSysVersionList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysVersion/downloadVersionJson", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysVersion/exportVersion", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysVersion/importVersion", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysVersion/deleteSysVersion", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysVersion/deleteSysVersionByIds", V2: "DELETE"},

		// 实例管理相关权限
		{Ptype: "p", V0: "888", V1: "/instance/createInstance", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/instance/deleteInstance", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/instance/deleteInstanceByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/instance/updateInstance", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/instance/findInstance", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/instance/getInstanceList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/instance/getAvailableNodes", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/instance/startContainer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/instance/stopContainer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/instance/restartContainer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/instance/getContainerLogs", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/instance/getContainerStats", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/instance/terminal", V2: "GET"},

		// 算力节点相关权限
		{Ptype: "p", V0: "888", V1: "/computeNode/createComputeNode", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/computeNode/deleteComputeNode", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/computeNode/deleteComputeNodeByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/computeNode/updateComputeNode", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/computeNode/findComputeNode", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/computeNode/getComputeNodeList", V2: "GET"},

		// PCDN相关权限
		{Ptype: "p", V0: "888", V1: "/pcdn/node/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/pcdn/node/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/pcdn/node/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/pcdn/node/find", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/pcdn/node/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/pcdn/node/online", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/pcdn/node/weight", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/find", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/grayRelease", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/pcdn/policy/switch", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/pcdn/task/list", V2: "GET"},

		// 镜像库相关权限
		{Ptype: "p", V0: "888", V1: "/imageRegistry/createImageRegistry", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/imageRegistry/deleteImageRegistry", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/imageRegistry/deleteImageRegistryByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/imageRegistry/updateImageRegistry", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/imageRegistry/findImageRegistry", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/imageRegistry/getImageRegistryList", V2: "GET"},

		// 产品规格相关权限
		{Ptype: "p", V0: "888", V1: "/productSpec/createProductSpec", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/productSpec/deleteProductSpec", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/productSpec/deleteProductSpecByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/productSpec/updateProductSpec", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/productSpec/findProductSpec", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/productSpec/getProductSpecList", V2: "GET"},

		// 端口转发相关权限
		{Ptype: "p", V0: "888", V1: "/portForward/createPortForward", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/portForward/deletePortForward", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/portForward/deletePortForwardByIds", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/portForward/updatePortForward", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/portForward/updatePortForwardStatus", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/portForward/findPortForward", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/portForward/getPortForwardList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/portForward/getServerIP", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/portForward/getForwarderStatus", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/portForward/getAllForwarderStatus", V2: "GET"},

		{Ptype: "p", V0: "8881", V1: "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/getAllApis", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/authority/setDataAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/importURL", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/system/getJumpboxConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/customer/customerList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/user/getUserInfo", V2: "GET"},

		{Ptype: "p", V0: "9528", V1: "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/importURL", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/system/getJumpboxConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/customer/customerList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
