CREATE TABLE transaction_voucher (
    id uuid NOT NULL,
    transaction_id uuid NOT NULL,
    voucher_id uuid NOT NULL,
    quantity INT NOT NULL,
    subtotal_points INT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL,
    CONSTRAINT transaction_voucher_pkey PRIMARY KEY(id), 
    CONSTRAINT transaction_voucher_transaction_fk FOREIGN KEY (transaction_id) REFERENCES transaction(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT transaction_voucher_voucher_fk FOREIGN KEY (voucher_id) REFERENCES voucher(id) ON DELETE CASCADE ON UPDATE CASCADE
);
