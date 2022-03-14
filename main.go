package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"tiktok/bootstrap"
	"tiktok/pkg/config"
)

type JsonBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Profiles []struct {
			CreatorID       string `json:"creator_id"`
			CreatorName     string `json:"creator_name"`
			CreatorNickname string `json:"creator_nickname"`
			Avatar          struct {
				ThumbURLList []string `json:"thumb_url_list"`
				URLList      []string `json:"url_list"`
			} `json:"avatar"`
			IsVerified bool   `json:"is_verified"`
			Region     string `json:"region"`
			Contact    struct {
				Emails       []string `json:"emails"`
				Whatsapps    []string `json:"whatsapps"`
				CountryCodes []string `json:"country_codes"`
			} `json:"contact"`
			ProductCategories        []string `json:"product_categories"`
			FollowerCnt              int      `json:"follower_cnt"`
			FollowerTopGender        int      `json:"follower_top_gender"`
			FollowerTopGenderShare   string   `json:"follower_top_gender_share"`
			FollowerTopAgeGroup      int      `json:"follower_top_age_group"`
			FollowerTopAgeGroupShare string   `json:"follower_top_age_group_share"`
			VideoAvgViewCnt          int      `json:"video_avg_view_cnt"`
			VideoPubCnt              int      `json:"video_pub_cnt"`
			ProductCnt               int      `json:"product_cnt"`
		} `json:"profiles"`
		NextPagination struct {
			HasMore   bool   `json:"has_more"`
			NextPage  int    `json:"next_page"`
			TotalPage int    `json:"total_page"`
			Total     int    `json:"total"`
			SearchKey string `json:"search_key"`
		} `json:"next_pagination"`
	} `json:"data"`
}

// main 入口程序.
func main() {
	bootstrap.Setup()
	api := config.GetString("tiktok.api")
	body := config.GetString("tiktok.body")
	cookie := config.GetString("tiktok.cookie")

	header := &http.Header{}
	header.Add("cookie", cookie)
	GetApiData(api, body, header, 0)
}

// GetApiData 获取 TIKTOK API 接口数据.
func GetApiData(api string, body string, header *http.Header, page int) error {
	// 请求接口
	api = fmt.Sprintf(api, page)
	bytes, err := Fetcher(http.MethodPost, api, body, header)
	if err != nil {
		return err
	}
	// 解析 Json
	respBody := JsonBody{}
	err = json.Unmarshal(bytes, &respBody)
	if err != nil {
		return err
	}
	if respBody.Message != "success" {
		return fmt.Errorf("request failed, unintended message: %s", respBody.Message)
	}
	for _, profile := range respBody.Data.Profiles {
		fmt.Println(profile)
	}

	// 递归函数
	if respBody.Data.NextPagination.HasMore {
		return GetApiData(api, body, header, page+1)
	}

	return nil
}

// Fetcher 获取网页内容.
func Fetcher(method string, api string, body string, header *http.Header) ([]byte, error) {
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}
	header.Add("Host", api)
	header.Add("Content-Length", strconv.Itoa(len(u.Host)))
	header.Add("Content-type", "application/json")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	req, err := http.NewRequest(method, api, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = *header
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, unintended status code: %d", resp.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}
