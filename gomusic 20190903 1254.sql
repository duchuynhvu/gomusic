-- MySQL Administrator dump 1.4
--
-- ------------------------------------------------------
-- Server version	5.5.5-10.3.16-MariaDB


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


--
-- Create schema gomusic
--

CREATE DATABASE IF NOT EXISTS gomusic;
USE gomusic;

--
-- Definition of table `customers`
--

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `firstname` varchar(50) NOT NULL,
  `lastname` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `pass` varchar(100) NOT NULL,
  `cc_customerid` varchar(50) NOT NULL,
  `loggedin` tinyint(4) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `customers`
--

/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` (`id`,`firstname`,`lastname`,`email`,`pass`,`cc_customerid`,`loggedin`,`created_at`,`updated_at`,`deleted_at`) VALUES 
 (1,'huynh','vu','duchuynhvu@gmail.com','$2a$10$WShAk/kmU5JCOIhT.pAPx.wchdmLe5cu3eFCudzprD1BCCHbFAFwe','',0,'2019-08-28 12:50:14','2019-08-28 12:50:14','0000-00-00 00:00:00'),
 (2,'Linh','Nguyen','linhnguyen@test.com','$2a$10$WShAk/kmU5JCOIhT.pAPx.wchdmLe5cu3eFCudzprD1BCCHbFAFwe','cus_FjwHRVcaRCw8R2',1,'2019-08-29 04:57:22','2019-08-29 04:57:22','2019-08-29 11:57:22'),
 (3,'Duy','Vu','duyvu@test.com','$2a$10$WShAk/kmU5JCOIhT.pAPx.wchdmLe5cu3eFCudzprD1BCCHbFAFwe','',1,'2019-08-29 05:00:47','2019-08-29 05:00:47','2019-08-29 12:00:47'),
 (4,'Duy','Vu','duyvu@test.com','$2a$10$WShAk/kmU5JCOIhT.pAPx.wchdmLe5cu3eFCudzprD1BCCHbFAFwe','',1,'2019-08-29 05:04:15','2019-08-29 05:04:15','2019-08-29 12:04:15'),
 (5,'Test','Vu','testvu@test.com','$2a$10$WShAk/kmU5JCOIhT.pAPx.wchdmLe5cu3eFCudzprD1BCCHbFAFwe','',1,'2019-08-29 05:50:03','2019-08-29 05:50:03','2019-08-29 12:50:03'),
 (6,'Test','Vu','testvu@test.com','$2a$10$4nkC0XUoUTLCUZ6I/0zDGOVRpmJO9msciBjR1FAGT2rSbXqfJiE0O','',1,'2019-09-03 09:47:06','2019-09-03 09:47:06','2019-09-03 09:47:06');
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;


--
-- Definition of table `orders`
--

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `sell_price` int(11) NOT NULL,
  `purchase_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `orders`
--

/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` (`id`,`customer_id`,`product_id`,`sell_price`,`purchase_date`,`created_at`,`updated_at`,`deleted_at`) VALUES 
 (1,1,1,600,'0000-00-00 00:00:00','2019-08-29 10:34:40','2019-08-29 10:34:40','2019-08-29 17:34:40'),
 (2,2,1,600,'0000-00-00 00:00:00','2019-08-29 10:35:11','2019-08-29 10:35:11','2019-08-29 17:35:11'),
 (3,2,1,600,'0000-00-00 00:00:00','2019-08-29 10:43:25','2019-08-29 10:43:25','2019-08-29 17:43:25'),
 (4,2,1,600,'0000-00-00 00:00:00','2019-08-29 10:46:18','2019-08-29 10:46:18','2019-08-29 17:46:18'),
 (5,2,1,600,'0000-00-00 00:00:00','2019-08-29 23:33:35','2019-08-29 23:33:35','2019-08-29 23:33:35'),
 (6,2,1,600,'0000-00-00 00:00:00','2019-09-03 09:48:37','2019-09-03 09:48:37','2019-09-03 09:48:37'),
 (7,2,1,600,'0000-00-00 00:00:00','2019-09-03 09:57:21','2019-09-03 09:57:21','2019-09-03 09:57:21'),
 (8,2,1,600,'0000-00-00 00:00:00','2019-09-03 10:09:06','2019-09-03 10:09:06','2019-09-03 10:09:06'),
 (9,2,1,600,'0000-00-00 00:00:00','2019-09-03 10:09:41','2019-09-03 10:09:41','2019-09-03 10:09:41');
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;


--
-- Definition of table `products`
--

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `image` varchar(100) DEFAULT NULL,
  `imgalt` varchar(50) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `productname` varchar(50) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `promotion` float DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `products`
--

/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` (`id`,`image`,`imgalt`,`description`,`productname`,`price`,`promotion`,`created_at`,`updated_at`,`deleted_at`) VALUES 
 (1,'img/strings.png','string','A very authentic and beautiful instrument!!','Strings',100,95,'2019-08-29 11:33:44','0000-00-00 00:00:00','0000-00-00 00:00:00'),
 (2,'img/redguitar.jpeg','redg','A really cool red guitar that can produce super cool music!!','Red Guitar',299,0,'2019-08-28 13:47:05','0000-00-00 00:00:00','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;




/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
