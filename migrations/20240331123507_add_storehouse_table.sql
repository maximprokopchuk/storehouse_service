-- +goose Up
-- +goose StatementBegin
CREATE TABLE storehouse (
  id      BIGSERIAL PRIMARY KEY,
  city_id INTEGER   NOT NULL,
  name    text      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE storehouse;
-- +goose StatementEnd
