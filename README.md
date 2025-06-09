Payslip Generation System
This project implements a scalable payslip generation system in Go, designed to handle employee salaries, attendance, overtime, and reimbursement requests. It provides clear API endpoints for both administrative tasks and employee self-service, ensuring data traceability and adhering to predefined business rules.

Table of Contents
Project Overview

Features

Technical Requirements

Plus Points (Traceability)

Software Architecture

Setup Guide

Prerequisites

Database Setup

Configuration

Running the Application

API Usage

Authentication

Admin Endpoints

Employee Endpoints

Automated Testing

Future Enhancements

Contribution

License

Project Overview
The Payslip Generation System automates the calculation of employee take-home pay based on monthly salaries, prorated attendance, approved overtime, and reimbursements. It is built with Go and PostgreSQL, focusing on scalability, data integrity, and auditability.

Features
Employee Management:

Seeding of 100 fake employees with various salaries, usernames, and passwords.

Seeding of 1 fake admin user.

Payroll Period Management (Admin):

Define start and end dates for payroll periods.

Prevent overlapping payroll periods.

Ability to run payroll for a specific period (one-time processing).

Employee Self-Service:

Attendance Submission: Employees can submit their daily attendance (weekdays only, one submission per day counts).

Overtime Submission: Employees can propose overtime hours (max 3 hours per day, any day).

Reimbursement Requests: Employees can submit reimbursement requests with amounts and descriptions.

Payslip Generation:

Employees can generate detailed payslips including:

Breakdown of prorated salary based on attendance.

Breakdown of overtime compensation (twice the prorated hourly rate).

List of approved reimbursements.

Total take-home pay.

Payroll Summary (Admin):

Generate a summary of all employee payslips for a given period, including individual take-home pay and total company-wide payout.

Authentication & Authorization: JWT-based authentication with role-based access control (Admin vs. Employee).

Technical Requirements
Backend: Golang

Database: PostgreSQL

API: HTTP with JSON data format

Automated Testing: Unit and Integration tests for all core functionalities.

Plus Points (Traceability)
The system is designed with comprehensive traceability in mind:

Timestamps: created_at and updated_at columns are present in almost all tables, automatically recording when a record was created or last modified.

User Tracking: created_by and updated_by columns (UUIDs referencing the users table) track which user performed the create or update action.

IP Address Logging: The ip_address column in tables stores the IP address from which the last action was performed.

Audit Log Table: A dedicated audit_logs table captures significant actions (e.g., creating users, running payroll, major data updates) with details like user_id, action, entity_type, entity_id, and timestamp.

Request ID: Each incoming HTTP request is assigned a unique request_id (added to X-Request-ID header and context), which is also stored in audit_logs for end-to-end request tracing.

Software Architecture
The application follows a layered architecture, promoting separation of concerns and maintainability:

main.go: The entry point of the application, responsible for loading configuration, connecting to the database, initializing handlers, and setting up the HTTP router with middleware.

config/: Manages application configuration, loaded from environment variables and .env files.

db/: Handles database connection establishment (postgres.go).

models/: Defines the Go structs that map to database tables, representing the data models (e.g., User, Employee, Payslip).

utils/: Contains utility functions such as password hashing, JWT token generation/parsing, JSON response helpers, and general helper functions (e.g., date calculations, float rounding).

middlewares/: Implements HTTP middleware for cross-cutting concerns like:

Authentication (AuthMiddleware).

Role-based authorization (AdminRoleMiddleware, EmployeeRoleMiddleware).

Request ID generation (RequestIDMiddleware).

IP address logging (IPLoggerMiddleware).

Audit logging for all requests (AuditMiddleware).

services/: Contains the core business logic. Each service (e.g., AuthService, AdminService, EmployeeService) interacts with the database and encapsulates specific domain operations. This layer is independent of the HTTP context.

handlers/: Responsible for handling incoming HTTP requests, decoding JSON payloads, calling appropriate service methods, and sending JSON responses. They act as the interface between the HTTP layer and the business logic layer.

scripts/: Contains one-off scripts, such as seed_data.go, used for populating the database with initial or fake data for development and testing.

