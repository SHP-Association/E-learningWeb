# Base Python image
FROM python:3.10

# Environment
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV PATH="/app/.local/bin:$PATH"

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    libpq-dev \
    default-libmysqlclient-dev \
    postgresql-client \
    netcat-traditional \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY requirements.txt /app/
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install --no-cache-dir celery[redis,beat]

# Copy project code
COPY . /app/

# Expose port
EXPOSE 8000
CMD ["sh", "-c", "python ./manage.py makemigrations && \
                   python ./manage.py migrate && \
                   python manage.py setup_roles && \
                   python ./manage.py shell -c \"import os; from django.contrib.auth import get_user_model; User = get_user_model(); User.objects.filter(username=os.environ.get('DJANGO_SUPERUSER_USERNAME')).exists() or User.objects.create_superuser(username=os.environ.get('DJANGO_SUPERUSER_USERNAME'), email=os.environ.get('DJANGO_SUPERUSER_EMAIL'), password=os.environ.get('DJANGO_SUPERUSER_PASSWORD'))\" && \
                   python ./manage.py runserver 0.0.0.0:8001"]