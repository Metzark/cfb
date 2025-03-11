from psycopg2 import pool
from os import getenv


def create_pg_pool():
    conn_pool = pool.SimpleConnectionPool(
        minconn=1,
        maxconn=10,
        dbname=getenv("CFB_DB_NAME"),
        user=getenv("CFB_DB_USER"),
        password=getenv("CFB_DB_PASSWORD"),
        host=getenv("MLEARNING_CFB_DB_HOST"),
        port=int(getenv("CFB_DB_PORT"))
    )

    return conn_pool