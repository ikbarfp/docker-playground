# `ADD` Instruction
`ADD` is an instruction that can be used to add files from the source to the destination folder in a Docker Image.

- The `ADD` command can detect whether a source file is a compressed file such as tar.gz, gzip, and others. If detected as a compressed file, the file will be automatically extracted into the destination folder.
- The `ADD` command can also support adding multiple files at once.
- Adding multiple files at once in the `ADD` instruction uses the pattern in Golang: https://pkg.go.dev/path/filepath#Match

## Format Instruction
```shell
ADD {source} {destination}
```

### ***example*** 
- Add `helloworld.txt` file to `playground` directory
```bash
ADD helloworld.txt playground
```

- Add all files with suffix/extension `.txt` to `playground` directory  
```bash
ADD *.txt playground
```

- Add single-character name of `.txt` file to `playground` directory
```bash
ADD ?.txt playground
```

## Step by step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/add-example:latest add
```

2. See all the image to make sure our image is created properly
```shell
docker image ls
```

3. Create a container from the image
```shell
docker container create --name add-example rambokong/add-example
```

4. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

5. Start the container
```shell
docker container start add-example
```

6. See container logs output
```shell
docker container logs add-example
```



