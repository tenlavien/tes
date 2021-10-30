CREATE SCHEMA IF NOT EXISTS `test_hub` DEFAULT CHARACTER SET utf8;
USE `test_hub`;

CREATE TABLE IF NOT EXISTS `test_hub`.`steps`
(
    `id`                   BIGINT(64)   NOT NULL AUTO_INCREMENT,
    `step_code`            VARCHAR(400) NOT NULL,
    `run_id`               VARCHAR(400) NOT NULL,
    `case_id`              BIGINT(64)   NOT NULL,
    `description`          LONGTEXT     NOT NULL,
    `status`               VARCHAR(100) NOT NULL,
    `step_in`              LONGTEXT     NOT NULL,
    `step_out`             LONGTEXT     NOT NULL,
    `elapsed_milliseconds` BIGINT(64)   NOT NULL,
    `created_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `test_hub`.`cases`
(
    `id`                   BIGINT(64)   NOT NULL AUTO_INCREMENT,
    `case_code`            VARCHAR(400) NOT NULL,
    `run_id`               VARCHAR(400) NOT NULL,
    `description`          LONGTEXT     NOT NULL,
    `status`               VARCHAR(100) NOT NULL,
    `case_in`              LONGTEXT     NOT NULL,
    `case_out`             LONGTEXT     NOT NULL,
    `elapsed_milliseconds` BIGINT(64)   NOT NULL,
    `created_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
