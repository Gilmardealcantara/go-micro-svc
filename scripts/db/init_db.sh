#!/bin/bash

# this script is needed to created db with user since we are setting up postgres
# for the go-temaplate application as an example.
# Feel free to edit/remove it if not needed.
set -e
set -u

function create_user_and_database() {
	local database=$1
	echo "Creating user and database '$database'"
	psql -v --username "$DB_HOST" <<-EOSQL # must use name of container for host, which in this case is postgres
	    CREATE USER $database;
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $database;
		ALTER USER $database CREATEDB;
EOSQL
}

if [ -n "$DB_NAME" ]; then
	echo "Database creation requested: $DB_NAME"
    create_user_and_database $DB_NAME
	echo "Database created"
fi
