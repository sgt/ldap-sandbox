package main

import (
	// "crypto/tls"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func main() {
	// The username and password we want to check
	// username := "foo"
	// password := "meow"

	l, err := ldap.DialURL("ldap://localhost:10389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Reconnect with TLS
	// err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// First bind with a read only user
	// err = l.Bind(username, password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	l.err = l.UnauthenticatedBind("readonly")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("user authenticated")

}
