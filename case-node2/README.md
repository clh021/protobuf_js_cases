## Use

```shell
  npm install
  node test.js
```

## Result

```shell
make run
node test.js
<Buffer 08 01 12 02 08 01>
```

## build to use in gecko engine
```shell
# browserify index.js --bare > dist.js
npm i -g @vercel/ncc
ncc build index.js -o dist
```