-- =====================================================
-- 天启算力管理平台 - 完整菜单和API权限SQL
-- =====================================================
-- 说明：此SQL包含系统所有主要菜单的完整配置
-- 使用方法：在数据库中直接执行此SQL即可
-- 注意：执行前请先备份数据库！
-- =====================================================

-- =====================================================
-- 第一部分：父级菜单插入
-- =====================================================

-- 1. 仪表盘（dashboard）- Sort: 1
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '{"title":"仪表盘","icon":"odometer"}', '仪表盘')
ON DUPLICATE KEY UPDATE `title` = '仪表盘';

-- 2. 项目说明（project）- Sort: 2
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'project', 'project', 0, 'view/about/project.vue', 2, '{"title":"项目说明","icon":"document"}', '项目说明')
ON DUPLICATE KEY UPDATE `title` = '项目说明';

-- 3. 超级管理员（superAdmin）- Sort: 3
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '{"title":"超级管理员","icon":"user"}', '超级管理员')
ON DUPLICATE KEY UPDATE `title` = '超级管理员';

-- 4. 个人信息（person）- Sort: 4（隐藏菜单）
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'person', 'person', 1, 'view/person/person.vue', 4, '{"title":"个人信息","icon":"message"}', '个人信息')
ON DUPLICATE KEY UPDATE `title` = '个人信息';

-- 5. 系统工具（systemTools）- Sort: 5
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '{"title":"系统工具","icon":"tools"}', '系统工具')
ON DUPLICATE KEY UPDATE `title` = '系统工具';

-- 6. 插件系统（plugin）- Sort: 6
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '{"title":"插件系统","icon":"cherry"}', '插件系统')
ON DUPLICATE KEY UPDATE `title` = '插件系统';

-- 7. 示例文件（example）- Sort: 7
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'example', 'example', 0, 'view/example/index.vue', 7, '{"title":"示例文件","icon":"management"}', '示例文件')
ON DUPLICATE KEY UPDATE `title` = '示例文件';

-- 8. 服务器状态（state）- Sort: 8
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'state', 'state', 0, 'view/system/state.vue', 8, '{"title":"服务器状态","icon":"cloudy"}', '服务器状态')
ON DUPLICATE KEY UPDATE `title` = '服务器状态';

-- 9. 关于我们（about）- Sort: 9
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'about', 'about', 0, 'view/about/index.vue', 9, '{"title":"关于我们","icon":"info-filled"}', '关于我们')
ON DUPLICATE KEY UPDATE `title` = '关于我们';

-- 10. 端口转发（portForward）- Sort: 10
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'portForward', 'portForward', 0, 'view/routerHolder.vue', 10, '{"title":"端口转发","icon":"position"}', '端口转发')
ON DUPLICATE KEY UPDATE `title` = '端口转发';

-- 业务菜单 - 镜像库（imageRegistry）
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'imageRegistry', 'imageRegistry', 0, 'view/imageregistry/imageRegistry/imageRegistry.vue', 11, '{"title":"镜像库","icon":"crop"}', '镜像库')
ON DUPLICATE KEY UPDATE `title` = '镜像库';

-- 业务菜单 - 算力节点（computeNode）
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'computeNode', 'computeNode', 0, 'view/computenode/computeNode/computeNode.vue', 12, '{"title":"算力节点","icon":"chicken"}', '算力节点')
ON DUPLICATE KEY UPDATE `title` = '算力节点';

-- 业务菜单 - 产品规格（productSpec）
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'productSpec', 'productSpec', 0, 'view/product/productSpec/productSpec.vue', 13, '{"title":"产品规格","icon":"baseball"}', '产品规格')
ON DUPLICATE KEY UPDATE `title` = '产品规格';

-- 业务菜单 - 实例管理（instance）
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'instance', 'instance', 0, 'view/instance/instance/instance.vue', 14, '{"title":"实例管理","icon":"briefcase"}', '实例管理')
ON DUPLICATE KEY UPDATE `title` = '实例管理';

-- K8s 管理父菜单 - Sort: 15
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES (NOW(), NOW(), 0, 0, 'k8s', 'k8s', 0, 'view/routerHolder.vue', 15, '{"title":"K8s管理","icon":"cpu-line"}', 'K8s管理')
ON DUPLICATE KEY UPDATE `title` = 'K8s管理';

