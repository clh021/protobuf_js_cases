# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/window.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13protos/window.proto\x12\x04lnks\"X\n\x08WindowId\x12\x18\n\x0b\x63hromium_id\x18\x01 \x01(\x05H\x00\x88\x01\x01\x12\x15\n\x08gecko_id\x18\x02 \x01(\x05H\x01\x88\x01\x01\x42\x0e\n\x0c_chromium_idB\x0b\n\t_gecko_id\"=\n\x0eMainWindowList\x12\n\n\x02ok\x18\x01 \x01(\x08\x12\x1f\n\x07windows\x18\x02 \x03(\x0b\x32\x0e.lnks.WindowId\"\xef\x03\n\x0fMainWindowProps\x12\x1a\n\x02id\x18\x01 \x01(\x0b\x32\x0e.lnks.WindowId\x12\n\n\x02ok\x18\x02 \x01(\x08\x12\x12\n\x05title\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x17\n\nvisibility\x18\x04 \x01(\x08H\x01\x88\x01\x01\x12\x0e\n\x01x\x18\x05 \x01(\x05H\x02\x88\x01\x01\x12\x0e\n\x01y\x18\x06 \x01(\x05H\x03\x88\x01\x01\x12\x12\n\x05width\x18\x07 \x01(\x05H\x04\x88\x01\x01\x12\x13\n\x06height\x18\x08 \x01(\x05H\x05\x88\x01\x01\x12\x19\n\x0cwindow_state\x18\t \x01(\tH\x06\x88\x01\x01\x12\x1c\n\x0ftop_level_focus\x18\n \x01(\x08H\x07\x88\x01\x01\x12\x11\n\x04icon\x18\x0b \x01(\tH\x08\x88\x01\x01\x12\x10\n\x03url\x18\x0c \x01(\tH\t\x88\x01\x01\x12\x13\n\x06\x61pp_id\x18\r \x01(\tH\n\x88\x01\x01\x12\x15\n\x08skip_bar\x18\x0e \x01(\x08H\x0b\x88\x01\x01\x12\x1a\n\ralways_on_top\x18\x0f \x01(\x08H\x0c\x88\x01\x01\x42\x08\n\x06_titleB\r\n\x0b_visibilityB\x04\n\x02_xB\x04\n\x02_yB\x08\n\x06_widthB\t\n\x07_heightB\x0f\n\r_window_stateB\x12\n\x10_top_level_focusB\x07\n\x05_iconB\x06\n\x04_urlB\t\n\x07_app_idB\x0b\n\t_skip_barB\x10\n\x0e_always_on_topb\x06proto3')



_WINDOWID = DESCRIPTOR.message_types_by_name['WindowId']
_MAINWINDOWLIST = DESCRIPTOR.message_types_by_name['MainWindowList']
_MAINWINDOWPROPS = DESCRIPTOR.message_types_by_name['MainWindowProps']
WindowId = _reflection.GeneratedProtocolMessageType('WindowId', (_message.Message,), {
  'DESCRIPTOR' : _WINDOWID,
  '__module__' : 'protos.window_pb2'
  # @@protoc_insertion_point(class_scope:lnks.WindowId)
  })
_sym_db.RegisterMessage(WindowId)

MainWindowList = _reflection.GeneratedProtocolMessageType('MainWindowList', (_message.Message,), {
  'DESCRIPTOR' : _MAINWINDOWLIST,
  '__module__' : 'protos.window_pb2'
  # @@protoc_insertion_point(class_scope:lnks.MainWindowList)
  })
_sym_db.RegisterMessage(MainWindowList)

MainWindowProps = _reflection.GeneratedProtocolMessageType('MainWindowProps', (_message.Message,), {
  'DESCRIPTOR' : _MAINWINDOWPROPS,
  '__module__' : 'protos.window_pb2'
  # @@protoc_insertion_point(class_scope:lnks.MainWindowProps)
  })
_sym_db.RegisterMessage(MainWindowProps)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _WINDOWID._serialized_start=29
  _WINDOWID._serialized_end=117
  _MAINWINDOWLIST._serialized_start=119
  _MAINWINDOWLIST._serialized_end=180
  _MAINWINDOWPROPS._serialized_start=183
  _MAINWINDOWPROPS._serialized_end=678
# @@protoc_insertion_point(module_scope)