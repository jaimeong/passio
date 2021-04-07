# passio-REST

Develop a REST API service using Golang for managing User entities (CRUD operations) stored in a PostgreSQL database.  The solution must be submitted with full source code, a docker stack/compose file for running the full system, and an integration test suite that will run against the docker stack and prove that the implementation works.  Include documentation and scripts as appropriate to build/run your solution (points are awarded for ease of validating your work).  The solution will be validated on Mac OS X, but points are awarded for solutions that take into account other platforms (Windows, Linux etc).

The User entity should contain the following basic information.  You can add others if you wish, but these are the minimum:

    Username – a login name the user will use
    Password – the password to associate with the account. Hint: think about this one.
    First Name – the first name of the person
    Middle Name – middle name of the person
    Last Name – last name of the person
    Email – Email address for the person
    Telephone – primary phone number for the person


------

## Dev Log
04/06
- 12:01 PM - Start research on Docker #2 
- 1:00 PM - Pull Postgress container
- 1:30 PM - Basic experimentation 
- 3:00 PM - Barebones containerized HTTP Server created
- 3:30 pm - Implement user model, local backend api
- 4:00 pm - working through postgress blocker: create database command doesn't do anything
    - might be a non-issue in the long run as the goal is to use Go to interact with DB
    - unblocked, SQL is case sensitive; run CREATE DATABASE <name> instead
 - 4:30 pm - DB created, break time
 - 6:30 PM - implement create and update. postgres is pretty straightforwards

04/07
- ran into critical error I have no idea how to solve
    - local backend (go run ./) connects to DB just find and CRUD opss work perfectly
    - however, Docker image of my app fails to connect to DB
    - panic: dial tcp 127.0.0.1:5432: connect: connection refused
    - there may just be a piece of the puzzle that i'm missing


------ 
## Issues and Solutions
- Issue: docker: Error response from daemon: OCI runtime create failed: container_linux.go:367: starting container process caused: exec: "app": executable file not found in $PATH: unknown.
    - Solution: set up Dockerfile correctly. https://golangdocs.com/golang-docker

- Issue: Couldn't connect to Docker daemon at http+docker://localhost - is it running? If it's at a non-standard location, specify the URL with the DOCKER_HOST environment variable.
    - Solution: didn't have permssions, run "docker-compose up" as sudo
