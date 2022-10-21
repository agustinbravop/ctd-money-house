-- users definition OK
CREATE TABLE Users
(
    id          BIGINT UNSIGNED auto_increment NOT NULL,
    keycloak_id varchar(36)  NOT NULL,
    name        varchar(150) NOT NULL,
    last_name   varchar(150) NOT NULL,
    dni         varchar(10)  NOT NULL,
    email       varchar(150) NOT NULL,
    telephone   varchar(20)  NOT NULL,
    cvu         int          NOT NULL,
    alias       varchar(60)  NOT NULL,
    CONSTRAINT id_PK PRIMARY KEY (id)
) ENGINE=INNODB
 DEFAULT CHARSET=utf8mb4
 COLLATE=utf8mb4_0900_ai_ci
 COMMENT='user information';