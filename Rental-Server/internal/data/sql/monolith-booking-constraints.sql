-- Add foreign key: fk_booking_customer

-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- Add foreign key: fk_booking_car
DO $$
BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'fk_booking_car'
        ) THEN
ALTER TABLE booking
    ADD CONSTRAINT fk_booking_car
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
            SELECT 1 FROM pg_constraint WHERE conname = 'chk_start_time_before_end_time'
        ) THEN
ALTER TABLE booking
    ADD CONSTRAINT chk_start_time_before_end_time
        CHECK (start_time <= end_time);
END IF;
END
$$;

-- Check: status is valid
DO $$
BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'chk_valid_status'
        ) THEN
ALTER TABLE booking
    ADD CONSTRAINT chk_valid_status
        CHECK (status IN ('pending', 'confirmed', 'completed', 'canceled'));
END IF;
END
$$;

-- Exclusion constraint: prevent overlapping bookings
DO $$
BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'car_not_overbooked'
        ) THEN
ALTER TABLE booking
    ADD CONSTRAINT car_not_overbooked
    EXCLUDE USING GIST (
                    car_vin WITH =,
                    daterange(start_time, end_time, '[]') WITH &&
                    );
END IF;
END
$$;


DO $$
BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'fk_booking_customer'
        ) THEN
ALTER TABLE booking
    ADD CONSTRAINT fk_booking_customer
        FOREIGN KEY (customer_id)
            REFERENCES rental_user(id)
            ON DELETE CASCADE;
END IF;
END
$$;