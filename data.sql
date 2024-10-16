-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: skool_saver
-- ------------------------------------------------------
-- Server version	8.0.39-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `consultations`
--

DROP TABLE IF EXISTS `consultations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consultations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fullname` varchar(1000) DEFAULT NULL,
  `email` varchar(1000) DEFAULT NULL,
  `mobilenumber` varchar(50) DEFAULT NULL,
  `service` varchar(1000) DEFAULT NULL,
  `pricerange` varchar(1000) DEFAULT NULL,
  `location` varchar(1000) DEFAULT NULL,
  `aboutproject` varchar(5000) DEFAULT NULL,
  `deleteflag` tinyint DEFAULT '0',
  `upserteddateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `lastupdateddateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consultations`
--

LOCK TABLES `consultations` WRITE;
/*!40000 ALTER TABLE `consultations` DISABLE KEYS */;
INSERT INTO `consultations` VALUES (1,'narendra modi','modi@gmail.com','696969696969','Web page design','10000rs-20000rs','Gujarat','Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of de Finibus Bonorum et Malorum',0,'2024-09-28 04:19:24','2024-09-28 04:19:24');
/*!40000 ALTER TABLE `consultations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobappliedcandidates`
--

DROP TABLE IF EXISTS `jobappliedcandidates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobappliedcandidates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `candidatename` varchar(1000) DEFAULT NULL,
  `email` varchar(1000) DEFAULT NULL,
  `mobno` varchar(50) DEFAULT NULL,
  `current_emplyeement` varchar(1000) DEFAULT NULL,
  `location` varchar(1000) DEFAULT NULL,
  `skills` varchar(4000) DEFAULT NULL,
  `yearofexp` varchar(1000) DEFAULT NULL,
  `blobname` varchar(1000) DEFAULT NULL,
  `ext` varchar(100) DEFAULT NULL,
  `filename` varchar(1000) DEFAULT NULL,
  `deleteflag` tinyint DEFAULT '0',
  `applieddateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `lastupdatedateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `jobid` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `job_id_fk_idx` (`jobid`),
  CONSTRAINT `job_id_fk` FOREIGN KEY (`jobid`) REFERENCES `jobs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobappliedcandidates`
--

LOCK TABLES `jobappliedcandidates` WRITE;
/*!40000 ALTER TABLE `jobappliedcandidates` DISABLE KEYS */;
INSERT INTO `jobappliedcandidates` VALUES (1,'vinayak','abc@gmail.com','1234567890','Fresher','banglore','java,python,mysql','fresher','iFIX-API (1)_20240925.pdf','pdf','iFIX-API (1)',1,'2024-09-25 18:23:24','2024-09-25 18:46:49',NULL),(2,'vinayak','abc@gmail.com','1234567890','Fresher','banglore','java,python,mysql','fresher','iFIX-API (1)_20240926.pdf','pdf','iFIX-API (1)',1,'2024-09-25 18:35:53','2024-09-25 18:47:00',NULL),(3,'vinayak','abc@gmail.com','1234567890','Fresher','banglore','java,python,mysql','fresher','iFIX-API (1)_20240926.pdf','pdf','iFIX-API (1)',1,'2024-09-25 18:37:01','2024-09-28 09:23:15',NULL),(4,'vinayak','abc@gmail.com','1234567890','Fresher','banglore','java,python,mysql','fresher','medic_20240928.pdf','pdf','medic',1,'2024-09-28 04:11:30','2024-09-28 09:23:15',NULL),(5,'vinayak','abc@gmail.com','1234567890','Fresher','banglore','java,python,mysql','fresher','medic_20240928.pdf','pdf','medic',0,'2024-09-28 09:25:29','2024-09-28 09:25:29',1);
/*!40000 ALTER TABLE `jobappliedcandidates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `designation` varchar(2000) DEFAULT NULL,
  `experience` varchar(1000) DEFAULT NULL,
  `location` varchar(1000) DEFAULT NULL,
  `posteddateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `lastupdatedateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleteflag` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
INSERT INTO `jobs` VALUES (1,'Sr.Developer','6','banglore','2024-09-28 09:13:50','2024-09-28 09:13:50',0),(2,'Devops Engineer','3','banglore','2024-09-28 09:15:03','2024-09-28 09:15:03',0);
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `skoolsaverusers`
--

DROP TABLE IF EXISTS `skoolsaverusers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `skoolsaverusers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(4000) DEFAULT NULL,
  `password` varchar(4000) DEFAULT NULL,
  `logintime` varchar(1000) DEFAULT NULL,
  `logouttime` varchar(1000) DEFAULT NULL,
  `deleteflag` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `skoolsaverusers`
--

LOCK TABLES `skoolsaverusers` WRITE;
/*!40000 ALTER TABLE `skoolsaverusers` DISABLE KEYS */;
INSERT INTO `skoolsaverusers` VALUES (1,'admin','YWRtaW4=','2024-09-28 09:38:53',NULL,0);
/*!40000 ALTER TABLE `skoolsaverusers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userdetails`
--

DROP TABLE IF EXISTS `userdetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userdetails` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(1000) DEFAULT NULL,
  `user_email` varchar(1000) DEFAULT NULL,
  `user_mobileno` varchar(1000) DEFAULT NULL,
  `user_message` varchar(5000) DEFAULT NULL,
  `added_date_and_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `lastupdateddateandtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleteflag` tinyint DEFAULT '0',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userdetails`
--

LOCK TABLES `userdetails` WRITE;
/*!40000 ALTER TABLE `userdetails` DISABLE KEYS */;
INSERT INTO `userdetails` VALUES (1,'ajay','ajay@gmail.com','9089902345','What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.','2024-09-19 18:49:50','2024-09-19 18:49:50',0),(2,'ajay','ajay@gmail.com','9089902345','What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.','2024-09-20 14:21:44','2024-09-20 14:21:44',0),(3,'ajay','ajay@gmail.com','9089902345','What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.','2024-09-20 14:24:09','2024-09-20 14:24:09',0);
/*!40000 ALTER TABLE `userdetails` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-10 16:42:32
