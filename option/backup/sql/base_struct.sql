
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `mini_sick_000` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `mini_sick_000`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cycle` (
  `userid` char(15) NOT NULL,
  `cycle_seq` int(11) NOT NULL,
  `name` varchar(10) NOT NULL,
  `out_hospital_time` char(17) DEFAULT '',
  `follow_over` char(1) DEFAULT '2',
  `date` char(10) NOT NULL,
  `time` char(5) NOT NULL,
  `timestamp` char(13) NOT NULL,
  PRIMARY KEY (`userid`,`cycle_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `follow` (
  `userid` char(15) NOT NULL,
  `cycle_seq` int(11) NOT NULL DEFAULT '1',
  `follow_seq` int(11) NOT NULL,
  `hight_risk` char(3) NOT NULL,
  `emesis_grade` char(3) NOT NULL,
  `nausea_grade` char(3) NOT NULL,
  `out_content` varchar(20) NOT NULL,
  `out_content_diy` varchar(100) DEFAULT NULL,
  `follow_over` char(3) NOT NULL,
  `satisfaction_1` char(3) DEFAULT NULL,
  `satisfaction_2` char(3) DEFAULT NULL,
  `satisfaction_3` char(3) DEFAULT NULL,
  `satisfaction_4` char(3) DEFAULT NULL,
  `satisfaction_5` char(3) DEFAULT NULL,
  `satisfaction_total` char(3) DEFAULT NULL,
  `writer` varchar(10) NOT NULL,
  `follow_follow_date` char(10) NOT NULL,
  `follow_follow_time` char(5) NOT NULL,
  `follow_follow_timestamp` char(13) NOT NULL,
  PRIMARY KEY (`userid`,`cycle_seq`,`follow_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `nurse` (
  `userid` char(15) NOT NULL,
  `cycle_seq` int(11) NOT NULL,
  `nurse_seq` int(11) NOT NULL,
  `nausea_assessment` char(1) NOT NULL,
  `emesis_assessment` char(1) NOT NULL,
  `measure` varchar(30) NOT NULL,
  `comment` varchar(50) DEFAULT NULL,
  `out_hospital` varchar(3) DEFAULT NULL,
  `writer` varchar(10) NOT NULL,
  `assessment_date` char(10) NOT NULL,
  `assessment_time` char(5) NOT NULL,
  `assessment_timestamp` char(13) NOT NULL,
  PRIMARY KEY (`userid`,`cycle_seq`,`nurse_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `risk` (
  `userid` char(15) NOT NULL,
  `cycle_seq` int(11) NOT NULL,
  `program` varchar(30) DEFAULT '',
  `not_medication` varchar(20) DEFAULT '',
  `medication` varchar(2) DEFAULT '',
  `grand` varchar(2) DEFAULT '',
  `pre_program` varchar(20) DEFAULT '',
  `pre_program_diy` varchar(50) DEFAULT NULL,
  `comment` varchar(20) DEFAULT NULL,
  `comment_diy` varchar(50) DEFAULT NULL,
  `need_nurse` char(1) DEFAULT '1',
  `writer` varchar(10) NOT NULL,
  `assessment_date` char(10) NOT NULL,
  `assessment_time` char(5) NOT NULL,
  `assessment_timestamp` char(13) NOT NULL,
  `chemotherapy_date` char(10) NOT NULL,
  `chemotherapy_time` char(5) NOT NULL,
  `chemotherapy_timestamp` char(13) NOT NULL,
  PRIMARY KEY (`userid`,`cycle_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sicker` (
  `userid` char(15) NOT NULL,
  `name` varchar(6) DEFAULT NULL,
  `age` char(3) DEFAULT NULL,
  `gender` char(2) DEFAULT NULL,
  `telphone` char(11) DEFAULT NULL,
  `hospital_number` varchar(12) DEFAULT NULL,
  `attandance_number` varchar(12) DEFAULT NULL,
  `disease` varchar(20) DEFAULT NULL,
  `know` char(11) DEFAULT '0',
  `cycle_seq` int(11) DEFAULT '0',
  `writer` varchar(10) NOT NULL,
  `write_data` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `account` varchar(15) NOT NULL,
  `name` varchar(10) DEFAULT '',
  `password` varchar(20) NOT NULL,
  `type` tinyint(4) NOT NULL,
  PRIMARY KEY (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
