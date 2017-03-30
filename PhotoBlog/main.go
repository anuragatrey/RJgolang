package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var tp1 *template.Template
var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	tp1, _ = template.ParseGlob("assets/templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("assets/images/", http.StripPrefix("assets/images/", http.FileServer(http.Dir("./assets/images"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}

func index(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	//Authenticate
	if session.Values["loggedin"] == "false" || session.Values["loggedin"] == nil {
		http.Redirect(res, req, "/login", 302)
		return
	}
	//upload Photo
	src, hdr, err := req.FormFile("data")
	if req.Method == "POST" && err == nil {
		uploadPhoto(src, hdr, session)
	}
	//Save session
	session.Save(req, res)
	//get Photos
	data := getPhotos(session)
	//Execute template
	tp1.ExecuteTemplate(res, "index.html", data)
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	session.Values["loggedin"] = "flase"
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}

func login(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if req.Method == "POST" && req.FormValue("userName") != "" && req.FormValue("password") == "secret" {
		session.Values["loggedin"] = "true"
		session.Save(req, res)
		http.Redirect(res, req, "/", 302)
	}
	tp1.ExecuteTemplate(res, "login.html", nil)
}

func uploadPhoto(src multipart.File, hdr *multipart.FileHeader, session *sessions.Session) {
	defer src.Close()
	fname := getSha(src) + ".jpg"
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "assets", "images", fname)
	dst, _ := os.Create(path)
	defer dst.Close()
	src.Seek(0, 0)
	io.Copy(dst, src)
	addPhoto(fname, session)
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func addPhoto(fname string, session *sessions.Session) {
	data := getPhotos(session)
	data = append(data, fname)
	bs, _ := json.Marshal(data)
	session.Values["data"] = string(bs)
}

func getPhotos(session *sessions.Session) []string {
	var data []string
	jsonData := session.Values["data"]
	if jsonData != nil {
		json.Unmarshal([]byte(jsonData.(string)), &data)
	}
	return data
}
