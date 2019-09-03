#CHANGELOG

###v1.0.0
 - foundations of a micro-services architecture project
 - docker-compose used to orchestrate micro-services deployment
 - one basic CRUD service
   - **magic** : dockerized **Go** webservice with REST api
     - API documented with `swagger` (http://localhost:9090/)
     - handmade logging system
     - missing tests (unit + integration) that must be implemented
   - **magicinventory** : dockerized postgres database with one table (**wizards**)