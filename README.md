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
## How to Use
TL;DR run this in root folder
```
cd postgres && sudo docker-compose up -d && cd .. && sudo docker run -it --rm -p 8001:8000 application-tag
```

1. Spin up postgres container by changing directories into /postgres and running your equivalent docker-compose command
    - "sudo docker-compose up -d"
3. Spin up Go rest api by returning to root folder, and running your equivalent docker run command
    - "sudo docker run -it --rm -p 8001:8000 application-tag"

## Verifying CRUD
- Post/Create
    - curl -d @request.json -H "Content-Type: application/json" http://localhost:8001/api/user
    - will return a JSON of newly created account
- Put/Update
    - curl -d @request.json -H "Content-Type: application/json" -X PUT  http://localhost:8001/api/user/Marge
    - will return JSON of updated account
- Read/Get/Gets
    - curl -v http://localhost:8001/api/users
    - curl -v http://localhost:8001/api/user/Marge
    - will return all users or specified user
- Delete
    - curl -X DELETE  http://localhost:8001/api/user/Marge
    - will return a message "account deleted"


## Technical Considerations
1. Creating and updating passwords will store the hashed + salted password in Postgres
2. Get/Gets currently returns ALL account info. I'd imagine in a production environment there would be more scrunity about what information is sensitive. I'm sure that the hashed + salted passwords are secure, but it just feels like bad form.
3. Postgres is in a countainer and the API is in another container. I know that docker-compose can be used for multi-container apps, but I couldn't figure out how to build the yml file.
4. getUser could be improved a bit, currently it's working almost exactly the same as plural getUsers. In MongoDB there would be a FindOne function, but I'm not sure what the equivalent would be for SQL.
5. Username is the primary key. I think a primary key/identity tied with an account id would be good too; this allows for users to change their username.
6. Users table looks like this
```
                         Table "public.users"
    Column    |         Type          | Collation | Nullable | Default 
--------------+-----------------------+-----------+----------+---------
 username     | character varying(40) |           | not null | 
 passwordhash | character varying(96) |           | not null | 
 firstname    | character varying(40) |           | not null | 
 middlename   | character varying(40) |           | not null | 
 lastname     | character varying(40) |           | not null | 
 email        | character varying(40) |           | not null | 
```

## Extras
Accesssing Postgres container from local machine
```
sudo docker exec -it postgres_postgres_1 psql -U root
```

Building API (run in root)
```
sudo docker build -t application-tag .
```
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
    - also don't forget semicolon
 - 4:30 pm - DB created, break time
 - 6:30 PM - implement create and update. postgres is pretty straightforwards

04/07
- 1:00 pm ran into critical error I have no idea how to solve
    - local backend (go run ./) connects to DB just fine and CRUD ops work perfectly
    - however, Docker image of my app fails to connect to DB
    - panic: dial tcp 127.0.0.1:5432: connect: connection refused
    - there may just be a piece of the puzzle that i'm missing
- 1:30 pm remove refused connection blocker, solution was to change connection stringg
    - Docker's internal host IP on linux is default to 172.17.0.1
- 2:00 pm clean up loose ends before submitting

------ 
## Issues and Solutions
- Issue: docker: Error response from daemon: OCI runtime create failed: container_linux.go:367: starting container process caused: exec: "app": executable file not found in $PATH: unknown.
    - Solution: set up Dockerfile correctly. https://golangdocs.com/golang-docker

- Issue: Couldn't connect to Docker daemon at http+docker://localhost - is it running? If it's at a non-standard location, specify the URL with the DOCKER_HOST environment variable.
    - Solution: didn't have permssions, run "docker-compose up" as sudo

- Issue : panic: dial tcp 127.0.0.1:5432: connect: connection refused
    - Solution:: connect to container's internal host ip, 127.17.0.1
