# Dross - your hackable memory system/agent

Work in progress

# Misc

(re)Generate SQL queries by running `sqlc generate`. Must naturally have `sqlc` installed.

Run migrations with `migrate -database "sqlite3://test.db" -path internal/migrations up`. This requires go-migrate installed, with the sqlite3 module.

# TODO

* It would be cool to not require cgo to run this project, and right now our hard requirement is the mattn sqlite3 driver (it's used by go-migrate)
