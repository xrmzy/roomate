CREATE TABLE "users" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(100) NOT NULL,
  "email" VARCHAR(50) NOT NULL,
  "password" VARCHAR(100) NOT NULL,
  "role_id" INT NOT NULL,
  "role_name" VARCHAR(30) NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

CREATE TABLE "customers" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(100) NOT NULL,
  "email" VARCHAR(50) NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(50) NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

CREATE TABLE "rooms" (
  "id" VARCHAR(100),
  "room_number" INT NOT NULL,
  "room_type" VARCHAR(100) NOT NULL,
  "capacity" INT NOT NULL,
  "facility" VARCHAR(100) NOT NULL,
  "price" BIGINT NOT NULL,
  "status" VARCHAR(30),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

CREATE TABLE "services" (
  "id" VARCHAR(100),
  "name" VARCHAR(100),
  "price" BIGINT NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

CREATE TABLE "bookings" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "night" INT,
  "check_in" DATE NOT NULL,
  "check_out" DATE NOT NULL,
  "user_id" UUID NOT NULL,
  "customer_id" UUID NOT NULL,
  "customer_name" VARCHAR(255),
  "status" BOOL,
  "information" VARCHAR(100),
  "total_price" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

CREATE TABLE "booking_details" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "booking_id" UUID NOT NULL,
  "room_id" UUID NOT NULL,
  "services_id" UUID NOT NULL,
  "sub_total" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

CREATE TABLE "roles" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "role_name" VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL,
  "is_deleted" BOOL
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "booking_details" ADD FOREIGN KEY ("booking_id") REFERENCES "bookings" ("id");

ALTER TABLE "booking_details" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "booking_details" ADD FOREIGN KEY ("services_id") REFERENCES "services" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
