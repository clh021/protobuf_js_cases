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


var windowList2 = new message.MainWindowList();
windowList2.setOk(1);
windowList2.addWindows(windowId);

console.log(windowList2);

var encode2 = windowList2.serializeBinary();
console.log("encode2:", encode2);

var decode2 = message.MainWindowList2.deserializeBinary(encode2);
console.log("decode2:", decode2);

var winList3 = new message.MainWindowList3();
winList3.setOk(true);
winList3.addChromiumIds(1);
console.log(winList3);
var encode3 = winList3.serializeBinary();
console.log("encode3:", encode3);
var decode3 = message.MainWindowList3.deserializeBinary(encode3);
console.log("decode3:", decode3);


var winList4 = new message.MainWindowList4();
winList4.addChromiumIds(1);
winList4.addChromiumIds(1);
console.log(winList4);
var encode4 = winList4.serializeBinary();
console.log("encode4:", encode4);


var winList5 = new message.MainWindowList5();
winList5.setChromiumIds(1);
winList5.setChromiumId(1);
console.log(winList5);
var encode5 = winList5.serializeBinary();
console.log("encode5:", encode5);
