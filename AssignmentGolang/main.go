package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct{
	id int   `json:"id"`
	subject string  `json:"subject"`
	marks int `json:"marks"`
}
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info1 := Info{
		id : 1,
		subject : "Maths",
		marks : 95,
	}
	
	requestType := r.Method
	fmt.Println(requestType)
	if requestType=="GET"{
		err := json.NewEncoder(w).Encode(info1)
		if err != nil {
			fmt.Println("Error in Encoding")
			
		}
		fmt.Println(info1)
		fmt.Fprint(w)

	}else if requestType=="POST"{
		info2 := Info{}
		err := json.NewDecoder(r.Body).Decode(&info2)
		if err != nil {
			fmt.Println("Error in Decoding")
			
		}
		fmt.Println(info2)
		fmt.Fprint(w,info2)

	}else{
		fmt.Fprint(w, "Not Defined")
	}
	
}
func main() {
	fmt.Println("starting server")
	http.HandleFunc("/", ExampleHandler)
	http.ListenAndServe("localhost:8080", nil)
}




