SELECT * FROM dbarquitectura.detalleEvento;

CREATE TABLE `dbarquitectura`.`detalleEvento` (
  `id_detalle` INT NOT NULL,
  `id_evento` varchar(45) NOT NULL,
  `descripcion` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`id_detalle`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

INSERT INTO `dbarquitectura`.`detalleEvento` (`id_detalle`, `id_evento`, `descripcion`) VALUES ('1', '4', 'Asalto a mano armada a GANA');
INSERT INTO `dbarquitectura`.`detalleEvento` (`id_detalle`, `id_evento`, `descripcion`) VALUES ('2', '3', 'Violencia sexual por parte de dos personas identifiacadas como homosexuales');
INSERT INTO `dbarquitectura`.`detalleEvento` (`id_detalle`, `id_evento`, `descripcion`) VALUES ('3', '2', 'Accidente de transito por exceso de velocidad');
INSERT INTO `dbarquitectura`.`detalleEvento` (`id_detalle`, `id_evento`, `descripcion`) VALUES ('4', '1', 'Violencia y robo de un celular con arma blanca');