Setup Guide
Prerequisites
Go: Version 1.18 or higher.

PostgreSQL: Server installed and running (e.g., Docker, local installation).

Database Setup
Create Databases: Create two PostgreSQL databases: one for development (e.g., payslip_db) and one for testing (e.g., payslip_test_db).

CREATE DATABASE payslip_db;
CREATE DATABASE payslip_test_db;

Apply Schema: Apply the database schema from the db_schema_regenerate immersive to both payslip_db and payslip_test_db.

Connect to payslip_db (or payslip_test_db) using your preferred client (e.g., psql, pgAdmin, DBeaver).

Execute the SQL commands provided in the db_schema_regenerate immersive.

Configuration
Project Structure:
Create the following directory structure in your project root (payslip-system/):

payslip-system/
├── config/
│   └── config.go
├── db/
│   └── postgres.go
├── handlers/
│   ├── admin_handler.go
│   ├── auth_handler.go
│   └── employee_handler.go
├── middlewares/
│   └── auth_middleware.go
├── models/
│   └── user.go
├── scripts/
│   └── seed_data.go
├── services/
│   ├── admin_service.go
│   ├── auth_service.go
│   └── employee_service.go
├── utils/
│   ├── helpers.go
│   ├── jwt.go
│   ├── password.go
│   └── response.go
├── main.go
├── .env
├── test.env
├── go.mod        <-- Will be generated by 'go mod init'
└── go.sum        <-- Will be generated by 'go mod tidy'
└── README.md     <-- This file

Go Module Initialization:
Open your terminal in the payslip-system directory and run:

go mod init payslip-system
go mod tidy

This will initialize your Go module and download all necessary dependencies (github.com/gorilla/mux, github.com/lib/pq, github.com/google/uuid, github.com/joho/godotenv, github.com/golang-jwt/jwt/v5, golang.org/x/crypto/bcrypt).

Environment Files:
Create the .env and test.env files in your payslip-system root directory as provided in the corresponding immersives earlier (env_file_regenerate and test_env_regenerate). Remember to replace placeholder values with your actual database credentials.

Running the Application
Navigate to the payslip-system directory in your terminal.

Run the application:

go run main.go

You should see log messages indicating successful database connection and data seeding.

API Usage
All API requests should use Content-Type: application/json for request bodies and expect application/json in responses.

Authentication
Login

Endpoint: POST /api/v1/login

Description: Authenticates a user and returns a JWT token. This token must be included in the Authorization header (Bearer YOUR_TOKEN) for all subsequent protected API calls.

Request Body:

{
    "username": "admin",
    "password": "adminpass"
}

Successful Response (200 OK):

{
    "message": "Login successful",
    "data": {
        "token": "eyJhbGciOiJIUzI1Ni...",
        "user_id": "...",
        "username": "admin",
        "role": "admin"
    }
}

Admin Endpoints
Requires Authorization: Bearer YOUR_ADMIN_JWT_TOKEN and the user to have the admin role.

Add Payroll Period

Endpoint: POST /api/v1/admin/payroll-periods

Description: Defines a new payroll period.

Request Body:

{
    "start_date": "2024-06-01",
    "end_date": "2024-06-30"
}

Successful Response (201 Created): Returns the newly created payroll period details.

Run Payroll

Endpoint: POST /api/v1/admin/run-payroll/{payrollPeriodId}

Description: Processes all attendance, overtime, and reimbursement records for the specified payroll period and generates payslips for all employees. This can only be run once per period.

Path Parameter: payrollPeriodId (UUID)

Example: POST /api/v1/admin/run-payroll/a1b2c3d4-e5f6-7890-1234-567890abcdef

Successful Response (200 OK):

{
    "message": "Payroll for period a1b2c3d4-e5f6-7890-1234-567890abcdef completed successfully"
}

Get Payslips Summary

Endpoint: GET /api/v1/admin/payslips-summary

Description: Retrieves a summary of payslips for all employees within a given payroll period, including individual take-home pays and the total company payout.

Query Parameter: payrollPeriodId (UUID, required)

Example: GET /api/v1/admin/payslips-summary?payrollPeriodId=a1b2c3d4-e5f6-7890-1234-567890abcdef

