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

<!-- 20221103 -->

- Buildkit example `build-using-dockerfile`, take 2
  - It does compile ✔️
  - One **must** bear in mind the example is a *client*. In other words the environment must comply with: <!-- fun story, I burned must of my study time realizing this ^ -->
    - A buildkit daemon
    - A container runtime (runc, crun or containerd)
  - I tried to let it connect to the buildkit container that `dagger` spun up for me weeks ago.
    - It didn't work
    - I used the the `BUILDKIT_HOST` env and the `--buildkit-addr` with no luck [ref](https://github.com/moby/buildkit/blob/master/examples/build-using-dockerfile/main.go#L37-L42)
    - The buildkit URI used was `docker-container://<container-name>`
      - ^ I blindy tried it since that is the way dagger does it [ref](https://github.com/dagger/dagger/blob/main/internal/buildkitd/buildkitd.go#L180)
      - ^ I'm not sure if that should work or there's a "connection helper" within dagger's source I am overlooking
      - ^ I bet there's one, since [gRPC Name Resolution](https://github.com/grpc/grpc/blob/master/doc/naming.md) does not mention the `docker-container` schema
  - Thanks to `delv`, I was able to debug the code, then I understood it won't worked until I fulfill the aforementioned dependencies.
    - ^ Although the 3 go-routines run within a `errgroup.Go` the main thread still hangs. I thought it would cancel the whole group as soon one of those failed.

Next Goal: Try approaching the following:

  - ~~Compile it `build-using-dockerfile` then~~ try it out
  - Try swaping `dockerfile.v0` with `gateway`
  - Look around why they couldn't use `containerd` from the get-go

Next Goal: Compile and run buildkit's examples
