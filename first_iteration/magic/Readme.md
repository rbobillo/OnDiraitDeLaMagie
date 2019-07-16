# Magic Service

### Always remember
 - Look at the [magic documentation](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20magic.pdf)
 
### What has been done already
 - First setup `magic` CRUD
   - Using **docker-compose**, and **Docker**:
     - `magic` can run locally (http://localhost:9090) in a Docker container named `magic`
     - A **postgreSQL** database can run locally (http://localhost:5432) in a Docker container named `magicinventory`
     - `magicinventory` is locally persisted _(so data is not lost if the container stops)_
     - `magic` knows every `magicinventory` properties (`host`, `password`...)
     - Everything can run simply using `docker-compose up`
   - `magic` is a simple service done in **Go**, that can interact with a **postgreSQL** DB
  - First implementation of `magic`
    - The service listens for actions on `localhost`, on port `9090`
    - It uses some cool Go libraries (`pq`, `gorilla/mux`...) to handle data reading/writing
    - The service exposes an API with some routes
      - GET  `/` _(should output a Swagger documentation of the API)_
      - GET  `/wizards` _(should output a JSON `{ "wizards" : [ ... ] }` with every wizards from `magicinventory`)_
      - GET  `/wizards/{id}` _(should output a JSON `{ "wizard" : {} }` with a single wizard (filtered by ID) from `magicinventory`)_
      - POST `/wizards/spawn` _(should create and insert a new wizard into `magicinventory`, and return a JSON describing him)_
      