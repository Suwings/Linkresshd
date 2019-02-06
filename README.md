Linkresshd
------
使用 Go 语言编写的 ssh 服务端程序，可以只通过一个二进制文件运行来实现 ssh 服务，一个小型轻量级的 ssh 服务只需要一个文件即可完成，并且还拥有高度的可自定义化，方便不同场景使用。

你可以通过 ssh 协议管理各种各样的程序，不一定非得链接 shell 程序，比如管理一个 docker 容器也可以，而不会给予使用者太高的权限。

<br />

使用方法
------

- 下载[最新发行版](https://github.com/Suwings/Linkresshd/releases/latest)的二进制程序直接运行。
- 下载或克隆源代码自行编译，可使用 go build 编译。

感谢使用！


<br />

配置文件
------
在 `config/config.json` 中 (此文件必须存在)，您可以自行修改。
```
"name": "root",         登录的用户名
"password": "toor",     登录的密码
"command": "/bin/bash", 验证后使用者需要运行程序
"port": 2222            监听的端口
```
<br />


开源协议
------
使用 `MIT License` 开源协议。

