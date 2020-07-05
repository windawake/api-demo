/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.131
Source Server Version : 50725
Source Host           : 192.168.1.131:3306
Source Database       : gorm

Target Server Type    : MYSQL
Target Server Version : 50725
File Encoding         : 65001

Date: 2020-07-06 00:46:32
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `city_id` int(11) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `hobby` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of profile
-- ----------------------------
INSERT INTO `profile` VALUES ('1', '1', '1', '11', '篮球', null, null, null);
INSERT INTO `profile` VALUES ('2', '2', '2', '22', '足球', null, null, null);
INSERT INTO `profile` VALUES ('3', '3', '3', '33', '滑板', null, null, null);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `is_super` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '超级管理员', '1');
INSERT INTO `role` VALUES ('2', '售前主管', '0');
INSERT INTO `role` VALUES ('3', '售前客服', '0');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'zhangsan', 'zs@qq.com', '1123', '3', '1', null, null, null);
INSERT INTO `user` VALUES ('2', 'lisi', 'lisi@qq.com', '3344', '1', '0', null, null, null);
INSERT INTO `user` VALUES ('3', 'wangwu', 'ww@qq.com', '5556', '2', '0', null, null, null);
