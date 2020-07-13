DROP DATABASE IF EXISTS `mcvt`;
CREATE DATABASE `mcvt`;

USE `mcvt`;

DROP TABLE IF EXISTS `HM_CTRLDECSN`;
CREATE TABLE `HM_CTRLDECSN` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENDER_ID` smallint(6) NOT NULL,
  `RECEIVER_ID` smallint(6) NOT NULL,
  `COMMAND` enum('COMMAND0', 'COMMAND1', 'COMMAND2')
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HM_CTRLDECSN` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `HM_RT_THERMFLUX`;
CREATE TABLE `HM_RT_THERMFLUX` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENSOR_ID` smallint(6) NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `THERMFLUX` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HM_RT_THERMFLUX` WRITE;
UNLOCK TABLES;
/*
DROP TABLE IF EXISTS `HM_THERMFLUX`;
CREATE TABLE `HM_THERMFLUX` (
  `SENSOR_ID` smallint(6) NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `THERMFLUX` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HM_THERMFLUX` WRITE;
UNLOCK TABLES;
*/
DROP TABLE IF EXISTS `SS_RT_SENSORS`;
CREATE TABLE `SS_RT_SENSORS` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENSOR_ID` smallint(6) NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `SENSOR_TYPE` enum("STRAIN", "ACCL", "TEMP", "DISP", "VISUAL", "DAMAGE", "THERMFLUX", "PRESSURE") NOT NULL,
  `TEMPERATURE` double(5, 2),
  `STRAIN` double(5, 2),
  `ACCL_X` double(5, 2),
  `ACCL_Y` double(5,2),
  `ACCL_Z` double(5, 2),
  `DISP_X` double(5, 2),
  `DISP_Y` double(5, 2),
  `DISP_Z` double(5, 2),
  `VISUAL` double(5, 2),
  `DAMAGE` double(5, 2),
  `PRESSURE` double(5, 2)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_RT_SENSORS` WRITE;
UNLOCK TABLES;
/*
DROP TABLE IF EXISTS `SS_TEMP`;
CREATE TABLE `SS_TEMP` (
  `SENSOR_ID` smallint(6) NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `TEMPERATURE` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_TEMP` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_STRAIN`;
CREATE TABLE `SS_STRAIN` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `STRAIN` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_STRAIN` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_ACCL`;
CREATE TABLE `SS_ACCL` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `ACCL_X` double(5, 2) NOT NULL,
  `ACCL_Y` double(5, 2) NOT NULL,
  `ACCL_Z` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_ACCL` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_DISP`;
CREATE TABLE `SS_DISP` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `DISP_X` double(5, 2) NOT NULL,
  `DISP_Y` double(5, 2) NOT NULL,
  `DISP_Z` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_DISP` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_VISUAL`;
CREATE TABLE `SS_VISUAL` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `VISUAL` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_VISUAL` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_DAMAGE`;
CREATE TABLE `SS_DAMAGE` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `DAMAGE` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_DAMAGE` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_THERMFLUX`;
CREATE TABLE `SS_THERMFLUX` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `THERMFLUX` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_THERMFLUX` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `SS_PRESSURE`;
CREATE TABLE `SS_PRESSURE` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `PRESSURE` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `SS_PRESSURE` WRITE;
UNLOCK TABLES;
*/
DROP TABLE IF EXISTS `PWR_GENERATION`;
CREATE TABLE `PWR_GENERATION` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SOLAR_PANEL_ID` smallint(5) NOT NULL,
  `SOLAR_PANEL_STATUS` enum("SLEEP, HALF, FULL, DAMAGED") NOT NULL,
  `SOLAR_PANEL_EFFICIENCY` int(20) NOT NULL,
  `PANEL_X` int(20) NOT NULL,
  `PANEL_Y` int(20) NOT NULL,
  `PANEL_Z` int(20) NOT NULL,
  `POWER_GEN` double(5, 2) NOT NULL,
  `HEAT_GEN` double(5, 2) NOT NULL,
  `TOTAL_POWER_GEN` double(5, 2) NOT NULL,
  `TOTAL_HEAT_GEN` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `PWR_GENERATION` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `PWR_STORAGE`;
CREATE TABLE `PWR_STORAGE` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `BATTERY_ID` smallint(5) NOT NULL,
  `BATTERY_STATUS` enum("HEALTHY, DAMAGED") NOT NULL,
  `BATTERY_EFFICIENCY` int(20) NOT NULL,
  `BATTERY_PERCENT` double(5, 2) NOT NULL,
  `VOLTAGE` double(5, 2) NOT NULL,
  `TOTAL_VOLTAGE` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `PWR_STORAGE` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `PWR_CONSUMPTION`;
CREATE TABLE `PWR_CONSUMPTION` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `CONSUMER_ID` smallint(5) NOT NULL,
  `POWER_CONS` double(5, 2) NOT NULL,
  `PERCENT_TOTAL_CONSUME` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `PWR_CONSUMPTION` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `PWR_HEALTHSTATE`;
CREATE TABLE `PWR_HEALTHSTATE` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `OPERATIONUNIT_ID` smallint(5) NOT NULL,
  `OPERATIONUNIT_HEALTHSTATE` enum("HEALTHY, DAMAGED, LOW_POWER") NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `PWR_HEALTHSTATE` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `ECLSS_RT_OPERATION`;
CREATE TABLE `ECLSS_RT_OPERATION` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `UNIT_ID` smallint(5) NOT NULL,
  `UNIT_TYPE` enum("WRS", "OGS", "FMS"),
  `HEALTH_STATUS` enum("HEALTHY, DAMAGED") NOT NULL,
  `OPERATION_STATUS` enum("ACTIVE", "STANDBY", "SLEEP") NOT NULL,
  `HUMIDITY` double(5, 2), -- WRS
  `WATER_LEVEL` double(5, 2), -- WRS
  `OXYGEN_LEVEL` double(5, 2), -- OGS
  `CO2_LEVEL` double(5, 2), -- OGS
  `REFRIGERATOR_TEMP` double(5, 2), -- FMS
  `MOISTURE` double(5, 2), -- FMS
  `CIRCULATION` double(5, 2) -- FMS
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `ECLSS_RT_OPERATION` WRITE;
UNLOCK TABLES;
/*
DROP TABLE IF EXISTS `ECLSS_WRS`;
CREATE TABLE `ECLSS_WRS` (
  `UNIT_ID` smallint(5) unsigned NOT NULL,
  `HEALTH_STATUS` enum("HEALTHY, DAMAGED") NOT NULL,
  `OPERATION_STATUS` enum("ACTIVE", "STANDBY", "SLEEP") NOT NULL,
  `UNIT_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `HUMIDITY` double(5, 2) NOT NULL,
  `WATER_LEVEL` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `ECLSS_WRS` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `ECLSS_OGS`;
