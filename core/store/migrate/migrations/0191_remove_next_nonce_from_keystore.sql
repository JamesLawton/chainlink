-- +goose Up

ALTER TABLE evm_key_states DROP COLUMN next_nonce;
ALTER TABLE keys DROP COLUMN next_nonce;

-- +goose Down

ALTER TABLE evm_key_states ADD next_nonce bigint NOT NULL DEFAULT 0;
ALTER TABLE keys ADD next_nonce bigint NOT NULL DEFAULT 0;