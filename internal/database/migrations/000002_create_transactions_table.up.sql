CREATE TABLE "transactions" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "payer" uuid NOT NULL,
    "payee" uuid NOT NULL,
    "value" DECIMAL(18, 2) NOT NULL
);
