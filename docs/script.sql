DROP TABLE IF EXISTS `database`.operations

  
CREATE TABLE `database`.operations (
	id MEDIUMINT AUTO_INCREMENT NOT NULL,
	user_name varchar(255) NOT NULL,
	type varchar(255) NOT NULL,
	coin varchar(255) NOT NULL,
	value double NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
)
