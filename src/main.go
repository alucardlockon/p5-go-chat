package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
	"database/sql"
)

func index(w http.ResponseWriter, r *http.Request){
	/*
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	//fmt.Fprint(w,"Hello p5")
	t,_:=template.ParseFiles("web/main.html")
	t.Execute(w,nil)
	sql.Open("mysql","root:1234@/test?charset=utf8")
	*/
}

func main(){
	http.HandleFunc("/",index)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("ListenAndServe:",err)
	}
}
