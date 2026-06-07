from fastapi import APIRouter
from sqlmodel import Session, select, create_engine
from models import Usuario
from app.routes.encrypt import cifrar_password, verificar_password


engine = create_engine("sqlite:///usuarios.db")
router= APIRouter()

@router.post("/registro")
def registro(usuario: Usuario):

    usuario.password = cifrar_password(usuario.password)

    with Session(engine) as session:
        session.add(usuario)
        session.commit()

    return {"mensaje": "Usuario registrado"}


@router.post("/login")
def login(datos: dict):

    with Session(engine) as session:

        usuario = session.exec(
            select(Usuario).where(
                Usuario.nickname == datos["nickname"]
            )
        ).first()

        if not usuario:
            return {"mensaje": "Usuario no encontrado"}

        if verificar_password(
            datos["password"],
            usuario.password
        ):
            return {"mensaje": "Login correcto"}

        return {"mensaje": "Contraseña incorrecta"}