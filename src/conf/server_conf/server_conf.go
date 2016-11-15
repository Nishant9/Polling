package server_conf

import (
	"database/sql"
	"gopkg.in/redis.v2"
)

const (
	AUTH_SERVER      string = "ldap.cse.iitk.ac.in"
	AUTH_BASE_DN     string = "ou=cse,o=iitk,dc=ac,dc=in"
	AUTH_SERVER_PORT int    = 389
	FAST_SERVER      string = "localhost:6379"
	SQL_SERVER       string = "root:root@tcp(localhost:3306)/poll_1"
	LOGIN_SERVER     string = "local"
	//LOGIN_SERVER string = "ldap"
)
const (
	Cookiedb  string = "Tokens"
	Base_Path string = "/home/nis/coding/cs455/Polling/"
)
const (
	Cookie_Length           int    = 128
	Cookie_Alphabets        string = "abcdefghijklmnopqrstuvwxyz0123456789"
	Length_Cookie_Alphabets int    = 36
)

// server links
var Redisdb *redis.Client
var Sqldb *sql.DB

var SqlErrNoRows error = sql.ErrNoRows
