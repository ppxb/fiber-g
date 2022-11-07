-- CREATE DEPT DATABASE
DROP DATABASE IF EXISTS fiber_dept;
CREATE DATABASE fiber_dept;
USE fiber_dept;

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

-- CREATE USER DATABASE
DROP DATABASE IF EXISTS fiber_user;
CREATE DATABASE fiber_user;
USE fiber_user;

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
    KEY `index_id` (`id`)
) ENGINE INNODB
  DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

-- CREATE ASSET DATABASE
DROP DATABASE IF EXISTS fiber_asset;
CREATE DATABASE fiber_asset DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
USE fiber_asset;

CREATE TABLE `assets`
(
    `id`               VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '资产ID',
    `serial`           VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '编号',
    `name`             VARCHAR(200) NOT NULL DEFAULT '' COMMENT '名称',
    `project_id`       VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '项目ID',
    `son_project_id`   VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '子项目ID',
    `part_project_id`  VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '分部项目ID',
    `type`             VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '属性',
    `sub_district`     VARCHAR(255) NOT NULL DEFAULT '' COMMENT '街道',
    `brand`            VARCHAR(200) NOT NULL DEFAULT '' COMMENT '品牌',
    `model`            VARCHAR(255) NOT NULL DEFAULT '' COMMENT '型号',
    `unit`             VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '单位',
    `params`           TEXT         NULL COMMENT '参数',
    `value`            DECIMAL      NOT NULL DEFAULT 0 COMMENT '价值',
    `address`          VARCHAR(255) NOT NULL DEFAULT '' COMMENT '位置',
    `long`             VARCHAR(100) NOT NULL DEFAULT '' COMMENT '经度',
    `lat`              VARCHAR(100) NOT NULL DEFAULT '' COMMENT '纬度',
    `img_url`          VARCHAR(255) NOT NULL DEFAULT '' COMMENT '图片地址',
    `system_login_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '登录地址',
    `system_login_pwd` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '登录密码',
    `is_deleted`       INT          NOT NULL DEFAULT 0 COMMENT '软删除：0-未删除 1-已删除',
    `create_time`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `index_serial` (`serial`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4 COMMENT = '资产表';