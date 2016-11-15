package auth

import (
 SC "../../conf/server_conf" 
 "../../models/model"
 "github.com/jtblin/go-ldap-client"
 "net/http"
 "github.com/julienschmidt/httprouter"
 "html/template"
 "strings"
 //"fmt"
)


func login_from_server(username string, password string) int {
    
    pass := ""
    if err := SC.Sqldb.QueryRow("SELECT passwords FROM authdb WHERE username = \"" +username+"\"").Scan(pass); (err != nil){
        panic(err.Error()) 
    }  
    correct_pass:=strings.Split(pass,"@")
    //fmt.Println("err1",stmt) 
    if password==correct_pass[0] {
        return 0
    }
    for i := 1; i < len(correct_pass); i++ {
        if password==correct_pass[0] {
            return 1
        }
    }
    return 2
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Raw_User := model.Electorate_Login{}
    Raw_User.Username=r.FormValue("Username")
    Raw_User.Password=r.FormValue("Password")
   // fmt.Println("In login",Raw_User.Username)

    Category,str_err:=Raw_User.Validate()
    if str_err!=""{
        http.Error(w,str_err,401)
        return
    }
    guard := login_from_server(Raw_User.Username,Raw_User.Password)
    if guard==2 {
        http.Error(w,"Wrong Password",401)
        return
    }
//fmt.Println("Logged in from remote server")
    decoy := ""
    if(guard==1) decoy="decoy"
    cookie := model.Bake(Raw_User.Username + "@" + Category,decoy);
    http.SetCookie(w,&http.Cookie{Name:"ACAVote",Value:cookie})
    http.Redirect(w,r,"/paper",302)
  
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    var cookie string
    cookies := r.Cookies()
    for _,value:= range cookies{
        if value.Name=="Vote" {
            cookie = value.Value
            break;
        }
    }
    decoy:=""
    guard := model.Check_logged_in(cookie)
    if(guard==1) decoy:="decoy"
    model.Burn(cookie,decoy);
        http.Redirect(w,r,"/",302)
    
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var cookie string
    cookies := r.Cookies()
    for _,value:= range cookies{
        if value.Name=="ACAVote" {
            cookie = value.Value
            break;
        }
    }
    guard := model.Check_logged_in(cookie)
    if guard==2 {
        t, _ := template.ParseFiles(SC.Base_Path+"src/views/auth.html")
        t.Execute(w, &map[string]string{"Username":""} )
        return
    }else{
        http.Redirect(w,r,"/paper",302)
    }
}