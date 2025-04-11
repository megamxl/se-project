-- Step 1: Create the booking table without constraints
CREATE TABLE IF NOT EXISTS booking
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    car_vin     VARCHAR(17) NOT NULL,
    customer_id UUID        NOT NULL,
    start_time  DATE        NOT NULL,
    end_time    DATE        NOT NULL,
    status      VARCHAR(50)      DEFAULT 'pending',
    created_at  TIMESTAMP        DEFAULT NOW(),
    paidAmount  DECIMAL(10, 2),
    currency    Varchar(3)
);