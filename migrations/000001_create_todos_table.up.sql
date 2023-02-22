CREATE TABLE todos (
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "description" varchar NOT NULL,
    "author" integer NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW())
);


CREATE INDEX ON "todos" ("author");
