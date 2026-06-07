from fastapi import FastAPI
from sqlmodel import SQLModel
from app.routes.auth import router as auth_router, engine

app = FastAPI()

@app.on_event("startup")
def iniciar_base_datos():
    SQLModel.metadata.create_all(engine)

app.include_router(auth_router)