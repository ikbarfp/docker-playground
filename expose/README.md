# `EXPOSE` Instruction
`EXPOSE` is an instruction to indicate that the container will listen on a specific port number and protocol.

- Essentially, the `EXPOSE` instruction itself does not actually publish any ports.
- `EXPOSE` instruction is used as additional information/documentation for the creator of the Docker container, 
indicating that this Docker image will use a specific port when run as a Docker container.
- If you not specifed your preferred protocol, it will be a **default** which is `tcp`

## Format Instruction
```shell
EXPOSE {your_port}/{your_protocol}
```

### ***example***
- State port `8080` with protocol `tcp`
```shell
EXPOSE 8080/tcp
```

- Expose port `9090` with protocol `udp`
```shell
EXPOSE 9090/udp
```

œ