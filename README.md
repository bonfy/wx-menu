# wx-menu

快速创建微信公众号自定义菜单

## 关于使用

1. ~~ 打开 `menu.json` 配置菜单 ~~
2. 配置公众号的 app_id 与 app_key
3. 运行 `we-menu.py`

```cmd
$ pip3 install requests

$ export WX_APP_ID=APP_ID         # 输入公众号的 APP_ID
$ export WX_APP_SECRET=APP_SECRET # 输入公众号的 APP_SECRET

$ python3 we-menu.py
```

## TODO

- [ ] 解决中文支持，目前菜单名如果填中文还是会出错
