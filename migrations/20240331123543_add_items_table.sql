
-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
  id            BIGSERIAL PRIMARY KEY,
  storehouse_id    INTEGER NOT NULL,
  detail_id     INTEGER NOT NULL,
  count         INTEGER NOT NULL DEFAULT 1,

  FOREIGN KEY (storehouse_id) REFERENCES storehouse(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
