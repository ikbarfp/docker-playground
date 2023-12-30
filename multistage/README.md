# Multistage Concept
This is a concept where we aim to avoid excessively large image sizes. Therefore, we can leverage the multistage concept
in Docker. The key to this concept lies in using multiple `FROM` instructions, which serve as stages for creating an image.

## Example in the Golang:

- We can first build the binary using an available base image (for example, `golang:1.XX-alpine`).
> [!WARNING]  
> Make sure when you built executable with same operating system with following image.
> In this case we build golang app on top of golang-alpine image
> and then we create final image using base image alpine too.

- After building the executable binary of our Go application from the previous base image, we then copy the resulting 
executable file into a new image that will use a smaller-sized base image (for example, `alpine:XX`).

- The built image will be derived from the last `FROM` instruction stated in the `Dockerfile` 
(which is taken from `alpine:XX`), making the final image size smaller.

## Step by step
1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/multistage-example:latest multistage
```

2. See all images to make sure our image is created properly
```shell
docker image ls
```

3. Create a container from the image with custom port to see the differences
```shell
docker container create --name multistage-example -p 8080:8080 rambokong/multistage-example
```

4. See all the container to make sure our container is created properly
```shell
docker container ls -a
```

5. Start the container
```shell
docker container start multistage-example
```