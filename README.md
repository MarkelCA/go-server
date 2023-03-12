# go-server
Go backend with plain standard libraries to learn the language implementation on http, authentication, security and other protocols.

## Installation
Clone the repository:
```
git clone https://github.com/MarkelCA/go-server.git
```

## Usage
Change directory inside the repo and start the server:
```
cd go-server
docker compose up
```

Start making requests:
```
curl localhost:8080/hello
```

## Additional info
### Conventions
The conventions and practices applied in this project try to comply with the community standards. To accomplish that I've being ispired by the following projects:
- [Docker Compose](https://github.com/docker/compose)
- [Go Ethereum](https://github.com/ethereum/go-ethereum)
- [Go Chi](https://github.com/go-chi/chi)

### Multi-stage build
The docker container is built using [multi-stage builds](https://docs.docker.com/build/building/multi-stage/). This reduces the container's weight from 846MB to 13.7MB at this moment.
