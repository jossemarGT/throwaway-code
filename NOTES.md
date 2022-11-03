# Notes

## First round
<!-- 20221101 -->

- Read Mockerfile post and jump through the code
  - <https://matt-rickard.com/building-a-new-dockerfile-frontend>
  - <https://github.com/r2d4/mockerfile>

- Found Gockerfile, which offers a reduced DSL to interact with it
  - <https://github.com/po3rin/gockerfile>

- What is the Dockerfile frontend entrypoint?
  - Perhaps from the client reference? <https://github.com/tonistiigi/buildkit-pack/blob/master/manifest.go#L52>

- Buildkit's README introduces on `llb` <https://github.com/moby/buildkit#exploring-llb>
- It seems it is best to go w/buildkit examples. The first frontends I found were quite dated.

<!-- 20221102 -->

- Buildkit example `build-using-dockerfile`
  - Uses the moby's default frontend [import](https://github.com/moby/buildkit/blob/master/examples/build-using-dockerfile/main.go#L14)[impl](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/builder/build.go)
  - In a nutshell it:
    - Generates buildkit client Opts from cli's arguments [ref](https://github.com/moby/buildkit/blob/master/examples/build-using-dockerfile/main.go#L126)
    - ^ it mentions using `gateway` instead of `dockerfile.v0` [ref](https://github.com/moby/buildkit/blob/master/examples/build-using-dockerfile/main.go#L143)
    - ^ it mentions using `containerd` instead of `docker` for exports [ref](https://github.com/moby/buildkit/blob/master/examples/build-using-dockerfile/main.go#L166)
    - Creates a pipe. The writer goes to the builder, the reader goes to the "load" mechanism
    - Waits for writer, reader and "UI" routines to finish its work
  - As result you'll get the OCI image as tar which is loaded into docker

- It would be nice to:
  - First compile it then try it out.
  - Try swaping `dockerfile.v0` with `gateway`
  - Look around why they couldn't use `containerd` from the get-go

Next Goal: Tackle "nice" bullets above ^ the move on
Next Goal: Compile and run buildkit's examples
