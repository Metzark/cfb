# cfb

CFB.

# Postgres DB

### Running Postgres Locally

Start up docker container running postgres:

```
docker-compose -f docker/docker-compose.yaml up -d
```

Create necessary tables:

```
python3 db/scripts/create_pg_tables.py
```

Insert base data (venues, conferences, teams, games):

```
python3 db/scripts/insert_pg_base_from_files.py
```

Insert stats data:

```
python3 db/scripts/insert_pg_stats_from_files.py
```

If tables need to be deleted (optional):

```
python3 db/scripts/delete_pg_tables.py
```

Spin down docker container running postgres:

```
docker-compose -f docker/docker-compose.yaml down
```
