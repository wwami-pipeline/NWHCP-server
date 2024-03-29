CREATE DATABASE mydatabase;
USE mydatabase;

CREATE TABLE if not exists Users
(
    UserID
    INT
    NOT
    NULL
    AUTO_INCREMENT
    PRIMARY
    KEY,
    Email
    VARCHAR
(
    128
) NOT NULL UNIQUE,
    PassHash BINARY
(
    128
) NOT NULL,
    FirstName VARCHAR
(
    128
) NOT NULL,
    LastName VARCHAR
(
    128
) NOT NULL,
    JoinDate VARCHAR
(
    15
) NOT NULL
    );

CREATE TABLE if not exists SignIns
(
    SignInID
    INT
    NOT
    NULL
    AUTO_INCREMENT
    PRIMARY
    KEY,
    UserID
    INT
    NOT
    NULL,
    SignInDate
    DATETIME
    NOT
    NULL,
    IPAddress
    VARCHAR
(
    100
) NOT NULL UNIQUE
    );

CREATE TABLE if not exists user_org
(
    UserOrgID
    INT
    NOT
    NULL
    AUTO_INCREMENT
    PRIMARY
    KEY,
    UserID
    INT
    NOT
    NULL,
    OrgID
    INT
    NOT
    NULL,
    FOREIGN
    KEY
(
    UserID
) REFERENCES Users
(
    UserID
)
    );