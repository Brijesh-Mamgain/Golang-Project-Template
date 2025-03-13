project-main : This project has the dependency on the project-auth-service and project-common.
project-common This project reusable component of logger and exception 
               to be used as a common component for other project.
project-auth-service : this project is developed as a microservice in GO 
              to provide the authentication and autherization feature.
Database  : project-main & project-auth-service is confired to use the postgres DB


### Vim commands 
$ vi <filename> — Open or edit a file.
i — Switch to Insert mode.
Esc — Switch to Command mode.
:w — Save and continue editing.
:wq or ZZ — Save and quit/exit vi.
:q! — Quit vi and do not save changes.
yy — Yank (copy) a line of text.
p — Paste a line of yanked text below the current line.