CREATE TABLE "rooms"(
    "id" bigserial PRIMARY KEY,
    "room_name" VARCHAR NOT NULL,
    "user_id" BIGINT NOT NULL,
    CONSTRAINT fk_user_id
                    FOREIGN KEY(user_id)
                    REFERENCES users(id)
);