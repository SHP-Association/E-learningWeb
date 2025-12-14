# 🚀 Backend Restart Instructions

## ⚠️ IMPORTANT: Restart Required

The Django backend needs to be restarted for CSRF changes to take effect.

## How to Restart:

### Option 1: Using Terminal
1. Go to your backend terminal
2. Press `Ctrl + C` to stop the server
3. Run: `python manage.py runserver`

### Option 2: Using PowerShell Script
```powershell
cd backend
.\start.ps1
```

## What Was Fixed:

✅ Added `http://localhost:5173` to `CSRF_TRUSTED_ORIGINS`
✅ Login will work after restart

## Verify It's Working:

After restart, check the terminal output for:
```
Starting development server at http://127.0.0.1:8000/
```

Then try logging in at http://localhost:5173/login
