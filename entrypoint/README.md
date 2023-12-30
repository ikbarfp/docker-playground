# `ENTRYPOINT` Instruction
`ENTRYPOINT` is an instruction used to specify the executable file that will be run by the container.

- Typically, `ENTRYPOINT` is closely related to the `CMD` instruction.
- When we create a `CMD` instruction without an executable file, `CMD` will automatically use 
the `ENTRYPOINT` instruction.
- When you passed parameters into `CMD` instruction, these parameters will automatically sent 
to `ENTRYPOINT` instruction

## Format Instruction
```shell
ENTRYPOINT {YOUR_EXECUTABLE} {PARAM_1} {PARAM_2} . . .
```

### _**example**_
- Create an `ENTRYPOINT` combined with `CMD`
```shell
# container will be built with command go run /app/main.go
ENTRYPOINT ["go", "run"]
CMD ["/app/main.go"]
```

## Step by Step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/entrypoint-example:latest entrypoint
```

2. See all images to make sure our image is created properly
```shell
docker image ls
```

3. Inspect the image
```shell
docker image inspect rambokong/entrypoint-example
```

4. Create a container from the image with custom port to see the differences
```shell
docker container create --name entrypoint-example -p 8080:8080 rambokong/entrypoint-example
```

5. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

6. Start the container
```shell
docker container start entrypoint-example
```