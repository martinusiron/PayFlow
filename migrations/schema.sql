CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    salary NUMERIC NOT NULL,
    role TEXT NOT NULL, -- 'admin' or 'employee'
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER,
    updated_by INTEGER,
    ip_address TEXT,
    request_id TEXT
);

CREATE TABLE IF NOT EXISTS attendance (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    attendance_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER,
    updated_by INTEGER,
    ip_address TEXT,
    request_id TEXT,
    UNIQUE(user_id, attendance_date)
);

CREATE TABLE IF NOT EXISTS overtime (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    overtime_date DATE NOT NULL,
    hours NUMERIC CHECK (hours <= 3),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER,
    updated_by INTEGER,
    ip_address TEXT,
    request_id TEXT
);

CREATE TABLE IF NOT EXISTS reimbursements (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    description TEXT,
    amount NUMERIC,
    reimbursement_date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER,
    updated_by INTEGER,
    ip_address TEXT,
    request_id TEXT
);

CREATE TABLE IF NOT EXISTS payroll_periods (
    id SERIAL PRIMARY KEY,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    run_at DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER,
    updated_by INTEGER,
    ip_address TEXT,
    request_id TEXT
);

CREATE TABLE IF NOT EXISTS payrolls (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    payroll_period_id INTEGER REFERENCES payroll_periods(id),
    base_salary NUMERIC,
    workdays_present NUMERIC NOT NULL,
    prorated_salary NUMERIC,
    overtime_hours int4 NOT NULL DEFAULT 0,
    overtime_amount NUMERIC,
    reimbursement_amount NUMERIC,
    total_amount NUMERIC,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER,
    updated_by INTEGER,
    ip_address TEXT,
    request_id TEXT,
    UNIQUE(user_id, payroll_period_id)
);

CREATE TABLE IF NOT EXISTS audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    action TEXT,
    table_name TEXT,
    record_id INTEGER,
    ip_address TEXT,
    request_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
