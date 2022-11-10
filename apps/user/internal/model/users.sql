CREATE TABLE `users`
(
    `id`          VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '用户ID',
    `mobile`      VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '手机号码',
    `name`        VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '姓名',
    `email`       VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '电子邮件',
    `avatar`      VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像',
    `dept_id`     VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '部门ID',
    `role_id`     VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '角色ID',
    `status`      INT          NOT NULL DEFAULT 0 COMMENT '状态：0-正常 1-停用',
    `is_deleted`  INT          NOT NULL DEFAULT 0 COMMENT '软删除：0-未删除 1-已删除',
    `create_time` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `index_id` (`id`)
) ENGINE INNODB
  DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';