"""Конфигурация приложения."""

from pydantic import Field
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    """
    Настройки приложения.

    Загружаются из переменных окружения или .env файла.
    """

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        case_sensitive=False,
        extra="ignore",
    )

    # HTTP
    http_port: str = Field(default="8082", alias="HTTP_PORT")

    # Database
    db_host: str = Field(alias="DB_HOST")
    db_port: str = Field(alias="DB_PORT")
    db_user: str = Field(alias="DB_USER")
    db_password: str = Field(alias="DB_PASSWORD")
    db_name: str = Field(alias="DB_NAME")
    db_sslmode: str = Field(default="disable", alias="DB_SSLMODE")

    # gRPC
    geo_service_grpc_host: str = Field(alias="GEO_SERVICE_GRPC_HOST")

    # Kafka
    kafka_host: str = Field(alias="KAFKA_HOST")
    kafka_consumer_group: str = Field(alias="KAFKA_CONSUMER_GROUP")
    kafka_basket_confirmed_topic: str = Field(alias="KAFKA_BASKET_CONFIRMED_TOPIC")
    kafka_order_changed_topic: str = Field(alias="KAFKA_ORDER_CHANGED_TOPIC")

    @property
    def database_url(self) -> str:
        """
        Формирует URL для подключения к PostgreSQL.

        Returns:
            Строка подключения к базе данных.
        """
        return (
            f"postgresql+asyncpg://{self.db_user}:{self.db_password}"
            f"@{self.db_host}:{self.db_port}/{self.db_name}?sslmode={self.db_sslmode}"
        )

    @property
    def database_url_sync(self) -> str:
        """
        Формирует синхронный URL для подключения к PostgreSQL (для Alembic).

        Returns:
            Синхронная строка подключения к базе данных.
        """
        return (
            f"postgresql://{self.db_user}:{self.db_password}"
            f"@{self.db_host}:{self.db_port}/{self.db_name}?sslmode={self.db_sslmode}"
        )


settings = Settings()
