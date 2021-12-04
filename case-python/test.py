#coding=utf-8
# file: test.py

import protos.window_pb2 as window
# winId1 = window.WindowId()
# winId1.chromiumId = 1
winList = window.MainWindowList()
winList.ok = 1
winId1 = winList.windows.add()
# 无论 chromium_id 是否包含下划线，最终的字节流编码都是一致的。
winId1.chromium_id = 1

print(winList)

data = winList.SerializeToString()
print(data)
