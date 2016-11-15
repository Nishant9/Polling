package dial

import (
	SC "../../conf/server_conf"
    "gopkg.in/redis.v2"
	"database/sql" 
    //"fmt"
    _ "github.com/go-sql-driver/mysql"
) 

func Setup_sql() {
	var err error
    SC.Sqldb, err = sql.Open("mysql", SC.SQL_SERVER)
    if err != nil {
        panic(err.Error())  
    }
}


func Close_sql (){
	if err := SC.Sqldb.Close();err!=nil{
		panic(err.Error())
	}
}

//Function to dial connection to Redis Server
func Setup_redis() {

    SC.Redisdb = redis.NewTCPClient(&redis.Options{
        Addr:     SC.FAST_SERVER,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}

func Close_redis() {
	if err := SC.Redisdb.Close();err!=nil{
		panic(err.Error())
	}
}