-- =====================================================
-- 第二部分：子菜单插入
-- =====================================================

-- 获取父级菜单ID
SET @super_admin_id = (SELECT id FROM `sys_base_menus` WHERE `name` = 'superAdmin' LIMIT 1);
SET @system_tools_id = (SELECT id FROM `sys_base_menus` WHERE `name` = 'systemTools' LIMIT 1);
SET @example_id = (SELECT id FROM `sys_base_menus` WHERE `name` = 'example' LIMIT 1);
SET @plugin_id = (SELECT id FROM `sys_base_menus` WHERE `name` = 'plugin' LIMIT 1);
SET @port_forward_id = (SELECT id FROM `sys_base_menus` WHERE `name` = 'portForward' LIMIT 1);
SET @k8s_id = (SELECT id FROM `sys_base_menus` WHERE `name` = 'k8s' LIMIT 1);

-- 超级管理员子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES
(NOW(), NOW(), 1, @super_admin_id, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '{"title":"角色管理","icon":"avatar"}', '角色管理'),
(NOW(), NOW(), 1, @super_admin_id, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '{"title":"菜单管理","icon":"tickets","keepAlive":true}', '菜单管理'),
(NOW(), NOW(), 1, @super_admin_id, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '{"title":"api管理","icon":"platform","keepAlive":true}', 'api管理'),
(NOW(), NOW(), 1, @super_admin_id, 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '{"title":"用户管理","icon":"coordinate"}', '用户管理'),
(NOW(), NOW(), 1, @super_admin_id, 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '{"title":"字典管理","icon":"notebook"}', '字典管理'),
(NOW(), NOW(), 1, @super_admin_id, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '{"title":"操作历史","icon":"pie-chart"}', '操作历史'),
(NOW(), NOW(), 1, @super_admin_id, 'sysParams', 'sysParams', 0, 'view/superAdmin/params/sysParams.vue', 7, '{"title":"参数管理","icon":"compass"}', '参数管理')
ON DUPLICATE KEY UPDATE `title` = VALUES(`title`);

-- 系统工具子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES
(NOW(), NOW(), 1, @system_tools_id, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '{"title":"代码生成器","icon":"cpu","keepAlive":true}', '代码生成器'),
(NOW(), NOW(), 1, @system_tools_id, 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 2, '{"title":"自动化代码管理","icon":"magic-stick"}', '自动化代码管理'),
(NOW(), NOW(), 1, @system_tools_id, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 3, '{"title":"表单生成器","icon":"magic-stick","keepAlive":true}', '表单生成器'),
(NOW(), NOW(), 1, @system_tools_id, 'system', 'system', 0, 'view/systemTools/system/system.vue', 4, '{"title":"系统配置","icon":"operation"}', '系统配置'),
(NOW(), NOW(), 1, @system_tools_id, 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 5, '{"title":"模板配置","icon":"folder"}', '模板配置'),
(NOW(), NOW(), 1, @system_tools_id, 'exportTemplate', 'exportTemplate', 0, 'view/systemTools/exportTemplate/exportTemplate.vue', 6, '{"title":"导出模板","icon":"reading"}', '导出模板'),
(NOW(), NOW(), 1, @system_tools_id, 'picture', 'picture', 0, 'view/systemTools/autoCode/picture.vue', 7, '{"title":"AI页面绘制","icon":"picture-filled"}', 'AI页面绘制'),
(NOW(), NOW(), 1, @system_tools_id, 'mcpTool', 'mcpTool', 0, 'view/systemTools/autoCode/mcp.vue', 8, '{"title":"Mcp Tools模板","icon":"magnet"}', 'Mcp Tools模板'),
(NOW(), NOW(), 1, @system_tools_id, 'mcpTest', 'mcpTest', 0, 'view/systemTools/autoCode/mcpTest.vue', 9, '{"title":"Mcp Tools测试","icon":"partly-cloudy"}', 'Mcp Tools测试'),
(NOW(), NOW(), 1, @system_tools_id, 'sysVersion', 'sysVersion', 0, 'view/systemTools/version/version.vue', 10, '{"title":"版本管理","icon":"server"}', '版本管理')
ON DUPLICATE KEY UPDATE `title` = VALUES(`title`);

