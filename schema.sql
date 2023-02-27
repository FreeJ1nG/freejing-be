CREATE TABLE blogs (
  id BIGSERIAL PRIMARY KEY,
  title text NOT NULL,
  content text NOT NULL,
  create_date DATE NOT NULL
);

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username text NOT NULL,
  email text NOT NULL,
  password_hash text NOT NULL
);

CREATE TABLE chat_history (
  id BIGSERIAL PRIMARY KEY,
  sender text NOT NULL,
  message text NOT NULL,
  create_date DATE NOT NULL
);