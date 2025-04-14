-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- Add foreign key: fk_booking_car
DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'fk_booking_car_ms'
        ) THEN
            ALTER TABLE booking
                ADD CONSTRAINT fk_booking_car_ms
                    FOREIGN KEY (car_vin)
                        REFERENCES car(vin)
                        ON DELETE CASCADE;
        END IF;
    END
$$;


-- Check: start_time <= end_time
DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'chk_start_time_before_end_time_ms'
        ) THEN
            ALTER TABLE booking
                ADD CONSTRAINT chk_start_time_before_end_time_ms
                    CHECK (start_time <= end_time);
        END IF;
    END
$$;

-- Check: status is valid
DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'chk_valid_status_ms'
        ) THEN
            ALTER TABLE booking
                ADD CONSTRAINT chk_valid_status_ms
                    CHECK (status IN ('pending', 'confirmed', 'completed', 'canceled'));
        END IF;
    END
$$;

-- Exclusion constraint: prevent overlapping bookings
DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'car_not_overbooked_ms'
        ) THEN
            ALTER TABLE booking
                ADD CONSTRAINT car_not_overbooked_ms
                    EXCLUDE USING GIST (
                    car_vin WITH =,
                    daterange(start_time, end_time, '[]') WITH &&
                    );
        END IF;
    END
$$;
