-- Clear users if needed
DELETE FROM users;

-- Insert 1 admin
INSERT INTO users (username, password, salary, role)
VALUES (
  'admin',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZag0u6n4uKdK4KpN4Be1G5M0r8fO6', -- password123
  0,
  'admin'
);

-- Insert 100 fake employees
DO $$
DECLARE
  i INT;
  uname TEXT;
  sal NUMERIC;
BEGIN
  FOR i IN 1..100 LOOP
    uname := 'employee' || i;
    sal := round(random() * 5000 + 5000, 2);  -- Salary between 5000 and 10000
    INSERT INTO users (username, password, salary, role)
    VALUES (
      uname,
      '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZag0u6n4uKdK4KpN4Be1G5M0r8fO6',
      sal,
      'employee'
    );
  END LOOP;
END $$;
