CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "role_name" varchar NOT NULL,
  "description" varchar
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "code_verify_email" varchar,
  "code_reset_password" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "hashed_password" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_verified_email" bool NOT NULL DEFAULT 'false',
  "full_name" varchar NOT NULL,
  "role_id" bigint NOT NULL,
  "token" varchar UNIQUE,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "paths" (
  "id" bigserial PRIMARY KEY,
  "path_name" varchar NOT NULL,
  "path" varchar UNIQUE NOT NULL,
  "path_description" varchar
);

CREATE TABLE "access_paths" (
  "id" bigserial PRIMARY KEY,
  "role_id" bigint NOT NULL,
  "path_id" bigint NOT NULL
);

CREATE INDEX ON "roles" ("id");

CREATE INDEX ON "users" ("username");

CREATE UNIQUE INDEX ON "users" ("email");

CREATE INDEX ON "paths" ("id");

CREATE INDEX ON "access_paths" ("id");

CREATE INDEX ON "access_paths" ("role_id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "access_paths" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "access_paths" ADD FOREIGN KEY ("path_id") REFERENCES "paths" ("id");