CREATE TABLE `ECLSS_OGS` (
  `UNIT_ID` smallint(5) unsigned NOT NULL,
  `HEALTH_STATUS` enum("HEALTHY, DAMAGED") NOT NULL,
  `OPERATION_STATUS` enum("ACTIVE", "STANDBY", "SLEEP") NOT NULL,
  `UNIT_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `OXYGEN_LEVEL` double(5, 2) NOT NULL,
  `CO2_LEVEL` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `ECLSS_OGS` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `ECLSS_FMS`;
CREATE TABLE `ECLSS_FMS` (
  `UNIT_ID` smallint(5) unsigned NOT NULL,
  `HEALTH_STATUS` enum("HEALTHY, DAMAGED") NOT NULL,
  `OPERATION_STATUS` enum("ACTIVE", "STANDBY", "SLEEP") NOT NULL,
  `UNIT_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `REFRIGERATOR_TEMP` double(5, 2) NOT NULL,
  `MOISTURE` double(5, 2) NOT NULL,
  `CIRCULATION` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `ECLSS_FMS` WRITE;
UNLOCK TABLES;
*/

DROP TABLE IF EXISTS `HUMAN_RT_UPDATES`;
CREATE TABLE `HUMAN_RT_UPDATES` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `HUMAN_ID` smallint(6) NOT NULL,
  `HUMAN_STATUS` enum("NORMAL", "PANIC", "UNCONSCIOUS", "DECEASED") NOT NULL,
  `HR` double(5, 2) NOT NULL,
  `BP` double(5, 2) NOT NULL,
  `O2_SAT` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HUMAN_RT_UPDATES` WRITE;
UNLOCK TABLES;
/*
DROP TABLE IF EXISTS `HUMAN_CURRENT`;
CREATE TABLE `HUMAN_CURRENT` (
  `HUMAN_ID` smallint(6) NOT NULL,
  `HUMAN_STATUS` enum("NORMAL", "PANIC", "UNCONSCIOUS", "DECEASED") NOT NULL,
  `LAST_UPDATE` bigint(20) NOT NULL,
  `HR` double(5, 2) NOT NULL,
  `BP` double(5, 2) NOT NULL,
  `O2_SAT` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HUMAN_CURRENT` WRITE;
UNLOCK TABLES;
*/
DROP TABLE IF EXISTS `HUMAN_SETPTCTRL`;
CREATE TABLE `HUMAN_SETPTCTRL` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENDER_ID` smallint(6) NOT NULL,
  `RECEIVER_ID` smallint(6) NOT NULL,
  `SET_POINT` int(20)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HUMAN_SETPTCTRL` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `ROBOT_RT_UPDATES`;
CREATE TABLE `ROBOT_RT_UPDATES` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `ROBOT_ID` smallint(6) NOT NULL,
  `ROBOT_STATUS` enum("ACTIVE", "STANDBY", "SLEEP") NOT NULL,
  `ROBOT_HEALTH` enum("FULL FUNC", "SOME DAMAGE", "CRITICAL") NOT NULL,
  `BATTERY` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `ROBOT_RT_UPDATES` WRITE;
UNLOCK TABLES;
/*
DROP TABLE IF EXISTS `ROBOT_CURRENT`;
CREATE TABLE `ROBOT_CURRENT` (
  `ROBOT_ID` smallint(6) NOT NULL,
  `ROBOT_STATUS` enum("ACTIVE", "STANDBY", "SLEEP") NOT NULL,
  `ROBOT_HEALTH` enum("FULL FUNC", "SOME DAMAGE", "CRITICAL") NOT NULL,
  `LAST_UPDATE` bigint(20) NOT NULL,
  `HR` double(5, 2) NOT NULL,
  `BP` double(5, 2) NOT NULL,
  `O2_SAT` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `HUMAN_CURRENT` WRITE;
UNLOCK TABLES;
*/
DROP TABLE IF EXISTS `ROBOT_INTERVENTIONCTRL`;
CREATE TABLE `ROBOT_INTERVENTIONCTRL` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENDER_ID` smallint(6) NOT NULL,
  `RECEIVER_ID` smallint(6) NOT NULL,
  `INTERVENTION` enum("IN0", "IN1", "IN2")
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `ROBOT_INTERVENTIONCTRL` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `IE_RT_SENSORS`;
CREATE TABLE `IE_RT_SENSORS` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENSOR_ID` smallint(6) NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `SENSOR_TYPE` enum("TEMP", "PRESSURE", "HUMIDITY") NOT NULL,
  `TEMPERATURE` double(5, 2),
  `PRESSURE` double(5, 2),
  `HUMIDITY` double(5, 2)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `IE_RT_SENSORS` WRITE;
UNLOCK TABLES;
/*
DROP TABLE IF EXISTS `IE_TEMP`;
CREATE TABLE `IE_TEMP` (
  `SENSOR_ID` smallint(6) NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `TEMPERATURE` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `IE_TEMP` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `IE_PRESSURE`;
CREATE TABLE `IE_PRESSURE` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `PRESSURE` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `IE_PRESSURE` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `IE_HUMIDITY`;
CREATE TABLE `IE_HUMIDITY` (
  `SENSOR_ID` smallint(5) unsigned NOT NULL,
  `SENSOR_STATUS` enum("REPLACE", "POOR", "GOOD", "EXCELLENT") NOT NULL,
  `SENSOR_POWER_CONSUME` double(10, 2),
  `LAST_UPDATE` bigint(20),
  `SENSOR_X` int(20) NOT NULL,
  `SENSOR_Y` int(20) NOT NULL,
  `SENSOR_Z` int(20) NOT NULL,
  `HUMIDITY` double(5, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `IE_HUMIDITY` WRITE;
UNLOCK TABLES;
*/
DROP TABLE IF EXISTS `EE_DISTURBANCE_DETECTION`;
CREATE TABLE `EE_DISTURBANCE_DETECTION` (
  `DATETIME` datetime NOT NULL,
  `TIMESTAMP` bigint(20) NOT NULL,
  `SENDER_ID` smallint(6) NOT NULL,
  `RECEIVER_ID` smallint(6) NOT NULL,
  `SENDER_X` int(20) NOT NULL,
  `SENDER_Y` int(20) NOT NULL,
  `SENDER_Z` int(20) NOT NULL,
  `DISTURBANCE` enum("MOONQUAKE", "TEMP_FLUC", "METEOROID_IMPACT", "DUST_POTENTIAL", "DUST_STORM", "SOLAR_RADIATION")
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `EE_DISTURBANCE_DETECTION` WRITE;
UNLOCK TABLES;

DROP TABLE IF EXISTS `INVENTORY`;
CREATE TABLE `INVENTORY` (
  `ITEM_ID` smallint(6) NOT NULL,
  `QUANTITY` smallint(6) NOT NULL,
  `RESTOCK_ORDER` enum("SENT", "NOT SENT"),
  `LAST_UPDATE` bigint(20) NOT NULL,
  `RECEIVER_ID` smallint(6)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
LOCK TABLES `EE_DISTURBANCE_DETECTION` WRITE;
UNLOCK TABLES;