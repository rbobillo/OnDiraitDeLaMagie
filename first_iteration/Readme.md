# First Iteration

### Create *'Magic'*: a simple CRUD web service, with a rest API
 - The service should be implemented in Go
 - It should run in a `Docker` container
 - It should interact with *'Magic Inventory'*: a PostgreSQL database (which should run in another `Docker`)
 - The service and the DB, should be deployed on Docker containers, using `ansible`, `terraform` or `docker-compose`
 - According to the [magic documentation](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20magic.pdf)
	- The API should implement routes (`/wizards`, `/spawn`, `/age`...) with the right HTTP verbs (`GET`, `POST`...)
	- The API should be documented, using swagger
	- The service should populate *'Magic Inventory'* with the right fields (`id`, `first_name`...)
	- Why not put data in the DB, using the [wizard generator](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/reference/generate_random_wizards.go), using it to request the *'Magic'* service

### Run (the fully Dockerized app) with
```
docker-compose up
```

### Run 'magic' locally, with a database container
```shell
# in one term, start the database container, and wait for it to run
docker-compose up magic_inventory

# in another term, load env variables, and run the 'magic' service locally
source magic/magic.env && go run magic/main.go
```

### Once the containers run
 - First, request this URL: `http://localhost:9090/` (it should give you the API description)
 - Then, this URL: `http://localhost:9090/wizards` (it should give you a list of existing Wizards, from the `magic_inventory`)
 - And now, just implement everything, according to the [magic Readme](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/reference/first_iteration/magic/Readme.md)