package config

// JSONBody 返回 Response 的 JSON 格式.
// 从这里可以生成 https://mholt.github.io/json-to-go/
type JSONBody struct {
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

// FaStarsModel fa_stars 数据表.
type FaStarsModel struct {
	ID                       int    `gorm:"type:int(11); primaryKey; autoIncrement"`
	CreatorId                uint64 `gorm:"type:bigint(20)"`
	CreatorName              string `gorm:"type:varchar(255)"`
	CreatorNickname          string `gorm:"type:varchar(255)"`
	ContactMails             string `gorm:"type:varchar(255)"`
	ContactWhatsapps         string `gorm:"type:varchar(255)"`
	ContactCountryCodes      string `gorm:"type:varchar(100)"`
	Avatar                   string `gorm:"type:text"`
	IsVerified               int    `gorm:"type:tinyint(2)"`
	Region                   string `gorm:"type:varchar(255)"`
	ProductCategories        string `gorm:"type:varchar(255)"`
	Whatsappswitch           int    `gorm:"type:int(11)"`
	Mailswitch               int    `gorm:"type:int(11)"`
	FollowerCnt              int    `gorm:"type:int(11)"`
	FollowerTopGender        int    `gorm:"type:int(11)"`
	FollowerTopGenderShare   string `gorm:"type:varchar(255)"`
	FollowerTopAgeGroup      int    `gorm:"type:int(11)"`
	FollowerTopAgeGroupShare string `gorm:"type:varchar(255)"`
	VideoAvgViewCnt          int    `gorm:"type:int(11)"`
	VideoPubCnt              int    `gorm:"type:int(11)"`
	ProductCnt               int    `gorm:"type:int(11)"`
	Createtime               int64  `gorm:"type:int(11)"`
	Updatetime               int64  `gorm:"type:int(11)"`
}
