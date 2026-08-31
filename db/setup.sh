#!/bin/bash

psql -U postgres -f /db_scripts/create_tables.sql
psql -U postgres -f /db_scripts/create_tables.sql
psql -U postgres -f /dumps/test.sql 
