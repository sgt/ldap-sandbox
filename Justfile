default:
    just -l

# Golang server
server:
    go run server/main.go

# Golang client
client:
    go run client/main.go

# glauth server
glauth:
    glauth -c glauth.cfg
