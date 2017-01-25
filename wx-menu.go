package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/levigross/grequests"
)

// use byte.Buffer to create url
func make_url(lst ...string) string {
	b := bytes.Buffer{}
	for _, s := range lst {
		b.WriteString(s)
	}
	url := b.String()
	return url
}

// get the access_token from weixin.qq.com by appid and app_secret
func get_access(app_id string, app_secret string) string {
	access_url := make_url("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=", app_id, "&secret=", app_secret)
	resp, err := grequests.Get(access_url, nil)
	if err != nil {
		panic(err)
	}
	if resp.Ok != true {
		panic("Request did not return OK")
	}

	json, _ := ioutil.ReadAll(resp)
	result, err := simplejson.NewJson(json)
	if err != nil {
		panic(err)
	}
	access_token := result.Get("access_token").MustString()
	return access_token
}

// crate menu
// config file default : menu.json
func create_menu(access_token string) {
	path := "menu.json"
	menu, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	menu_url := make_url("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=", access_token)

	resp, _ := grequests.Post(menu_url, &grequests.RequestOptions{JSON: menu, IsAjax: true})
	if resp.Error != nil {
		panic(resp.Error)
	}

	if resp.Ok != true {
		panic("Request did not return OK")
	}

	json, _ := ioutil.ReadAll(resp)
	result, err := simplejson.NewJson(json)
	if err != nil {
		panic(err)
	}
	errmsg := result.Get("errmsg").MustString()
	if errmsg == "ok" {
		fmt.Println("Set menu success")
	} else {
		fmt.Println("Set menu failed, try again")
	}
}

// get menu
func get_menu(access_token string) {
	menu_url := make_url("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=", access_token)
	resp, err := grequests.Get(menu_url, nil)
	if err != nil {
		panic(err)
	}
	if resp.Ok != true {
		panic("Request did not return OK")
	}

	fmt.Println(resp)
}

// delete menu 慎用
func delete_menu(access_token string) {
	menu_url := make_url("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=", access_token)
	resp, err := grequests.Get(menu_url, nil)
	if err != nil {
		panic(err)
	}
	if resp.Ok != true {
		panic("Request did not return OK")
	}

	json, _ := ioutil.ReadAll(resp)
	result, err := simplejson.NewJson(json)
	if err != nil {
		panic(err)
	}
	errmsg := result.Get("errmsg").MustString()
	if errmsg == "ok" {
		fmt.Println("delete menu success")
	} else {
		fmt.Println("delete menu failed, try again")
	}
}

func main() {
	app_id := os.Getenv("WX_APP_ID")
	app_secret := os.Getenv("WX_APP_SECRET")

	access_token := get_access(app_id, app_secret)
	create_menu(access_token)
	get_menu(access_token)
	// delete_menu(access_token)
}
