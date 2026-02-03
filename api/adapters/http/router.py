from fastapi import APIRouter

from api.adapters.http import health

router = APIRouter()

router.include_router(health.router, tags=["health"])
