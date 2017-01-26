# wx-menu

快速创建微信公众号自定义菜单

## 关于使用

### 配置环境变量

配置环境变量 `WX_APP_ID` 和 `WX_APP_SECRET`

```cmd
$ export WX_APP_ID=APP_ID         # 输入公众号的 APP_ID
$ export WX_APP_SECRET=APP_SECRET # 输入公众号的 APP_SECRET
```

### Python

> 注意： python 在 菜单中文配置方面有问题，输入英文则没有问题

```cmd
$ pip3 install requests
$ python3 we-menu.py
```


### Go

```
$ go get -u github.com/bitly/go-simplejson
$ go get -u github.com/levigross/grequests

$ go run wx-menu.go
```

## TODO

- [ ] 解决python中文支持，目前菜单名如果填中文还是会出错
