DROP TABLE IF EXISTS blogs;

CREATE TABLE blogs (
  id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(10000) NOT NULL,
  create_date DATE NOT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS chat_history;

CREATE TABLE chat_history (
  id VARCHAR(255) NOT NULL,
  sender VARCHAR(255) NOT NULL,
  message VARCHAR(255) NOT NULL,
  create_date DATE NOT NULL,
  PRIMARY KEY (id)
);

/*
to load to database:
psql -h localhost -U freejing -d portofolio -f ~/p/freejing-be/create_table.sql
*/