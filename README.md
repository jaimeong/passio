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
- 12:01 PM - Start research on Docker #2 
- 1:00 PM - Pull Postgress container
- 1:30 PM - Basic experimentation 


------ 
## Issues and Solutions
- Issue: docker: Error response from daemon: OCI runtime create failed: container_linux.go:367: starting container process caused: exec: "app": executable file not found in $PATH: unknown.
    -Solution: set up Dockerfile correctly. https://golangdocs.com/golang-docker