-- 插件系统子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES
(NOW(), NOW(), 1, @plugin_id, 'https://plugin.tianqi.com/', 'https://plugin.tianqi.com/', 0, 'https://plugin.tianqi.com/', 0, '{"title":"插件市场","icon":"shop"}', '插件市场'),
(NOW(), NOW(), 1, @plugin_id, 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '{"title":"插件安装","icon":"box"}', '插件安装'),
(NOW(), NOW(), 1, @plugin_id, 'pubPlug', 'pubPlug', 0, 'view/systemTools/pubPlug/pubPlug.vue', 3, '{"title":"打包插件","icon":"files"}', '打包插件'),
(NOW(), NOW(), 1, @plugin_id, 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 4, '{"title":"邮件插件","icon":"message"}', '邮件插件'),
(NOW(), NOW(), 1, @plugin_id, 'anInfo', 'anInfo', 0, 'plugin/announcement/view/info.vue', 5, '{"title":"公告管理[示例]","icon":"scaleToOriginal"}', '公告管理[示例]')
ON DUPLICATE KEY UPDATE `title` = VALUES(`title`);

-- 端口转发子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES
(NOW(), NOW(), 1, @port_forward_id, 'portForwardRules', 'portForwardRules', 0, 'plugin/portforward/view/portForward.vue', 1, '{"title":"转发规则","icon":"list"}', '转发规则')
ON DUPLICATE KEY UPDATE `title` = VALUES(`title`);

-- 示例文件子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES
(NOW(), NOW(), 1, @example_id, 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '{"title":"媒体库（上传下载）","icon":"upload"}', '媒体库（上传下载）'),
(NOW(), NOW(), 1, @example_id, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '{"title":"断点续传","icon":"upload-filled"}', '断点续传'),
(NOW(), NOW(), 1, @example_id, 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '{"title":"客户列表（资源示例）","icon":"avatar"}', '客户列表（资源示例）')
ON DUPLICATE KEY UPDATE `title` = VALUES(`title`);

-- K8s 管理子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`, `title`)
VALUES
(NOW(), NOW(), 1, @k8s_id, 'k8sCluster', 'k8sCluster', 0, 'plugin/k8smanager/view/cluster.vue', 1, '{"title":"集群管理","icon":"server-line"}', '集群管理'),
(NOW(), NOW(), 1, @k8s_id, 'k8sPod', 'k8sPod', 0, 'plugin/k8smanager/view/pod.vue', 2, '{"title":"Pod管理","icon":"apps-line"}', 'Pod管理'),
(NOW(), NOW(), 1, @k8s_id, 'k8sDeployment', 'k8sDeployment', 0, 'plugin/k8smanager/view/deployment.vue', 3, '{"title":"Deployment管理","icon":"stack-line"}', 'Deployment管理'),
(NOW(), NOW(), 1, @k8s_id, 'k8sService', 'k8sService', 0, 'plugin/k8smanager/view/service.vue', 4, '{"title":"Service管理","icon":"links-line"}', 'Service管理'),
(NOW(), NOW(), 1, @k8s_id, 'k8sNamespace', 'k8sNamespace', 0, 'plugin/k8smanager/view/namespace.vue', 5, '{"title":"Namespace管理","icon":"folder-line"}', 'Namespace管理'),
(NOW(), NOW(), 1, @k8s_id, 'k8sEvent', 'k8sEvent', 0, 'plugin/k8smanager/view/event.vue', 6, '{"title":"事件管理","icon":"notification-line"}', '事件管理')
ON DUPLICATE KEY UPDATE `title` = VALUES(`title`);

-- =====================================================
-- 第三部分：为角色分配菜单权限
-- =====================================================

-- 1. 为超级管理员(role_id = 888)分配所有菜单
-- 注意：超级管理员通常通过代码自动拥有所有权限，这里不需要额外配置
-- 如果需要，可以执行：
-- INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
-- SELECT 888, id, NOW(), NOW() FROM `sys_base_menus`
-- ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2. 为普通用户(role_id = 8881)分配基础菜单
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
SELECT 8881, id, NOW(), NOW()
FROM `sys_base_menus`
WHERE `name` IN ('dashboard', 'project', 'about', 'person', 'state')
AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus` am
    WHERE am.`sys_authority_authority_id` = 8881
    AND am.`sys_base_menu_id` = `sys_base_menus`.id
);

