# Cryptowatch protobuf definitions

This repo contains protobuf definitions for publicly accessible APIs.

Refer to our [Websocket Streaming API docs](https://cryptowat.ch/docs/streaming-api) for more information.

## Deprecated Values
All `float` types are deprecated in favor of using `string` for cross-platform consistency. Every `float` value now has a corresponding `string` value. For example, `price` would now be `priceStr`, and parsing is left to the client.
