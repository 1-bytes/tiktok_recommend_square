package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"tiktok/bootstrap"
	configs "tiktok/config"
	"tiktok/pkg/config"
	"time"
)

// main 入口程序.
func main() {
	bootstrap.Setup()
	api := config.GetString("tiktok.api")
	body := config.GetString("tiktok.body")
	cookie := config.GetString("tiktok.cookie")

	header := &http.Header{}
	header.Add("cookie", cookie)
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

// GetAPIData 获取 TIKTOK API 接口数据.
func GetAPIData(api string, body string, header *http.Header, page int) (int, error) {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	var data []configs.FaStarsModel
	for {
		// 请求接口
		log.Printf("Currently collecting page %d...\n", page+1)
		bytes, err := Fetcher(http.MethodPost, api, fmt.Sprintf(body, page), header)
		if err != nil {
			return 0, err
		}
		// 解析 Json
		respBody := configs.JSONBody{}
		err = json.Unmarshal(bytes, &respBody)
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
		return int(tx.RowsAffected), fmt.Errorf("error! inserts failed, error message: %v", tx.Error)
	}
	return int(tx.RowsAffected), nil
}

// Fetcher 获取网页内容.
func Fetcher(method string, api string, body string, header *http.Header) ([]byte, error) {
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}
	if header.Get("Host") != u.Host {
		header.Add("Host", u.Host)
		header.Add("Content-Length", strconv.Itoa(len(body)))
		header.Add("Content-type", "application/json")
		header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	}
	req, err := http.NewRequest(method, api, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = *header
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, unintended status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
