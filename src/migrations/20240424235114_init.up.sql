CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS create_order (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    amount character varying NOT NULL,
    currency character varying NOT NULL,
    provider_type integer NOT NULL,
    created_at bigint NOT NULL
);
CREATE TABLE IF NOT EXISTS create_order_provider (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    create_order_id uuid NOT NULL REFERENCES create_order (id) ON DELETE CASCADE,
    payload jsonb
)