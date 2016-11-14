package scrapbook

import (
 "net/http"
 SC "../../conf/server_conf" 
 "github.com/julienschmidt/httprouter"
 "../../models/model"
 "html/template"
 //"fmt"
 "strings"

)

func Paper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var cookie string
    cookies := r.Cookies()
    for _,value:= range cookies{
        if value.Name=="IITKvote" {
            cookie = value.Value
            break
        }
    }
    guard := model.Check_logged_in(cookie)
    if guard==2 {
        http.Redirect(w,r,"/",302)
        return
    }

    User := model.Pass_Profile{}
    User.Cookie = cookie
    User.New_Pass=r.FormValue("Password")

    guard2 := User.Validate()
    if guard2 {
        http.Redirect(w,r,"/book",302)
        return
        // not logged in redirect to auth
    }
   // fmt.Println("Validated Votes")
    s := strings.Split(User.Cookie, "@")
    username := s[0]
    //category := s[1]
   

    correct_password:=""
    if err := SC.Sqldb.QueryRow("SELECT passwords FROM authdb WHERE username = \"" +username+"\"").Scan(&correct_password); (err != nil){
        if err != SC.SqlErrNoRows {
            panic(err.Error()) 
        }
    } 
    correct_pass:=strings.Split(correct_password,"@")
    stmt, err := SC.Sqldb.Prepare("UPDATE authdb set passwords = \""+strings.Join(correct_pass,"@")+ "@" +User.New_Pass + "\"" + " WHERE username =\""+username + "\"")
    if err != nil {
        panic(err.Error()) 
    }  
    if guard==0 {
        stmt, err = SC.Sqldb.Prepare("UPDATE authdb set passwords = \""+User.New_Pass+ "@" + strings.Join(correct_pass,"@")+"\"" + " WHERE username =\""+username + "\"")
        if err != nil {
            panic(err.Error()) 
        }  
    }
    _, err2 := stmt.Exec()
    if err2 != nil {
        panic(err.Error()) 
    }
    http.Redirect(w,r,"/note",302)

    return



}
func Note(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var cookie string
    cookies := r.Cookies()
    for _,value:= range cookies{
        if value.Name=="IITKvote" {
            cookie = value.Value
            break;
        }
    }
    guard := model.Check_logged_in(cookie)
    if guard==2 {
        http.Redirect(w,r,"/",302)
        return
    }else{
        t, _ := template.ParseFiles(SC.Base_Path+"src/views/scrapbook.html")
        t.Execute(w, nil )
    }

}