#!/bin/sh
set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "where current file"
pwd

echo "list folder"
ls -l

echo "add env"
echo DB_SOURCE=$DB_SOURCE >> .env
echo DB_DRIVER=$DB_DRIVER >> .env
echo DB_SOURCE=$SERVER_ADDRESS >> .env
echo DB_DRIVER=$ACCESS_TOKEN_PRIVATE_KEY >> .env
echo DB_SOURCE=$ACCESS_TOKEN_PUBLIC_KEY >> .env
echo DB_DRIVER=$ACCESS_TOKEN_EXPIRED_IN >> .env
echo DB_SOURCE=$ACCESS_TOKEN_MAXAGE >> .env
echo DB_DRIVER=$REFRESH_TOKEN_PRIVATE_KEY >> .env
echo DB_DRIVER=$REFRESH_TOKEN_PUBLIC_KEY >> .env
echo DB_DRIVER=$REFRESH_TOKEN_EXPIRED_IN >> .env
echo DB_DRIVER=$REFRESH_TOKEN_MAXAGE >> .env

echo "start the app"
exec "$@"