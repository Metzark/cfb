from psycopg2.extras import execute_values

# Run a SQL statement with a given connection and cursor (EXCEPTIONS RAISED BY EXECUTE OR COMMIT ARE NOT CAUGHT HERE)
def execute_stmt(cur, conn, stmt, values = None, batch_size = 250):
    # If there are not values (Non insertion)
    if values == None:
        cur.execute(stmt)
        conn.commit()
        return
    
    # Only insert batch size at a time
    for i in range(0, len(values), batch_size):
        batch = values[i:i + batch_size]
        execute_values(cur, stmt, batch)
    
    conn.commit()

    


    