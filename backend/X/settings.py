from pathlib import Path

# Build paths inside the project like this: BASE_DIR / 'subdir'.
BASE_DIR = Path(__file__).resolve().parent.parent

# --- Core Settings ---

# SECURITY WARNING: keep the secret key used in production secret!
# Hardcoded for demonstration, but NEVER do this in real production.
SECRET_KEY = 'django-insecure-@pqg#nf+l=6lp)x5(+d8-(u$ud^1mujk5^(+@5lv3t)azj+jw*'

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True # Hardcoded to True for development

# ALLOWED_HOSTS needs to be a list of strings.
# Hardcoded for development and your Render deployment.
ALLOWED_HOSTS = ['127.0.0.1', 'localhost',]

# Application definition
INSTALLED_APPS = [
    'corsheaders',  # <-- Add this line at the top
    'jazzmin',
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'rest_framework',
    'courses', # Assuming 'courses' is your app name
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

# IMPORTANT: Replace 'X' with your actual Django project's main module name (e.g., 'myproject')
ROOT_URLCONF = 'X.urls' # Example: 'myproject.urls'

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

# IMPORTANT: Replace 'X' with your actual Django project's main module name (e.g., 'myproject')
WSGI_APPLICATION = 'X.wsgi.application' # Example: 'myproject.wsgi.application'

# --- Database Configuration ---
# Hardcoded to SQLite.
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': BASE_DIR / 'db.sqlite3',
    }
}

# Password validation
AUTH_PASSWORD_VALIDATORS = [
    {'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator'},
    {'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator'},
    {'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator'},
    {'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator'},
]

# Internationalization
LANGUAGE_CODE = 'en-us'
TIME_ZONE = 'UTC' # Consider 'Asia/Kolkata' if your primary users are in India for local time awareness
USE_I18N = True
USE_TZ = True # Django's recommended way to handle datetimes

# Static and media files
STATIC_URL = '/static/'
STATICFILES_DIRS = [BASE_DIR / 'static'] # Django will look for static files here in development
STATIC_ROOT = BASE_DIR / 'staticfiles' # Directory where `python manage.py collectstatic` will gather all static files for deployment
MEDIA_URL = '/media/'
MEDIA_ROOT = BASE_DIR / 'media' # Directory where user-uploaded media will be stored

# Default primary key field type
DEFAULT_AUTO_FIELD = 'django.db.models.BigAutoField'

# Custom User Model
AUTH_USER_MODEL = 'courses.CustomUser' # Ensure this model exists and is correctly defined in your 'courses' app

# Django REST Framework Settings
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
# SECURITY WARNING: Never hardcode sensitive information like passwords in a public repo!
# Hardcoded email settings.
EMAIL_BACKEND = 'django.core.mail.backends.smtp.EmailBackend'
EMAIL_HOST = 'smtp.gmail.com'
EMAIL_PORT = 587
EMAIL_USE_TLS = True
EMAIL_HOST_USER = 'sandeshpatel.sp.93@gmail.com'
EMAIL_HOST_PASSWORD = 'fkorislsvoxviqcn'
DEFAULT_FROM_EMAIL = 'sandesh.patel@reak.in'

# --- Security Settings ---
# These are crucial for production environments.
# Hardcoded to reflect DEBUG = True for local development.
SECURE_SSL_REDIRECT = False
SESSION_COOKIE_SECURE = False
CSRF_COOKIE_SECURE = False
SECURE_BROWSER_XSS_FILTER = True
SECURE_CONTENT_TYPE_NOSNIFF = True

# CSRF_TRUSTED_ORIGINS should include all domains that serve your site.
CSRF_TRUSTED_ORIGINS = [
    # 'https://sh-2eas.onrender.com',
    'http://127.0.0.1',
    'http://localhost',
    'http://localhost:5173',
    'http://127.0.0.1:5173',
]

# CORS settings for development (allow all origins)
# CORS_ALLOW_ALL_ORIGINS = True

