# Golang OOP Concepts
blah blah about OOP and Go! (to be completed...)

In order to follow the instruction and run the test inside the container, you need to have the docker installed on your host machine.

* **Version:** [1.2](https://golang.org/doc/go1.12)
* **Assertion, Mock & Test Suits:** [Testify](https://github.com/stretchr/testify)

## Makefile commands
You could run the commands by starting: `~$ make up`
* `build` to build the container
* `build-force` to build the container ignoring cached packages
* `up` and `down` used to up and down the Golang container
* `ssh` to gain an access to tty of running container

### Docker Cheat sheet
__Build and Run the Golang container:__
```bash
~$ docker-compose -f docker/docker-compose.yml up -d
```

__Enter the Golang Container:__
```bash
~$ docker-compose -f docker/docker-compose.yml exec golang bash
```

__Shutdown the Golang container:__
```bash
~$ docker-compose -f docker/docker-compose.yml down
```

### IDE Cheat sheet
In order for IDE to detect the location based on replace path in `go.mod`. Please create a symlink like command below:
`ln -s ln -s ~/dev/go/oop/src ~/dev/go/src/github.com/neg0/go-oop`