admin :
	python adminroot/manage.py runserver 0.0.0.0:8194

configure :
	cp ./election_conf.go src/conf/election_conf/
	mysql -h 'localhost' -u 'root' -p < sql_setup.sql

run :
	cd src/controllers && go run main.go 8080

result :
	mysql -h 'localhost' -u 'root' -p < result.sql
	
