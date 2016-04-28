# demo  [![GoDoc](https://godoc.org/github.com/lessgo/lessgo?status.svg)](https://godoc.org/github.com/lessgo/lessgo) [![GitHub release](https://img.shields.io/github/release/lessgo/lessgo.svg)](https://github.com/lessgo/lessgo/releases)
A demo of Lessgo.


![Lessgo Favicon](https://github.com/lessgo/lessgo/raw/master/doc/favicon.png)

##概述
Lessgo 是一款Go 语言开发的简单、稳定、高效、灵活的 web开发框架。它的项目组织形式经过精心设计，实现前后端分离、系统与业务分离，完美兼容MVC与MVVC等多种开发模式，非常利于企业级应用与API接口的开发。当然，最值得关注的是它突破性地支持了运行时路由重建，开发者可在Admin后台轻松实现启用/禁用模块与操作，添加/移除中间件等功能！同时，它推荐以HandlerFunc与MiddlewareFunc为基础的函数式编程，也令开发变得更加灵活富有趣味性。
此外它也博采众长，核心架构基于[echo v2](https://github.com/labstack/echo)并增强优化，数据库引擎内置为[xorm](https://github.com/go-xorm/xorm)，模板引擎内置为[pongo2](https://github.com/flosch/pongo2)，其他某些功能模块改写自[beego](https://github.com/astaxie/beego)以及其他优秀开源项目。（在此感谢这些优秀的开源项目）

##适用场景
- 网站
- web应用
- Restful API服务应用
- 企业应用

##项目架构
![Lessgo Web Framework](https://github.com/lessgo/lessgo/raw/master/doc/LessgoWebFramework.jpg)

##框架构成
- 核心框架：[lessgo](https://github.com/lessgo/lessgo)
- 框架扩展：[lessgoext](https://github.com/lessgo/lessgoext)
- 项目Demo：[demo](https://github.com/lessgo/demo)

##框架下载

```sh
go get -u github.com/lessgo/lessgo
go get -u github.com/lessgo/lessgoext
```

##系统文档
- [综述](doc/Introduction.md)
- [安装部署](doc/Install.md)
- [开始lessgo之旅](doc/Develop01.md)


##贡献者名单
贡献者                          |贡献概要
--------------------------------|--------------------------------------------------
[henrylee2cn](https://github.com/henrylee2cn)|第一作者 (主要代码实现者) 
[changyu72](https://github.com/changyu72)|第二作者 (主要架构设计者) 

##开源协议
Lessgo 项目采用商业应用友好的 [MIT](https://github.com/lessgo/lessgo/raw/master/doc/LICENSE) 协议发布。
 
##项目联系
* 官方QQ群：Go-Web 编程 42730308    [![Go-Web 编程群](http://pub.idqqimg.com/wpa/images/group.png)](http://jq.qq.com/?_wv=1027&k=fzi4p1)
![Lessgo Server](https://github.com/lessgo/lessgo/raw/master/doc/server.jpg)
