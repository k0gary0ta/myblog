-- DROP SCHEMA IF EXISTS dev;
-- CREATE SCHEMA dev;
USE dev;

-- DROP TABLE IF EXISTS album;
-- CREATE TABLE album (
--   ID         INT AUTO_INCREMENT NOT NULL,
--   Title      VARCHAR(128) NOT NULL,
--   artist     VARCHAR(255) NOT NULL,
--   price      DECIMAL(5,2) NOT NULL,
--   PRIMARY KEY (`ID`)
-- );

DROP TABLE IF EXISTS article;
DROP TABLE IF EXISTS blog;

CREATE TABLE blog (
  ID         VARBINARY(26) NOT NULL PRIMARY KEY,
  Title      VARCHAR(80) NOT NULL,
  Body   TEXT NOT NULL,
  CreatedAt DATETIME NOT NULL
  -- PRIMARY KEY (`ID`)
);

INSERT INTO blog (ID, Title, Body, CreatedAt) VALUES (
  '01G54DRCKVXP20M78QVKKVDT51',
  'blog 1',
  'blog 1 Body',
  '2022-06-09 00:00:00'
);

INSERT INTO blog (ID, Title, Body, CreatedAt) VALUES (
  '01G55RHE2WJS1JTZ7Q7YB8SA34',
  'blog 2',
  'blog 2 Body',
  '2022-06-10 09:00:00'
);

INSERT INTO blog (ID, Title, Body, CreatedAt) VALUES (
  '01G55RRGH1B36AX4YBWRGYM9VQ',
  'blog 3',
  'blog 3 Body',
  '2022-06-11 15:00:00'
);