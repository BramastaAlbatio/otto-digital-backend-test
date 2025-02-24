CREATE TABLE customer (
    id uuid NOT NULL,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT NOW(),           
	updated_at timestamptz DEFAULT NOW(),
    CONSTRAINT customer_pkey PRIMARY KEY(id) 
);
CREATE INDEX customer_name_idx ON "customer" USING btree (name);