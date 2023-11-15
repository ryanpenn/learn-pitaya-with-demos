# Learn Pitaya with Demos

## Pitaya框架介绍

- [官方文档](./docs/Pitaya%E5%AE%98%E6%96%B9%E6%96%87%E6%A1%A3.pdf) (翻译)
- [框架分析](./docs/pitaya%E6%A1%86%E6%9E%B6%E5%88%86%E6%9E%90.pdf)

## 示例

1. [聊天](./01_chat/README.md)
2. [简单集群](./02_cluster/README.md) 区分前、后端服务器
3. [服务集群](./03_cluster_chat/README.md) 多组后端服务器
4. [采用protobuf序列化消息](./04_protobuf/README.md)
5. [采用msgpack序列化消息](./05_msgp/README.md)
6. [tadpole服务端](./06_tadpole_server/README.md) 分别通过 [nano](https://github.com/lonng/nano) 和 [pitaya](https://github.com/topfreegames/pitaya) 实现

## 引用

- [pitaya-cli](https://github.com/topfreegames/pitaya-cli)
- [pitaya-protos](https://github.com/topfreegames/pitaya-protos)

```bash
git submodule add https://github.com/topfreegames/pitaya-cli.git pitaya-cli
git submodule add https://github.com/topfreegames/pitaya-protos.git pitaya-protos
```

## Tips

- pull with submodules

```bash
# clone with submodules
git clone https://github.com/ryanpenn/learn-pitaya-with-demos.git --recursive

# (or) clone and update for the first time
git clone https://github.com/ryanpenn/learn-pitaya-with-demos.git
git submodule update --init --recursive

# update submodules (git version > 1.8.2)
git pull --recurse-submodules

# update submodules (git version > 1.7.3)
git submodule update --recursive --remote

```

- go module `SECURITY ERROR`

```bash
# 1. 修改 go.mod 升级 pitaya 版本: github.com/topfreegames/pitaya/v2 v2.8.0
# 2. 执行 tidy 命令
go mod tidy
```