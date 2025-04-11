CREATE TABLE IF NOT EXISTS car (
                     vin VARCHAR(17) PRIMARY KEY,
                     model VARCHAR(100) NOT NULL,
                     brand VARCHAR(100) NOT NULL,
                     image_url TEXT,
                     price_per_day DECIMAL(10,2) NOT NULL CHECK (price_per_day >= 0),
                     created_at TIMESTAMP DEFAULT NOW()
);
