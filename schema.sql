CREATE TABLE storehouse (
  id      BIGSERIAL PRIMARY KEY,
  city_id INTEGER   NOT NULL,
  name    text      NOT NULL
);

CREATE TABLE storehouse_item (
  id            BIGSERIAL PRIMARY KEY,
  storehouse_id    INTEGER NOT NULL,
  component_id     INTEGER NOT NULL,
  count         INTEGER NOT NULL DEFAULT 1,

  FOREIGN KEY (storehouse_id) REFERENCES storehouse(id) ON DELETE CASCADE
);