#!/bin/bash
set -x

unset PGDATABASE

pghost="${1-localhost}"

if [ ! -z $1 ]; then
    pghost=$1
elif [ ! -z $POSTGRES_HOST ]; then
    pghost=$POSTGRES_HOST
else
    pghost="localhost"
fi

pgusername="${POSTGRES_USER-`whoami`}"
pgpassword="$POSTGRES_PASSWORD"

psqlopts="-h $pghost -U $pgusername -v ON_ERROR_STOP=on"


echo "Creating databases..."
PGPASSWORD=$pgpassword psql $psqlopts -f create.sql 1> /dev/null
psql_exit_status=$?
if [ $psql_exit_status != 0 ]; then
  echo "critical psql failure, check sudo apt install postgresql-client" 1>&2
  exit $psql_exit_status
fi
echo "Done"