CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "first_name" varchar(100),
  "last_name" varchar(100),
  "email" varchar,
  "username" varchar,
  "password" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "stores" (
  "id" serial PRIMARY KEY,
  "store_name" varchar,
  "description" text,
  "contact_info" varchar,
  "owner_id" integer
);

CREATE TABLE "addresses" (
  "id" serial PRIMARY KEY,
  "user_id" integer,
  "store_id" integer,
  "street_address" varchar,
  "city" varchar,
  "state" varchar,
  "country" varchar,
  "zip_code" varchar
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "store_id" integer,
  "name" varchar,
  "description" varchar,
  "price" numeric(12,2),
  "image_url" varchar,
  "stock" integer,
  "category_id" integer,
  "created_at" timestamp
);

CREATE TABLE "parent_category" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "description" varchar
);

CREATE TABLE "category" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "parent_category_id" integer
);

CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "user_id" integer,
  "store_id" integer,
  "order_date" timestamp DEFAULT (now()),
  "total_amount" numeric(12,2),
  "status" varchar
);

CREATE TABLE "order_items" (
  "id" serial PRIMARY KEY,
  "order_id" integer,
  "product_id" integer,
  "quantity" integer,
  "price" numeric(12,2),
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "cart" (
  "id" serial PRIMARY KEY,
  "user_id" integer,
  "product_id" integer,
  "quantity" integer,
  "price" numeric(12,2)
);

ALTER TABLE "stores" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("parent_category_id") REFERENCES "parent_category" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "cart" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "cart" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
