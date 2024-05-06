#!/bin/bash
SH_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR=$(dirname $SH_DIR)

brew install flyway

cd REPO_DIR

echo "" > flyway.conf
echo "flyway.url=jdbc:postgresql://${POSTGRES_HOST}:$POSTGRES_PORT/$POSTGRES_DB" >> flyway.conf
echo "flyway.user=$POSTGRES_USER" >> flyway.conf
echo "flyway.password=$POSTGRES_PASSWORD" >> flyway.conf
echo "" >> flyway.conf 
flyway migrate -locations=filesystem:$REPO_DIR/sql/ddl/
