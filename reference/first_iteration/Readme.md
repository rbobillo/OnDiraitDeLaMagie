# First Iteration
===

### Create *'Magic'*: a simple CRUD web service, with a rest API
 - The service should be implemented in Go
 - It should run in a `Docker` container
 - It should interact with *'Magic Inventory'*: a PostgreSQL database (which should run in another `Docker`)
 - The service and the DB, should be deployed on Docker containers, using `ansible`, `terraform` or `docker-compose`

### According to the [documentation](documentation/On Dirait De La Magie.pdf)
 - The API should implement routes (`/wizards`, `/spawn`, `/age`...) with the right HTTP verbs (`GET`, `POST`...)
 - The API should be documented, using swagger
 - The service should populate *'Magic Inventory'* with the right fields (`ID`, `FIRST_NAME`...)
 - Why not put data in the DB, using the [wizard generator](generate_random_wizards.go), using it to request the *'Magic'* service
