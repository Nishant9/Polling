package poll

import (
 "net/http"
 SC "../../conf/server_conf" 
 EC "../../conf/election_conf" 
 "github.com/julienschmidt/httprouter"
 "../../models/model"
 "strings"
 "strconv"
 "html/template"
 //"fmt"

)






func Vote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var cookie string
    cookies := r.Cookies()
    for _,value:= range cookies{
        if value.Name=="Vote" {
            cookie = value.Value
            break;
        }
    }
    guard := model.Check_logged_in(cookie)
    if guard==2 {
        http.Redirect(w,r,"/",302)
        return
    }

	User := model.Electorate_Profile{}
    User.Cookie = cookie
    for i := 0; i < EC.Number_of_votes; i++ {
        User.Votes=append(User.Votes,r.FormValue(strconv.Itoa(i+1)))
    }

    guard2 := User.Validate()
    if guard2 {
        http.Redirect(w,r,"/ballot",302)
        return
    	// not logged in redirect to auth
    }
   // fmt.Println("Validated Votes")
    s := strings.Split(User.Cookie, "@")
    username := s[0]
    //category := s[1]
   
   for i := 0; i < EC.Number_of_votes; i++ {
    	 
        stmt, err := SC.Sqldb.Prepare("UPDATE ballot set vote_"+ strconv.Itoa(i) +" = \""+User.Votes[i]+"\"" + " WHERE username =\" "+username + "\"")
    	if err != nil {
            panic(err.Error()) 
        }  
        //fmt.Println("err1",stmt) 

        if guard==0{
            _, err2 := stmt.Exec()
            if err2 != nil {
                panic(err.Error()) 
            }
        }
        //fmt.Println("err2",g) 
   }

    http.Redirect(w,r,"/paper",302)

    return


}
func Paper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var cookie string
    cookies := r.Cookies()
    for _,value:= range cookies{
        if value.Name=="Vote" {
            cookie = value.Value
            break;
        }
    }
    guard := model.Check_logged_in(cookie)
    if guard==2 {
        http.Redirect(w,r,"/",302)
        return
    }else{
        // fmt.Println("CurreAAAAAAAAAAAAAAAAaa",current_votes)
        t, _ := template.ParseFiles(SC.Base_Path+"src/views/poll.html")
        t.Execute(w ,nil)
    }

}