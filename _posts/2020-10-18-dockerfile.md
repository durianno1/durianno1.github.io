---
layout: post
title: 'dockerfile'
subtitle: 'dockerfile基础'
date: 2020-10-18
categories: 技术
cover: 'http://on2171g4d.bkt.clouddn.com/jekyll-theme-h2o-postcover.jpg'
tags:  docker
---

### dockerflie ###

通过dockerfile可以生成镜像文件

构建步骤：

1. 编写dockerfile文件
2. docker build -f dockerfile文件名 -t 目标镜像名:tag
3. docker run 创建容器 
4. Docker push发布镜像(docker hub,阿里云)





### 指令 ###

```shell
FROM					# 指定基础镜像 
MAINTAINAER		# 表明镜像维护者
RUN 					# 执行命令并创建新的镜像层通常用于安装软件包,直接对镜像作用
CMD 					# 设置容器启动后默认执行的命令和参数，可以被docker run的参数覆盖，与RUN不同的是CMD不作用于镜像，而是创建容器时执行，作用于实体容器,配置多个的话只有最后的会生效
ENTRYPOINT		# 与CMD类似不过不会被覆盖始终会执行
COPY 					# 将本地文件复制到容器中，如果复制的是一个目录，则不复制目录，只复制目录中文件
WORKDIR				# 工作目录,进入到容器后的目录
ADD 					# 与COPY类似，功能更多更灵活
EXPOSE 				# 监听的端口，可在docker run 时候用-p覆盖
VOLUME 				# 配置数据卷挂载
ONBUILD 			# 被继承时候触发
ENV						# 设置环境变量
```



### 写一个centos的dockerfile ###

```shell
FROM centos
MAINTAINER Severus<durianno1@aliyun.com>

ENV MYPATH /usr/local
WORKDIR $MYPATH

RUN yum -y install vim
RUN yum -y install net-tools

EXPOSE 80

CMD echo $MYPATH
CMD echo "------end-------"
CMD /bin/bash
```

> docker build -f dockerfile -t mycentos:0.1 . 

