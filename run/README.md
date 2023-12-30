# `RUN` Instruction
`RUN` is an instruction which only executed during the build process of a Docker image. 
- Once the Docker image is formed, `RUN` instruction will not be executed.
- In single dockerfile, it can contains multiple `RUN` instruction

## Format Instruction
```shell
RUN {your_command}
```

### ***example***
- Create a directory `playground`
```shell
RUN mkdir playground
```

- Create new `helloworld.txt` with text `hello world!` in it
```shell
RUN echo "hello world!" > helloworld.txt
```

- See a file with name `helloworld.txt`
```shell
RUN cat
```

## Step by Step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/run-example:latest run
```

2. See all the image to make sure our image is created properly
```shell
docker image ls
```

3. We will re-create the image with `--progress` and `--no-cache` flags to see the details when image is being build, and it won't set any cache
by removing the previous image first
```shell
docker image rm rambokong/run-example

docker build -t rambokong/run-example:latest dockerfile/run --progress=plain --no-cache
```

4. See all the image to make sure our image is created properly
```shell
docker image ls
```