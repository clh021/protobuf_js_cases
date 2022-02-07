/* jshint devel: true */

'use strict';

var message = require('./protocol_buffers/messages/window_pb');

var test1 = new message.Test1();
test1.setA(150);
console.log(test1);
console.log("test1:", test1.serializeBinary());
var test2 = new message.Test2();
test2.setB('test');
console.log(test2);
console.log("test2:", test2.serializeBinary());

// console.log(window);
// var windowId = new message.WindowId();
// windowId.setChromiumId(1)

// var windowList = new message.MainWindowList();
// windowList.setOk(true);
// windowList.addWindows(windowId);

// console.log(windowList);

// var encode = windowList.serializeBinary();
// console.log("encode:", encode);


// var winProps = new message.MainWindowProps();
// winProps.setOk(true);
// winProps.setId(windowId);
// winProps.setTitle("test");

// var encode = winProps.serializeBinary();
// console.log("encode:", encode);


// var decode = message.MainWindowList.deserializeBinary(encode);
// console.log("decode:", decode);

