#!/usr/bin/bash
if [ $1 -gt 1 ]
then
    sed -i '6d' sql_setup.sql
    for i in $(seq 1 $1); do sed  -i "$((i + 4)) a create table ballot (username VARCHAR(50) PRIMARY KEY, vote_$((i - 1)) VARCHAR(40));  " sql_setup.sql ; done
fi
