run-all:
	cd api-gateway && nohup make run-server &
	cd driver-service && nohup make run-server &
	cd trip-service && nohup make run-server &

view-ports:
	netstat -ano | findstr :3000
	netstat -ano | findstr :50052
	netstat -ano | findstr :50053

docker-build:
	cd api-gateway && docker build -t uitgo/api-gateway .
	cd driver-service && docker build -t uitgo/driver-service .
	cd trip-service && docker build -t uitgo/trip-service .

docker-push:
	cd api-gateway && docker push uitgo/api-gateway
	cd driver-service && docker push uitgo/driver-service
	cd trip-service && docker push uitgo/trip-service

run:
	docker compose down --remove-orphans
	docker system prune -f
	docker compose build --no-cache
	docker compose up
