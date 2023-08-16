#!/bin/sh
set -e

echo "download curl"
apk add curl

echo "download migration"
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "add env"
echo DB_SOURCE=$DB_SOURCE >> app.env
echo DB_DRIVER=$DB_DRIVER >> app.env
echo SERVER_ADDRESS=$SERVER_ADDRESS >> app.env
echo ACCESS_TOKEN_PRIVATE_KEY=$ACCESS_TOKEN_PRIVATE_KEY >> app.env
echo ACCESS_TOKEN_PUBLIC_KEY=$ACCESS_TOKEN_PUBLIC_KEY >> app.env
echo ACCESS_TOKEN_EXPIRED_IN=$ACCESS_TOKEN_EXPIRED_IN >> app.env
echo ACCESS_TOKEN_MAXAGE=$ACCESS_TOKEN_MAXAGE >> app.env
echo REFRESH_TOKEN_PRIVATE_KEY=$REFRESH_TOKEN_PRIVATE_KEY >> app.env
echo REFRESH_TOKEN_PUBLIC_KEY=$REFRESH_TOKEN_PUBLIC_KEY >> app.env
echo DB_DREFRESH_TOKEN_EXPIRED_INRIVER=$REFRESH_TOKEN_EXPIRED_IN >> app.env
echo REFRESH_TOKEN_MAXAGE=$REFRESH_TOKEN_MAXAGE >> app.env

echo "start the app"
exec "$@"