FROM python:3.12-slim

WORKDIR /app

RUN pip install --no-cache-dir uvicorn[standard]==0.27.0 fastapi==0.109.0 grpcio==1.62.1 pydantic==2.5.3 redis==5.0.1 httpx==0.26.0

COPY . .

EXPOSE 9016 9116

CMD ["python", "-m", "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "9016"]
