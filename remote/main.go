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
	log.Println(password)

	l, err := ldap.DialURL("ldap://172.16.10.101")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	/*

		/usr/bin/ldapsearch \
		    -H "ldap://172.16.10.101" \
		    -l "8" \
		    -x -w "${pass}" \
		    -D "uid=${user},cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com" \
		    -b "cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com" \
		    "uid=${user}"
	*/

	ldapUsername := fmt.Sprintf("uid=%s,cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com", username)
	err = l.Bind(ldapUsername, password)
	if err != nil {
		log.Fatal(err)
	}

	sr, err := l.Search(&ldap.SearchRequest{
		BaseDN: "cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com",
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
