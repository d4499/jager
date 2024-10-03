-- +goose Up
CREATE TABLE users (
  id varchar(30) PRIMARY KEY,
  email varchar(255) NOT NULL UNIQUE,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
  id varchar(30) PRIMARY KEY,
  user_id varchar(30) references users (id) NOT NULL,
  expires_at timestamp
);

-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE sessions;

DROP TABLE users;

-- +goose StatementBegin
-- +goose StatementEnd
