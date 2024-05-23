
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
$ make prepare

# migrate table and seed data
$ make migrate

# run the application
$ make run

# run from the docker (make sure all table is migrate)
$ make up
```

## How to test
```
make test
```

## Sample Documentation
### Register
```
POST {{base_url}}/register

request body
{
    "name": "test name", // required string, max 50 char
    "email": "test_mail@yopmail.com", // required string, max 50 char
    "password": "123456" // required string, min 6 char
}
```
### Login
```
POST {{base_url}}/login

request body
{
    "email": "test_mail@yopmail.com", // required string
    "password": "123456" // required string
}
```
### List User Profiles
```
GET {{base_url}}/candidates
Authorization: Bearer {{access_token}}

query params
- "limit": 5, // optional integer
```
### Swipe
```
POST {{base_url}}/pair
Authorization: Bearer {{access_token}}

request body
{
    "pair_user_id": 1, // required integer
    "status": 1 // required integer. 0 = pass, 1 = like
}
```
### List Packages
```
GET {{base_url}}/packages
```
### Purchase
```
POST {{base_url}}/purchase
Authorization: Bearer {{access_token}}

request body
{
    "package_id": 1 // required integer
}
```