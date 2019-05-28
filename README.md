# Golang OOP Concepts
blah blah about OOP and Go!....

In order to follow the instruction and run the test inside the container, you need to have the docker installed on your host machine.

* **Version:** [1.2](https://golang.org/doc/go1.12)
* **Assertion, Mock & Test Suits:** [Testify](https://github.com/stretchr/testify)

### Makefile
`Up` and `Down` used to up and down the Golang container; and `shell` to gain an access to tty of running container.

#### Docker cheat sheet
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
