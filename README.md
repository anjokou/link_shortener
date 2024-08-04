# Running

The project is configured with a docker-compose, so both services can be run with the docker-compose command:

```sh
docker compose up
```

If you want to run the services on the host machine, they can be run with the command

```sh
go run .
```

Building link_shortener requires gcc to be installed and available though the PATH variable.

Individual runtime details are in each service's README file.