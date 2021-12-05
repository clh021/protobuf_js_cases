/* jshint devel: true */

'use strict';

var message = require('./protocol_buffers/messages/window_pb');

// console.log(window);
var windowId = new message.WindowId();
windowId.setChromiumId(1)

var windowList = new message.MainWindowList();
windowList.setOk(true);
windowList.addWindows(windowId);

console.log(windowList);

var encode = windowList.serializeBinary();
console.log("encode:", encode);

var decode = message.MainWindowList.deserializeBinary(encode);
console.log("decode:", decode);

