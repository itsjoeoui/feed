set dotenv-load := true

init:
  pnpm install

build:
  just assets
  templ generate
  go build -o ./tmp/main ./cmd/main.go

assets:
  pnpm exec tailwindcss -i ./internal/assets/tailwind.css -o ./internal/assets/dist/styles.css

migrateup:
  migrate -database ${DB_URL} -path internal/database/migrations up

migratedown:
  migrate -database ${DB_URL} -path internal/database/migrations down
