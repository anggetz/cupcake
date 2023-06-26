tags=v1

run-api:
	go run -tags $(tags) app\apis\$(run)\main.go