-- +goose Up
-- +goose StatementBegin
CREATE TABLE "chapters" (
    "id" BIGSERIAL PRIMARY KEY,
    "content" text,
    "book_id" int,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "chapters";
-- +goose StatementEnd
