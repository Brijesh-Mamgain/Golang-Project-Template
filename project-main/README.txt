
project-main : This project has the dependency on the project-auth-service and project-common.
project-common This project reusable component of logger and exception 
               to be used as a common component for other project.
project-auth-service : this project is developed as a microservice in GO 
              to provide the authentication and autherization feature.
Database  : project-main & project-auth-service is confired to use the postgres DB


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