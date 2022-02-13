#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_NAME" <<-EOSQL
    CREATE USER commentify WITH PASSWORD 'secret';
    CREATE DATABASE commentify;
    GRANT ALL PRIVILEGES ON DATABASE commentify TO commentify;
EOSQL