Successful Response (200 OK):

{
    "message": "Payslips summary retrieved successfully",
    "data": {
        "payrollPeriodId": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
        "totalTakeHomePay": 12345678.90,
        "employeePayslips": [
            {
                "username": "employee1",
                "takeHomePay": 1234567.89
            },
            {
                "username": "employee2",
                "takeHomePay": 987654.32
            }
        ]
    }
}

Employee Endpoints
Requires Authorization: Bearer YOUR_EMPLOYEE_JWT_TOKEN and the user to have the employee role. The employeeID for these operations is derived from the authenticated user's ID.

Submit Attendance

Endpoint: POST /api/v1/employee/attendance

Description: Allows an employee to submit their attendance for a specific date. Submissions on weekends are not allowed, and only one submission per day is counted.

Request Body:

{
    "attendance_date": "2024-06-03"
}

Successful Response (201 Created): Returns the newly created attendance record.

Submit Overtime

Endpoint: POST /api/v1/employee/overtime

Description: Allows an employee to submit overtime hours for a specific date. Overtime cannot exceed 3 hours per day.

Request Body:

{
    "overtime_date": "2024-06-03",
    "hours": 2
}

Successful Response (201 Created): Returns the newly created overtime record.

Submit Reimbursement

Endpoint: POST /api/v1/employee/reimbursement

Description: Allows an employee to submit a reimbursement request.

Request Body:

{
    "amount": 75.50,
    "description": "Travel expenses for client meeting"
}

Successful Response (201 Created): Returns the newly created reimbursement record.

Get Payslip

Endpoint: GET /api/v1/employee/payslips

Description: Retrieves the detailed payslip for the authenticated employee for a specific payroll period.

Query Parameter: payrollPeriodId (UUID, required)

Example: GET /api/v1/employee/payslips?payrollPeriodId=a1b2c3d4-e5f6-7890-1234-567890abcdef

Successful Response (200 OK):

{
    "message": "Payslip retrieved successfully",
    "data": {
        "id": "...",
        "employee_id": "...",
        "payroll_period_id": "...",
        "base_salary": 5000000.00,
        "pro_rated_salary_component": 4545454.55,
        "overtime_component": 227272.73,
        "reimbursement_component": 150.75,
        "total_take_home_pay": 4772878.03,
        "details": {
            "attendedDays": 20,
            "baseSalary": 5000000.00,
            "overtimeCalculation": "28409.09 (hourly_rate) * 8 (overtime_hours) * 2 = 454545.46",
            "payrollPeriodEndDate": "2024-06-30",
            "payrollPeriodStartDate": "2024-06-01",
            "proRatedSalaryCalculation": "5000000.00 (base) / 22 (avg_working_days) * 20 (attended_days) = 4545454.55",
            "totalOvertimeHours": 8,
            "totalPossibleWorkingDays": 22,
            "totalReimbursementAmount": 150.75
        },
        "created_at": "2024-06-10T10:00:00Z"
    }
}

Automated Testing
The project includes unit and integration tests to ensure the correctness and reliability of the application's functionality.

Test Environment Setup:

Ensure your test.env file is configured correctly for a separate PostgreSQL test database.

Apply the database schema (from db_schema_regenerate immersive) to your payslip_test_db.

Running Tests:
Navigate to the payslip-system directory in your terminal and run:

go test ./...

This command will discover and run all tests in the current module. The TestMain function in the services/auth_service_test.go file handles setting up and tearing down the test database for a clean test run.

Future Enhancements
API Documentation: Integrate a tool like Swagger/OpenAPI for interactive API documentation.

Error Handling: Implement more granular error handling and custom error types.

Input Validation: Add more robust input validation (e.g., using a validation library) to handler inputs.

Concurrency: Explore Go routines and channels for concurrent processing of payslip calculations for very large datasets, if performance becomes a bottleneck.

Security: Implement rate limiting, proper TLS/SSL, and more advanced security headers.

Dockerization: Provide Dockerfiles for easy deployment and local development setup.

UI: Develop a simple web-based UI for interacting with the system.

Contribution
Feel free to fork this repository, submit pull requests, or open issues for any