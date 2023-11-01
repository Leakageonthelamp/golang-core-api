dev:
	APP_HOST=:3000 APP_ENV=dev air

migrate:
	APP_SERVICE=migration go run main.go