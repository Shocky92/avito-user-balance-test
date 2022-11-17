# avito-user-balance-test
Test project for avito intern backend program

## Local environment

Docker compose with Go + Postgres + PgAdmin

To start docker containers:
`docker-compose up -d`

To stop docker containers:
`docker-compose stop`

Application available on URL:
`localhost:3000`

PgAdmin available on URL:
`localhost:5050`

## List of HTTP methods

Get user balance
`[GET] /user/balance/:id`

Increase user balance 
`[PUT] /user/balance/:id/increase`

Reserve user balance for order
`[POST] /user/:id/order/reserve`

Proceed reserved money for order
`[POST] /user/:id/order/proceed`
