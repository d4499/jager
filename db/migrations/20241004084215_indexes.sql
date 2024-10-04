-- +goose Up
CREATE UNIQUE INDEX idx_users_email ON users (email);

CREATE UNIQUE INDEX idx_magic_links_token ON magic_links (token);

-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP INDEX idx_magic_links_token;

DROP INDEX idx_users_email;

-- +goose StatementBegin
-- +goose StatementEnd
