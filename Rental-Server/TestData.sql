--- Setup

INSERT INTO rental_user (name, email, password, created_at)
VALUES
    ('John Doe', 'john.doe@example.com', 'password123', NOW()),
    ( 'Jane Smith', 'jane.smith@example.com', 'password456', NOW());

INSERT INTO car (vin, model, brand, image_url, kilometers, price_per_day, created_at)
VALUES
    ('1HGBH41JXMN109186', 'Civic', 'Honda', 'http://example.com/civic.jpg', 10000, 30.00, NOW()),
    ('2HGBH41JXMN109187', 'Model 3', 'Tesla', 'http://example.com/model3.jpg', 5000, 80.00, NOW());

--- should work
INSERT INTO booking (car_vin, customer_id, start_time, end_time, status, created_at)
VALUES
    ( '1HGBH41JXMN109186',
      (SELECT id FROM rental_user WHERE email = 'john.doe@example.com'),
      '2025-03-16', '2025-03-18', 'pending', NOW());

INSERT INTO booking (car_vin, customer_id, start_time, end_time, status, created_at)
VALUES
    ( '1HGBH41JXMN109186',
      (SELECT id FROM rental_user WHERE email = 'john.doe@example.com'),
      '2025-03-20', '2025-03-25', 'pending', NOW());

--- should fail
INSERT INTO booking (car_vin, customer_id, start_time, end_time, status, created_at)
VALUES
    ( '1HGBH41JXMN109186',
      (SELECT id FROM rental_user WHERE email = 'john.doe@example.com'),
      '2025-03-17', '2025-03-18', 'pending', NOW());

INSERT INTO booking (car_vin, customer_id, start_time, end_time, status, created_at)
VALUES
    ( '1HGBH41JXMN109186',
      (SELECT id FROM rental_user WHERE email = 'john.doe@example.com'),
      '2025-03-19', '2025-03-18', 'pending', NOW());