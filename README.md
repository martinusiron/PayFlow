# 🧾 Payroll System

A complete payroll management backend system built with Go and PostgreSQL.

## Features

- Employee attendance tracking
- Overtime & reimbursement submission
- Admin-managed attendance periods
- Payroll processing with breakdowns
- Payslip generation
- Admin summaries
- Fully auditable system with logs

## Architecture

- **Clean Architecture** (domain, usecase, repository, delivery)
- **PostgreSQL** as the main DB
- **JSON-based HTTP API**
- **TDD** with unit + integration tests
- **Audit logging** and request traceability

## Technologies

- Go (1.23+)
- PostgreSQL
- Chi Router (or your router of choice)
- SQL migrations
- Docker (optional)

## Running Locally

1. Start PostgreSQL:
   ```bash
   docker-compose up -d
   ```
