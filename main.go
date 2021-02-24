package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
	Last  string
}

func basicAuth() {
	// this func might be used to make a string for an Auth header
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("jed:A fun password")))
}

func hashPW(password string) ([]byte, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("error in comparing pw, %w", err)
	}

	return pw, nil
}

func checkPW(password string, hashed []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid Password: %w", err)
	}
	return nil
}

func main() {
	// p1 := person{
	// 	First: "Jed",
	// 	Last:  "Zeins",
	// }
	// p2 := person{
	// 	First: "Claude",
	// 	Last:  "Zeins",
	// }

	// persons := []person{p1, p2}

	// bs, err := json.Marshal(persons)
	// // turns to json string?////

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println(string(bs))

	// persons2 := []person{}

	// // needs pointer of object
	// err = json.Unmarshal(bs, &persons2)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Back into a go data structure", persons2)

	// basicAuth()

	pw := "12345678"

	hash, err := hashPW(pw)

	if err != nil {
		panic(err)
	}

	err = checkPW(pw, hash)

	if err != nil {
		log.Fatalln("not logged in")
	}

	log.Println("Logged in!")

	fmt.Println(err)

	// http.HandleFunc("/quiz1E", encodeHandler)
	// http.HandleFunc("/quiz1D", decodeHandler)
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8081", nil)

}
func foo(w http.ResponseWriter, r *http.Request) {

	p1 := person{
		First: "Jed",
		Last:  "Zeins",
	}

	err := json.NewEncoder(w).Encode(p1)

	if err != nil {
		log.Println("encoding issue, ", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person

	err := json.NewDecoder(r.Body).Decode(&p1)

	if err != nil {
		log.Println("decoding issue, ", err)
	}

	log.Println(p1)
}
