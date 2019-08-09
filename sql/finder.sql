/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50718
Source Host           : 127.0.0.1:3306
Source Database       : finder

Target Server Type    : MYSQL
Target Server Version : 50718
File Encoding         : 65001

Date: 2019-08-09 11:43:34
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='线索表';

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('1', '1', '2', '3', '/static/img/ef3f8a5b7079f7ce2726e3d1676f89da.png', '2');
INSERT INTO `comment` VALUES ('4', '7', '2323', '23', '', '1565257673');
INSERT INTO `comment` VALUES ('5', '7', '2323', '23', '/static/img/ef3f8a5b7079f7ce2726e3d1676f89da.png', '1565257702');
INSERT INTO `comment` VALUES ('6', '7', '2323', '23', '/static/img/5eee1c52391b7ee5c787bf403ef98e16.png', '1565258156');
INSERT INTO `comment` VALUES ('7', '7', '2323', '23', '/static/img/c5d42b1d74fe4325b3187400d51a4dc3.png', '1565258161');

-- ----------------------------
-- Table structure for `record`
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `publisher_id` int(11) NOT NULL DEFAULT '0' COMMENT '发布人id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '失踪者姓名',
  `sex` tinyint(1) NOT NULL DEFAULT '1' COMMENT '失踪者性别 1 男 2 女',
  `photo` varchar(255) NOT NULL DEFAULT '' COMMENT '照片',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '失踪地点',
  `date` varchar(32) NOT NULL DEFAULT '' COMMENT '失踪日期',
  `remark` text COMMENT '备注',
  `isfind` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否被找到 1 未找到 2 已找到',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '审核状态 1 未审核 2 审核未通过 3 审核通过 ',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '发布时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `publisher_id` (`publisher_id`),
  KEY `status_name` (`status`,`name`),
  KEY `isfind_status` (`isfind`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='发布记录表';

-- ----------------------------
-- Records of record
-- ----------------------------
INSERT INTO `record` VALUES ('4', '2', '2', '1', '', '23', '2019-07-07', 'dsf', '1', '3', '1564134513', '1564134513');
INSERT INTO `record` VALUES ('5', '3', '2', '1', '', '2', '2019-07-05', 'sdf', '1', '3', '1564134605', '1565245906');
INSERT INTO `record` VALUES ('6', '3', 'tt', '1', '', 'tt', '2019-08-18', 'sdf', '2', '3', '1564714594', '1565245901');
INSERT INTO `record` VALUES ('7', '3', '2323', '1', '/static/img/ef3f8a5b7079f7ce2726e3d1676f89da.png', 'ff', '2019-08-30', 'sdf', '1', '3', '1564717331', '1565245895');
INSERT INTO `record` VALUES ('8', '3', '123', '1', '/static/img/ef3f8a5b7079f7ce2726e3d1676f89da.png', '23', '2019-08-18', 'fdsfdsf', '2', '3', '1564725173', '1565245888');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键id',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '姓名',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '手机号码',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT '电子邮件',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态 1 正常 2 非正常',
  `role` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户角色 1 普通用户 2 志愿者 3 管理员',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('3', '123456', 'c4ca4238a0b923820dcc509a6f75849b', 'ww', '231', '123456@qq.com', '3313', '1', '3', '1564123430', '1563955300');
INSERT INTO `user` VALUES ('7', 'test', '098f6bcd4621d373cade4e832627b4f6', 'test1', 'test2', '22@qq.com', '2', '1', '2', '1564123670', '1564048167');
INSERT INTO `user` VALUES ('8', '12345623', 'e10adc3949ba59abbe56e057f20f883e', '123456', '123456', '1234526@qq.com', 'dd', '1', '1', '1564122077', '1564122077');
INSERT INTO `user` VALUES ('9', 'qq', '099b3b060154898840f0ebdfb46ec78f', 'qq', 'qq', 'qq@qq.com', '', '1', '1', '1564125530', '1564125530');
