/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.34 : Database - wldz
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`wldz` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `wldz`;

/*Table structure for table `casbin_rule` */

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  KEY `IDX_casbin_rule_v1` (`v1`),
  KEY `IDX_casbin_rule_v2` (`v2`),
  KEY `IDX_casbin_rule_v3` (`v3`),
  KEY `IDX_casbin_rule_v4` (`v4`),
  KEY `IDX_casbin_rule_v5` (`v5`),
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_v0` (`v0`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

/*Data for the table `casbin_rule` */

insert  into `casbin_rule`(`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) values ('p','用户管理角色','/api/v1/system/user/list','GET','','',''),('p','用户管理角色','/api/v1/system/user/add','GET','','',''),('p','用户管理角色','/api/v1/system/user/change','GET','','',''),('p','用户管理角色','/api/v1/system/user/:id','GET','','',''),('p','删除用户角色333','/api/v1/system/user/:id','GET','','',''),('p','删除居民角色','/api/v1/medical/people/:id','GET','','',''),('p','删除居民角色','/api/v1/medical/people/list','GET','','',''),('p','新增居民角色','/api/v1/medical/people/add','GET','','',''),('p','新增居民角色','/api/v1/medical/people/list','GET','','',''),('p','超级管理员root','/api/v1/system/user/add','GET','','',''),('p','超级管理员root','/api/v1/system/user/list','GET','','',''),('p','超级管理员root','/api/v1/system/user/change','GET','','',''),('p','超级管理员root','/api/v1/system/user/:id','GET','','',''),('p','超级管理员root','/api/v1/system/role/add','GET','','',''),('p','超级管理员root','/api/v1/system/role/change','GET','','',''),('p','超级管理员root','/api/v1/system/role/list','GET','','',''),('p','超级管理员root','/api/v1/system/role/delete','GET','','',''),('p','超级管理员root','/api/v1/system/permission/add','GET','','',''),('p','超级管理员root','/api/v1/system/permission/change','GET','','',''),('p','超级管理员root','/api/v1/system/permission/:id','GET','','',''),('p','超级管理员root','/api/v1/medical/goods/put','GET','','',''),('p','超级管理员root','/api/v1/medical/goods/out','GET','','',''),('p','超级管理员root','/api/v1/medical/result/add','GET','','',''),('p','超级管理员root','/api/v1/medical/result/list','GET','','',''),('p','超级管理员root','/api/v1/medical/result/:id','GET','','',''),('p','超级管理员root','/api/v1/medical/community/add','GET','','',''),('p','超级管理员root','/api/v1/medical/community/list','GET','','',''),('p','超级管理员root','/api/v1/medical/community/:id','GET','','',''),('p','超级管理员root','/api/v1/medical/people/add','GET','','',''),('p','超级管理员root','/api/v1/medical/people/:id','GET','','',''),('p','超级管理员root','/api/v1/medical/people/list','GET','','',''),('p','超级管理员root','/api/v1/medical/push/add','GET','','',''),('p','超级管理员root','/api/v1/medical/push/change','GET','','',''),('p','超级管理员root','/api/v1/medical/push/list','GET','','',''),('p','超级管理员root','/api/v1/medical/api/v1/medical','GET','','',''),('p','超级管理员root','/api/v1/medical/push/:id','GET','','',''),('p','超级管理员root','/api/v1/medical/feedback/add','GET','','',''),('p','超级管理员root','/api/v1/medical/feedback/:id','GET','','',''),('p','超级管理员root','/api/v1/medical/feedback/list','GET','','',''),('p','推送管理角色','/api/v1/medical/push/add','GET','','',''),('p','推送管理角色','/api/v1/medical/push/change','GET','','',''),('p','推送管理角色','/api/v1/medical/push/list','GET','','',''),('p','推送管理角色','/api/v1/medical/push/:id','GET','','',''),('p','推送管理角色','/api/v1/medical/api/v1/medical','GET','','',''),('p','超级管理员root','/api/v1/system/permission/common','GET',NULL,NULL,NULL),('p','超级管理员root','/api/v1/system/permission/menu','GET',NULL,NULL,NULL),('p','超级管理员root','/api/v1/system/permission/tree','GET',NULL,NULL,NULL),('p','居民管理角色','/api/v1/medical/people/add','GET','','',''),('p','居民管理角色','/api/v1/medical/people/:id','GET','','',''),('p','居民管理角色','/api/v1/medical/people/list','GET','','','');

/*Table structure for table `community` */

DROP TABLE IF EXISTS `community`;

