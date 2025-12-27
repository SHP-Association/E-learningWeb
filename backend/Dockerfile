# Base Python image
FROM python:3.10
# Environment
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV PATH="/app/.local/bin:$PATH"

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    libpq-dev \
    default-libmysqlclient-dev \
    postgresql-client \
    netcat-traditional \
    curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY requirements.txt /app/
RUN pip install --upgrade pip
RUN pip install -r requirements.txt
RUN pip install celery[redis,beat] gunicorn

# Copy project code
COPY . /app/

# Expose port
EXPOSE 8001

# Startup script that handles migrations and runs appropriate server
CMD sh -c "python ./manage.py collectstatic --noinput && \
           python ./manage.py makemigrations && \
           python ./manage.py migrate && \
           python manage.py setup_roles && \
           python ./manage.py shell -c \"import os; from django.contrib.auth import get_user_model; User = get_user_model(); User.objects.filter(username=os.environ.get('DJANGO_SUPERUSER_USERNAME')).exists() or User.objects.create_superuser(username=os.environ.get('DJANGO_SUPERUSER_USERNAME'), email=os.environ.get('DJANGO_SUPERUSER_EMAIL'), password=os.environ.get('DJANGO_SUPERUSER_PASSWORD'))\" && \
           if [ \"$MODE\" = \"prod\" ]; then \
               gunicorn X.wsgi:application --bind 0.0.0.0:8001 --workers 4 --timeout 120; \
           else \
               python ./manage.py runserver 0.0.0.0:8001; \
           fi"