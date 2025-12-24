# Quick Start Script for E-Learning Platform Backend

Write-Host "🚀 Starting E-Learning Platform Backend Setup..." -ForegroundColor Cyan
Write-Host ""

# Check if virtual environment exists
if (Test-Path ".\.venv") {
    Write-Host "✅ Virtual environment found" -ForegroundColor Green
} else {
    Write-Host "❌ Virtual environment not found. Creating one..." -ForegroundColor Yellow
    python -m venv .venv
    Write-Host "✅ Virtual environment created" -ForegroundColor Green
}

# Activate virtual environment
Write-Host "🔄 Activating virtual environment..." -ForegroundColor Cyan
& .\.venv\Scripts\Activate.ps1

# Install dependencies
Write-Host "📦 Installing dependencies..." -ForegroundColor Cyan
python -m pip install --upgrade pip
pip install -r requirements.txt

# Run migrations
Write-Host "🗄️  Running database migrations..." -ForegroundColor Cyan
python manage.py migrate

# Check if superuser exists
Write-Host ""
Write-Host "👤 Do you want to create a superuser? (Y/N)" -ForegroundColor Yellow
$createSuperuser = Read-Host
if ($createSuperuser -eq "Y" -or $createSuperuser -eq "y") {
    python manage.py createsuperuser
}

# Start server
Write-Host ""
Write-Host "✅ Setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "🌐 Starting development server..." -ForegroundColor Cyan
Write-Host "   Backend will be available at: http://localhost:8000" -ForegroundColor White
Write-Host "   Admin panel: http://localhost:8000/admin/" -ForegroundColor White
Write-Host ""
Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
Write-Host ""

python manage.py runserver