CREATE TABLE `community` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '社区编号',
  `POI_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'POI编号',
  `POI_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'POI名称',
  `POI_district` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所属区域',
  `POI_adcode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '区域编码',
  `POI_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '中心坐标',
  `POI_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '详细地址',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `community` */

insert  into `community`(`ID`,`POI_id`,`POI_name`,`POI_district`,`POI_adcode`,`POI_location`,`POI_address`,`created_at`,`updated_at`,`deleted_at`) values (1,'B0FFG6TQT3','美里湖街道社会治理综合服务中心','山东省济南市槐荫区','370104','116.931176,36.704761','新沙北路28号','2024-03-27 20:09:16','2024-03-27 20:09:16','2024-06-04 19:28:59'),(2,'B0GRXR7VM8','匡山街道办事处社区卫生服务中心','山东省济南市槐荫区','370104','116.947781,36.684456','济齐路112号','2024-03-27 20:25:30','2024-03-27 20:25:30',NULL),(3,'B0FFFOLNE9','鲍山花园','山东省济南市历城区','370112','117.191065,36.709074','鲍山街道鲍山社区(鲍山地铁站B口步行390米)','2024-03-27 20:27:18','2024-03-27 20:27:18',NULL),(4,'B0GUK9GRKY','山东科创新材料有限公司','山东省德州市齐河县','371425','116.803812,36.821475','经济开发区名嘉西路(中州电缆西门)','2024-03-27 20:27:57','2024-03-27 20:27:57','2024-06-10 23:59:41'),(5,'B0IB27UAN4','国贸花园A区','山东省济南市天桥区','370105','117.016928,36.688200','生产路与刘家桥街交叉口东北约200米','2024-03-27 20:29:04','2024-03-27 20:29:04',NULL),(6,'B0FFHM85EF','信用社宿舍','山东省济南市槐荫区','370104','116.911948,36.695810','济齐路附近','2024-03-27 20:34:10','2024-03-27 20:34:10',NULL),(7,'B021307DVQ','龙腾小区','山东省济南市槐荫区','370104','116.912659,36.694829','济齐路136号','2024-03-27 20:34:40','2024-03-27 20:34:40','2024-03-27 20:58:39'),(8,'53','广识面会起起','-','55','est sint dolore','贵州省沧州市西沙群岛','2024-04-21 15:17:54','2024-04-21 15:17:54','2024-04-21 15:19:02'),(9,'B0FFHIWAFJ','鲍山花园南区(南2门)','山东省济南市历城区','370112','117.190574,36.706277','鲍山街道凤鸣路鲍山花园南区','2024-06-04 19:56:37','2024-06-04 19:56:37',NULL),(10,'B0FFF03ZCB','舜泰广场3号楼','山东省济南市历城区','370112','117.147055,36.663803','舜华路2000号','2024-06-04 19:57:48','2024-06-04 19:57:48',NULL),(11,'B0FFHIWAFI','鲍山花园南区(北2门)','山东省济南市历城区','370112','117.190587,36.705907','鲍山街道凤鸣路鲍山花园南区','2024-06-04 20:01:45','2024-06-04 20:01:45',NULL),(12,'B0FFLHSHTD','哈尔滨师范大学松北校区学生四公寓','黑龙江省哈尔滨市呼兰区','230111','126.562921,45.865627','师大路1号','2024-06-04 20:18:31','2024-06-04 20:18:31',NULL),(13,'BV10885903','德州万达广场西门(公交站)','山东省德州市德城区','371402','116.309988,37.430676','11路;13路;36路;4路;68路;K916路B线','2024-06-04 20:18:51','2024-06-04 20:18:51',NULL),(14,'B0JUV552DI','高唐县博物馆','山东省聊城市高唐县','371526','116.231360,36.842541','高唐书画小镇5号楼','2024-06-04 20:19:04','2024-06-04 20:19:04',NULL),(15,'B0KGOZKT82','哈啰电动车出租(东大街店)','陕西省西安市碑林区','610103','108.960469,34.258769','马厂子59号(大差市地铁站F口步行190米)','2024-06-10 23:59:29','2024-06-10 23:59:29','2024-06-11 00:01:31'),(16,'B0FFG8QYFV','大连周水子国际机场航站楼','辽宁省大连市甘井子区','210211','121.542160,38.962519','迎客路100号','2024-06-11 00:02:34','2024-06-11 00:02:34',NULL);

/*Table structure for table `feedback` */

DROP TABLE IF EXISTS `feedback`;

CREATE TABLE `feedback` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '反馈编号',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '反馈内容',
  `status` int DEFAULT NULL COMMENT '反馈状态（0为未处理，1为已处理）',
  `people_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `feedback` */

insert  into `feedback`(`ID`,`content`,`status`,`people_id`,`created_at`,`updated_at`,`deleted_at`) values (1,'字太小 看不清',53,1,'2024-06-11 09:37:58','2024-06-11 09:37:58',NULL),(2,'minim dolore',21,NULL,'2024-06-11 09:38:03','2024-06-11 09:38:03','2024-06-11 09:39:08'),(3,'qui ex',49,NULL,'2024-06-11 09:38:04','2024-06-11 09:38:04','2024-06-11 11:22:44'),(4,'sit id ullamco eiusmod dolore',18,NULL,'2024-06-11 09:38:06','2024-06-11 09:38:06','2024-06-11 09:57:21'),(5,'xxxxxxx',1,4,'2024-06-11 11:24:33','2024-06-12 17:10:38',NULL);

