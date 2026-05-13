.PHONY: all help setup git-setup check-git \
        docker-up docker-down docker-logs docker-ps docker-build docker-pull docker-restart \
        deploy-all deploy-db deploy-redis deploy-backend deploy-frontend stop-db stop-redis stop-backend stop-frontend \
        backend-run frontend-setup frontend-dev frontend-build frontend-preview \
        clean

# --- OS Detection ---
ifeq ($(OS),Windows_NT)
    RM = rmdir /s /q
else
    RM = rm -rf
endif

# --- Configuration ---
BACKEND_DIR = backend
FRONTEND_DIR = web

# --- Default Target ---
all: help

help:
	@echo "🚀 Project Optimization Makefile"
	@echo "-------------------------------"
	@echo "setup              - Full optimized setup (Parallel)"
	@echo "git-setup          - Configure local Git identity"
	@echo "backend-run        - Start Go backend server"
	@echo "frontend-setup     - Optimized Node environment setup"
	@echo "frontend-dev       - Start frontend development"
	@echo "docker-up          - Spin up Docker services"
	@echo "deploy-all         - Build and deploy all containers"
	@echo "deploy-db          - Deploy only database container"
	@echo "deploy-redis       - Deploy only redis container"
	@echo "deploy-backend     - Build and deploy only backend container"
	@echo "deploy-frontend    - Build and deploy only frontend container"
	@echo "docker-ps          - Show running compose services"
	@echo "docker-logs        - Tail compose logs"
	@echo "docker-restart     - Restart all compose services"
	@echo "clean              - Remove build artifacts and environments"

# --- Main Setup (Optimized for Parallel Execution) ---
setup: git-setup
	@$(MAKE) -j 1 frontend-setup
	@echo "✅ All systems ready."

git-setup:
	@echo "📝 Configuring Git..."
	@printf "Enter Git Name: "; read name; \
	printf "Enter Git Email: "; read email; \
	git config user.name "$$name"; \
	git config user.email "$$email"
	@echo "Git configured as: $$(git config user.name) <$$(git config user.email)>"

backend-run:
	cd $(BACKEND_DIR) && go run ./cmd/web

# --- Frontend Optimization (File-based dependencies) ---
frontend-setup: $(FRONTEND_DIR)/node_modules/.bin

$(FRONTEND_DIR)/node_modules/.bin: $(FRONTEND_DIR)/package.json
	@echo "⚛️ Setting up Frontend Environment..."
	cd $(FRONTEND_DIR) && npm install
	@touch $@

# --- Development Commands ---
frontend-dev:
	cd $(FRONTEND_DIR) && npm run dev

frontend-build:
	cd $(FRONTEND_DIR) && npm run build

frontend-preview:
	cd $(FRONTEND_DIR) && npm run preview

# --- Docker Integration ---
docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-build:
	docker compose build

docker-pull:
	docker compose pull

docker-ps:
	docker compose ps

docker-logs:
	docker compose logs -f --tail=200

docker-restart:
	docker compose restart

deploy-all:
	docker compose up -d --build

deploy-db:
	docker compose up -d db

deploy-redis:
	docker compose up -d redis

deploy-backend:
	docker compose up -d --build backend

deploy-frontend:
	docker compose up -d --build frontend

stop-db:
	docker compose stop db

stop-redis:
	docker compose stop redis

stop-backend:
	docker compose stop backend

stop-frontend:
	docker compose stop frontend

db-shell:
	docker exec -it shp-postgres psql -U $${DATABASE_USER:-postgres} -d $${DATABASE_NAME:-shp_db}

# --- Cleanup ---
clean:
	@echo "🧹 Cleaning up..."
	@$(RM) $(FRONTEND_DIR)/node_modules
	@$(RM) $(FRONTEND_DIR)/dist
