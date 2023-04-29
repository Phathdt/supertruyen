-- +goose Up
-- +goose StatementBegin
CREATE TABLE "books" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text,
    "owner_id" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "books";
-- +goose StatementEnd
