# 使用Gin搭建的IM聊天Demo

# ToDo  
- [x] 注册  
- [x] 登录  
- [x] 加载朋友  
- [x] 加载群组  
- [x] 加入群组  
- [x] 创建群组  
- [x] 添加朋友  
- [x] 私聊 
- [x] 群聊 
- [x] 发送表情
- [x] 发送图片
- [x] 发送语言
- [x] 发送视频文件
- [x] 多节点时发送信息通过UDP广播到各个节点然后发送消息
- [x] 用户状态
- [ ] 获取历史信息
- [ ] 优化消息推送问题，用户不在线，用户不在聊天页面等情况
- [x] 优化多节点部署的方案，不使用UDP广播，可能选择节点通过RPC通信
- [ ] 解决网络丢包问题，ACK响应机制，重传机制，数据包校验
- [ ] 优化重传机制可能导致消息重复的问题，唯一messageId
- [x] 强踢无响应心态连接
- [ ] 文件类使用对象存储


# 安装方法
## 1.下载项目
```
git clone git@github.com:PPPHUANG/ginchat.git
```
## 2.修改配置
修改配置文件/common/config.yaml
```
server:
  # Protocol (http or https)
  protocol: http
  ip: 127.0.0.1                            //多节点时不能使用本地回环地址
  port: 8080
  enforce_domain: true
  log_path: "./logger/ginchat/"             //日志路径
  # HTTP log output stdout
  log_stdout: false
  # file upload path
  attach_path: "./mnt"                      //文件上传缓存路径
  # server debug
  debug: false         
  nodes:                                    //节点IP
      - "192.168.0.100"
      - "192.168.0.101"                  
rpc:
  port: 8090                                //rpc监听端口
mysql:                                      //数据库配置自行修改
  ip: 127.0.0.1
  port: 3306
  user: root
  pwd: root
  db_name: chat
  show_sql: true
  max_open_conns : 2
redis:
  ip: 192.168.0.100                         //多节点时为redis-proxy入口
  port: 6666
```
## 3.编译运行
```
直接编译运行main.go入口文件，项目使用go mod，自行下载依赖库。
GOPROXY=https://goproxy.io go build //直接编译运行，注意可执行文件与asset、view文件夹需在同一级目录下。
```
## 4.访问
```
http://127.0.0.1:8080/user/login.shtml
注册接口没有前端对接，需要通过接口添加，然后通过页面登录即可。
```
# 接口介绍
```
POST /user/login                    登录
POST /user/register                 注册
POST /contact/loadcommunity/        加载群组
POST /contact/loadfriend            加载朋友
GET /contact/joincommunity          加入群组
GET /contact/createcommunity        创建群组
POST /contact/addfriend             添加朋友
GET /chat                           聊天创建长连接
POST /attach/upload                 附件上传
```
# 项目文件介绍
```
.
├── args                        定义的结构体
├── asset                       前端模板样式
│   ├── attach
│   ├── css
│   ├── fonts
│   ├── images
│   ├── js
│   └── plugins
├── client                      grpc客户端
├── common                      配置管理
├── db_conn                     数据库连接
├── httphandler                 Controller
│   ├── attach
│   ├── chat
│   ├── contact
│   └── user
├── httphandlerpack             logic
│   ├── contact
│   └── user
├── logger                      日志
│   └── ginchat
├── mnt                         文件缓存目录
├── model                       model
├── proto                       proto文件
├── router                      路由
├── service                     grpc服务
├── util                        工具类
├── version                     版本接口
└── view                        前端
    ├── chat
    └── user
```