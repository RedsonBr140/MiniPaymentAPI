-- Create enum type "funcao"
CREATE TYPE "funcao" AS ENUM ('COMUM', 'LOJISTA');
-- Create "users" table
CREATE TABLE "users" ("user_id" uuid NOT NULL DEFAULT gen_random_uuid(), "first_name" character varying(255) NOT NULL, "last_name" character varying(255) NOT NULL, "cpf" character varying(14) NOT NULL, "email" character varying(255) NOT NULL, "user_type" "funcao" NOT NULL, PRIMARY KEY ("user_id"));
-- Create index "users_cpf_key" to table: "users"
CREATE UNIQUE INDEX "users_cpf_key" ON "users" ("cpf");
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
CREATE TABLE "transactions" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "payer" uuid NOT NULL, "payee" uuid NOT NULL, "value" VARCHAR(255) NOT NULL)