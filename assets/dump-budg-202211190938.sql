-- MySQL dump 10.13  Distrib 8.0.31, for macos12.6 (x86_64)
--
-- Host: nicocartalla.com    Database: budg
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Budget`
--

DROP TABLE IF EXISTS `Budget`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Budget` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `user_id` int NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `current_budget` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `Budget_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Budget`
--

LOCK TABLES `Budget` WRITE;
/*!40000 ALTER TABLE `Budget` DISABLE KEYS */;
INSERT INTO `Budget` VALUES (1,'Enero',2,1000.00,'2018-01-01 00:00:00','2022-02-01 00:00:00',1);
/*!40000 ALTER TABLE `Budget` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Category`
--

DROP TABLE IF EXISTS `Category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `Category_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Category`
--

LOCK TABLES `Category` WRITE;
/*!40000 ALTER TABLE `Category` DISABLE KEYS */;
INSERT INTO `Category` VALUES (1,'Income',1),(2,'Food',1),(3,'Transportation',1),(4,'Enterntainment',1),(5,'Utilities',1),(6,'Healthcare',1),(7,'Clothing',1),(8,'Education',1),(9,'Savings',1),(10,'Personal Spendings',1),(11,'Miscellaneous',1),(12,'Enterntainment',1),(13,'Utilities',1),(14,'Healthcare',1),(15,'Clothing',1),(16,'Education',1),(17,'Savings',1),(18,'Personal Spendings',1),(19,'Miscellaneous',1);
/*!40000 ALTER TABLE `Category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Expense`
--

DROP TABLE IF EXISTS `User_transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User_transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `budget_id` int NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `description` varchar(255) NOT NULL,
  `category_id` int NOT NULL,
  `date` datetime NOT NULL,
  `type` varchar(10) NOT NULL DEFAULT 'expense',
  `filepath` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `budget_id` (`budget_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `User_transaction` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`),
  CONSTRAINT `User_transaction_ibfk_2` FOREIGN KEY (`budget_id`) REFERENCES `Budget` (`id`),
  CONSTRAINT `User_transaction_ibfk_3` FOREIGN KEY (`category_id`) REFERENCES `Category` (`id`),
  CONSTRAINT `User_transaction_type_check` CHECK (`type` IN ('expense', 'income'))
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User_transaction`
--

LOCK TABLES `User_transaction` WRITE;
/*!40000 ALTER TABLE `User_transaction` DISABLE KEYS */;
INSERT INTO `User_transaction` VALUES (1,2,1,105.00,'Groceries',2,'2022-01-01 00:00:00'),(2,2,1,1.00,'Alfajor',2,'2022-01-01 00:00:00'),(3,2,1,5.00,'Cerveza',2,'2022-01-01 00:00:00'),(4,2,1,50.00,'Vodka',2,'2022-01-01 00:00:00'),(5,2,1,100.00,'Tshirt',7,'2022-01-01 00:00:00'),(6,2,1,150.00,'Jeans',7,'2022-01-01 00:00:00'),(8,2,1,23.40,'Dinner with friends',4,'2022-11-14 00:00:00'),(9,2,1,31200.00,'Rent',5,'2022-07-04 00:00:00'),(10,2,1,23000.00,'UCU',8,'2022-11-24 00:00:00'),(11,2,1,1250.00,'What I could manage to save',9,'2022-11-24 00:00:00'),(12,2,1,1700.00,'Health insurance',8,'2022-11-20 00:00:00'),(13,2,1,300.00,'Cinema',4,'2022-10-16 00:00:00'),(14,2,1,17000.00,'Sold NFT',1,'2022-11-06 00:00:00'),(15,2,1,600.00,'Bus tickets',8,'2022-11-24 00:00:00');
/*!40000 ALTER TABLE `User_transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `last_login` datetime DEFAULT NULL,
  `active` tinyint(1) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,'admin','admin','admin@admin','admin@admin.com','$2a$09$0BxHCT2cE/V3JurhuJQKM.vN4FrFKExYldvmvBWLpJGSGTULJO2iS','2022-01-01 00:00:00',1,'https://ui-avatars.com/api/?name=Admin+Admin?length=2'),(2,'John','Smith','johnsmith','johnsmith@gmail.com','$2a$09$0BxHCT2cE/V3JurhuJQKM.vN4FrFKExYldvmvBWLpJGSGTULJO2iS','2022-01-01 00:00:00',1,'https://ui-avatars.com/api/?name=John+Smith?length=2');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'budg'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-11-19  9:38:50
