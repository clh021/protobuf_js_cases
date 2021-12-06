
const PROTO_PATH = `./window.proto`;

const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
  },
);
const lnks_proto = grpc.loadPackageDefinition(packageDefinition).lnks;
const rpc = lnks_proto.Test.service.send
const rpc2 = lnks_proto.Test2.service.send

// console.log("lnks_proto:", lnks_proto)
// var winId1 = lnks_proto.WindowId;
// winId1.setOk(true);
// var winList = new lnks_proto.MainWindowList();

// console.log(winId1);

console.log(rpc.requestSerialize({ ok: true, windows: [{ chromium_id: 1 }] }));

console.log(rpc2.requestSerialize({ ok: true, id: { chromium_id: 1 }, title:"test" }));
