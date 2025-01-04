import psycopg2
from dotenv import dotenv_values

# Load environment variables
config = dotenv_values("scripts/.env")

conn_params = {
    "dbname": config["CFB_DB_NAME"],
    "user": config["CFB_DB_USER"],
    "password": config["CFB_DB_PASSWORD"],
    "host": config["CFB_DB_HOST"],
    "port": int(config["CFB_DB_PORT"])
}
#region SQL Statements 

drop_cfb_schema_stmt = "DROP SCHEMA IF EXISTS cfb;"

drop_venues_table_stmt = "DROP TABLE IF EXISTS cfb.venues;"

drop_conferences_table_stmt = "DROP TABLE IF EXISTS cfb.conferences;"

drop_teams_table_stmt = "DROP TABLE IF EXISTS cfb.teams;"

drop_games_table_stmt = "DROP TABLE IF EXISTS cfb.games;"

drop_stats_table_stmt = "DROP TABLE IF EXISTS cfb.stats;"

#endregion

def execute_stmt(cur, conn, stmt):
    # Both of these throw exceptions, let main except catch them
    cur.execute(stmt)
    conn.commit()

try:
     # Set up db connection
    conn = psycopg2.connect(**conn_params)
    cur = conn.cursor()

    # Drop cfb.stats table
    execute_stmt(cur, conn, drop_stats_table_stmt)

    # Drop cfb.games table
    execute_stmt(cur, conn, drop_games_table_stmt)

    # Drop cfb.teams table
    execute_stmt(cur, conn, drop_teams_table_stmt)

    # Drop cfb.conferences table
    execute_stmt(cur, conn, drop_conferences_table_stmt)

    # Drop cfb.venues table
    execute_stmt(cur, conn, drop_venues_table_stmt)

    # Drop cfb schema
    execute_stmt(cur, conn, drop_cfb_schema_stmt)

    # Close db connection
    cur.close()
    conn.close()

    print("Dropped all tables")

except Exception as e:
    print(f"{e}")