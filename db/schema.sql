CREATE DATABASE accounts_db
USE accounts_db

CREATE TABLE [user] (
    [user_id] INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(128) NOT NULL UNIQUE,
    passhash BINARY(128) NOT NULL,
    firstname VARCHAR(128) NOT NULL,
    lastname VARCHAR(128) NOT NULL,
    dob DATE NOT NULL,
	join_date DATE NOT NULL,
    PRIMARY KEY ([user_id]),
	FOREIGN KEY (inst_id) REFERENCES institution(inst_id)
);

CREATE TABLE organization (
	org_id INT NOT NULL AUTO_INCREMENT,
	org_title VARCHAR(128) NOT NULL,
	PRIMARY KEY (org_id)
);

CREATE TABLE user_org (
	id INT NOT NULL AUTO_INCREMENT,
	[user_id] INT NOT NULL,
	org_id INT NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY ([user_id]) REFERENCES [user]([user_id]),
	FOREIGN KEY (org_id) REFERENCES organization(org_id)
);