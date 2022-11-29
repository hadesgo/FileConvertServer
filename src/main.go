package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hadesgo/FileConvertServer/src/session"
	_ "github.com/hadesgo/FileConvertServer/src/session/providers/memory"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	// var sessionToken string
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// timestamp := strconv.Itoa(time.Now().Nanosecond())
		// hashWr := md5.New()
		// hashWr.Write([]byte(timestamp))
		// token := fmt.Sprintf("%x", hashWr.Sum(nil))
		// sessionToken = token

		t, _ := template.ParseFiles("webpage/login.html")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
		// t.Execute(w, token)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
		// token := r.Form.Get("token")
		// if token != "" {
		// 	if token != sessionToken {
		// 		fmt.Println("token error: token check failed")
		// 		fmt.Println("token: ", token)
		// 		fmt.Println("sessionToken: ", sessionToken)
		// 	} else {
		// 		fmt.Println("token: ", token)
		// 	}
		// } else {
		// 	fmt.Println("token error: token is null")
		// }
		// fmt.Println("username length:", len(r.Form["username"][0]))
		// fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		// fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		// template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("webpage/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, hander, err := r.FormFile("uploadfile")
		if err != nil {
			log.Fatal("uploadfile error: ", err)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", hander.Header)
		f, err := os.OpenFile("./test/"+hander.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal("OpenFile error: ", err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