/*Table structure for table `goods` */

DROP TABLE IF EXISTS `goods`;

CREATE TABLE `goods` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '物品id',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '物品名称',
  `num` int DEFAULT NULL COMMENT '物品数量',
  `uint` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '物品单位',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '物品备注',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `goods` */

insert  into `goods`(`ID`,`name`,`num`,`uint`,`remark`,`created_at`,`updated_at`,`deleted_at`) values (1,'利器盒',100,'个',NULL,NULL,NULL,'2024-03-17 17:05:07'),(2,'针管',201,'个',NULL,NULL,'2024-06-13 01:24:52',NULL),(3,'专南该到说想真',5,'个','nis','2024-04-21 15:32:53','2024-04-21 15:32:53','2024-04-21 15:33:28'),(4,'市前中标带影',-1661429002,'dolor','dolo','2024-04-25 21:11:30','2024-04-25 21:11:30','2024-04-25 21:16:30'),(5,'管何治业众育',-1001631005,'nisi','cillum reprehenderit est non','2024-04-25 21:11:34','2024-04-25 21:11:34','2024-04-25 21:16:40'),(8,'即就石素外料料',155,'se','Duis esse ipsum','2024-05-07 21:09:22','2024-05-07 21:09:22','2024-06-13 01:23:16'),(9,'指根且',2940,'a','culpa dolore in irure n','2024-05-07 21:09:32','2024-05-07 21:09:32','2024-06-10 22:57:14'),(10,'fsa',0,'saff','asfas','2024-06-10 23:07:20','2024-06-10 23:07:20','2024-06-13 01:23:17'),(11,'test',0,'个','testssssss','2024-06-10 23:08:53','2024-06-10 23:08:53','2024-06-13 01:23:19'),(12,'de',0,'1','123156165','2024-06-10 23:09:09','2024-06-10 23:09:09','2024-06-10 23:09:25'),(13,'aaa',20,'vvv','ccc','2024-06-10 23:30:48','2024-06-10 23:30:48','2024-06-13 01:23:20'),(14,'利器盒',51,'个','','2024-06-13 01:23:32','2024-06-13 01:24:58',NULL);

/*Table structure for table `people` */

DROP TABLE IF EXISTS `people`;

CREATE TABLE `people` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '居民编号',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '姓名',
  `sex` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '性别',
  `nation` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '民族',
  `birthday` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '出生日期',
  `idnumber` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身份证号',
  `email` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `age` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '年龄',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `people` */

insert  into `people`(`ID`,`name`,`sex`,`nation`,`birthday`,`idnumber`,`email`,`phone`,`age`,`created_at`,`updated_at`,`deleted_at`) values (1,'张三','女','汉族','2013-08-30','232103197804172580','e.ckbpj@qq.com','19883826343','67','2024-06-11 09:29:51','2024-06-11 09:29:51',NULL),(2,'第品写','女','dolore','2013-08-30','96','e.ckbpj@qq.com','19883826343','aut','2024-06-11 09:29:52','2024-06-11 09:29:52','2024-06-11 09:30:57'),(3,'as','女','sfa','','fasf','fasfa','asfas','','2024-06-11 10:51:23','2024-06-11 10:51:23','2024-06-11 10:57:06'),(4,'李四','男','汉族','2024-06-04','232103197804212580','asfasfas@126.com','15978452521','50','2024-06-11 10:56:59','2024-06-11 10:56:59',NULL),(5,'王五','男','汉族','2024-06-14','1','1','1','','2024-06-14 01:53:52','2024-06-14 01:53:52','2024-06-14 02:20:26'),(6,'1','男','1','2024-06-02','1','1','1','','2024-06-14 03:00:52','2024-06-14 03:00:52',NULL);

/*Table structure for table `push` */

DROP TABLE IF EXISTS `push`;

CREATE TABLE `push` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '推送标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '推送内容',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `push` */

