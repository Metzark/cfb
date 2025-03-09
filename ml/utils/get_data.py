import pandas as pd
import numpy as np
from pg.execute_stmt import execute_stmt

query = """
SELECT
	g.winner_id,
	g.loser_id,
	g.home_id,
	g.away_id,
	CASE 
		WHEN g.home_id = winner_id THEN g.home_final_score
		ELSE g.away_final_score
	END AS "winner_score",
	CASE 
		WHEN g.home_id = loser_id THEN g.home_final_score
		ELSE g.away_final_score
	END AS "loser_score",
	%s
FROM cfb.games g
JOIN cfb.stats ws ON ws.team_id = g.winner_id
JOIN cfb.stats ls ON ls.team_id = g.loser_id
WHERE
	g.season = %i
	AND g.winner_id != g.loser_id;
"""

# Target columns to exclude from features dataframe
target_col_names = ["winner_id", "loser_id", "home_id", "away_id", "winner_score", "loser_score"]

def get_data(pool, model: str, target: str, features: list[str], week: int, season: int):
    x = None
    y = None
    error = None
    conn = None
    cur = None

    try:
        # Get a connection from pg pool
        conn = pool.getconn()

        # Get cursor
        cur = conn.cursor()

        parsed_features, error = parse_features(features)

        if error:
            return x, y, error

        # Execute query to get feature and target values
        execute_stmt(cur, conn, query % (parsed_features, season))

        # Keep all rows
        rows = cur.fetchall()

        # Get column names
        col_names = [desc[0] for desc in cur.description]

        # New dataframe for all columns
        d = pd.DataFrame(data=rows, columns=col_names)

        # New dataframe containing only features
        x = d.drop(columns=target_col_names).to_numpy()

        # List containing target values
        y = d.apply(lambda row: get_target(row, model, target), axis=1).to_numpy()

        # Close cursor
        cur.close()

    except Exception as e:
        print(e)
        error = str(e)
    finally:
        # Close pg connection
        pool.putconn(conn)
    
    return x, y, error

# Get the target (y) list depending on the type of model
def get_target(row, model, target):
    match target:
        case "home_team_win":
            return 1 if row["winner_id"] == row["home_id"] else 0
        case "away_team_win":
            return 1 if row["winner_id"] == row["away_id"] else 0
        case _:
            return 1
        
def parse_features(features):
    try:
        parsed_features = []
        for i in range(len(features)):
            eq = features[i]["equation"]
            stats =features[i]["stats"]

            for j in range(len(stats)):
                eq = eq.replace("$" + str(j), stats[j])
            parsed_features.append("(" + eq + ")" + " AS stat_" + str(i))

            if "$" in eq:
                raise "Not all equation variables were replaced"
        
        return ",\n".join(parsed_features), None
    except Exception as e:
        return None, "Features contains an invalid feature equation"