/*
 Navicat Premium Data Transfer

 Source Server         : taotao
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : blog_service

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 28/07/2020 21:48:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned NOT NULL,
  `tag_id` int(10) unsigned DEFAULT NULL,
  `created_on` int(10) DEFAULT '0',
  `created_by` varchar(100) DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT NULL,
  `modified_by` varchar(100) DEFAULT NULL,
  `deleted_on` int(10) unsigned DEFAULT NULL,
  `is_del` tinyint(3) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_article_id` (`article_id`) USING BTREE,
  KEY `idx_tag_id` (`tag_id`) USING BTREE,
  KEY `idx_is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for blog_articles
-- ----------------------------
DROP TABLE IF EXISTS `blog_articles`;
CREATE TABLE `blog_articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL COMMENT '文章标题',
  `desc` varchar(255) DEFAULT NULL COMMENT '文章简述',
  `cover_image_url` varchar(255) DEFAULT NULL COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT NULL COMMENT '创建人',
  `modified_on` int(10) DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT NULL COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`) USING BTREE,
  KEY `idx_state` (`state`) USING BTREE,
  KEY `idx_is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '',
  `app_secret` varchar(50) DEFAULT '',
  `created_by` varchar(100) DEFAULT NULL,
  `modified_by` varchar(100) DEFAULT NULL,
  `created_on` int(10) unsigned DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT NULL,
  `deleted_on` int(10) unsigned DEFAULT NULL,
  `is_del` tinyint(3) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
BEGIN;
INSERT INTO `blog_auth` VALUES (1, 'eddycjy', 'go-programming-tour-book', 'eddycjy', NULL, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for blog_tags
-- ----------------------------
DROP TABLE IF EXISTS `blog_tags`;
CREATE TABLE `blog_tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `state` tinyint(3) unsigned DEFAULT '1',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modifed_on` int(10) unsigned DEFAULT '0' COMMENT '更新时间',
  `modified_by` varchar(100) DEFAULT NULL COMMENT '更新人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:未删除,1: 删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_deleted_is_del` (`is_del`) USING BTREE,
  KEY `idx_state` (`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
