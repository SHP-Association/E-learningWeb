FROM python:3.11-slim

# Install system dependencies for psycopg2 and mysqlclient
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libmariadb-dev \
    libmariadb-dev-compat \
    libpq-dev \
    pkg-config \
    gcc \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy requirements first to leverage Docker cache
COPY requirements.txt .

# Upgrade pip, install psycopg and all requirements
RUN pip install --upgrade pip && \
    pip install psycopg[binary] && \
    pip install -r requirements.txt

# Copy the rest of the backend code
COPY . .

EXPOSE 8001

CMD ["python", "manage.py", "runserver", "0.0.0.0:8001"]
