CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TABLE rental_user (
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      name VARCHAR(100) NOT NULL,
                      email VARCHAR(255) UNIQUE NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE car (
                     vin VARCHAR(17) PRIMARY KEY,
                     model VARCHAR(100) NOT NULL,
                     brand VARCHAR(100) NOT NULL,
                     image_url TEXT,
                     price_per_day DECIMAL(10,2) NOT NULL CHECK (price_per_day >= 0),
                     created_at TIMESTAMP DEFAULT NOW()
                     ---year INT CHECK (year >= 1900 AND year <= EXTRACT(YEAR FROM NOW()))
);


CREATE TABLE booking (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         car_vin VARCHAR(17) NOT NULL,
                         customer_id UUID NOT NULL, -- Assuming customers have a UUID
                         start_time date NOT NULL,
                         end_time date NOT NULL,
                         status VARCHAR(50) DEFAULT 'pending' CHECK (status IN ('pending', 'confirmed', 'completed', 'canceled')),
                         created_at TIMESTAMP DEFAULT NOW(),

                         CONSTRAINT fk_booking_car FOREIGN KEY (car_vin) REFERENCES car(vin) ON DELETE CASCADE,
                         CONSTRAINT fk_booking_customer FOREIGN KEY (customer_id) REFERENCES rental_user(id) ON DELETE CASCADE,

                         CONSTRAINT chk_start_time_before_end_time CHECK (start_time < end_time),

                        -- Prevent overlapping bookings for the same car (PostgreSQL)
                        -- Instead of '[]' (inclusive range), we could use '[)' (start-inclusive, end-exclusive) to allow back-to-back bookings without gaps.
                         CONSTRAINT car_not_overbooked
                             EXCLUDE USING GIST (
                            car_vin WITH =,
                             DATERANGE(start_time, end_time, '[]') WITH &&
                        )
);