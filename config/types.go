package config

// ApiJSONBody 获取达人接口的 JSON 结构.
// 从这里可以生成 https://mholt.github.io/json-to-go/
type ApiJSONBody struct {
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

// LoginJSONBody 登录接口的 JSON 结构.
type LoginJSONBody struct {
	Data struct {
		AppID                      int           `json:"app_id"`
		UserID                     int64         `json:"user_id"`
		UserIDStr                  string        `json:"user_id_str"`
		OdinUserType               int           `json:"odin_user_type"`
		Name                       string        `json:"name"`
		ScreenName                 string        `json:"screen_name"`
		AvatarURL                  string        `json:"avatar_url"`
		UserVerified               bool          `json:"user_verified"`
		VerifiedContent            string        `json:"verified_content"`
		VerifiedAgency             string        `json:"verified_agency"`
		IsBlocked                  int           `json:"is_blocked"`
		IsBlocking                 int           `json:"is_blocking"`
		BgImgURL                   string        `json:"bg_img_url"`
		Gender                     int           `json:"gender"`
		MediaID                    int           `json:"media_id"`
		UserAuthInfo               string        `json:"user_auth_info"`
		Industry                   string        `json:"industry"`
		Area                       string        `json:"area"`
		CanBeFoundByPhone          int           `json:"can_be_found_by_phone"`
		Mobile                     string        `json:"mobile"`
		Birthday                   string        `json:"birthday"`
		Description                string        `json:"description"`
		Email                      string        `json:"email"`
		NewUser                    int           `json:"new_user"`
		SessionKey                 string        `json:"session_key"`
		IsRecommendAllowed         int           `json:"is_recommend_allowed"`
		RecommendHintMessage       string        `json:"recommend_hint_message"`
		Connects                   []interface{} `json:"connects"`
		FollowingsCount            int           `json:"followings_count"`
		FollowersCount             int           `json:"followers_count"`
		VisitCountRecent           int           `json:"visit_count_recent"`
		SkipEditProfile            int           `json:"skip_edit_profile"`
		IsManualSetUserInfo        bool          `json:"is_manual_set_user_info"`
		DeviceID                   int           `json:"device_id"`
		CountryCode                int           `json:"country_code"`
		HasPassword                int           `json:"has_password"`
		ShareToRepost              int           `json:"share_to_repost"`
		UserDecoration             string        `json:"user_decoration"`
		UserPrivacyExtend          int           `json:"user_privacy_extend"`
		OldUserID                  int64         `json:"old_user_id"`
		OldUserIDStr               string        `json:"old_user_id_str"`
		SecUserID                  string        `json:"sec_user_id"`
		SecOldUserID               string        `json:"sec_old_user_id"`
		VcdAccount                 int           `json:"vcd_account"`
		VcdRelation                int           `json:"vcd_relation"`
		CanBindVisitorAccount      bool          `json:"can_bind_visitor_account"`
		IsVisitorAccount           bool          `json:"is_visitor_account"`
		IsOnlyBindIns              bool          `json:"is_only_bind_ins"`
		UserDeviceRecordStatus     int           `json:"user_device_record_status"`
		IsKidsMode                 int           `json:"is_kids_mode"`
		IsEmployee                 bool          `json:"is_employee"`
		PassportEnterpriseUserType int           `json:"passport_enterprise_user_type"`
		NeedDeviceCreate           int           `json:"need_device_create"`
		NeedTtwidMigration         int           `json:"need_ttwid_migration"`
		ExpiredUIDList             interface{}   `json:"expired_uid_list"`
	} `json:"data"`
	Message string `json:"message"`
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
	IsFirst                  int    `gorm:"type:(2)"`
	Createtime               int64  `gorm:"type:int(11)"`
	Updatetime               int64  `gorm:"type:int(11)"`
}
