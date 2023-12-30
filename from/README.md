# `FROM` Instruction
`FROM` is an instruction that can be used to as the definition of base image for the image to be created.
- 
- You can indeed create an image from scratch without a base image, but that is rarely used in practice. 

## Format Instruction
```shell
FROM {base_image}:{base_image_tag}
```

### example
- Get base image from Golang 1.21 version
```shell
FROM go:1.21
```

## Step by step
1. Build new image based on our Dockerfile 
```shell
docker build -t rambokong/from-example:latest from
```

2. See all the image to make sure our image is created properly
```shell
docker image ls
```