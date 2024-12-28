CREATE DATABASE IF NOT EXISTS `payment`;
USE `payment`;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE accounts (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  user_id VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  currency VARCHAR(10) NOT NULL,
  balance DECIMAL(10, 2) DEFAULT 0.0,
  status VARCHAR(20) DEFAULT 'active',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `account_address_information`;
CREATE TABLE account_address_information (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    account_id INTEGER NOT NULL,
    type VARCHAR(20) DEFAULT 'primary',
    status VARCHAR(20) DEFAULT 'inactive',
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
CREATE FULLTEXT INDEX idx_address ON account_address_information(address);

DROP TABLE IF EXISTS `external_banks`;
CREATE TABLE external_banks (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  api_url VARCHAR(255) NOT NULL,
  auth_credentials VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE transactions (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  from_account_id INTEGER NOT NULL,
  to_account_id INTEGER,
  address VARCHAR(50),
  amount DECIMAL(10,2) NOT NULL,
  currency VARCHAR(10) NOT NULL,
  transaction_type VARCHAR(20) NOT NULL,
  status VARCHAR(20) DEFAULT 'pending',
  description VARCHAR(255),
  external_bank_id INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (from_account_id) REFERENCES accounts(id),
  FOREIGN KEY (external_bank_id) REFERENCES external_banks(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;