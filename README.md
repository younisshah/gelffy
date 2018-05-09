## gelffy

**gelffy** is a UDP-based GELF ([Greylog Extended Log Format](http://docs.graylog.org/en/2.2/pages/gelf.html)) aggregator with an HTTP POST hook to the LaaS - ([Loggly](https://www.loggly.com)).

gelffy is supposed to be used in the context of Docker. A Docker container (including a *service* in Docker Swarm mode) can configured to do logging using various *[drivers](https://docs.docker.com/config/containers/logging/configure/#supported-logging-drivers)*. Whenever anything
is written to the *stdout* or *stderr* of a Docker container, the Docker Engine captures the log entries and processes them using the configured driver.

 **gelffy** uses **gelf** Docker log driver.

 **gelffy** runs on port **12123**.

When a Docker container or a Docker Swarm service is configured to use *gelf* as its logging driver with *gelffy* as the endpoint for the GELF log entries,
the Docker Engine sends the log entries to the gelffy's UDP endpoint. gelffy processes the log entries and does an HTTP POST to Loggly.

### Pre-requisites

[A Loggly customer token](https://www.loggly.com/docs/logging-setup/)


#### Installation and Usage

1) Clone the repository
```ini
    git clone https://github.com/younisshah/gelffy.git
```

2) Change your working directory to *gelffy*

3) Edit the *Dockerfile* and update the value of *LOGGLY_CUSTOMER_TOKEN Env* variable to your Loggly Customer Token.

3) Run
```ini
dep ensure
```

4) Next, run
```ini
make build-bin
```

5) Then, do a
```ini
make build
```

6) Finally,
```ini
make run
```

7) Get the container IP address of *gelffy*
```ini
docker inspect --format '{{ .NetworkSettings.IPAddress }}' gelffy
```

8) Run your Docker container (or Docker Swarm service) which *gelf* log driver and *gelffy* configured
```ini
docker run -p IN_PORT:OUT_PORT --rm --log-driver=gelf --log-opt gelf-address=udp://GELFFY_CONTAINER_IP_ADDRESS:12123 --log-opt gelf-compression-type=none IMAGE_NAME:TAG
```

That's it!

Now all the logs of your Docker container (or Docker Swarm service) will be pushed to gelffy which POSTs to Loggly.

### TODOs
- [ ] Support additional drivers


#### License

MIT
