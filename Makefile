run-nats:
	nats-server
run-service-config:
	go run app\services\config\main.go
run-service-auth:
	go run app\services\auth\main.go
run-caddy:
	caddy run --config=deploy\local\Caddyfile --envfile=.env