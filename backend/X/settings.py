import os
from pathlib import Path
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

BASE_DIR = Path(__file__).resolve().parent.parent

# =========================
# 🔐 CORE SETTINGS
# =========================
SECRET_KEY = os.environ["DJANGO_SECRET_KEY"]
DEBUG = os.environ["DJANGO_DEBUG"].lower() in ("true", "1", "t")
ALLOWED_HOSTS = [h.strip() for h in os.environ["DJANGO_ALLOWED_HOSTS"].split(",")]

ROOT_URLCONF = os.environ["DJANGO_ROOT_URLCONF"]

# =========================
# 📦 INSTALLED APPS
# =========================
INSTALLED_APPS = [
    "jazzmin",
    "corsheaders",
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
    "rest_framework",
    "courses",
    "Account",
    "Category",
    "Certificate",
    "Enrollment",
    "FAQ",
    "Lesson",
    "Question",
    "Quiz",
    "Review",
]

# =========================
# ⚙️ MIDDLEWARE
# =========================
MIDDLEWARE = [
    "corsheaders.middleware.CorsMiddleware",
    "django.middleware.security.SecurityMiddleware",
    "django.contrib.sessions.middleware.SessionMiddleware",
    "django.middleware.common.CommonMiddleware",
    "django.middleware.csrf.CsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
]

# =========================
# 🎨 TEMPLATES
# =========================
TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [BASE_DIR / "templates"],
        "APP_DIRS": True,
        "OPTIONS": {
            "context_processors": [
                "django.template.context_processors.debug",
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
            ],
        },
    },
]

# =========================
# 🗄️ DATABASE
# =========================
DATABASES = {
    "default": {
        "ENGINE": os.environ.get("DJANGO_DB_ENGINE", "django.db.backends.sqlite3"),
        "NAME": os.environ.get("DJANGO_DB_NAME", BASE_DIR / "db.sqlite3"),
        "USER": os.environ.get("DJANGO_DB_USER", ""),
        "PASSWORD": os.environ.get("DJANGO_DB_PASSWORD", ""),
        "HOST": os.environ.get("DJANGO_DB_HOST", ""),
        "PORT": os.environ.get("DJANGO_DB_PORT", ""),
        "CONN_MAX_AGE": 600,  # Connection pooling (10 minutes)
        "OPTIONS": {
            "connect_timeout": 10,
        } if os.environ.get("DJANGO_DB_ENGINE") != "django.db.backends.sqlite3" else {},
    }
}

# =========================
# 🔒 PASSWORD VALIDATION
# =========================
AUTH_PASSWORD_VALIDATORS = [
    {"NAME": "django.contrib.auth.password_validation.UserAttributeSimilarityValidator"},
    {"NAME": "django.contrib.auth.password_validation.MinimumLengthValidator"},
    {"NAME": "django.contrib.auth.password_validation.CommonPasswordValidator"},
    {"NAME": "django.contrib.auth.password_validation.NumericPasswordValidator"},
]

# =========================
# 🌍 INTERNATIONALIZATION
# =========================
LANGUAGE_CODE = "en-us"
TIME_ZONE = "UTC"
USE_I18N = True
USE_TZ = True

# =========================
# 🖼️ STATIC & MEDIA
# =========================
STATIC_URL = "/static/"
STATICFILES_DIRS = [BASE_DIR / "static"]
STATIC_ROOT = BASE_DIR / "staticfiles"

MEDIA_URL = "/media/"
MEDIA_ROOT = BASE_DIR / "media"

DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
AUTH_USER_MODEL = "Account.CustomUser"

# =========================
# ⚙️ REST FRAMEWORK
# =========================
REST_FRAMEWORK = {
    "DEFAULT_AUTHENTICATION_CLASSES": [
        "rest_framework.authentication.SessionAuthentication",
        "rest_framework.authentication.TokenAuthentication",
    ],
    "DEFAULT_PERMISSION_CLASSES": [
        "rest_framework.permissions.IsAuthenticatedOrReadOnly",
    ],
    # Pagination
    "DEFAULT_PAGINATION_CLASS": "rest_framework.pagination.PageNumberPagination",
    "PAGE_SIZE": 20,
    # Rate Limiting (Throttling)
    "DEFAULT_THROTTLE_CLASSES": [
        "rest_framework.throttling.AnonRateThrottle",
        "rest_framework.throttling.UserRateThrottle",
    ],
    "DEFAULT_THROTTLE_RATES": {
        "anon": "100/hour",  # Anonymous users
        "user": "1000/hour",  # Authenticated users
    },
}

# =========================
# 📧 EMAIL SETTINGS
# =========================
EMAIL_BACKEND = os.environ.get("EMAIL_BACKEND", "django.core.mail.backends.console.EmailBackend")
EMAIL_HOST = os.environ.get("EMAIL_HOST", "smtp.gmail.com")
EMAIL_PORT = int(os.environ.get("EMAIL_PORT", "587"))
EMAIL_USE_TLS = os.environ.get("EMAIL_USE_TLS", "True").lower() in ("true", "1", "t")
EMAIL_HOST_USER = os.environ.get("EMAIL_HOST_USER", "noreply@shp-learner.com")
EMAIL_HOST_PASSWORD = os.environ.get("EMAIL_HOST_PASSWORD", "dummy-password")
DEFAULT_FROM_EMAIL = EMAIL_HOST_USER

