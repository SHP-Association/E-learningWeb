import os
from pathlib import Path
from dotenv import load_dotenv

load_dotenv()

BASE_DIR = Path(__file__).resolve().parent.parent

# --- Core Settings ---

SECRET_KEY = os.getenv('DJANGO_SECRET_KEY')

DEBUG = os.getenv('DJANGO_DEBUG', 'False').lower() in ('true', '1', 't')

ALLOWED_HOSTS_STR = os.getenv('DJANGO_ALLOWED_HOSTS')
ALLOWED_HOSTS = [host.strip() for host in ALLOWED_HOSTS_STR.split(',') if host.strip()] if ALLOWED_HOSTS_STR else []

INSTALLED_APPS = [
    'jazzmin',
    'corsheaders',  # <-- Add this line at the top
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'rest_framework',
    'courses', 
    'Account',
    'Category',
    'Certificate',
    'Enrollment',
    'FAQ',
    'Lesson',
    'Question',
    'Quiz',
    'Review',
]

MIDDLEWARE = [
    'corsheaders.middleware.CorsMiddleware',  # <-- Add this line at the top
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
]

ROOT_URLCONF = os.getenv('DJANGO_ROOT_URLCONF')

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [BASE_DIR / 'templates'], # Ensure you have a 'templates' directory at project root
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]

WSGI_APPLICATION = os.getenv('DJANGO_WSGI_APPLICATION')

# --- Database Configuration ---
DATABASE_ENGINE = os.getenv('DJANGO_DB_ENGINE')
DATABASE_NAME = os.getenv('DJANGO_DB_NAME')

DATABASES = {
    'default': {
        'ENGINE': DATABASE_ENGINE,
        'NAME': DATABASE_NAME,
    }
}

AUTH_PASSWORD_VALIDATORS = [
    {'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator'},
    {'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator'},
    {'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator'},
    {'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator'},
]

LANGUAGE_CODE = 'en-us'
TIME_ZONE = os.getenv('DJANGO_TIME_ZONE')
USE_I18N = True
USE_TZ = True

STATIC_URL = '/static/'
STATICFILES_DIRS = [BASE_DIR / 'static']
STATIC_ROOT = BASE_DIR / 'staticfiles'
MEDIA_URL = '/media/'
MEDIA_ROOT = BASE_DIR / 'media'

DEFAULT_AUTO_FIELD = 'django.db.models.BigAutoField'
AUTH_USER_MODEL = 'Account.CustomUser'

REST_FRAMEWORK = {
    'DEFAULT_AUTHENTICATION_CLASSES': [
        'rest_framework.authentication.SessionAuthentication',
        'rest_framework.authentication.TokenAuthentication',
    ],
    'DEFAULT_PERMISSION_CLASSES': [
        'rest_framework.permissions.IsAuthenticatedOrReadOnly',
    ],
}

# --- Email Configuration ---
EMAIL_BACKEND = os.getenv('EMAIL_BACKEND')
EMAIL_HOST = os.getenv('EMAIL_HOST')
EMAIL_PORT = int(os.getenv('EMAIL_PORT')) if os.getenv('EMAIL_PORT') else None
EMAIL_USE_TLS = os.getenv('EMAIL_USE_TLS', 'False').lower() in ('true', '1', 't')
EMAIL_HOST_USER = os.getenv('EMAIL_HOST_USER')
EMAIL_HOST_PASSWORD = os.getenv('EMAIL_HOST_PASSWORD')
DEFAULT_FROM_EMAIL = os.getenv('DEFAULT_FROM_EMAIL')

# --- Security Settings ---
SECURE_SSL_REDIRECT = os.getenv('SECURE_SSL_REDIRECT', 'False').lower() in ('true', '1', 't') and not DEBUG
SESSION_COOKIE_SECURE = os.getenv('SESSION_COOKIE_SECURE', 'False').lower() in ('true', '1', 't') and not DEBUG
CSRF_COOKIE_SECURE = os.getenv('CSRF_COOKIE_SECURE', 'False').lower() in ('true', '1', 't') and not DEBUG
SECURE_BROWSER_XSS_FILTER = os.getenv('SECURE_BROWSER_XSS_FILTER', 'True').lower() in ('true', '1', 't')
SECURE_CONTENT_TYPE_NOSNIFF = os.getenv('SECURE_CONTENT_TYPE_NOSNIFF', 'True').lower() in ('true', '1', 't')

CSRF_TRUSTED_ORIGINS_STR = os.getenv('CSRF_TRUSTED_ORIGINS')
CSRF_TRUSTED_ORIGINS = [origin.strip() for origin in CSRF_TRUSTED_ORIGINS_STR.split(',') if origin.strip()] if CSRF_TRUSTED_ORIGINS_STR else []

CORS_ALLOWED_ORIGINS_STR = os.getenv('CORS_ALLOWED_ORIGINS')
CORS_ALLOWED_ORIGINS = [origin.strip() for origin in CORS_ALLOWED_ORIGINS_STR.split(',') if origin.strip()] if CORS_ALLOWED_ORIGINS_STR else []
CORS_ALLOW_CREDENTIALS = os.getenv('CORS_ALLOW_CREDENTIALS', 'False').lower() in ('true', '1', 't')

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
        {
            "name": "View Live Site",
            "url": os.getenv("FRONTEND_URL"),
            "new_window": True,
        },
        {
            "name": "Support & Docs",
            "url": "https://github.com/farridav/django-jazzmin/issues",
            "new_window": True,
        },
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
    "show_ui_builder": os.getenv('JAZZMIN_SHOW_UI_BUILDER', 'False').lower() in ('true', '1', 't'),

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
        "success": "btn-outline-success"
    }
}
