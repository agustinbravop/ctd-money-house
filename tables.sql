DROP SCHEMA IF EXISTS DMH;
CREATE SCHEMA IF NOT EXISTS DMH;
USE DMH;

--
-- Table structure for table `users`
--
-- users definition OK
CREATE TABLE IF NOT EXISTS Users
(
    id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(150)    NOT NULL,
    last_name  VARCHAR(150)    NOT NULL,
    dni        VARCHAR(10)     NOT NULL,
    email      VARCHAR(150)    NOT NULL,
    telephone  VARCHAR(20)     NOT NULL,
    cvu        VARCHAR(22)     NOT NULL,
    alias      VARCHAR(60)     NOT NULL,
    CONSTRAINT id_PK PRIMARY KEY (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
    COMMENT ='user information';