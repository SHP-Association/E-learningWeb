#!/bin/bash

# Script to run the Go backend server locally
# This loads .env from project root and overrides database host to localhost

echo "🚀 Starting E-Learning Go Backend (Local Development Mode)"
echo "📍 Loading environment from root .env file"
echo "📍 Database: localhost:5433"
echo ""

# Get the project root directory (two levels up from backend-go)
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

# Load environment variables from root .env file
if [ -f "$PROJECT_ROOT/.env" ]; then
    echo "✅ Loading environment from $PROJECT_ROOT/.env"
    set -a  # automatically export all variables
    source "$PROJECT_ROOT/.env"
    set +a
else
    echo "⚠️  Warning: $PROJECT_ROOT/.env not found"
fi

# Navigate to the server directory
cd "$(dirname "$0")/cmd/server"

# Override database host for local development and run the server
echo "🔧 Overriding GO_DB_HOST=localhost for local development"
echo ""
GO_DB_HOST=localhost GO_DB_PORT=5433 go run . serve

