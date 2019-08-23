# Azkaban Service

### Always remember
 - Look at the [magic documentation](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20magic.pdf)

### What needs to be done
 - `owls`, a dockerized **RabbitMQ** message broker (http:localhost:5672)
 - A simple app, `azkaban`, that interacts with `owls` and other services, running in a Docker container
 - According to [documentation](https://github.com/rbobillo/OnDiraitDeLaMagie/blob/master/documentation/On%20Dirait%20De%20La%20Magie%20-%20azkaban.pdf) it should interact with other services, using **RabbitMQ**, or http requests
   - Send / Receive JSON messages via RabbitMQ to communicate with `ministry` and `magic`
     - Receive:
       - Arrest: { "id": UUID, "action": "arrest", "wizard_id": UUID }
     - Send:
       - Arrested: { "id": UUID, "action": "arrested", "wizard_id": UUID }
   - Send http requests to `magic` to interact with it
       - PATCH: http://localhost:9090/wizards/{id}/jail { "arrested": true }

### What has been done already
 - A **RabbitMQ** message broker can run locally (http://localhost:5672) in a Docker container named `owls`
