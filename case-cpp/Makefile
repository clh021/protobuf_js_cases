TARGET = test

CXX = g++
CXXFLAGS = -O0 -g -std=c++11 -I../ # for Debug
#CXXFLAGS = -W -Wall -O3 -std=c++11 # for Release
Protobuf_INCLUDE_DIR = ~/.local/bin/
LDFLAGS =
RM = rm -rf
HEADERS = $(wildcard *.h)
SRCS = $(wildcard *.cc)
OBJS = $(SRCS:.cc=.o)

LIBS = protobuf
ifneq ($(LIBS),)
CXXFLAGS += $(shell pkg-config --cflags $(LIBS))
LDFLAGS += $(shell pkg-config --libs $(LIBS))
endif


.PHONY: all clean

all: gen build

gen:
	protoc --cpp_out=. *.proto

build: $(OBJS)
	$(CXX) $(OBJS) -o $@ $(LDFLAGS)

buildit:
	cmake -DProtobuf_PROTOC_EXECUTABLE=~/.local/bin/protoc

%.o: %.cpp $(HEADERS)
	$(CXX) -c $(CXXFLAGS) -o $@ $<

clean:
	$(RM) $(OBJS) build

