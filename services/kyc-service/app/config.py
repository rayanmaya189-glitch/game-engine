from pydantic_settings import BaseSettings
from typing import List


class Settings(BaseSettings):
    # Database
    database_url: str = "postgresql+asyncpg://postgres:postgres@localhost:5432/kyc_db"
    database_pool_size: int = 20
    database_max_overflow: int = 10

    # Redis
    redis_url: str = "redis://localhost:6379/0"

    # Server
    host: str = "0.0.0.0"
    port: int = 9031
    debug: bool = False

    # CORS
    cors_origins: List[str] = ["*"]

    # KYC Providers
    onfido_enabled: bool = False
    onfido_api_key: str = ""
    onfido_webhook_token: str = ""

    jumio_enabled: bool = False
    jumio_api_token: str = ""
    jumio_api_secret: str = ""

    sumsub_enabled: bool = False
    sumsub_app_token: str = ""
    sumsub_secret_key: str = ""

    # KYC Levels
    basic_max_deposit: float = 2000.0
    intermediate_max_deposit: float = 10000.0
    full_max_deposit: float = 50000.0

    # Document types
    allowed_document_types: List[str] = ["PASSPORT", "DRIVERS_LICENSE", "NATIONAL_ID", "RESIDENCE_PERMIT"]

    class Config:
        env_file = ".env"
        case_sensitive = False


settings = Settings()
