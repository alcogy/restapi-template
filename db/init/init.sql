DROP TABLE IF EXISTS `my_table`;
CREATE TABLE `my_table` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `latest_update` datetime NOT NULL,
  PRIMARY KEY(`id`)
);

INSERT INTO `my_table` values (
  0,
  'Test data 1',
  'Test value 1',
  now()
);

INSERT INTO `my_table` values (
  0,
  'Test data 2',
  'Test value 2',
  now()
);

INSERT INTO `my_table` values (
  0,
  'Test data 3',
  'Test value 3',
  now()
);
