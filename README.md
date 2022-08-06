<p align="center">
  <h1 align="center">tobeg</h1>
</p>
<p align="center">
  <a href="https://tobeg.asksowhat.cn">
    <img alt="asksowhat" src="https://bl6pap003files.storage.live.com/y4m3oxa18JTyLQHRtExGB4yqCUrzwydU7vfXLXbid-k7b708X8MzTgKSqisBAlceRISxF6ebyllOvN2xs2w3X3C81PqUw9VyidgzRMIpB4BGgfp-sPKVOAQ9oos4l9b_-OyYDztnibR4qcqWVyNCr2PAKNNWhLzXMdBt53mFmTf3Xl85jrJuHcwQTU5AzXireol?width=128&height=128&cropmode=none">
  </a>
</p>

<p align="center">
  叫花子要饭新姿势，简洁优雅，使用 gin 做为后端框架，前端使用 vue + element-ui 构建。
</p>
<p align="center">
<img src="https://img.shields.io/github/issues/asksowhat/tobeg?color=green">
<img src="https://img.shields.io/github/stars/asksowhat/tobeg?color=yellow">
<img src="https://img.shields.io/github/forks/asksowhat/tobeg?color=orange">
<img src="https://img.shields.io/github/license/asksowhat/tobeg?color=ff69b4">
<img src="https://img.shields.io/github/search/asksowhat/tobeg/main?color=blue">
<img src="https://img.shields.io/github/v/release/asksowhat/tobeg?color=blueviolet">
<img src="https://img.shields.io/github/languages/code-size/asksowhat/tobeg?color=critical">
</p>

