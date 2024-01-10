-- go-migrate up

CREATE TYPE "funcao" AS ENUM ('COMUM', 'LOJISTA');

CREATE TABLE "users" (
    "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "first_name" character varying(255) NOT NULL,
    "last_name" character varying(255) NOT NULL,
    "cpf" character varying(14) NOT NULL,
    "email" character varying(255) NOT NULL,
    "user_type" "funcao" NOT NULL,
    PRIMARY KEY ("user_id")
);

CREATE UNIQUE INDEX "users_cpf_key" ON "users" ("cpf");
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");