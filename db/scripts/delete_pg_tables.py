from functions.create_pg_conn import create_pg_conn
from functions.execute_stmt import execute_stmt

#region SQL Statements 

drop_cfb_schema_stmt = "DROP SCHEMA IF EXISTS cfb;"

drop_venues_table_stmt = "DROP TABLE IF EXISTS cfb.venues;"

drop_conferences_table_stmt = "DROP TABLE IF EXISTS cfb.conferences;"

drop_teams_table_stmt = "DROP TABLE IF EXISTS cfb.teams;"

drop_games_table_stmt = "DROP TABLE IF EXISTS cfb.games;"

drop_stats_table_stmt = "DROP TABLE IF EXISTS cfb.stats;"

#endregion

try:
     # Set up db connection
    conn, cur = create_pg_conn()

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