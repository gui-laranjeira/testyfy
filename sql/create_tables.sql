CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "passwords" (
  "user_id" bigserial NOT NULL,
  "password" varchar NOT NULL,
  "code" SERIAL NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "groups" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "guests" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "guest_on_group" (
  "guest_id" bigserial NOT NULL,
  "group_id" bigserial NOT NULL
);

CREATE TABLE IF NOT EXISTS "tests" (
  "id" bigserial PRIMARY KEY,
  "group_id" bigserial NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "questions" (
  "id" bigserial PRIMARY KEY,
  "test_id" bigserial NOT NULL,
  "question" varchar NOT NULL,
  "answer_type_code" smallint NOT NULL,
  "first_option" varchar,
  "second_option" varchar,
  "third_option" varchar,
  "fourth_option" varchar,
  "fifth_option" varchar,
  "image_url" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "answers_type" (
  "answer_type_code" smallint NOT NULL,
  "answer_type_description" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "guest_on_test" (
  "id" bigserial PRIMARY KEY,
  "guest_id" bigserial NOT NULL,
  "test_id" bigserial NOT NULL,
  "test_url" varchar NOT NULL,
  "total_score" real,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "guest_on_question" (
  "guest_id" bigserial NOT NULL,
  "question_id" bigserial NOT NULL,
  "answer" varchar NOT NULL,
  "score" real,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "passwords" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "guests" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "guest_on_group" ADD FOREIGN KEY ("guest_id") REFERENCES "guests" ("id");

ALTER TABLE "guest_on_group" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "tests" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "questions" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");

ALTER TABLE "guest_on_test" ADD FOREIGN KEY ("guest_id") REFERENCES "guests" ("id");

ALTER TABLE "guest_on_test" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");

ALTER TABLE "guest_on_question" ADD FOREIGN KEY ("guest_id") REFERENCES "guests" ("id");

ALTER TABLE "guest_on_question" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");