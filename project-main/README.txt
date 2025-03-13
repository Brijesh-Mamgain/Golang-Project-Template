
## to Set the environment variable for the project-main
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

// run the application
go run main.go