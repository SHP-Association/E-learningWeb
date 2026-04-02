.PHONY: all help setup git-setup check-git \
        docker-up docker-down docker-logs \
        backend-setup backend-run backend-superuser backend-migrate backend-makemigrations \
        frontend-setup frontend-dev frontend-build frontend-preview \
        clean

# --- OS Detection ---
ifeq ($(OS),Windows_NT)
    PYTHON = python
    UV_INSTALL = pip install uv
    VENV_ACTIVATE = backend\.venv\Scripts\activate
    RM = rmdir /s /q
else
    PYTHON = python3
    UV_INSTALL = pip3 install uv --break-system-packages
    VENV_ACTIVATE = . backend/.venv/bin/activate
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
	@echo "backend-setup      - Optimized Python environment setup"
	@echo "backend-run        - Start backend Django server (uv)"
	@echo "backend-superuser  - Create a Django superuser (uv)"
	@echo "frontend-setup     - Optimized Node environment setup"
	@echo "frontend-dev       - Start frontend development"
	@echo "docker-up          - Spin up Docker services"
	@echo "clean              - Remove build artifacts and environments"

# --- Main Setup (Optimized for Parallel Execution) ---
setup: git-setup
	@$(MAKE) -j 2 backend-setup frontend-setup
	@echo "✅ All systems ready."

git-setup:
	@echo "📝 Configuring Git..."
	@printf "Enter Git Name: "; read name; \
	printf "Enter Git Email: "; read email; \
	git config user.name "$$name"; \
	git config user.email "$$email"
	@echo "Git configured as: $$(git config user.name) <$$(git config user.email)>"

# --- Backend Optimization ---
backend-setup: $(BACKEND_DIR)/.venv/touchfile

$(BACKEND_DIR)/.venv/touchfile: $(BACKEND_DIR)/requirements.txt
	@echo "🐍 Setting up Backend Environment..."
	@$(UV_INSTALL) || true
	@cd $(BACKEND_DIR) && if [ ! -d .venv ]; then uv venv; else echo "venv already exists, skipping creation."; fi
	cd $(BACKEND_DIR) && uv pip install -r requirements.txt
	@touch $@

backend-run:
	cd $(BACKEND_DIR) && uv run manage.py makemigrations
	cd $(BACKEND_DIR) && uv run manage.py migrate
	cd $(BACKEND_DIR) && uv run manage.py runserver 8001

backend-superuser:
	cd $(BACKEND_DIR) && uv run manage.py createsuperuser

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

backend-migrate:
	docker exec -it shp-backend python manage.py migrate

backend-makemigrations:
	docker exec -it shp-backend python manage.py makemigrations

db-shell:
	docker exec -it shp-postgres psql -U postgres -d postgres

# --- Cleanup ---
clean:
	@echo "🧹 Cleaning up..."
	@$(RM) $(BACKEND_DIR)/.venv
	@$(RM) $(FRONTEND_DIR)/node_modules
	@$(RM) $(FRONTEND_DIR)/dist
