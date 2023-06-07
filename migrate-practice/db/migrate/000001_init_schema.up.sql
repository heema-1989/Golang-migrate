CREATE TABLE IF NOT EXISTS "users"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "update_at" TIMESTAMP NOT NULL DEFAULT (now())
);