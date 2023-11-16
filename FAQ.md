# FAQ

## pull with submodules

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

## go module `SECURITY ERROR`

```bash
# 1. 修改 go.mod 升级 pitaya 版本: github.com/topfreegames/pitaya/v2 v2.8.0
# 2. 执行 tidy 命令
go mod tidy
```

## git submodules

```bash
# 作为子模块添加到项目
git submodule add https://github.com/topfreegames/pitaya-cli.git pitaya-cli
git submodule add https://github.com/topfreegames/pitaya-protos.git pitaya-protos

# 带子模块一起克隆项目
git clone https://github.com/ryanpenn/learn-pitaya-with-demos.git --recursive

# 克隆项目后，再更新子模块
git clone https://github.com/ryanpenn/learn-pitaya-with-demos.git
git submodule update --init --recursive

# 更新子模块
 git pull --recurse-submodules
```

## 执行 `git submodule update --init --recursive` 报错

```bash
fatal: remote error: upload-pack: not our ref fc7223ca00124e8f5b5b354457379071e2fd091b
Fetched in submodule path 'proto', but it did not contain fc7223ca00124e8f5b5b354457379071e2fd091b. Direct fetching of that commit failed.
```

解决方法:
```bash
cd ./proto
git reset --hard origin/master
cd ../
git clean -n
git add ./proto
git commit -m "reset submodule"
git submodule update --init --recursive
```

