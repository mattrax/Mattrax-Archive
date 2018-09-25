# Mattrax [![Go Report Card](https://goreportcard.com/badge/github.com/mattrax/Mattrax)](https://goreportcard.com/report/github.com/mattrax/Mattrax)
Mattrax Is An Open Source Device Management System

### [Project Feature Tracker/Roadmap](https://github.com/mattrax/Mattrax/projects/1)

# This Project Is In Early Development
This Project is Currently Not Ready To Be Used. At The Current Rate A Version With Usable Functionality And Security Will Be Avaiable By The End Of 2018 So Please Check Back Here. This Project Is Planned To Be Installed Via Docker To Make Updates & Management Easier. There Will Also Be Tests Added So That Updates Can Be Safely Push Without Breaking Deployments. If You Are Interested In Helping Out Contact Me [Oscar Beaumont](mailto:oscartbeaumont@gmail.com)

# Below Is Stuff For My Reference

## Notes For Other Developers
If You Do Go Development And Don't Agree With These Decision Please Create An Issue To Discuss It.

### Project Structure
This Projects File Structure Is Modelled After [Go Standard Project Layout](https://github.com/golang-standards/project-layout)

### Web Frameworks/Routers
It Is Unclear To Me Whether It Is Idiomatic To Use A Web Router (Mux). I Eventually Decided On Using [httprouter](https://github.com/julienschmidt/httprouter) Because It Is Fast And Makes Routing Alot Easier Than The Built In `net/http`.



### Objectives/TODO With The Rebuild
* Logging
* HTTP Routing/Handling
* HTTP Middlewear -> Logging, XSS, CORS, Authentication
* Database -> Access In HTTP Handlers
* Postges Triggers

### Database Driver Decision
#### Benchmarks
SQLX:
  Connected To Database In 11.105445ms
  Queried Multple In 2.741305ms
  Queried Single In 1.224454ms

Raw SQL:
  Connected To Database In 22.672µs
  Queried Multple In 25.992115ms
  Queried Single In 1.24597ms

Upper DB:
  Connected To Database In 20.273257ms
  Queried Multple In 11.091147ms
  Queried Single In 1.34252ms