# =========================
# 🔐 SECURITY SETTINGS
# =========================
SECURE_SSL_REDIRECT = os.environ.get("DJANGO_SECURE_SSL_REDIRECT", "False").lower() in ("true", "1", "t")
SESSION_COOKIE_SECURE = os.environ.get("DJANGO_SESSION_COOKIE_SECURE", "False").lower() in ("true", "1", "t")
CSRF_COOKIE_SECURE = os.environ.get("DJANGO_CSRF_COOKIE_SECURE", "False").lower() in ("true", "1", "t")

SECURE_BROWSER_XSS_FILTER = True
SECURE_CONTENT_TYPE_NOSNIFF = True

# CSRF Trusted Origins - Support both HTTP (dev) and HTTPS (prod)
allowed_hosts = os.environ.get("DJANGO_ALLOWED_HOSTS", "localhost").split(",")
CSRF_TRUSTED_ORIGINS = []

for host in allowed_hosts:
    host = host.strip()
    if host:
        # Always add HTTPS for production
        CSRF_TRUSTED_ORIGINS.append(f"https://{host}")
        # Add HTTP only in development
        if DEBUG:
            CSRF_TRUSTED_ORIGINS.append(f"http://{host}")

# Add frontend URL to CSRF trusted origins
if DEBUG:
    CSRF_TRUSTED_ORIGINS += [
        "http://localhost:5173",
        "http://127.0.0.1:5173",
        "http://localhost:8001",
        "http://127.0.0.1:8001",
        "http://localhost:80",
        "http://127.0.0.1:80",
        "http://localhost",
        "http://127.0.0.1"
    ]

# CORS Settings - CRITICAL: Must allow credentials for session authentication
CORS_ALLOWED_ORIGINS = [os.environ.get("VITE_API_URL", "http://localhost:")]
CORS_ALLOW_CREDENTIALS = True  # Required for session auth with credentials: 'include'

# Add development origins if in DEBUG mode
if DEBUG:
    CORS_ALLOWED_ORIGINS += [
        "http://localhost:5173",
        "http://127.0.0.1:5173",
        "http://localhost:80",
        "http://127.0.0.1:80",
        "http://localhost",
        "http://127.0.0.1",
    ]

# =========================
# 🎛️ JAZZMIN ADMIN UI
# =========================
JAZZMIN_SETTINGS = {
    "site_title": "SHP-Learner Admin",
    "site_header": "SHP-Learner",
    "site_brand": "SHP-Learner",
    "site_logo": "/img/logo.png",
    "login_logo": "/img/loginlogo.png",
    "site_icon": "/img/logo.png",
    "welcome_sign": "Welcome to the SHP-Learner Administration Panel",
    "copyright": "SHP-Learner Platform © 2025",
    "show_sidebar": True,
    "navigation_expanded": True,
    "order_with_respect_to": ["courses", "auth"],
    "search_model": ["courses.Course", "Account.CustomUser"],
    "topmenu_links": [
        {"name": "Dashboard", "url": "admin:index", "permissions": ["auth.view_user"]},
        {"name": "View Live Site", "url": os.environ.get("VITE_API_URL", "http://localhost:8001"), "new_window": True},
        {"name": "Support & Docs", "url": "https://github.com/farridav/django-jazzmin/issues", "new_window": True},
    ],
    "usermenu_links": [
        {"name": "My Profile", "url": "admin:Account_customuser_change", "permissions": ["Account.change_customuser"]},
    ],
    "user_avatar": "/img/logo.png",
    "custom_links": {
        "courses": [
            {
                "name": "Course Analytics",
                "url": "course_analytics",
                "icon": "fas fa-chart-bar",
                "permissions": ["courses.view_course"],
            }
        ]
    },
    "icons": {
        "auth": "fas fa-users-cog",
        "auth.Group": "fas fa-users",
        "Account.CustomUser": "fas fa-user-graduate",
        "courses.Course": "fas fa-book-open",
        "Category.Category": "fas fa-layer-group",
        "Certificate.Certificate": "fas fa-certificate",
        "Enrollment.Enrollment": "fas fa-user-check",
        "FAQ.FAQ": "fas fa-question",
        "Lesson.Lesson": "fas fa-chalkboard-teacher",
        "Question.Question": "fas fa-question-circle",
        "Question.AnswerChoice": "fas fa-check-circle",
        "Quiz.Quiz": "fas fa-poll",
        "Quiz.UserQuizAttempt": "fas fa-user-clock",
        "Review.Review": "fas fa-star",
    },
    "default_icon_parents": "fas fa-folder-open",
    "default_icon_children": "fas fa-file-alt",
    "related_modal_active": True,
    "use_google_fonts_cdn": True,
    "show_ui_builder": False,
    "changeform_format": "horizontal_tabs",
    "changeform_format_single": "single",
    "language_chooser": False,
    "themes": ["flatly", "cosmo", "litera", "lumen", "minty"],
    "default_theme": "flatly",
    "dark_mode_theme": "darkly",
    "button_classes": {
        "primary": "btn-outline-primary",
        "secondary": "btn-outline-secondary",
        "info": "btn-outline-info",
        "warning": "btn-outline-warning",
        "danger": "btn-outline-danger",
        "success": "btn-outline-success",
    },
}
