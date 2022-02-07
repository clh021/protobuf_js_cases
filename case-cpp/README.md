# case-cpp

> 本测试案例，还未完成

## 步骤

```bash
protoc --cpp_out=. *.proto
#make build # make build -j 16
cmake -DProtobuf_PROTOC_EXECUTABLE=~/.local/bin/protoc
```
