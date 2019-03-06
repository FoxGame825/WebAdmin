CREATE TABLE IF NOT EXISTS `admin_userlog`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `userid` INT(11) NOT NULL,
   `action` VARCHAR(1024),
   `log` VARCHAR (2048),
   `created` DATE,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;