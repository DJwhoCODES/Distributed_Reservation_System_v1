Domain (internal/domain/)

The domain represents your business entities and business logic.
It is written in Go, not SQL.

type Movie struct {
ID string
Name string
DurationMin int
Language string
}

The domain should model the business, not the database.

---

Schema (Database Schema)

The schema is the structure of your database.

It defines:

tables
columns
types
constraints
indexes
foreign keys

Schema is the actual layout of your PostgreSQL database.

---

Migrations (migrations/\*.sql)

Migrations are version-controlled changes to the schema.
Instead of manually editing the database, you create migration files.
Running all migrations sequentially constructs the current database schema.

---

                 Your Go Code
                       |
        +--------------+--------------+
        |                             |
    Domain Models              Repositories
        |                             |
        +--------------+--------------+
                       |
                    PostgreSQL
                       |
          +------------+------------+
          |                         |
      Database Schema         Migration Files
