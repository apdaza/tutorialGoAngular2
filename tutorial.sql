CREATE DATABASE  IF NOT EXISTS `contactos_db`;
USE `contactos_db`;
CREATE TABLE `contacto` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) DEFAULT NULL,
  `apellido` varchar(45) DEFAULT NULL,
  `celular` varchar(45) DEFAULT NULL,
  `mail` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