insert  into `push`(`ID`,`title`,`content`,`created_at`,`updated_at`,`deleted_at`) values (1,'二金金关问就','laboris aliquipasfgibubbuuuuuuuuuuuuuuuuuuuuuuu','2024-06-10 21:58:00','2024-06-13 02:03:42',NULL),(2,'离书发个商极百','consequat culpa minim nulla','2024-06-10 21:58:46','2024-06-10 21:58:46','2024-06-10 21:59:09'),(3,'条亲况动全','laboris aliquipasfgibubbuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu','2024-06-10 21:58:48','2024-06-13 02:03:47',NULL),(4,'saf','laboris aliquipasfgibubbuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu','2024-06-10 23:50:25','2024-06-13 02:03:51',NULL),(5,'fsa','asfasfasfa','2024-06-10 23:51:31','2024-06-10 23:51:31','2024-06-10 23:52:09'),(6,'','','2024-06-11 09:29:26','2024-06-11 09:29:26','2024-06-13 01:27:33'),(7,'','','2024-06-11 09:29:32','2024-06-11 09:29:32','2024-06-13 01:27:35'),(8,'','','2024-06-11 09:29:34','2024-06-11 09:29:34','2024-06-13 01:27:37');

/*Table structure for table `relate_role_permission` */

DROP TABLE IF EXISTS `relate_role_permission`;

