DROP TABLE IF EXISTS blog;

CREATE TABLE blog (
  id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(10000) NOT NULL,
  create_date VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);