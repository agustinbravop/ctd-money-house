DROP SCHEMA IF EXISTS USERSDB;
CREATE SCHEMA IF NOT EXISTS USERSDB;
USE USERSDB;

--
-- Table structure for table `users`
--
-- users definition OK
CREATE TABLE IF NOT EXISTS users
(
    id         VARCHAR(36)  NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name  VARCHAR(150) NOT NULL,
    dni        VARCHAR(10)  NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    telephone  VARCHAR(20)  NOT NULL,
    CONSTRAINT id_PK PRIMARY KEY (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
    COMMENT ='user information';