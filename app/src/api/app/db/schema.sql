-- name: load-schema
CREATE TABLE IF NOT EXISTS items(
	id MEDIUMINT(11) NOT NULL AUTO_INCREMENT,
	name CHAR(30),
	description TEXT,
	PRIMARY KEY (id)
);