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
    `id`                bigint(20)     NOT NULL AUTO_INCREMENT,
    `uuid`              varchar(64)    NOT NULL DEFAULT '' COMMENT '资产UUID',
    `project_id`        varchar(64)    NOT NULL DEFAULT '' COMMENT '项目ID',
    `son_project_id`    varchar(64)    NOT NULL DEFAULT '' COMMENT '子项目ID',
    `part_project_id`   varchar(64)    NOT NULL DEFAULT '' COMMENT '分部项目ID',
    `project_name`      varchar(255)   NOT NULL DEFAULT '' COMMENT '项目名称：冗余',
    `son_project_name`  varchar(255)   NOT NULL DEFAULT '' COMMENT '子项目名称：冗余',
    `part_project_name` varchar(255)   NOT NULL DEFAULT '' COMMENT '分部项目名称：冗余',
    `serial`            varchar(64)    NOT NULL DEFAULT '' COMMENT '资产编号',
    `name`              varchar(200)   NOT NULL DEFAULT '' COMMENT '名称',
    `type`              varchar(32)    NOT NULL DEFAULT '' COMMENT '资产类别',
    `kind`              varchar(50)    NOT NULL DEFAULT '' COMMENT '资产属性',
    `sub_district`      varchar(255)   NOT NULL DEFAULT '' COMMENT '街道',
    `brand`             varchar(200)   NOT NULL DEFAULT '' COMMENT '品牌',
    `model`             varchar(255)   NOT NULL DEFAULT '' COMMENT '型号',
    `unit`              varchar(20)    NOT NULL DEFAULT '' COMMENT '单位',
    `params`            text COMMENT '参数',
    `value`             decimal(10, 0) NOT NULL DEFAULT '0' COMMENT '价值',
    `address`           varchar(255)   NOT NULL DEFAULT '' COMMENT '位置',
    `long`              varchar(100)   NOT NULL DEFAULT '' COMMENT '经度',
    `lat`               varchar(100)   NOT NULL DEFAULT '' COMMENT '纬度',
    `img_url`           varchar(255)   NOT NULL DEFAULT '' COMMENT '图片地址',
    `system_login_url`  varchar(255)   NOT NULL DEFAULT '' COMMENT '登录地址',
    `system_login_pwd`  varchar(255)   NOT NULL DEFAULT '' COMMENT '登录密码',
    `iot_net_serial`    varchar(32)    NOT NULL DEFAULT '' COMMENT '物联网卡号/宽带号',
    `net_status`        varchar(255)   NOT NULL DEFAULT '' COMMENT '网络资费情况',
    `emeter_serial`     varchar(64)    NOT NULL DEFAULT '' COMMENT '电费使用情况',
    `is_deleted`        int(11)        NOT NULL DEFAULT '0' COMMENT '软删除：0-未删除 1-已删除',
    `create_time`       timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`       timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `index_serial` (`serial`),
    KEY `index_name` (`name`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='资产表';