# cutie

cutie is an opinionated go code generator for database models born out of a zealous drive to enforce separation and types for maintainability and testability.

Core tenants:

- Separation of Concern
- Type safety
- Transparency
- Testable at every layer

## Components

### pg_dump

pg_dump produces the SQL schema required for sqlc to generate model code.

### sqlc

sqlc generates type-safe go model and query code from raw SQL.

### gomock

gomock creates mock implementations of query interfaces generated by sqlc.
