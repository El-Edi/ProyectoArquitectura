SELECT * FROM dbarquitectura.ubicacion;

CREATE TABLE `eventos` (
  `id_evento` varchar(45) NOT NULL,
  `tipo` varchar(45) NOT NULL,
  `descripcion` VARCHAR(100) NOT NULL,
  `id_ubicacion` INT NOT NULL,
  PRIMARY KEY (`id_evento`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--faltan los insert