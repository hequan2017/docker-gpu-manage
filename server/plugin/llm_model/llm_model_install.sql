-- =====================================================
-- LLM Model 插件初始化 SQL
-- =====================================================

CREATE TABLE IF NOT EXISTS `llm_models` (
  `id` bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(100) NOT NULL COMMENT '模型名称',
  `publisher` varchar(100) COMMENT '发布者/机构',
  `type` varchar(50) DEFAULT 'general_llm' COMMENT '模型类型',
  `parameters` varchar(50) COMMENT '参数量(如7B, 13B)',
  `url` varchar(255) NOT NULL COMMENT '魔搭社区地址',
  `description` varchar(500) COMMENT '模型简介',
  PRIMARY KEY (`id`),
  INDEX `idx_llm_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
