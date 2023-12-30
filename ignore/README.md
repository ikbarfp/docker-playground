# `.dockerignore` File
When we execute the `ADD` or `COPY` instructions from the source file, Docker first reads the file named `.dockerignore`.

- The `.dockerignore` file is similar to the `.gitignore` file, where we can specify the files we want to ignore.
- If a file is mentioned in the `.dockerignore` file, then automatically won't be added or copied.
- The `.dockerignore` file also supports ignoring folders or using regular expressions.

## Step by step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/ignore-example:latest ignore
```

2. See all the image to make sure our image is created properly
```shell
docker image ls
```

3. Create a container from the image
```shell
docker container create --name ignore-example rambokong/ignore-example
```

4. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

5. Start the container
```shell
docker container start ignore-example
```

6. See container logs output
```shell
docker container logs ignore-example
```