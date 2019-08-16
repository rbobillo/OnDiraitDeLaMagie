# Next Iteration

### Create a micro-services architecture (according to [documentation](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie.pdf)), with:
 - Dockerized RabbitMQ Message Broker:
   - **owls** http://localhost:5672
 - Dockerized PostgreSQL databases:
   - **magicinventory** http://localhost:5432
   - **hogwartsinventory** http://localhost:5433
   - **azkabaninventory** http://localhost:5434
 - Dockerized REST Web Services (`swagger` documented API):
   - **magic** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20magic.pdf)): http://localhost:9090
   - **hogwarts** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20hogwarts.pdf)): http://localhost:9091
   - **families** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20families.pdf)): http://localhost:9092
   - **guests** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20guests.pdf)): http://localhost:9093
   - **villains** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20villains.pdf)): http://localhost:9094
 - Dockerized Applications:
   - **ministry** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20ministry.pdf))
   - **azkaban** ([doc](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20azkaban.pdf))

### Each service must be as much responsive and resilient as possible (cf [Reactive Manifesto](https://www.reactivemanifesto.org/))
 - If an error occurs during an action, this action should be replayed
 - If another service is not available, this should not be considered as an error for the current service
   - For example: interaction between `ministry` and `hogwarts`:
     - `hogwarts` asks for Help via `owls`, and crashes because of an unhandled error (container stops, some other error...)
     - `ministry` receives the Help message, and tries to PATCH http://hogwarts:9091/actions/protect, but `hogwarts` is unavailable
     - `ministry` must be able to replay this action when `hogwarts` is available
     - this action must still be replayable if `ministry` crashes before replaying the action

### Run with
```
docker-compose up
```
