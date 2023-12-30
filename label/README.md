# `LABEL` Instruction
`LABEL` is an instruction to specify our additional metadata for our docker container, but it's not necessary.

1. Build new image based on our Dockerfile
```shell
docker build -t rambokong/label-example:latest label
```

2. See all the image to make sure our image is created properly
```shell
docker image ls
```