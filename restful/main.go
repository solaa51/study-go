package main

import (
    "net/http"
    "fmt"
    "github.com/julienschmidt/httprouter"
    "log"
)

/**
get post put delete
put和delete 可能会被拦截 可以采用post+参数(标记) 隐式使用
 */

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "welcome\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "Hello %s\n", ps.ByName("name"))
}

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    uid := ps.ByName("uid")
    fmt.Fprintf(w, "you are get user %s\n", uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    uid := ps.ByName("uid")
    fmt.Fprintf(w, "add user %s\n", uid)
}

func moduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    uid := ps.ByName("uid")
    fmt.Fprintf(w, "modify user %s\n", uid)
}


func deluser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    uid := ps.ByName("uid")
    fmt.Fprintf(w, "delete user %s\n", uid)
}


func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    router.GET("/user/:uid", getuser)
    router.POST("/adduser/:uid", adduser)
    router.PUT("/moduser/:uid", moduser)
    router.DELETE("/deluser/:uid", deluser)

    log.Fatal(http.ListenAndServe(":8080", router))
}