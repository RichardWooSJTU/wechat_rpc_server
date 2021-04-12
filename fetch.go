package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wechatpro/gen-go/wechat"
)

const (
	prefix = "http://81.71.124.134:8765/"
)

func CreateReq(bodyJson []byte, addr string ) (*http.Request, error){
	req, err := http.NewRequest(http.MethodPost, prefix+ addr, bytes.NewReader(bodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Authorization", auth)
	return req, nil
}

func VerifyWechatOnline(wId string) error {
	bodyJson, err := json.Marshal(map[string]string{
		"wId": wId,
	})
	if err != nil {
		return err
	}
	req, err := CreateReq(bodyJson,"getIPadLoginInfo")
	if err != nil {
		return err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBody := resp.Body
	defer respBody.Close()
	respBytes, _ := ioutil.ReadAll(respBody)
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return err
	}
	if respMap["code"] == "1000" {
		log.Println("确认登录")
	} else {
		log.Println("确认登录失败")
	}
	return nil
}

func InitAddressList(wId string) error {
	bodyJson, err := json.Marshal(map[string]string{
		"wId": wId,
	})
	if err != nil {
		return err
	}
	req, err := CreateReq(bodyJson, "initAddressList")
	if err != nil {
		return err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBody := resp.Body
	defer respBody.Close()
	respBytes, _ := ioutil.ReadAll(respBody)
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return err
	}
	if respMap["code"] == "1000" {
		log.Println("初始化通讯录列表成功")
	} else {
		log.Println("初始化通讯录列表失败")
	}
	return nil
}

func QueryGroups(wId string) ([]*wechat.Group, error) {
	bodyJson, err := json.Marshal(map[string]string{
		"wId": wId,
	})
	if err != nil {
		return nil, err
	}
	req, err := CreateReq(bodyJson, "getAddressList")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody := resp.Body
	defer respBody.Close()
	respBytes, _ := ioutil.ReadAll(respBody)
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return nil, err
	}
	var groups []*wechat.Group
	if respMap["code"] == "1000" {
		dataInterface := respMap["data"]
		dataMap := dataInterface.(map[string]interface{})
		if _, ok := dataMap["chatrooms"]; ok {
			chatRoomsInterface := dataMap["chatrooms"]
			chatRooms := chatRoomsInterface.([]interface{})
			for _, chatRoom := range chatRooms {
				group := wechat.Group{
					GroupID:   chatRoom.(string),
					GroupName: "",
				}
				groups = append(groups, &group)
			}
		}
	}
	return groups, nil
}

func GetGroupDetail(wId string, groups []*wechat.Group) error {
	var groupIds string
	for i := 0; i < len(groups); i++ {
		groupIds += groups[i].GroupID
		if i != len(groups)-1 {
			groupIds += ","
		}
	}
	bodyJson, err := json.Marshal(map[string]string{
		"wId": wId,
		"wcId": groupIds,
	})
	if err != nil {
		return err
	}
	req, err := CreateReq(bodyJson, "getContact")
	if err != nil {
		return err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBody := resp.Body
	defer respBody.Close()
	respBytes, _ := ioutil.ReadAll(respBody)
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return  err
	}
	if respMap["code"] == "1000" {
		dataInterface := respMap["data"]
		dataSlice := dataInterface.([]interface{})
		for i:=0; i< len(dataSlice); i++ {
			m := dataSlice[i]
			groupDetail := m.(map[string]interface{})
			groupName := groupDetail["nickName"].(string)
			groups[i].GroupName = groupName
		}
	}
	return nil
}