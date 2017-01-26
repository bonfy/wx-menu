# wx-menu

快速创建微信公众号自定义菜单

## 关于使用

### 配置环境变量

配置环境变量 `WX_APP_ID` 和 `WX_APP_SECRET`

```cmd
$ export WX_APP_ID=APP_ID         # 输入公众号的 APP_ID
$ export WX_APP_SECRET=APP_SECRET # 输入公众号的 APP_SECRET
```

### Go

Go 支持 中文菜单 & 英文菜单

```
$ go get -u github.com/bitly/go-simplejson
$ go get -u github.com/levigross/grequests

$ go run wx-menu.go
```

### Python

注意： python 目前支持 英文菜单， 有中文字符会报错

```cmd
$ pip3 install requests
$ python3 we-menu.py
```

## TODO

- [ ] 解决python中文支持，目前菜单名如果填中文还是会出错
