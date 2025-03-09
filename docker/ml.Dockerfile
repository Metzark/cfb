# Use Python 3.13 as base image
FROM python:3.13-bookworm

# Set the working directory inside the container
WORKDIR /usr/app

# Copy requirements file and install dependencies
COPY ../ml/requirements.txt .

COPY ../.env .

RUN pip install --no-cache-dir -r requirements.txt

# Keep the container running
CMD ["sleep", "infinity"]
