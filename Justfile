default:
    just -l

# Golang server
server:
    go run server/main.go

# Golang client
client:
    go run client/main.go

# test against remote server
remote *ARGS:
    go run remote/main.go {{ARGS}}

# glauth server
glauth:
    glauth -c glauth.cfg
