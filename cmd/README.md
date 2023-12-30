# CMD Instruction
`CMD` is an instruction that will only be executed when the Docker container is running.

- During the build process of a Docker image, instruction with prefixed with the `CMD` syntax will not be executed.
- In a single Dockerfile, only **one** `CMD` command can be run. 
- If you add more than one command, only the last `CMD` command will be executed.

## Step by step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/cmd-example:latest cmd
```

2. See the details of the image we've had created
```shell
docker image inspect rambokong/cmd-example
```

3. We will create a container with created image above to see if `CMD` instruction we prompted had executed or not
```shell
docker container create --name cmd-example rambokong/cmd-example:latest
```

4. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

5. Start the container
```shell
docker container start cmd-example
```

6. See logs from our container
```shell
docker container logs cmd-example
```