-- 3. 为测试角色(role_id = 9528)分配部分菜单
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
SELECT 9528, id, NOW(), NOW()
FROM `sys_base_menus`
WHERE `parent_id` = 0  -- 所有父级菜单
AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus` am
    WHERE am.`sys_authority_authority_id` = 9528
    AND am.`sys_base_menu_id` = `sys_base_menus`.id
);

-- 为测试角色添加系统工具和示例的子菜单
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
SELECT 9528, id, NOW(), NOW()
FROM `sys_base_menus`
WHERE `parent_id` IN (
    SELECT id FROM `sys_base_menus` WHERE `name` IN ('systemTools', 'example')
)
AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus` am
    WHERE am.`sys_authority_authority_id` = 9528
    AND am.`sys_base_menu_id` = `sys_base_menus`.id
);

-- =====================================================
-- 第四部分：验证查询
-- =====================================================

-- 查询所有父级菜单
-- SELECT id, name, title, path, sort FROM `sys_base_menus` WHERE `parent_id` = 0 ORDER BY sort;

-- 查询"项目说明"菜单
-- SELECT * FROM `sys_base_menus` WHERE `name` = 'project';

-- 查询各角色的菜单数量
-- SELECT
--     a.authority_id,
--     a.authority_name,
--     COUNT(am.sys_base_menu_id) as menu_count
-- FROM `sys_authorities` a
-- LEFT JOIN `sys_authority_menus` am ON a.authority_id = am.sys_authority_authority_id
-- GROUP BY a.authority_id, a.authority_name
-- ORDER BY a.authority_id;

-- 查询普通用户的菜单
-- SELECT
--     m.id,
--     m.title,
--     m.path,
--     m.sort
-- FROM `sys_authority_menus` am
-- JOIN `sys_base_menus` m ON am.sys_base_menu_id = m.id
-- WHERE am.sys_authority_authority_id = 8881
-- ORDER BY m.sort;

-- =====================================================
-- 第五部分：清理SQL（谨慎使用！）
-- =====================================================

-- 删除所有菜单（慎用！）
-- DELETE FROM `sys_authority_menus`;
-- DELETE FROM `sys_base_menus`;

-- 删除特定角色的菜单权限
-- DELETE FROM `sys_authority_menus` WHERE `sys_authority_authority_id` = 8881;

-- 删除"项目说明"菜单及其权限关联
-- DELETE FROM `sys_authority_menus` WHERE `sys_base_menu_id` = (SELECT id FROM `sys_base_menus` WHERE `name` = 'project');
-- DELETE FROM `sys_base_menus` WHERE `name` = 'project';

-- =====================================================
-- 第六部分：K8s 管理插件 - API 和数据表
-- =====================================================

