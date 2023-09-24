# nats-exploration

Just an exploration into NATS.

The goal is to create a limited-size multi-consumer queue that allows replay and at-least-once delivery (per subject).

## Setup

```bash
# Mac
brew tap nats-io/nats-tools
brew install nats-io/nats-tools/nats
```

## Installing NATS

Install NATS from the Github releases page.

Then start NATS + Jetstream with:

```bash
nats-server -js
```

## Creating the first stream

```bash
# follow guide and create the new stream
nats stream add my-stream

# print info about the stream
nats stream info my-stream
```
