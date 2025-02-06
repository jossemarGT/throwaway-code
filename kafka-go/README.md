# Kafka with Go

**Status**: Done

## Topics to explore (plus notes)

- [x] Basic producer / consumer
  - This is silly but use `.newConfig()` instead of instantiating `Config{}`, otherwise you won't get the safe default. (smh)
  - `IBM/sarama` abstracts enough, however the details are in the configurations prior to instantiate the client. Examples
    - SyncProducer requires to return success and error
    - It will set the client ID to `sarama`, and it doesn't complain about it unless you set it in debug mode
  - To enable sarama debug mode you simply need to pass a "logger" to `Logger` variable on its package
- [x] Use **dev containers** to spin all deps including a single node Kafka
  - Either you run everything within a single image (initial purpose) or spin a docker compose.
  - The "all in a single image" seems to have more capabilities like:
    - _features_, think of lambda layers
    - _args_, think of the container's command override
    - _dind_, when set it shares the docker socket within the container without any hazzle
  - Some `.devcontainer` configurations rely on the values of specific members, as for example:
    - _service_, only works when `dockerComposeFile` is set and selects the container to be used for development
    - _postCreateCommand_, will only run on the selected service when using docker compose
  - The `customizations` are vendor specific, so not everything will behave the same across IDEs
  - Turns out intellij has support for this, but still in beta and will need to download a 1GB+ client (wth?)
- [x] Solve/Search: How to horizontally scale consumers w/o task assignation collision?
  - No need to code, the answer is kafka consumer groups and there's a working example on sarama tools + docs
    - <https://github.com/IBM/sarama/blob/main/examples/consumergroup/main.go>
    - <https://github.com/IBM/sarama/blob/v1.45.0/consumer_group.go>
    - <https://kafka.apache.org/documentation/#basic_ops_consumer_group>
  - The key concept is that all the consumer must belong to the same group and the message be spread in multiple partitions.
- [x] Use Go workspaces, since consumer doesn't use any external lib than the kafka client
  - TIL that workspaces aren't for what I had in mind, since you can't "force" common libraries to the `go.work` file.
  - Instead, what it is good for is when you have multiple modules within the same source code and you need to reutilize their logic within the same repository.
    - This saves you from multiple `replace` statements on the `go.mod` file
    - An hypothetical would be that I needed to use some exported components from the producer module into the consumer one or viceversa (beware of circular deps)
- [ ] **Dropped** dig most frequent issues on kafka ie: topic partitioning
  - Requires a deep dive, gotta move to separate topic.
