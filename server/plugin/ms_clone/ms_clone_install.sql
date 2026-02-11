-- =====================================================
-- ModelScope Clone 插件初始化 SQL
-- =====================================================

-- 模型库表
CREATE TABLE IF NOT EXISTS `gva_ms_models` (
  `id` bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(100) NOT NULL COMMENT '模型名称',
  `cover` varchar(255) COMMENT '封面图',
  `description` varchar(500) COMMENT '简介',
  `task_type` varchar(191) COMMENT '任务类型',
  `publisher` varchar(191) COMMENT '发布者',
  `readme` mediumtext COMMENT '详情文档',
  `download_count` int DEFAULT 0 COMMENT '下载量',
  PRIMARY KEY (`id`),
  INDEX `idx_gva_ms_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据集表
CREATE TABLE IF NOT EXISTS `gva_ms_datasets` (
  `id` bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(100) NOT NULL COMMENT '数据集名称',
  `cover` varchar(255) COMMENT '封面图',
  `description` varchar(500) COMMENT '简介',
  `size` varchar(191) COMMENT '数据集大小',
  `publisher` varchar(191) COMMENT '发布者',
  `readme` mediumtext COMMENT '详情文档',
  PRIMARY KEY (`id`),
  INDEX `idx_gva_ms_datasets_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创空间表
CREATE TABLE IF NOT EXISTS `gva_ms_spaces` (
  `id` bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(100) NOT NULL COMMENT '空间名称',
  `cover` varchar(255) COMMENT '封面图',
  `description` varchar(500) COMMENT '简介',
  `sdk` varchar(191) COMMENT 'SDK类型(Gradio/Streamlit)',
  `status` varchar(191) COMMENT '状态',
  `app_file` varchar(191) COMMENT '入口文件路径',
  PRIMARY KEY (`id`),
  INDEX `idx_gva_ms_spaces_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 社区讨论表
CREATE TABLE IF NOT EXISTS `gva_ms_discussions` (
  `id` bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `content` mediumtext COMMENT '内容',
  `user_id` bigint unsigned COMMENT '用户ID',
  `related_id` bigint unsigned COMMENT '关联ID',
  `related_type` varchar(191) COMMENT '关联类型(Model/Dataset/Space)',
  PRIMARY KEY (`id`),
  INDEX `idx_gva_ms_discussions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
