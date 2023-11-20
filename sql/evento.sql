SELECT * FROM dbarquitectura.eventos;

CREATE TABLE `eventos` (
  `id_evento` varchar(45) NOT NULL,
  `tipo` varchar(45) NOT NULL,
  `descripcion` VARCHAR(100) NOT NULL,
  `id_ubicacion` INT NOT NULL,
  PRIMARY KEY (`id_evento`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `dbarquitectura`.`evento` (`id_evento`, `tipo`, `descripcion`,`id_ubicacion`) VALUES ('1', '4', 'Asalto a mano armada a GANA','1');
INSERT INTO `dbarquitectura`.`evento` (`id_evento`, `tipo`, `descripcion`,`id_ubicacion`) VALUES ('2', '3', 'Violencia sexual por parte de dos personas identifiacadas como homosexuales','2');
INSERT INTO `dbarquitectura`.`evento` (`id_evento`, `tipo`, `descripcion`,`id_ubicacion`) VALUES ('3', '3', 'Accidente de transito por exceso de velocidad','3');
INSERT INTO `dbarquitectura`.`evento` (`id_evento`, `tipo`, `descripcion`,`id_ubicacion`) VALUES ('4', '1', 'Violencia y robo de un celular con arma blanca','4');