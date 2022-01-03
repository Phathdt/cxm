CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" text NOT NULL,
    "password" text NOT NULL,
    "created_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ON "users" ("username");
