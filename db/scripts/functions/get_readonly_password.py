from dotenv import dotenv_values

# Get the password for pg user 'readonly'
def get_readonly_password():
    # Load environment variables
    config = dotenv_values(".env")

    return config["CFB_DB_READONLY_PASSWORD"]