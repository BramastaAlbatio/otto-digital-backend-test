CREATE TABLE voucher (
    id uuid NOT NULL,
    brand_id uuid NOT NULL,
    name VARCHAR NOT NULL,
    cost_in_point INT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW(),   
    CONSTRAINT voucher_pkey PRIMARY KEY(id), 
    CONSTRAINT vouchers_brand_fk FOREIGN KEY (brand_id) REFERENCES brand(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX voucher_name_idx ON "voucher" USING btree (name);
