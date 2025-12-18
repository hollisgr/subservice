# Subscription Service

## Overview

This project implements a simple **REST API** for managing subscriptions. It provides functionality for basic CRUDL operations. The backend uses **PostgreSQL** as the primary database and utilizes **Golang**, specifically the **pgx** library, for efficient communication with the database. Additionally, the **Gin** web framework is employed to handle HTTP requests effectively.

## Features

- Full CRUDL support for managing subscription entries via HTTP endpoints. Each record contains:
  1. Service name providing the subscription.
  2. Monthly subscription fee in rubles.
  3. User's unique identifier in UUID format.
  4. Subscription start date (month and year).
  5. Optional subscription end date.

- Expose an additional HTTP endpoint to calculate the total cost of all active subscriptions within a specified period filtered by both user ID and service name.

- Utilize PostgreSQL as the relational database system. Provide migrations for initializing the database structure.

- Cover application logic with comprehensive logging.

- Extract configuration parameters (such as DB credentials, API ports, etc.) into either `.env` or `.yaml` files.

- Document the implemented API using Swagger specification and generate interactive documentation accordingly.

- Allow running the entire application stack using Docker Compose.

## Quick Start Guide

### Prerequisites:
- Go (≥ 1.24.6);
- Postgresql;
- Goose (optional);
- Docker (optional).

### Running Locally:

- **Step 1**: Create a `config.env` file with environment variables, for example:

```bash
BIND_IP=127.0.0.1
LISTEN_PORT=8888
PSQL_HOST=your_db_host
PSQL_PORT=your_db_port
PSQL_NAME=your_db_name
PSQL_USER=your_db_user
PSQL_PASSWORD=your_db_password
LOG_LEVEL=warn
CORS_ALLOW_ORIGINS=http://127.0.0.1:8888
```

- **Step 2**: Install `goose` migration tool (optional):

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- **Step 3**: Apply database migrations (optional):

```bash
goose -dir=migrations postgres \
"host=your_db_host port=your_db_port dbname=your_db_name user=your_db_user password=your_db_password sslmode=disable" up
```

- **Step 4**: Build and run the server:

```bash
make build
make run
```

---

### Running with Docker:

- **Step 1**: Create a `config.env` file with environment variables, for example:

```bash
BIND_IP=0.0.0.0
LISTEN_PORT=8888
PSQL_HOST=subservice-db
PSQL_PORT=5432
PSQL_NAME=postgres
PSQL_USER=postgres
PSQL_PASSWORD=postgres
LOG_LEVEL=warn
CORS_ALLOW_ORIGINS=http://127.0.0.1:8888
DOCKER_SERVICE_PORT=8888
DOCKER_PSQL_PORT=25432
```

- **Step 2**: Start the containerized application:

```bash
make docker-compose-up
```

---