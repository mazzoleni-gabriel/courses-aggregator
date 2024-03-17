# courses-aggregator

This application is responsible for aggregating courses information with enrollments users

## To run this application locally

- Run `Make run`

By now your application should be up and running! Try it out:

#### Ping

`curl --location --request GET 'localhost:8080/ping`

#### Aggregate courses

`curl --location --request GET 'localhost:8080/courses' \
--header 'apiKey: your-token-here'`
