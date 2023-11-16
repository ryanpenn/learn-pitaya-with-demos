# 使用 grpc/protobuf 通信

> 遵循 [proto3](https://developers.google.com/protocol-buffers/docs/proto3) Language Guide

## 工具安装

1. 下载 [protobuf](https://github.com/protocolbuffers/protobuf/releases) 源码

2. 解压缩到目录 `protobuf-xx.x`

3. 进入目录，编译 `protobuf`

4. 拷贝到`$GOPATH`
```bash
mv protoc $GOPATH/bin
```

5. 验证
```bash
$ protoc --version
libprotoc x.x.x
```

6. 安装`protoc-gen-go`
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

7. 将`$GOPATH/bin`添加到`$PATH`环境变量中

8. 执行命令
```bash
# gen pb files
protoc --go_out=. -I=./proto  ./proto/*.proto
# (或)执行项目make指令
make protos
```

## 协议定义规范

> 遵守 **约定优于配置**（convention over configuration）的原则，规范通信协议的定义。

- **路由**由三级结构组成，第1层为**服务器**，第2层为**处理器**，第3层为**api**
  如`chat.handler.say`,说明这个服务由`chat`这台服务器提供，由`handler`这个处理器在处理，里面的方法`say`来接收这台路由对应的消息
- 路由全部由**小写**字符组成,注册`strings.ToLower`作为`pitaya.Register()`方法的第三个参数
- proto的命名要求：协议以`Req`、`Resp`、`Input`、`Sync`、`Arg`、`Reply`开头
- 消息分为3种类型，即`Call`、`Notify`、`Push`
  > - `Call`类型的：对应的proto数据结构为`Req` + `Resp`，普通消息，有返回值，发送后等待返回值。
  > - `Notify`消息，对应的proto数据结构为`Input`，通知消息，无返回值，发送后立即返回。
  > - `Push`消息，对应的proto数据结构为`Sync`，推送消息，由服务端发起。
- 服务器之间的通信通过 `RPC` 实现，RPC同样遵循**路由**约定，例如：`app.RPC(ctx, "web.remote.playerinfo", &ArgPlayerInfo, &ReplyPlayerInfo)`，RPC参数类型为：
  > - 参数以 `Arg` 开头，如：`ArgPlayerInfo`
  > - 返回值以 `Reply` 开头，如：`ReplyPlayerInfo`

## proto3

### 1.定义消息类型

使用# `proto3` 需要在 `.proto` 文件的最开始添加声明： `syntax = "proto3"`; 否则将按照 `proto2` 来编译。

```proto
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

#### 1.1. 字段类型

见下面的标量类型，枚举等。

#### 1.2. 分配字段编号

消息中定义的字段每个都要有一个 **唯一的编号** ，在编码的时候会用来确定字段，一旦类型定义之后不可以更改。

注意: `1 ~ 15` 之间的字段编号需要一个字节来编码（包括字段编号和字段类型）， `16 ~ 2047` 需要 2 个字节。因此，应该为经常使用的字段预留 1-15 数字。

字段编号的最小值是 `1`，最大是 `2^29 - 1` 或者 `536,870,911`。 `19000 ~ 19999` (`FieldDescriptor::kFirstReservedNumber ~ FieldDescriptor::kLastReservedNumber`) 之间的数字预留给 Protocol Buffers 实现的。

#### 1.3. 字段规则

字段可以是：
- 单数：单个值
- repeated 对应的是数组
  在 proto3 中，标量数字类型的 `repeated` 字段，默认使用 `packed` 编码。

#### 1.4. 消息类型

一个 `.proto` 文件中，可以有多个 `message` 定义。

#### 1.5. 注释

在 `.proto` 中，注释使用 C/C++ 风格的 `//` 和 `/* ... */`

#### 1.6. 保留字段

message 中的字段可能被删除或者注释，一旦未来之前的字段编号被复用。可能会导致严重的问题，比如数据损坏和一些隐藏的 bug 等。 所以要确保已经废弃的字段编号不会被再次使用。一种解决办法是显式的用 reserved 指定已经被删除的字段编号。如果将来有人使用了，编译器会做出提示。

```proto
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

甚至可以使用 `max` 关键字来表示最大字段编号值。比如 `40 to max` 表示 `40` 到 `max` 之间的全部保留。注意，在 `reserved` 值中不可混用字段编码和字段名。

### 2. 标量类型

| `.proto` Type | Notes | C++ Type | Python Type | Go Type |
| --- | --- | --- | --- | --- |
| double |   | double | float | float64 |
| float |   | float | float | float32 |
| int32 | 可变长度编码，负数编码效率低，如果值可能是负的，请用 sint32 代替 | int32 | int | int32 |
| int64 | 可变长度编码，负数编码效率低，如果值可能是负的，请用 sint64 代替 | int64 | int/long | int64 |
| uint32 | 可变长度编码 | uint32 | int/long | uint32 |
| uint64 | 可变长度编码 | uint64 | int/long | uint64 |
| sint32 | 可变长度编码，有符号的整型值 | int32 | int | int32 |
| sint64 | 可变长度编码，有符号的整型值 | int64 | int/long | int64 |
| fixed32 | 总是 4 字节，如果值大于 2^28 比 uint32 更高效 | uint32 | int/long | uint32 |
| fixed64 | 总是 8 字节，如果值大于 2^56 比 uint64 更高效 | uint64 | int/long | uint64 |
| sfixed32 | 总是 4 字节 | int32 | int | int32 |
| sfixed64 | 总是 8 字节 | int64 | int/long | int64 |
| bool |   | bool | bool | bool |
| string | 字符串必须是 UTF-8 编码或者 7-bit 的 ASCII 文本，不能超过 2^32 | string | str/unicode | string |
| bytes | 任意长度不超过 2^32 的字节序列 | string | str | []byte |

### 3. 默认值

消息解析时，如果不包含特定的字段，会用该字段的默认值代替。不同的类型，默认值不同：

- 字符串，默认为空字符串
- 字节，默认为空字节
- 布尔值，默认为 `false`
- 数值类型，默认为 `0`
- 枚举类型，默认是第一个定义的枚举项，必须为 `0`
- 对于 `message` 字段，默认值取决于具体的编程语言
- `repeated` 字段的默认值为空（通常为对应编程语言的空列表）。

对于标量类型，一旦消息解析之后，无法判断该值是默认值还是未被设置的值。在设计时应该明确这一点，避免产生与设想不符的行为。

### 4. 枚举

如下定义：

```proto
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {
	UNIVERSAL = 0;
	WEB = 1;
	IMAGES = 2;
	LOCAL = 3;
	NEWS = 4;
	PRODUCTS = 5;
	VIDEO = 6;
  }
  Corpus corpus = 4;
}
```

每个枚举定义 **必须** 将第一个元素的值设置为常量 `0` 。因为：

- 必须要有一个零值，来保证数值类型有默认值
- 零值必须是第一个元素，为了跟 proto2 兼容
  通过设置 `allow_alias=true` ，你可以定义相同的值分配给不同的枚举常量，否则出现相同的值编译不通过：

```proto
message MyMessage1 {
  enum EnumAllowingAlias {
	option allow_alias = true;
	UNKNOWN = 0;
	STARTED = 1;
	RUNNING = 1;
  }
}
```

枚举的值必须在 32-bit 整型范围内，不建议使用复数，编码效率不高。你可以在消息的内部、外部定义枚举。还可以使用 `MessageType.EnumType` 添加声明，将消息中的枚举类型公开。

### 5. 使用其它消息类型

你可以使用已经定义的消息类型作为另外一个消息的字段类型。如：

```proto
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

#### 5.1. 导入定义

如果消息定义在不同的 `.proto` 文件中，你在使用 `import` 关键字来导入，如：

```proto
import "myproject/other_protos.proto";
```

默认情况下只能使用直接导入的 `.proto` 文件定义，比如 B import A，C import B，这种情况下 C 是没有 import A。 如果想要实现 C import A 的效果，则需要在 B import A 的时候指定 public 。如：

```proto
import public "a.proto"
```

protobuf 编译器搜索导入的文件是基于 `-I` 或者 `--proto_path` 指定的路径的。一般将项目的根目录作为 `--proto_path` 的值。

### 6. 内嵌类型

你可以在消息类型中定义和使用消息类型，如下：

```proto
message SearchResponse {
  message Result {
	string url = 1;
	string title = 2;
	repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```
如果想要在父消息外面使用内部定义的消息，需要加一层引用，如 `Parent.Type` ：

```proto
message SomeOtherMessage {
  SearchResponse.Result result = 1;
}
```

消息定义可以内嵌很多层。

### 7. 更新消息类型

如果一个已经存在的消息类型不再能满足需求，比如，添加额外的字段等。在不破坏现有的消息类型更新非常简单，但是要遵守一下规则：

- 不要更改现有的任何字段的字段编号。
- 添加新字段时，老的消息格式序列化仍旧可以被新的解析，新的字段会以默认值出现。同样新的消息格式序列化也可以被旧的解析，但是会忽略新字段。 兼容性
- 删除字段时，要保证新的字段编号不与删除的相同。重命名该字段，或者添加 `OBSOLETE_` 前缀，或者使用 `reserved` 关键字。 以确保将来的用户不会复用之前的字段编号。
- `int32` `uint32` `int64` `uint64` 和 `bool` 都是兼容的, 也就是说你可以在它们之间修改字段的类型，而不会破坏向前或者向后兼容性。 如果解析中的字段类型不同，会发生自动类型转换。如果字节数变少了，会自动截断。
- `sint32` 和 `sint64` 相互兼容，但与其它类型不兼容。
- `string` 和 `bytes` 只要是有效的 `UTF-8` ，相互兼容。
- 如果字节包含消息的编码版本，则 `bytes` 和内嵌消息兼容。
- `fixed32` 跟 `sfixed32` 兼容， `fixed64` 和 `sfixed64` 兼容。
- `enum` 和 `int32` `uint32` `int64` `uint64` 兼容（如果值不同，自动截断）。但要注意，反序列化消息时，客户端代码可能会以不同的方式对待它们： 比如，无法识别的 proto3 enum 类型会保留在消息中，在反序列化消息时如何表达取决于具体的语言。 int 字段只是保留其值。
- Changing a single value into a member of a new oneof is safe and binary compatible. Moving multiple fields into a new oneof may be safe if you are sure that no code sets more than one at a time. Moving any fields into an existing oneof is not safe.

### 8. 未知的字段

未知(Unknown)的字段表示在序列化话数据时，解析器无法识别的字段。比如说，用旧的二进制数据使用新的二进制解析时，新的字段变成旧二进制数据中的未知字段。

原本，proto3 消息在解析过程中会始终丢弃未知字段，但是 3.5 版本之后，我们重新引入和保留未知字段以匹配 proto2 行为的功能。 在 3.5 或者更高版本中，未知字段将在解析期间保留并包含在序列化输出中。

### 9. Any

`Any` 消息类型可以作为嵌入类型，而无须定义。 `Any` 以 `bytes` 为单位，包含任何序列化消息，扮演着该消息类型的全局唯一标识符的 URL。 要使用 `Any` 类型，需要先导入 `google/protobuf/any.proto` 。

```proto
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

给定类型默认的 URL 是 `type.googleapis.com/packagename.messagename` 。

不同语言的实现会在运行时库帮助下以类型安全的方式 `pack` 和 `unpack Any` 值。比如，在 Java 中，`Any` 类型用特殊的 `pack()` 和 `unpack()` 存取器， C++ 是 `PackForm()` 和 `UnpackTo()` 方法。

**当前用于 Any 类型的运行时库正在开发中。**

### 10. Oneof

如果一个消息有很多的字段，但是在同一个时间内只会被设置其中的一个，你可以使用 `oneof` 特性来实现它，以节省内存。

Oneof 除了所有的字段共享一块内存之外的行为和其它的普通字段一样，最多可以同时设置一个字段，设置 `oneof` 中的任何一个字段会清空其它的成员。 你可以使用特殊的 `case()` 或者 `WhichOneof()` 方法来判断设置了哪个值，具体依赖于你选择语言的实现。

#### 10.1. 使用 Oneof

在 `.proto` 文件中使用 `oneof` 关键字来定义：

```proto
message SampleMessage {
  oneof test_oneof {
	string name = 4;
	SubMessage sub_message = 9;
  }
}
```

然后将 `oneof` 字段添加到定义，但是不可以定义成 `repeated` 字段。

在生成的代码中，`oneof` 字段具有相同的 `getters` 和 `setters`。额外还有一种的特殊的方法来判断哪个字段被设置了。 具体看语言的 API 参考文档：https://developers.google.com/protocol-buffers/docs/reference/overview。

#### 10.2. Oneof 特性

- 设置一个 `oneof` 字段会自动清空其它的。多次设置，只有最后一次生效。
- 如果解析器看到相同 `oneof` 的多个成员，只有最后一个看到的成员被解析。
- 不可以使用 `repeated` 。
- 反射 APIs 对 oneof 字段有效。
- If you set a oneof field to the default value (such as setting an int32 oneof field to 0), the "case" of that oneof field will be set, and the value will be serialized on the wire.
- 如果你使用 C++，确保你的代码不会导致内存崩溃。如下代码会导致崩溃，因为 `set_name()` 方法已经删除了 `sub_message` :

```cpp
SampleMessage message;
SubMessage* sub_message = message.mutable_sub_message();
message.set_name("name");      // Will delete sub_message
sub_message->set_...            // Crashes here
```

- 还是 C++，如果你 `Swap()` 有 oneof 的两个消息，每个消息都会最终以对方为准：下面的例子中 `msg1` 会有一个 `sub_message` 而 `msg2` 会有一个 `name` 。

```cpp
SampleMessage msg1;
msg1.set_name("name");
SampleMessage msg2;
msg2.mutable_sub_message();
msg1.swap(&msg2);
CHECK(msg1.has_sub_message());
CHECK(msg2.has_name());
```

#### 10.3. 向后兼容问题

> 这玩意一般用不到，而且我感觉使用的时候问题比带来的好处要多。

### 11. Maps

如果想要在数据定义中创建关联映射，protobuf 提供了方便的快捷语法：

```proto
map<key_type, value_type> map_field = N;
```

`key_type` 可以使用任何整型和字符串类型（也就是说除了标量类型浮点型和 `bytes` 之外的）。枚举不是一个有效的 `key_type` ， `value_type` 可以是除了 `map` 以外的所有类型。

使用注意：

- Map 字段不可以是 `repeated` 。 字段不可以，不是 `value` 不可以
- Map 的 `key` 和 `value` 顺序是不确定的，因为你不可以依赖与 `map` 中的元素的特定顺序。
- 当生成 `.proto` 的文本格式时，`maps` 是按 `key` 排序，数值类型的 `key` 是按照数值排序。
- 编码的时候重复的键会使用最后一个看到的值。解码的时候，出现相同的 `keys` 会解码失败。
- 如果提供了一个 `key` 但是没有值，序列化行为取决于语言。C++，Java，Python 会使用类型的默认值，其它的语言什么都不做。

#### 11.1. 向后兼容

编码的时候 `map` 语法等价于下面这样，因此 protobuf 实现就算不支持 `map` 也可以处理你的数据：

```proto
message MapFieldEntry {
  key_type key = 1;
  value_type value = 2;
}

repeated MapFieldEntry map_field = N;
```

### 12. Package

你可以在 `.proto` 文件中添加一个可选的 `package` 指示符用来放置协议消息类型质检的命名冲突。

```proto
package foo.bar;
message Open { ... }
```

然后你可以在另外的消息定义中使用：

```proto
message Foo {
  ...
  foo.bar.Open open = 1;
  ...
}
```

`package` 不同的语言生成的代码不同：

- **C++** 等价于 `namespace`
- **Python** 会被路忽略，因为 Python 的模块是根据它在文件系统中的位置进行组织的。
- **Go** 当做 Go 的 `package` 名称，除非你显式的提供了 `option go_package` 。

### 13. 定义服务(Defining Services)

如果你要与 RPC 系统一起使用你的消息类型，你可以在 `.proto` 文件中定义一个 RPC 服务接口，protobuf 编译器会根据你选择的语言自动生成服务接口代码和存根。

如下，一个 RPC 服务请求是 `SearchRequest` 消息，返回是 `SearchResponse` 消息，你可以这样定义：

```proto
service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}
```
与 protobuf 一起使用的直接使用 RPC 系统是 gRPC：谷歌开发的一个语言和平台无关的开源的 RPC 系统。gRPC 和 protobuf 是最佳搭档， 使用特殊的编译插件可以直接生成相关的 RPC 代码。

在 [这里](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md) 列出了很多第三方的 protobuf 插件。

### 14. JSON 映射

Proto3 支持 JSON 编码规范，这使得在系统之间共享数据更加容易。

https://developers.google.com/protocol-buffers/docs/proto3#json

### 15. 选项(Options)

`.proto` 文件中的各种声明可以使用许多的 `options` 来注释。`Options` 不会改变声明的整体定义，但可能会影响在特定上下文中处理声明的方式。 可用的 `Options` 的完整定义在 `google/protobuf/descriptor.proto` 中。

- 一些 `options` 是在文件级别定义，这意味着它们在顶层范围定义，而不是在任何的消息、枚举或者服务定义中。
- 一些 `options` 是在消息级别定义，意味着它们应该写在消息内部定义中。
- 一些 `options` 是字段级别的选项，它们应该写在字段定义上。

`Options` 可以写在枚举类型、枚举值、oneof 字段，服务类型和服务方法；但是目前没有有用的 `options` 给它们用。

下面是常用的选项：

- `java_package`
- `java_multiple_files`
- `java_outer_classname`
- `optimize_for`
- `cc_enable_arenas`
- `objc_class_prefix`
- `deprecated` 字段选项，如果设置为 `true` ，表示该字段的已经不推荐使用，不应该在新的代码中使用。在大多数语言中，都不会有实际的效果。 Java 中，会添加 `@Deprecated` 注解。如果一个字段没被用过，但是不想新用户使用它，可以使用 `reserved` 声明，而不是 `deprecated`

```proto
int32 old_field = 6 [deprecated = true];
```

## 使用`Makefile`生成协议

```makefile
ifeq ($(OS),Windows_NT)
RM = rmdir /s /q
else
RM = rm -rf
endif

# 生成协议
.PHONY: protos
protos:
	@echo "step1: clean pd folder"
	$(RM) pb
	@mkdir pb
	@echo "step2: gen *.pb.go files"
	@protoc --go_out=. -I=./proto  ./proto/*.proto
	@echo "step2: done!"
```
