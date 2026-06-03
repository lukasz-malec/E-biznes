from fastapi import FastAPI
from pydantic import BaseModel
import requests

app = FastAPI()

historia = [
    {"role": "system", "content": "Jesteś pomocnym asystentem. Odpowiadaj po polsku, krótko i konkretnie."}
]


class Query(BaseModel):
    message: str


@app.post("/ask")
def ask(query: Query):
    historia.append({"role": "user", "content": query.message})

    wiadomosci = [historia[0]] + historia[-20:]

    try:
        response = requests.post("http://localhost:11434/api/chat", json={
            "model": "llama3.2:1b",
            "messages": wiadomosci,
            "stream": False
        })
        odpowiedz = response.json()["message"]["content"]
    except Exception as e:
        historia.pop()
        return {"answer": f"Błąd: {e}"}

    historia.append({"role": "assistant", "content": odpowiedz})
    return {"answer": odpowiedz}