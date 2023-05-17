default:
    just -l

# Golang server
server:
    go run server/main.go

# Golang client
client *ARGS:
    go run client/main.go {{ARGS}}

# glauth server
glauth:
    glauth -c glauth.cfg
