DROP TABLE IF EXISTS contracts CASCADE;

CREATE TABLE contracts (
    id SERIAL PRIMARY KEY,
    uuid TEXT NOT NULL UNIQUE,
    status TEXT CHECK (status IN ('currently', 'canceled', 'expired')) NOT NULL,
    stripe_subscription_id TEXT NOT NULL,
    stripe_price_id TEXT NOT NULL,
    stripe_product_id TEXT NOT NULL,
    signing_url TEXT NOT NULL,
    driver_cnh TEXT NOT NULL,
    school_cnpj TEXT NOT NULL,
    kid_rg TEXT NOT NULL,
    responsible_cpf TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    expire_at BIGINT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    anual_amount NUMERIC(10, 2) NOT NULL
);
