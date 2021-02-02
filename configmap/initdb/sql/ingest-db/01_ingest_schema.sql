SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 ;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 ;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL' ;


CREATE DATABASE qservIngest;
USE qservIngest;

-- --------------------------------------------------------------
-- Table `task`
-- --------------------------------------------------------------
--
-- The list of chunks to load inside a Qserv database
-- Used as a queue by loader jobs
CREATE TABLE `chunkfile` (

  `id`                    INTEGER UNSIGNED    NOT NULL AUTO_INCREMENT,
  `chunk_id`              INTEGER UNSIGNED    NOT NULL ,                  -- the id of the chunk to load
  `chunk_file_path`       VARCHAR(255)        NOT NULL ,                  -- the path of the chunk file to load
  `database`              VARCHAR(255)        NOT NULL ,                  -- the name of the target database
  `is_overlap`            BOOLEAN             NOT NULL ,                  -- is this file an overlap
  `table`                 VARCHAR(255)        NOT NULL ,                  -- the name of the target table
  `ingest_time`           TIMESTAMP           NULL ,                      -- the date when this file is ingested in the current transaction

  PRIMARY KEY (`id`),
  UNIQUE KEY (`chunk_id`, `chunk_file_path`, `database`, `is_overlap`, `table`)
)
ENGINE = InnoDB;

create table `chunkfile_transaction` (
    `chunkfile_id`   INTEGER UNSIGNED    NOT NULL ,
    `transaction_id` INTEGER UNSIGNED    NOT NULL ,

    UNIQUE KEY (`chunkfile_id`, `transaction_id`),

    CONSTRAINT `transaction_fk1`
      FOREIGN KEY (`transaction_id`)
      REFERENCES `transaction` (`id`),

    CONSTRAINT `chunkfile_fk1`
      FOREIGN KEY (`chunkfile_id`)
      REFERENCES `chunkfile` (`id`),
)
ENGINE = InnoDB;

CREATE TABLE `transaction` (

  `id`         INTEGER      NOT NULL ,                  -- the id of a replication service super-transaction
  `pod`        VARCHAR(255) NOT NULL ,                  -- the id of the pod which run the transaction
  -- `state`      VARCHAR(255) NOT NULL ,                  -- the latest returned state from the replication REST service
  -- `begin_time` TIMESTAMP    NOT NULL ,                  -- the date when the replication REST service returns when starting a super-transaction
  -- `end_time`   TIMESTAMP    DEFAULT NULL ,              -- the date when the replication REST service returns when stopping a super-transaction
)
ENGINE = InnoDB;
