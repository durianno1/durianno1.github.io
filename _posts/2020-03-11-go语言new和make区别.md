---
layout: post
title: 'golang中的make和new'
subtitle: 'make和new创建变量的区别和特点'
date: 2020-03-11
categories: 技术
cover: 'http://on2171g4d.bkt.clouddn.com/jekyll-theme-h2o-postcover.jpg'
tags:  golang
---
首先二者都是在堆中分配内存的关键字
## new

new返回一个指向制定类型的指针，并且new在go中创建的变量都是zero-type的，就是该类型的零值，比如int类型为&0，string类型为&“”，struct类型为{},但是也有例外，比如bool类型则为false。一般map、slice、channel类型不使用new来创建，如果使用了new来创建，则会返回一个空类型，因为这三种数据在底层结构体上都有引用类型的字段，必须初始化。但是强制使用new来创建slice是可以的，而map不行.


### make
make是专门创建map、slice、channel时候使用的关键字，它返回的是一个初始化后的实体而非指针。
