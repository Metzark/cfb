import psycopg2
from dotenv import dotenv_values

# Create pg connection and cursor (using environment variables)
def create_pg_conn():
    # Load environment variables
    config = dotenv_values("scripts/.env")

    conn_params = {
        "dbname": config["CFB_DB_NAME"],
        "user": config["CFB_DB_USER"],
        "password": config["CFB_DB_PASSWORD"],
        "host": config["CFB_DB_HOST"],
        "port": int(config["CFB_DB_PORT"])
    }

    conn = psycopg2.connect(**conn_params)
    
    cur = conn.cursor()

    return (conn, cur)