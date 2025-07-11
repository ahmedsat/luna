
-- ENGINEERS
CREATE TABLE engineers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    regions TEXT[] NOT NULL DEFAULT '{}'
);

-- FARMERS
CREATE TABLE farmers (
    national_id BIGINT PRIMARY KEY CHECK (char_length(national_id::text) = 14),
    full_name TEXT NOT NULL,
    gender TEXT NOT NULL CHECK (gender IN ('Male', 'Female')),
    personal_image TEXT,
    id_image_front TEXT,
    id_image_back TEXT,
    phone_number VARCHAR(20)
);
-- FARMS
CREATE TABLE farms (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    arabic_name TEXT NOT NULL,
    english_name TEXT NOT NULL,
    engineer_id INTEGER REFERENCES engineers(id) ON DELETE SET NULL,
    manager_id INTEGER REFERENCES farmers(id) ON DELETE SET NULL,
    total_area DOUBLE PRECISION NOT NULL,
    cultivated_area DOUBLE PRECISION NOT NULL DEFAULT 0,
    year_of_reclamation INTEGER,
    ownership_document TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- OWNERSHIP
CREATE TABLE ownership (
    farm_id INTEGER NOT NULL REFERENCES farms(id) ON DELETE CASCADE,
    farmer_id INTEGER NOT NULL REFERENCES farmers(id) ON DELETE CASCADE,
    share DOUBLE PRECISION NOT NULL CHECK (share >= 0 AND share <= 100),
    is_manager BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (farm_id, farmer_id)
);

-- LOCATIONS
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL REFERENCES farms(id) ON DELETE CASCADE,
    description TEXT NOT NULL
);

-- LIVESTOCK
CREATE TABLE livestock (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL REFERENCES farms(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    count INTEGER NOT NULL CHECK (count >= 0)
);

-- DEMOGRAPHICS
CREATE TABLE demographics (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL REFERENCES farms(id) ON DELETE CASCADE,
    age_group TEXT NOT NULL, -- e.g. "<18", "18-60", "60+"
    gender TEXT NOT NULL CHECK (gender IN ('Male', 'Female')),
    count INTEGER NOT NULL CHECK (count >= 0)
);

-- ATTACHMENTS
CREATE TABLE attachments (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL REFERENCES farms(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    file_path TEXT NOT NULL
);

-- AUDIT LOG
CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    action TEXT NOT NULL CHECK (action IN ('CREATE', 'UPDATE', 'DELETE')),
    table_name TEXT NOT NULL,
    user_id TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    old_value JSONB,
    new_value JSONB
);
