#!/usr/bin/bash

if [ $1 -gt 0 ]
then
    sed -i '6d' sql_setup.sql
    ballot="create table ballot (username VARCHAR(50) PRIMARY KEY"
    poll_string="<br><br><label for=\"1\">To</label>	<br/> {{ range \$key, \$value := . }} <input type=\"radio\" name=\"1\" value=\"{{\$key}}\" checked> {{\$key}} <img src=\"/res/photos/{{\$value}}\" width=\"60px\" height=\"60px\"><br> {{ end }}"
    poll_string_final=""
    for i in $(seq 1 $1); do 
    		temp=", vote_$((i - 1)) VARCHAR(40)"
   			ballot=$ballot$temp
   			temp=$poll_string
   			temp=$(echo $temp | sed -e 's/for\=\"[0-9]*\"/for\=\"'$i'\"/g')
   			temp=$(echo $temp | sed -e 's/name\=\"[0-9]*\"/name\=\"'$i'\"/g')
   			poll_string_final=$poll_string_final$temp
		done
	sed -i -e "s%.*radio.*%TEXT_TO_BE_REPLACED%" ./src/views/poll.html	
	sed -i -e "s%TEXT_TO_BE_REPLACED.*%$poll_string_final\n%" ./src/views/poll.html	
	temp=");"
    ballot=$ballot$temp
    sed  -i "5 a $ballot" sql_setup.sql
    sed -i -E "s/const Number\_of\_votes int \=[0-9]+/const Number\_of\_votes int \=$1/g" election_conf.go
fi
