CREATE TABLE IF NOT EXISTS categories
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX categories_name_unique ON categories (LOWER(name));

CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_timestamp
    BEFORE UPDATE
    ON categories
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();