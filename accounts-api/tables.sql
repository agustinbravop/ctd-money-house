DROP DATABASE IF EXISTS ACCOUNTSDB;
CREATE DATABASE IF NOT EXISTS ACCOUNTSDB;
USE ACCOUNTSDB;

--
-- Table structure for table `accounts`
--
-- accounts definition OK
CREATE TABLE IF NOT EXISTS accounts
(
    id      BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    cvu     VARCHAR(22)     NOT NULL UNIQUE,
    alias   VARCHAR(60)     NOT NULL UNIQUE,
    amount  DECIMAL         NOT NULL,
    user_id VARCHAR(36)     NOT NULL,
    CONSTRAINT id_PK PRIMARY KEY (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
    COMMENT ='accounts information';

--
-- Table structure for table `transaction`
--
-- transaction definition OK
CREATE TABLE IF NOT EXISTS transactions
(
    id               BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    amount           DECIMAL         NOT NULL,
    transaction_date DATETIME        NOT NULL,
    description      VARCHAR(60)     NOT NULL,
    origin_cvu       VARCHAR(22)     NOT NULL,
    destination_cvu  VARCHAR(22)     NOT NULL,
    account_id       BIGINT UNSIGNED NOT NULL,
    transaction_type VARCHAR(10)     NOT NULL,
    CONSTRAINT id_PK PRIMARY KEY (id),
    CONSTRAINT FK_TRANSACTIONS_ACCOUNTS FOREIGN KEY (account_id)
      REFERENCES accounts (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
    COMMENT ='transactions information';


--
-- Table structure for table `cards`
--
-- cards definition OK
CREATE TABLE IF NOT EXISTS cards
(
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    card_number     VARCHAR(40)     NOT NULL,
    expiration_date VARCHAR(5)      NOT NULL,
    owner           VARCHAR(40)     NOT NULL,
    security_code   VARCHAR(5)      NOT NULL,
    brand           VARCHAR(40)     NOT NULL,
    account_id      BIGINT UNSIGNED NOT NULL,
    CONSTRAINT id_PK PRIMARY KEY (id),
    CONSTRAINT FK_CARDS_ACCOUNTS FOREIGN KEY (account_id)
        REFERENCES accounts (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
    COMMENT ='cards information';