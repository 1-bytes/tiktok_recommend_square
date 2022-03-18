package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"tiktok/bootstrap"
	configs "tiktok/config"
	"tiktok/fetcher"
	"tiktok/pkg/config"
	"time"
)

// main 入口程序.
func main() {
	bootstrap.Setup()
	api := config.GetString("tiktok.api")
	body := config.GetString("tiktok.body")
	email := config.GetString("tiktok.email")
	password := config.GetString("tiktok.password")

	_, guard, err := Login(email, password)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Login successfully! the guard is: %s", guard)
	header := &http.Header{}
	header.Add("cookie", guard)
	rows, err := GetAPIData(api, body, header, 0)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Affects %d rows of data\n", rows)

	fmt.Printf("The program ends after 10 seconds ...")
	time.Sleep(10 * time.Second)
	// fmt.Printf("Press any key to exit...")
	// b := make([]byte, 1)
	// os.Stdin.Read(b)
}

func Login(email string, password string) (userinfo string, guard string, err error) {
	api := "https://seller.tiktokglobalshop.com/passport/web/user/login"
	body := map[string]string{
		"mix_mode":           "1",
		"aid":                "6556",
		"language":           "zh",
		"account_sdk_source": "web",
		"email":              email,
		"mobile":             "",
		"account":            email,
		"password":           password,
	}
	header := &http.Header{}
	resp, err := fetcher.FormData(http.MethodPost, api, &body, header)
	if err != nil {
		return "", "", err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	respBody := configs.LoginJSONBody{}
	_ = json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "success" {
		return "", "", fmt.Errorf(
			"login failed, please check if the login API is invalid")
	}

	// 关闭请求
	_ = resp.Body.Close()

	// 获取返回的Cookie
	var respHeader []string
	respHeaders := resp.Header.Values("Set-Cookie")
	for _, v := range respHeaders {
		if !strings.Contains(v, "sid_guard") {
			continue
		}
		respHeader = strings.Split(v, ";")
	}
	return string(bodyBytes), respHeader[0], nil
}

// GetAPIData 获取 TIKTOK API 接口数据.
func GetAPIData(api string, body string, header *http.Header, page int) (int, error) {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	var data []configs.FaStarsModel
	for {
		// 请求接口
		log.Printf("Currently collecting page %d...\n", page+1)
		resp, err := fetcher.Json(http.MethodPost, api, fmt.Sprintf(body, page), header)
		if err != nil {
			return 0, err
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		// 关闭请求
		_ = resp.Body.Close()
		// 解析 Json
		respBody := configs.ApiJSONBody{}
		err = json.Unmarshal(bodyBytes, &respBody)
		if err != nil {
			return 0, err
		}
		if respBody.Message != "success" {
			return 0, fmt.Errorf("request failed, unintended message: %s", respBody.Message)
		}
		// 插入到数据库
		for _, profile := range respBody.Data.Profiles {
			// --------------- 字段处理 开始 ---------------
			// 用户ID
			creatorId, _ := strconv.ParseUint(profile.CreatorID, 10, 64)
			// 用户已验证
			isVerified := 0
			if profile.IsVerified {
				isVerified = 1
			}
			// 联系方式
			var email, whatsapps, CountryCodes string
			if profile.Contact.Emails != nil {
				email = profile.Contact.Emails[0]
			}
			if profile.Contact.Whatsapps != nil {
				whatsapps = profile.Contact.Whatsapps[0]
			}
			if profile.Contact.CountryCodes != nil {
				CountryCodes = profile.Contact.CountryCodes[0]
			}

			// --------------- 字段处理 结束 ---------------
			data = append(data, configs.FaStarsModel{
				CreatorId:                creatorId,
				CreatorName:              profile.CreatorName,
				CreatorNickname:          profile.CreatorNickname,
				ContactMails:             email,
				ContactWhatsapps:         whatsapps,
				ContactCountryCodes:      CountryCodes,
				Avatar:                   profile.Avatar.URLList[0],
				IsVerified:               isVerified,
				Region:                   profile.Region,
				ProductCategories:        strings.Join(profile.ProductCategories, ","),
				Whatsappswitch:           0,
				Mailswitch:               0,
				FollowerCnt:              profile.FollowerCnt,
				FollowerTopGender:        profile.FollowerTopGender,
				FollowerTopGenderShare:   profile.FollowerTopGenderShare,
				FollowerTopAgeGroup:      profile.FollowerTopAgeGroup,
				FollowerTopAgeGroupShare: profile.FollowerTopAgeGroupShare,
				VideoAvgViewCnt:          profile.VideoAvgViewCnt,
				VideoPubCnt:              profile.VideoPubCnt,
				ProductCnt:               profile.ProductCnt,
				Createtime:               tm1.Unix(),
				Updatetime:               tm1.Unix(),
			})
		}
		// 没有下一页了就跳出循环，有的话就继续找下一页，出错了直接 return 不会走到这里
		if !respBody.Data.NextPagination.HasMore {
			break
		}
		page++
	}
	table := config.GetString("mysql.table")
	tx := bootstrap.DB.Table(table).Create(data)
	if tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf(
			"error! inserts failed, error message: %w", tx.Error)
	}
	return int(tx.RowsAffected), nil
}
