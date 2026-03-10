from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.api import fraud_routes

app = FastAPI(title="Fraud Detection Service", version="1.0.0")

app.add_middleware(CORSMiddleware, allow_origins=["*"], allow_credentials=True, allow_methods=["*"], allow_headers=["*"])
app.include_router(fraud_routes.router, prefix="/api/v1/fraud", tags=["Fraud"])

@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "fraud-service"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=9033)
