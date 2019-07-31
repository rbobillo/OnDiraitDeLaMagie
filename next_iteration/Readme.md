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

### Run with
```
docker-compose up
```