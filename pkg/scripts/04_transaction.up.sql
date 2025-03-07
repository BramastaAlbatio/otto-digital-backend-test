CREATE TABLE transaction (
    id uuid NOT NULL,
    customer_id uuid NOT NULL,
    total_point INT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NULL,
    CONSTRAINT transaction_pkey PRIMARY KEY(id), 
    CONSTRAINT transaction_customer_fk FOREIGN KEY (customer_id) REFERENCES customer(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX transaction_total_points_idx ON "transaction" USING btree (total_point);