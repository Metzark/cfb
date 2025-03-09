from fastapi import FastAPI, Request
from pg.create_pg_pool import create_pg_pool
from handlers.generation import handle_generation

app = FastAPI()

pool = create_pg_pool()

@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.post("/generate")
async def generate(request: Request):

    body = await request.json()

    model = body.get("model")
    season = body.get("season")
    week = body.get("week")
    target = body.get("target")
    features = body.get("features")

    if model == None or season == None or week == None or features == None:
        return {"error": "Must provide a model, season, week, and features list"}

    accuracy, error = await handle_generation(pool, model, target, features, week, season)

    return {"accuracy": accuracy, "error": error}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)