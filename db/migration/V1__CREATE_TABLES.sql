CREATE TABLE "url" (
  "id" bigserial PRIMARY KEY,
  "url_encoded" varchar NOT NULL,
  "url_original" varchar NOT NULL,
  "expiration_date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Create a unique index on url_encoded and url_original
CREATE UNIQUE INDEX url_encoded_url_original_unique_idx ON "url" ("url_encoded", "url_original");

CREATE INDEX ON "url" ("created_at");

CREATE INDEX ON "url" ("url_encoded");

CREATE INDEX ON "url" ("expiration_date");