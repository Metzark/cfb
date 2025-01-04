from functions.create_pg_conn import create_pg_conn
from functions.execute_stmt import execute_stmt

#region SQL Statements

create_cfb_schema_stmt = "CREATE SCHEMA IF NOT EXISTS cfb;"

create_venues_table_stmt = """
    CREATE TABLE IF NOT EXISTS cfb.venues (
        id INTEGER PRIMARY KEY,
        name TEXT,
        city TEXT,
        state TEXT,
        zip TEXT,
        country_code TEXT,
        location_x DOUBLE PRECISION,
        location_y DOUBLE PRECISION,
        elevation DOUBLE PRECISION,
        year_constructed INTEGER,
        capacity INTEGER,
        dome BOOLEAN,
        grass BOOLEAN,
        timezone TEXT
    );
"""

create_conferences_table_stmt = """
    CREATE TABLE IF NOT EXISTS cfb.conferences (
        id INTEGER PRIMARY KEY,
        name TEXT,
        short_name TEXT,
        abbreviation TEXT
    );
"""

create_teams_table_stmt = """
    CREATE TABLE IF NOT EXISTS cfb.teams (
        id INTEGER PRIMARY KEY,
        school TEXT,
        mascot TEXT,
        abbreviation TEXT,
        alt_name1 TEXT,
        alt_name2 TEXT,
        alt_name3 TEXT,
        color TEXT,
        alt_color TEXT,
        twitter TEXT,
        logo TEXT,
        conference_id INTEGER REFERENCES cfb.conferences(id),
        home_venue_id INTEGER REFERENCES cfb.venues(id)
    );
"""

create_games_table_stmt = """
    CREATE TABLE IF NOT EXISTS cfb.games (
        id BIGINT PRIMARY KEY,
        season INTEGER,
        week INTEGER,
        season_type TEXT,
        start_date TIMESTAMPTZ,
        neutral_site BOOLEAN,
        conference_game BOOLEAN,
        venue_id INTEGER REFERENCES cfb.venues(id),
        home_id INTEGER REFERENCES cfb.teams(id),
        away_id INTEGER REFERENCES cfb.teams(id),
        home_qtr1_score INTEGER,
        home_qtr2_score INTEGER,
        home_qtr3_score INTEGER,
        home_qtr4_score INTEGER,
        home_final_score INTEGER,
        away_qtr1_score INTEGER,
        away_qtr2_score INTEGER,
        away_qtr3_score INTEGER,
        away_qtr4_score INTEGER,
        away_final_score INTEGER,
        winner_id INTEGER REFERENCES cfb.teams(id),
        loser_id INTEGER REFERENCES cfb.teams(id),
        excitement_index DOUBLE PRECISION
    );
"""

create_stats_table_stmt = """
    CREATE TABLE IF NOT EXISTS cfb.stats (
        team_id INTEGER REFERENCES cfb.teams(id),
        season INTEGER,
        -- Basic Stats (from API)
        interception_yards INTEGER,
        turnovers INTEGER,
        punt_return_yards INTEGER,
        fumbles_lost INTEGER,
        possession_time INTEGER,
        fumbles_recovered INTEGER,
        rushing_tds INTEGER,
        kick_return_yards INTEGER,
        pass_attempts INTEGER,
        interception_tds INTEGER,
        fourth_down_conversions INTEGER,
        punt_returns INTEGER,
        interceptions INTEGER,
        third_down_conversions INTEGER,
        passing_tds INTEGER,
        rushing_yards INTEGER,
        penalty_yards INTEGER,
        kick_return_tds INTEGER,
        fourth_downs INTEGER,
        penalties INTEGER,
        third_downs INTEGER,
        kick_returns INTEGER,
        total_yards INTEGER,
        passes_intercepted INTEGER,
        sacks INTEGER,
        punt_return_tds INTEGER,
        tackles_for_loss INTEGER,
        pass_completions INTEGER,
        first_downs INTEGER,
        rushing_attempts INTEGER,
        net_passing_yards INTEGER,

        -- Advanced Stats (from API) o = offensive, d = defensive
        o_plays INTEGER,
        o_drives INTEGER,
        o_ppa DOUBLE PRECISION,
        o_total_ppa DOUBLE PRECISION,
        o_explosiveness DOUBLE PRECISION,
        o_power_success DOUBLE PRECISION,
        o_line_yards DOUBLE PRECISION,
        o_line_yards_total INTEGER,
        o_second_level_yards DOUBLE PRECISION,
        o_second_level_yards_total INTEGER,
        o_open_field_yards DOUBLE PRECISION,
        o_open_field_yards_total INTEGER,
        o_avg_field_position_start REAL,
        o_avg_field_position_predicted_points REAL,
        o_points_per_opportunity DOUBLE PRECISION,
        o_total_opportunities INTEGER,
        d_plays INTEGER,
        d_drives INTEGER,
        d_ppa DOUBLE PRECISION,
        d_total_ppa DOUBLE PRECISION,
        d_explosiveness DOUBLE PRECISION,
        d_power_success DOUBLE PRECISION,
        d_line_yards DOUBLE PRECISION,
        d_line_yards_total INTEGER,
        d_second_level_yards DOUBLE PRECISION,
        d_second_level_yards_total INTEGER,
        d_open_field_yards DOUBLE PRECISION,
        d_open_field_yards_total INTEGER,
        d_avg_field_position_start REAL,
        d_avg_field_position_predicted_points REAL,
        d_points_per_opportunity DOUBLE PRECISION,
        d_total_opportunities INTEGER,
        PRIMARY KEY (team_id, season)
    );
"""

#endregion

try:
     # Set up db connection
    conn, cur = create_pg_conn()

    # Create cfb schema
    execute_stmt(cur, conn, create_cfb_schema_stmt)

    # Create cfb.venues table
    execute_stmt(cur, conn, create_venues_table_stmt)

    # Create cfb.conferences table
    execute_stmt(cur, conn, create_conferences_table_stmt)

    # Create cfb.teams table
    execute_stmt(cur, conn, create_teams_table_stmt)

    # Create cfb.games table
    execute_stmt(cur, conn, create_games_table_stmt)

    # Create cfb.stats table
    execute_stmt(cur, conn, create_stats_table_stmt)

    # Close db connection
    cur.close()
    conn.close()

    print("Created all tables")

except Exception as e:
    print(f"{e}")