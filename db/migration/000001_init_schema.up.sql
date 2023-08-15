CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reports" (
  "id" bigserial PRIMARY KEY,
  "date" date NOT NULL,
  "outlet" varchar NOT NULL,
  "item" varchar NOT NULL,
  "problem" varchar NOT NULL,
  "status" varchar NOT NULL,
  "remark" varchar NOT NULL,
  "owner" bigint NOT NULL,
  "company_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "companies" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "owner" bigint NOT NULL
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "reports" ("id");

CREATE INDEX ON "reports" ("owner", "company_id");

CREATE INDEX ON "companies" ("id");

CREATE INDEX ON "companies" ("owner");

ALTER TABLE "reports" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "reports" ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id");

ALTER TABLE "companies" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");
