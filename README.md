
# Project Name

A brief, one-sentence description of what your Go and PostgreSQL application does (e.g., *A RESTful CRUD API built with Go and PostgreSQL*).

## Tech Stack

* **Language:** Go (1.23+)
* **Database:** PostgreSQL (16+)
* **Router / Framework:** `net/http` / Gin / Echo / Chi *(update as needed)*
* **Database Driver / ORM:** `sqlc`
* **Migrations:**  Goose 

## Prerequisites

Before running this project, ensure you have the following installed:

* [Go](https://go.dev) (version 1.23 or higher)
* [PostgreSQL](https://postgresql.org) (or [Docker](https://docker.com) to run Postgres in a container)
* [Goose](https://github.com/pressly/goose) *(optional, if running CLI migrations)*

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/montacore/blog_aggregator.git
cd blog_aggregator
```

### 2. Configuration

Create a configuration file in your home directory named `.gatorconfig.json`

```json
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
    "current_user_name": "your_username"
```

### 3. Start PostgreSQL

**Option A: Using Docker**
```bash
docker run --name gator-postgres \
  -e POSTGRES_PASSWORD=yourpassword \
  -p 5432:5432 -d postgres
```

**Option B: Local PostgreSQL Instance**
```bash
psql -U postgres -c "CREATE DATABASE your_db_name;"
```

### 4. Run Database Migrations

Apply the SQL schema migrations using Goose from the directory containing your migration files:

```bash
cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
```

- If you don't already have Goose installed on your system, you can install it by following the instructions at [Goose](https://github.com/pressly/goose)

### 5. Usage

```text
gator <command> [arguments]

Commands
register <username> - Register a new user and log in as them.

login <username> - Switch the current user.

users - List all registered users.

addfeed <name> <url> - Add a new RSS feed and follow it.

feeds - List all feeds in the system.

follow <url> - Follow an existing feed.

unfollow <url> - Unfollow a feed.

following - List all feeds followed by the current user.

agg <time_between_reqs> - Start the aggregator loop (e.g., gator agg 1m).

browse [limit] - View recent posts from followed feeds (default limit is 2).

```markdown
### Quick Start

1. Register your user:
   ```bash
   gator register alice

Add a feed:

gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"

Fetch posts:

gator agg 30s

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
