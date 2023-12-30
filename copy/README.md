# `COPY` Instruction
`COPY` is an instruction that can be used to add files from 
the source to the destination folder in a Docker Image.

- The fundamental difference from the `ADD` instruction is that `COPY` only copies the files, 
 while `ADD`, in addition to copying, also downloads the source from a URL and automatically extracts compressed files.
- As a best practice, it is advisable to use `COPY` whenever possible. 
- If it is necessary to extract compressed files, use the `RUN` instruction and 
run the application to extract those compressed files.
- Copying multiple files at once in the `COPY` instruction uses the pattern 
in Golang: https://pkg.go.dev/path/filepath#Match

## Format Instruction
```shell
COPY {source} {destination}
```

### ***example***
- Copy helloworld.txt file to `playground` directory
```shell
COPY helloworld.txt playground
```

- Copy all files with suffix/extension `.txt` to `playground` directory
```shell
COPY *.txt playground
```

- Copy single-character name of `.txt` file to `playground` directory
```bash
COPY ?.txt playground
```

## Step by step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/copy-example:latest copy
```

2. See all the image to make sure our image is created properly
```shell
docker image ls
```

3. Create a container from the image
```shell
docker container create --name copy-example rambokong/copy-example
```

4. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

5. Start the container
```shell
docker container start copy-example
```

6. See container logs output
```shell
docker container logs copy-example
```