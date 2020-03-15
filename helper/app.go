package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	redisLen, err := GetListLength()

	fmt.Printf("Redis len: %d", redisLen)
	if err != nil {
		fmt.Fprintf(w, "Something is not working :<")
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "Redis list len: %d", redisLen)
	}
	fmt.Println("Endpoint Hit: homePage")
}

// StartWebserver start simple HTTP server
func StartWebserver() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3232", nil))
}
