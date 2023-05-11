package main

import (
	// "crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func main() {
	// The username and password we want to check
	username := "someuser"
	password := "userpassword"

	bindUsername := "readonly"
	bindPassword := "password"

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
	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
	}

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", ldap.EscapeFilter(username)),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
		log.Fatal(err)
	}

	// Rebind as the read only user for any further queries
	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
	}
}
