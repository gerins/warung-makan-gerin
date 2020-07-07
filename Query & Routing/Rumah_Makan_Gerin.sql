-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema warung_makan_gerin
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema warung_makan_gerin
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `warung_makan_gerin` DEFAULT CHARACTER SET utf8 ;
USE `warung_makan_gerin` ;

-- -----------------------------------------------------
-- Table `warung_makan_gerin`.`category_menu`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `warung_makan_gerin`.`category_menu` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nama` VARCHAR(45) NOT NULL,
  `status` VARCHAR(5) NOT NULL DEFAULT 'A',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 5
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `warung_makan_gerin`.`daftar_transaksi`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `warung_makan_gerin`.`daftar_transaksi` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `peralatan_makan` INT NOT NULL DEFAULT '0',
  `total` INT NOT NULL DEFAULT '0',
  `status` VARCHAR(5) NOT NULL DEFAULT 'A',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 22
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `warung_makan_gerin`.`menu`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `warung_makan_gerin`.`menu` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nama` VARCHAR(45) NOT NULL,
  `harga` VARCHAR(45) NOT NULL,
  `stock` INT NOT NULL DEFAULT '0',
  `category_menu_id` INT NOT NULL DEFAULT '1',
  `status` VARCHAR(5) NOT NULL DEFAULT 'A',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_menu_category_menu_idx` (`category_menu_id` ASC) VISIBLE,
  CONSTRAINT `fk_menu_category_menu`
    FOREIGN KEY (`category_menu_id`)
    REFERENCES `warung_makan_gerin`.`category_menu` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 6
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `warung_makan_gerin`.`detail_transaksi`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `warung_makan_gerin`.`detail_transaksi` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `daftar_transaksi_id` INT NOT NULL,
  `menu_id` INT NOT NULL,
  `kuantiti` INT NOT NULL,
  `extra_pedas` INT NOT NULL,
  `total` INT NOT NULL,
  `status` VARCHAR(5) NOT NULL DEFAULT 'A',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_detail_transaksi_menu1_idx` (`menu_id` ASC) VISIBLE,
  INDEX `fk_detail_transaksi_daftar_transaksi1_idx` (`daftar_transaksi_id` ASC) VISIBLE,
  CONSTRAINT `fk_detail_transaksi_daftar_transaksi1`
    FOREIGN KEY (`daftar_transaksi_id`)
    REFERENCES `warung_makan_gerin`.`daftar_transaksi` (`id`),
  CONSTRAINT `fk_detail_transaksi_menu1`
    FOREIGN KEY (`menu_id`)
    REFERENCES `warung_makan_gerin`.`menu` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 37
DEFAULT CHARACTER SET = utf8;

USE `warung_makan_gerin` ;

-- -----------------------------------------------------
-- Placeholder table for view `warung_makan_gerin`.`detail_transaksi_idx`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `warung_makan_gerin`.`detail_transaksi_idx` (`trans_id` INT, `nama` INT, `kategori` INT, `harga` INT, `extra_pedas` INT, `kuantiti` INT, `total` INT);

-- -----------------------------------------------------
-- Placeholder table for view `warung_makan_gerin`.`menu_idx`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `warung_makan_gerin`.`menu_idx` (`id` INT, `nama` INT, `harga` INT, `stock` INT, `jenis` INT, `status` INT, `created_at` INT, `updated_at` INT);

-- -----------------------------------------------------
-- View `warung_makan_gerin`.`detail_transaksi_idx`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `warung_makan_gerin`.`detail_transaksi_idx`;
USE `warung_makan_gerin`;
CREATE  OR REPLACE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `warung_makan_gerin`.`detail_transaksi_idx` AS select `dt`.`daftar_transaksi_id` AS `trans_id`,`m`.`nama` AS `nama`,`cm`.`nama` AS `kategori`,`m`.`harga` AS `harga`,`dt`.`extra_pedas` AS `extra_pedas`,`dt`.`kuantiti` AS `kuantiti`,`dt`.`total` AS `total` from ((`warung_makan_gerin`.`detail_transaksi` `dt` join `warung_makan_gerin`.`menu` `m` on((`dt`.`menu_id` = `m`.`id`))) join `warung_makan_gerin`.`category_menu` `cm` on((`cm`.`id` = `m`.`category_menu_id`)));

-- -----------------------------------------------------
-- View `warung_makan_gerin`.`menu_idx`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `warung_makan_gerin`.`menu_idx`;
USE `warung_makan_gerin`;
CREATE  OR REPLACE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `warung_makan_gerin`.`menu_idx` AS select `m`.`id` AS `id`,`m`.`nama` AS `nama`,`m`.`harga` AS `harga`,`m`.`stock` AS `stock`,`cm`.`nama` AS `jenis`,`m`.`status` AS `status`,`m`.`created_at` AS `created_at`,`m`.`updated_at` AS `updated_at` from (`warung_makan_gerin`.`menu` `m` join `warung_makan_gerin`.`category_menu` `cm` on((`cm`.`id` = `m`.`category_menu_id`)));

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