CORS_ALLOWED_ORIGINS = [
    "http://localhost:5173",
    "http://127.0.0.1:5173",
]
CORS_ALLOW_CREDENTIALS = True

# For production, use:
# CORS_ALLOWED_ORIGINS = [
#     "http://localhost:5173",
#     "http://127.0.0.1:5173",
# ]

# --- Jazzmin Settings ---
# Reference: https://django-jazzmin.readthedocs.io/
JAZZMIN_SETTINGS = {
    "site_title": "SHP-Learner",
    "site_header": "SHP-Learner",
    "site_brand": "SHP-Learner",
    # IMPORTANT: Use STATIC_URL prefix for images that are collected by collectstatic
    "site_logo": "/img/logo.png",
    "login_logo": "/static/img/logo.png", # Adjusted path for consistency
    "site_icon": "/img/logo.png",
    "welcome_sign": "Welcome to SHP-Learner Admin",
    "copyright": "SHP-Learner Platform 2025",
    "search_model": ["courses.Course", "courses.CustomUser"],
    "user_avatar": "/static/img/logo.png",
    "topmenu_links": [
        {"name": "Home", "url": "admin:index", "permissions": ["auth.view_user"]},
        {"name": "View Site", "url": "/", "new_window": True},
        {"name": "Support", "url": "https://github.com/farridav/django-jazzmin/issues", "new_window": True},
    ],
    "usermenu_links": [
        # Ensure 'courses.change_customuser' permission is correctly set for the custom user model
        {"name": "Profile", "url": "admin:courses_customuser_change", "permissions": ["courses.change_customuser"]},
    ],
    "show_sidebar": True,
    "navigation_expanded": True,
    "hide_apps": [],
    "hide_models": [],
    "order_with_respect_to": ["courses", "auth"],
    "custom_links": {
        "courses": [
            {
                "name": "Course Analytics",
                "url": "course_analytics", # Ensure you have a URL named 'course_analytics' in your urls.py
                "icon": "fas fa-chart-line",
                "permissions": ["courses.view_course"],
            }
        ]
    },
    "icons": {
        "auth": "fas fa-users-cog",
        "courses.CustomUser": "fas fa-user",
        "auth.Group": "fas fa-users",
        "courses.Course": "fas fa-book",
        "courses.Lesson": "fas fa-chalkboard",
        "courses.Quiz": "fas fa-question-circle",
        "courses.Enrollment": "fas fa-user-check",
    },
    "default_icon_parents": "fas fa-chevron-circle-right",
    "default_icon_children": "fas fa-circle",
    "related_modal_active": True,
    "custom_css": None,
    "custom_js": None,
    "use_google_fonts_cdn": True,
    "show_ui_builder": True, # Set to False in production for security/performance
    "changeform_format": "horizontal_tabs",
    "changeform_format_single": "single",
    "language_chooser": False,
}

JAZZMIN_UI_TWEAKS = {
    "navbar_small_text": False,
    "footer_small_text": False,
    "body_small_text": False,
    "brand_small_text": False,
    "brand_colour": "navbar-primary",
    "accent": "accent-primary",
    "navbar": "navbar-dark",
    "no_navbar_border": False,
    "navbar_fixed": True,
    "layout_boxed": False,
    "footer_fixed": False,
    "sidebar_fixed": True,
    "sidebar": "sidebar-dark-primary",
    "sidebar_nav_small_text": False,
    "sidebar_disable_expand": False,
    "sidebar_nav_child_indent": True,
    "sidebar_nav_compact_style": False,
    "sidebar_nav_legacy_style": False,
    "sidebar_nav_flat_style": False,
    "theme": "default",
    "button_classes": {
        "primary": "btn-primary",
        "secondary": "btn-secondary",
        "info": "btn-info",
        "warning": "btn-warning",
        "danger": "btn-danger",
        "success": "btn-success"
    }
}