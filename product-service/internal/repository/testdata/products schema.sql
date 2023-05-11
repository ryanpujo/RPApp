CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "store_id" integer NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" numeric(12,2) NOT NULL,
  "image_url" varchar NOT NULL,
  "stock" integer NOT NULL,
  "category_id" integer NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "stores" (
  "id" serial PRIMARY KEY,
  "store_name" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "category" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "parent_category_id" integer NOT NULL
);

CREATE TABLE "parent_category" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL
);

ALTER TABLE "products" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("parent_category_id") REFERENCES "parent_category" ("id");