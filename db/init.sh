#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    \connect $POSTGRES_DB
    BEGIN;
        CREATE TABLE IF NOT EXISTS  pairs (
            id int4 NOT NULL,
            fsym varchar(10) NOT NULL,
            tsym varchar(10) NOT NULL,
            updated_at timestamp NOT NULL DEFAULT now(),
            created_at timestamp NOT NULL DEFAULT now(),
            raw text NOT NULL,
            display text NOT NULL,
            CONSTRAINT pairs_pk PRIMARY KEY (id)
        );
    COMMIT;
EOSQL
