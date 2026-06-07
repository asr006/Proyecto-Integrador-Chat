import base64
import os
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from dotenv import load_dotenv

load_dotenv()

PEPPER= os.getenv("SECRET_PEPPER")
AES_KEY= os.getenv("AES_KEY").encode('utf-8')


def cifrar_password(password:str) -> str:
    iv = os.urandom(16)
    texto_para_cifrar= password + PEPPER

    cipher = AES.new(AES_KEY, AES.MODE_CBC, iv)
    texto_cifrado = cipher.encrypt(pad(texto_para_cifrar.encode('utf-8'), AES.block_size))

    texto_completo= iv + texto_cifrado
    return base64.b64encode(texto_completo).decode('utf-8')


def verificar_password(password_intento: str, password_guardada: str) -> bool:
    try:
        datos_completos = base64.b64decode(password_guardada)
        
        iv = datos_completos[:16]
        texto_cifrado = datos_completos[16:]
        cipher = AES.new(AES_KEY, AES.MODE_CBC, iv)
        texto_descifrado = unpad(cipher.decrypt(texto_cifrado), AES.block_size).decode('utf-8')
        
        return texto_descifrado == (password_intento + PEPPER)
    except Exception:
        return False


