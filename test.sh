#!/usr/bin/env bash

user=$1
pass=$2

/usr/bin/ldapsearch \
    -H "ldap://172.16.10.101" \
    -l "8" \
    -x -w "${pass}" \
    -D "uid=${user},cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com" \
    -b "cn=users,cn=accounts,dc=ipa,dc=abc-p,dc=com" \
    "uid=${user}"

echo $?
