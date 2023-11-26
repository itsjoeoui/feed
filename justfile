set dotenv-load := true

migrateup:
  migrate -database ${DB_URL} -path internal/database/migrations up

migratedown:
  migrate -database ${DB_URL} -path internal/database/migrations down
