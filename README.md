
# Description

This is the API for dating APP build with golang, gin gonic, and postgresql.

## Structure

- app
    - domain
        - entity : domain entities/model located in here
        - service : business logic happen
        - repository : the adapter for data needs (database, other service)
    - handler : controller/handler is located in here


## How to run
```
# copy the example.env to .env
$ cp example.env .env

# prepare all dependencies
make prepare

# migrate table and seed data
$ make migrate

# run the application
make run
```

## How to test
```
make test
```

