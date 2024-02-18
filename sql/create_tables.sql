CREATE TABLE IF NOT EXISTS "users" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "passwords" (
  "user_id" varchar NOT NULL,
  "password" varchar NOT NULL,
  "code" SERIAL NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "groups" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "guests" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "guest_on_group" (
  "guest_id" varchar NOT NULL,
  "group_id" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "tests" (
  "id" varchar PRIMARY KEY,
  "group_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "questions" (
  "id" varchar PRIMARY KEY,
  "test_id" varchar NOT NULL,
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
  "id" varchar PRIMARY KEY,
  "guest_id" varchar NOT NULL,
  "test_id" varchar NOT NULL,
  "test_url" varchar NOT NULL,
  "total_score" real,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "guest_on_question" (
  "guest_id" varchar NOT NULL,
  "question_id" varchar NOT NULL,
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