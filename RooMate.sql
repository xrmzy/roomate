CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "role_id" BIGINT NOT NULL,
  "role_name" VARCHAR,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "roles" (
  "id" BIGSERIAL PRIMARY KEY,
  "role_name" VARCHAR NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "customers" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "phone_number" VARCHAR NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "rooms" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "room_number" INT NOT NULL,
  "room_type" VARCHAR NOT NULL,
  "capacity" INT NOT NULL,
  "facility" VARCHAR NOT NULL,
  "price" BIGINT NOT NULL,
  "status" VARCHAR,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "services" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR,
  "price" BIGINT NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "bookings" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "night" INT,
  "check_in" DATE NOT NULL,
  "check_out" DATE NOT NULL,
  "user_id" UUID NOT NULL,
  "customer_id" UUID NOT NULL,
  "is_agree" BOOL,
  "information" VARCHAR,
  "total_price" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "booking_details" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "booking_id" UUID NOT NULL,
  "room_id" UUID NOT NULL,
  "sub_total" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

CREATE TABLE "booking_detail_services" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "booking_detail_id" UUID NOT NULL,
  "service_id" UUID NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL DEFAULT false
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "booking_details" ADD FOREIGN KEY ("booking_id") REFERENCES "bookings" ("id");

ALTER TABLE "booking_details" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "booking_detail_services" ADD FOREIGN KEY ("booking_detail_id") REFERENCES "booking_details" ("id");