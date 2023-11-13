SELECT * FROM bdarquitectura.usuario;

CREATE TABLE `usuario` (
  `id_usuario` int NOT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `apellido` varchar(45) DEFAULT NULL,
  `telefono` varchar(45) DEFAULT NULL,
  `correo` varchar(45) DEFAULT NULL,
  `fecha_nacimiento` varchar(45) DEFAULT NULL,
  `deleted_at` varchar(45) DEFAULT NULL,
  `created_at` varchar(45) DEFAULT NULL,
  `updated_at` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_usuario`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `bdarquitectura`.`usuario` (`id_usuario`, `nombre`, `apellido`, `telefono`, `correo`, `fecha_nacimiento`) VALUES ('1', 'David', 'Florez', '353466', 'florez@gmail.com', '21/12/2001');
INSERT INTO `bdarquitectura`.`usuario` (`id_usuario`, `nombre`, `apellido`, `telefono`, `correo`, `fecha_nacimiento`) VALUES ('2', 'Estefan', 'Torres', '3452345', 'torres@gmail.com', '10/08/2001');
INSERT INTO `bdarquitectura`.`usuario` (`id_usuario`, `nombre`, `apellido`, `telefono`, `correo`, `fecha_nacimiento`) VALUES ('3', 'Sebastian', 'Chavarria', '6876965', 'chavarria@gmail.com', '17/11/1999');
INSERT INTO `bdarquitectura`.`usuario` (`id_usuario`, `nombre`, `apellido`, `telefono`, `correo`, `fecha_nacimiento`) VALUES ('4', 'Sara', 'Cano', '2345', 'cano@gmail.com', '11/11/1994');

