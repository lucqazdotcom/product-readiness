begin
;

CREATE TABLE products (
    sku TEXT PRIMARY KEY,
    upc TEXT NOT NULL UNIQUE,
    product_name TEXT,
    brand TEXT,
    category TEXT,
    subcategory TEXT,
    color TEXT,
    size TEXT,
    unit_price NUMERIC(12, 2),
    cost_price NUMERIC(12, 2),
    currency TEXT,
    taxable BOOLEAN,
    status TEXT,
    supplier_id TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE product_descriptions (
    description_record_id TEXT PRIMARY KEY,
    sku TEXT NOT NULL REFERENCES products (sku),
    locale TEXT NOT NULL,
    short_description TEXT,
    long_description TEXT,
    material TEXT,
    care_instructions TEXT,
    bullet_1 TEXT,
    bullet_2 TEXT,
    seo_title TEXT,
    seo_description TEXT,
    approval_status TEXT,
    updated_at TIMESTAMPTZ,
    UNIQUE (sku, locale)
);

CREATE TABLE images (
    image_record_id TEXT PRIMARY KEY,
    sku TEXT NOT NULL REFERENCES products (sku),
    image_url TEXT,
    image_role TEXT,
    sort_order INTEGER,
    alt_text TEXT,
    width_px INTEGER,
    height_px INTEGER,
    is_active BOOLEAN,
    updated_at TIMESTAMPTZ
);

CREATE INDEX images_sku_idx ON images (sku);

CREATE TABLE inventory (
    inventory_record_id TEXT PRIMARY KEY,
    sku TEXT NOT NULL REFERENCES products (sku),
    warehouse_code TEXT NOT NULL,
    on_hand_qty INTEGER,
    reserved_qty INTEGER,
    available_qty INTEGER,
    reorder_point INTEGER,
    backorder_allowed BOOLEAN,
    inventory_status TEXT,
    last_counted_at TIMESTAMPTZ,
    feed_updated_at TIMESTAMPTZ,
    UNIQUE (sku, warehouse_code)
);

-- Keep UPCs as text to preserve leading zeros.
CREATE TABLE launch_dates (
    upc TEXT PRIMARY KEY REFERENCES products (upc),
    launch_week TEXT,
    launch_date DATE
);

commit
;
