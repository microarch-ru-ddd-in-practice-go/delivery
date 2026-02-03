# Демо проект к курсу "Domain Driven Design и Clean Architecture на языке Python"
📚 Подробнее о курсе: [microarch.ru/courses/ddd/languages/go](https://microarch.ru/courses/ddd/languages/go?utm_source=gitlab&utm_medium=repository&utm_content=basket)

---

## 🚀 Технологии

- **Python 3.13**
- **FastAPI** — веб-фреймворк
- **SQLAlchemy 2.0** (async) — ORM для PostgreSQL
- **Alembic** — миграции БД
- **Pydantic v2** — валидация данных
- **aiokafka** — async Kafka client
- **grpcio** — gRPC клиенты
- **pytest** — тестирование

## 📁 Структура проекта

```
app/
├── main.py                    # FastAPI entrypoint
├── core/                      # Ядро приложения
│   ├── config.py             # Настройки (Pydantic Settings)
│   ├── database.py           # SQLAlchemy engine, session
│   └── dependencies.py       # FastAPI Depends
├── api/                       # API слой
│   └── v1/
│       ├── router.py         # Главный роутер
│       └── endpoints/        # API endpoints
├── domain/                    # Доменный слой
│   ├── entities/             # Доменные сущности
│   ├── events/               # Доменные события
│   └── exceptions.py         # Доменные исключения
├── application/              # Слой приложения
│   ├── services/             # Бизнес-логика
│   ├── dto/                  # Data Transfer Objects
│   └── mediatr.py            # Медиатор событий
├── infrastructure/           # Инфраструктурный слой
│   ├── repositories/        # Реализация репозиториев
│   ├── external/             # Внешние интеграции
│   └── outbox/               # Outbox паттерн
└── consumers/                # Background workers
    ├── kafka_consumer.py     # Kafka consumers
    └── outbox_processor.py   # Outbox processor
```

## 🛠 Установка и запуск

### Локальная разработка

#### Вариант 1: С использованием uv (рекомендуется)

1. **Установите uv** (если ещё не установлен):
   ```bash
   # Windows (PowerShell)
   powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
   
   # Linux/Mac
   curl -LsSf https://astral.sh/uv/install.sh | sh
   ```

2. **Создайте virtualenv и установите зависимости:**

   **Windows (PowerShell):**
   ```powershell
   .\setup.ps1 -Install    # Создать venv и установить зависимости
   # или по отдельности:
   .\setup.ps1 -Venv       # Только создать virtualenv
   ```

   **Linux/Mac:**
   ```bash
   chmod +x setup.sh
   ./setup.sh install      # Создать venv и установить зависимости
   # или по отдельности:
   ./setup.sh venv         # Только создать virtualenv
   ```

   **Или используйте make (если установлен):**
   ```bash
   make venv        # Создать virtualenv
   make install     # Установить зависимости
   ```

   **Или вручную:**
   ```bash
   uv venv --python 3.13
   # Активируйте virtualenv:
   # Windows PowerShell: .venv\Scripts\Activate.ps1
   # Windows CMD: .venv\Scripts\activate.bat
   # Linux/Mac: source .venv/bin/activate
   uv pip install -r requirements.txt
   ```
   
   **Примечание для Windows:** Если получаете ошибку "running scripts is disabled", используйте один из вариантов:
   - CMD вместо PowerShell: `.venv\Scripts\activate.bat`
   - Или разрешите выполнение скриптов: `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser`

#### Вариант 2: С использованием pip

1. **Создайте virtualenv:**
   ```bash
   python -m venv .venv
   # Активируйте virtualenv:
   # Windows: .venv\Scripts\activate
   # Linux/Mac: source .venv/bin/activate
   ```

2. **Установите зависимости:**
   ```bash
   pip install -r requirements.txt
   ```

2. **Настройте переменные окружения:**
   ```bash
   cp .env.example .env
   # Отредактируйте .env файл
   ```

3. **Запустите миграции:**
   ```bash
   alembic upgrade head
   ```

4. **Запустите приложение:**
   ```bash
   uvicorn app.main:app --reload
   ```

5. **Запустите consumers (в отдельных терминалах):**
   ```bash
   # Kafka consumer
   python -m app.consumers.kafka_consumer

   # Outbox processor
   python -m app.consumers.outbox_processor
   ```

### Docker Compose

1. **Запустите все сервисы:**
   ```bash
   docker-compose up -d
   ```

2. **Примените миграции:**
   ```bash
   docker-compose exec api alembic upgrade head
   ```

3. **Проверьте здоровье сервиса:**
   ```bash
   curl http://localhost:8082/api/v1/health
   ```

## 📊 База данных

### Запросы к БД

```sql
-- Выборки
SELECT * FROM public.couriers;
SELECT * FROM public.storage_places;
SELECT * FROM public.orders;
SELECT * FROM public.outbox;

-- Очистка БД (все кроме справочников)
DELETE FROM public.couriers;
DELETE FROM public.storage_places;
DELETE FROM public.orders;
DELETE FROM public.outbox;

-- Добавить курьеров
-- Пеший
INSERT INTO public.couriers(
    id, name, speed, location_x, location_y)
VALUES ('bf79a004-56d7-4e5f-a21c-0a9e5e08d10d', 'Пеший', 1, 1, 1);

INSERT INTO storage_places (id, name, order_id, total_volume, courier_id)
VALUES 
  ('ed58fa74-b8fb-4a8c-a84b-e5c29ca9b0c6', 'Сумка', NULL, 10, 'bf79a004-56d7-4e5f-a21c-0a9e5e08d10d');

-- Вело
INSERT INTO public.couriers(
    id, name, speed, location_x, location_y)
VALUES ('db18375d-59a7-49d1-bd96-a1738adcee93', 'Вело', 2, 2, 2);

INSERT INTO storage_places (id, name, order_id, total_volume, courier_id)
VALUES 
  ('b96a9d83-aefa-4d06-99fb-e630d17c3868', 'Вело-Сумка', NULL, 10, 'db18375d-59a7-49d1-bd96-a1738adcee93'),
  ('838ac7aa-3f39-4b8a-b2be-f75fc3e35d34', 'Вело-Багажник', NULL, 30, 'db18375d-59a7-49d1-bd96-a1738adcee93');

-- Авто
INSERT INTO public.couriers(
    id, name, speed, location_x, location_y)
VALUES ('0f860f2c-d76a-4140-99b3-fcc63f27a826', 'Авто', 3, 3, 3);

INSERT INTO storage_places (id, name, order_id, total_volume, courier_id)
VALUES 
  ('f15b0f8c-dd93-4be6-a95a-3afd3a9f199e', 'Авто-Сумка', NULL, 10, '0f860f2c-d76a-4140-99b3-fcc63f27a826'),
  ('84e1ccae-555d-439c-8c87-dae080c82d29', 'Авто-Багажник', NULL, 50, '0f860f2c-d76a-4140-99b3-fcc63f27a826'),
  ('11fc6c0a-fc58-4718-b32d-8ce82e002201', 'Авто-Прицеп', NULL, 100, '0f860f2c-d76a-4140-99b3-fcc63f27a826');
```

### Миграции

```bash
# Создать новую миграцию
alembic revision --autogenerate -m "Описание изменений"

# Применить миграции
alembic upgrade head

# Откатить последнюю миграцию
alembic downgrade -1
```

## 🧪 Тестирование

```bash
# Запустить все тесты
pytest

# С покрытием
pytest --cov=app --cov-report=html

# Конкретный тест
pytest tests/unit/test_example.py
```

## 🔧 Разработка

### Code Quality

```bash
# Линтинг
make lint
# или
ruff check app/

# Форматирование
make format
# или
ruff format app/

# Проверка типов
make type-check
# или
mypy app/
```

### Полезные команды

**Windows (PowerShell):**
```powershell
.\setup.ps1 -Install    # Создать venv и установить зависимости
.\setup.ps1 -Venv       # Только создать virtualenv
.\setup.ps1 -Help       # Показать справку
```

**Linux/Mac или если установлен make:**
```bash
make help          # Показать все доступные команды
make venv          # Создать virtualenv через uv
make install       # Установить зависимости через uv
make test          # Запустить тесты
make run           # Запустить приложение
make migrate MESSAGE="описание"  # Создать миграцию
make upgrade       # Применить миграции
make downgrade     # Откатить миграцию
```

**Или используйте скрипты напрямую:**
```bash
# Linux/Mac
./setup.sh install
./setup.sh venv
```

### gRPC (генерация клиентов)

```bash
# Установите protoc и grpcio-tools
pip install grpcio-tools

# Скачайте proto файлы
curl -o ./api/proto/geo_service.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/geo/contracts/contract.proto

# Сгенерируйте Python код
python -m grpc_tools.protoc \
    --proto_path=./api/proto \
    --python_out=./app/infrastructure/external/grpc \
    --grpc_python_out=./app/infrastructure/external/grpc \
    ./api/proto/geo_service.proto
```

### Kafka (генерация интеграционных сообщений)

```bash
# Скачайте proto файлы
curl -o ./api/proto/basket_confirmed.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/basket/contracts/basket_confirmed.proto

curl -o ./api/proto/order_status_changed.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/delivery/contracts/order_status_changed.proto

# Сгенерируйте Python код (аналогично gRPC)
```

## 📚 Документация API

После запуска приложения доступна автоматическая документация:

- **Swagger UI**: http://localhost:8082/docs
- **ReDoc**: http://localhost:8082/redoc

## 📦 Используемые библиотеки

- [FastAPI](https://fastapi.tiangolo.com/) — современный веб-фреймворк
- [SQLAlchemy](https://www.sqlalchemy.org/) — ORM для работы с БД
- [Alembic](https://alembic.sqlalchemy.org/) — миграции БД
- [Pydantic](https://docs.pydantic.dev/) — валидация данных
- [aiokafka](https://aiokafka.readthedocs.io/) — async Kafka client
- [gRPC](https://grpc.io/docs/languages/python/) — gRPC для Python
- [pytest](https://docs.pytest.org/) — фреймворк для тестирования
- [testcontainers](https://testcontainers-python.readthedocs.io/) — тестовые контейнеры

## 📝 Лицензия

Код распространяется под лицензией [MIT](./LICENSE).  
© 2025 microarch.ru
