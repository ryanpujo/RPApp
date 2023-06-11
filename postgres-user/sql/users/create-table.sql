CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(100) NOT NULL,
  "last_name" varchar(100) NOT NULL,
  "username" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);
