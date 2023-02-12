from fastapi import APIRouter, HTTPException
from src.consumer import start_consumption

router = APIRouter(prefix="/v1")


@router.get(
    "/start_kafka",
)
async def start_kafka():
    try:     
      start_consumption()
    except Exception as exception:
        print(f"failed due to {exception}")
        raise HTTPException(
            status_code=403,
            detail=f"An unexpected error occurred while fetching events:{exception}",
        ) from exception
