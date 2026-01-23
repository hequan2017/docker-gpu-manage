#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
权限同步脚本
功能：自动检查并补充缺失的菜单和API权限
"""

import pymysql
import json
import sys
from typing import List, Dict, Optional

# ============================================================================
# 配置区域 - 从项目 config.yaml 自动读取
# ============================================================================
DB_CONFIG = {
    'host': '192.168.112.155',
    'port': 3307,
    'user': 'hequan',
    'password': '123456',
    'database': 'docker-gpu',
    'charset': 'utf8mb4'
}

# 管理员角色ID
AUTHORITY_ID = 888

# ============================================================================
# 插件配置 - 定义所有需要的菜单和API
# ============================================================================

PLUGINS_CONFIG = {
    'k8smanager': {
        'parent_menu': {
            'path': 'k8s',
            'name': 'k8s',
            'title': 'K8s管理',
            'icon': 'cpu-line',
            'sort': 8,
            'component': 'view/routerHolder.vue',
            'parent_id': 0,
        },
        'sub_menus': [
            {
                'path': 'cluster',
                'name': 'k8sCluster',
                'title': '集群管理',
                'icon': 'server-line',
                'sort': 1,
                'component': 'plugin/k8smanager/view/cluster.vue',
                'apis': [
                    ('/k8s/cluster/create', 'POST', '创建K8s集群', 'k8s-cluster'),
                    ('/k8s/cluster/delete', 'DELETE', '删除K8s集群', 'k8s-cluster'),
                    ('/k8s/cluster/deleteByIds', 'DELETE', '批量删除K8s集群', 'k8s-cluster'),
                    ('/k8s/cluster/update', 'PUT', '更新K8s集群', 'k8s-cluster'),
                    ('/k8s/cluster/get', 'GET', '获取K8s集群详情', 'k8s-cluster'),
                    ('/k8s/cluster/list', 'GET', '获取K8s集群列表', 'k8s-cluster'),
                    ('/k8s/cluster/refresh', 'POST', '刷新集群状态', 'k8s-cluster'),
                    ('/k8s/cluster/all', 'GET', '获取所有集群', 'k8s-cluster'),
                ]
            },
            {
                'path': 'pod',
                'name': 'k8sPod',
                'title': 'Pod管理',
                'icon': 'apps-line',
                'sort': 2,
                'component': 'plugin/k8smanager/view/pod.vue',
                'apis': [
                    ('/k8s/pod/list', 'GET', '获取Pod列表', 'k8s-pod'),
                    ('/k8s/pod/get', 'GET', '获取Pod详情', 'k8s-pod'),
                    ('/k8s/pod/delete', 'DELETE', '删除Pod', 'k8s-pod'),
                    ('/k8s/pod/log', 'POST', '获取Pod日志', 'k8s-pod'),
                    ('/k8s/pod/containers', 'GET', '获取Pod容器列表', 'k8s-pod'),
                    ('/k8s/pod/events', 'GET', '获取Pod事件', 'k8s-pod'),
                ]
            },
            {
                'path': 'deployment',
                'name': 'k8sDeployment',
                'title': 'Deployment管理',
                'icon': 'stack-line',
                'sort': 3,
                'component': 'plugin/k8smanager/view/deployment.vue',
                'apis': [
                    ('/k8s/deployment/list', 'GET', '获取Deployment列表', 'k8s-deployment'),
                    ('/k8s/deployment/get', 'GET', '获取Deployment详情', 'k8s-deployment'),
                    ('/k8s/deployment/scale', 'POST', '扩缩容Deployment', 'k8s-deployment'),
                    ('/k8s/deployment/restart', 'POST', '重启Deployment', 'k8s-deployment'),
                    ('/k8s/deployment/delete', 'DELETE', '删除Deployment', 'k8s-deployment'),
                    ('/k8s/deployment/pods', 'GET', '获取Deployment关联的Pods', 'k8s-deployment'),
                ]
            },
            {
                'path': 'service',
                'name': 'k8sService',
                'title': 'Service管理',
                'icon': 'links-line',
                'sort': 4,
                'component': 'plugin/k8smanager/view/service.vue',
                'apis': [
                    ('/k8s/service/list', 'GET', '获取Service列表', 'k8s-service'),
                    ('/k8s/service/get', 'GET', '获取Service详情', 'k8s-service'),
                    ('/k8s/service/delete', 'DELETE', '删除Service', 'k8s-service'),
                    ('/k8s/service/endpoints', 'GET', '获取Service的Endpoints', 'k8s-service'),
                ]
            },
            {
                'path': 'namespace',
                'name': 'k8sNamespace',
                'title': 'Namespace管理',
                'icon': 'folder-line',
                'sort': 5,
                'component': 'plugin/k8smanager/view/namespace.vue',
                'apis': [
                    ('/k8s/namespace/list', 'GET', '获取Namespace列表', 'k8s-namespace'),
                    ('/k8s/namespace/get', 'GET', '获取Namespace详情', 'k8s-namespace'),
                    ('/k8s/namespace/create', 'POST', '创建Namespace', 'k8s-namespace'),
                    ('/k8s/namespace/delete', 'DELETE', '删除Namespace', 'k8s-namespace'),
                ]
            },
            {
                'path': 'event',
                'name': 'k8sEvent',
                'title': '事件管理',
                'icon': 'notification-line',
                'sort': 6,
                'component': 'plugin/k8smanager/view/event.vue',
                'apis': [
                    ('/k8s/event/list', 'POST', '获取Event列表', 'k8s-event'),
                ]
            },
            # ===== 缺失的菜单 =====
            {
                'path': 'metrics',
                'name': 'k8sMetrics',
                'title': '监控指标',
                'icon': 'line-chart-line',
                'sort': 7,
                'component': 'plugin/k8smanager/view/metrics.vue',
                'apis': [
                    ('/k8s/metrics/cluster', 'GET', '获取集群指标', 'k8s-metrics'),
                    ('/k8s/metrics/cluster/refresh', 'POST', '刷新集群指标', 'k8s-metrics'),
                    ('/k8s/metrics/nodes', 'GET', '获取节点指标', 'k8s-metrics'),
                    ('/k8s/metrics/pods', 'GET', '获取Pod指标', 'k8s-metrics'),
                    ('/k8s/metrics/summary', 'GET', '获取指标汇总', 'k8s-metrics'),
                    ('/k8s/metrics/collector/start', 'POST', '启动指标收集器', 'k8s-metrics'),
                    ('/k8s/metrics/collector/stop', 'POST', '停止指标收集器', 'k8s-metrics'),
                ]
            },
            {
                'path': 'audit',
                'name': 'k8sAudit',
                'title': '审计日志',
                'icon': 'file-list-line',
                'sort': 8,
                'component': 'plugin/k8smanager/view/audit.vue',
                'apis': [
                    ('/k8s/audit/list', 'GET', '获取审计日志列表', 'k8s-audit'),
                    ('/k8s/audit/stats', 'GET', '获取审计统计信息', 'k8s-audit'),
                    ('/k8s/audit/client-stats', 'GET', '获取客户端统计信息', 'k8s-audit'),
                    ('/k8s/audit/cleanup', 'DELETE', '清理审计日志', 'k8s-audit'),
                    ('/k8s/audit/export', 'GET', '导出审计日志', 'k8s-audit'),
                ]
            },
        ]
    },
    'portforward': {
        'parent_menu': {
            'path': 'portForward',
            'name': 'portForward',
            'title': '端口转发',
            'icon': 'position-line',
            'sort': 9,
            'component': 'view/routerHolder.vue',
            'parent_id': 0,
        },
        'sub_menus': [
            {
                'path': 'rules',
                'name': 'portForwardRules',
                'title': '转发规则',
                'icon': 'route-line',
                'sort': 1,
                'component': 'plugin/portforward/view/portForward.vue',
                'apis': [
                    ('/portForward/createPortForward', 'POST', '创建端口转发规则', 'portForward'),
                    ('/portForward/deletePortForward', 'DELETE', '删除端口转发规则', 'portForward'),
                    ('/portForward/deletePortForwardByIds', 'DELETE', '批量删除端口转发规则', 'portForward'),
                    ('/portForward/updatePortForward', 'PUT', '更新端口转发规则', 'portForward'),
                    ('/portForward/updatePortForwardStatus', 'PUT', '更新端口转发规则状态', 'portForward'),
                    ('/portForward/findPortForward', 'GET', '根据ID获取端口转发规则', 'portForward'),
                    ('/portForward/getPortForwardList', 'GET', '获取端口转发规则列表', 'portForward'),
                    ('/portForward/getServerIP', 'GET', '获取服务器IP地址', 'portForward'),
                    ('/portForward/getForwarderStatus', 'GET', '获取端口转发运行状态', 'portForward'),
                    ('/portForward/getAllForwarderStatus', 'GET', '获取所有端口转发运行状态', 'portForward'),
                ]
            }
        ]
    },
    'aiagent': {
        'parent_menu': {
            'path': 'aiagent',
            'name': 'aiagent',
            'title': 'AI Agent',
            'icon': 'chat-dot-square',
            'sort': 10,
            'component': 'view/routerHolder.vue',
            'parent_id': 0,  # 顶级菜单
        },
        'sub_menus': [
            {
                'path': 'chat',
                'name': 'aiagentChat',
                'title': 'AI 对话',
                'icon': 'chat-line-round',
                'sort': 1,
                'component': 'plugin/aiagent/view/chat.vue',
                'apis': [
                    ('/aiagent/conversation/createConversation', 'POST', '创建会话', 'AI Agent'),
                    ('/aiagent/conversation/deleteConversation', 'DELETE', '删除会话', 'AI Agent'),
                    ('/aiagent/conversation/updateConversation', 'PUT', '更新会话', 'AI Agent'),
                    ('/aiagent/conversation/findConversation', 'GET', '根据ID获取会话', 'AI Agent'),
                    ('/aiagent/conversation/getConversationList', 'GET', '获取会话列表', 'AI Agent'),
                    ('/aiagent/conversation/setActive', 'POST', '设置激活状态', 'AI Agent'),
                    ('/aiagent/conversation/getActive', 'GET', '获取激活的会话', 'AI Agent'),
                    ('/aiagent/message/getMessageList', 'GET', '获取消息列表', 'AI Agent'),
                    ('/aiagent/message/deleteMessage', 'DELETE', '删除消息', 'AI Agent'),
                    ('/aiagent/chat/sendMessage', 'POST', '发送消息', 'AI Agent'),
                ]
            },
            {
                'path': 'config',
                'name': 'aiagentConfig',
                'title': 'AI 配置',
                'icon': 'setting',
                'sort': 2,
                'component': 'plugin/aiagent/view/config.vue',
                'apis': [
                    ('/aiagent/config/createConfig', 'POST', '创建AI配置', 'AI Agent'),
                    ('/aiagent/config/deleteConfig', 'DELETE', '删除AI配置', 'AI Agent'),
                    ('/aiagent/config/updateConfig', 'PUT', '更新AI配置', 'AI Agent'),
                    ('/aiagent/config/findConfig', 'GET', '根据ID获取AI配置', 'AI Agent'),
                    ('/aiagent/config/getConfigList', 'GET', '获取AI配置列表', 'AI Agent'),
                    ('/aiagent/config/setActive', 'POST', '设置AI配置激活状态', 'AI Agent'),
                    ('/aiagent/config/getActive', 'GET', '获取激活的AI配置', 'AI Agent'),
                ]
            }
        ]
    },
    'finetuning': {
        'parent_menu': {
            'path': 'finetuning',
            'name': 'finetuning',
            'title': '算法微调',
            'icon': 'cpu',
            'sort': 11,
            'component': 'view/routerHolder.vue',
            'parent_id': 0,  # 顶级菜单
        },
        'sub_menus': [
            {
                'path': 'taskList',
                'name': 'finetuningTaskList',
                'title': '微调任务',
                'icon': 'list',
                'sort': 1,
                'component': 'plugin/finetuning/view/taskList.vue',
                'apis': [
                    ('/finetuning/createTask', 'POST', '创建微调任务', 'Finetuning'),
                    ('/finetuning/deleteTask', 'DELETE', '删除微调任务', 'Finetuning'),
                    ('/finetuning/stopTask', 'POST', '停止微调任务', 'Finetuning'),
                    ('/finetuning/getTask', 'GET', '根据ID获取微调任务', 'Finetuning'),
                    ('/finetuning/getTaskList', 'GET', '获取微调任务列表', 'Finetuning'),
                    ('/finetuning/getTaskLog', 'GET', '获取微调任务日志', 'Finetuning'),
                ]
            },
            {
                'path': 'taskDetail',
                'name': 'finetuningTaskDetail',
                'title': '任务详情',
                'icon': 'document',
                'sort': 2,
                'component': 'plugin/finetuning/view/taskDetail.vue',
                'hidden': 1,
                'apis': []
            }
        ]
    },
    'dellasset': {
        'parent_menu': {
            'path': 'dellAsset',
            'name': 'dellAsset',
            'title': '戴尔资产管理',
            'icon': 'cpu',
            'sort': 12,
            'component': 'plugin/dellasset/view/dellAsset.vue',
            'parent_id': 0,  # 顶级菜单
        },
        'sub_menus': [
            {
                'path': 'dellAssetList',
                'name': 'dellAssetList',
                'title': '资产管理',
                'icon': 'list',
                'sort': 1,
                'component': 'plugin/dellasset/view/dellAsset.vue',
                'apis': [
                    ('/dellAsset/createDellAsset', 'POST', '创建戴尔服务器资产', '戴尔资产'),
                    ('/dellAsset/deleteDellAsset', 'DELETE', '删除戴尔服务器资产', '戴尔资产'),
                    ('/dellAsset/deleteDellAssetByIds', 'DELETE', '批量删除戴尔服务器资产', '戴尔资产'),
                    ('/dellAsset/updateDellAsset', 'PUT', '更新戴尔服务器资产', '戴尔资产'),
                    ('/dellAsset/findDellAsset', 'GET', '查询戴尔服务器资产', '戴尔资产'),
                    ('/dellAsset/getDellAssetList', 'GET', '获取戴尔服务器资产列表', '戴尔资产'),
                    ('/dellAsset/getStatistics', 'GET', '获取资产统计信息', '戴尔资产'),
                ]
            }
        ]
    },
}


class PermissionSyncer:
    """权限同步器"""

    def __init__(self, db_config: Dict):
        self.db_config = db_config
        self.conn = None
        self.cursor = None
        self.stats = {
            'menus_added': 0,
            'menus_updated': 0,
            'apis_added': 0,
            'api_permissions_added': 0,
            'menu_permissions_added': 0,
            'errors': []
        }

    def connect(self):
        """连接数据库"""
        try:
            self.conn = pymysql.connect(**self.db_config)
            self.cursor = self.conn.cursor()
            print("[OK] Successfully connected to database: %s" % self.db_config['database'])
        except Exception as e:
            print("[ERROR] Database connection failed: %s" % e)
            sys.exit(1)

    def close(self):
        """关闭数据库连接"""
        if self.cursor:
            self.cursor.close()
        if self.conn:
            self.conn.close()

    def get_menu_id_by_path(self, path: str, parent_id: int = None) -> Optional[int]:
        """根据路径获取菜单ID"""
        if parent_id is not None:
            sql = "SELECT id FROM sys_base_menus WHERE path = %s AND parent_id = %s LIMIT 1"
            self.cursor.execute(sql, (path, parent_id))
        else:
            sql = "SELECT id FROM sys_base_menus WHERE path = %s LIMIT 1"
            self.cursor.execute(sql, (path,))
        result = self.cursor.fetchone()
        return result[0] if result else None

    def get_api_id_by_path(self, path: str, method: str) -> Optional[int]:
        """根据路径和方法获取API ID"""
        sql = "SELECT id FROM sys_apis WHERE path = %s AND method = %s LIMIT 1"
        self.cursor.execute(sql, (path, method))
        result = self.cursor.fetchone()
        return result[0] if result else None

    def check_casbin_permission(self, v0: str, v1: str, v2: str) -> bool:
        """检查 casbin 权限是否存在"""
        sql = "SELECT id FROM casbin_rule WHERE ptype = 'p' AND v0 = %s AND v1 = %s AND v2 = %s LIMIT 1"
        self.cursor.execute(sql, (v0, v1, v2))
        return self.cursor.fetchone() is not None

    def check_menu_authority(self, menu_id: int, authority_id: int) -> bool:
        """检查菜单权限是否存在"""
        sql = "SELECT * FROM sys_authority_menus WHERE sys_base_menu_id = %s AND sys_authority_authority_id = %s LIMIT 1"
        self.cursor.execute(sql, (menu_id, authority_id))
        return self.cursor.fetchone() is not None

    def create_menu(self, parent_id: int, path: str, name: str, title: str,
                    icon: str, sort: int, component: str, hidden: int = 0) -> int:
        """创建菜单"""
        # 确定menu_level: 0表示顶级菜单，1表示子菜单
        menu_level = 0 if parent_id == 0 else 1
        sql = """
            INSERT INTO sys_base_menus
            (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
            VALUES (NOW(), NOW(), %s, %s, %s, %s, %s, %s, %s, %s, %s)
        """
        try:
            self.cursor.execute(sql, (parent_id, path, name, hidden, component, sort, menu_level, title, icon))
            self.conn.commit()
            self.stats['menus_added'] += 1
            return self.cursor.lastrowid
        except Exception as e:
            self.stats['errors'].append("Failed to create menu %s: %s" % (name, e))
            return None

    def update_menu(self, menu_id: int, **kwargs):
        """更新菜单"""
        if not kwargs:
            return
        set_clause = ", ".join(["%s = %%s" % k for k in kwargs.keys()])
        values = list(kwargs.values())
        values.append(menu_id)
        sql = "UPDATE sys_base_menus SET %s, updated_at = NOW() WHERE id = %%s" % set_clause
        try:
            self.cursor.execute(sql, values)
            self.conn.commit()
            self.stats['menus_updated'] += 1
        except Exception as e:
            self.stats['errors'].append("Failed to update menu %s: %s" % (menu_id, e))

    def create_api(self, path: str, description: str, api_group: str, method: str) -> int:
        """创建API"""
        sql = """
            INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
            VALUES (NOW(), NOW(), %s, %s, %s, %s)
        """
        try:
            self.cursor.execute(sql, (path, description, api_group, method))
            self.conn.commit()
            self.stats['apis_added'] += 1
            return self.cursor.lastrowid
        except Exception as e:
            self.stats['errors'].append("Failed to create API %s: %s" % (path, e))
            return None

    def create_casbin_permission(self, authority_id: str, v1: str, v2: str):
        """创建 casbin 权限"""
        if self.check_casbin_permission(str(authority_id), v1, v2):
            return
        sql = """
            INSERT INTO casbin_rule (ptype, v0, v1, v2)
            VALUES ('p', %s, %s, %s)
        """
        try:
            self.cursor.execute(sql, (str(authority_id), v1, v2))
            self.conn.commit()
            self.stats['api_permissions_added'] += 1
        except Exception as e:
            self.stats['errors'].append("Failed to create API permission %s %s: %s" % (v1, v2, e))

    def create_menu_authority(self, menu_id: int, authority_id: int):
        """创建菜单权限"""
        if self.check_menu_authority(menu_id, authority_id):
            return
        sql = """
            INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
            VALUES (%s, %s)
        """
        try:
            self.cursor.execute(sql, (authority_id, menu_id))
            self.conn.commit()
            self.stats['menu_permissions_added'] += 1
        except Exception as e:
            self.stats['errors'].append("Failed to create menu permission %s: %s" % (menu_id, e))

    def sync_plugin(self, plugin_name: str, config: Dict):
        """同步单个插件的权限"""
        print("\n" + "=" * 60)
        print("[SYNC] Starting plugin: %s" % plugin_name)
        print("=" * 60)

        parent_menu_id = None

        # 处理父菜单
        if config.get('parent_menu'):
            parent = config['parent_menu']
            parent_menu_id = self.get_menu_id_by_path(parent['path'], parent.get('parent_id'))

            if not parent_menu_id:
                print("  [CREATE] Parent menu: %s" % parent['title'])
                parent_menu_id = self.create_menu(
                    parent_id=parent.get('parent_id', 0),
                    path=parent['path'],
                    name=parent['name'],
                    title=parent['title'],
                    icon=parent['icon'],
                    sort=parent['sort'],
                    component=parent['component']
                )
                if parent_menu_id:
                    print("     [OK] Parent menu created, ID: %s" % parent_menu_id)
            else:
                print("  [EXISTS] Parent menu: %s (ID: %s)" % (parent['title'], parent_menu_id))

            # 为父菜单授权
            if parent_menu_id:
                self.create_menu_authority(parent_menu_id, AUTHORITY_ID)

        # 处理子菜单
        for sub_menu in config.get('sub_menus', []):
            print("\n  [MENU] Processing: %s" % sub_menu['title'])

            # 确定父菜单ID
            if sub_menu.get('parent_id'):
                sm_parent_id = sub_menu['parent_id']
            elif parent_menu_id:
                sm_parent_id = parent_menu_id
            else:
                print("     [SKIP] Cannot determine parent menu ID")
                continue

            # 查找或创建子菜单
            menu_id = self.get_menu_id_by_path(sub_menu['path'], sm_parent_id)

            if not menu_id:
                print("     [CREATE] Sub menu")
                menu_id = self.create_menu(
                    parent_id=sm_parent_id,
                    path=sub_menu['path'],
                    name=sub_menu['name'],
                    title=sub_menu['title'],
                    icon=sub_menu['icon'],
                    sort=sub_menu['sort'],
                    component=sub_menu['component'],
                    hidden=sub_menu.get('hidden', 0)
                )
                if menu_id:
                    print("     [OK] Sub menu created, ID: %s" % menu_id)
            else:
                print("     [EXISTS] Sub menu, ID: %s" % menu_id)

            # 为子菜单授权
            if menu_id:
                self.create_menu_authority(menu_id, AUTHORITY_ID)

            # 处理API
            for api_path, method, description, api_group in sub_menu.get('apis', []):
                print("     [API] %s %s" % (method, api_path))

                # 查找或创建API
                api_id = self.get_api_id_by_path(api_path, method)
                if not api_id:
                    api_id = self.create_api(api_path, description, api_group, method)
                    if api_id:
                        print("         [OK] API created")
                else:
                    print("         [EXISTS] API")

                # 创建API权限
                self.create_casbin_permission(AUTHORITY_ID, api_path, method)

    def sync_all(self):
        """同步所有插件"""
        print("\n" + "=" * 60)
        print("[START] Starting permission sync")
        print("=" * 60)

        for plugin_name, config in PLUGINS_CONFIG.items():
            self.sync_plugin(plugin_name, config)

    def print_summary(self):
        """打印统计信息"""
        print("\n" + "=" * 60)
        print("[SUMMARY] Sync Statistics")
        print("=" * 60)
        print("  [+] Menus added: %s" % self.stats['menus_added'])
        print("  [~] Menus updated: %s" % self.stats['menus_updated'])
        print("  [+] APIs added: %s" % self.stats['apis_added'])
        print("  [+] API permissions added: %s" % self.stats['api_permissions_added'])
        print("  [+] Menu permissions added: %s" % self.stats['menu_permissions_added'])

        if self.stats['errors']:
            print("\n[ERRORS] List:")
            for error in self.stats['errors']:
                print("     - %s" % error)
        else:
            print("\n[SUCCESS] All operations completed without errors!")

        print("=" * 60 + "\n")


def main():
    """主函数"""
    syncer = PermissionSyncer(DB_CONFIG)
    syncer.connect()

    try:
        syncer.sync_all()
        syncer.print_summary()
    except Exception as e:
        print("\n[ERROR] An error occurred: %s" % e)
        import traceback
        traceback.print_exc()
    finally:
        syncer.close()


if __name__ == '__main__':
    main()
