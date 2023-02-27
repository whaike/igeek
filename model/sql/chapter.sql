CREATE TABLE `chapter_info` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `course_id` int NOT NULL DEFAULT '0' COMMENT '课程id',
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '章节标题',
    `audio_id` int NOT NULL DEFAULT '0' COMMENT '音频id',
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
    `create_at` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
    `update_at` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='章节信息表';

-- goctl model mysql ddl -src chapter.sql -dir ./.. --style go_zero