CREATE TABLE `depts`
(
    `id`          VARCHAR(64) NOT NULL DEFAULT '' COMMENT '部门ID',
    `name`        VARCHAR(50) NOT NULL DEFAULT '' COMMENT '部门名称',
    `level`       INT         NOT NULL DEFAULT 0 COMMENT '层级',
    `parent_id`   VARCHAR(64) NOT NULL DEFAULT '' COMMENT '上级部门ID',
    `header_id`   VARCHAR(64) NOT NULL DEFAULT '' COMMENT '部门负责人ID',
    `status`      INT         NOT NULL DEFAULT 0 COMMENT '状态：0-正常 1-禁用',
    `create_time` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `index_id` (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4 COMMENT = '部门表';