CREATE TABLE `relate_role_permission` (
  `sys_role_id` bigint NOT NULL COMMENT '角色编号',
  `sys_permission_id` bigint NOT NULL COMMENT '权限编号',
  PRIMARY KEY (`sys_role_id`,`sys_permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `relate_role_permission` */

insert  into `relate_role_permission`(`sys_role_id`,`sys_permission_id`) values (28,21),(28,22),(28,25),(28,26),(31,26),(33,34),(33,35),(34,33),(34,35),(35,21),(35,22),(35,25),(35,26),(35,27),(35,28),(35,29),(35,30),(35,31),(35,32),(35,33),(35,34),(35,35),(35,36),(35,37),(35,38),(35,39),(35,40),(35,41),(35,42),(35,43),(35,44),(35,45),(35,46),(35,47),(35,48),(35,49),(35,50),(35,51),(35,52),(35,53),(35,54),(35,55),(35,56),(35,57),(35,58),(36,47),(36,48),(36,49),(36,50),(36,51),(37,33),(37,34),(37,35);

/*Table structure for table `relate_user_role` */

DROP TABLE IF EXISTS `relate_user_role`;

CREATE TABLE `relate_user_role` (
  `sys_user_id` bigint NOT NULL,
  `sys_role_id` bigint NOT NULL,
  PRIMARY KEY (`sys_user_id`,`sys_role_id`) USING BTREE,
  KEY `relate_user_role_sys_role_ID_fk` (`sys_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `relate_user_role` */

insert  into `relate_user_role`(`sys_user_id`,`sys_role_id`) values (12,0),(21,0),(22,0),(25,0),(10,21),(0,22),(17,22),(18,22),(25,22),(12,25),(21,26),(22,26),(25,26),(35,33),(34,34),(36,35),(0,36),(37,36);

/*Table structure for table `result` */

DROP TABLE IF EXISTS `result`;

CREATE TABLE `result` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '结果编号',
  `sfz_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '姓名',
  `sfz_sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '性别',
  `sfz_nation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '民族',
  `sfz_birthday` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '出生日期',
  `sfz_idnumber` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身份证号',
  `sfz_age` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '年龄',
  `hw_height` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身高',
  `hw_weight` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '体重',
  `hw_bmi` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'BMI',
  `fat_zflv` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '脂肪率',
  `fat_jcdx` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '基础代谢',
  `fat_tsfl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '体水份量',
  `fat_tsflv` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '体水份率',
  `fat_zfl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '脂肪量',
  `fat_jrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '肌肉量',
  `fat_jrlv` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '肌肉率',
  `fat_gy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '骨盐',
  `fat_qztz` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '去脂体重',
  `fat_dbzlv` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '蛋白质率',
  `fat_xbnyl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '细胞内液量',
  `fat_xbwyl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '细胞外液量',
  `fat_xbnylv` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '细胞内液率',
  `fat_xbwylv` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '细胞外液率',
  `fat_dbz` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '蛋白质',
  `fat_nzzf` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '内脏脂肪等级',
  `fat_gl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '骨量',
  `blood_high` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '高压',
  `blood_low` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '低压',
  `blood_rate` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '心率',
  `blood_rhigh` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '右高压(血压双臂时发送)',
  `blood_rlow` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '右低压(血压双臂时发送)',
  `spo2_sp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '血氧',
  `tiwen` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '体温',
  `ecg_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '心电结果',
  `ecg_data` longblob COMMENT '心电波形图片数据',
  `ecg_len` bigint DEFAULT NULL COMMENT '心电波形图片数据长度',
  `xt_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '血糖测量类型',
  `xt_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '血糖值',
  `ns` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '尿酸值',
  `dgc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '胆固醇值',
  `zybs` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '中医体质类型',
  `shili_left_eye` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '左眼视力',
  `shili_right_eye` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '右眼视力',
  `semang` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '色盲结果',
  `xlcp_ucla` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'UCLA孤独量表得分',
  `xlcp_lnyy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '老年抑郁量表得分',
  `xlcp_zpyy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '自评抑郁量表得分',
  `xlcp_hmdjl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '汉密顿焦虑量表得分',
  `xlcp_qxjkd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '情绪健康度测试得分',
  `xlcp_zcjkpd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '自测健康评定量表得分',
  `xlcp_shmyd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '生活满意度量表得分',
  `xlcp_rgza` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '人格障碍性格测试得分',
  `xlcp_pstr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'PSTR成人心理压力测试得分',
  `xlcp_hfxx` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '哈佛性向测试得分',
  `xlcp_eq` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '情商(EQ)测试得分',
  `xlcp_smzkpg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '睡眠状况评估得分',
  `deviceID` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备码',
  `examNo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '体检编号',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `result` */

insert  into `result`(`ID`,`sfz_name`,`sfz_sex`,`sfz_nation`,`sfz_birthday`,`sfz_idnumber`,`sfz_age`,`hw_height`,`hw_weight`,`hw_bmi`,`fat_zflv`,`fat_jcdx`,`fat_tsfl`,`fat_tsflv`,`fat_zfl`,`fat_jrl`,`fat_jrlv`,`fat_gy`,`fat_qztz`,`fat_dbzlv`,`fat_xbnyl`,`fat_xbwyl`,`fat_xbnylv`,`fat_xbwylv`,`fat_dbz`,`fat_nzzf`,`fat_gl`,`blood_high`,`blood_low`,`blood_rate`,`blood_rhigh`,`blood_rlow`,`spo2_sp`,`tiwen`,`ecg_result`,`ecg_data`,`ecg_len`,`xt_type`,`xt_value`,`ns`,`dgc`,`zybs`,`shili_left_eye`,`shili_right_eye`,`semang`,`xlcp_ucla`,`xlcp_lnyy`,`xlcp_zpyy`,`xlcp_hmdjl`,`xlcp_qxjkd`,`xlcp_zcjkpd`,`xlcp_shmyd`,`xlcp_rgza`,`xlcp_pstr`,`xlcp_hfxx`,`xlcp_eq`,`xlcp_smzkpg`,`deviceID`,`examNo`,`created_at`,`updated_at`,`deleted_at`) values (15,'李明','男','汉族','1991-08-02','232120197705152512','32','174','75','14.7','15','1395','35','50','7','42','66','3.8','30','16','20','17','35','21','14','5','18','100','60','80','110','70','98','36.8','窦性心律不齐','est dolore in',84,'0','4.2','0.252','4.1','气虚质','4.5','4.7','色盲','36','5','10','5','15','385','19','7','52','85','128','18','73','20240606063152','2024-06-06 06:31:16','2024-06-06 06:31:16',NULL),(16,'张红','女','汉族','1991-08-02','232120197705152512','32','174','75','14.7','15','1395','35','50','7','42','66','3.8','30','16','20','17','35','21','14','5','18','100','60','80','110','70','98','36.8','窦性心律不齐','est dolore in',84,'0','4.2','0.252','4.1','气虚质','4.5','4.7','色盲','36','5','10','5','15','385','19','7','52','85','128','18','73','20240606063252','2024-06-06 06:31:34','2024-06-06 06:31:34',NULL);

/*Table structure for table `sys_captcha` */

DROP TABLE IF EXISTS `sys_captcha`;

CREATE TABLE `sys_captcha` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '验证码编号',
  `verifyCode` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '验证码内容',
  `verifyKey` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '验证码标识',
  `verifyImg` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '验证码图片（base64）',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_captcha` */

insert  into `sys_captcha`(`ID`,`verifyCode`,`verifyKey`,`verifyImg`) values (1,NULL,'quis','http://dummyimage.com/400x400'),(2,NULL,'dolor sunt','http://dummyimage.com/400x400'),(4,NULL,'veniam non d','http://dummyimage.com/400x400');

/*Table structure for table `sys_permission` */

DROP TABLE IF EXISTS `sys_permission`;

CREATE TABLE `sys_permission` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单编号',
  `parentId` bigint DEFAULT NULL COMMENT '父菜单编号',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单图标',
  `type` int DEFAULT NULL COMMENT '菜单类型（0为目录，1为菜单，2为功能）',
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '功能的uri',
  `sort` int DEFAULT NULL COMMENT '菜单排序',
  `router` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单的路由',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单的组件路径',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_permission` */

insert  into `sys_permission`(`ID`,`parentId`,`name`,`icon`,`type`,`uri`,`sort`,`router`,`component`,`created_at`,`updated_at`,`deleted_at`) values (10,NULL,'系统',NULL,0,NULL,NULL,NULL,'',NULL,NULL,NULL),(11,NULL,'医疗管理',NULL,0,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(12,10,'用户管理',NULL,1,NULL,NULL,'/user','system/UserManage',NULL,NULL,NULL),(13,10,'角色管理',NULL,1,NULL,NULL,'/role','system/RoleManage',NULL,NULL,NULL),(14,11,'库存管理',NULL,1,NULL,NULL,'/good','medical/GoodManage',NULL,NULL,NULL),(15,11,'检测结果',NULL,1,NULL,NULL,'/result','medical/ResultManage',NULL,NULL,NULL),(16,10,'权限管理',NULL,1,NULL,NULL,'/permission','system/PermissionManage',NULL,NULL,NULL),(17,11,'社区管理',NULL,1,NULL,NULL,'/community','medical/CommunityManage',NULL,NULL,NULL),(18,11,'居民管理',NULL,1,NULL,NULL,'/people','medical/PeopleManage',NULL,NULL,NULL),(19,11,'推送管理',NULL,1,NULL,NULL,'/message','medical/PushMessage',NULL,NULL,NULL),(20,11,'反馈管理',NULL,1,NULL,NULL,'/feedback','medical/feedBackManage',NULL,NULL,NULL),(21,12,'获取用户列表',NULL,2,'/api/v1/system/user/list',NULL,NULL,NULL,NULL,NULL,NULL),(22,12,'新增用户',NULL,2,'/api/v1/system/user/add',NULL,NULL,NULL,NULL,'2024-06-12 15:04:01',NULL),(25,12,'修改用户',NULL,2,'/api/v1/system/user/change',NULL,NULL,NULL,NULL,NULL,NULL),(26,12,'删除用户',NULL,2,'/api/v1/system/user/:id',NULL,NULL,NULL,NULL,NULL,NULL),(27,13,'新增角色',NULL,2,'/api/v1/system/role/add',NULL,NULL,NULL,NULL,NULL,NULL),(28,13,'修改角色',NULL,2,'/api/v1/system/role/change',NULL,NULL,NULL,NULL,NULL,NULL),(29,13,'查询角色列表',NULL,2,'/api/v1/system/role/list',NULL,NULL,NULL,NULL,NULL,NULL),(30,13,'删除角色',NULL,2,'/api/v1/system/role/delete',NULL,NULL,NULL,NULL,NULL,NULL),(31,14,'入库',NULL,2,'/api/v1/medical/goods/put',NULL,NULL,NULL,NULL,NULL,NULL),(32,14,'出库',NULL,2,'/api/v1/medical/goods/out',NULL,NULL,NULL,NULL,NULL,NULL),(33,18,'新增居民',NULL,2,'/api/v1/medical/people/add',NULL,NULL,'','2024-06-14 01:57:22','2024-06-14 01:57:22',NULL),(34,18,'删除居民','',2,'/api/v1/medical/people/:id',0,NULL,'','2024-06-14 02:07:41','2024-06-14 02:08:11',NULL),(35,18,'获取居民信息','',2,'/api/v1/medical/people/list',0,'/api/v1/medical/people/list','','2024-06-14 02:15:58','2024-06-14 02:16:11',NULL),(36,16,'新增权限','',2,'/api/v1/system/permission/add',0,'','','2024-06-14 09:33:39','2024-06-14 09:33:39',NULL),(37,16,'根据id获取权限','',2,'/api/v1/system/permission/:id',0,'','','2024-06-14 09:34:37','2024-06-14 09:34:37',NULL),(38,16,'更新权限','',2,'/api/v1/system/permission/change',0,'','','2024-06-14 09:35:27','2024-06-14 09:35:27',NULL),(39,16,'删除权限','',2,'/api/v1/system/permission/:id',0,'','','2024-06-14 09:36:19','2024-06-14 09:36:19',NULL),(40,15,'新增检测结果','',2,'/api/v1/medical/result/add',0,'','','2024-06-14 09:38:11','2024-06-14 09:38:11',NULL),(41,15,'获取检测结果列表','',2,'/api/v1/medical/result/list',0,'','','2024-06-14 09:38:58','2024-06-14 09:38:58',NULL),(42,15,'根据id获取检测结果','',2,'/api/v1/medical/result/:id',0,'','','2024-06-14 09:40:24','2024-06-14 09:40:24',NULL),(43,15,'删除检测结果','',2,'/api/v1/medical/result/:id',0,'','','2024-06-14 09:41:16','2024-06-14 09:41:16',NULL),(44,17,'添加社区管理','',2,'/api/v1/medical/community/add',0,'','','2024-06-14 09:42:30','2024-06-14 09:42:30',NULL),(45,17,'获取社区管理列表','',2,'/api/v1/medical/community/list',0,'','','2024-06-14 09:43:17','2024-06-14 09:43:17',NULL),(46,17,'删除社区管理','',2,'/api/v1/medical/community/:id',0,'','','2024-06-14 09:43:53','2024-06-14 09:43:53',NULL),(47,19,'新增推送','',2,'/api/v1/medical/push/add',0,'','','2024-06-14 09:45:31','2024-06-14 09:45:31',NULL),(48,19,'更新推送','',2,'/api/v1/medical/push/change',0,'','','2024-06-14 09:47:20','2024-06-14 09:47:20',NULL),(49,19,'获取推送列表','',2,'/api/v1/medical/push/list',0,'','','2024-06-14 09:48:08','2024-06-14 09:48:08',NULL),(50,19,'根据推送id获取详情','',2,'/api/v1/medical/push/:id',0,'','','2024-06-14 09:49:21','2024-06-14 09:49:21',NULL),(51,19,'删除推送','',2,'/api/v1/medical/api/v1/medical',0,'','','2024-06-14 09:49:48','2024-06-14 09:49:48',NULL),(52,20,'新建反馈','',2,'/api/v1/medical/feedback/add',0,'','','2024-06-14 11:03:55','2024-06-14 11:03:55',NULL),(53,20,'更新反馈','',2,'/api/v1/medical/feedback/:id',0,'','','2024-06-14 11:04:52','2024-06-14 11:04:52',NULL),(54,20,'获取反馈列表','',2,'/api/v1/medical/feedback/list',0,'','','2024-06-14 11:05:31','2024-06-14 11:05:31',NULL),(55,20,'删除反馈','',2,'/api/v1/medical/feedback/:id',0,'','','2024-06-14 11:06:03','2024-06-14 11:06:03',NULL),(56,16,'查询普通列表',NULL,2,'/api/v1/system/permission/common',0,NULL,NULL,'2024-06-14 21:58:43','2024-06-14 21:58:45',NULL),(57,16,'查询树状菜单列表',NULL,2,'/api/v1/system/permission/menu',0,NULL,NULL,'2024-06-14 22:01:04','2024-06-14 22:01:06',NULL),(58,16,'查询树状功能列表',NULL,2,'/api/v1/system/permission/tree',0,NULL,NULL,'2024-06-14 22:01:04','2024-06-14 22:01:06',NULL),(59,15,'测试修改','',2,'/ceshi2',0,'','','2024-06-15 11:43:19','2024-06-15 11:45:10','2024-06-15 11:45:18');

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '角色编号',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名称',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色描述',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_role` */

insert  into `sys_role`(`ID`,`name`,`desc`,`created_at`,`updated_at`,`deleted_at`) values (21,'超级管理员','拥有最高权限','2024-06-13 01:09:50',NULL,NULL),(22,'测试角色','查看结果权限','2024-06-13 01:10:27',NULL,NULL),(28,'用户管理角色','用户管理','2024-06-13 18:00:05','2024-06-13 18:00:05',NULL),(29,'测试角色','测试','2024-06-13 18:14:46','2024-06-13 18:14:46','2024-06-13 18:37:51'),(30,'超级管理员2','root','2024-06-13 20:29:34','2024-06-13 20:29:34','2024-06-14 11:02:12'),(31,'删除用户角色333','333','2024-06-13 23:28:44','2024-06-13 23:28:44',NULL),(32,'删除居民角色','1','2024-06-14 02:10:11','2024-06-14 02:10:11','2024-06-14 02:14:06'),(33,'删除居民角色','a','2024-06-14 02:18:11','2024-06-14 02:18:11',NULL),(34,'新增居民角色','a','2024-06-14 02:18:36','2024-06-14 02:18:36',NULL),(35,'超级管理员root','最高权限','2024-06-14 11:09:00','2024-06-14 11:09:00',NULL),(36,'推送管理角色','推送管理。。。','2024-06-14 11:14:54','2024-06-14 11:14:54',NULL),(37,'居民管理角色','居民管理','2024-06-14 23:02:23','2024-06-14 23:02:23',NULL);

/*Table structure for table `sys_user` */

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户姓名',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户密码',
  `sex` int DEFAULT NULL COMMENT '用户性别',
  `birthday` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户生日',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_user` */

insert  into `sys_user`(`ID`,`name`,`username`,`password`,`sex`,`birthday`,`phone`,`email`,`created_at`,`updated_at`,`deleted_at`) values (10,'超级管理员','admin','$2a$10$WjdE7lCdUX.HPvApKAqelOMJzcCQzpWiPlGHycBTHehqzbJn9P1Ae',1,'1986-01-21','18137846437','218516354@qq.com','2024-04-07 19:50:20','2024-04-07 19:50:20',NULL),(17,'测试用户','test','$2a$10$WjdE7lCdUX.HPvApKAqelOMJzcCQzpWiPlGHycBTHehqzbJn9P1Ae',0,'2001-10-25','18845125124','6749555@qq.com','2024-05-28 01:14:15',NULL,NULL),(18,'社区管理','community','$2a$10$WjdE7lCdUX.HPvApKAqelOMJzcCQzpWiPlGHycBTHehqzbJn9P1Ae',1,'1988-01-01','15945125478','546816@qq.com','2024-06-20 01:16:02',NULL,NULL),(19,'测试','slh','$2a$10$WjdE7lCdUX.HPvApKAqelOMJzcCQzpWiPlGHycBTHehqzbJn9P1Ae',1,'2001-09-09','18348563173','67490009@qq.com','2024-06-04 01:16:46',NULL,'2024-06-13 17:23:43'),(20,'','','$2a$10$eC8znZC4scS9aC0V7ySKuOCgQkvOuk1DyFfjJRLvV.zPJnCGLSxMq',0,'','','','2024-06-13 01:17:35','2024-06-13 01:17:35','2024-06-13 01:17:39'),(21,'张三','zhangsan','$2a$10$9qEqEqgwsYG1LPTKPsF3b.Emp93S2Q7AhEfAj.NHe.l.qHoBOyCz2',1,'','11122223333','11122223333@163.com','2024-06-13 15:12:09','2024-06-13 15:12:09','2024-06-13 15:26:58'),(23,'李四','lisi','$2a$10$CIZWr0bsI1BpK18R/GO63ecvzcQ5xOfq22ZkwAbmQjGpbiHkaR3QO',1,'','1','1','2024-06-13 17:04:08','2024-06-13 17:04:08','2024-06-14 02:09:37'),(24,'张三','zhangsan','$2a$10$MOZ4.zPQWSelyHg4tUIhqeT0kaLFZlWIqwmQ9MIYoO/.QQrrf72u6',1,'','1','1','2024-06-13 17:54:44','2024-06-13 17:54:44','2024-06-14 00:54:20'),(26,'测试','ceshi','$2a$10$CNFNcTZD6N0Fp2t7Q5b/rexLH6HLz.ySXQ3V0YKQhP0KnPaWGBxye',1,'','1','1','2024-06-13 19:05:15','2024-06-13 19:05:15','2024-06-13 19:11:01'),(27,'超级管理员root','root','$2a$10$rzvJsumFjg3y29.22vz5HuH0zDVGibb4wZO7scka8Bmt6JQ2zDTMS',1,'','1','1','2024-06-13 20:29:58','2024-06-13 20:29:58','2024-06-14 00:54:25'),(29,'666','666','$2a$10$ibQjhQrdGPT6CoXaO.pZ6ejUZqKycpPZI1xAT88zplXOrtEQRkcua',1,'','1','1','2024-06-14 00:47:04','2024-06-14 00:47:04','2024-06-14 00:49:46'),(30,'666','666','$2a$10$y7nC9SAsiOJDNso7g/FVFeCRrvvAHlWxeCDlK9GA7FVQbf5QKK886',1,'','6','6','2024-06-14 00:52:09','2024-06-14 00:52:09','2024-06-14 02:09:34'),(31,'root','root','$2a$10$9kqdbT.p0OxC0LatWoPQSO7ytCY5ZyAW.hUKP54vHZ36P3WEJyjDq',1,'','t','t','2024-06-14 00:54:48','2024-06-14 00:54:48','2024-06-14 11:06:41'),(32,'777','777','$2a$10$9Mqqjzy9k7DfVxB2gyKVleirTbA5vfPJHIyFfJu.H5hUcL0gll4P6',1,'','7','7','2024-06-14 01:33:53','2024-06-14 01:33:53','2024-06-14 02:09:30'),(33,'删除居民用户','666','$2a$10$v.97423pQuz.kQ0KnV3CueBySkW85VJ1tpaqEou9CI2Y194kbOGpi',1,'','6','6','2024-06-14 02:10:38','2024-06-14 02:10:38','2024-06-14 02:19:21'),(34,'3','3','$2a$10$3.9DEhGGfhws4202GuqwYO3cD1r9N.2N.688HGNsOp6nkz1ns7RIS',1,'','3','3','2024-06-14 02:19:17','2024-06-14 02:19:17',NULL),(35,'4','4','$2a$10$Ge6r4.rPngPPAqidspuoB.GacZJSXxRtTHS3BMVaGS8KnTJi4Uyyu',1,'','4','4','2024-06-14 02:19:41','2024-06-14 02:19:41',NULL),(36,'超级管理员','root','$2a$10$4MGLvo4g9lWndiHK1BxrYevsIFuW/kcbVrWnc98SiNKHD3qMRMy4S',1,'','1','1','2024-06-14 11:11:39','2024-06-14 11:11:39',NULL),(37,'推送管理员','1','$2a$10$C1smR2LOu/Dn4GqPGmnQLub6rUffBw3ne19Mn9q2uf2iKWVJgfR8K',1,'','1','1','2024-06-14 11:15:29','2024-06-14 11:15:29',NULL),(38,'测试用户修改6','6','',1,'2024-06-14','6','6','2024-06-14 23:13:00','2024-06-14 23:51:18','2024-06-15 00:16:34');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
