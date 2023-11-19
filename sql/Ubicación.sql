SELECT * FROM bdarquitectura.ubicacion;

CREATE TABLE `bdarquitectura`.`ubicacion` (
  `id_ubicacion` INT NOT NULL,
  `longitud` VARCHAR(45) NULL,
  `latitud` VARCHAR(45) NULL,
  `precision` INT NULL,
  `deleted_at` VARCHAR(45) NULL,
  `created_at` VARCHAR(45) NULL,
  `updated_at` VARCHAR(45) NULL,
  `id_usuario` INT NULL,
  PRIMARY KEY (`id_ubicacion`),
  INDEX (id_usuario),
  FOREIGN KEY (id_usuario) REFERENCES usuario(id_usuario)
  )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

INSERT INTO `bdarquitectura`.`ubicacion` (`id_ubicacion`, `longitud`, `latitud`, `precision`,`id_usuario`) VALUES ('1', ' -75.562114', '6.336277', '5','1');
INSERT INTO `bdarquitectura`.`ubicacion` (`id_ubicacion`, `longitud`, `latitud`, `precision`,`id_usuario`) VALUES ('2', '-75.559968', '6.311836', '3','2');
INSERT INTO `bdarquitectura`.`ubicacion` (`id_ubicacion`, `longitud`, `latitud`, `precision`,`id_usuario`) VALUES ('3', '-75.570954', '6.277369', '2','3');
INSERT INTO `bdarquitectura`.`ubicacion` (`id_ubicacion`, `longitud`, `latitud`, `precision`,`id_usuario`) VALUES ('4', '-75.643432', '6.175946', '1','4');
