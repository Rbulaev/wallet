DROP TABLE IF EXISTS wallets CASCADE;

CREATE TABLE wallets
(
    id           UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    amount       NUMERIC(10,2),
)