-- 创建 K8s 集群配置表
CREATE TABLE IF NOT EXISTS `k8s_clusters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '集群名称',
  `kube_config` longtext COMMENT 'kubeconfig配置内容',
  `endpoint` varchar(500) DEFAULT NULL COMMENT 'API Server地址',
  `version` varchar(50) DEFAULT NULL COMMENT 'K8s版本',
  `status` varchar(20) DEFAULT 'unknown' COMMENT '集群状态',
  `description` varchar(500) DEFAULT NULL COMMENT '集群描述',
  `region` varchar(100) DEFAULT NULL COMMENT '区域',
  `provider` varchar(50) DEFAULT NULL COMMENT '云服务商',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '是否默认集群',
  `node_count` int DEFAULT '0' COMMENT '节点数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='K8s集群配置表';

-- 创建 K8s API 接口
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
-- K8s集群相关API (8个)
(NOW(), NOW(), '/k8s/cluster/create', '创建K8s集群', 'K8s集群', 'POST'),
(NOW(), NOW(), '/k8s/cluster/delete', '删除K8s集群', 'K8s集群', 'DELETE'),
(NOW(), NOW(), '/k8s/cluster/deleteByIds', '批量删除K8s集群', 'K8s集群', 'DELETE'),
(NOW(), NOW(), '/k8s/cluster/update', '更新K8s集群', 'K8s集群', 'PUT'),
(NOW(), NOW(), '/k8s/cluster/get', '获取K8s集群详情', 'K8s集群', 'GET'),
(NOW(), NOW(), '/k8s/cluster/list', '获取K8s集群列表', 'K8s集群', 'GET'),
(NOW(), NOW(), '/k8s/cluster/refresh', '刷新K8s集群状态', 'K8s集群', 'POST'),
(NOW(), NOW(), '/k8s/cluster/all', '获取所有K8s集群', 'K8s集群', 'GET'),
-- Pod相关API (6个)
(NOW(), NOW(), '/k8s/pod/list', '获取Pod列表', 'K8s Pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/get', '获取Pod详情', 'K8s Pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/delete', '删除Pod', 'K8s Pod', 'DELETE'),
(NOW(), NOW(), '/k8s/pod/log', '获取Pod日志', 'K8s Pod', 'POST'),
(NOW(), NOW(), '/k8s/pod/containers', '获取Pod容器列表', 'K8s Pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/events', '获取Pod事件', 'K8s Pod', 'GET'),
-- Deployment相关API (6个)
(NOW(), NOW(), '/k8s/deployment/list', '获取Deployment列表', 'K8s Deployment', 'GET'),
(NOW(), NOW(), '/k8s/deployment/get', '获取Deployment详情', 'K8s Deployment', 'GET'),
(NOW(), NOW(), '/k8s/deployment/scale', '扩缩容Deployment', 'K8s Deployment', 'POST'),
(NOW(), NOW(), '/k8s/deployment/restart', '重启Deployment', 'K8s Deployment', 'POST'),
(NOW(), NOW(), '/k8s/deployment/delete', '删除Deployment', 'K8s Deployment', 'DELETE'),
(NOW(), NOW(), '/k8s/deployment/pods', '获取Deployment关联的Pods', 'K8s Deployment', 'GET'),
-- Service相关API (4个)
(NOW(), NOW(), '/k8s/service/list', '获取Service列表', 'K8s Service', 'GET'),
(NOW(), NOW(), '/k8s/service/get', '获取Service详情', 'K8s Service', 'GET'),
(NOW(), NOW(), '/k8s/service/delete', '删除Service', 'K8s Service', 'DELETE'),
(NOW(), NOW(), '/k8s/service/endpoints', '获取Service的Endpoints', 'K8s Service', 'GET'),
-- Namespace相关API (4个)
(NOW(), NOW(), '/k8s/namespace/list', '获取Namespace列表', 'K8s Namespace', 'GET'),
(NOW(), NOW(), '/k8s/namespace/get', '获取Namespace详情', 'K8s Namespace', 'GET'),
(NOW(), NOW(), '/k8s/namespace/create', '创建Namespace', 'K8s Namespace', 'POST'),
(NOW(), NOW(), '/k8s/namespace/delete', '删除Namespace', 'K8s Namespace', 'DELETE'),
-- Event相关API (1个)
(NOW(), NOW(), '/k8s/event/list', '获取Event列表', 'K8s Event', 'POST')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 为管理员角色授权 K8s 菜单（authority_id = 888）
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`)
SELECT 888, id FROM `sys_base_menus` WHERE `name` LIKE 'k8s%'
AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus` am
    WHERE am.`sys_authority_authority_id` = 888
    AND am.`sys_base_menu_id` = `sys_base_menus`.id
);

-- 为管理员角色授权 K8s API（使用 sys_casbin 表）
INSERT INTO `sys_casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
SELECT NULL, 'p', '888', path, method, '', '', '', ''
FROM `sys_apis`
WHERE `api_group` LIKE 'K8s%'
AND NOT EXISTS (
    SELECT 1 FROM `sys_casbin` c
    WHERE c.v0 = '888'
    AND c.v1 = `sys_apis`.path
    AND c.v2 = `sys_apis`.method
);

-- 验证 K8s 插件安装
SELECT '✅ K8s 管理插件集成完成！' as status;
SELECT COUNT(*) as 'K8s 菜单数量' FROM `sys_base_menus` WHERE `name` LIKE 'k8s%';
SELECT COUNT(*) as 'K8s API数量' FROM `sys_apis` WHERE `api_group` LIKE 'K8s%';
SELECT COUNT(*) as 'K8s 菜单权限数' FROM `sys_authority_menus` WHERE `sys_base_menu_id` IN (
    SELECT id FROM `sys_base_menus` WHERE `name` LIKE 'k8s%'
);
SELECT COUNT(*) as 'K8s API权限数' FROM `sys_casbin` WHERE `v1` LIKE '/k8s%' AND `v0` = '888';
