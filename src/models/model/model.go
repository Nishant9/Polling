package model

import (
	"fmt"
 	EC "../../conf/election_conf" 
 	ScC "../../conf/scrapbook_conf" 
 	SC "../../conf/server_conf" 
 	"sort"
 	"time"
/* 	"html"
 	"net/url"*/
 	"math/rand"
)


type Candidate_Profile struct{
	Name string
	ID int
	Homepage string
	Image_Link string
}

type Electorate_Profile struct{
	Cookie string
	Votes []string
}




type Pass_Profile struct{
	Cookie string
	New_Pass string
}

type Electorate_Login struct{
	Username string
	Password string 
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = SC.Cookie_Alphabets[rand.Intn(SC.Length_Cookie_Alphabets)]
	}
	return string(result)
}

func (raw *Electorate_Login) Validate() (string,string){
	for Category,User_List := range EC.List {
		index := sort.SearchStrings(User_List, raw.Username)
		if index < len(User_List) && User_List[index] == raw.Username {
			return Category,""
		}
		//fmt.Println("In model",raw.Username,Category)
	}
	return "","Electorate Not Registered"
}

func Check_logged_in(cookie string) int {
	//fmt.Println("Just checking ",SC.Redisdb.SIsMember(SC.Cookiedb,cookie))
	if(SC.Redisdb.SIsMember(SC.Cookiedb,cookie).Val()){
		return 0
	}
	if(SC.Redisdb.SIsMember(SC.Cookiedb,cookie+"decoy").Val()) {
		return 1
	}
	return 2
}

func (raw *Electorate_Profile) Validate() bool {
	tmpmap := make(map[string]bool)
	for i := 0; i < EC.Number_of_votes; i++ {
    	tmpmap[raw.Votes[i]] = true
	}
	if len(tmpmap) != EC.Number_of_votes {
		return true
	}
	for key,_ := range tmpmap{
		if EC.Candidates[key] == "" {
			return true
		}
	}
	return false
}

func (raw *Pass_Profile) Validate() bool {
	if len(raw.New_Pass)>=EC.Pass_Length && ScC.Pass_allowed.MatchString(raw.New_Pass) {
			fmt.Println("Here3")
			return false
		}
	return true
}

//SAdd to allow only one session per user // there is no expiry for the cookie
// @ as seperator is safe as key can't have it .. checked already in checking from voter list
func Bake(key string,decoy string) string{
	cookie := key + "@" + RandomString(SC.Cookie_Length)
	SC.Redisdb.SAdd(SC.Cookiedb,cookie+decoy)
	return cookie
}

func Burn(cookie string,decoy string) {
	SC.Redisdb.SRem(SC.Cookiedb,cookie+decoy)
}