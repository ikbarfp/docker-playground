# `VOLUME` Instruction
`VOLUME` is an instruction used to automatically create a volume when creating a Docker Container.

- All files present in the volume will be automatically copied to the Docker Volume, even if we do not explicitly 
create a Docker Volume when creating the Docker Container.
- It is particularly useful in cases where the application needs to store data in files, ensuring that the data is 
automatically secured within the Docker Volume.

## Format Instruction
```shell
VOLUME {yourDirectory/location}
```

### ***example***
- Create a volume
```shell
VOLUME myVolume/directory1
```

- Create multiple volume
```shell
# single line declaration
VOLUME myVolume/directory1 myVolume/directory2

# use array of string to make it more clear
VOLUME ["myVolume/directory1", "myVolume/directory2"]

```

## Step by step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/volume-example:latest volume
```

2. See the details of image we've been created, to make sure this image is volume at certain port
```shell
docker image inspect rambokong/volume-example
```

3. See all the image to make sure our image is created properly
```shell
docker image ls
```

4. Create a container from the image with custom port to see the differences
```shell
docker container create --name volume-example -p 8080:8080 rambokong/volume-example
```

5. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

6. Start the container
```shell
docker container start volume-example
```

7. See the container logs to make sure our container is running.
```shell
docker container logs volume-example
```

8. Access the `http://localhost:9999/:input_name`

9. See the details of container
```shell
docker container inspect volume-example
```
