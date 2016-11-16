#!/usr/bin/bash
if [ $1 -gt 0 ]
then
    sed -i '6d' sql_setup.sql
    ballot="create table ballot (username VARCHAR(50) PRIMARY KEY"
    for i in $(seq 1 $1); do 
    		temp=", vote_$((i - 1)) VARCHAR(40)"
   			ballot=$ballot$temp
		done
	temp=");"
    ballot=$ballot$temp
    sed  -i "5 a $ballot" sql_setup.sql
    sed -i -E "s/const Number\_of\_votes int \=[0-9]+/const Number\_of\_votes int \=$1/g" election_conf.go
fi
