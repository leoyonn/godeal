-- sql for create all tables
-- @author leo
-- @data 2014-03-14

DROP TABLE IF EXISTS `demo`;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
CREATE TABLE `demo` (
  `v` varchar(128) NOT NULL,
  PRIMARY KEY (`v`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account` varchar(32) NOT NULL,
  `name` varchar(32) NOT NULL,
  `desc` varchar(128),
  `gender` tinyint,
  `email` varchar(32) NOT NULL,
  `phone` varchar(32) NOT NULL,
  `avatar` varchar(64) NULL,
  `password` varchar(32) NOT NULL,
  `passtoken` varchar(128),
  `createAt` timestamp,
  PRIMARY KEY (`id`), UNIQUE(`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
SET character_set_client = @saved_cs_client;

DROP TABLE IF EXISTS `good`;
CREATE TABLE `good` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `category` varchar(32),
  `desc` varchar(128),
  `owner` bigint(20),
  `buyer` bigint(20),
  `ttl` bigint(20),
  `price` float,
  `status` tinyint,
  `icon` varchar(64),
  `createAt` timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `need`;
CREATE TABLE `need` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `category` varchar(32),
  `desc` varchar(128),
  `owner` bigint(20),
  `seller` bigint(20),
  `ttl` bigint(20),
  `price` float,
  `status` tinyint,
  `icon` varchar(64),
  `createAt` timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `name` varchar(32) NOT NULL,
  `desc` varchar(128),
  `parent` varchar(32),
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET character_set_client = @saved_cs_client;
