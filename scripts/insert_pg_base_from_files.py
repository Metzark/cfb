import json
from functions.create_pg_conn import create_pg_conn
from functions.execute_stmt import execute_stmt

#region SQL Statements

insert_pg_venues =  """
    INSERT INTO cfb.venues (id, name, city, state, zip, country_code, location_x, location_y, elevation, year_constructed, capacity, dome, grass, timezone)
    VALUES %s
    ON CONFLICT (id) DO NOTHING;
"""

insert_pg_conferences = """
    INSERT INTO cfb.conferences (id, name, short_name, abbreviation)
    VALUES %s
    ON CONFLICT (id) DO NOTHING;
"""

insert_pg_teams = """
    INSERT INTO cfb.teams (id, school, mascot, abbreviation, alt_name1, alt_name2, alt_name3, color, alt_color, twitter, logo, conference_id, home_venue_id)
    VALUES %s
    ON CONFLICT (id) DO NOTHING;
"""

insert_pg_games = """
    INSERT INTO cfb.games (id, season, week, season_type, start_date, neutral_site, conference_game, venue_id, home_id, away_id, home_qtr1_score,
        home_qtr2_score, home_qtr3_score, home_qtr4_score, home_final_score, away_qtr1_score, away_qtr2_score, away_qtr3_score, away_qtr4_score,
        away_final_score, winner_id, loser_id, excitement_index)
    VALUES %s
    ON CONFLICT (id) DO NOTHING;
"""

#endregion

try:
    # Set up db connection
    conn, cur = create_pg_conn()

    #region Insert venues from venues.json

    with open("data/venues.json", "r") as file:
        data = json.load(file)["venues"]
    
    # Format venues for insertion values
    venues = [
        (
            venue["id"],
            venue["name"],
            venue["city"],
            venue["state"],
            venue["zip"],
            venue["country_code"],
            venue["location"]["x"] if venue["location"] else None,
            venue["location"]["y"] if venue["location"] else None,
            float(venue["elevation"]) if venue.get("elevation") else None,
            venue.get("year_constructed"),
            venue.get("capacity"),
            venue["dome"],
            venue.get("grass"),
            venue.get("timezone"),
        )
        for venue in data
    ]

    # Insert venues
    execute_stmt(cur, conn, insert_pg_venues, venues)

    print("Inserted venues")

    #endregion
    
    #region Insert conferences from conferences.json
    with open("data/conferences.json", "r") as file:
        data = json.load(file)["conferences"]
    
    # Format conferences for insertion values
    conferences = [
        (
            conference["id"],
            conference["name"],
            conference["short_name"],
            conference["abbreviation"]
        )
        for conference in data
    ]

    # Insert conferences
    execute_stmt(cur, conn, insert_pg_conferences, conferences)

    print("Inserted conferences")

    # Map conferences names to id's
    conference_name_to_id = {conference["name"]: conference["id"] for conference in data }

    #endregion

    #region Insert teams from teams.json
    with open("data/teams.json", "r") as file:
        data = json.load(file)["teams"]
    
    # Format teams for insertion values
    teams = [
        (
            team["id"],
            team["school"],
            team["mascot"],
            team["abbreviation"],
            team["alt_name1"],
            team["alt_name2"],
            team["alt_name3"],
            team["color"],
            team["alt_color"],
            team["twitter"],
            team["logos"][0] if team["logos"] and len(team["logos"]) > 0 else None,
            conference_name_to_id[team["conference"]] if team["conference"] else None,
            team["location"]["venue_id"] if team["location"] else None
        )
        for team in data
    ]

    # Insert teams
    execute_stmt(cur, conn, insert_pg_teams, teams)

    print("Inserted teams")

    #endregion

    #region Insert games from games.json
    with open("data/games.json", "r") as file:
        data = json.load(file)["games"]

    # Format games for insertion values
    games = [
        (
            game["id"],
            game["season"],
            game["week"],
            game["season_type"],
            game["start_date"],
            game["neutral_site"],
            game["conference_game"],
            game["venue_id"],
            game["home_id"],
            game["away_id"],
            game["home_line_scores"][0] if game["home_line_scores"] and len(game["home_line_scores"]) > 0 else None,
            game["home_line_scores"][1] if game["home_line_scores"] and len(game["home_line_scores"]) > 1 else None,
            game["home_line_scores"][2] if game["home_line_scores"] and len(game["home_line_scores"]) > 2 else None,
            game["home_line_scores"][3] if game["home_line_scores"] and len(game["home_line_scores"]) > 3 else None,
            game["home_line_scores"][0] + game["home_line_scores"][1] + game["home_line_scores"][2] + game["home_line_scores"][3] if game["home_line_scores"] and len(game["home_line_scores"]) > 3 else None,
            game["away_line_scores"][0] if game["away_line_scores"] and len(game["away_line_scores"]) > 0 else None,
            game["away_line_scores"][1] if game["away_line_scores"] and len(game["away_line_scores"]) > 1 else None,
            game["away_line_scores"][2] if game["away_line_scores"] and len(game["away_line_scores"]) > 2 else None,
            game["away_line_scores"][3] if game["away_line_scores"] and len(game["away_line_scores"]) > 3 else None,
            game["away_line_scores"][0] + game["away_line_scores"][1] + game["away_line_scores"][2] + game["away_line_scores"][3] if game["away_line_scores"] and len(game["away_line_scores"]) > 3 else None,
            game["home_id"] if game["home_line_scores"] and len(game["home_line_scores"]) > 3 and game["home_line_scores"][0] + game["home_line_scores"][1] + game["home_line_scores"][2] + game["home_line_scores"][3] > game["away_line_scores"][0] + game["away_line_scores"][1] + game["away_line_scores"][2] + game["away_line_scores"][3] else game["away_id"],
            game["home_id"] if game["home_line_scores"] and len(game["home_line_scores"]) > 3 and game["home_line_scores"][0] + game["home_line_scores"][1] + game["home_line_scores"][2] + game["home_line_scores"][3] < game["away_line_scores"][0] + game["away_line_scores"][1] + game["away_line_scores"][2] + game["away_line_scores"][3] else game["away_id"],
            game["excitement_index"]

        )
        for game in data
    ]

    # Insert games
    execute_stmt(cur, conn, insert_pg_games, games)

    print("Inserted games")

    #endregion

    # Close db connection
    conn.close()
    cur.close()

except Exception as e:
    print(e.s)
