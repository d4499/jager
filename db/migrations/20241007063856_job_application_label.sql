-- +goose Up
CREATE TABLE labels (
  id varchar(30) PRIMARY KEY,
  title varchar(32) NOT NULL,
  user_id varchar(30) REFERENCES users (id) NOT NULL
);

CREATE TABLE job_applications (
  id varchar(30) PRIMARY KEY,
  title varchar(32) NOT NULL,
  company varchar(32) NOT NULL,
  applied_date timestamp,
  user_id varchar(30) REFERENCES users (id) NOT NULL,
  label_id varchar(30) REFERENCES labels (id),
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE job_applications;

DROP TABLE labels;

-- +goose StatementBegin
-- +goose StatementEnd
