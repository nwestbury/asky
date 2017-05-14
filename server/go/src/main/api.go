package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
//	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"github.com/asaskevich/govalidator"
)

type Response struct {
    Message string `json:"msg"`
	Success bool   `json:"success"`
}


var oauthConf = &oauth2.Config{
    ClientID:     os.Getenv("ENV_FB_CLIENT_ID"),
    ClientSecret: os.Getenv("ENV_FB_CLIENT_SECRET"),
    RedirectURL:  os.Getenv("ENV_FB_REDIRECT_URL"),
    Scopes:       []string{"email", "public_profile"},
    Endpoint:     facebook.Endpoint,
}

var oauthStateString = "thisshouldberandom"

func writeReply(w http.ResponseWriter, resp *Response){
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(*resp)

	if err != nil {
		log.Print(err)
	}
	
	w.Write(bytes)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	resp := Response{Message: "OK!", Success: true}
	writeReply(w, &resp)
})

var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	email := query.Get("email")
	token := query.Get("token")

	resp := Response{Message: "Failed", Success: false}

	if len(email) == 0 || len(token) == 0 {
		resp.Message = "Missing email or token"
		writeReply(w, &resp)
		return
	}

	if !govalidator.IsEmail(email) {
		resp.Message = "Not a valid email"
		writeReply(w, &resp)
		return
	}
	
	stmt, err := db.Prepare(`INSERT INTO main_schema.user (email, token, creation_time) 
                             VALUES ($1, $2, NOW())`)

	if err != nil {
		log.Print(err)
		resp.Message = "Internal fail"
		writeReply(w, &resp)
		return
	}
	
	_, err = stmt.Exec(email, token)

	if err != nil {
		log.Print(err)
		writeReply(w, &resp)
		return
	}

	resp.Success = true
	resp.Message = "OK!"
	writeReply(w, &resp)
}) 

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Login"))
	// http://stackoverflow.com/questions/27368973/golang-facebook-authentication-using-golang-org-x-oauth2
})
