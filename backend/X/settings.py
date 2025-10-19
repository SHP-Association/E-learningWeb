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
        "ENGINE": os.environ["DJANGO_DB_ENGINE"],
        "NAME": os.environ["DJANGO_DB_NAME"],
        "USER": os.environ["DJANGO_DB_USER"],
        "PASSWORD": os.environ["DJANGO_DB_PASSWORD"],
        "HOST": os.environ["DJANGO_DB_HOST"],
        "PORT": os.environ["DJANGO_DB_PORT"],
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
}

# =========================
# 📧 EMAIL SETTINGS
# =========================
EMAIL_BACKEND = os.environ["EMAIL_BACKEND"]
EMAIL_HOST = os.environ["EMAIL_HOST"]
EMAIL_PORT = int(os.environ["EMAIL_PORT"])
EMAIL_USE_TLS = os.environ["EMAIL_USE_TLS"].lower() in ("true", "1", "t")
EMAIL_HOST_USER = os.environ["EMAIL_HOST_USER"]
EMAIL_HOST_PASSWORD = os.environ["EMAIL_HOST_PASSWORD"]
DEFAULT_FROM_EMAIL = EMAIL_HOST_USER

# =========================
# 🔐 SECURITY SETTINGS
# =========================
SECURE_SSL_REDIRECT = os.environ["DJANGO_SECURE_SSL_REDIRECT"].lower() in ("true", "1", "t")
SESSION_COOKIE_SECURE = os.environ["DJANGO_SESSION_COOKIE_SECURE"].lower() in ("true", "1", "t")
CSRF_COOKIE_SECURE = os.environ["DJANGO_CSRF_COOKIE_SECURE"].lower() in ("true", "1", "t")

SECURE_BROWSER_XSS_FILTER = True
SECURE_CONTENT_TYPE_NOSNIFF = True

CSRF_TRUSTED_ORIGINS = [f"http://{h.strip()}" for h in os.environ["DJANGO_ALLOWED_HOSTS"].split(",")]
CORS_ALLOWED_ORIGINS = [os.environ["VITE_API_URL"]]
CORS_ALLOW_CREDENTIALS = False

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
        {"name": "View Live Site", "url": os.environ["VITE_API_URL"], "new_window": True},
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
