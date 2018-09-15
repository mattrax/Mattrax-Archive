# Mattrax [![Go Report Card](https://goreportcard.com/badge/github.com/mattrax/Mattrax)](https://goreportcard.com/report/github.com/mattrax/Mattrax)
Mattrax Is An Open Source Device Management System

### [Project Feature Tracker/Roadmap](https://github.com/mattrax/Mattrax/projects/1)

# Notes For Other Developers
If You Do Go Development And Don't Agree With These Decision Please Create An Issue To Discuss It.

## Project Structure
This Projects File Structure Is Modelled After [Go Standard Project Layout](https://github.com/golang-standards/project-layout)

## Web Frameworks/Routers
It Is Unclear To Me Whether It Is Idiomatic To Use A Web Router (Mux). I Eventually Decided On Using [httprouter](https://github.com/julienschmidt/httprouter) Because It Is Fast And Makes Routing Alot Easier Than The Built In `net/http`.
