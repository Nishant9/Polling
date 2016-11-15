port=8080
host='localhost'

admin :
	python adminroot/manage.py runserver 0.0.0.0:8194

configure :
	cp ./election_conf.go src/conf/election_conf/
	cp ./server_conf.go src/conf/server_conf/
	if [ -d "./photos" ]; then cp -r ./photos src/views/; fi
	mysql -h $(host) -u 'root' -p < sql_setup.sql

run :
	cd src/controllers && go run main.go $(port)

result :
	mysql -h $(host) -u 'root' -p < result.sql
