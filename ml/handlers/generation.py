from utils.get_data import get_data
from models.logistic import create_logistic_model

async def handle_generation(pool, model: str, target: str, features: list[str], week: int, season: int):
    # Get features and target dataframes
    x, y, error = get_data(pool, model, target, features, week, season)

    if (error):
        return None, error
        
    # Switch on model and generate that model
    match model:
        case "logistic":
            return create_logistic_model(x, y)
            
        case _:
            return None, "Invalid model"



