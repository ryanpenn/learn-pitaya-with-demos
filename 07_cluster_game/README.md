# game server cluster

## Layouts

```bash
cluster_game
├─resources             资源目录: 服务器配置,配置表等
├─build                 项目构建脚本: dockerfile,deploy.sh等
├─cmd                   项目入口
├─chat                  聊天服务器
├─game                  游戏服务器
├─gate                  网关服务器
├─login                 登录服务器
├─modules               公共模块
├─pkg                   公共包
├─protos                通信协议
├─settings              配置表
├─main.go               入口
├─go.mod                项目模块
└─README.md             项目说明
```
