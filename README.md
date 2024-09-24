# NoAuth

2024.9.24
有空改下，修下小bug。。。
有空改下，最开始的目的是自己总审代码所以就简单辅助下白盒啥的。。。
有空改下，把bypass规则与判断规则完善下，再加下代理功能啥的



## \[简介]

审Java代码的时候，经常会遇到接口因鉴权问题被组合拳getshell，例如泛微和海康安防这些系统。为了平时审代码与绕鉴权时节省点时间，花了点时间分析总结了下网上所有与Java鉴权有关的问题，写了这工具，主要用于动态生成可能用于绕过的payload进行fuzz测试。默认发送Get、Post-Form-datas、POST-json 数据包。工具没跑出来大概率也就不存在了，估计就得深入代码了。

## \[功能]

![](image/image_KtBTT0LNqX.png)

```bash


Usage:  [-unat] [-u url] [-n interface without authentication] [-a interface An interface that requires authentication] [-t thread] [-debug choose start debug] [-h help]

Options:
  -a string
        An interface that requires authentication, such as /admin/adduser
  -debug int
        choose start debug, such -debug 1
  -h    This help
  -n string
        An interface without authentication, such as /login
  -t int
        Thread Num (default 8)
  -u string
        A target url(Please add http or https)


```

## \[使用方法]

NoAuth -n 不需要鉴权的接口地址(如/login、/register、/index.jsp、/index.html、/js/background.js等) -a 需要鉴权的接口地址 -u url地址

例：NoAuth -n /login -a /admin/adduser -u [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")&#x20;

如图，成功利用[http://localhost:8080/Firstfilter/](http://localhost:8080/Firstfilter/ "http://localhost:8080/Firstfilter/")..;/FirstServlet  绕过鉴权

![](image/image_rvuUWUCy2w.png)

![](image/image_jbIueinWPF.png)

NoAuth  -n /login -a /admin/adduser -u [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/") -debug 1 ，添加 -debug 1 参数可查看所有请求 ：

![](image/image_rOjvXpoojL.png)
