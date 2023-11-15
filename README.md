# Learn Pitaya with Demos

## Pitaya框架介绍

- [官方文档](./docs/Pitaya%E5%AE%98%E6%96%B9%E6%96%87%E6%A1%A3.pdf) (翻译)
- [框架分析](./docs/pitaya%E6%A1%86%E6%9E%B6%E5%88%86%E6%9E%90.pdf)

## 示例

1. [聊天](./01_chat/README.md)


## 引用

- [pitaya-cli](https://github.com/topfreegames/pitaya-cli)
- [pitaya-protos](https://github.com/topfreegames/pitaya-protos)

```bash
git submodule add https://github.com/topfreegames/pitaya-cli.git pitaya-cli
git submodule add https://github.com/topfreegames/pitaya-protos.git pitaya-protos
```

## Tips

- go module `SECURITY ERROR`

```bash
# 1. 修改 go.mod 升级 pitaya 版本: github.com/topfreegames/pitaya/v2 v2.8.0
# 2. 执行 tidy 命令
go mod tidy
```