CREATE TABLE IF NOT EXISTS `admin_usertoken`(
   `userid` INT(11) NOT NULL,
   `token` VARCHAR(1024),
   `ip` VARCHAR(15),
   `created` DATE,
   PRIMARY KEY ( `userid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;