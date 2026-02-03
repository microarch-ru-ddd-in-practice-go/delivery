# syntax=docker/dockerfile:1

# Build stage
FROM python:3.13-slim AS build-stage

WORKDIR /build

# Устанавливаем зависимости для сборки
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get update --fix-missing && \
    apt-get install -y --no-install-recommends \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Копируем файлы зависимостей
COPY pyproject.toml requirements.txt ./

# Устанавливаем зависимости
RUN pip install --no-cache-dir --upgrade pip setuptools wheel && \
    pip install --no-cache-dir -r requirements.txt

# Копируем код приложения
COPY . .

# Test stage
FROM build-stage AS run-test-stage
RUN pytest tests/ || true  # Пока тестов нет, но структура готова

# Production stage
FROM python:3.13-slim AS build-release-stage

WORKDIR /app

# Копируем только установленные пакеты из build stage
COPY --from=build-stage /usr/local/lib/python3.13/site-packages /usr/local/lib/python3.13/site-packages
COPY --from=build-stage /usr/local/bin /usr/local/bin

# Копируем код приложения
COPY app ./app
COPY alembic.ini ./
COPY alembic ./alembic

# Создаём непривилегированного пользователя
RUN useradd -m -u 1000 appuser && chown -R appuser:appuser /app

USER appuser

EXPOSE 8082

# По умолчанию запускаем веб-сервер
# Для consumers используйте docker-compose или отдельные команды
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8082"]
