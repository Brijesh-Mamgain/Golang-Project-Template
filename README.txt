
// go mod init project-common  
//go install go.uber.org/zap@latest 

 go mod init project-main-service

Envirobment variable
SERVER_ADDRESS = localhost
SERVER_PORT = 8000
DB_USER =postgres
DB_PASSWD =
DB_ADDR =localhost
DB_PORT=5432 postgres
DB_NAME =project-main

go run main.go

set SERVER_ADDRESS=localhost
set SERVER_PORT=8000
set DB_USER=postgres
set DB_PASSWD=postgres
set DB_ADDR=localhost
set DB_PORT=5432
set DB_NAME=project-main

$env:SERVER_ADDRESS="localhost"
$env:SERVER_PORT="8000"
$env:DB_USER="postgres"
$env:DB_PASSWD="postgres"
$env:DB_ADDR="localhost"
$env:DB_PORT="5432"
$env:DB_NAME="project-main"