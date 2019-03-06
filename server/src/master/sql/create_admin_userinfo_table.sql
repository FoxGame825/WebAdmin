CREATE TABLE IF NOT EXISTS `admin_userinfo`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `username` VARCHAR(100) NOT NULL UNIQUE ,
   `password` VARCHAR(40) NOT NULL,
		`roles` VARCHAR(40) NOT NULL,
		`permission` INT,
   `created` DATE,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;