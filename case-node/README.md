# JavaScript protocol buffers example

Simple protocol buffers example written in JavaScript.

## Requirements

- Node.js and npm
- protoc

## Installation

```shell
  npm install
```

## Usage

First you have to generate the JavaScript classes from the `.proto` file:

```shell
  protoc \
    --proto_path=protocol_buffers/definitions \
    --js_out=import_style=commonjs,binary:protocol_buffers/messages \
    protocol_buffers/definitions/*.proto
```

Now you can use the generated getters, setters and methods for the serialization.
Run `index.js` for the example:

```shell
  node index.js
```

## Testing

Run JSHint checks:

```shell
  docker-compose run js-tools jshint index.js
```

Run JSCS checks:

```shell
  docker-compose run js-tools jscs index.js
```

## License

This project is licensed under the terms of the [MIT License (MIT)](LICENSE).

## Result

node test.js
{
wrappers*: { '2': [ [Object] ] },
messageId*: undefined,
arrayIndexOffset*: -1,
array: [ true, [ [Array] ] ],
pivot*: 1.7976931348623157e+308,
convertedPrimitiveFields*: {}
}
encode: Uint8Array(6) [ 8, 1, 18, 2, 8, 1 ]
decode: {
wrappers*: { '2': [ [Object] ] },
messageId*: undefined,
arrayIndexOffset*: -1,
array: [ true, [ [Array] ] ],
pivot*: 1.7976931348623157e+308,
convertedPrimitiveFields*: {}
}
