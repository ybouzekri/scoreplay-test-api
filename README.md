# Test Golang API

## Start application

```sh
go run ./main.go
```

You can use the `-h` flag to list all the possible flags.

## Architecture
This API is built using clean architecture approche based on this article [The Clean Architecture by Robert C. Martin (Uncle Bob)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Thus there is three main parts in the app:
- adapters:
    - handlers: contains the controllers for the different requests, each handler handles one request with an extra one for errors
    - repositories: contains the actual repositories used 
- business:
    - entities: contains the entities used, media and tags in our case
    - repositories: contains the repositories interface
    - usecases: this is were we handle interaction between elements of the domain
- driver:
    - rest: contains all the routes and server for the rest api

I chose to go with clean archi as it largely used in the industry today which makes it well documented.
It also allows to dive quickly into the business code, making it clearer, easy to ready, maintain and evolve

I chose also to keep the repositories as in memory instead of a bdd, so I can focus more on the architure of the code which is important and using clean architecture makes adding a bdd easier

## What to improve
- writing more tests, adding specific cases in functional tests
- Secure file upload
- Adding a real DB instead of in memory
- Adding cache on top with decorators over the repositories
- Adding metrics with prom or open telemetry to monitor the system and how it behaves
- Dockerise the app
- Add readiness and liveness endpoints in case of a kubernetes deploy
- We can add FindByIDs for tags to replace [this part](https://github.com/ybouzekri/scoreplay-test-api/blob/master/internal/business/usecases/create_media.go#L35)