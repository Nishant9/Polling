package server_conf

import (
    "gopkg.in/redis.v2"
	"database/sql" 
) 

const (
	AUTH_SERVER string = "ldap.cse.iitk.ac.in"
	AUTH_BASE_DN string ="ou=cse,o=iitk,dc=ac,dc=in"
	AUTH_SERVER_PORT int=389
	FAST_SERVER string = "localhost:6379"
	SQL_SERVER string = "root:password@tcp(localhost:3306)/aca"
)
const (
	Cookiedb string = "Tokens"
	Base_Path string = "/media/harpreet/winstuff/ACA/Best_Faculty_Voting/"
)
const (
	Cookie_Length int = 128
	Cookie_Alphabets string = "abcdefghijklmnopqrstuvwxyz0123456789"
	Length_Cookie_Alphabets int = 36
)

// server links
var Redisdb *redis.Client
var Sqldb *sql.DB

var  SqlErrNoRows error = sql.ErrNoRows 