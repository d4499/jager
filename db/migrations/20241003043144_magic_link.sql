-- +goose Up
CREATE TABLE magic_links (
  id varchar(30) PRIMARY KEY,
  email varchar(255) NOT NULL,
  token varchar NOT NULL UNIQUE,
  expires_at timestamp
);

-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE magic_links;

-- +goose StatementBegin
-- +goose StatementEnd
