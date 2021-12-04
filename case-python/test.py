#coding=utf-8
# file: test.py

import protos.window_pb2 as window

winList = window.MainWindowList()
winList.ok = 1
winId1 = winList.windows.add()
# 无论 chromium_id 是否包含下划线，最终的字节流编码都是一致的。
winId1.chromium_id = 1

print(winList)

data = winList.SerializeToString()
print(data)


winList = window.MainWindowList2()
winList.ok = 1
winId1 = winList.windows.add()
winId1.chromium_id = 1

print(winList)

data = winList.SerializeToString()
print(data)


winList = window.MainWindowList3()
winList.ok = 1
winList.chromium_ids.append(1)

print(winList)

data = winList.SerializeToString()
print(data)




winList = window.MainWindowList4()
winList.chromium_ids.append(1)
winList.chromium_ids.append(1)

print(winList)

data = winList.SerializeToString()
print(data)



winList = window.MainWindowList5()
winList.chromium_ids = 1
winList.chromium_id = 1

print(winList)

data = winList.SerializeToString()
print(data)
