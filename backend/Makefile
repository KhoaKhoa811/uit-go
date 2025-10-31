run-all:
	@echo "Starting all local Go services..."
	cd api-gateway && nohup make run-server &
	cd driver-service && nohup make run-server &
	cd trip-service && nohup make run-server &
	@echo "All local Go services started in background."

# Check if main service ports are active
view-ports:
	@echo "Checking active ports..."
	netstat -ano | findstr :3000 || true
	netstat -ano | findstr :50052 || true
	netstat -ano | findstr :50053 || true
	netstat -ano | findstr :8080 || true
	@echo "Done."

# Build Docker images for all services
docker-build:
	@echo "Building Docker images..."
	cd api-gateway && docker build -t uitgo/api-gateway:latest .
	cd driver-service && docker build -t uitgo/driver-service:latest .
	cd trip-service && docker build -t uitgo/trip-service:latest .
	cd user_service && docker build -t uitgo/user-service:latest .
	@echo "All Docker images built successfully."

# Push Docker images to registry
docker-push:
	@echo "Pushing Docker images to registry..."
	docker push uitgo/api-gateway:latest
	docker push uitgo/driver-service:latest
	docker push uitgo/trip-service:latest
	docker push uitgo/user-service:latest
	@echo "All images pushed successfully."

# Run full system with Docker Compose
run:
	@echo "Cleaning up old containers and images..."
	docker compose down --remove-orphans
	docker system prune -f
	@echo "Rebuilding all containers..."
	docker compose build --no-cache
	@echo "Starting all containers..."
	docker compose up -d

# Run system in detached mode
run-detach:
	@echo "Starting Docker Compose in detached mode..."
	docker compose up -d
	@echo "Containers are running in background."

# Stop and clean
clean:
	@echo "Stopping and removing all containers..."
	docker compose down --remove-orphans
	docker system prune -f
	@echo "Clean complete."
