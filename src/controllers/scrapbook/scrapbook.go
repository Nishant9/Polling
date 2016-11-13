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
        if value.Name=="Vote" {
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
    User.New_Pass=r.FormValue("New_Password")

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
   

    correct_pass:=[]string{}
    if err := SC.Sqldb.QueryRow("SELECT passwords FROM authdb WHERE username = \"" +username+"\"").Scan(correct_pass); (err != nil){
        panic(err.Error()) 
    } 

    stmt, err := SC.Sqldb.Prepare("UPDATE authdb set passwords = \""+strings.Join(correct_pass,"@")+User.New_Pass + "\"" + " WHERE username =\" "+username + "\"")
    if err != nil {
        panic(err.Error()) 
    }  
    if guard==1 {
    }
    if guard==0 {
        stmt, err = SC.Sqldb.Prepare("UPDATE authdb set passwords = \""+User.New_Pass+strings.Join(correct_pass,"@")+"\"" + " WHERE username =\" "+username + "\"")
        if err != nil {
            panic(err.Error()) 
        }  
    }
    _, err2 := stmt.Exec()
    if err2 != nil {
        panic(err.Error()) 
    }
    http.Redirect(w,r,"/note",302)
    //update as per guard
/*    stmt, err := SC.Sqldb.Prepare("INSERT INTO Notes (username,category,Friend,Text) VALUES (\""+username+"\""+","+"\""+category+"\""+","+"\""+User.For+"\""+","+"\""+User.Text+"\""+") ON DUPLICATE KEY UPDATE Text=\"" + User.Text +"\"")
    if err != nil {
        panic(err.Error()) 
    }  
    //fmt.Println("err1",stmt) 
    //fmt.Println("err2",g) 

*/
    return



}
func Note(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
        t, _ := template.ParseFiles(SC.Base_Path+"src/views/scrapbook.html")
        t.Execute(w, nil )
    }

}