![tobeg](https://bl6pap003files.storage.live.com/y4mBREf1EIU0h9IW47iLsjHx1r-fkuj8cF3mx5D76bwZA7CF-xNtkKTBhjeJfOzNjiTty3ZLU9bhcYGjHJyQZRS2h1FjSVVCVEyMFqiER1Lv4IFJrxkRJRpXRyfRd-BiGs-ErCFKRA6iaa_lTOsR66HHqfA5teq4mKUF0A44lbc3kmlBYPfBxZymB8LXHldlpDi?width=1069&height=663&cropmode=none)


# 项目说明

目前，该项目仅仅完成了基本功能，还存在着许多不足和 bug ，这些我都会继续完善。下一阶段，准备引入数据库，保存“赏钱”信息。

# 项目结构

```bash
.
├── api #api接口
│   ├── alipay.go
│   ├── flow.go
│   └── tmpl.go
├── cert #证书实体
│   └── cert.go
├── config #配置项实体
│   └── config.go
├── config.yml #配置文件
├── db #db数据库操作
│   └── sqlite3.go
├── docker-compose.yml
├── Dockerfile
├── flow.db #sqlite存储
├── forms
│   └── alipay.go
├── global #全局变量配置
│   └── global.go
├── go.mod
├── go.sum
├── initialize #初始化
│   ├── cert.go
│   ├── config.go
│   ├── logger.go
│   ├── router.go
│   └── sqlite3.go
├── LICENSE
├── main.go
├── middlerwares #中间件
│   └── cors.go
├── model #数据库表实体
│   └── flow.go
├── public #静态文件
│   ├── css
│   │   ├── elementui.index.css
│   │   └── fonts
│   │       ├── element-icons.ttf
│   │       └── element-icons.woff
│   ├── images
│   │   └── favicon.png
│   └── js
│       ├── elementui.index.js
│       ├── jquery.min.js
│       ├── qrcode.min.js
│       └── vue.min.js
├── README.md
├── router #路由配置
│   ├── alipay.go
│   ├── flow.go
│   └── tmpl.go
├── templates #模版引擎
│   └── index.html
└── utils #工具模块
    └── db.go
```

# 准备

目前，该项目仅仅对接了支付宝当面支付（支付扫码付款场景理论都支持）,需要完成以下几件事

- 申请开通支付宝当面付功能

- appId(应用ID)

- private_key(支付宝私钥)

- public_content_rsa2(支付宝公钥)

- app_public_content(应用公钥)

## 支付宝开通当面付

开通当面付，可参考官方说明进行申请。[当面付介绍及准入条件](https://opendocs.alipay.com/open/194/105072?ref=api)

## 进入应用管理

当开通完当面付功能之后，进入[应用管理](https://open.alipay.com/develop/manage)页面，点击网页/移动应用，可以看到已经有一个应用，点击进入详情。

红色箭头所指便是 **appId(应用ID)**

![](https://bl6pap003files.storage.live.com/y4mdzsosELSOl9LbdDywpWivb77IhvP4XEMiP1n2uk5TOTo-36oj6Qe7shzabb8wxi8t_k19aa7qen3D5OyuyHSRcqkYiKnosweL_5nvtWB14MR52pvMTP_2Be1VTocZF8ggvJJZU5ALqC94TnpB9dWAENpDdc25boVtQrQGXLsz2l_gjOhcJkPTcqW5xDwBPEt?width=794&height=273&cropmode=none)

## 生成密钥

参考支付宝官方文档 [生成并配置密钥](https://opendocs.alipay.com/common/02kdnc)，这里需要注意的是，生成密钥的流程相同，配置密钥不同的应用界面可能存在差异。

密钥生成之后，会得到两个文件 

- 应用公钥2048.txt

- 应用私钥2048.txt

应用私钥2048.txt 中的内容便是 **private_key(支付宝私钥)**

## 配置密钥

![](https://bl6pap003files.storage.live.com/y4myGbHKhhql36n05KBUBYsGUHqGc_bDvuAezGzKNXJnp217nddmNwaSsLn9RH4CKN_uVcyHstWetYNdi17tfrP8EDzAO-8fFAASHU5OSbALd_R7obcHrezRvPABQWXHzJZtYnSMHyqHxcfyvB8Y2CIoA6p1k-VwLnZCFh5jSdOgEJVDQA1uP1JzZd4HYXRAjdW?width=1589&height=473&cropmode=none)

将 应用公钥2048.txt 中的公钥填写在 圈2 位置，填写好之后你就得到了 **public_content_rsa2(支付宝公钥)** 和 **app_public_content(应用公钥)**。

如果上述表示有不明确之处，可参考支付宝官方文档 [当面付接入准备](https://opendocs.alipay.com/open/01csp3?ref=api)

至此，准备完毕。

# 配置

tobeg 支持用户在 config.yml 填写自定义配置，配置文件说明

base 模块下主要包含网站一些基础配置如网站图标，标题等等

字段 | 示例 | 说明
-|-|-
title | "tobeg" | 网站的title
url | "http://baidu.com" | 网站对外访问地址，关乎到支付回调判断，如站点访问地址是 https://google.com ，那么这个字段就要填写为 https://google.com 。
favicon | "favicon.png" | 站点图标，存储位置在/public/images/favicon.png，新图片可覆盖也可代替，但要注意这里要与实际图片文件名相同。
toSells | 字符串列表 | 在走马灯处显示
thank | "谢谢你, 好人" | 支付成功后的提示

alipay 主要是配置支付宝相关的信息

字段 | 示例 | 说明
-|-|-
appid | "2032423423" | 支付宝应用Id
private_key | "MIIEpAIBAAKCAQEAsw6yoLZ" | 支付宝密钥
public_content_rsa2 | "MIIBIjANBgdqhkiG9w" | 支付宝公钥
app_public_content | "MIIBIjANBgkqhkiG9wfBAQEFA" | 应用公钥

server 主要主要配置后端服务的相关信息

字段 | 示例 | 说明
-|-|-
port | 10020 | 配置后端服务的启动端口

完整示例

```yml
base:
  title: "tobeg"
  url: "http://baidu.com"
  favicon: "favicon.png"
  thank: "谢谢你, 好人"
  toSells:
    - "瞧一瞧、看一看、这里来个穷要饭哎"
    - "大娘好，大爷善，可怜可怜俺这个穷光蛋"
    - "给个摸，给口汤，祝恁长命又健康"
    - "行行好，行行秒，施舍的人呱呱叫"

alipay:
  appid: "2032423423"
  private_key: "MIIEpAIBAAKCAQEAsw6yoLZ"
  public_content_rsa2: "MIIBIjANBgdqhkiG9w"
  app_public_content: "MIIBIjANBgkqhkiG9wfBAQEFA"

server:
  port: 10020
```

需要注意的是，在 config.yml 配置文件中，你至少要两个地方 base.url 和 alipay 相关的证书信息，至此配置结束。

# 部署

该项目支持使用Docker-Compose进行部署，推荐自己构建最新镜像，也可以使用我构建的镜像，但可能不是最新的。

## 自己构建镜像

- 下载项目

```bash
git clone https://github.com/asksowhat/tobeg.git && cd tobeg
```

- 运行

相关配置完成之后

```bash
docker-compose up -d
```

这里会产生一个无用镜像，使用下面命令，便可删除

```bash
docker rmi $(docker images -f "dangling=true" -q)
```

## 使用我的镜像

将下面内容，覆盖 docker-compose.yml

```yml
version: "3"

services:
  tobeg:
    image: asksowhat/tobeg:latest
    container_name: tobeg
    restart: always
    ports:
      - 10020:10020
    volumes:
      - $PWD/config.yml:/config.yml
      - $PWD/flow.db:/flow.db
```

- 运行

```bash
docker-compose up -d
```

这里有一点要注意的是，如果你在 config.yml 更改了端口配置，需要将 docker-compose.yml 的ports配置，也要做相应的调整。

# 致谢（不分先后）

- [gin](https://github.com/gin-gonic/gin)

- [gopay](https://github.com/go-pay/gopay)

- [element-ui](https://github.com/ElemeFE/element)

- [qrcodejs](https://github.com/davidshimjs/qrcodejs)
