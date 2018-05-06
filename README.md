## OpenList-Backend
This repository will contain the Go implementation for the OpenList backend. The intention of this project is to create a fully dockerized open source service to manage tasks and lists.

The end goal for this project is to create a repo which anybody can fork or clone and put up on any service which supports docker ie Now-CLI, AWS, Vultr, ect. and have their own local list manager. The UI for this lives [here](https://github.com/NickLewanowicz/OpenList-app) as a react native application and there will be a desktop web app made in react but there is no projected timeline for completion for those.

See `Projects` tab for activity and goals with this project!

### Installation
Please Install Go:

https://golang.org/doc/install

After I complete the dockerfile this hopefully wont be necessary
For mac/unix install run the following in bash:  
- `git clone https://github.com/NickLewanowicz/OpenList-backend.git`
- `cd OpenList-backend`
- `go build`
- `OpenList-backend`

The server is now running :)



### Configuration 
**TBD**


### Resources Used
- [Restfull Go API Article] (https://thenewstack.io/make-a-restful-json-api-go/)
- [Go Documentation] (https://golang.org/doc/)
- [MySQL driver] (https://github.com/go-sql-driver/mysql#examples)
- [MySQL with Go] (https://tutorialedge.net/golang/golang-mysql-tutorial/)
