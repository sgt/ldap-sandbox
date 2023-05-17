package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("need 2 arguments: user and password")
	}

	username := os.Args[1]
	password := os.Args[2]

	l, err := ldap.DialURL("ldap://172.16.10.101")
	// l, err := ldap.DialURL("ldap://127.0.0.1:10389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	baseDN := "cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com"

	ldapUsername := fmt.Sprintf("uid=%s,%s", username, baseDN)
	err = l.Bind(ldapUsername, password)
	if err != nil {
		log.Fatal(err)
	}

	sr, err := l.Search(&ldap.SearchRequest{
		BaseDN: baseDN,
		Filter: "(uid=sgt)",
		// Filter:     fmt.Sprintf("(&(uid=%s))", username),
		Attributes: []string{"telephoneNumber"},
		Scope:      ldap.ScopeWholeSubtree,
		TimeLimit:  8,
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("no single entry found")
	}

	entry := sr.Entries[0]
	if len(entry.Attributes) != 1 || entry.Attributes[0].Name != "telephoneNumber" {
		log.Fatal("no telephone number entries")
	}

	phoneNumbers := entry.Attributes[0].Values
	if len(phoneNumbers) == 0 {
		log.Fatal("no telephone numbers")
	}

	log.Println(phoneNumbers[0])

}
