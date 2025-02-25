CREATE TABLE brand (
    id uuid NOT NULL,
    name VARCHAR NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NULL,
    CONSTRAINT brand_pkey PRIMARY KEY(id) 
);
CREATE INDEX brand_name_idx ON "brand" USING btree (name);