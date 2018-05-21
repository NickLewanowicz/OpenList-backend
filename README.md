<b>WARNING: </b><i>Progress on the Golang backend for Kickit has been delayed in favour of the Hapi.JS backend [here](http://www.github.com/nicklewanowicz/kickit-backend) as it will reduce barrier to entry for contributors to the project by providing a full javascript stack. The intention is to circle back and build out the golang backend in the future to offer options for the backend of your Kickit install.</i>

## kickit-Backend

This repository will contain the Go implementation for the kickit backend. The intention of this project is to create a fully dockerized open source service to manage users, projects, and tasks. 

The end goal for this project is to create a repo which anybody can fork or clone and put up on any service which supports docker ie Now-CLI, AWS, Vultr, ect. and have their own project/task manager. The UI for this lives [here](https://github.com/NickLewanowicz/OpenList-app) as a react native application and there will be a desktop web app made in react but there is no projected timeline for completion for those.

<b>tldr; Open source Asana/Trello</b>

See `Projects` tab for activity and goals with this project!

### Prerequisites

Go: https://golang.org/doc/install
MySQL: https://dev.mysql.com/downloads/mysql/

### Installation

After I complete the dockerfile this hopefully wont be necessary
For mac/unix install run the following in bash:  
- `go get github.com/NickLewanowicz/kickit-backend`
- `cd kickit-backend`
- `go build`
- `kickit-backend`

The server is now running :)

Ex.
![](https://image.ibb.co/hC9Rtn/carbon_19.png)


### Configuration 
**TBD**


### Resources Used
- [Restfull Go API Article] (https://thenewstack.io/make-a-restful-json-api-go/)
- [Go Documentation] (https://golang.org/doc/)
- [MySQL driver] (https://github.com/go-sql-driver/mysql#examples)
- [MySQL with Go] (https://tutorialedge.net/golang/golang-mysql-tutorial/)
