package main

import "net/http"

func allApis() {
	http.HandleFunc("/get", getAll)
	http.HandleFunc("/post", postAll)
}
