/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50718
Source Host           : 127.0.0.1:3306
Source Database       : finder

Target Server Type    : MYSQL
Target Server Version : 50718
File Encoding         : 65001

Date: 2019-06-25 17:03:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comment`
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `record_id` int(11) NOT NULL DEFAULT '0' COMMENT '失踪记录信息id',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '线索发布者联系方式',
  `remark` text COMMENT '线索描述',
  `photo` varchar(255) DEFAULT NULL,
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `record_id` (`record_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='线索表';

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for `record`
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `publisher_id` int(11) NOT NULL DEFAULT '0' COMMENT '发布人id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '失踪者姓名',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '失踪者性别 0 男 1 女',
  `photo` varchar(255) NOT NULL DEFAULT '' COMMENT '照片',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '失踪地点',
  `date` varchar(32) NOT NULL DEFAULT '' COMMENT '失踪日期',
  `remark` text COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核状态 0 未审核 1 审核通过 2 审核未通过',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `publisher_id` (`publisher_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='发布记录表';

-- ----------------------------
-- Records of record
-- ----------------------------

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键id',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(32) NOT NULL DEFAULT '',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '手机号码',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT '电子邮件',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户状态 0 正常 1 非正常',
  `role` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户角色 1 普通用户 2 志愿者 3 管理员',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
