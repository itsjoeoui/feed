CREATE TABLE IF NOT EXISTS "tweets" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "content" varchar NOT NULL,
  "user_id" INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "tweets" ("content");

CREATE INDEX ON "tweets" ("user_id");
