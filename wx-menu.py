# coding: utf-8

# Author BONFY<i@bonfy.im>
# createdAt 2017.01.20

import requests
import json
import os

# 使用之前需要修改下面的app_id 和 app_secret
app_id = os.environ.get('WX_APP_ID', '')
app_secret = os.environ.get('WX_APP_SECRET','')

def get_access(app_id, app_secret):
    """
    获取 access_token

        http请求方式: GET
        https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
        正常： {"access_token":"ACCESS_TOKEN","expires_in":7200}
        错误： {"errcode":40013,"errmsg":"invalid appid"}

    :param str app_id:      APP_ID
    :param str app_secret:  APP_SECRET
    :return str token: weixin access token
    """
    access_url = 'https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid={app_id}&secret={app_secret}'.format(app_id=app_id, app_secret=app_secret) # noqa
    r = requests.get(access_url)
    assert r.status_code == 200

    res = r.json()
    return res.get('access_token', None)

def create_menu(access_token, filename='menu.json'):
    """
    创建 menu

        http请求方式：POST
        （请使用https协议） https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN

    :param str access_token: ACCESS_TOKEN
    :param str filename:     default `menu.json`
    :rtype: None
    """
    # get menu json
    with open(filename, 'r') as f:
        data = json.load(f)
    # post to menu_url
    menu_url = 'https://api.weixin.qq.com/cgi-bin/menu/create?access_token={access_token}'.format(access_token=access_token)
    r = requests.post(menu_url, json=data)
    assert r.status_code == 200
    assert r.json().get('errcode') == 0

def get_menu(access_token):
    """
    获取 创建的 menu

        http请求方式：GET
        https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN

    :param str access_token: ACCESS_TOKEN
    :rtype: json
    """
    menu_url = 'https://api.weixin.qq.com/cgi-bin/menu/get?access_token={access_token}'.format(access_token=access_token) # noqa
    r = requests.get(menu_url)
    assert r.status_code == 200
    return r.json()

def delete_menu(access_token):
    """
    删除菜单 请谨慎使用

        http请求方式：GET
        https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN

    :param str access_token: ACCESS_TOKEN
    :rtype: json
    """
    menu_url = 'https://api.weixin.qq.com/cgi-bin/menu/delete?access_token={access_token}'.format(access_token=access_token)
    r = requests.get(menu_url)
    assert r.status_code == 200
    return r.json()

if __name__ == '__main__':
    access_token = get_access(app_id, app_secret)
    print(access_token)
    create_menu(access_token)
    menu = get_menu(access_token)
    print(menu)
