# MEND 

This is a simple CRUD application that supports SQL(postgres) and NoSQL(mongodb) database.

## Local setup

Make sure that you have go installed.

1. Make sure that you are running go version `1.17`.

2.  Install dependencies.
```shell
go mod download & go mod vendor
```

3. TLS certs generation
```shell
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

4. Use docker-compose setup. App will run on `8080` port receiving only `https` traffic.
```shell
docker-compose -f docker-compose-mongo.yaml
```
or 
```shell
docker-compose -f docker-compose-psql.yaml
```


## Development

1. Linter
```shell
make lint
```

2. Tests
```shell
make test
```

3. Build
```shell
make build
```

4. Generate mocks
```shell
make generate
```

## Thoughts

### Take a look at `NOTE` in comments

### Project structure
I picked MVC-like structure for that although I like to structure my project based on domains when it's supposed to be a bigger project. 

### DB clients
I did not use orm for psql approach because I kind of like the control over run queries and abstraction provided in those lib sometimes is just a pain in the ass if you want to do something that is not just simple queries.

I have no experience in MongoDB so I'm not sure that holding `*mongo.Collection` as client is a way to go but it seems like nice separation - `Mongo` structure in this project is only responsible for managing `user` collection. 

### Place to define interfaces
As stated in `NOTE` it's a tricky thing to have only one approach. Read `NOTE` and check those articles for more info:
```
https://github.com/golang/go/wiki/CodeReviewComments#interfaces
https://rakyll.org/interface-pollution/
https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html
```

### Graceful shutdown
The implementation is missing some parts e.g. graceful shutdown - an app should listen for os signals and gracefully shutdown whole thing(close db client connections). It was skipped for simplicity and this thing took me even more as expected :D 

Nice article about that: https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/