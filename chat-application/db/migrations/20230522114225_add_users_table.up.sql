CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "users"(
    "id" bigserial PRIMARY KEY,
    "verify_id" UUID NOT NULL,
    "full_name" VARCHAR NOT NULL,
    "user_name" VARCHAR NOT NULL,
    "email" varchar NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "password_create_date" timestamp NOT NULL DEFAULT (now()),
    "password_update_date" timestamp
);

