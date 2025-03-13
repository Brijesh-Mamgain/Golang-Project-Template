
// go mod init project-common  
//go install go.uber.org/zap@latest 

 go mod init project-auth-service

Envirobment variable

go run main.go

set SERVER_ADDRESS=localhost
set SERVER_PORT=8082
set DB_USER=postgres
set DB_PASSWD=postgres
set DB_ADDR=localhost
set DB_PORT=5432
set DB_NAME=project-main

$env:SERVER_ADDRESS="localhost"
$env:SERVER_PORT="8082"
$env:DB_USER="postgres"
$env:DB_PASSWD="postgres"
$env:DB_ADDR="localhost"
$env:DB_PORT="5432"
$env:DB_NAME="project-main"