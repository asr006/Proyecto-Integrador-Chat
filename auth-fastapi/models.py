from sqlmodel import SQLModel, Field
from typing import Optional


class Usuario(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)

    nombre: str
    apellido: str
    fecha_nacimiento: str
    correo: str
    nickname: str
    password: str