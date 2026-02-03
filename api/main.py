from contextlib import asynccontextmanager

from fastapi import FastAPI

from api.adapters.http.router import router as v1_router
from config.config import settings


@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup
    yield
    # Shutdown


app = FastAPI(
    title="Delivery Service",
    description="Сервис доставки на основе DDD и Clean Architecture",
    version="1.0.0",
    lifespan=lifespan,
)

app.include_router(v1_router, prefix="/api/v1")

if __name__ == "__main__":
    import uvicorn

    uvicorn.run(
        "api.main:app",
        host="0.0.0.0",
        port=int(settings.http_port),
        reload=True,
    )
