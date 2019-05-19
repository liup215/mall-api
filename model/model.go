package model

type Ewei_shop_live_goods_option struct {
	Id        int     `orm:"id" json:"id"`
	Uniacid   int     `orm:"uniacid" json:"uniacid"`
	Goodsid   int     `orm:"goodsid" json:"goodsid"`
	Optionid  int     `orm:"optionid" json:"optionid"`
	Liveid    int     `orm:"liveid" json:"liveid"`
	Liveprice float64 `orm:"liveprice" json:"liveprice"`
}

func (*Ewei_shop_live_goods_option) TableName() string {
	return "ewei_shop_live_goods_option"
}

type Ewei_shop_poster_scan struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Posterid   int    `orm:"posterid" json:"posterid"`
	Openid     string `orm:"openid" json:"openid"`
	FromOpenid string `orm:"from_openid" json:"from_openid"`
	Scantime   int    `orm:"scantime" json:"scantime"`
}

func (*Ewei_shop_poster_scan) TableName() string {
	return "ewei_shop_poster_scan"
}

type Ewei_shop_quick_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Merchid      int    `orm:"merchid" json:"merchid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_quick_adv) TableName() string {
	return "ewei_shop_quick_adv"
}

type Ewei_shop_task_extension struct {
	Id           int    `orm:"id" json:"id"`
	Taskname     string `orm:"taskname" json:"taskname"`
	Taskclass    string `orm:"taskclass" json:"taskclass"`
	Status       int    `orm:"status" json:"status"`
	Classify     string `orm:"classify" json:"classify"`
	ClassifyName string `orm:"classify_name" json:"classify_name"`
	Verb         string `orm:"verb" json:"verb"`
	Unit         string `orm:"unit" json:"unit"`
}

func (*Ewei_shop_task_extension) TableName() string {
	return "ewei_shop_task_extension"
}

type Mc_mapping_ucenter struct {
	Id        int `orm:"id" json:"id"`
	Uniacid   int `orm:"uniacid" json:"uniacid"`
	Uid       int `orm:"uid" json:"uid"`
	Centeruid int `orm:"centeruid" json:"centeruid"`
}

func (*Mc_mapping_ucenter) TableName() string {
	return "mc_mapping_ucenter"
}

type Rule_keyword struct {
	Id           int    `orm:"id" json:"id"`
	Rid          int    `orm:"rid" json:"rid"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Module       string `orm:"module" json:"module"`
	Content      string `orm:"content" json:"content"`
	Type         int    `orm:"type" json:"type"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
}

func (*Rule_keyword) TableName() string {
	return "rule_keyword"
}

type Stat_msg_history struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Rid        int    `orm:"rid" json:"rid"`
	Kid        int    `orm:"kid" json:"kid"`
	FromUser   string `orm:"from_user" json:"from_user"`
	Module     string `orm:"module" json:"module"`
	Message    string `orm:"message" json:"message"`
	Type       string `orm:"type" json:"type"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Stat_msg_history) TableName() string {
	return "stat_msg_history"
}

type Users_permission struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Uid        int    `orm:"uid" json:"uid"`
	Type       string `orm:"type" json:"type"`
	Permission string `orm:"permission" json:"permission"`
	Url        string `orm:"url" json:"url"`
}

func (*Users_permission) TableName() string {
	return "users_permission"
}

type Wxapp_versions struct {
	Id             int    `orm:"id" json:"id"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Multiid        int    `orm:"multiid" json:"multiid"`
	Version        string `orm:"version" json:"version"`
	Description    string `orm:"description" json:"description"`
	Modules        string `orm:"modules" json:"modules"`
	DesignMethod   int    `orm:"design_method" json:"design_method"`
	Template       int    `orm:"template" json:"template"`
	Quickmenu      string `orm:"quickmenu" json:"quickmenu"`
	Createtime     int    `orm:"createtime" json:"createtime"`
	Redirect       string `orm:"redirect" json:"redirect"`
	Connection     string `orm:"connection" json:"connection"`
	Type           int    `orm:"type" json:"type"`
	EntryId        int    `orm:"entry_id" json:"entry_id"`
	Appjson        string `orm:"appjson" json:"appjson"`
	DefaultAppjson string `orm:"default_appjson" json:"default_appjson"`
	UseDefault     int    `orm:"use_default" json:"use_default"`
}

func (*Wxapp_versions) TableName() string {
	return "wxapp_versions"
}

type Ewei_shop_area_config struct {
	Id            int `orm:"id" json:"id"`
	Uniacid       int `orm:"uniacid" json:"uniacid"`
	NewArea       int `orm:"new_area" json:"new_area"`
	AddressStreet int `orm:"address_street" json:"address_street"`
	Createtime    int `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_area_config) TableName() string {
	return "ewei_shop_area_config"
}

type Ewei_shop_bargain_actor struct {
	Id           int     `orm:"id" json:"id"`
	GoodsId      int     `orm:"goods_id" json:"goods_id"`
	NowPrice     float64 `orm:"now_price" json:"now_price"`
	CreatedTime  string  `orm:"created_time" json:"created_time"`
	UpdateTime   string  `orm:"update_time" json:"update_time"`
	BargainTimes int     `orm:"bargain_times" json:"bargain_times"`
	Openid       string  `orm:"openid" json:"openid"`
	Nickname     string  `orm:"nickname" json:"nickname"`
	HeadImage    string  `orm:"head_image" json:"head_image"`
	BargainPrice float64 `orm:"bargain_price" json:"bargain_price"`
	Status       int     `orm:"status" json:"status"`
	AccountId    int     `orm:"account_id" json:"account_id"`
	Initiate     int     `orm:"initiate" json:"initiate"`
	Order        int     `orm:"order" json:"order"`
}

func (*Ewei_shop_bargain_actor) TableName() string {
	return "ewei_shop_bargain_actor"
}

type Ewei_shop_task_record struct {
	Id               int    `orm:"id" json:"id"`
	Uniacid          int    `orm:"uniacid" json:"uniacid"`
	Taskid           int    `orm:"taskid" json:"taskid"`
	Tasktitle        string `orm:"tasktitle" json:"tasktitle"`
	Taskimage        string `orm:"taskimage" json:"taskimage"`
	Tasktype         string `orm:"tasktype" json:"tasktype"`
	TaskProgress     int    `orm:"task_progress" json:"task_progress"`
	TaskDemand       int    `orm:"task_demand" json:"task_demand"`
	Openid           string `orm:"openid" json:"openid"`
	Nickname         string `orm:"nickname" json:"nickname"`
	Picktime         string `orm:"picktime" json:"picktime"`
	Stoptime         string `orm:"stoptime" json:"stoptime"`
	Finishtime       string `orm:"finishtime" json:"finishtime"`
	RewardData       string `orm:"reward_data" json:"reward_data"`
	FollowrewardData string `orm:"followreward_data" json:"followreward_data"`
	DesignData       string `orm:"design_data" json:"design_data"`
	DesignBg         string `orm:"design_bg" json:"design_bg"`
	RequireGoods     string `orm:"require_goods" json:"require_goods"`
	Level1           int    `orm:"level1" json:"level1"`
	RewardData1      string `orm:"reward_data1" json:"reward_data1"`
	Level2           int    `orm:"level2" json:"level2"`
	RewardData2      string `orm:"reward_data2" json:"reward_data2"`
	MemberGroup      string `orm:"member_group" json:"member_group"`
	AutoPick         int    `orm:"auto_pick" json:"auto_pick"`
}

func (*Ewei_shop_task_record) TableName() string {
	return "ewei_shop_task_record"
}

type Uni_account_menus struct {
	Id                 int    `orm:"id" json:"id"`
	Uniacid            int    `orm:"uniacid" json:"uniacid"`
	Menuid             int    `orm:"menuid" json:"menuid"`
	Type               int    `orm:"type" json:"type"`
	Title              string `orm:"title" json:"title"`
	Sex                int    `orm:"sex" json:"sex"`
	GroupId            int    `orm:"group_id" json:"group_id"`
	ClientPlatformType int    `orm:"client_platform_type" json:"client_platform_type"`
	Area               string `orm:"area" json:"area"`
	Data               string `orm:"data" json:"data"`
	Status             int    `orm:"status" json:"status"`
	Createtime         int    `orm:"createtime" json:"createtime"`
	Isdeleted          int    `orm:"isdeleted" json:"isdeleted"`
}

func (*Uni_account_menus) TableName() string {
	return "uni_account_menus"
}

type Ewei_shop_commission_log struct {
	Id             int     `orm:"id" json:"id"`
	Uniacid        int     `orm:"uniacid" json:"uniacid"`
	Applyid        int     `orm:"applyid" json:"applyid"`
	Mid            int     `orm:"mid" json:"mid"`
	Commission     float64 `orm:"commission" json:"commission"`
	Createtime     int     `orm:"createtime" json:"createtime"`
	CommissionPay  float64 `orm:"commission_pay" json:"commission_pay"`
	Realmoney      float64 `orm:"realmoney" json:"realmoney"`
	Charge         float64 `orm:"charge" json:"charge"`
	Deductionmoney float64 `orm:"deductionmoney" json:"deductionmoney"`
	Type           int     `orm:"type" json:"type"`
}

func (*Ewei_shop_commission_log) TableName() string {
	return "ewei_shop_commission_log"
}

type Ewei_shop_member_level struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Level      int     `orm:"level" json:"level"`
	Levelname  string  `orm:"levelname" json:"levelname"`
	Ordermoney float64 `orm:"ordermoney" json:"ordermoney"`
	Ordercount int     `orm:"ordercount" json:"ordercount"`
	Discount   float64 `orm:"discount" json:"discount"`
	Enabled    int     `orm:"enabled" json:"enabled"`
	Enabledadd int     `orm:"enabledadd" json:"enabledadd"`
	Buygoods   int     `orm:"buygoods" json:"buygoods"`
	Goodsids   string  `orm:"goodsids" json:"goodsids"`
}

func (*Ewei_shop_member_level) TableName() string {
	return "ewei_shop_member_level"
}

type Ewei_shop_quick_cart struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Quickid       int     `orm:"quickid" json:"quickid"`
	Openid        string  `orm:"openid" json:"openid"`
	Goodsid       int     `orm:"goodsid" json:"goodsid"`
	Total         int     `orm:"total" json:"total"`
	Marketprice   float64 `orm:"marketprice" json:"marketprice"`
	Deleted       int     `orm:"deleted" json:"deleted"`
	Optionid      int     `orm:"optionid" json:"optionid"`
	Createtime    int     `orm:"createtime" json:"createtime"`
	Diyformdataid int     `orm:"diyformdataid" json:"diyformdataid"`
	Diyformdata   string  `orm:"diyformdata" json:"diyformdata"`
	Diyformfields string  `orm:"diyformfields" json:"diyformfields"`
	Diyformid     int     `orm:"diyformid" json:"diyformid"`
	Selected      int     `orm:"selected" json:"selected"`
	Merchid       int     `orm:"merchid" json:"merchid"`
	Selectedadd   int     `orm:"selectedadd" json:"selectedadd"`
}

func (*Ewei_shop_quick_cart) TableName() string {
	return "ewei_shop_quick_cart"
}

type Ewei_shop_system_link struct {
	Id           int    `orm:"id" json:"id"`
	Name         string `orm:"name" json:"name"`
	Url          string `orm:"url" json:"url"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_link) TableName() string {
	return "ewei_shop_system_link"
}

type Site_slide struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Multiid      int    `orm:"multiid" json:"multiid"`
	Title        string `orm:"title" json:"title"`
	Url          string `orm:"url" json:"url"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Site_slide) TableName() string {
	return "site_slide"
}

type Stat_rule struct {
	Id         int `orm:"id" json:"id"`
	Uniacid    int `orm:"uniacid" json:"uniacid"`
	Rid        int `orm:"rid" json:"rid"`
	Hit        int `orm:"hit" json:"hit"`
	Lastupdate int `orm:"lastupdate" json:"lastupdate"`
	Createtime int `orm:"createtime" json:"createtime"`
}

func (*Stat_rule) TableName() string {
	return "stat_rule"
}

type Ewei_shop_globonus_billo struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Billid     int     `orm:"billid" json:"billid"`
	Orderid    int     `orm:"orderid" json:"orderid"`
	Ordermoney float64 `orm:"ordermoney" json:"ordermoney"`
}

func (*Ewei_shop_globonus_billo) TableName() string {
	return "ewei_shop_globonus_billo"
}

type Ewei_shop_postera_qr struct {
	Id         int    `orm:"id" json:"id"`
	Acid       int    `orm:"acid" json:"acid"`
	Openid     string `orm:"openid" json:"openid"`
	Posterid   int    `orm:"posterid" json:"posterid"`
	Type       int    `orm:"type" json:"type"`
	Sceneid    int    `orm:"sceneid" json:"sceneid"`
	Mediaid    string `orm:"mediaid" json:"mediaid"`
	Ticket     string `orm:"ticket" json:"ticket"`
	Url        string `orm:"url" json:"url"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Goodsid    int    `orm:"goodsid" json:"goodsid"`
	Qrimg      string `orm:"qrimg" json:"qrimg"`
	Expire     int    `orm:"expire" json:"expire"`
	Endtime    int    `orm:"endtime" json:"endtime"`
	Qrtime     string `orm:"qrtime" json:"qrtime"`
}

func (*Ewei_shop_postera_qr) TableName() string {
	return "ewei_shop_postera_qr"
}

type Site_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Nid          int    `orm:"nid" json:"nid"`
	Name         string `orm:"name" json:"name"`
	Parentid     int    `orm:"parentid" json:"parentid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Icon         string `orm:"icon" json:"icon"`
	Description  string `orm:"description" json:"description"`
	Styleid      int    `orm:"styleid" json:"styleid"`
	Linkurl      string `orm:"linkurl" json:"linkurl"`
	Ishomepage   int    `orm:"ishomepage" json:"ishomepage"`
	Icontype     int    `orm:"icontype" json:"icontype"`
	Css          string `orm:"css" json:"css"`
	Multiid      int    `orm:"multiid" json:"multiid"`
}

func (*Site_category) TableName() string {
	return "site_category"
}

type Uni_account struct {
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Groupid      int    `orm:"groupid" json:"groupid"`
	Name         string `orm:"name" json:"name"`
	Description  string `orm:"description" json:"description"`
	DefaultAcid  int    `orm:"default_acid" json:"default_acid"`
	Rank         int    `orm:"rank" json:"rank"`
	TitleInitial string `orm:"title_initial" json:"title_initial"`
}

func (*Uni_account) TableName() string {
	return "uni_account"
}

type Uni_account_group struct {
	Id      int `orm:"id" json:"id"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Groupid int `orm:"groupid" json:"groupid"`
}

func (*Uni_account_group) TableName() string {
	return "uni_account_group"
}

type Ewei_shop_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Shopid       int    `orm:"shopid" json:"shopid"`
	Iswxapp      int    `orm:"iswxapp" json:"iswxapp"`
}

func (*Ewei_shop_adv) TableName() string {
	return "ewei_shop_adv"
}

type Ewei_shop_system_plugingrant_adv struct {
	Id           int    `orm:"id" json:"id"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_system_plugingrant_adv) TableName() string {
	return "ewei_shop_system_plugingrant_adv"
}

type Ewei_shop_system_setting struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Background string `orm:"background" json:"background"`
	Casebanner string `orm:"casebanner" json:"casebanner"`
	Contact    string `orm:"contact" json:"contact"`
}

func (*Ewei_shop_system_setting) TableName() string {
	return "ewei_shop_system_setting"
}

type Ewei_shop_virtual_data struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Typeid     int     `orm:"typeid" json:"typeid"`
	Pvalue     string  `orm:"pvalue" json:"pvalue"`
	Fields     string  `orm:"fields" json:"fields"`
	Openid     string  `orm:"openid" json:"openid"`
	Usetime    int     `orm:"usetime" json:"usetime"`
	Orderid    int     `orm:"orderid" json:"orderid"`
	Ordersn    string  `orm:"ordersn" json:"ordersn"`
	Price      float64 `orm:"price" json:"price"`
	Merchid    int     `orm:"merchid" json:"merchid"`
	Createtime int     `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_virtual_data) TableName() string {
	return "ewei_shop_virtual_data"
}

type Voice_reply struct {
	Id         int    `orm:"id" json:"id"`
	Rid        int    `orm:"rid" json:"rid"`
	Title      string `orm:"title" json:"title"`
	Mediaid    string `orm:"mediaid" json:"mediaid"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Voice_reply) TableName() string {
	return "voice_reply"
}

type Account_wxapp struct {
	Acid           int    `orm:"acid" json:"acid"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Token          string `orm:"token" json:"token"`
	Encodingaeskey string `orm:"encodingaeskey" json:"encodingaeskey"`
	Level          int    `orm:"level" json:"level"`
	Account        string `orm:"account" json:"account"`
	Original       string `orm:"original" json:"original"`
	Key            string `orm:"key" json:"key"`
	Secret         string `orm:"secret" json:"secret"`
	Name           string `orm:"name" json:"name"`
	Appdomain      string `orm:"appdomain" json:"appdomain"`
}

func (*Account_wxapp) TableName() string {
	return "account_wxapp"
}

type Ewei_shop_banner struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Bannername   string `orm:"bannername" json:"bannername"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Shopid       int    `orm:"shopid" json:"shopid"`
	Iswxapp      int    `orm:"iswxapp" json:"iswxapp"`
}

func (*Ewei_shop_banner) TableName() string {
	return "ewei_shop_banner"
}

type Ewei_shop_sendticket struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Cpid       string `orm:"cpid" json:"cpid"`
	Expiration int    `orm:"expiration" json:"expiration"`
	Starttime  int    `orm:"starttime" json:"starttime"`
	Endtime    int    `orm:"endtime" json:"endtime"`
	Status     int    `orm:"status" json:"status"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Title      string `orm:"title" json:"title"`
}

func (*Ewei_shop_sendticket) TableName() string {
	return "ewei_shop_sendticket"
}

type Activity_stores struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	BusinessName string `orm:"business_name" json:"business_name"`
	BranchName   string `orm:"branch_name" json:"branch_name"`
	Category     string `orm:"category" json:"category"`
	Province     string `orm:"province" json:"province"`
	City         string `orm:"city" json:"city"`
	District     string `orm:"district" json:"district"`
	Address      string `orm:"address" json:"address"`
	Longitude    string `orm:"longitude" json:"longitude"`
	Latitude     string `orm:"latitude" json:"latitude"`
	Telephone    string `orm:"telephone" json:"telephone"`
	PhotoList    string `orm:"photo_list" json:"photo_list"`
	AvgPrice     int    `orm:"avg_price" json:"avg_price"`
	Recommend    string `orm:"recommend" json:"recommend"`
	Special      string `orm:"special" json:"special"`
	Introduction string `orm:"introduction" json:"introduction"`
	OpenTime     string `orm:"open_time" json:"open_time"`
	LocationId   int    `orm:"location_id" json:"location_id"`
	Status       int    `orm:"status" json:"status"`
	Source       int    `orm:"source" json:"source"`
	Message      string `orm:"message" json:"message"`
}

func (*Activity_stores) TableName() string {
	return "activity_stores"
}

type Ewei_message_mass_template struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Title       string `orm:"title" json:"title"`
	TemplateId  string `orm:"template_id" json:"template_id"`
	First       string `orm:"first" json:"first"`
	Firstcolor  string `orm:"firstcolor" json:"firstcolor"`
	Data        string `orm:"data" json:"data"`
	Remark      string `orm:"remark" json:"remark"`
	Remarkcolor string `orm:"remarkcolor" json:"remarkcolor"`
	Url         string `orm:"url" json:"url"`
	Createtime  int    `orm:"createtime" json:"createtime"`
	Sendtimes   int    `orm:"sendtimes" json:"sendtimes"`
	Sendcount   int    `orm:"sendcount" json:"sendcount"`
	Miniprogram int    `orm:"miniprogram" json:"miniprogram"`
	Appid       string `orm:"appid" json:"appid"`
	Pagepath    string `orm:"pagepath" json:"pagepath"`
}

func (*Ewei_message_mass_template) TableName() string {
	return "ewei_message_mass_template"
}

type Ewei_shop_article_report struct {
	Id      int    `orm:"id" json:"id"`
	Mid     int    `orm:"mid" json:"mid"`
	Openid  string `orm:"openid" json:"openid"`
	Aid     int    `orm:"aid" json:"aid"`
	Cate    string `orm:"cate" json:"cate"`
	Cons    string `orm:"cons" json:"cons"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
}

func (*Ewei_shop_article_report) TableName() string {
	return "ewei_shop_article_report"
}

type Ewei_shop_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Parentid     int    `orm:"parentid" json:"parentid"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
	Description  string `orm:"description" json:"description"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Ishome       int    `orm:"ishome" json:"ishome"`
	Advimg       string `orm:"advimg" json:"advimg"`
	Advurl       string `orm:"advurl" json:"advurl"`
	Level        int    `orm:"level" json:"level"`
}

func (*Ewei_shop_category) TableName() string {
	return "ewei_shop_category"
}

type Ewei_shop_goods_comment struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Goodsid    int    `orm:"goodsid" json:"goodsid"`
	Openid     string `orm:"openid" json:"openid"`
	Nickname   string `orm:"nickname" json:"nickname"`
	Headimgurl string `orm:"headimgurl" json:"headimgurl"`
	Content    string `orm:"content" json:"content"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_goods_comment) TableName() string {
	return "ewei_shop_goods_comment"
}

type Ewei_shop_system_plugingrant_plugin struct {
	Id           int    `orm:"id" json:"id"`
	Pluginid     int    `orm:"pluginid" json:"pluginid"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Data         string `orm:"data" json:"data"`
	State        int    `orm:"state" json:"state"`
	Content      string `orm:"content" json:"content"`
	Sales        int    `orm:"sales" json:"sales"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Plugintype   int    `orm:"plugintype" json:"plugintype"`
	Name         string `orm:"name" json:"name"`
}

func (*Ewei_shop_system_plugingrant_plugin) TableName() string {
	return "ewei_shop_system_plugingrant_plugin"
}

type Ewei_shop_task_log struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	FromOpenid string `orm:"from_openid" json:"from_openid"`
	JoinId     int    `orm:"join_id" json:"join_id"`
	Taskid     int    `orm:"taskid" json:"taskid"`
	TaskType   int    `orm:"task_type" json:"task_type"`
	Subdata    string `orm:"subdata" json:"subdata"`
	Recdata    string `orm:"recdata" json:"recdata"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_task_log) TableName() string {
	return "ewei_shop_task_log"
}

type Uni_account_users struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Uid     int    `orm:"uid" json:"uid"`
	Role    string `orm:"role" json:"role"`
	Rank    int    `orm:"rank" json:"rank"`
}

func (*Uni_account_users) TableName() string {
	return "uni_account_users"
}

type Core_cron struct {
	Id          int    `orm:"id" json:"id"`
	Cloudid     int    `orm:"cloudid" json:"cloudid"`
	Module      string `orm:"module" json:"module"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Type        int    `orm:"type" json:"type"`
	Name        string `orm:"name" json:"name"`
	Filename    string `orm:"filename" json:"filename"`
	Lastruntime int    `orm:"lastruntime" json:"lastruntime"`
	Nextruntime int    `orm:"nextruntime" json:"nextruntime"`
	Weekday     int    `orm:"weekday" json:"weekday"`
	Day         int    `orm:"day" json:"day"`
	Hour        int    `orm:"hour" json:"hour"`
	Minute      string `orm:"minute" json:"minute"`
	Extra       string `orm:"extra" json:"extra"`
	Status      int    `orm:"status" json:"status"`
	Createtime  int    `orm:"createtime" json:"createtime"`
}

func (*Core_cron) TableName() string {
	return "core_cron"
}

type Ewei_shop_diypage_plu struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Type         int    `orm:"type" json:"type"`
	Status       int    `orm:"status" json:"status"`
	Name         string `orm:"name" json:"name"`
	Data         string `orm:"data" json:"data"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Lastedittime int    `orm:"lastedittime" json:"lastedittime"`
	Merch        int    `orm:"merch" json:"merch"`
}

func (*Ewei_shop_diypage_plu) TableName() string {
	return "ewei_shop_diypage_plu"
}

type Ewei_shop_system_category struct {
	Id           int    `orm:"id" json:"id"`
	Name         string `orm:"name" json:"name"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_category) TableName() string {
	return "ewei_shop_system_category"
}

type Wxapp_general_analysis struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	SessionCnt      int    `orm:"session_cnt" json:"session_cnt"`
	VisitPv         int    `orm:"visit_pv" json:"visit_pv"`
	VisitUv         int    `orm:"visit_uv" json:"visit_uv"`
	VisitUvNew      int    `orm:"visit_uv_new" json:"visit_uv_new"`
	Type            int    `orm:"type" json:"type"`
	StayTimeUv      string `orm:"stay_time_uv" json:"stay_time_uv"`
	StayTimeSession string `orm:"stay_time_session" json:"stay_time_session"`
	VisitDepth      string `orm:"visit_depth" json:"visit_depth"`
	RefDate         string `orm:"ref_date" json:"ref_date"`
}

func (*Wxapp_general_analysis) TableName() string {
	return "wxapp_general_analysis"
}

type Activity_exchange_trades_shipping struct {
	Tid        int    `orm:"tid" json:"tid"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Exid       int    `orm:"exid" json:"exid"`
	Uid        int    `orm:"uid" json:"uid"`
	Status     int    `orm:"status" json:"status"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Province   string `orm:"province" json:"province"`
	City       string `orm:"city" json:"city"`
	District   string `orm:"district" json:"district"`
	Address    string `orm:"address" json:"address"`
	Zipcode    string `orm:"zipcode" json:"zipcode"`
	Mobile     string `orm:"mobile" json:"mobile"`
	Name       string `orm:"name" json:"name"`
}

func (*Activity_exchange_trades_shipping) TableName() string {
	return "activity_exchange_trades_shipping"
}

type Ewei_shop_coupon_goodsendtask struct {
	Id        int `orm:"id" json:"id"`
	Uniacid   int `orm:"uniacid" json:"uniacid"`
	Goodsid   int `orm:"goodsid" json:"goodsid"`
	Couponid  int `orm:"couponid" json:"couponid"`
	Starttime int `orm:"starttime" json:"starttime"`
	Endtime   int `orm:"endtime" json:"endtime"`
	Sendnum   int `orm:"sendnum" json:"sendnum"`
	Num       int `orm:"num" json:"num"`
	Sendpoint int `orm:"sendpoint" json:"sendpoint"`
	Status    int `orm:"status" json:"status"`
}

func (*Ewei_shop_coupon_goodsendtask) TableName() string {
	return "ewei_shop_coupon_goodsendtask"
}

type Ewei_shop_express struct {
	Id           int    `orm:"id" json:"id"`
	Name         string `orm:"name" json:"name"`
	Express      string `orm:"express" json:"express"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Code         string `orm:"code" json:"code"`
}

func (*Ewei_shop_express) TableName() string {
	return "ewei_shop_express"
}

type Ewei_shop_pc_link struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Linkname     string `orm:"linkname" json:"linkname"`
	Url          string `orm:"url" json:"url"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_pc_link) TableName() string {
	return "ewei_shop_pc_link"
}

type Ewei_shop_sendticket_share struct {
	Id          int     `orm:"id" json:"id"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Sharetitle  string  `orm:"sharetitle" json:"sharetitle"`
	Shareicon   string  `orm:"shareicon" json:"shareicon"`
	Sharedesc   string  `orm:"sharedesc" json:"sharedesc"`
	Expiration  int     `orm:"expiration" json:"expiration"`
	Starttime   int     `orm:"starttime" json:"starttime"`
	Endtime     int     `orm:"endtime" json:"endtime"`
	Paycpid1    int     `orm:"paycpid1" json:"paycpid1"`
	Paycpid2    int     `orm:"paycpid2" json:"paycpid2"`
	Paycpid3    int     `orm:"paycpid3" json:"paycpid3"`
	Paycpnum1   int     `orm:"paycpnum1" json:"paycpnum1"`
	Paycpnum2   int     `orm:"paycpnum2" json:"paycpnum2"`
	Paycpnum3   int     `orm:"paycpnum3" json:"paycpnum3"`
	Sharecpid1  int     `orm:"sharecpid1" json:"sharecpid1"`
	Sharecpid2  int     `orm:"sharecpid2" json:"sharecpid2"`
	Sharecpid3  int     `orm:"sharecpid3" json:"sharecpid3"`
	Sharecpnum1 int     `orm:"sharecpnum1" json:"sharecpnum1"`
	Sharecpnum2 int     `orm:"sharecpnum2" json:"sharecpnum2"`
	Sharecpnum3 int     `orm:"sharecpnum3" json:"sharecpnum3"`
	Status      int     `orm:"status" json:"status"`
	Createtime  int     `orm:"createtime" json:"createtime"`
	Order       int     `orm:"order" json:"order"`
	Enough      float64 `orm:"enough" json:"enough"`
	Issync      int     `orm:"issync" json:"issync"`
}

func (*Ewei_shop_sendticket_share) TableName() string {
	return "ewei_shop_sendticket_share"
}

type Modules_recycle struct {
	Id         int    `orm:"id" json:"id"`
	Modulename string `orm:"modulename" json:"modulename"`
}

func (*Modules_recycle) TableName() string {
	return "modules_recycle"
}

type Ewei_shop_datatransfer struct {
	Id          int `orm:"id" json:"id"`
	Fromuniacid int `orm:"fromuniacid" json:"fromuniacid"`
	Touniacid   int `orm:"touniacid" json:"touniacid"`
	Status      int `orm:"status" json:"status"`
}

func (*Ewei_shop_datatransfer) TableName() string {
	return "ewei_shop_datatransfer"
}

type Ewei_shop_diypage_menu struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Data         string `orm:"data" json:"data"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Lastedittime int    `orm:"lastedittime" json:"lastedittime"`
	Merch        int    `orm:"merch" json:"merch"`
}

func (*Ewei_shop_diypage_menu) TableName() string {
	return "ewei_shop_diypage_menu"
}

type Ewei_shop_system_casecategory struct {
	Id           int    `orm:"id" json:"id"`
	Name         string `orm:"name" json:"name"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_casecategory) TableName() string {
	return "ewei_shop_system_casecategory"
}

type Activity_exchange_trades struct {
	Tid        int `orm:"tid" json:"tid"`
	Uniacid    int `orm:"uniacid" json:"uniacid"`
	Uid        int `orm:"uid" json:"uid"`
	Exid       int `orm:"exid" json:"exid"`
	Type       int `orm:"type" json:"type"`
	Createtime int `orm:"createtime" json:"createtime"`
}

func (*Activity_exchange_trades) TableName() string {
	return "activity_exchange_trades"
}

type Ewei_shop_merch_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Catename     string `orm:"catename" json:"catename"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
}

func (*Ewei_shop_merch_category) TableName() string {
	return "ewei_shop_merch_category"
}

type Ewei_shop_merch_clearing struct {
	Id                int     `orm:"id" json:"id"`
	Uniacid           int     `orm:"uniacid" json:"uniacid"`
	Merchid           int     `orm:"merchid" json:"merchid"`
	Clearno           string  `orm:"clearno" json:"clearno"`
	Goodsprice        float64 `orm:"goodsprice" json:"goodsprice"`
	Dispatchprice     float64 `orm:"dispatchprice" json:"dispatchprice"`
	Deductprice       float64 `orm:"deductprice" json:"deductprice"`
	Deductcredit2     float64 `orm:"deductcredit2" json:"deductcredit2"`
	Discountprice     float64 `orm:"discountprice" json:"discountprice"`
	Deductenough      float64 `orm:"deductenough" json:"deductenough"`
	Merchdeductenough float64 `orm:"merchdeductenough" json:"merchdeductenough"`
	Isdiscountprice   float64 `orm:"isdiscountprice" json:"isdiscountprice"`
	Price             float64 `orm:"price" json:"price"`
	Createtime        int     `orm:"createtime" json:"createtime"`
	Starttime         int     `orm:"starttime" json:"starttime"`
	Endtime           int     `orm:"endtime" json:"endtime"`
	Status            int     `orm:"status" json:"status"`
	Realprice         float64 `orm:"realprice" json:"realprice"`
	Realpricerate     float64 `orm:"realpricerate" json:"realpricerate"`
	Finalprice        float64 `orm:"finalprice" json:"finalprice"`
	Remark            string  `orm:"remark" json:"remark"`
	Paytime           int     `orm:"paytime" json:"paytime"`
	Payrate           float64 `orm:"payrate" json:"payrate"`
}

func (*Ewei_shop_merch_clearing) TableName() string {
	return "ewei_shop_merch_clearing"
}

type Ewei_shop_exhelper_esheet_temp struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Esheetid     int    `orm:"esheetid" json:"esheetid"`
	Esheetname   string `orm:"esheetname" json:"esheetname"`
	Customername string `orm:"customername" json:"customername"`
	Customerpwd  string `orm:"customerpwd" json:"customerpwd"`
	Monthcode    string `orm:"monthcode" json:"monthcode"`
	Sendsite     string `orm:"sendsite" json:"sendsite"`
	Paytype      int    `orm:"paytype" json:"paytype"`
	Templatesize string `orm:"templatesize" json:"templatesize"`
	Isnotice     int    `orm:"isnotice" json:"isnotice"`
	Merchid      int    `orm:"merchid" json:"merchid"`
	Issend       int    `orm:"issend" json:"issend"`
	Isdefault    int    `orm:"isdefault" json:"isdefault"`
}

func (*Ewei_shop_exhelper_esheet_temp) TableName() string {
	return "ewei_shop_exhelper_esheet_temp"
}

type Ewei_shop_lottery_default struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Data    string `orm:"data" json:"data"`
	Addtime int    `orm:"addtime" json:"addtime"`
}

func (*Ewei_shop_lottery_default) TableName() string {
	return "ewei_shop_lottery_default"
}

type Ewei_shop_member_message_template_default struct {
	Id         int    `orm:"id" json:"id"`
	Typecode   string `orm:"typecode" json:"typecode"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Templateid string `orm:"templateid" json:"templateid"`
}

func (*Ewei_shop_member_message_template_default) TableName() string {
	return "ewei_shop_member_message_template_default"
}

type Userapi_cache struct {
	Id         int    `orm:"id" json:"id"`
	Key        string `orm:"key" json:"key"`
	Content    string `orm:"content" json:"content"`
	Lastupdate int    `orm:"lastupdate" json:"lastupdate"`
}

func (*Userapi_cache) TableName() string {
	return "userapi_cache"
}

type Core_refundlog struct {
	Id             int     `orm:"id" json:"id"`
	Uniacid        int     `orm:"uniacid" json:"uniacid"`
	RefundUniontid string  `orm:"refund_uniontid" json:"refund_uniontid"`
	Reason         string  `orm:"reason" json:"reason"`
	Uniontid       string  `orm:"uniontid" json:"uniontid"`
	Fee            float64 `orm:"fee" json:"fee"`
	Status         int     `orm:"status" json:"status"`
}

func (*Core_refundlog) TableName() string {
	return "core_refundlog"
}

type Ewei_shop_bargain_record struct {
	Id           int     `orm:"id" json:"id"`
	ActorId      int     `orm:"actor_id" json:"actor_id"`
	BargainPrice float64 `orm:"bargain_price" json:"bargain_price"`
	Openid       string  `orm:"openid" json:"openid"`
	Nickname     string  `orm:"nickname" json:"nickname"`
	HeadImage    string  `orm:"head_image" json:"head_image"`
	BargainTime  string  `orm:"bargain_time" json:"bargain_time"`
}

func (*Ewei_shop_bargain_record) TableName() string {
	return "ewei_shop_bargain_record"
}

type Ewei_shop_goods_spec_item struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Specid       int    `orm:"specid" json:"specid"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Show         int    `orm:"show" json:"show"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	ValueId      string `orm:"valueId" json:"valueId"`
	Virtual      int    `orm:"virtual" json:"virtual"`
}

func (*Ewei_shop_goods_spec_item) TableName() string {
	return "ewei_shop_goods_spec_item"
}

type Ewei_shop_notice struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Link         string `orm:"link" json:"link"`
	Detail       string `orm:"detail" json:"detail"`
	Status       int    `orm:"status" json:"status"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Shopid       int    `orm:"shopid" json:"shopid"`
	Iswxapp      int    `orm:"iswxapp" json:"iswxapp"`
}

func (*Ewei_shop_notice) TableName() string {
	return "ewei_shop_notice"
}

type Ewei_shop_sns_level struct {
	Id        int    `orm:"id" json:"id"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Levelname string `orm:"levelname" json:"levelname"`
	Credit    int    `orm:"credit" json:"credit"`
	Enabled   int    `orm:"enabled" json:"enabled"`
	Post      int    `orm:"post" json:"post"`
	Color     string `orm:"color" json:"color"`
	Bg        string `orm:"bg" json:"bg"`
}

func (*Ewei_shop_sns_level) TableName() string {
	return "ewei_shop_sns_level"
}

type Ewei_shop_article struct {
	Id                      int     `orm:"id" json:"id"`
	ArticleTitle            string  `orm:"article_title" json:"article_title"`
	RespDesc                string  `orm:"resp_desc" json:"resp_desc"`
	RespImg                 string  `orm:"resp_img" json:"resp_img"`
	ArticleContent          string  `orm:"article_content" json:"article_content"`
	ArticleCategory         int     `orm:"article_category" json:"article_category"`
	ArticleDateV            string  `orm:"article_date_v" json:"article_date_v"`
	ArticleDate             string  `orm:"article_date" json:"article_date"`
	ArticleMp               string  `orm:"article_mp" json:"article_mp"`
	ArticleAuthor           string  `orm:"article_author" json:"article_author"`
	ArticleReadnumV         int     `orm:"article_readnum_v" json:"article_readnum_v"`
	ArticleReadnum          int     `orm:"article_readnum" json:"article_readnum"`
	ArticleLikenumV         int     `orm:"article_likenum_v" json:"article_likenum_v"`
	ArticleLikenum          int     `orm:"article_likenum" json:"article_likenum"`
	ArticleLinkurl          string  `orm:"article_linkurl" json:"article_linkurl"`
	ArticleRuleDaynum       int     `orm:"article_rule_daynum" json:"article_rule_daynum"`
	ArticleRuleAllnum       int     `orm:"article_rule_allnum" json:"article_rule_allnum"`
	ArticleRuleCredit       int     `orm:"article_rule_credit" json:"article_rule_credit"`
	ArticleRuleMoney        float64 `orm:"article_rule_money" json:"article_rule_money"`
	PageSetOptionNocopy     int     `orm:"page_set_option_nocopy" json:"page_set_option_nocopy"`
	PageSetOptionNoshareTl  int     `orm:"page_set_option_noshare_tl" json:"page_set_option_noshare_tl"`
	PageSetOptionNoshareMsg int     `orm:"page_set_option_noshare_msg" json:"page_set_option_noshare_msg"`
	ArticleKeyword          string  `orm:"article_keyword" json:"article_keyword"`
	ArticleReport           int     `orm:"article_report" json:"article_report"`
	ProductAdvsType         int     `orm:"product_advs_type" json:"product_advs_type"`
	ProductAdvsTitle        string  `orm:"product_advs_title" json:"product_advs_title"`
	ProductAdvsMore         string  `orm:"product_advs_more" json:"product_advs_more"`
	ProductAdvsLink         string  `orm:"product_advs_link" json:"product_advs_link"`
	ProductAdvs             string  `orm:"product_advs" json:"product_advs"`
	ArticleState            int     `orm:"article_state" json:"article_state"`
	NetworkAttachment       string  `orm:"network_attachment" json:"network_attachment"`
	Uniacid                 int     `orm:"uniacid" json:"uniacid"`
	ArticleRuleCredittotal  int     `orm:"article_rule_credittotal" json:"article_rule_credittotal"`
	ArticleRuleMoneytotal   float64 `orm:"article_rule_moneytotal" json:"article_rule_moneytotal"`
	ArticleRuleCredit2      int     `orm:"article_rule_credit2" json:"article_rule_credit2"`
	ArticleRuleMoney2       float64 `orm:"article_rule_money2" json:"article_rule_money2"`
	ArticleRuleCreditm      int     `orm:"article_rule_creditm" json:"article_rule_creditm"`
	ArticleRuleMoneym       float64 `orm:"article_rule_moneym" json:"article_rule_moneym"`
	ArticleRuleCreditm2     int     `orm:"article_rule_creditm2" json:"article_rule_creditm2"`
	ArticleRuleMoneym2      float64 `orm:"article_rule_moneym2" json:"article_rule_moneym2"`
	ArticleReadtime         int     `orm:"article_readtime" json:"article_readtime"`
	ArticleAreas            string  `orm:"article_areas" json:"article_areas"`
	ArticleEndtime          int     `orm:"article_endtime" json:"article_endtime"`
	ArticleHasendtime       int     `orm:"article_hasendtime" json:"article_hasendtime"`
	Displayorder            int     `orm:"displayorder" json:"displayorder"`
	ArticleKeyword2         string  `orm:"article_keyword2" json:"article_keyword2"`
	ArticleAdvance          int     `orm:"article_advance" json:"article_advance"`
	ArticleVirtualadd       int     `orm:"article_virtualadd" json:"article_virtualadd"`
	ArticleVisit            int     `orm:"article_visit" json:"article_visit"`
	ArticleVisitLevel       string  `orm:"article_visit_level" json:"article_visit_level"`
	ArticleVisitTip         string  `orm:"article_visit_tip" json:"article_visit_tip"`
}

func (*Ewei_shop_article) TableName() string {
	return "ewei_shop_article"
}

type Ewei_shop_author_billo struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Billid     int     `orm:"billid" json:"billid"`
	Authorid   int     `orm:"authorid" json:"authorid"`
	Orderid    string  `orm:"orderid" json:"orderid"`
	Ordermoney float64 `orm:"ordermoney" json:"ordermoney"`
}

func (*Ewei_shop_author_billo) TableName() string {
	return "ewei_shop_author_billo"
}

type Ewei_shop_commission_level struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Levelname       string  `orm:"levelname" json:"levelname"`
	Commission1     float64 `orm:"commission1" json:"commission1"`
	Commission2     float64 `orm:"commission2" json:"commission2"`
	Commission3     float64 `orm:"commission3" json:"commission3"`
	Commissionmoney float64 `orm:"commissionmoney" json:"commissionmoney"`
	Ordermoney      float64 `orm:"ordermoney" json:"ordermoney"`
	Downcount       string  `orm:"downcount" json:"downcount"`
	Ordercount      int     `orm:"ordercount" json:"ordercount"`
	Withdraw        float64 `orm:"withdraw" json:"withdraw"`
	Repurchase      float64 `orm:"repurchase" json:"repurchase"`
	Goodsids        string  `orm:"goodsids" json:"goodsids"`
}

func (*Ewei_shop_commission_level) TableName() string {
	return "ewei_shop_commission_level"
}

type Ewei_shop_groups_paylog struct {
	Plid        int     `orm:"plid" json:"plid"`
	Type        string  `orm:"type" json:"type"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Acid        int     `orm:"acid" json:"acid"`
	Openid      string  `orm:"openid" json:"openid"`
	Tid         string  `orm:"tid" json:"tid"`
	Credit      int     `orm:"credit" json:"credit"`
	Creditmoney float64 `orm:"creditmoney" json:"creditmoney"`
	Fee         float64 `orm:"fee" json:"fee"`
	Status      int     `orm:"status" json:"status"`
	Module      string  `orm:"module" json:"module"`
	Tag         string  `orm:"tag" json:"tag"`
	IsUsecard   int     `orm:"is_usecard" json:"is_usecard"`
	CardType    int     `orm:"card_type" json:"card_type"`
	CardId      string  `orm:"card_id" json:"card_id"`
	CardFee     float64 `orm:"card_fee" json:"card_fee"`
	EncryptCode string  `orm:"encrypt_code" json:"encrypt_code"`
	Uniontid    string  `orm:"uniontid" json:"uniontid"`
}

func (*Ewei_shop_groups_paylog) TableName() string {
	return "ewei_shop_groups_paylog"
}

type Ewei_shop_member_mergelog struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Mergetime   int    `orm:"mergetime" json:"mergetime"`
	OpenidA     string `orm:"openid_a" json:"openid_a"`
	OpenidB     string `orm:"openid_b" json:"openid_b"`
	MidA        int    `orm:"mid_a" json:"mid_a"`
	MidB        int    `orm:"mid_b" json:"mid_b"`
	DetailA     string `orm:"detail_a" json:"detail_a"`
	DetailB     string `orm:"detail_b" json:"detail_b"`
	DetailC     string `orm:"detail_c" json:"detail_c"`
	Fromuniacid int    `orm:"fromuniacid" json:"fromuniacid"`
}

func (*Ewei_shop_member_mergelog) TableName() string {
	return "ewei_shop_member_mergelog"
}

type Ewei_shop_member_rank struct {
	Id      int `orm:"id" json:"id"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Status  int `orm:"status" json:"status"`
	Num     int `orm:"num" json:"num"`
}

func (*Ewei_shop_member_rank) TableName() string {
	return "ewei_shop_member_rank"
}

type Ewei_shop_virtual_type struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Cate     int    `orm:"cate" json:"cate"`
	Title    string `orm:"title" json:"title"`
	Fields   string `orm:"fields" json:"fields"`
	Usedata  int    `orm:"usedata" json:"usedata"`
	Alldata  int    `orm:"alldata" json:"alldata"`
	Merchid  int    `orm:"merchid" json:"merchid"`
	Linktext string `orm:"linktext" json:"linktext"`
	Linkurl  string `orm:"linkurl" json:"linkurl"`
	Recycled int    `orm:"recycled" json:"recycled"`
}

func (*Ewei_shop_virtual_type) TableName() string {
	return "ewei_shop_virtual_type"
}

type Mc_handsel struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Touid       int    `orm:"touid" json:"touid"`
	Fromuid     string `orm:"fromuid" json:"fromuid"`
	Module      string `orm:"module" json:"module"`
	Sign        string `orm:"sign" json:"sign"`
	Action      string `orm:"action" json:"action"`
	CreditValue int    `orm:"credit_value" json:"credit_value"`
	Createtime  int    `orm:"createtime" json:"createtime"`
}

func (*Mc_handsel) TableName() string {
	return "mc_handsel"
}

type News_reply struct {
	Id           int    `orm:"id" json:"id"`
	Rid          int    `orm:"rid" json:"rid"`
	ParentId     int    `orm:"parent_id" json:"parent_id"`
	Title        string `orm:"title" json:"title"`
	Author       string `orm:"author" json:"author"`
	Description  string `orm:"description" json:"description"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Content      string `orm:"content" json:"content"`
	Url          string `orm:"url" json:"url"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Incontent    int    `orm:"incontent" json:"incontent"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	MediaId      string `orm:"media_id" json:"media_id"`
}

func (*News_reply) TableName() string {
	return "news_reply"
}

type Ewei_shop_diypage_template_category struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
	Merch   int    `orm:"merch" json:"merch"`
}

func (*Ewei_shop_diypage_template_category) TableName() string {
	return "ewei_shop_diypage_template_category"
}

type Ewei_shop_groups_goods_atlas struct {
	Id    int    `orm:"id" json:"id"`
	GId   int    `orm:"g_id" json:"g_id"`
	Thumb string `orm:"thumb" json:"thumb"`
}

func (*Ewei_shop_groups_goods_atlas) TableName() string {
	return "ewei_shop_groups_goods_atlas"
}

type Ewei_shop_sns_like struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Openid  string `orm:"openid" json:"openid"`
	Pid     int    `orm:"pid" json:"pid"`
}

func (*Ewei_shop_sns_like) TableName() string {
	return "ewei_shop_sns_like"
}

type Mc_fans_groups struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Acid    int    `orm:"acid" json:"acid"`
	Groups  string `orm:"groups" json:"groups"`
}

func (*Mc_fans_groups) TableName() string {
	return "mc_fans_groups"
}

type Modules_bindings struct {
	Eid          int    `orm:"eid" json:"eid"`
	Module       string `orm:"module" json:"module"`
	Entry        string `orm:"entry" json:"entry"`
	Call         string `orm:"call" json:"call"`
	Title        string `orm:"title" json:"title"`
	Do           string `orm:"do" json:"do"`
	State        string `orm:"state" json:"state"`
	Direct       int    `orm:"direct" json:"direct"`
	Url          string `orm:"url" json:"url"`
	Icon         string `orm:"icon" json:"icon"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Modules_bindings) TableName() string {
	return "modules_bindings"
}

type Ewei_shop_coupon_taskdata struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Openid        string `orm:"openid" json:"openid"`
	Taskid        int    `orm:"taskid" json:"taskid"`
	Couponid      int    `orm:"couponid" json:"couponid"`
	Sendnum       int    `orm:"sendnum" json:"sendnum"`
	Tasktype      int    `orm:"tasktype" json:"tasktype"`
	Orderid       int    `orm:"orderid" json:"orderid"`
	Parentorderid int    `orm:"parentorderid" json:"parentorderid"`
	Createtime    int    `orm:"createtime" json:"createtime"`
	Status        int    `orm:"status" json:"status"`
	Sendpoint     int    `orm:"sendpoint" json:"sendpoint"`
}

func (*Ewei_shop_coupon_taskdata) TableName() string {
	return "ewei_shop_coupon_taskdata"
}

type Ewei_shop_fullback_goods struct {
	Id                     int     `orm:"id" json:"id"`
	Uniacid                int     `orm:"uniacid" json:"uniacid"`
	Type                   int     `orm:"type" json:"type"`
	Goodsid                int     `orm:"goodsid" json:"goodsid"`
	Titles                 string  `orm:"titles" json:"titles"`
	Thumb                  string  `orm:"thumb" json:"thumb"`
	Marketprice            float64 `orm:"marketprice" json:"marketprice"`
	Minallfullbackallprice float64 `orm:"minallfullbackallprice" json:"minallfullbackallprice"`
	Maxallfullbackallprice float64 `orm:"maxallfullbackallprice" json:"maxallfullbackallprice"`
	Minallfullbackallratio float64 `orm:"minallfullbackallratio" json:"minallfullbackallratio"`
	Maxallfullbackallratio float64 `orm:"maxallfullbackallratio" json:"maxallfullbackallratio"`
	Day                    int     `orm:"day" json:"day"`
	Fullbackprice          float64 `orm:"fullbackprice" json:"fullbackprice"`
	Fullbackratio          float64 `orm:"fullbackratio" json:"fullbackratio"`
	Status                 int     `orm:"status" json:"status"`
	Displayorder           int     `orm:"displayorder" json:"displayorder"`
	Hasoption              int     `orm:"hasoption" json:"hasoption"`
	Optionid               string  `orm:"optionid" json:"optionid"`
	Startday               int     `orm:"startday" json:"startday"`
	Refund                 int     `orm:"refund" json:"refund"`
}

func (*Ewei_shop_fullback_goods) TableName() string {
	return "ewei_shop_fullback_goods"
}

type Ewei_shop_system_company_category struct {
	Id           int    `orm:"id" json:"id"`
	Name         string `orm:"name" json:"name"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_company_category) TableName() string {
	return "ewei_shop_system_company_category"
}

type Mc_member_fields struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Fieldid      int    `orm:"fieldid" json:"fieldid"`
	Title        string `orm:"title" json:"title"`
	Available    int    `orm:"available" json:"available"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Mc_member_fields) TableName() string {
	return "mc_member_fields"
}

type Menu_event struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Keyword    string `orm:"keyword" json:"keyword"`
	Type       string `orm:"type" json:"type"`
	Picmd5     string `orm:"picmd5" json:"picmd5"`
	Openid     string `orm:"openid" json:"openid"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Menu_event) TableName() string {
	return "menu_event"
}

type Activity_clerk_menu struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Pid          int    `orm:"pid" json:"pid"`
	GroupName    string `orm:"group_name" json:"group_name"`
	Title        string `orm:"title" json:"title"`
	Icon         string `orm:"icon" json:"icon"`
	Url          string `orm:"url" json:"url"`
	Type         string `orm:"type" json:"type"`
	Permission   string `orm:"permission" json:"permission"`
	System       int    `orm:"system" json:"system"`
}

func (*Activity_clerk_menu) TableName() string {
	return "activity_clerk_menu"
}

type Ewei_shop_package_goods_option struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Goodsid      int     `orm:"goodsid" json:"goodsid"`
	Optionid     int     `orm:"optionid" json:"optionid"`
	Pid          int     `orm:"pid" json:"pid"`
	Title        string  `orm:"title" json:"title"`
	Packageprice float64 `orm:"packageprice" json:"packageprice"`
	Marketprice  float64 `orm:"marketprice" json:"marketprice"`
	Commission1  float64 `orm:"commission1" json:"commission1"`
	Commission2  float64 `orm:"commission2" json:"commission2"`
	Commission3  float64 `orm:"commission3" json:"commission3"`
}

func (*Ewei_shop_package_goods_option) TableName() string {
	return "ewei_shop_package_goods_option"
}

type Ewei_shop_system_plugingrant_setting struct {
	Id         int    `orm:"id" json:"id"`
	Com        string `orm:"com" json:"com"`
	Adv        string `orm:"adv" json:"adv"`
	Plugin     string `orm:"plugin" json:"plugin"`
	Customer   string `orm:"customer" json:"customer"`
	Contact    string `orm:"contact" json:"contact"`
	Servertime string `orm:"servertime" json:"servertime"`
	Weixin     int    `orm:"weixin" json:"weixin"`
	Appid      string `orm:"appid" json:"appid"`
	Mchid      string `orm:"mchid" json:"mchid"`
	Apikey     string `orm:"apikey" json:"apikey"`
	Alipay     int    `orm:"alipay" json:"alipay"`
	Account    string `orm:"account" json:"account"`
	Partner    string `orm:"partner" json:"partner"`
	Secret     string `orm:"secret" json:"secret"`
}

func (*Ewei_shop_system_plugingrant_setting) TableName() string {
	return "ewei_shop_system_plugingrant_setting"
}

type Ewei_shop_virtual_category struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
	Merchid int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_virtual_category) TableName() string {
	return "ewei_shop_virtual_category"
}

type Account_webapp struct {
	Acid    int    `orm:"acid" json:"acid"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
}

func (*Account_webapp) TableName() string {
	return "account_webapp"
}

type Coupon_activity struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	MsgId       int    `orm:"msg_id" json:"msg_id"`
	Status      int    `orm:"status" json:"status"`
	Title       string `orm:"title" json:"title"`
	Type        int    `orm:"type" json:"type"`
	Thumb       string `orm:"thumb" json:"thumb"`
	Coupons     string `orm:"coupons" json:"coupons"`
	Description string `orm:"description" json:"description"`
	Members     string `orm:"members" json:"members"`
}

func (*Coupon_activity) TableName() string {
	return "coupon_activity"
}

type Ewei_shop_commission_bank struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Bankname     string `orm:"bankname" json:"bankname"`
	Content      string `orm:"content" json:"content"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_commission_bank) TableName() string {
	return "ewei_shop_commission_bank"
}

type Ewei_shop_funbar struct {
	Id      int    `orm:"id" json:"id"`
	Uid     int    `orm:"uid" json:"uid"`
	Datas   string `orm:"datas" json:"datas"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
}

func (*Ewei_shop_funbar) TableName() string {
	return "ewei_shop_funbar"
}

type Ewei_shop_refund_address struct {
	Id        int    `orm:"id" json:"id"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Title     string `orm:"title" json:"title"`
	Name      string `orm:"name" json:"name"`
	Tel       string `orm:"tel" json:"tel"`
	Mobile    string `orm:"mobile" json:"mobile"`
	Province  string `orm:"province" json:"province"`
	City      string `orm:"city" json:"city"`
	Area      string `orm:"area" json:"area"`
	Address   string `orm:"address" json:"address"`
	Isdefault int    `orm:"isdefault" json:"isdefault"`
	Zipcode   string `orm:"zipcode" json:"zipcode"`
	Content   string `orm:"content" json:"content"`
	Deleted   int    `orm:"deleted" json:"deleted"`
	Openid    string `orm:"openid" json:"openid"`
	Merchid   int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_refund_address) TableName() string {
	return "ewei_shop_refund_address"
}

type Ewei_shop_sms struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Name     string `orm:"name" json:"name"`
	Type     string `orm:"type" json:"type"`
	Template int    `orm:"template" json:"template"`
	Smstplid string `orm:"smstplid" json:"smstplid"`
	Smssign  string `orm:"smssign" json:"smssign"`
	Content  string `orm:"content" json:"content"`
	Data     string `orm:"data" json:"data"`
	Status   int    `orm:"status" json:"status"`
}

func (*Ewei_shop_sms) TableName() string {
	return "ewei_shop_sms"
}

type Modules struct {
	Mid            int    `orm:"mid" json:"mid"`
	Name           string `orm:"name" json:"name"`
	Type           string `orm:"type" json:"type"`
	Title          string `orm:"title" json:"title"`
	Version        string `orm:"version" json:"version"`
	Ability        string `orm:"ability" json:"ability"`
	Description    string `orm:"description" json:"description"`
	Author         string `orm:"author" json:"author"`
	Url            string `orm:"url" json:"url"`
	Settings       int    `orm:"settings" json:"settings"`
	Subscribes     string `orm:"subscribes" json:"subscribes"`
	Handles        string `orm:"handles" json:"handles"`
	Isrulefields   int    `orm:"isrulefields" json:"isrulefields"`
	Issystem       int    `orm:"issystem" json:"issystem"`
	Target         int    `orm:"target" json:"target"`
	Iscard         int    `orm:"iscard" json:"iscard"`
	Permissions    string `orm:"permissions" json:"permissions"`
	TitleInitial   string `orm:"title_initial" json:"title_initial"`
	WxappSupport   int    `orm:"wxapp_support" json:"wxapp_support"`
	AppSupport     int    `orm:"app_support" json:"app_support"`
	WelcomeSupport int    `orm:"welcome_support" json:"welcome_support"`
	OauthType      int    `orm:"oauth_type" json:"oauth_type"`
	WebappSupport  int    `orm:"webapp_support" json:"webapp_support"`
}

func (*Modules) TableName() string {
	return "modules"
}

type Music_reply struct {
	Id          int    `orm:"id" json:"id"`
	Rid         int    `orm:"rid" json:"rid"`
	Title       string `orm:"title" json:"title"`
	Description string `orm:"description" json:"description"`
	Url         string `orm:"url" json:"url"`
	Hqurl       string `orm:"hqurl" json:"hqurl"`
}

func (*Music_reply) TableName() string {
	return "music_reply"
}

type Ewei_shop_coupon struct {
	Id                 int     `orm:"id" json:"id"`
	Uniacid            int     `orm:"uniacid" json:"uniacid"`
	Catid              int     `orm:"catid" json:"catid"`
	Couponname         string  `orm:"couponname" json:"couponname"`
	Gettype            int     `orm:"gettype" json:"gettype"`
	Getmax             int     `orm:"getmax" json:"getmax"`
	Usetype            int     `orm:"usetype" json:"usetype"`
	Returntype         int     `orm:"returntype" json:"returntype"`
	Bgcolor            string  `orm:"bgcolor" json:"bgcolor"`
	Enough             float64 `orm:"enough" json:"enough"`
	Timelimit          int     `orm:"timelimit" json:"timelimit"`
	Coupontype         int     `orm:"coupontype" json:"coupontype"`
	Timedays           int     `orm:"timedays" json:"timedays"`
	Timestart          int     `orm:"timestart" json:"timestart"`
	Timeend            int     `orm:"timeend" json:"timeend"`
	Discount           float64 `orm:"discount" json:"discount"`
	Deduct             float64 `orm:"deduct" json:"deduct"`
	Backtype           int     `orm:"backtype" json:"backtype"`
	Backmoney          string  `orm:"backmoney" json:"backmoney"`
	Backcredit         string  `orm:"backcredit" json:"backcredit"`
	Backredpack        string  `orm:"backredpack" json:"backredpack"`
	Backwhen           int     `orm:"backwhen" json:"backwhen"`
	Thumb              string  `orm:"thumb" json:"thumb"`
	Desc               string  `orm:"desc" json:"desc"`
	Createtime         int     `orm:"createtime" json:"createtime"`
	Total              int     `orm:"total" json:"total"`
	Status             int     `orm:"status" json:"status"`
	Money              float64 `orm:"money" json:"money"`
	Respdesc           string  `orm:"respdesc" json:"respdesc"`
	Respthumb          string  `orm:"respthumb" json:"respthumb"`
	Resptitle          string  `orm:"resptitle" json:"resptitle"`
	Respurl            string  `orm:"respurl" json:"respurl"`
	Credit             int     `orm:"credit" json:"credit"`
	Usecredit2         int     `orm:"usecredit2" json:"usecredit2"`
	Remark             string  `orm:"remark" json:"remark"`
	Descnoset          int     `orm:"descnoset" json:"descnoset"`
	Pwdkey             string  `orm:"pwdkey" json:"pwdkey"`
	Pwdsuc             string  `orm:"pwdsuc" json:"pwdsuc"`
	Pwdfail            string  `orm:"pwdfail" json:"pwdfail"`
	Pwdurl             string  `orm:"pwdurl" json:"pwdurl"`
	Pwdask             string  `orm:"pwdask" json:"pwdask"`
	Pwdstatus          int     `orm:"pwdstatus" json:"pwdstatus"`
	Pwdtimes           int     `orm:"pwdtimes" json:"pwdtimes"`
	Pwdfull            string  `orm:"pwdfull" json:"pwdfull"`
	Pwdwords           string  `orm:"pwdwords" json:"pwdwords"`
	Pwdopen            int     `orm:"pwdopen" json:"pwdopen"`
	Pwdown             string  `orm:"pwdown" json:"pwdown"`
	Pwdexit            string  `orm:"pwdexit" json:"pwdexit"`
	Pwdexitstr         string  `orm:"pwdexitstr" json:"pwdexitstr"`
	Displayorder       int     `orm:"displayorder" json:"displayorder"`
	Pwdkey2            string  `orm:"pwdkey2" json:"pwdkey2"`
	Merchid            int     `orm:"merchid" json:"merchid"`
	Limitgoodtype      int     `orm:"limitgoodtype" json:"limitgoodtype"`
	Limitgoodcatetype  int     `orm:"limitgoodcatetype" json:"limitgoodcatetype"`
	Limitgoodcateids   string  `orm:"limitgoodcateids" json:"limitgoodcateids"`
	Limitgoodids       string  `orm:"limitgoodids" json:"limitgoodids"`
	Islimitlevel       int     `orm:"islimitlevel" json:"islimitlevel"`
	Limitmemberlevels  string  `orm:"limitmemberlevels" json:"limitmemberlevels"`
	Limitagentlevels   string  `orm:"limitagentlevels" json:"limitagentlevels"`
	Limitpartnerlevels string  `orm:"limitpartnerlevels" json:"limitpartnerlevels"`
	Limitaagentlevels  string  `orm:"limitaagentlevels" json:"limitaagentlevels"`
	Tagtitle           string  `orm:"tagtitle" json:"tagtitle"`
	Settitlecolor      int     `orm:"settitlecolor" json:"settitlecolor"`
	Titlecolor         string  `orm:"titlecolor" json:"titlecolor"`
	Limitdiscounttype  int     `orm:"limitdiscounttype" json:"limitdiscounttype"`
	Quickget           int     `orm:"quickget" json:"quickget"`
}

func (*Ewei_shop_coupon) TableName() string {
	return "ewei_shop_coupon"
}

type Ewei_shop_dispatch struct {
	Id                  int     `orm:"id" json:"id"`
	Uniacid             int     `orm:"uniacid" json:"uniacid"`
	Dispatchname        string  `orm:"dispatchname" json:"dispatchname"`
	Dispatchtype        int     `orm:"dispatchtype" json:"dispatchtype"`
	Displayorder        int     `orm:"displayorder" json:"displayorder"`
	Firstprice          float64 `orm:"firstprice" json:"firstprice"`
	Secondprice         float64 `orm:"secondprice" json:"secondprice"`
	Firstweight         int     `orm:"firstweight" json:"firstweight"`
	Secondweight        int     `orm:"secondweight" json:"secondweight"`
	Express             string  `orm:"express" json:"express"`
	Areas               string  `orm:"areas" json:"areas"`
	Carriers            string  `orm:"carriers" json:"carriers"`
	Enabled             int     `orm:"enabled" json:"enabled"`
	Calculatetype       int     `orm:"calculatetype" json:"calculatetype"`
	Firstnum            int     `orm:"firstnum" json:"firstnum"`
	Secondnum           int     `orm:"secondnum" json:"secondnum"`
	Firstnumprice       float64 `orm:"firstnumprice" json:"firstnumprice"`
	Secondnumprice      float64 `orm:"secondnumprice" json:"secondnumprice"`
	Isdefault           int     `orm:"isdefault" json:"isdefault"`
	Shopid              int     `orm:"shopid" json:"shopid"`
	Merchid             int     `orm:"merchid" json:"merchid"`
	Nodispatchareas     string  `orm:"nodispatchareas" json:"nodispatchareas"`
	NodispatchareasCode string  `orm:"nodispatchareas_code" json:"nodispatchareas_code"`
	Isdispatcharea      int     `orm:"isdispatcharea" json:"isdispatcharea"`
	Freeprice           float64 `orm:"freeprice" json:"freeprice"`
}

func (*Ewei_shop_dispatch) TableName() string {
	return "ewei_shop_dispatch"
}

type Ewei_shop_diyform_type struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Cate     int    `orm:"cate" json:"cate"`
	Title    string `orm:"title" json:"title"`
	Fields   string `orm:"fields" json:"fields"`
	Usedata  int    `orm:"usedata" json:"usedata"`
	Alldata  int    `orm:"alldata" json:"alldata"`
	Status   int    `orm:"status" json:"status"`
	Savedata int    `orm:"savedata" json:"savedata"`
}

func (*Ewei_shop_diyform_type) TableName() string {
	return "ewei_shop_diyform_type"
}

type Ewei_shop_merch_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_adv) TableName() string {
	return "ewei_shop_merch_adv"
}

type Ewei_shop_qa_set struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Showmember int    `orm:"showmember" json:"showmember"`
	Showtype   int    `orm:"showtype" json:"showtype"`
	Keyword    string `orm:"keyword" json:"keyword"`
	EnterTitle string `orm:"enter_title" json:"enter_title"`
	EnterImg   string `orm:"enter_img" json:"enter_img"`
	EnterDesc  string `orm:"enter_desc" json:"enter_desc"`
	Share      int    `orm:"share" json:"share"`
}

func (*Ewei_shop_qa_set) TableName() string {
	return "ewei_shop_qa_set"
}

type Mc_mapping_fans struct {
	Fanid        int    `orm:"fanid" json:"fanid"`
	Acid         int    `orm:"acid" json:"acid"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Uid          int    `orm:"uid" json:"uid"`
	Openid       string `orm:"openid" json:"openid"`
	Nickname     string `orm:"nickname" json:"nickname"`
	Groupid      string `orm:"groupid" json:"groupid"`
	Salt         string `orm:"salt" json:"salt"`
	Follow       int    `orm:"follow" json:"follow"`
	Followtime   int    `orm:"followtime" json:"followtime"`
	Unfollowtime int    `orm:"unfollowtime" json:"unfollowtime"`
	Tag          string `orm:"tag" json:"tag"`
	Updatetime   int    `orm:"updatetime" json:"updatetime"`
	Unionid      string `orm:"unionid" json:"unionid"`
}

func (*Mc_mapping_fans) TableName() string {
	return "mc_mapping_fans"
}

type Article_news struct {
	Id           int    `orm:"id" json:"id"`
	Cateid       int    `orm:"cateid" json:"cateid"`
	Title        string `orm:"title" json:"title"`
	Content      string `orm:"content" json:"content"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Source       string `orm:"source" json:"source"`
	Author       string `orm:"author" json:"author"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	IsDisplay    int    `orm:"is_display" json:"is_display"`
	IsShowHome   int    `orm:"is_show_home" json:"is_show_home"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Click        int    `orm:"click" json:"click"`
}

func (*Article_news) TableName() string {
	return "article_news"
}

type Ewei_shop_coupon_usesendtasks struct {
	Id          int `orm:"id" json:"id"`
	Uniacid     int `orm:"uniacid" json:"uniacid"`
	Usecouponid int `orm:"usecouponid" json:"usecouponid"`
	Couponid    int `orm:"couponid" json:"couponid"`
	Starttime   int `orm:"starttime" json:"starttime"`
	Endtime     int `orm:"endtime" json:"endtime"`
	Sendnum     int `orm:"sendnum" json:"sendnum"`
	Num         int `orm:"num" json:"num"`
	Status      int `orm:"status" json:"status"`
}

func (*Ewei_shop_coupon_usesendtasks) TableName() string {
	return "ewei_shop_coupon_usesendtasks"
}

type Ewei_shop_diyform_data struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Typeid        int    `orm:"typeid" json:"typeid"`
	Cid           int    `orm:"cid" json:"cid"`
	Diyformfields string `orm:"diyformfields" json:"diyformfields"`
	Fields        string `orm:"fields" json:"fields"`
	Openid        string `orm:"openid" json:"openid"`
	Type          int    `orm:"type" json:"type"`
}

func (*Ewei_shop_diyform_data) TableName() string {
	return "ewei_shop_diyform_data"
}

type Ewei_shop_groups_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Advimg       string `orm:"advimg" json:"advimg"`
	Advurl       string `orm:"advurl" json:"advurl"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
}

func (*Ewei_shop_groups_category) TableName() string {
	return "ewei_shop_groups_category"
}

type Ewei_shop_lottery struct {
	LotteryId     int    `orm:"lottery_id" json:"lottery_id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	LotteryTitle  string `orm:"lottery_title" json:"lottery_title"`
	LotteryIcon   string `orm:"lottery_icon" json:"lottery_icon"`
	LotteryBanner string `orm:"lottery_banner" json:"lottery_banner"`
	LotteryCannot string `orm:"lottery_cannot" json:"lottery_cannot"`
	LotteryType   int    `orm:"lottery_type" json:"lottery_type"`
	IsDelete      int    `orm:"is_delete" json:"is_delete"`
	Addtime       int    `orm:"addtime" json:"addtime"`
	LotteryData   string `orm:"lottery_data" json:"lottery_data"`
	IsGoods       int    `orm:"is_goods" json:"is_goods"`
	LotteryDays   int    `orm:"lottery_days" json:"lottery_days"`
	TaskType      int    `orm:"task_type" json:"task_type"`
	TaskData      string `orm:"task_data" json:"task_data"`
	StartTime     int    `orm:"start_time" json:"start_time"`
	EndTime       int    `orm:"end_time" json:"end_time"`
	AwardStart    int    `orm:"award_start" json:"award_start"`
	AwardEnd      int    `orm:"award_end" json:"award_end"`
}

func (*Ewei_shop_lottery) TableName() string {
	return "ewei_shop_lottery"
}

type Mc_member_address struct {
	Id        int    `orm:"id" json:"id"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Uid       int    `orm:"uid" json:"uid"`
	Username  string `orm:"username" json:"username"`
	Mobile    string `orm:"mobile" json:"mobile"`
	Zipcode   string `orm:"zipcode" json:"zipcode"`
	Province  string `orm:"province" json:"province"`
	City      string `orm:"city" json:"city"`
	District  string `orm:"district" json:"district"`
	Address   string `orm:"address" json:"address"`
	Isdefault int    `orm:"isdefault" json:"isdefault"`
}

func (*Mc_member_address) TableName() string {
	return "mc_member_address"
}

type Core_cron_record struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Module     string `orm:"module" json:"module"`
	Type       string `orm:"type" json:"type"`
	Tid        int    `orm:"tid" json:"tid"`
	Note       string `orm:"note" json:"note"`
	Tag        string `orm:"tag" json:"tag"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Core_cron_record) TableName() string {
	return "core_cron_record"
}

type Custom_reply struct {
	Id     int `orm:"id" json:"id"`
	Rid    int `orm:"rid" json:"rid"`
	Start1 int `orm:"start1" json:"start1"`
	End1   int `orm:"end1" json:"end1"`
	Start2 int `orm:"start2" json:"start2"`
	End2   int `orm:"end2" json:"end2"`
}

func (*Custom_reply) TableName() string {
	return "custom_reply"
}

type Ewei_shop_coupon_sendshow struct {
	Id           int    `orm:"id" json:"id"`
	Showkey      string `orm:"showkey" json:"showkey"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Openid       string `orm:"openid" json:"openid"`
	Coupondataid int    `orm:"coupondataid" json:"coupondataid"`
}

func (*Ewei_shop_coupon_sendshow) TableName() string {
	return "ewei_shop_coupon_sendshow"
}

type Ewei_shop_creditshop_goods struct {
	Id               int     `orm:"id" json:"id"`
	Uniacid          int     `orm:"uniacid" json:"uniacid"`
	Displayorder     int     `orm:"displayorder" json:"displayorder"`
	Title            string  `orm:"title" json:"title"`
	Cate             int     `orm:"cate" json:"cate"`
	Thumb            string  `orm:"thumb" json:"thumb"`
	Price            float64 `orm:"price" json:"price"`
	Type             int     `orm:"type" json:"type"`
	Credit           int     `orm:"credit" json:"credit"`
	Money            float64 `orm:"money" json:"money"`
	Total            int     `orm:"total" json:"total"`
	Totalday         int     `orm:"totalday" json:"totalday"`
	Chance           int     `orm:"chance" json:"chance"`
	Chanceday        int     `orm:"chanceday" json:"chanceday"`
	Detail           string  `orm:"detail" json:"detail"`
	Rate1            int     `orm:"rate1" json:"rate1"`
	Rate2            int     `orm:"rate2" json:"rate2"`
	Endtime          int     `orm:"endtime" json:"endtime"`
	Joins            int     `orm:"joins" json:"joins"`
	Views            int     `orm:"views" json:"views"`
	Createtime       int     `orm:"createtime" json:"createtime"`
	Status           int     `orm:"status" json:"status"`
	Deleted          int     `orm:"deleted" json:"deleted"`
	Showlevels       string  `orm:"showlevels" json:"showlevels"`
	Buylevels        string  `orm:"buylevels" json:"buylevels"`
	Showgroups       string  `orm:"showgroups" json:"showgroups"`
	Buygroups        string  `orm:"buygroups" json:"buygroups"`
	Vip              int     `orm:"vip" json:"vip"`
	Istop            int     `orm:"istop" json:"istop"`
	Isrecommand      int     `orm:"isrecommand" json:"isrecommand"`
	Istime           int     `orm:"istime" json:"istime"`
	Timestart        int     `orm:"timestart" json:"timestart"`
	Timeend          int     `orm:"timeend" json:"timeend"`
	ShareTitle       string  `orm:"share_title" json:"share_title"`
	ShareIcon        string  `orm:"share_icon" json:"share_icon"`
	ShareDesc        string  `orm:"share_desc" json:"share_desc"`
	Followneed       int     `orm:"followneed" json:"followneed"`
	Followtext       string  `orm:"followtext" json:"followtext"`
	Subtitle         string  `orm:"subtitle" json:"subtitle"`
	Subdetail        string  `orm:"subdetail" json:"subdetail"`
	Noticedetail     string  `orm:"noticedetail" json:"noticedetail"`
	Usedetail        string  `orm:"usedetail" json:"usedetail"`
	Goodsdetail      string  `orm:"goodsdetail" json:"goodsdetail"`
	Isendtime        int     `orm:"isendtime" json:"isendtime"`
	Usecredit2       int     `orm:"usecredit2" json:"usecredit2"`
	Area             string  `orm:"area" json:"area"`
	Dispatch         float64 `orm:"dispatch" json:"dispatch"`
	Storeids         string  `orm:"storeids" json:"storeids"`
	Noticeopenid     string  `orm:"noticeopenid" json:"noticeopenid"`
	Noticetype       int     `orm:"noticetype" json:"noticetype"`
	Isverify         int     `orm:"isverify" json:"isverify"`
	Goodstype        int     `orm:"goodstype" json:"goodstype"`
	Couponid         int     `orm:"couponid" json:"couponid"`
	Goodsid          int     `orm:"goodsid" json:"goodsid"`
	Merchid          int     `orm:"merchid" json:"merchid"`
	Productprice     float64 `orm:"productprice" json:"productprice"`
	Mincredit        float64 `orm:"mincredit" json:"mincredit"`
	Minmoney         float64 `orm:"minmoney" json:"minmoney"`
	Maxcredit        float64 `orm:"maxcredit" json:"maxcredit"`
	Maxmoney         float64 `orm:"maxmoney" json:"maxmoney"`
	Dispatchtype     int     `orm:"dispatchtype" json:"dispatchtype"`
	Dispatchid       int     `orm:"dispatchid" json:"dispatchid"`
	Verifytype       int     `orm:"verifytype" json:"verifytype"`
	Verifynum        int     `orm:"verifynum" json:"verifynum"`
	Grant1           float64 `orm:"grant1" json:"grant1"`
	Grant2           float64 `orm:"grant2" json:"grant2"`
	Goodssn          string  `orm:"goodssn" json:"goodssn"`
	Productsn        string  `orm:"productsn" json:"productsn"`
	Weight           int     `orm:"weight" json:"weight"`
	Showtotal        int     `orm:"showtotal" json:"showtotal"`
	Totalcnf         int     `orm:"totalcnf" json:"totalcnf"`
	Usetime          int     `orm:"usetime" json:"usetime"`
	Hasoption        int     `orm:"hasoption" json:"hasoption"`
	Noticedetailshow int     `orm:"noticedetailshow" json:"noticedetailshow"`
	Detailshow       int     `orm:"detailshow" json:"detailshow"`
	Packetmoney      float64 `orm:"packetmoney" json:"packetmoney"`
	Surplusmoney     float64 `orm:"surplusmoney" json:"surplusmoney"`
	Packetlimit      float64 `orm:"packetlimit" json:"packetlimit"`
	Packettype       int     `orm:"packettype" json:"packettype"`
	Minpacketmoney   float64 `orm:"minpacketmoney" json:"minpacketmoney"`
	Packettotal      int     `orm:"packettotal" json:"packettotal"`
	Packetsurplus    int     `orm:"packetsurplus" json:"packetsurplus"`
	Maxpacketmoney   float64 `orm:"maxpacketmoney" json:"maxpacketmoney"`
}

func (*Ewei_shop_creditshop_goods) TableName() string {
	return "ewei_shop_creditshop_goods"
}

type Ewei_shop_polyapi_key struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Merchid    int    `orm:"merchid" json:"merchid"`
	Appkey     string `orm:"appkey" json:"appkey"`
	Token      string `orm:"token" json:"token"`
	Appsecret  string `orm:"appsecret" json:"appsecret"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Updatetime int    `orm:"updatetime" json:"updatetime"`
	Status     int    `orm:"status" json:"status"`
}

func (*Ewei_shop_polyapi_key) TableName() string {
	return "ewei_shop_polyapi_key"
}

type Ewei_shop_poster_log struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Openid       string  `orm:"openid" json:"openid"`
	Posterid     int     `orm:"posterid" json:"posterid"`
	FromOpenid   string  `orm:"from_openid" json:"from_openid"`
	Subcredit    int     `orm:"subcredit" json:"subcredit"`
	Submoney     float64 `orm:"submoney" json:"submoney"`
	Reccredit    int     `orm:"reccredit" json:"reccredit"`
	Recmoney     float64 `orm:"recmoney" json:"recmoney"`
	Createtime   int     `orm:"createtime" json:"createtime"`
	Reccouponid  int     `orm:"reccouponid" json:"reccouponid"`
	Reccouponnum int     `orm:"reccouponnum" json:"reccouponnum"`
	Subcouponid  int     `orm:"subcouponid" json:"subcouponid"`
	Subcouponnum int     `orm:"subcouponnum" json:"subcouponnum"`
}

func (*Ewei_shop_poster_log) TableName() string {
	return "ewei_shop_poster_log"
}

type Ewei_shop_qa_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_qa_adv) TableName() string {
	return "ewei_shop_qa_adv"
}

type Ewei_shop_seckill_task_room struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Taskid       int    `orm:"taskid" json:"taskid"`
	Title        string `orm:"title" json:"title"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	PageTitle    string `orm:"page_title" json:"page_title"`
	ShareTitle   string `orm:"share_title" json:"share_title"`
	ShareDesc    string `orm:"share_desc" json:"share_desc"`
	ShareIcon    string `orm:"share_icon" json:"share_icon"`
	Oldshow      int    `orm:"oldshow" json:"oldshow"`
	Tag          string `orm:"tag" json:"tag"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Diypage      int    `orm:"diypage" json:"diypage"`
}

func (*Ewei_shop_seckill_task_room) TableName() string {
	return "ewei_shop_seckill_task_room"
}

type Images_reply struct {
	Id          int    `orm:"id" json:"id"`
	Rid         int    `orm:"rid" json:"rid"`
	Title       string `orm:"title" json:"title"`
	Description string `orm:"description" json:"description"`
	Mediaid     string `orm:"mediaid" json:"mediaid"`
	Createtime  int    `orm:"createtime" json:"createtime"`
}

func (*Images_reply) TableName() string {
	return "images_reply"
}

type Mc_member_property struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Property string `orm:"property" json:"property"`
}

func (*Mc_member_property) TableName() string {
	return "mc_member_property"
}

type Site_styles struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Templateid int    `orm:"templateid" json:"templateid"`
	Name       string `orm:"name" json:"name"`
}

func (*Site_styles) TableName() string {
	return "site_styles"
}

type Stat_fans struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	New      int    `orm:"new" json:"new"`
	Cancel   int    `orm:"cancel" json:"cancel"`
	Cumulate int    `orm:"cumulate" json:"cumulate"`
	Date     string `orm:"date" json:"date"`
}

func (*Stat_fans) TableName() string {
	return "stat_fans"
}

type Coupon_location struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Acid         int    `orm:"acid" json:"acid"`
	Sid          int    `orm:"sid" json:"sid"`
	LocationId   int    `orm:"location_id" json:"location_id"`
	BusinessName string `orm:"business_name" json:"business_name"`
	BranchName   string `orm:"branch_name" json:"branch_name"`
	Category     string `orm:"category" json:"category"`
	Province     string `orm:"province" json:"province"`
	City         string `orm:"city" json:"city"`
	District     string `orm:"district" json:"district"`
	Address      string `orm:"address" json:"address"`
	Longitude    string `orm:"longitude" json:"longitude"`
	Latitude     string `orm:"latitude" json:"latitude"`
	Telephone    string `orm:"telephone" json:"telephone"`
	PhotoList    string `orm:"photo_list" json:"photo_list"`
	AvgPrice     int    `orm:"avg_price" json:"avg_price"`
	OpenTime     string `orm:"open_time" json:"open_time"`
	Recommend    string `orm:"recommend" json:"recommend"`
	Special      string `orm:"special" json:"special"`
	Introduction string `orm:"introduction" json:"introduction"`
	OffsetType   int    `orm:"offset_type" json:"offset_type"`
	Status       int    `orm:"status" json:"status"`
	Message      string `orm:"message" json:"message"`
}

func (*Coupon_location) TableName() string {
	return "coupon_location"
}

type Ewei_shop_creditshop_option struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Goodsid       int     `orm:"goodsid" json:"goodsid"`
	Title         string  `orm:"title" json:"title"`
	Thumb         string  `orm:"thumb" json:"thumb"`
	Credit        int     `orm:"credit" json:"credit"`
	Money         float64 `orm:"money" json:"money"`
	Total         int     `orm:"total" json:"total"`
	Weight        float64 `orm:"weight" json:"weight"`
	Displayorder  int     `orm:"displayorder" json:"displayorder"`
	Specs         string  `orm:"specs" json:"specs"`
	SkuId         string  `orm:"skuId" json:"skuId"`
	Goodssn       string  `orm:"goodssn" json:"goodssn"`
	Productsn     string  `orm:"productsn" json:"productsn"`
	Virtual       int     `orm:"virtual" json:"virtual"`
	ExchangeStock int     `orm:"exchange_stock" json:"exchange_stock"`
}

func (*Ewei_shop_creditshop_option) TableName() string {
	return "ewei_shop_creditshop_option"
}

type Ewei_shop_exchange_code struct {
	Id            int    `orm:"id" json:"id"`
	Groupid       int    `orm:"groupid" json:"groupid"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Endtime       string `orm:"endtime" json:"endtime"`
	Status        int    `orm:"status" json:"status"`
	Openid        string `orm:"openid" json:"openid"`
	Count         int    `orm:"count" json:"count"`
	Key           string `orm:"key" json:"key"`
	Type          int    `orm:"type" json:"type"`
	Scene         int    `orm:"scene" json:"scene"`
	QrcodeUrl     string `orm:"qrcode_url" json:"qrcode_url"`
	Serial        string `orm:"serial" json:"serial"`
	Balancestatus int    `orm:"balancestatus" json:"balancestatus"`
	Redstatus     int    `orm:"redstatus" json:"redstatus"`
	Scorestatus   int    `orm:"scorestatus" json:"scorestatus"`
	Couponstatus  int    `orm:"couponstatus" json:"couponstatus"`
	Goodsstatus   int    `orm:"goodsstatus" json:"goodsstatus"`
	Repeatcount   int    `orm:"repeatcount" json:"repeatcount"`
}

func (*Ewei_shop_exchange_code) TableName() string {
	return "ewei_shop_exchange_code"
}

type Ewei_shop_groups_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_groups_adv) TableName() string {
	return "ewei_shop_groups_adv"
}

type Ewei_shop_order_goods struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Orderid         int     `orm:"orderid" json:"orderid"`
	Goodsid         int     `orm:"goodsid" json:"goodsid"`
	Price           float64 `orm:"price" json:"price"`
	Total           int     `orm:"total" json:"total"`
	Optionid        int     `orm:"optionid" json:"optionid"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Optionname      string  `orm:"optionname" json:"optionname"`
	Commission1     string  `orm:"commission1" json:"commission1"`
	Applytime1      int     `orm:"applytime1" json:"applytime1"`
	Checktime1      int     `orm:"checktime1" json:"checktime1"`
	Paytime1        int     `orm:"paytime1" json:"paytime1"`
	Invalidtime1    int     `orm:"invalidtime1" json:"invalidtime1"`
	Deletetime1     int     `orm:"deletetime1" json:"deletetime1"`
	Status1         int     `orm:"status1" json:"status1"`
	Content1        string  `orm:"content1" json:"content1"`
	Commission2     string  `orm:"commission2" json:"commission2"`
	Applytime2      int     `orm:"applytime2" json:"applytime2"`
	Checktime2      int     `orm:"checktime2" json:"checktime2"`
	Paytime2        int     `orm:"paytime2" json:"paytime2"`
	Invalidtime2    int     `orm:"invalidtime2" json:"invalidtime2"`
	Deletetime2     int     `orm:"deletetime2" json:"deletetime2"`
	Status2         int     `orm:"status2" json:"status2"`
	Content2        string  `orm:"content2" json:"content2"`
	Commission3     string  `orm:"commission3" json:"commission3"`
	Applytime3      int     `orm:"applytime3" json:"applytime3"`
	Checktime3      int     `orm:"checktime3" json:"checktime3"`
	Paytime3        int     `orm:"paytime3" json:"paytime3"`
	Invalidtime3    int     `orm:"invalidtime3" json:"invalidtime3"`
	Deletetime3     int     `orm:"deletetime3" json:"deletetime3"`
	Status3         int     `orm:"status3" json:"status3"`
	Content3        string  `orm:"content3" json:"content3"`
	Realprice       float64 `orm:"realprice" json:"realprice"`
	Goodssn         string  `orm:"goodssn" json:"goodssn"`
	Productsn       string  `orm:"productsn" json:"productsn"`
	Nocommission    int     `orm:"nocommission" json:"nocommission"`
	Changeprice     float64 `orm:"changeprice" json:"changeprice"`
	Oldprice        float64 `orm:"oldprice" json:"oldprice"`
	Commissions     string  `orm:"commissions" json:"commissions"`
	Diyformid       int     `orm:"diyformid" json:"diyformid"`
	Diyformdataid   int     `orm:"diyformdataid" json:"diyformdataid"`
	Diyformdata     string  `orm:"diyformdata" json:"diyformdata"`
	Diyformfields   string  `orm:"diyformfields" json:"diyformfields"`
	Openid          string  `orm:"openid" json:"openid"`
	Printstate      int     `orm:"printstate" json:"printstate"`
	Printstate2     int     `orm:"printstate2" json:"printstate2"`
	Rstate          int     `orm:"rstate" json:"rstate"`
	Refundtime      int     `orm:"refundtime" json:"refundtime"`
	Merchid         int     `orm:"merchid" json:"merchid"`
	Parentorderid   int     `orm:"parentorderid" json:"parentorderid"`
	Merchsale       int     `orm:"merchsale" json:"merchsale"`
	Isdiscountprice float64 `orm:"isdiscountprice" json:"isdiscountprice"`
	Canbuyagain     int     `orm:"canbuyagain" json:"canbuyagain"`
	Seckill         int     `orm:"seckill" json:"seckill"`
	SeckillTaskid   int     `orm:"seckill_taskid" json:"seckill_taskid"`
	SeckillRoomid   int     `orm:"seckill_roomid" json:"seckill_roomid"`
	SeckillTimeid   int     `orm:"seckill_timeid" json:"seckill_timeid"`
	IsMake          int     `orm:"is_make" json:"is_make"`
	Sendtype        int     `orm:"sendtype" json:"sendtype"`
	Expresscom      string  `orm:"expresscom" json:"expresscom"`
	Expresssn       string  `orm:"expresssn" json:"expresssn"`
	Express         string  `orm:"express" json:"express"`
	Sendtime        int     `orm:"sendtime" json:"sendtime"`
	Finishtime      int     `orm:"finishtime" json:"finishtime"`
	Remarksend      string  `orm:"remarksend" json:"remarksend"`
	Prohibitrefund  int     `orm:"prohibitrefund" json:"prohibitrefund"`
	Storeid         string  `orm:"storeid" json:"storeid"`
	TradeTime       int     `orm:"trade_time" json:"trade_time"`
	Optime          string  `orm:"optime" json:"optime"`
	TdateTime       int     `orm:"tdate_time" json:"tdate_time"`
	Dowpayment      float64 `orm:"dowpayment" json:"dowpayment"`
	Peopleid        int     `orm:"peopleid" json:"peopleid"`
	Esheetprintnum  int     `orm:"esheetprintnum" json:"esheetprintnum"`
	Ordercode       string  `orm:"ordercode" json:"ordercode"`
}

func (*Ewei_shop_order_goods) TableName() string {
	return "ewei_shop_order_goods"
}

type Ewei_shop_quick struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Title      string `orm:"title" json:"title"`
	Keyword    string `orm:"keyword" json:"keyword"`
	Datas      string `orm:"datas" json:"datas"`
	Cart       int    `orm:"cart" json:"cart"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Lasttime   int    `orm:"lasttime" json:"lasttime"`
	ShareTitle string `orm:"share_title" json:"share_title"`
	ShareDesc  string `orm:"share_desc" json:"share_desc"`
	ShareIcon  string `orm:"share_icon" json:"share_icon"`
	EnterTitle string `orm:"enter_title" json:"enter_title"`
	EnterDesc  string `orm:"enter_desc" json:"enter_desc"`
	EnterIcon  string `orm:"enter_icon" json:"enter_icon"`
	Status     int    `orm:"status" json:"status"`
	Merchid    int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_quick) TableName() string {
	return "ewei_shop_quick"
}

type Ewei_shop_sns_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_sns_adv) TableName() string {
	return "ewei_shop_sns_adv"
}

type Ewei_shop_sns_complain struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Type          int    `orm:"type" json:"type"`
	Postsid       int    `orm:"postsid" json:"postsid"`
	Defendant     string `orm:"defendant" json:"defendant"`
	Complainant   string `orm:"complainant" json:"complainant"`
	ComplaintType int    `orm:"complaint_type" json:"complaint_type"`
	ComplaintText string `orm:"complaint_text" json:"complaint_text"`
	Images        string `orm:"images" json:"images"`
	Createtime    int    `orm:"createtime" json:"createtime"`
	Checkedtime   int    `orm:"checkedtime" json:"checkedtime"`
	Checked       int    `orm:"checked" json:"checked"`
	CheckedNote   string `orm:"checked_note" json:"checked_note"`
	Deleted       int    `orm:"deleted" json:"deleted"`
}

func (*Ewei_shop_sns_complain) TableName() string {
	return "ewei_shop_sns_complain"
}

type Ewei_shop_system_site struct {
	Id      int    `orm:"id" json:"id"`
	Type    string `orm:"type" json:"type"`
	Content string `orm:"content" json:"content"`
}

func (*Ewei_shop_system_site) TableName() string {
	return "ewei_shop_system_site"
}

type Mc_fans_tag_mapping struct {
	Id    int `orm:"id" json:"id"`
	Fanid int `orm:"fanid" json:"fanid"`
	Tagid int `orm:"tagid" json:"tagid"`
}

func (*Mc_fans_tag_mapping) TableName() string {
	return "mc_fans_tag_mapping"
}

type Business struct {
	Id         int    `orm:"id" json:"id"`
	Weid       int    `orm:"weid" json:"weid"`
	Title      string `orm:"title" json:"title"`
	Thumb      string `orm:"thumb" json:"thumb"`
	Content    string `orm:"content" json:"content"`
	Phone      string `orm:"phone" json:"phone"`
	Qq         string `orm:"qq" json:"qq"`
	Province   string `orm:"province" json:"province"`
	City       string `orm:"city" json:"city"`
	Dist       string `orm:"dist" json:"dist"`
	Address    string `orm:"address" json:"address"`
	Lng        string `orm:"lng" json:"lng"`
	Lat        string `orm:"lat" json:"lat"`
	Industry1  string `orm:"industry1" json:"industry1"`
	Industry2  string `orm:"industry2" json:"industry2"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Business) TableName() string {
	return "business"
}

type Core_menu struct {
	Id             int    `orm:"id" json:"id"`
	Pid            int    `orm:"pid" json:"pid"`
	Title          string `orm:"title" json:"title"`
	Name           string `orm:"name" json:"name"`
	Url            string `orm:"url" json:"url"`
	AppendTitle    string `orm:"append_title" json:"append_title"`
	AppendUrl      string `orm:"append_url" json:"append_url"`
	Displayorder   int    `orm:"displayorder" json:"displayorder"`
	Type           string `orm:"type" json:"type"`
	IsDisplay      int    `orm:"is_display" json:"is_display"`
	IsSystem       int    `orm:"is_system" json:"is_system"`
	PermissionName string `orm:"permission_name" json:"permission_name"`
	GroupName      string `orm:"group_name" json:"group_name"`
	Icon           string `orm:"icon" json:"icon"`
}

func (*Core_menu) TableName() string {
	return "core_menu"
}

type Ewei_shop_member_favorite struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Goodsid    int    `orm:"goodsid" json:"goodsid"`
	Openid     string `orm:"openid" json:"openid"`
	Deleted    int    `orm:"deleted" json:"deleted"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Merchid    int    `orm:"merchid" json:"merchid"`
	Type       int    `orm:"type" json:"type"`
}

func (*Ewei_shop_member_favorite) TableName() string {
	return "ewei_shop_member_favorite"
}

type Ewei_shop_print struct {
	Id         int    `orm:"id" json:"id"`
	Status     int    `orm:"status" json:"status"`
	Name       string `orm:"name" json:"name"`
	PrintNo    string `orm:"print_no" json:"print_no"`
	Key        string `orm:"key" json:"key"`
	PrintNums  int    `orm:"print_nums" json:"print_nums"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Sid        int    `orm:"sid" json:"sid"`
	PrintType  int    `orm:"print_type" json:"print_type"`
	QrcodeLink string `orm:"qrcode_link" json:"qrcode_link"`
}

func (*Ewei_shop_print) TableName() string {
	return "ewei_shop_print"
}

type Ewei_shop_sns_post struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Bid         int    `orm:"bid" json:"bid"`
	Pid         int    `orm:"pid" json:"pid"`
	Rpid        int    `orm:"rpid" json:"rpid"`
	Openid      string `orm:"openid" json:"openid"`
	Avatar      string `orm:"avatar" json:"avatar"`
	Nickname    string `orm:"nickname" json:"nickname"`
	Title       string `orm:"title" json:"title"`
	Content     string `orm:"content" json:"content"`
	Images      string `orm:"images" json:"images"`
	Voice       string `orm:"voice" json:"voice"`
	Createtime  int    `orm:"createtime" json:"createtime"`
	Replytime   int    `orm:"replytime" json:"replytime"`
	Credit      int    `orm:"credit" json:"credit"`
	Views       int    `orm:"views" json:"views"`
	Islock      int    `orm:"islock" json:"islock"`
	Istop       int    `orm:"istop" json:"istop"`
	Isboardtop  int    `orm:"isboardtop" json:"isboardtop"`
	Isbest      int    `orm:"isbest" json:"isbest"`
	Isboardbest int    `orm:"isboardbest" json:"isboardbest"`
	Deleted     int    `orm:"deleted" json:"deleted"`
	Deletedtime int    `orm:"deletedtime" json:"deletedtime"`
	Checked     int    `orm:"checked" json:"checked"`
	Checktime   int    `orm:"checktime" json:"checktime"`
	Isadmin     int    `orm:"isadmin" json:"isadmin"`
}

func (*Ewei_shop_sns_post) TableName() string {
	return "ewei_shop_sns_post"
}

type Ewei_shop_system_plugingrant_package struct {
	Id           int    `orm:"id" json:"id"`
	Pluginid     string `orm:"pluginid" json:"pluginid"`
	Text         string `orm:"text" json:"text"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Data         string `orm:"data" json:"data"`
	State        int    `orm:"state" json:"state"`
	Rec          int    `orm:"rec" json:"rec"`
	Desc         string `orm:"desc" json:"desc"`
	Content      string `orm:"content" json:"content"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_system_plugingrant_package) TableName() string {
	return "ewei_shop_system_plugingrant_package"
}

type Ewei_shop_cashier_clearing struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Cashierid  int     `orm:"cashierid" json:"cashierid"`
	Clearno    string  `orm:"clearno" json:"clearno"`
	Status     int     `orm:"status" json:"status"`
	Money      float64 `orm:"money" json:"money"`
	Realmoney  float64 `orm:"realmoney" json:"realmoney"`
	Remark     string  `orm:"remark" json:"remark"`
	Orderids   string  `orm:"orderids" json:"orderids"`
	Createtime int     `orm:"createtime" json:"createtime"`
	Paytime    int     `orm:"paytime" json:"paytime"`
	Deleted    int     `orm:"deleted" json:"deleted"`
	Paytype    int     `orm:"paytype" json:"paytype"`
	Payinfo    string  `orm:"payinfo" json:"payinfo"`
	Charge     float64 `orm:"charge" json:"charge"`
}

func (*Ewei_shop_cashier_clearing) TableName() string {
	return "ewei_shop_cashier_clearing"
}

type Ewei_shop_coupon_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_coupon_category) TableName() string {
	return "ewei_shop_coupon_category"
}

type Ewei_shop_exchange_group struct {
	Id           int     `orm:"id" json:"id"`
	Title        string  `orm:"title" json:"title"`
	Type         int     `orm:"type" json:"type"`
	Endtime      string  `orm:"endtime" json:"endtime"`
	Mode         int     `orm:"mode" json:"mode"`
	Status       int     `orm:"status" json:"status"`
	Max          int     `orm:"max" json:"max"`
	Value        float64 `orm:"value" json:"value"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Starttime    string  `orm:"starttime" json:"starttime"`
	Goods        string  `orm:"goods" json:"goods"`
	Score        int     `orm:"score" json:"score"`
	Coupon       string  `orm:"coupon" json:"coupon"`
	Use          int     `orm:"use" json:"use"`
	Total        int     `orm:"total" json:"total"`
	Red          float64 `orm:"red" json:"red"`
	Balance      float64 `orm:"balance" json:"balance"`
	BalanceLeft  float64 `orm:"balance_left" json:"balance_left"`
	BalanceRight float64 `orm:"balance_right" json:"balance_right"`
	RedLeft      float64 `orm:"red_left" json:"red_left"`
	RedRight     float64 `orm:"red_right" json:"red_right"`
	ScoreLeft    int     `orm:"score_left" json:"score_left"`
	ScoreRight   int     `orm:"score_right" json:"score_right"`
	BalanceType  int     `orm:"balance_type" json:"balance_type"`
	RedType      int     `orm:"red_type" json:"red_type"`
	ScoreType    int     `orm:"score_type" json:"score_type"`
	TitleReply   string  `orm:"title_reply" json:"title_reply"`
	Img          string  `orm:"img" json:"img"`
	Content      string  `orm:"content" json:"content"`
	Rule         string  `orm:"rule" json:"rule"`
	CouponType   string  `orm:"coupon_type" json:"coupon_type"`
	BasicContent string  `orm:"basic_content" json:"basic_content"`
	ReplyType    int     `orm:"reply_type" json:"reply_type"`
	CodeType     int     `orm:"code_type" json:"code_type"`
	Binding      int     `orm:"binding" json:"binding"`
	Showcount    int     `orm:"showcount" json:"showcount"`
	Postage      float64 `orm:"postage" json:"postage"`
	PostageType  int     `orm:"postage_type" json:"postage_type"`
	Banner       string  `orm:"banner" json:"banner"`
	KeywordReply int     `orm:"keyword_reply" json:"keyword_reply"`
	ReplyStatus  int     `orm:"reply_status" json:"reply_status"`
	ReplyKeyword string  `orm:"reply_keyword" json:"reply_keyword"`
	InputBanner  string  `orm:"input_banner" json:"input_banner"`
	Diypage      int     `orm:"diypage" json:"diypage"`
	Sendname     string  `orm:"sendname" json:"sendname"`
	Wishing      string  `orm:"wishing" json:"wishing"`
	Actname      string  `orm:"actname" json:"actname"`
	Remark       string  `orm:"remark" json:"remark"`
	Repeat       int     `orm:"repeat" json:"repeat"`
	Koulingstart string  `orm:"koulingstart" json:"koulingstart"`
	Koulingend   string  `orm:"koulingend" json:"koulingend"`
	Kouling      int     `orm:"kouling" json:"kouling"`
	Chufa        string  `orm:"chufa" json:"chufa"`
	Chufaend     string  `orm:"chufaend" json:"chufaend"`
}

func (*Ewei_shop_exchange_group) TableName() string {
	return "ewei_shop_exchange_group"
}

type Ewei_shop_system_copyright struct {
	Id        int    `orm:"id" json:"id"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Copyright string `orm:"copyright" json:"copyright"`
	Bgcolor   string `orm:"bgcolor" json:"bgcolor"`
	Ismanage  int    `orm:"ismanage" json:"ismanage"`
	Logo      string `orm:"logo" json:"logo"`
	Title     string `orm:"title" json:"title"`
}

func (*Ewei_shop_system_copyright) TableName() string {
	return "ewei_shop_system_copyright"
}

type Site_nav struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Multiid      int    `orm:"multiid" json:"multiid"`
	Section      int    `orm:"section" json:"section"`
	Module       string `orm:"module" json:"module"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Name         string `orm:"name" json:"name"`
	Description  string `orm:"description" json:"description"`
	Position     int    `orm:"position" json:"position"`
	Url          string `orm:"url" json:"url"`
	Icon         string `orm:"icon" json:"icon"`
	Css          string `orm:"css" json:"css"`
	Status       int    `orm:"status" json:"status"`
	Categoryid   int    `orm:"categoryid" json:"categoryid"`
}

func (*Site_nav) TableName() string {
	return "site_nav"
}

type Ewei_shop_seckill_task_goods struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Displayorder int     `orm:"displayorder" json:"displayorder"`
	Taskid       int     `orm:"taskid" json:"taskid"`
	Roomid       int     `orm:"roomid" json:"roomid"`
	Timeid       int     `orm:"timeid" json:"timeid"`
	Goodsid      int     `orm:"goodsid" json:"goodsid"`
	Optionid     int     `orm:"optionid" json:"optionid"`
	Price        float64 `orm:"price" json:"price"`
	Total        int     `orm:"total" json:"total"`
	Maxbuy       int     `orm:"maxbuy" json:"maxbuy"`
	Totalmaxbuy  int     `orm:"totalmaxbuy" json:"totalmaxbuy"`
	Commission1  float64 `orm:"commission1" json:"commission1"`
	Commission2  float64 `orm:"commission2" json:"commission2"`
	Commission3  float64 `orm:"commission3" json:"commission3"`
}

func (*Ewei_shop_seckill_task_goods) TableName() string {
	return "ewei_shop_seckill_task_goods"
}

type Ewei_shop_task_set struct {
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Entrance    int    `orm:"entrance" json:"entrance"`
	Keyword     string `orm:"keyword" json:"keyword"`
	CoverTitle  string `orm:"cover_title" json:"cover_title"`
	CoverImg    string `orm:"cover_img" json:"cover_img"`
	CoverDesc   string `orm:"cover_desc" json:"cover_desc"`
	MsgPick     string `orm:"msg_pick" json:"msg_pick"`
	MsgProgress string `orm:"msg_progress" json:"msg_progress"`
	MsgFinish   string `orm:"msg_finish" json:"msg_finish"`
	MsgFollow   string `orm:"msg_follow" json:"msg_follow"`
	Isnew       int    `orm:"isnew" json:"isnew"`
	BgImg       string `orm:"bg_img" json:"bg_img"`
	TopNotice   int    `orm:"top_notice" json:"top_notice"`
}

func (*Ewei_shop_task_set) TableName() string {
	return "ewei_shop_task_set"
}

type Mobilenumber struct {
	Id       int `orm:"id" json:"id"`
	Rid      int `orm:"rid" json:"rid"`
	Enabled  int `orm:"enabled" json:"enabled"`
	Dateline int `orm:"dateline" json:"dateline"`
}

func (*Mobilenumber) TableName() string {
	return "mobilenumber"
}

type Basic_reply struct {
	Id      int    `orm:"id" json:"id"`
	Rid     int    `orm:"rid" json:"rid"`
	Content string `orm:"content" json:"content"`
}

func (*Basic_reply) TableName() string {
	return "basic_reply"
}

type Ewei_shop_customer struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	KfId         string `orm:"kf_id" json:"kf_id"`
	KfAccount    string `orm:"kf_account" json:"kf_account"`
	KfNick       string `orm:"kf_nick" json:"kf_nick"`
	KfPwd        string `orm:"kf_pwd" json:"kf_pwd"`
	KfHeadimgurl string `orm:"kf_headimgurl" json:"kf_headimgurl"`
	Createtime   int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_customer) TableName() string {
	return "ewei_shop_customer"
}

type Ewei_shop_globonus_billp struct {
	Id          int     `orm:"id" json:"id"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Billid      int     `orm:"billid" json:"billid"`
	Openid      string  `orm:"openid" json:"openid"`
	Payno       string  `orm:"payno" json:"payno"`
	Paytype     int     `orm:"paytype" json:"paytype"`
	Bonus       float64 `orm:"bonus" json:"bonus"`
	Money       float64 `orm:"money" json:"money"`
	Realmoney   float64 `orm:"realmoney" json:"realmoney"`
	Paymoney    float64 `orm:"paymoney" json:"paymoney"`
	Charge      float64 `orm:"charge" json:"charge"`
	Chargemoney float64 `orm:"chargemoney" json:"chargemoney"`
	Status      int     `orm:"status" json:"status"`
	Reason      string  `orm:"reason" json:"reason"`
	Paytime     int     `orm:"paytime" json:"paytime"`
}

func (*Ewei_shop_globonus_billp) TableName() string {
	return "ewei_shop_globonus_billp"
}

type Ewei_shop_goods_param struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Goodsid      int    `orm:"goodsid" json:"goodsid"`
	Title        string `orm:"title" json:"title"`
	Value        string `orm:"value" json:"value"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_goods_param) TableName() string {
	return "ewei_shop_goods_param"
}

type Ewei_shop_poster struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Type          int     `orm:"type" json:"type"`
	Title         string  `orm:"title" json:"title"`
	Bg            string  `orm:"bg" json:"bg"`
	Data          string  `orm:"data" json:"data"`
	Keyword       string  `orm:"keyword" json:"keyword"`
	Times         int     `orm:"times" json:"times"`
	Follows       int     `orm:"follows" json:"follows"`
	Isdefault     int     `orm:"isdefault" json:"isdefault"`
	Resptitle     string  `orm:"resptitle" json:"resptitle"`
	Respthumb     string  `orm:"respthumb" json:"respthumb"`
	Createtime    int     `orm:"createtime" json:"createtime"`
	Respdesc      string  `orm:"respdesc" json:"respdesc"`
	Respurl       string  `orm:"respurl" json:"respurl"`
	Waittext      string  `orm:"waittext" json:"waittext"`
	Oktext        string  `orm:"oktext" json:"oktext"`
	Subcredit     int     `orm:"subcredit" json:"subcredit"`
	Submoney      float64 `orm:"submoney" json:"submoney"`
	Reccredit     int     `orm:"reccredit" json:"reccredit"`
	Recmoney      float64 `orm:"recmoney" json:"recmoney"`
	Paytype       int     `orm:"paytype" json:"paytype"`
	Scantext      string  `orm:"scantext" json:"scantext"`
	Subtext       string  `orm:"subtext" json:"subtext"`
	Beagent       int     `orm:"beagent" json:"beagent"`
	Bedown        int     `orm:"bedown" json:"bedown"`
	Isopen        int     `orm:"isopen" json:"isopen"`
	Opentext      string  `orm:"opentext" json:"opentext"`
	Openurl       string  `orm:"openurl" json:"openurl"`
	Templateid    string  `orm:"templateid" json:"templateid"`
	Subpaycontent string  `orm:"subpaycontent" json:"subpaycontent"`
	Recpaycontent string  `orm:"recpaycontent" json:"recpaycontent"`
	Entrytext     string  `orm:"entrytext" json:"entrytext"`
	Reccouponid   int     `orm:"reccouponid" json:"reccouponid"`
	Reccouponnum  int     `orm:"reccouponnum" json:"reccouponnum"`
	Subcouponid   int     `orm:"subcouponid" json:"subcouponid"`
	Subcouponnum  int     `orm:"subcouponnum" json:"subcouponnum"`
	Resptype      int     `orm:"resptype" json:"resptype"`
	Resptext      string  `orm:"resptext" json:"resptext"`
	Keyword2      string  `orm:"keyword2" json:"keyword2"`
	Resptext11    string  `orm:"resptext11" json:"resptext11"`
	RewardTotle   string  `orm:"reward_totle" json:"reward_totle"`
	Ismembergroup int     `orm:"ismembergroup" json:"ismembergroup"`
	Membergroupid int     `orm:"membergroupid" json:"membergroupid"`
}

func (*Ewei_shop_poster) TableName() string {
	return "ewei_shop_poster"
}

type Ewei_shop_task_default struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Data    string `orm:"data" json:"data"`
	Addtime int    `orm:"addtime" json:"addtime"`
	Bgimg   string `orm:"bgimg" json:"bgimg"`
	Open    int    `orm:"open" json:"open"`
}

func (*Ewei_shop_task_default) TableName() string {
	return "ewei_shop_task_default"
}

type Uni_settings struct {
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Passport        string `orm:"passport" json:"passport"`
	Oauth           string `orm:"oauth" json:"oauth"`
	JsauthAcid      int    `orm:"jsauth_acid" json:"jsauth_acid"`
	Uc              string `orm:"uc" json:"uc"`
	Notify          string `orm:"notify" json:"notify"`
	Creditnames     string `orm:"creditnames" json:"creditnames"`
	Creditbehaviors string `orm:"creditbehaviors" json:"creditbehaviors"`
	Welcome         string `orm:"welcome" json:"welcome"`
	Default         string `orm:"default" json:"default"`
	DefaultMessage  string `orm:"default_message" json:"default_message"`
	Payment         string `orm:"payment" json:"payment"`
	Stat            string `orm:"stat" json:"stat"`
	DefaultSite     int    `orm:"default_site" json:"default_site"`
	Sync            int    `orm:"sync" json:"sync"`
	Recharge        string `orm:"recharge" json:"recharge"`
	Tplnotice       string `orm:"tplnotice" json:"tplnotice"`
	Grouplevel      int    `orm:"grouplevel" json:"grouplevel"`
	Mcplugin        string `orm:"mcplugin" json:"mcplugin"`
	ExchangeEnable  int    `orm:"exchange_enable" json:"exchange_enable"`
	CouponType      int    `orm:"coupon_type" json:"coupon_type"`
	Menuset         string `orm:"menuset" json:"menuset"`
	Shortcuts       string `orm:"shortcuts" json:"shortcuts"`
	Statistics      string `orm:"statistics" json:"statistics"`
	BindDomain      string `orm:"bind_domain" json:"bind_domain"`
}

func (*Uni_settings) TableName() string {
	return "uni_settings"
}

type Core_cache struct {
	Key   string `orm:"key" json:"key"`
	Value string `orm:"value" json:"value"`
}

func (*Core_cache) TableName() string {
	return "core_cache"
}

type Core_sessions struct {
	Sid        string `orm:"sid" json:"sid"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Data       string `orm:"data" json:"data"`
	Expiretime int    `orm:"expiretime" json:"expiretime"`
}

func (*Core_sessions) TableName() string {
	return "core_sessions"
}

type Coupon_groups struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Couponid string `orm:"couponid" json:"couponid"`
	Groupid  int    `orm:"groupid" json:"groupid"`
}

func (*Coupon_groups) TableName() string {
	return "coupon_groups"
}

type Ewei_shop_diypage_template struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Type    int    `orm:"type" json:"type"`
	Name    string `orm:"name" json:"name"`
	Data    string `orm:"data" json:"data"`
	Preview string `orm:"preview" json:"preview"`
	Tplid   int    `orm:"tplid" json:"tplid"`
	Cate    int    `orm:"cate" json:"cate"`
	Deleted int    `orm:"deleted" json:"deleted"`
	Merch   int    `orm:"merch" json:"merch"`
}

func (*Ewei_shop_diypage_template) TableName() string {
	return "ewei_shop_diypage_template"
}

type Ewei_shop_goods_spec struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Goodsid      int    `orm:"goodsid" json:"goodsid"`
	Title        string `orm:"title" json:"title"`
	Description  string `orm:"description" json:"description"`
	Displaytype  int    `orm:"displaytype" json:"displaytype"`
	Content      string `orm:"content" json:"content"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	PropId       string `orm:"propId" json:"propId"`
}

func (*Ewei_shop_goods_spec) TableName() string {
	return "ewei_shop_goods_spec"
}

type Ewei_shop_invitation_log struct {
	Id               int    `orm:"id" json:"id"`
	Uniacid          int    `orm:"uniacid" json:"uniacid"`
	InvitationId     int    `orm:"invitation_id" json:"invitation_id"`
	Openid           string `orm:"openid" json:"openid"`
	InvitationOpenid string `orm:"invitation_openid" json:"invitation_openid"`
	ScanTime         int    `orm:"scan_time" json:"scan_time"`
	Follow           int    `orm:"follow" json:"follow"`
}

func (*Ewei_shop_invitation_log) TableName() string {
	return "ewei_shop_invitation_log"
}

type Ewei_shop_order_comment struct {
	Id                 int    `orm:"id" json:"id"`
	Uniacid            int    `orm:"uniacid" json:"uniacid"`
	Orderid            int    `orm:"orderid" json:"orderid"`
	Goodsid            int    `orm:"goodsid" json:"goodsid"`
	Openid             string `orm:"openid" json:"openid"`
	Nickname           string `orm:"nickname" json:"nickname"`
	Headimgurl         string `orm:"headimgurl" json:"headimgurl"`
	Level              int    `orm:"level" json:"level"`
	Content            string `orm:"content" json:"content"`
	Images             string `orm:"images" json:"images"`
	Createtime         int    `orm:"createtime" json:"createtime"`
	Deleted            int    `orm:"deleted" json:"deleted"`
	AppendContent      string `orm:"append_content" json:"append_content"`
	AppendImages       string `orm:"append_images" json:"append_images"`
	ReplyContent       string `orm:"reply_content" json:"reply_content"`
	ReplyImages        string `orm:"reply_images" json:"reply_images"`
	AppendReplyContent string `orm:"append_reply_content" json:"append_reply_content"`
	AppendReplyImages  string `orm:"append_reply_images" json:"append_reply_images"`
	Istop              int    `orm:"istop" json:"istop"`
	Checked            int    `orm:"checked" json:"checked"`
	Replychecked       int    `orm:"replychecked" json:"replychecked"`
}

func (*Ewei_shop_order_comment) TableName() string {
	return "ewei_shop_order_comment"
}

type Ewei_shop_store struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Storename     string `orm:"storename" json:"storename"`
	Address       string `orm:"address" json:"address"`
	Tel           string `orm:"tel" json:"tel"`
	Lat           string `orm:"lat" json:"lat"`
	Lng           string `orm:"lng" json:"lng"`
	Status        int    `orm:"status" json:"status"`
	Realname      string `orm:"realname" json:"realname"`
	Mobile        string `orm:"mobile" json:"mobile"`
	Fetchtime     string `orm:"fetchtime" json:"fetchtime"`
	Type          int    `orm:"type" json:"type"`
	Logo          string `orm:"logo" json:"logo"`
	Saletime      string `orm:"saletime" json:"saletime"`
	Desc          string `orm:"desc" json:"desc"`
	Displayorder  int    `orm:"displayorder" json:"displayorder"`
	OrderPrinter  string `orm:"order_printer" json:"order_printer"`
	OrderTemplate int    `orm:"order_template" json:"order_template"`
	Ordertype     string `orm:"ordertype" json:"ordertype"`
	Banner        string `orm:"banner" json:"banner"`
	Label         string `orm:"label" json:"label"`
	Tag           string `orm:"tag" json:"tag"`
	Classify      int    `orm:"classify" json:"classify"`
	Perms         string `orm:"perms" json:"perms"`
	Citycode      string `orm:"citycode" json:"citycode"`
	Opensend      int    `orm:"opensend" json:"opensend"`
	Province      string `orm:"province" json:"province"`
	City          string `orm:"city" json:"city"`
	Area          string `orm:"area" json:"area"`
	Provincecode  string `orm:"provincecode" json:"provincecode"`
	Areacode      string `orm:"areacode" json:"areacode"`
	Diypage       int    `orm:"diypage" json:"diypage"`
	DiypageIspage int    `orm:"diypage_ispage" json:"diypage_ispage"`
	DiypageList   string `orm:"diypage_list" json:"diypage_list"`
	Storegroupid  int    `orm:"storegroupid" json:"storegroupid"`
	Cates         string `orm:"cates" json:"cates"`
}

func (*Ewei_shop_store) TableName() string {
	return "ewei_shop_store"
}

type Ewei_shop_wxapp_page struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Type       int    `orm:"type" json:"type"`
	Name       string `orm:"name" json:"name"`
	Data       string `orm:"data" json:"data"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Lasttime   int    `orm:"lasttime" json:"lasttime"`
	Status     int    `orm:"status" json:"status"`
	Isdefault  int    `orm:"isdefault" json:"isdefault"`
}

func (*Ewei_shop_wxapp_page) TableName() string {
	return "ewei_shop_wxapp_page"
}

type Mc_cash_record struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Uid        int     `orm:"uid" json:"uid"`
	ClerkId    int     `orm:"clerk_id" json:"clerk_id"`
	StoreId    int     `orm:"store_id" json:"store_id"`
	ClerkType  int     `orm:"clerk_type" json:"clerk_type"`
	Fee        float64 `orm:"fee" json:"fee"`
	FinalFee   float64 `orm:"final_fee" json:"final_fee"`
	Credit1    int     `orm:"credit1" json:"credit1"`
	Credit1Fee float64 `orm:"credit1_fee" json:"credit1_fee"`
	Credit2    float64 `orm:"credit2" json:"credit2"`
	Cash       float64 `orm:"cash" json:"cash"`
	ReturnCash float64 `orm:"return_cash" json:"return_cash"`
	FinalCash  float64 `orm:"final_cash" json:"final_cash"`
	Remark     string  `orm:"remark" json:"remark"`
	Createtime int     `orm:"createtime" json:"createtime"`
	TradeType  string  `orm:"trade_type" json:"trade_type"`
}

func (*Mc_cash_record) TableName() string {
	return "mc_cash_record"
}

type Mc_credits_recharge struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Uid        int    `orm:"uid" json:"uid"`
	Openid     string `orm:"openid" json:"openid"`
	Tid        string `orm:"tid" json:"tid"`
	Transid    string `orm:"transid" json:"transid"`
	Fee        string `orm:"fee" json:"fee"`
	Type       string `orm:"type" json:"type"`
	Tag        string `orm:"tag" json:"tag"`
	Status     int    `orm:"status" json:"status"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Backtype   int    `orm:"backtype" json:"backtype"`
}

func (*Mc_credits_recharge) TableName() string {
	return "mc_credits_recharge"
}

type Modules_plugin struct {
	Id         int    `orm:"id" json:"id"`
	Name       string `orm:"name" json:"name"`
	MainModule string `orm:"main_module" json:"main_module"`
}

func (*Modules_plugin) TableName() string {
	return "modules_plugin"
}

type Uni_verifycode struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Receiver   string `orm:"receiver" json:"receiver"`
	Verifycode string `orm:"verifycode" json:"verifycode"`
	Total      int    `orm:"total" json:"total"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Uni_verifycode) TableName() string {
	return "uni_verifycode"
}

type Account_wechats struct {
	Acid             int    `orm:"acid" json:"acid"`
	Uniacid          int    `orm:"uniacid" json:"uniacid"`
	Token            string `orm:"token" json:"token"`
	Encodingaeskey   string `orm:"encodingaeskey" json:"encodingaeskey"`
	Level            int    `orm:"level" json:"level"`
	Name             string `orm:"name" json:"name"`
	Account          string `orm:"account" json:"account"`
	Original         string `orm:"original" json:"original"`
	Signature        string `orm:"signature" json:"signature"`
	Country          string `orm:"country" json:"country"`
	Province         string `orm:"province" json:"province"`
	City             string `orm:"city" json:"city"`
	Username         string `orm:"username" json:"username"`
	Password         string `orm:"password" json:"password"`
	Lastupdate       int    `orm:"lastupdate" json:"lastupdate"`
	Key              string `orm:"key" json:"key"`
	Secret           string `orm:"secret" json:"secret"`
	Styleid          int    `orm:"styleid" json:"styleid"`
	Subscribeurl     string `orm:"subscribeurl" json:"subscribeurl"`
	AuthRefreshToken string `orm:"auth_refresh_token" json:"auth_refresh_token"`
	AccessToken      string `orm:"access_token" json:"access_token"`
}

func (*Account_wechats) TableName() string {
	return "account_wechats"
}

type Ewei_shop_article_sys struct {
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	ArticleMessage string `orm:"article_message" json:"article_message"`
	ArticleTitle   string `orm:"article_title" json:"article_title"`
	ArticleImage   string `orm:"article_image" json:"article_image"`
	ArticleShownum int    `orm:"article_shownum" json:"article_shownum"`
	ArticleKeyword string `orm:"article_keyword" json:"article_keyword"`
	ArticleTemp    int    `orm:"article_temp" json:"article_temp"`
	ArticleSource  string `orm:"article_source" json:"article_source"`
}

func (*Ewei_shop_article_sys) TableName() string {
	return "ewei_shop_article_sys"
}

type Ewei_shop_author_level struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Levelname       string  `orm:"levelname" json:"levelname"`
	Bonus           float64 `orm:"bonus" json:"bonus"`
	Ordermoney      float64 `orm:"ordermoney" json:"ordermoney"`
	Ordercount      int     `orm:"ordercount" json:"ordercount"`
	Commissionmoney float64 `orm:"commissionmoney" json:"commissionmoney"`
	Bonusmoney      float64 `orm:"bonusmoney" json:"bonusmoney"`
	Downcount       int     `orm:"downcount" json:"downcount"`
	BonusFg         string  `orm:"bonus_fg" json:"bonus_fg"`
}

func (*Ewei_shop_author_level) TableName() string {
	return "ewei_shop_author_level"
}

type Ewei_shop_seckill_task struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Cateid     int    `orm:"cateid" json:"cateid"`
	Title      string `orm:"title" json:"title"`
	Enabled    int    `orm:"enabled" json:"enabled"`
	PageTitle  string `orm:"page_title" json:"page_title"`
	ShareTitle string `orm:"share_title" json:"share_title"`
	ShareDesc  string `orm:"share_desc" json:"share_desc"`
	ShareIcon  string `orm:"share_icon" json:"share_icon"`
	Tag        string `orm:"tag" json:"tag"`
	Closesec   int    `orm:"closesec" json:"closesec"`
	Oldshow    int    `orm:"oldshow" json:"oldshow"`
	Times      string `orm:"times" json:"times"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Overtimes  int    `orm:"overtimes" json:"overtimes"`
}

func (*Ewei_shop_seckill_task) TableName() string {
	return "ewei_shop_seckill_task"
}

type Article_notice struct {
	Id           int    `orm:"id" json:"id"`
	Cateid       int    `orm:"cateid" json:"cateid"`
	Title        string `orm:"title" json:"title"`
	Content      string `orm:"content" json:"content"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	IsDisplay    int    `orm:"is_display" json:"is_display"`
	IsShowHome   int    `orm:"is_show_home" json:"is_show_home"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Click        int    `orm:"click" json:"click"`
	Style        string `orm:"style" json:"style"`
}

func (*Article_notice) TableName() string {
	return "article_notice"
}

type Ewei_shop_cashier_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Catename     string `orm:"catename" json:"catename"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_cashier_category) TableName() string {
	return "ewei_shop_cashier_category"
}

type Ewei_shop_creditshop_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_creditshop_adv) TableName() string {
	return "ewei_shop_creditshop_adv"
}

type Ewei_shop_live_status struct {
	Id        int `orm:"id" json:"id"`
	Uniacid   int `orm:"uniacid" json:"uniacid"`
	Roomid    int `orm:"roomid" json:"roomid"`
	Starttime int `orm:"starttime" json:"starttime"`
	Endtime   int `orm:"endtime" json:"endtime"`
}

func (*Ewei_shop_live_status) TableName() string {
	return "ewei_shop_live_status"
}

type Ewei_shop_pc_menu struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Type         int    `orm:"type" json:"type"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Title        string `orm:"title" json:"title"`
	Link         string `orm:"link" json:"link"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Createtime   int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_pc_menu) TableName() string {
	return "ewei_shop_pc_menu"
}

type Ewei_shop_sign_records struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Time    int    `orm:"time" json:"time"`
	Openid  string `orm:"openid" json:"openid"`
	Credit  int    `orm:"credit" json:"credit"`
	Log     string `orm:"log" json:"log"`
	Type    int    `orm:"type" json:"type"`
	Day     int    `orm:"day" json:"day"`
}

func (*Ewei_shop_sign_records) TableName() string {
	return "ewei_shop_sign_records"
}

type Qrcode struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Acid       int    `orm:"acid" json:"acid"`
	Type       string `orm:"type" json:"type"`
	Extra      int    `orm:"extra" json:"extra"`
	Qrcid      int    `orm:"qrcid" json:"qrcid"`
	SceneStr   string `orm:"scene_str" json:"scene_str"`
	Name       string `orm:"name" json:"name"`
	Keyword    string `orm:"keyword" json:"keyword"`
	Model      int    `orm:"model" json:"model"`
	Ticket     string `orm:"ticket" json:"ticket"`
	Url        string `orm:"url" json:"url"`
	Expire     int    `orm:"expire" json:"expire"`
	Subnum     int    `orm:"subnum" json:"subnum"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Status     int    `orm:"status" json:"status"`
}

func (*Qrcode) TableName() string {
	return "qrcode"
}

type Stat_visit struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Module  string `orm:"module" json:"module"`
	Count   int    `orm:"count" json:"count"`
	Date    int    `orm:"date" json:"date"`
}

func (*Stat_visit) TableName() string {
	return "stat_visit"
}

type Wechat_news struct {
	Id               int    `orm:"id" json:"id"`
	Uniacid          int    `orm:"uniacid" json:"uniacid"`
	AttachId         int    `orm:"attach_id" json:"attach_id"`
	ThumbMediaId     string `orm:"thumb_media_id" json:"thumb_media_id"`
	ThumbUrl         string `orm:"thumb_url" json:"thumb_url"`
	Title            string `orm:"title" json:"title"`
	Author           string `orm:"author" json:"author"`
	Digest           string `orm:"digest" json:"digest"`
	Content          string `orm:"content" json:"content"`
	ContentSourceUrl string `orm:"content_source_url" json:"content_source_url"`
	ShowCoverPic     int    `orm:"show_cover_pic" json:"show_cover_pic"`
	Url              string `orm:"url" json:"url"`
	Displayorder     int    `orm:"displayorder" json:"displayorder"`
}

func (*Wechat_news) TableName() string {
	return "wechat_news"
}

type Ewei_shop_cashier_qrcode struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Cashierid  int     `orm:"cashierid" json:"cashierid"`
	Title      string  `orm:"title" json:"title"`
	Goodstitle string  `orm:"goodstitle" json:"goodstitle"`
	Money      float64 `orm:"money" json:"money"`
	Createtime int     `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_cashier_qrcode) TableName() string {
	return "ewei_shop_cashier_qrcode"
}

type Ewei_shop_diyform_temp struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Typeid          int    `orm:"typeid" json:"typeid"`
	Cid             int    `orm:"cid" json:"cid"`
	Diyformfields   string `orm:"diyformfields" json:"diyformfields"`
	Fields          string `orm:"fields" json:"fields"`
	Openid          string `orm:"openid" json:"openid"`
	Type            int    `orm:"type" json:"type"`
	Diyformid       int    `orm:"diyformid" json:"diyformid"`
	Diyformdata     string `orm:"diyformdata" json:"diyformdata"`
	CarrierRealname string `orm:"carrier_realname" json:"carrier_realname"`
	CarrierMobile   string `orm:"carrier_mobile" json:"carrier_mobile"`
}

func (*Ewei_shop_diyform_temp) TableName() string {
	return "ewei_shop_diyform_temp"
}

type Ewei_shop_merch_bill struct {
	Id                int     `orm:"id" json:"id"`
	Uniacid           int     `orm:"uniacid" json:"uniacid"`
	Applyno           string  `orm:"applyno" json:"applyno"`
	Merchid           int     `orm:"merchid" json:"merchid"`
	Orderids          string  `orm:"orderids" json:"orderids"`
	Realprice         float64 `orm:"realprice" json:"realprice"`
	Realpricerate     float64 `orm:"realpricerate" json:"realpricerate"`
	Finalprice        float64 `orm:"finalprice" json:"finalprice"`
	Payrateprice      float64 `orm:"payrateprice" json:"payrateprice"`
	Payrate           float64 `orm:"payrate" json:"payrate"`
	Money             float64 `orm:"money" json:"money"`
	Applytime         int     `orm:"applytime" json:"applytime"`
	Checktime         int     `orm:"checktime" json:"checktime"`
	Paytime           int     `orm:"paytime" json:"paytime"`
	Invalidtime       int     `orm:"invalidtime" json:"invalidtime"`
	Refusetime        int     `orm:"refusetime" json:"refusetime"`
	Remark            string  `orm:"remark" json:"remark"`
	Status            int     `orm:"status" json:"status"`
	Ordernum          int     `orm:"ordernum" json:"ordernum"`
	Orderprice        float64 `orm:"orderprice" json:"orderprice"`
	Price             float64 `orm:"price" json:"price"`
	Passrealprice     float64 `orm:"passrealprice" json:"passrealprice"`
	Passrealpricerate float64 `orm:"passrealpricerate" json:"passrealpricerate"`
	Passorderids      string  `orm:"passorderids" json:"passorderids"`
	Passordernum      int     `orm:"passordernum" json:"passordernum"`
	Passorderprice    float64 `orm:"passorderprice" json:"passorderprice"`
	Alipay            string  `orm:"alipay" json:"alipay"`
	Bankname          string  `orm:"bankname" json:"bankname"`
	Bankcard          string  `orm:"bankcard" json:"bankcard"`
	Applyrealname     string  `orm:"applyrealname" json:"applyrealname"`
	Applytype         int     `orm:"applytype" json:"applytype"`
	Handpay           int     `orm:"handpay" json:"handpay"`
}

func (*Ewei_shop_merch_bill) TableName() string {
	return "ewei_shop_merch_bill"
}

type Ewei_shop_saler struct {
	Id         int    `orm:"id" json:"id"`
	Storeid    int    `orm:"storeid" json:"storeid"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Status     int    `orm:"status" json:"status"`
	Salername  string `orm:"salername" json:"salername"`
	Username   string `orm:"username" json:"username"`
	Pwd        string `orm:"pwd" json:"pwd"`
	Salt       string `orm:"salt" json:"salt"`
	Lastvisit  string `orm:"lastvisit" json:"lastvisit"`
	Lastip     string `orm:"lastip" json:"lastip"`
	Isfounder  int    `orm:"isfounder" json:"isfounder"`
	Mobile     string `orm:"mobile" json:"mobile"`
	Getmessage int    `orm:"getmessage" json:"getmessage"`
	Getnotice  int    `orm:"getnotice" json:"getnotice"`
	Roleid     int    `orm:"roleid" json:"roleid"`
}

func (*Ewei_shop_saler) TableName() string {
	return "ewei_shop_saler"
}

type Ewei_shop_task_joiner struct {
	CompleteId int    `orm:"complete_id" json:"complete_id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	TaskUser   string `orm:"task_user" json:"task_user"`
	JoinerId   string `orm:"joiner_id" json:"joiner_id"`
	JoinId     int    `orm:"join_id" json:"join_id"`
	TaskId     int    `orm:"task_id" json:"task_id"`
	TaskType   int    `orm:"task_type" json:"task_type"`
	JoinStatus int    `orm:"join_status" json:"join_status"`
	Addtime    int    `orm:"addtime" json:"addtime"`
}

func (*Ewei_shop_task_joiner) TableName() string {
	return "ewei_shop_task_joiner"
}

type Ewei_shop_wxcard struct {
	Id                      int     `orm:"id" json:"id"`
	Uniacid                 int     `orm:"uniacid" json:"uniacid"`
	CardId                  string  `orm:"card_id" json:"card_id"`
	Displayorder            int     `orm:"displayorder" json:"displayorder"`
	Catid                   int     `orm:"catid" json:"catid"`
	CardType                string  `orm:"card_type" json:"card_type"`
	LogoUrl                 string  `orm:"logo_url" json:"logo_url"`
	Wxlogourl               string  `orm:"wxlogourl" json:"wxlogourl"`
	BrandName               string  `orm:"brand_name" json:"brand_name"`
	CodeType                string  `orm:"code_type" json:"code_type"`
	Title                   string  `orm:"title" json:"title"`
	Color                   string  `orm:"color" json:"color"`
	Notice                  string  `orm:"notice" json:"notice"`
	ServicePhone            string  `orm:"service_phone" json:"service_phone"`
	Description             string  `orm:"description" json:"description"`
	Datetype                string  `orm:"datetype" json:"datetype"`
	BeginTimestamp          int     `orm:"begin_timestamp" json:"begin_timestamp"`
	EndTimestamp            int     `orm:"end_timestamp" json:"end_timestamp"`
	FixedTerm               int     `orm:"fixed_term" json:"fixed_term"`
	FixedBeginTerm          int     `orm:"fixed_begin_term" json:"fixed_begin_term"`
	Quantity                int     `orm:"quantity" json:"quantity"`
	TotalQuantity           string  `orm:"total_quantity" json:"total_quantity"`
	UseLimit                int     `orm:"use_limit" json:"use_limit"`
	GetLimit                int     `orm:"get_limit" json:"get_limit"`
	UseCustomCode           int     `orm:"use_custom_code" json:"use_custom_code"`
	BindOpenid              int     `orm:"bind_openid" json:"bind_openid"`
	CanShare                int     `orm:"can_share" json:"can_share"`
	CanGiveFriend           int     `orm:"can_give_friend" json:"can_give_friend"`
	CenterTitle             string  `orm:"center_title" json:"center_title"`
	CenterSubTitle          string  `orm:"center_sub_title" json:"center_sub_title"`
	CenterUrl               string  `orm:"center_url" json:"center_url"`
	Setcustom               int     `orm:"setcustom" json:"setcustom"`
	CustomUrlName           string  `orm:"custom_url_name" json:"custom_url_name"`
	CustomUrlSubTitle       string  `orm:"custom_url_sub_title" json:"custom_url_sub_title"`
	CustomUrl               string  `orm:"custom_url" json:"custom_url"`
	Setpromotion            int     `orm:"setpromotion" json:"setpromotion"`
	PromotionUrlName        string  `orm:"promotion_url_name" json:"promotion_url_name"`
	PromotionUrlSubTitle    string  `orm:"promotion_url_sub_title" json:"promotion_url_sub_title"`
	PromotionUrl            string  `orm:"promotion_url" json:"promotion_url"`
	Source                  string  `orm:"source" json:"source"`
	CanUseWithOtherDiscount int     `orm:"can_use_with_other_discount" json:"can_use_with_other_discount"`
	Setabstract             int     `orm:"setabstract" json:"setabstract"`
	Abstract                string  `orm:"abstract" json:"abstract"`
	Abstractimg             string  `orm:"abstractimg" json:"abstractimg"`
	IconUrlList             string  `orm:"icon_url_list" json:"icon_url_list"`
	AcceptCategory          string  `orm:"accept_category" json:"accept_category"`
	RejectCategory          string  `orm:"reject_category" json:"reject_category"`
	LeastCost               float64 `orm:"least_cost" json:"least_cost"`
	ReduceCost              float64 `orm:"reduce_cost" json:"reduce_cost"`
	Discount                float64 `orm:"discount" json:"discount"`
	Limitgoodtype           int     `orm:"limitgoodtype" json:"limitgoodtype"`
	Limitgoodcatetype       int     `orm:"limitgoodcatetype" json:"limitgoodcatetype"`
	Limitgoodcateids        string  `orm:"limitgoodcateids" json:"limitgoodcateids"`
	Limitgoodids            string  `orm:"limitgoodids" json:"limitgoodids"`
	Limitdiscounttype       int     `orm:"limitdiscounttype" json:"limitdiscounttype"`
	Merchid                 int     `orm:"merchid" json:"merchid"`
	Gettype                 int     `orm:"gettype" json:"gettype"`
	Islimitlevel            int     `orm:"islimitlevel" json:"islimitlevel"`
	Limitmemberlevels       string  `orm:"limitmemberlevels" json:"limitmemberlevels"`
	Limitagentlevels        string  `orm:"limitagentlevels" json:"limitagentlevels"`
	Limitpartnerlevels      string  `orm:"limitpartnerlevels" json:"limitpartnerlevels"`
	Limitaagentlevels       string  `orm:"limitaagentlevels" json:"limitaagentlevels"`
	Settitlecolor           int     `orm:"settitlecolor" json:"settitlecolor"`
	Titlecolor              string  `orm:"titlecolor" json:"titlecolor"`
	Tagtitle                string  `orm:"tagtitle" json:"tagtitle"`
	UseCondition            int     `orm:"use_condition" json:"use_condition"`
}

func (*Ewei_shop_wxcard) TableName() string {
	return "ewei_shop_wxcard"
}

type Stat_keyword struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Rid        string `orm:"rid" json:"rid"`
	Kid        int    `orm:"kid" json:"kid"`
	Hit        int    `orm:"hit" json:"hit"`
	Lastupdate int    `orm:"lastupdate" json:"lastupdate"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Stat_keyword) TableName() string {
	return "stat_keyword"
}

type Core_settings struct {
	Key   string `orm:"key" json:"key"`
	Value string `orm:"value" json:"value"`
}

func (*Core_settings) TableName() string {
	return "core_settings"
}

type Ewei_shop_abonus_level struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Levelname  string  `orm:"levelname" json:"levelname"`
	Bonus1     float64 `orm:"bonus1" json:"bonus1"`
	Bonus2     float64 `orm:"bonus2" json:"bonus2"`
	Bonus3     float64 `orm:"bonus3" json:"bonus3"`
	Ordermoney float64 `orm:"ordermoney" json:"ordermoney"`
	Ordercount int     `orm:"ordercount" json:"ordercount"`
	Bonusmoney float64 `orm:"bonusmoney" json:"bonusmoney"`
	Downcount  int     `orm:"downcount" json:"downcount"`
}

func (*Ewei_shop_abonus_level) TableName() string {
	return "ewei_shop_abonus_level"
}

type Ewei_shop_designer_menu struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Menuname   string `orm:"menuname" json:"menuname"`
	Isdefault  int    `orm:"isdefault" json:"isdefault"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Menus      string `orm:"menus" json:"menus"`
	Params     string `orm:"params" json:"params"`
}

func (*Ewei_shop_designer_menu) TableName() string {
	return "ewei_shop_designer_menu"
}

type Ewei_shop_live_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Merchid      int    `orm:"merchid" json:"merchid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_live_adv) TableName() string {
	return "ewei_shop_live_adv"
}

type Ewei_shop_member_log struct {
	Id             int     `orm:"id" json:"id"`
	Uniacid        int     `orm:"uniacid" json:"uniacid"`
	Openid         string  `orm:"openid" json:"openid"`
	Type           int     `orm:"type" json:"type"`
	Logno          string  `orm:"logno" json:"logno"`
	Title          string  `orm:"title" json:"title"`
	Createtime     int     `orm:"createtime" json:"createtime"`
	Status         int     `orm:"status" json:"status"`
	Money          float64 `orm:"money" json:"money"`
	Rechargetype   string  `orm:"rechargetype" json:"rechargetype"`
	Gives          float64 `orm:"gives" json:"gives"`
	Couponid       int     `orm:"couponid" json:"couponid"`
	Transid        string  `orm:"transid" json:"transid"`
	Realmoney      float64 `orm:"realmoney" json:"realmoney"`
	Charge         float64 `orm:"charge" json:"charge"`
	Deductionmoney float64 `orm:"deductionmoney" json:"deductionmoney"`
	Isborrow       int     `orm:"isborrow" json:"isborrow"`
	Borrowopenid   string  `orm:"borrowopenid" json:"borrowopenid"`
	Remark         string  `orm:"remark" json:"remark"`
	Apppay         int     `orm:"apppay" json:"apppay"`
	Alipay         string  `orm:"alipay" json:"alipay"`
	Bankname       string  `orm:"bankname" json:"bankname"`
	Bankcard       string  `orm:"bankcard" json:"bankcard"`
	Realname       string  `orm:"realname" json:"realname"`
	Applytype      int     `orm:"applytype" json:"applytype"`
	Sendmoney      float64 `orm:"sendmoney" json:"sendmoney"`
	Senddata       string  `orm:"senddata" json:"senddata"`
}

func (*Ewei_shop_member_log) TableName() string {
	return "ewei_shop_member_log"
}

type Ewei_shop_merch_perm_role struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Merchid  int    `orm:"merchid" json:"merchid"`
	Rolename string `orm:"rolename" json:"rolename"`
	Status   int    `orm:"status" json:"status"`
	Perms    string `orm:"perms" json:"perms"`
	Deleted  int    `orm:"deleted" json:"deleted"`
}

func (*Ewei_shop_merch_perm_role) TableName() string {
	return "ewei_shop_merch_perm_role"
}

type Ewei_shop_seckill_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_seckill_adv) TableName() string {
	return "ewei_shop_seckill_adv"
}

type Ewei_shop_system_adv struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Url          string `orm:"url" json:"url"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Module       string `orm:"module" json:"module"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_adv) TableName() string {
	return "ewei_shop_system_adv"
}

type Message_notice_log struct {
	Id         int    `orm:"id" json:"id"`
	Message    string `orm:"message" json:"message"`
	IsRead     int    `orm:"is_read" json:"is_read"`
	Uid        int    `orm:"uid" json:"uid"`
	Sign       string `orm:"sign" json:"sign"`
	Type       int    `orm:"type" json:"type"`
	Status     int    `orm:"status" json:"status"`
	CreateTime int    `orm:"create_time" json:"create_time"`
	EndTime    int    `orm:"end_time" json:"end_time"`
}

func (*Message_notice_log) TableName() string {
	return "message_notice_log"
}

type Ewei_shop_author_bill struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Billno          string  `orm:"billno" json:"billno"`
	Paytype         int     `orm:"paytype" json:"paytype"`
	Year            int     `orm:"year" json:"year"`
	Month           int     `orm:"month" json:"month"`
	Week            int     `orm:"week" json:"week"`
	Ordercount      int     `orm:"ordercount" json:"ordercount"`
	Ordermoney      float64 `orm:"ordermoney" json:"ordermoney"`
	Bonusordermoney float64 `orm:"bonusordermoney" json:"bonusordermoney"`
	Bonusrate       float64 `orm:"bonusrate" json:"bonusrate"`
	Bonusmoney      float64 `orm:"bonusmoney" json:"bonusmoney"`
	BonusmoneySend  float64 `orm:"bonusmoney_send" json:"bonusmoney_send"`
	BonusmoneyPay   float64 `orm:"bonusmoney_pay" json:"bonusmoney_pay"`
	Paytime         int     `orm:"paytime" json:"paytime"`
	Partnercount    int     `orm:"partnercount" json:"partnercount"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Status          int     `orm:"status" json:"status"`
	Starttime       int     `orm:"starttime" json:"starttime"`
	Endtime         int     `orm:"endtime" json:"endtime"`
	Confirmtime     int     `orm:"confirmtime" json:"confirmtime"`
}

func (*Ewei_shop_author_bill) TableName() string {
	return "ewei_shop_author_bill"
}

type Ewei_shop_commission_repurchase struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Openid     string  `orm:"openid" json:"openid"`
	Year       int     `orm:"year" json:"year"`
	Month      int     `orm:"month" json:"month"`
	Repurchase float64 `orm:"repurchase" json:"repurchase"`
	Applyid    int     `orm:"applyid" json:"applyid"`
}

func (*Ewei_shop_commission_repurchase) TableName() string {
	return "ewei_shop_commission_repurchase"
}

type Ewei_shop_lottery_join struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	JoinUser   string `orm:"join_user" json:"join_user"`
	LotteryId  int    `orm:"lottery_id" json:"lottery_id"`
	LotteryNum int    `orm:"lottery_num" json:"lottery_num"`
	LotteryTag string `orm:"lottery_tag" json:"lottery_tag"`
	Addtime    int    `orm:"addtime" json:"addtime"`
}

func (*Ewei_shop_lottery_join) TableName() string {
	return "ewei_shop_lottery_join"
}

type Ewei_shop_perm_log struct {
	Id         int    `orm:"id" json:"id"`
	Uid        int    `orm:"uid" json:"uid"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Name       string `orm:"name" json:"name"`
	Type       string `orm:"type" json:"type"`
	Op         string `orm:"op" json:"op"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Ip         string `orm:"ip" json:"ip"`
}

func (*Ewei_shop_perm_log) TableName() string {
	return "ewei_shop_perm_log"
}

type Users_failed_login struct {
	Id         int    `orm:"id" json:"id"`
	Ip         string `orm:"ip" json:"ip"`
	Username   string `orm:"username" json:"username"`
	Count      int    `orm:"count" json:"count"`
	Lastupdate int    `orm:"lastupdate" json:"lastupdate"`
}

func (*Users_failed_login) TableName() string {
	return "users_failed_login"
}

type Ewei_message_mass_task struct {
	Id             int    `orm:"id" json:"id"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Title          string `orm:"title" json:"title"`
	Status         int    `orm:"status" json:"status"`
	Processnum     int    `orm:"processnum" json:"processnum"`
	Sendnum        int    `orm:"sendnum" json:"sendnum"`
	Messagetype    int    `orm:"messagetype" json:"messagetype"`
	Templateid     int    `orm:"templateid" json:"templateid"`
	Resptitle      string `orm:"resptitle" json:"resptitle"`
	Respthumb      string `orm:"respthumb" json:"respthumb"`
	Respdesc       string `orm:"respdesc" json:"respdesc"`
	Respurl        string `orm:"respurl" json:"respurl"`
	Sendlimittype  int    `orm:"sendlimittype" json:"sendlimittype"`
	SendOpenid     string `orm:"send_openid" json:"send_openid"`
	SendLevel      int    `orm:"send_level" json:"send_level"`
	SendGroup      int    `orm:"send_group" json:"send_group"`
	SendAgentlevel int    `orm:"send_agentlevel" json:"send_agentlevel"`
	Customertype   int    `orm:"customertype" json:"customertype"`
	Resdesc2       string `orm:"resdesc2" json:"resdesc2"`
	Pagecount      int    `orm:"pagecount" json:"pagecount"`
	Successnum     int    `orm:"successnum" json:"successnum"`
	Failnum        int    `orm:"failnum" json:"failnum"`
}

func (*Ewei_message_mass_task) TableName() string {
	return "ewei_message_mass_task"
}

type Ewei_shop_author_billp struct {
	Id          int     `orm:"id" json:"id"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Billid      int     `orm:"billid" json:"billid"`
	Openid      string  `orm:"openid" json:"openid"`
	Payno       string  `orm:"payno" json:"payno"`
	Paytype     int     `orm:"paytype" json:"paytype"`
	Bonus       float64 `orm:"bonus" json:"bonus"`
	Money       float64 `orm:"money" json:"money"`
	Realmoney   float64 `orm:"realmoney" json:"realmoney"`
	Paymoney    float64 `orm:"paymoney" json:"paymoney"`
	Charge      float64 `orm:"charge" json:"charge"`
	Chargemoney float64 `orm:"chargemoney" json:"chargemoney"`
	Status      int     `orm:"status" json:"status"`
	Reason      string  `orm:"reason" json:"reason"`
	Paytime     int     `orm:"paytime" json:"paytime"`
}

func (*Ewei_shop_author_billp) TableName() string {
	return "ewei_shop_author_billp"
}

type Ewei_shop_cashier_user struct {
	Id               int     `orm:"id" json:"id"`
	Uniacid          int     `orm:"uniacid" json:"uniacid"`
	Storeid          int     `orm:"storeid" json:"storeid"`
	Merchid          int     `orm:"merchid" json:"merchid"`
	Setmeal          int     `orm:"setmeal" json:"setmeal"`
	Title            string  `orm:"title" json:"title"`
	Logo             string  `orm:"logo" json:"logo"`
	Manageopenid     string  `orm:"manageopenid" json:"manageopenid"`
	IsopenCommission int     `orm:"isopen_commission" json:"isopen_commission"`
	Name             string  `orm:"name" json:"name"`
	Mobile           string  `orm:"mobile" json:"mobile"`
	Categoryid       int     `orm:"categoryid" json:"categoryid"`
	WechatStatus     int     `orm:"wechat_status" json:"wechat_status"`
	Wechatpay        string  `orm:"wechatpay" json:"wechatpay"`
	AlipayStatus     int     `orm:"alipay_status" json:"alipay_status"`
	Alipay           string  `orm:"alipay" json:"alipay"`
	Withdraw         float64 `orm:"withdraw" json:"withdraw"`
	Openid           string  `orm:"openid" json:"openid"`
	Diyformfields    string  `orm:"diyformfields" json:"diyformfields"`
	Diyformdata      string  `orm:"diyformdata" json:"diyformdata"`
	Createtime       int     `orm:"createtime" json:"createtime"`
	Username         string  `orm:"username" json:"username"`
	Password         string  `orm:"password" json:"password"`
	Salt             string  `orm:"salt" json:"salt"`
	Lifetimestart    int     `orm:"lifetimestart" json:"lifetimestart"`
	Lifetimeend      int     `orm:"lifetimeend" json:"lifetimeend"`
	Status           int     `orm:"status" json:"status"`
	Set              string  `orm:"set" json:"set"`
	Deleted          int     `orm:"deleted" json:"deleted"`
	CanWithdraw      int     `orm:"can_withdraw" json:"can_withdraw"`
	ShowPaytype      int     `orm:"show_paytype" json:"show_paytype"`
	Couponid         string  `orm:"couponid" json:"couponid"`
	Management       string  `orm:"management" json:"management"`
	NoticeOpenids    string  `orm:"notice_openids" json:"notice_openids"`
}

func (*Ewei_shop_cashier_user) TableName() string {
	return "ewei_shop_cashier_user"
}

type Ewei_shop_commission_shop struct {
	Id             int    `orm:"id" json:"id"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Mid            int    `orm:"mid" json:"mid"`
	Name           string `orm:"name" json:"name"`
	Logo           string `orm:"logo" json:"logo"`
	Img            string `orm:"img" json:"img"`
	Desc           string `orm:"desc" json:"desc"`
	Selectgoods    int    `orm:"selectgoods" json:"selectgoods"`
	Selectcategory int    `orm:"selectcategory" json:"selectcategory"`
	Goodsids       string `orm:"goodsids" json:"goodsids"`
}

func (*Ewei_shop_commission_shop) TableName() string {
	return "ewei_shop_commission_shop"
}

type Ewei_shop_exchange_query struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Querykey   string `orm:"querykey" json:"querykey"`
	Querytime  int    `orm:"querytime" json:"querytime"`
	Unfreeze   int    `orm:"unfreeze" json:"unfreeze"`
	Errorcount int    `orm:"errorcount" json:"errorcount"`
}

func (*Ewei_shop_exchange_query) TableName() string {
	return "ewei_shop_exchange_query"
}

type Ewei_shop_exhelper_express struct {
	Id          int     `orm:"id" json:"id"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Type        int     `orm:"type" json:"type"`
	Expressname string  `orm:"expressname" json:"expressname"`
	Expresscom  string  `orm:"expresscom" json:"expresscom"`
	Express     string  `orm:"express" json:"express"`
	Width       float64 `orm:"width" json:"width"`
	Datas       string  `orm:"datas" json:"datas"`
	Height      float64 `orm:"height" json:"height"`
	Bg          string  `orm:"bg" json:"bg"`
	Isdefault   int     `orm:"isdefault" json:"isdefault"`
	Merchid     int     `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_exhelper_express) TableName() string {
	return "ewei_shop_exhelper_express"
}

type Ewei_shop_goods struct {
	Id                    int     `orm:"id" json:"id"`
	Uniacid               int     `orm:"uniacid" json:"uniacid"`
	Pcate                 int     `orm:"pcate" json:"pcate"`
	Ccate                 int     `orm:"ccate" json:"ccate"`
	Type                  int     `orm:"type" json:"type"`
	Status                int     `orm:"status" json:"status"`
	Displayorder          int     `orm:"displayorder" json:"displayorder"`
	Title                 string  `orm:"title" json:"title"`
	Thumb                 string  `orm:"thumb" json:"thumb"`
	Unit                  string  `orm:"unit" json:"unit"`
	Description           string  `orm:"description" json:"description"`
	Content               string  `orm:"content" json:"content"`
	Goodssn               string  `orm:"goodssn" json:"goodssn"`
	Productsn             string  `orm:"productsn" json:"productsn"`
	Productprice          float64 `orm:"productprice" json:"productprice"`
	Marketprice           float64 `orm:"marketprice" json:"marketprice"`
	Costprice             float64 `orm:"costprice" json:"costprice"`
	Originalprice         float64 `orm:"originalprice" json:"originalprice"`
	Total                 int     `orm:"total" json:"total"`
	Totalcnf              int     `orm:"totalcnf" json:"totalcnf"`
	Sales                 int     `orm:"sales" json:"sales"`
	Salesreal             int     `orm:"salesreal" json:"salesreal"`
	Spec                  string  `orm:"spec" json:"spec"`
	Createtime            int     `orm:"createtime" json:"createtime"`
	Weight                float64 `orm:"weight" json:"weight"`
	Credit                string  `orm:"credit" json:"credit"`
	Maxbuy                int     `orm:"maxbuy" json:"maxbuy"`
	Usermaxbuy            int     `orm:"usermaxbuy" json:"usermaxbuy"`
	Hasoption             int     `orm:"hasoption" json:"hasoption"`
	Dispatch              int     `orm:"dispatch" json:"dispatch"`
	ThumbUrl              string  `orm:"thumb_url" json:"thumb_url"`
	Isnew                 int     `orm:"isnew" json:"isnew"`
	Ishot                 int     `orm:"ishot" json:"ishot"`
	Isdiscount            int     `orm:"isdiscount" json:"isdiscount"`
	Isrecommand           int     `orm:"isrecommand" json:"isrecommand"`
	Issendfree            int     `orm:"issendfree" json:"issendfree"`
	Istime                int     `orm:"istime" json:"istime"`
	Iscomment             int     `orm:"iscomment" json:"iscomment"`
	Timestart             int     `orm:"timestart" json:"timestart"`
	Timeend               int     `orm:"timeend" json:"timeend"`
	Viewcount             int     `orm:"viewcount" json:"viewcount"`
	Deleted               int     `orm:"deleted" json:"deleted"`
	Hascommission         int     `orm:"hascommission" json:"hascommission"`
	Commission1Rate       float64 `orm:"commission1_rate" json:"commission1_rate"`
	Commission1Pay        float64 `orm:"commission1_pay" json:"commission1_pay"`
	Commission2Rate       float64 `orm:"commission2_rate" json:"commission2_rate"`
	Commission2Pay        float64 `orm:"commission2_pay" json:"commission2_pay"`
	Commission3Rate       float64 `orm:"commission3_rate" json:"commission3_rate"`
	Commission3Pay        float64 `orm:"commission3_pay" json:"commission3_pay"`
	Score                 float64 `orm:"score" json:"score"`
	Taobaoid              string  `orm:"taobaoid" json:"taobaoid"`
	Taotaoid              string  `orm:"taotaoid" json:"taotaoid"`
	Taobaourl             string  `orm:"taobaourl" json:"taobaourl"`
	Updatetime            int     `orm:"updatetime" json:"updatetime"`
	ShareTitle            string  `orm:"share_title" json:"share_title"`
	ShareIcon             string  `orm:"share_icon" json:"share_icon"`
	Cash                  int     `orm:"cash" json:"cash"`
	CommissionThumb       string  `orm:"commission_thumb" json:"commission_thumb"`
	Isnodiscount          int     `orm:"isnodiscount" json:"isnodiscount"`
	Showlevels            string  `orm:"showlevels" json:"showlevels"`
	Buylevels             string  `orm:"buylevels" json:"buylevels"`
	Showgroups            string  `orm:"showgroups" json:"showgroups"`
	Buygroups             string  `orm:"buygroups" json:"buygroups"`
	Isverify              int     `orm:"isverify" json:"isverify"`
	Storeids              string  `orm:"storeids" json:"storeids"`
	Noticeopenid          string  `orm:"noticeopenid" json:"noticeopenid"`
	Tcate                 int     `orm:"tcate" json:"tcate"`
	Noticetype            string  `orm:"noticetype" json:"noticetype"`
	Needfollow            int     `orm:"needfollow" json:"needfollow"`
	Followtip             string  `orm:"followtip" json:"followtip"`
	Followurl             string  `orm:"followurl" json:"followurl"`
	Deduct                float64 `orm:"deduct" json:"deduct"`
	Virtual               int     `orm:"virtual" json:"virtual"`
	Ccates                string  `orm:"ccates" json:"ccates"`
	Discounts             string  `orm:"discounts" json:"discounts"`
	Nocommission          int     `orm:"nocommission" json:"nocommission"`
	Hidecommission        int     `orm:"hidecommission" json:"hidecommission"`
	Pcates                string  `orm:"pcates" json:"pcates"`
	Tcates                string  `orm:"tcates" json:"tcates"`
	Cates                 string  `orm:"cates" json:"cates"`
	Artid                 int     `orm:"artid" json:"artid"`
	DetailLogo            string  `orm:"detail_logo" json:"detail_logo"`
	DetailShopname        string  `orm:"detail_shopname" json:"detail_shopname"`
	DetailBtntext1        string  `orm:"detail_btntext1" json:"detail_btntext1"`
	DetailBtnurl1         string  `orm:"detail_btnurl1" json:"detail_btnurl1"`
	DetailBtntext2        string  `orm:"detail_btntext2" json:"detail_btntext2"`
	DetailBtnurl2         string  `orm:"detail_btnurl2" json:"detail_btnurl2"`
	DetailTotaltitle      string  `orm:"detail_totaltitle" json:"detail_totaltitle"`
	Saleupdate42392       int     `orm:"saleupdate42392" json:"saleupdate42392"`
	Deduct2               float64 `orm:"deduct2" json:"deduct2"`
	Ednum                 int     `orm:"ednum" json:"ednum"`
	Edmoney               float64 `orm:"edmoney" json:"edmoney"`
	Edareas               string  `orm:"edareas" json:"edareas"`
	Diyformtype           int     `orm:"diyformtype" json:"diyformtype"`
	Diyformid             int     `orm:"diyformid" json:"diyformid"`
	Diymode               int     `orm:"diymode" json:"diymode"`
	Dispatchtype          int     `orm:"dispatchtype" json:"dispatchtype"`
	Dispatchid            int     `orm:"dispatchid" json:"dispatchid"`
	Dispatchprice         float64 `orm:"dispatchprice" json:"dispatchprice"`
	Manydeduct            int     `orm:"manydeduct" json:"manydeduct"`
	Shorttitle            string  `orm:"shorttitle" json:"shorttitle"`
	IsdiscountTitle       string  `orm:"isdiscount_title" json:"isdiscount_title"`
	IsdiscountTime        int     `orm:"isdiscount_time" json:"isdiscount_time"`
	IsdiscountDiscounts   string  `orm:"isdiscount_discounts" json:"isdiscount_discounts"`
	Commission            string  `orm:"commission" json:"commission"`
	Saleupdate37975       int     `orm:"saleupdate37975" json:"saleupdate37975"`
	Shopid                int     `orm:"shopid" json:"shopid"`
	Allcates              string  `orm:"allcates" json:"allcates"`
	Minbuy                int     `orm:"minbuy" json:"minbuy"`
	Invoice               int     `orm:"invoice" json:"invoice"`
	Repair                int     `orm:"repair" json:"repair"`
	Seven                 int     `orm:"seven" json:"seven"`
	Money                 string  `orm:"money" json:"money"`
	Minprice              float64 `orm:"minprice" json:"minprice"`
	Maxprice              float64 `orm:"maxprice" json:"maxprice"`
	Province              string  `orm:"province" json:"province"`
	City                  string  `orm:"city" json:"city"`
	Buyshow               int     `orm:"buyshow" json:"buyshow"`
	Buycontent            string  `orm:"buycontent" json:"buycontent"`
	Saleupdate51117       int     `orm:"saleupdate51117" json:"saleupdate51117"`
	Virtualsend           int     `orm:"virtualsend" json:"virtualsend"`
	Virtualsendcontent    string  `orm:"virtualsendcontent" json:"virtualsendcontent"`
	Verifytype            int     `orm:"verifytype" json:"verifytype"`
	Diyfields             string  `orm:"diyfields" json:"diyfields"`
	Diysaveid             int     `orm:"diysaveid" json:"diysaveid"`
	Diysave               int     `orm:"diysave" json:"diysave"`
	Quality               int     `orm:"quality" json:"quality"`
	Groupstype            int     `orm:"groupstype" json:"groupstype"`
	Showtotal             int     `orm:"showtotal" json:"showtotal"`
	Subtitle              string  `orm:"subtitle" json:"subtitle"`
	Minpriceupdated       int     `orm:"minpriceupdated" json:"minpriceupdated"`
	Sharebtn              int     `orm:"sharebtn" json:"sharebtn"`
	Catesinit3            string  `orm:"catesinit3" json:"catesinit3"`
	Showtotaladd          int     `orm:"showtotaladd" json:"showtotaladd"`
	Merchid               int     `orm:"merchid" json:"merchid"`
	Checked               int     `orm:"checked" json:"checked"`
	ThumbFirst            int     `orm:"thumb_first" json:"thumb_first"`
	Merchsale             int     `orm:"merchsale" json:"merchsale"`
	Keywords              string  `orm:"keywords" json:"keywords"`
	CatchId               string  `orm:"catch_id" json:"catch_id"`
	CatchUrl              string  `orm:"catch_url" json:"catch_url"`
	CatchSource           string  `orm:"catch_source" json:"catch_source"`
	Saleupdate40170       int     `orm:"saleupdate40170" json:"saleupdate40170"`
	Saleupdate35843       int     `orm:"saleupdate35843" json:"saleupdate35843"`
	Labelname             string  `orm:"labelname" json:"labelname"`
	Autoreceive           int     `orm:"autoreceive" json:"autoreceive"`
	Cannotrefund          int     `orm:"cannotrefund" json:"cannotrefund"`
	Saleupdate33219       int     `orm:"saleupdate33219" json:"saleupdate33219"`
	Bargain               int     `orm:"bargain" json:"bargain"`
	Buyagain              float64 `orm:"buyagain" json:"buyagain"`
	BuyagainIslong        int     `orm:"buyagain_islong" json:"buyagain_islong"`
	BuyagainCondition     int     `orm:"buyagain_condition" json:"buyagain_condition"`
	BuyagainSale          int     `orm:"buyagain_sale" json:"buyagain_sale"`
	BuyagainCommission    string  `orm:"buyagain_commission" json:"buyagain_commission"`
	Saleupdate32484       int     `orm:"saleupdate32484" json:"saleupdate32484"`
	Saleupdate36586       int     `orm:"saleupdate36586" json:"saleupdate36586"`
	Diypage               int     `orm:"diypage" json:"diypage"`
	Cashier               int     `orm:"cashier" json:"cashier"`
	Saleupdate53481       int     `orm:"saleupdate53481" json:"saleupdate53481"`
	Saleupdate30424       int     `orm:"saleupdate30424" json:"saleupdate30424"`
	Isendtime             int     `orm:"isendtime" json:"isendtime"`
	Usetime               int     `orm:"usetime" json:"usetime"`
	Endtime               int     `orm:"endtime" json:"endtime"`
	Merchdisplayorder     int     `orm:"merchdisplayorder" json:"merchdisplayorder"`
	ExchangeStock         int     `orm:"exchange_stock" json:"exchange_stock"`
	ExchangePostage       float64 `orm:"exchange_postage" json:"exchange_postage"`
	Ispresell             int     `orm:"ispresell" json:"ispresell"`
	Presellprice          float64 `orm:"presellprice" json:"presellprice"`
	Presellover           int     `orm:"presellover" json:"presellover"`
	Presellovertime       int     `orm:"presellovertime" json:"presellovertime"`
	Presellstart          int     `orm:"presellstart" json:"presellstart"`
	Preselltimestart      int     `orm:"preselltimestart" json:"preselltimestart"`
	Presellend            int     `orm:"presellend" json:"presellend"`
	Preselltimeend        int     `orm:"preselltimeend" json:"preselltimeend"`
	Presellsendtype       int     `orm:"presellsendtype" json:"presellsendtype"`
	Presellsendstatrttime int     `orm:"presellsendstatrttime" json:"presellsendstatrttime"`
	Presellsendtime       int     `orm:"presellsendtime" json:"presellsendtime"`
	EdareasCode           string  `orm:"edareas_code" json:"edareas_code"`
	UniteTotal            int     `orm:"unite_total" json:"unite_total"`
	BuyagainPrice         float64 `orm:"buyagain_price" json:"buyagain_price"`
	Threen                string  `orm:"threen" json:"threen"`
	Intervalfloor         int     `orm:"intervalfloor" json:"intervalfloor"`
	Intervalprice         string  `orm:"intervalprice" json:"intervalprice"`
	Isfullback            int     `orm:"isfullback" json:"isfullback"`
	Isstatustime          int     `orm:"isstatustime" json:"isstatustime"`
	Statustimestart       int     `orm:"statustimestart" json:"statustimestart"`
	Statustimeend         int     `orm:"statustimeend" json:"statustimeend"`
	Nosearch              int     `orm:"nosearch" json:"nosearch"`
	Showsales             int     `orm:"showsales" json:"showsales"`
	Islive                int     `orm:"islive" json:"islive"`
	Liveprice             float64 `orm:"liveprice" json:"liveprice"`
	Opencard              int     `orm:"opencard" json:"opencard"`
	Cardid                string  `orm:"cardid" json:"cardid"`
	Verifygoodsnum        int     `orm:"verifygoodsnum" json:"verifygoodsnum"`
	Verifygoodsdays       int     `orm:"verifygoodsdays" json:"verifygoodsdays"`
	Verifygoodslimittype  int     `orm:"verifygoodslimittype" json:"verifygoodslimittype"`
	Verifygoodslimitdate  int     `orm:"verifygoodslimitdate" json:"verifygoodslimitdate"`
	Minliveprice          float64 `orm:"minliveprice" json:"minliveprice"`
	Maxliveprice          float64 `orm:"maxliveprice" json:"maxliveprice"`
	Dowpayment            float64 `orm:"dowpayment" json:"dowpayment"`
	Tempid                int     `orm:"tempid" json:"tempid"`
	Isstoreprice          int     `orm:"isstoreprice" json:"isstoreprice"`
	Beforehours           int     `orm:"beforehours" json:"beforehours"`
	Saleupdate            int     `orm:"saleupdate" json:"saleupdate"`
	Newgoods              int     `orm:"newgoods" json:"newgoods"`
	Video                 string  `orm:"video" json:"video"`
	Officthumb            string  `orm:"officthumb" json:"officthumb"`
}

func (*Ewei_shop_goods) TableName() string {
	return "ewei_shop_goods"
}

type Ewei_shop_live_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Advimg       string `orm:"advimg" json:"advimg"`
	Advurl       string `orm:"advurl" json:"advurl"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
}

func (*Ewei_shop_live_category) TableName() string {
	return "ewei_shop_live_category"
}

type Ewei_shop_merch_category_swipe struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Title        string `orm:"title" json:"title"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Thumb        string `orm:"thumb" json:"thumb"`
}

func (*Ewei_shop_merch_category_swipe) TableName() string {
	return "ewei_shop_merch_category_swipe"
}

type Cover_reply struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Multiid     int    `orm:"multiid" json:"multiid"`
	Rid         int    `orm:"rid" json:"rid"`
	Module      string `orm:"module" json:"module"`
	Do          string `orm:"do" json:"do"`
	Title       string `orm:"title" json:"title"`
	Description string `orm:"description" json:"description"`
	Thumb       string `orm:"thumb" json:"thumb"`
	Url         string `orm:"url" json:"url"`
}

func (*Cover_reply) TableName() string {
	return "cover_reply"
}

type Ewei_shop_cashier_goods_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Cashierid    int    `orm:"cashierid" json:"cashierid"`
	Catename     string `orm:"catename" json:"catename"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_cashier_goods_category) TableName() string {
	return "ewei_shop_cashier_goods_category"
}

type Ewei_shop_cashier_order struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Ordersn    string  `orm:"ordersn" json:"ordersn"`
	Price      float64 `orm:"price" json:"price"`
	Openid     string  `orm:"openid" json:"openid"`
	Payopenid  string  `orm:"payopenid" json:"payopenid"`
	Createtime int     `orm:"createtime" json:"createtime"`
	Status     int     `orm:"status" json:"status"`
	Paytime    int     `orm:"paytime" json:"paytime"`
}

func (*Ewei_shop_cashier_order) TableName() string {
	return "ewei_shop_cashier_order"
}

type Ewei_shop_creditshop_spec struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Goodsid      int    `orm:"goodsid" json:"goodsid"`
	Title        string `orm:"title" json:"title"`
	Description  string `orm:"description" json:"description"`
	Displaytype  int    `orm:"displaytype" json:"displaytype"`
	Content      string `orm:"content" json:"content"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	PropId       string `orm:"propId" json:"propId"`
}

func (*Ewei_shop_creditshop_spec) TableName() string {
	return "ewei_shop_creditshop_spec"
}

type Ewei_shop_customer_category struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
}

func (*Ewei_shop_customer_category) TableName() string {
	return "ewei_shop_customer_category"
}

type Ewei_shop_postera struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Type          int     `orm:"type" json:"type"`
	Days          int     `orm:"days" json:"days"`
	Title         string  `orm:"title" json:"title"`
	Bg            string  `orm:"bg" json:"bg"`
	Data          string  `orm:"data" json:"data"`
	Keyword       string  `orm:"keyword" json:"keyword"`
	Isdefault     int     `orm:"isdefault" json:"isdefault"`
	Resptitle     string  `orm:"resptitle" json:"resptitle"`
	Respthumb     string  `orm:"respthumb" json:"respthumb"`
	Createtime    int     `orm:"createtime" json:"createtime"`
	Respdesc      string  `orm:"respdesc" json:"respdesc"`
	Respurl       string  `orm:"respurl" json:"respurl"`
	Waittext      string  `orm:"waittext" json:"waittext"`
	Oktext        string  `orm:"oktext" json:"oktext"`
	Subcredit     int     `orm:"subcredit" json:"subcredit"`
	Submoney      float64 `orm:"submoney" json:"submoney"`
	Reccredit     int     `orm:"reccredit" json:"reccredit"`
	Recmoney      float64 `orm:"recmoney" json:"recmoney"`
	Scantext      string  `orm:"scantext" json:"scantext"`
	Subtext       string  `orm:"subtext" json:"subtext"`
	Beagent       int     `orm:"beagent" json:"beagent"`
	Bedown        int     `orm:"bedown" json:"bedown"`
	Isopen        int     `orm:"isopen" json:"isopen"`
	Opentext      string  `orm:"opentext" json:"opentext"`
	Openurl       string  `orm:"openurl" json:"openurl"`
	Paytype       int     `orm:"paytype" json:"paytype"`
	Subpaycontent string  `orm:"subpaycontent" json:"subpaycontent"`
	Recpaycontent string  `orm:"recpaycontent" json:"recpaycontent"`
	Templateid    string  `orm:"templateid" json:"templateid"`
	Entrytext     string  `orm:"entrytext" json:"entrytext"`
	Reccouponid   int     `orm:"reccouponid" json:"reccouponid"`
	Reccouponnum  int     `orm:"reccouponnum" json:"reccouponnum"`
	Subcouponid   int     `orm:"subcouponid" json:"subcouponid"`
	Subcouponnum  int     `orm:"subcouponnum" json:"subcouponnum"`
	Timestart     int     `orm:"timestart" json:"timestart"`
	Timeend       int     `orm:"timeend" json:"timeend"`
	Status        int     `orm:"status" json:"status"`
	Goodsid       int     `orm:"goodsid" json:"goodsid"`
	Starttext     string  `orm:"starttext" json:"starttext"`
	Endtext       string  `orm:"endtext" json:"endtext"`
	Resptype      int     `orm:"resptype" json:"resptype"`
	Resptext      string  `orm:"resptext" json:"resptext"`
	Testflag      int     `orm:"testflag" json:"testflag"`
	Keyword2      string  `orm:"keyword2" json:"keyword2"`
	RewardTotle   string  `orm:"reward_totle" json:"reward_totle"`
}

func (*Ewei_shop_postera) TableName() string {
	return "ewei_shop_postera"
}

type Ewei_shop_task_adv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_task_adv) TableName() string {
	return "ewei_shop_task_adv"
}

type Ewei_shop_version struct {
	Id      int `orm:"id" json:"id"`
	Uid     int `orm:"uid" json:"uid"`
	Type    int `orm:"type" json:"type"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Version int `orm:"version" json:"version"`
}

func (*Ewei_shop_version) TableName() string {
	return "ewei_shop_version"
}

type Site_store_order struct {
	Id          int     `orm:"id" json:"id"`
	Orderid     string  `orm:"orderid" json:"orderid"`
	Goodsid     int     `orm:"goodsid" json:"goodsid"`
	Duration    int     `orm:"duration" json:"duration"`
	Buyer       string  `orm:"buyer" json:"buyer"`
	Buyerid     int     `orm:"buyerid" json:"buyerid"`
	Amount      float64 `orm:"amount" json:"amount"`
	Type        int     `orm:"type" json:"type"`
	Changeprice int     `orm:"changeprice" json:"changeprice"`
	Createtime  int     `orm:"createtime" json:"createtime"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Endtime     int     `orm:"endtime" json:"endtime"`
	Wxapp       int     `orm:"wxapp" json:"wxapp"`
}

func (*Site_store_order) TableName() string {
	return "site_store_order"
}

type Core_sendsms_log struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Mobile     string `orm:"mobile" json:"mobile"`
	Content    string `orm:"content" json:"content"`
	Result     string `orm:"result" json:"result"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Core_sendsms_log) TableName() string {
	return "core_sendsms_log"
}

type Ewei_shop_article_log struct {
	Id      int    `orm:"id" json:"id"`
	Aid     int    `orm:"aid" json:"aid"`
	Read    int    `orm:"read" json:"read"`
	Like    int    `orm:"like" json:"like"`
	Openid  string `orm:"openid" json:"openid"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
}

func (*Ewei_shop_article_log) TableName() string {
	return "ewei_shop_article_log"
}

type Ewei_shop_goodscode_good struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Goodsid      int    `orm:"goodsid" json:"goodsid"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Qrcode       string `orm:"qrcode" json:"qrcode"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_goodscode_good) TableName() string {
	return "ewei_shop_goodscode_good"
}

type Ewei_shop_member_history struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Goodsid    int    `orm:"goodsid" json:"goodsid"`
	Openid     string `orm:"openid" json:"openid"`
	Deleted    int    `orm:"deleted" json:"deleted"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Times      int    `orm:"times" json:"times"`
	Merchid    int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_member_history) TableName() string {
	return "ewei_shop_member_history"
}

type Core_performance struct {
	Id         int    `orm:"id" json:"id"`
	Type       int    `orm:"type" json:"type"`
	Runtime    string `orm:"runtime" json:"runtime"`
	Runurl     string `orm:"runurl" json:"runurl"`
	Runsql     string `orm:"runsql" json:"runsql"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Core_performance) TableName() string {
	return "core_performance"
}

type Ewei_shop_express_cache struct {
	Id        int    `orm:"id" json:"id"`
	Expresssn string `orm:"expresssn" json:"expresssn"`
	Express   string `orm:"express" json:"express"`
	Lasttime  int    `orm:"lasttime" json:"lasttime"`
	Datas     string `orm:"datas" json:"datas"`
}

func (*Ewei_shop_express_cache) TableName() string {
	return "ewei_shop_express_cache"
}

type Ewei_shop_goods_label struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Label        string `orm:"label" json:"label"`
	Labelname    string `orm:"labelname" json:"labelname"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_goods_label) TableName() string {
	return "ewei_shop_goods_label"
}

type Ewei_shop_member_message_template struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Title       string `orm:"title" json:"title"`
	TemplateId  string `orm:"template_id" json:"template_id"`
	First       string `orm:"first" json:"first"`
	Firstcolor  string `orm:"firstcolor" json:"firstcolor"`
	Data        string `orm:"data" json:"data"`
	Remark      string `orm:"remark" json:"remark"`
	Remarkcolor string `orm:"remarkcolor" json:"remarkcolor"`
	Url         string `orm:"url" json:"url"`
	Createtime  int    `orm:"createtime" json:"createtime"`
	Sendtimes   int    `orm:"sendtimes" json:"sendtimes"`
	Sendcount   int    `orm:"sendcount" json:"sendcount"`
	Typecode    string `orm:"typecode" json:"typecode"`
	Messagetype int    `orm:"messagetype" json:"messagetype"`
	SendDesc    string `orm:"send_desc" json:"send_desc"`
}

func (*Ewei_shop_member_message_template) TableName() string {
	return "ewei_shop_member_message_template"
}

type Ewei_shop_merch_notice struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Link         string `orm:"link" json:"link"`
	Detail       string `orm:"detail" json:"detail"`
	Status       int    `orm:"status" json:"status"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_notice) TableName() string {
	return "ewei_shop_merch_notice"
}

type Ewei_shop_poster_qr struct {
	Id         int    `orm:"id" json:"id"`
	Acid       int    `orm:"acid" json:"acid"`
	Openid     string `orm:"openid" json:"openid"`
	Type       int    `orm:"type" json:"type"`
	Sceneid    int    `orm:"sceneid" json:"sceneid"`
	Mediaid    string `orm:"mediaid" json:"mediaid"`
	Ticket     string `orm:"ticket" json:"ticket"`
	Url        string `orm:"url" json:"url"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Goodsid    int    `orm:"goodsid" json:"goodsid"`
	Qrimg      string `orm:"qrimg" json:"qrimg"`
	Scenestr   string `orm:"scenestr" json:"scenestr"`
	Posterid   int    `orm:"posterid" json:"posterid"`
}

func (*Ewei_shop_poster_qr) TableName() string {
	return "ewei_shop_poster_qr"
}

type Ewei_shop_sendticket_draw struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Cpid       string `orm:"cpid" json:"cpid"`
	Openid     string `orm:"openid" json:"openid"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_sendticket_draw) TableName() string {
	return "ewei_shop_sendticket_draw"
}

type Ewei_shop_goods_cards struct {
	Id                  int    `orm:"id" json:"id"`
	Uniacid             int    `orm:"uniacid" json:"uniacid"`
	CardId              string `orm:"card_id" json:"card_id"`
	CardTitle           string `orm:"card_title" json:"card_title"`
	CardBrandName       string `orm:"card_brand_name" json:"card_brand_name"`
	CardTotalquantity   int    `orm:"card_totalquantity" json:"card_totalquantity"`
	CardQuantity        int    `orm:"card_quantity" json:"card_quantity"`
	CardLogoimg         string `orm:"card_logoimg" json:"card_logoimg"`
	CardLogowxurl       string `orm:"card_logowxurl" json:"card_logowxurl"`
	CardBackgroundtype  int    `orm:"card_backgroundtype" json:"card_backgroundtype"`
	Color               string `orm:"color" json:"color"`
	CardBackgroundimg   string `orm:"card_backgroundimg" json:"card_backgroundimg"`
	CardBackgroundwxurl string `orm:"card_backgroundwxurl" json:"card_backgroundwxurl"`
	Prerogative         string `orm:"prerogative" json:"prerogative"`
	CardDescription     string `orm:"card_description" json:"card_description"`
	Freewifi            int    `orm:"freewifi" json:"freewifi"`
	Withpet             int    `orm:"withpet" json:"withpet"`
	Freepark            int    `orm:"freepark" json:"freepark"`
	Deliver             int    `orm:"deliver" json:"deliver"`
	CustomCell1         int    `orm:"custom_cell1" json:"custom_cell1"`
	CustomCell1Name     string `orm:"custom_cell1_name" json:"custom_cell1_name"`
	CustomCell1Tips     string `orm:"custom_cell1_tips" json:"custom_cell1_tips"`
	CustomCell1Url      string `orm:"custom_cell1_url" json:"custom_cell1_url"`
	Color2              string `orm:"color2" json:"color2"`
}

func (*Ewei_shop_goods_cards) TableName() string {
	return "ewei_shop_goods_cards"
}

type Ewei_shop_groups_order_refund struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Openid          string  `orm:"openid" json:"openid"`
	Orderid         int     `orm:"orderid" json:"orderid"`
	Refundno        string  `orm:"refundno" json:"refundno"`
	Refundstatus    int     `orm:"refundstatus" json:"refundstatus"`
	Refundaddressid int     `orm:"refundaddressid" json:"refundaddressid"`
	Refundaddress   string  `orm:"refundaddress" json:"refundaddress"`
	Content         string  `orm:"content" json:"content"`
	Reason          string  `orm:"reason" json:"reason"`
	Images          string  `orm:"images" json:"images"`
	Applytime       string  `orm:"applytime" json:"applytime"`
	Applycredit     int     `orm:"applycredit" json:"applycredit"`
	Applyprice      float64 `orm:"applyprice" json:"applyprice"`
	Reply           string  `orm:"reply" json:"reply"`
	Refundtype      string  `orm:"refundtype" json:"refundtype"`
	Rtype           int     `orm:"rtype" json:"rtype"`
	Refundtime      string  `orm:"refundtime" json:"refundtime"`
	Endtime         string  `orm:"endtime" json:"endtime"`
	Message         string  `orm:"message" json:"message"`
	Operatetime     string  `orm:"operatetime" json:"operatetime"`
	Realcredit      int     `orm:"realcredit" json:"realcredit"`
	Realmoney       float64 `orm:"realmoney" json:"realmoney"`
	Express         string  `orm:"express" json:"express"`
	Expresscom      string  `orm:"expresscom" json:"expresscom"`
	Expresssn       string  `orm:"expresssn" json:"expresssn"`
	Sendtime        string  `orm:"sendtime" json:"sendtime"`
	Returntime      int     `orm:"returntime" json:"returntime"`
	Rexpress        string  `orm:"rexpress" json:"rexpress"`
	Rexpresscom     string  `orm:"rexpresscom" json:"rexpresscom"`
	Rexpresssn      string  `orm:"rexpresssn" json:"rexpresssn"`
}

func (*Ewei_shop_groups_order_refund) TableName() string {
	return "ewei_shop_groups_order_refund"
}

type Ewei_shop_payment struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Title        string `orm:"title" json:"title"`
	Type         int    `orm:"type" json:"type"`
	Appid        string `orm:"appid" json:"appid"`
	MchId        string `orm:"mch_id" json:"mch_id"`
	Apikey       string `orm:"apikey" json:"apikey"`
	SubAppid     string `orm:"sub_appid" json:"sub_appid"`
	SubAppsecret string `orm:"sub_appsecret" json:"sub_appsecret"`
	SubMchId     string `orm:"sub_mch_id" json:"sub_mch_id"`
	CertFile     string `orm:"cert_file" json:"cert_file"`
	KeyFile      string `orm:"key_file" json:"key_file"`
	RootFile     string `orm:"root_file" json:"root_file"`
	IsRaw        int    `orm:"is_raw" json:"is_raw"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Paytype      int    `orm:"paytype" json:"paytype"`
	Alitype      int    `orm:"alitype" json:"alitype"`
	AlipaySec    string `orm:"alipay_sec" json:"alipay_sec"`
}

func (*Ewei_shop_payment) TableName() string {
	return "ewei_shop_payment"
}

type Ewei_shop_seckill_task_time struct {
	Id      int `orm:"id" json:"id"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Taskid  int `orm:"taskid" json:"taskid"`
	Time    int `orm:"time" json:"time"`
}

func (*Ewei_shop_seckill_task_time) TableName() string {
	return "ewei_shop_seckill_task_time"
}

type Ewei_shop_sign_user struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Openid        string `orm:"openid" json:"openid"`
	Order         int    `orm:"order" json:"order"`
	Orderday      int    `orm:"orderday" json:"orderday"`
	Sum           int    `orm:"sum" json:"sum"`
	Signdate      string `orm:"signdate" json:"signdate"`
	Isminiprogram int    `orm:"isminiprogram" json:"isminiprogram"`
}

func (*Ewei_shop_sign_user) TableName() string {
	return "ewei_shop_sign_user"
}

type Ewei_shop_task_poster_qr struct {
	Id         int    `orm:"id" json:"id"`
	Acid       int    `orm:"acid" json:"acid"`
	Openid     string `orm:"openid" json:"openid"`
	Posterid   int    `orm:"posterid" json:"posterid"`
	Type       int    `orm:"type" json:"type"`
	Sceneid    int    `orm:"sceneid" json:"sceneid"`
	Mediaid    string `orm:"mediaid" json:"mediaid"`
	Ticket     string `orm:"ticket" json:"ticket"`
	Url        string `orm:"url" json:"url"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Qrimg      string `orm:"qrimg" json:"qrimg"`
	Expire     int    `orm:"expire" json:"expire"`
	Endtime    int    `orm:"endtime" json:"endtime"`
}

func (*Ewei_shop_task_poster_qr) TableName() string {
	return "ewei_shop_task_poster_qr"
}

type Paycenter_order struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Uid           int     `orm:"uid" json:"uid"`
	Pid           int     `orm:"pid" json:"pid"`
	ClerkId       int     `orm:"clerk_id" json:"clerk_id"`
	StoreId       int     `orm:"store_id" json:"store_id"`
	ClerkType     int     `orm:"clerk_type" json:"clerk_type"`
	Uniontid      string  `orm:"uniontid" json:"uniontid"`
	TransactionId string  `orm:"transaction_id" json:"transaction_id"`
	Type          string  `orm:"type" json:"type"`
	TradeType     string  `orm:"trade_type" json:"trade_type"`
	Body          string  `orm:"body" json:"body"`
	Fee           string  `orm:"fee" json:"fee"`
	FinalFee      float64 `orm:"final_fee" json:"final_fee"`
	Credit1       int     `orm:"credit1" json:"credit1"`
	Credit1Fee    float64 `orm:"credit1_fee" json:"credit1_fee"`
	Credit2       float64 `orm:"credit2" json:"credit2"`
	Cash          float64 `orm:"cash" json:"cash"`
	Remark        string  `orm:"remark" json:"remark"`
	AuthCode      string  `orm:"auth_code" json:"auth_code"`
	Openid        string  `orm:"openid" json:"openid"`
	Nickname      string  `orm:"nickname" json:"nickname"`
	Follow        int     `orm:"follow" json:"follow"`
	Status        int     `orm:"status" json:"status"`
	CreditStatus  int     `orm:"credit_status" json:"credit_status"`
	Paytime       int     `orm:"paytime" json:"paytime"`
	Createtime    int     `orm:"createtime" json:"createtime"`
}

func (*Paycenter_order) TableName() string {
	return "paycenter_order"
}

type Article_category struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Type         string `orm:"type" json:"type"`
}

func (*Article_category) TableName() string {
	return "article_category"
}

type Core_queue struct {
	Qid      int    `orm:"qid" json:"qid"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Acid     int    `orm:"acid" json:"acid"`
	Message  string `orm:"message" json:"message"`
	Params   string `orm:"params" json:"params"`
	Keyword  string `orm:"keyword" json:"keyword"`
	Response string `orm:"response" json:"response"`
	Module   string `orm:"module" json:"module"`
	Type     int    `orm:"type" json:"type"`
	Dateline int    `orm:"dateline" json:"dateline"`
}

func (*Core_queue) TableName() string {
	return "core_queue"
}

type Ewei_shop_cashier_goods struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Cashierid  int     `orm:"cashierid" json:"cashierid"`
	Createtime int     `orm:"createtime" json:"createtime"`
	Title      string  `orm:"title" json:"title"`
	Image      string  `orm:"image" json:"image"`
	Categoryid int     `orm:"categoryid" json:"categoryid"`
	Price      float64 `orm:"price" json:"price"`
	Total      int     `orm:"total" json:"total"`
	Status     int     `orm:"status" json:"status"`
	Goodssn    string  `orm:"goodssn" json:"goodssn"`
}

func (*Ewei_shop_cashier_goods) TableName() string {
	return "ewei_shop_cashier_goods"
}

type Ewei_shop_coupon_sendtasks struct {
	Id        int     `orm:"id" json:"id"`
	Uniacid   int     `orm:"uniacid" json:"uniacid"`
	Enough    float64 `orm:"enough" json:"enough"`
	Couponid  int     `orm:"couponid" json:"couponid"`
	Starttime int     `orm:"starttime" json:"starttime"`
	Endtime   int     `orm:"endtime" json:"endtime"`
	Sendnum   int     `orm:"sendnum" json:"sendnum"`
	Num       int     `orm:"num" json:"num"`
	Sendpoint int     `orm:"sendpoint" json:"sendpoint"`
	Status    int     `orm:"status" json:"status"`
}

func (*Ewei_shop_coupon_sendtasks) TableName() string {
	return "ewei_shop_coupon_sendtasks"
}

type Ewei_shop_newstore_category struct {
	Id      int    `orm:"id" json:"id"`
	Name    string `orm:"name" json:"name"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
}

func (*Ewei_shop_newstore_category) TableName() string {
	return "ewei_shop_newstore_category"
}

type Activity_exchange struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Title       string `orm:"title" json:"title"`
	Description string `orm:"description" json:"description"`
	Thumb       string `orm:"thumb" json:"thumb"`
	Type        int    `orm:"type" json:"type"`
	Extra       string `orm:"extra" json:"extra"`
	Credit      int    `orm:"credit" json:"credit"`
	Credittype  string `orm:"credittype" json:"credittype"`
	Pretotal    int    `orm:"pretotal" json:"pretotal"`
	Num         int    `orm:"num" json:"num"`
	Total       int    `orm:"total" json:"total"`
	Status      int    `orm:"status" json:"status"`
	Starttime   int    `orm:"starttime" json:"starttime"`
	Endtime     int    `orm:"endtime" json:"endtime"`
}

func (*Activity_exchange) TableName() string {
	return "activity_exchange"
}

type Core_attachment struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Uid             int    `orm:"uid" json:"uid"`
	Filename        string `orm:"filename" json:"filename"`
	Attachment      string `orm:"attachment" json:"attachment"`
	Type            int    `orm:"type" json:"type"`
	Createtime      int    `orm:"createtime" json:"createtime"`
	ModuleUploadDir string `orm:"module_upload_dir" json:"module_upload_dir"`
}

func (*Core_attachment) TableName() string {
	return "core_attachment"
}

type Ewei_shop_merch_saler struct {
	Id        int    `orm:"id" json:"id"`
	Storeid   int    `orm:"storeid" json:"storeid"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Openid    string `orm:"openid" json:"openid"`
	Status    int    `orm:"status" json:"status"`
	Salername string `orm:"salername" json:"salername"`
	Merchid   int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_saler) TableName() string {
	return "ewei_shop_merch_saler"
}

type Ewei_shop_merch_store struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Storename       string  `orm:"storename" json:"storename"`
	Address         string  `orm:"address" json:"address"`
	Tel             string  `orm:"tel" json:"tel"`
	Lat             string  `orm:"lat" json:"lat"`
	Lng             string  `orm:"lng" json:"lng"`
	Status          int     `orm:"status" json:"status"`
	Type            int     `orm:"type" json:"type"`
	Realname        string  `orm:"realname" json:"realname"`
	Mobile          string  `orm:"mobile" json:"mobile"`
	Fetchtime       string  `orm:"fetchtime" json:"fetchtime"`
	Logo            string  `orm:"logo" json:"logo"`
	Saletime        string  `orm:"saletime" json:"saletime"`
	Desc            string  `orm:"desc" json:"desc"`
	Displayorder    int     `orm:"displayorder" json:"displayorder"`
	CommissionTotal float64 `orm:"commission_total" json:"commission_total"`
	Merchid         int     `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_store) TableName() string {
	return "ewei_shop_merch_store"
}

type Ewei_shop_order struct {
	Id                   int     `orm:"id" json:"id"`
	Uniacid              int     `orm:"uniacid" json:"uniacid"`
	Openid               string  `orm:"openid" json:"openid"`
	Agentid              int     `orm:"agentid" json:"agentid"`
	Ordersn              string  `orm:"ordersn" json:"ordersn"`
	Price                float64 `orm:"price" json:"price"`
	Goodsprice           float64 `orm:"goodsprice" json:"goodsprice"`
	Discountprice        float64 `orm:"discountprice" json:"discountprice"`
	Status               int     `orm:"status" json:"status"`
	Paytype              int     `orm:"paytype" json:"paytype"`
	Transid              string  `orm:"transid" json:"transid"`
	Remark               string  `orm:"remark" json:"remark"`
	Addressid            int     `orm:"addressid" json:"addressid"`
	Dispatchprice        float64 `orm:"dispatchprice" json:"dispatchprice"`
	Dispatchid           int     `orm:"dispatchid" json:"dispatchid"`
	Createtime           int     `orm:"createtime" json:"createtime"`
	Dispatchtype         int     `orm:"dispatchtype" json:"dispatchtype"`
	Carrier              string  `orm:"carrier" json:"carrier"`
	Refundid             int     `orm:"refundid" json:"refundid"`
	Iscomment            int     `orm:"iscomment" json:"iscomment"`
	Creditadd            int     `orm:"creditadd" json:"creditadd"`
	Deleted              int     `orm:"deleted" json:"deleted"`
	Userdeleted          int     `orm:"userdeleted" json:"userdeleted"`
	Finishtime           int     `orm:"finishtime" json:"finishtime"`
	Paytime              int     `orm:"paytime" json:"paytime"`
	Expresscom           string  `orm:"expresscom" json:"expresscom"`
	Expresssn            string  `orm:"expresssn" json:"expresssn"`
	Express              string  `orm:"express" json:"express"`
	Sendtime             int     `orm:"sendtime" json:"sendtime"`
	Fetchtime            int     `orm:"fetchtime" json:"fetchtime"`
	Cash                 int     `orm:"cash" json:"cash"`
	Canceltime           int     `orm:"canceltime" json:"canceltime"`
	Cancelpaytime        int     `orm:"cancelpaytime" json:"cancelpaytime"`
	Refundtime           int     `orm:"refundtime" json:"refundtime"`
	Isverify             int     `orm:"isverify" json:"isverify"`
	Verified             int     `orm:"verified" json:"verified"`
	Verifyopenid         string  `orm:"verifyopenid" json:"verifyopenid"`
	Verifycode           string  `orm:"verifycode" json:"verifycode"`
	Verifytime           int     `orm:"verifytime" json:"verifytime"`
	Verifystoreid        int     `orm:"verifystoreid" json:"verifystoreid"`
	Deductprice          float64 `orm:"deductprice" json:"deductprice"`
	Deductcredit         int     `orm:"deductcredit" json:"deductcredit"`
	Deductcredit2        float64 `orm:"deductcredit2" json:"deductcredit2"`
	Deductenough         float64 `orm:"deductenough" json:"deductenough"`
	Virtual              int     `orm:"virtual" json:"virtual"`
	VirtualInfo          string  `orm:"virtual_info" json:"virtual_info"`
	VirtualStr           string  `orm:"virtual_str" json:"virtual_str"`
	Address              string  `orm:"address" json:"address"`
	Sysdeleted           int     `orm:"sysdeleted" json:"sysdeleted"`
	Ordersn2             int     `orm:"ordersn2" json:"ordersn2"`
	Changeprice          float64 `orm:"changeprice" json:"changeprice"`
	Changedispatchprice  float64 `orm:"changedispatchprice" json:"changedispatchprice"`
	Oldprice             float64 `orm:"oldprice" json:"oldprice"`
	Olddispatchprice     float64 `orm:"olddispatchprice" json:"olddispatchprice"`
	Isvirtual            int     `orm:"isvirtual" json:"isvirtual"`
	Couponid             int     `orm:"couponid" json:"couponid"`
	Couponprice          float64 `orm:"couponprice" json:"couponprice"`
	Diyformdata          string  `orm:"diyformdata" json:"diyformdata"`
	Diyformfields        string  `orm:"diyformfields" json:"diyformfields"`
	Diyformid            int     `orm:"diyformid" json:"diyformid"`
	Storeid              int     `orm:"storeid" json:"storeid"`
	Printstate           int     `orm:"printstate" json:"printstate"`
	Printstate2          int     `orm:"printstate2" json:"printstate2"`
	AddressSend          string  `orm:"address_send" json:"address_send"`
	Refundstate          int     `orm:"refundstate" json:"refundstate"`
	Closereason          string  `orm:"closereason" json:"closereason"`
	Remarksaler          string  `orm:"remarksaler" json:"remarksaler"`
	Remarkclose          string  `orm:"remarkclose" json:"remarkclose"`
	Remarksend           string  `orm:"remarksend" json:"remarksend"`
	Ismr                 int     `orm:"ismr" json:"ismr"`
	Isdiscountprice      float64 `orm:"isdiscountprice" json:"isdiscountprice"`
	Isvirtualsend        int     `orm:"isvirtualsend" json:"isvirtualsend"`
	VirtualsendInfo      string  `orm:"virtualsend_info" json:"virtualsend_info"`
	Verifyinfo           string  `orm:"verifyinfo" json:"verifyinfo"`
	Verifytype           int     `orm:"verifytype" json:"verifytype"`
	Verifycodes          string  `orm:"verifycodes" json:"verifycodes"`
	Invoicename          string  `orm:"invoicename" json:"invoicename"`
	Merchid              int     `orm:"merchid" json:"merchid"`
	Ismerch              int     `orm:"ismerch" json:"ismerch"`
	Parentid             int     `orm:"parentid" json:"parentid"`
	Isparent             int     `orm:"isparent" json:"isparent"`
	Grprice              float64 `orm:"grprice" json:"grprice"`
	Merchshow            int     `orm:"merchshow" json:"merchshow"`
	Merchdeductenough    float64 `orm:"merchdeductenough" json:"merchdeductenough"`
	Couponmerchid        int     `orm:"couponmerchid" json:"couponmerchid"`
	Isglobonus           int     `orm:"isglobonus" json:"isglobonus"`
	Merchapply           int     `orm:"merchapply" json:"merchapply"`
	Isabonus             int     `orm:"isabonus" json:"isabonus"`
	Isborrow             int     `orm:"isborrow" json:"isborrow"`
	Borrowopenid         string  `orm:"borrowopenid" json:"borrowopenid"`
	Merchisdiscountprice float64 `orm:"merchisdiscountprice" json:"merchisdiscountprice"`
	Apppay               int     `orm:"apppay" json:"apppay"`
	Coupongoodprice      float64 `orm:"coupongoodprice" json:"coupongoodprice"`
	Buyagainprice        float64 `orm:"buyagainprice" json:"buyagainprice"`
	Authorid             int     `orm:"authorid" json:"authorid"`
	Isauthor             int     `orm:"isauthor" json:"isauthor"`
	Ispackage            int     `orm:"ispackage" json:"ispackage"`
	Packageid            int     `orm:"packageid" json:"packageid"`
	Taskdiscountprice    float64 `orm:"taskdiscountprice" json:"taskdiscountprice"`
	Seckilldiscountprice float64 `orm:"seckilldiscountprice" json:"seckilldiscountprice"`
	Verifyendtime        int     `orm:"verifyendtime" json:"verifyendtime"`
	Willcancelmessage    int     `orm:"willcancelmessage" json:"willcancelmessage"`
	Sendtype             int     `orm:"sendtype" json:"sendtype"`
	Lotterydiscountprice float64 `orm:"lotterydiscountprice" json:"lotterydiscountprice"`
	Contype              int     `orm:"contype" json:"contype"`
	Wxid                 int     `orm:"wxid" json:"wxid"`
	Wxcardid             string  `orm:"wxcardid" json:"wxcardid"`
	Wxcode               string  `orm:"wxcode" json:"wxcode"`
	Dispatchkey          string  `orm:"dispatchkey" json:"dispatchkey"`
	Quickid              int     `orm:"quickid" json:"quickid"`
	Istrade              int     `orm:"istrade" json:"istrade"`
	Isnewstore           int     `orm:"isnewstore" json:"isnewstore"`
	Liveid               int     `orm:"liveid" json:"liveid"`
	OrdersnTrade         string  `orm:"ordersn_trade" json:"ordersn_trade"`
	Tradestatus          int     `orm:"tradestatus" json:"tradestatus"`
	Tradepaytype         int     `orm:"tradepaytype" json:"tradepaytype"`
	Tradepaytime         int     `orm:"tradepaytime" json:"tradepaytime"`
	Dowpayment           float64 `orm:"dowpayment" json:"dowpayment"`
	Betweenprice         float64 `orm:"betweenprice" json:"betweenprice"`
	Isshare              int     `orm:"isshare" json:"isshare"`
	Officcode            string  `orm:"officcode" json:"officcode"`
	WxappPrepayId        string  `orm:"wxapp_prepay_id" json:"wxapp_prepay_id"`
	Cashtime             int     `orm:"cashtime" json:"cashtime"`
	Iswxappcreate        int     `orm:"iswxappcreate" json:"iswxappcreate"`
	RandomCode           string  `orm:"random_code" json:"random_code"`
	PrintTemplate        string  `orm:"print_template" json:"print_template"`
	CityExpressState     int     `orm:"city_express_state" json:"city_express_state"`
}

func (*Ewei_shop_order) TableName() string {
	return "ewei_shop_order"
}

type Ewei_shop_sns_board_follow struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Bid        int    `orm:"bid" json:"bid"`
	Openid     string `orm:"openid" json:"openid"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_sns_board_follow) TableName() string {
	return "ewei_shop_sns_board_follow"
}

type Ewei_shop_system_copyright_notice struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Author       string `orm:"author" json:"author"`
	Content      string `orm:"content" json:"content"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_copyright_notice) TableName() string {
	return "ewei_shop_system_copyright_notice"
}

type Ewei_shop_system_guestbook struct {
	Id         int    `orm:"id" json:"id"`
	Title      string `orm:"title" json:"title"`
	Content    string `orm:"content" json:"content"`
	Nickname   string `orm:"nickname" json:"nickname"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Email      string `orm:"email" json:"email"`
	Clientip   string `orm:"clientip" json:"clientip"`
	Mobile     string `orm:"mobile" json:"mobile"`
}

func (*Ewei_shop_system_guestbook) TableName() string {
	return "ewei_shop_system_guestbook"
}

type Ewei_shop_task_reward struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Taskid        int     `orm:"taskid" json:"taskid"`
	Tasktitle     string  `orm:"tasktitle" json:"tasktitle"`
	Tasktype      string  `orm:"tasktype" json:"tasktype"`
	Taskowner     string  `orm:"taskowner" json:"taskowner"`
	Ownernickname string  `orm:"ownernickname" json:"ownernickname"`
	Recordid      int     `orm:"recordid" json:"recordid"`
	Nickname      string  `orm:"nickname" json:"nickname"`
	Headimg       string  `orm:"headimg" json:"headimg"`
	Openid        string  `orm:"openid" json:"openid"`
	RewardType    string  `orm:"reward_type" json:"reward_type"`
	RewardTitle   string  `orm:"reward_title" json:"reward_title"`
	RewardData    float64 `orm:"reward_data" json:"reward_data"`
	Get           int     `orm:"get" json:"get"`
	Sent          int     `orm:"sent" json:"sent"`
	Gettime       string  `orm:"gettime" json:"gettime"`
	Senttime      string  `orm:"senttime" json:"senttime"`
	Isjoiner      int     `orm:"isjoiner" json:"isjoiner"`
	Price         float64 `orm:"price" json:"price"`
	Level         int     `orm:"level" json:"level"`
	Read          int     `orm:"read" json:"read"`
}

func (*Ewei_shop_task_reward) TableName() string {
	return "ewei_shop_task_reward"
}

type Coupon_record struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Acid         int    `orm:"acid" json:"acid"`
	CardId       string `orm:"card_id" json:"card_id"`
	Openid       string `orm:"openid" json:"openid"`
	FriendOpenid string `orm:"friend_openid" json:"friend_openid"`
	Givebyfriend int    `orm:"givebyfriend" json:"givebyfriend"`
	Code         string `orm:"code" json:"code"`
	Hash         string `orm:"hash" json:"hash"`
	Addtime      int    `orm:"addtime" json:"addtime"`
	Usetime      int    `orm:"usetime" json:"usetime"`
	Status       int    `orm:"status" json:"status"`
	ClerkName    string `orm:"clerk_name" json:"clerk_name"`
	ClerkId      int    `orm:"clerk_id" json:"clerk_id"`
	StoreId      int    `orm:"store_id" json:"store_id"`
	ClerkType    int    `orm:"clerk_type" json:"clerk_type"`
	Couponid     int    `orm:"couponid" json:"couponid"`
	Uid          int    `orm:"uid" json:"uid"`
	Grantmodule  string `orm:"grantmodule" json:"grantmodule"`
	Remark       string `orm:"remark" json:"remark"`
}

func (*Coupon_record) TableName() string {
	return "coupon_record"
}

type Ewei_shop_cashier_pay_log struct {
	Id             int     `orm:"id" json:"id"`
	Uniacid        int     `orm:"uniacid" json:"uniacid"`
	Cashierid      int     `orm:"cashierid" json:"cashierid"`
	Operatorid     int     `orm:"operatorid" json:"operatorid"`
	Openid         string  `orm:"openid" json:"openid"`
	Paytype        int     `orm:"paytype" json:"paytype"`
	Logno          string  `orm:"logno" json:"logno"`
	Title          string  `orm:"title" json:"title"`
	Createtime     int     `orm:"createtime" json:"createtime"`
	Status         int     `orm:"status" json:"status"`
	Money          float64 `orm:"money" json:"money"`
	Paytime        int     `orm:"paytime" json:"paytime"`
	IsApplypay     int     `orm:"is_applypay" json:"is_applypay"`
	Randommoney    float64 `orm:"randommoney" json:"randommoney"`
	Enough         float64 `orm:"enough" json:"enough"`
	Mobile         string  `orm:"mobile" json:"mobile"`
	Deduction      float64 `orm:"deduction" json:"deduction"`
	Discountmoney  float64 `orm:"discountmoney" json:"discountmoney"`
	Discount       float64 `orm:"discount" json:"discount"`
	Isgoods        int     `orm:"isgoods" json:"isgoods"`
	Orderid        int     `orm:"orderid" json:"orderid"`
	Orderprice     float64 `orm:"orderprice" json:"orderprice"`
	Goodsprice     float64 `orm:"goodsprice" json:"goodsprice"`
	Couponpay      float64 `orm:"couponpay" json:"couponpay"`
	Payopenid      string  `orm:"payopenid" json:"payopenid"`
	Nosalemoney    float64 `orm:"nosalemoney" json:"nosalemoney"`
	Coupon         int     `orm:"coupon" json:"coupon"`
	Usecoupon      int     `orm:"usecoupon" json:"usecoupon"`
	Usecouponprice float64 `orm:"usecouponprice" json:"usecouponprice"`
	PresentCredit1 int     `orm:"present_credit1" json:"present_credit1"`
	Refundsn       string  `orm:"refundsn" json:"refundsn"`
	Refunduser     int     `orm:"refunduser" json:"refunduser"`
}

func (*Ewei_shop_cashier_pay_log) TableName() string {
	return "ewei_shop_cashier_pay_log"
}

type Ewei_shop_customer_guestbook struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Realname   string `orm:"realname" json:"realname"`
	Mobile     string `orm:"mobile" json:"mobile"`
	Weixin     string `orm:"weixin" json:"weixin"`
	Images     string `orm:"images" json:"images"`
	Content    string `orm:"content" json:"content"`
	Remark     string `orm:"remark" json:"remark"`
	Status     int    `orm:"status" json:"status"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Deleted    int    `orm:"deleted" json:"deleted"`
}

func (*Ewei_shop_customer_guestbook) TableName() string {
	return "ewei_shop_customer_guestbook"
}

type Ewei_shop_diypage struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Type         int    `orm:"type" json:"type"`
	Name         string `orm:"name" json:"name"`
	Data         string `orm:"data" json:"data"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Lastedittime int    `orm:"lastedittime" json:"lastedittime"`
	Keyword      string `orm:"keyword" json:"keyword"`
	Diymenu      int    `orm:"diymenu" json:"diymenu"`
	Merch        int    `orm:"merch" json:"merch"`
	Diyadv       int    `orm:"diyadv" json:"diyadv"`
}

func (*Ewei_shop_diypage) TableName() string {
	return "ewei_shop_diypage"
}

type Ewei_shop_fullback_log struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Openid       string  `orm:"openid" json:"openid"`
	Orderid      int     `orm:"orderid" json:"orderid"`
	Price        float64 `orm:"price" json:"price"`
	Priceevery   float64 `orm:"priceevery" json:"priceevery"`
	Day          int     `orm:"day" json:"day"`
	Fullbackday  int     `orm:"fullbackday" json:"fullbackday"`
	Createtime   int     `orm:"createtime" json:"createtime"`
	Fullbacktime int     `orm:"fullbacktime" json:"fullbacktime"`
	Isfullback   int     `orm:"isfullback" json:"isfullback"`
	Goodsid      int     `orm:"goodsid" json:"goodsid"`
}

func (*Ewei_shop_fullback_log) TableName() string {
	return "ewei_shop_fullback_log"
}

type Ewei_shop_invitation struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Type       int    `orm:"type" json:"type"`
	Title      string `orm:"title" json:"title"`
	Data       string `orm:"data" json:"data"`
	Scan       int    `orm:"scan" json:"scan"`
	Follow     int    `orm:"follow" json:"follow"`
	Qrcode     int    `orm:"qrcode" json:"qrcode"`
	Status     int    `orm:"status" json:"status"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_invitation) TableName() string {
	return "ewei_shop_invitation"
}

type Ewei_shop_sns_manage struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Bid     int    `orm:"bid" json:"bid"`
	Openid  string `orm:"openid" json:"openid"`
	Enabled int    `orm:"enabled" json:"enabled"`
}

func (*Ewei_shop_sns_manage) TableName() string {
	return "ewei_shop_sns_manage"
}

type Mc_credits_record struct {
	Id         int     `orm:"id" json:"id"`
	Uid        int     `orm:"uid" json:"uid"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Credittype string  `orm:"credittype" json:"credittype"`
	Num        float64 `orm:"num" json:"num"`
	Operator   int     `orm:"operator" json:"operator"`
	Module     string  `orm:"module" json:"module"`
	ClerkId    int     `orm:"clerk_id" json:"clerk_id"`
	StoreId    int     `orm:"store_id" json:"store_id"`
	ClerkType  int     `orm:"clerk_type" json:"clerk_type"`
	Createtime int     `orm:"createtime" json:"createtime"`
	Remark     string  `orm:"remark" json:"remark"`
}

func (*Mc_credits_record) TableName() string {
	return "mc_credits_record"
}

type Profile_fields struct {
	Id             int    `orm:"id" json:"id"`
	Field          string `orm:"field" json:"field"`
	Available      int    `orm:"available" json:"available"`
	Title          string `orm:"title" json:"title"`
	Description    string `orm:"description" json:"description"`
	Displayorder   int    `orm:"displayorder" json:"displayorder"`
	Required       int    `orm:"required" json:"required"`
	Unchangeable   int    `orm:"unchangeable" json:"unchangeable"`
	Showinregister int    `orm:"showinregister" json:"showinregister"`
	FieldLength    int    `orm:"field_length" json:"field_length"`
}

func (*Profile_fields) TableName() string {
	return "profile_fields"
}

type Qrcode_stat struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Acid       int    `orm:"acid" json:"acid"`
	Qid        int    `orm:"qid" json:"qid"`
	Openid     string `orm:"openid" json:"openid"`
	Type       int    `orm:"type" json:"type"`
	Qrcid      int    `orm:"qrcid" json:"qrcid"`
	SceneStr   string `orm:"scene_str" json:"scene_str"`
	Name       string `orm:"name" json:"name"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Qrcode_stat) TableName() string {
	return "qrcode_stat"
}

type Users_founder_group struct {
	Id            int    `orm:"id" json:"id"`
	Name          string `orm:"name" json:"name"`
	Package       string `orm:"package" json:"package"`
	Maxaccount    int    `orm:"maxaccount" json:"maxaccount"`
	Maxsubaccount int    `orm:"maxsubaccount" json:"maxsubaccount"`
	Timelimit     int    `orm:"timelimit" json:"timelimit"`
	Maxwxapp      int    `orm:"maxwxapp" json:"maxwxapp"`
	Maxwebapp     int    `orm:"maxwebapp" json:"maxwebapp"`
}

func (*Users_founder_group) TableName() string {
	return "users_founder_group"
}

type Ewei_message_mass_sign struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Openid   string `orm:"openid" json:"openid"`
	Nickname string `orm:"nickname" json:"nickname"`
	Taskid   int    `orm:"taskid" json:"taskid"`
	Status   int    `orm:"status" json:"status"`
	Log      string `orm:"log" json:"log"`
}

func (*Ewei_message_mass_sign) TableName() string {
	return "ewei_message_mass_sign"
}

type Ewei_shop_abonus_bill struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Billno          string  `orm:"billno" json:"billno"`
	Paytype         int     `orm:"paytype" json:"paytype"`
	Year            int     `orm:"year" json:"year"`
	Month           int     `orm:"month" json:"month"`
	Week            int     `orm:"week" json:"week"`
	Ordercount      int     `orm:"ordercount" json:"ordercount"`
	Ordermoney      float64 `orm:"ordermoney" json:"ordermoney"`
	Paytime         int     `orm:"paytime" json:"paytime"`
	Aagentcount1    int     `orm:"aagentcount1" json:"aagentcount1"`
	Aagentcount2    int     `orm:"aagentcount2" json:"aagentcount2"`
	Aagentcount3    int     `orm:"aagentcount3" json:"aagentcount3"`
	Bonusmoney1     float64 `orm:"bonusmoney1" json:"bonusmoney1"`
	BonusmoneySend1 float64 `orm:"bonusmoney_send1" json:"bonusmoney_send1"`
	BonusmoneyPay1  float64 `orm:"bonusmoney_pay1" json:"bonusmoney_pay1"`
	Bonusmoney2     float64 `orm:"bonusmoney2" json:"bonusmoney2"`
	BonusmoneySend2 float64 `orm:"bonusmoney_send2" json:"bonusmoney_send2"`
	BonusmoneyPay2  float64 `orm:"bonusmoney_pay2" json:"bonusmoney_pay2"`
	Bonusmoney3     float64 `orm:"bonusmoney3" json:"bonusmoney3"`
	BonusmoneySend3 float64 `orm:"bonusmoney_send3" json:"bonusmoney_send3"`
	BonusmoneyPay3  float64 `orm:"bonusmoney_pay3" json:"bonusmoney_pay3"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Status          int     `orm:"status" json:"status"`
	Starttime       int     `orm:"starttime" json:"starttime"`
	Endtime         int     `orm:"endtime" json:"endtime"`
	Confirmtime     int     `orm:"confirmtime" json:"confirmtime"`
	Ceshi           int     `orm:"ceshi" json:"ceshi"`
}

func (*Ewei_shop_abonus_bill) TableName() string {
	return "ewei_shop_abonus_bill"
}

type Ewei_shop_author_team struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Teamno     string `orm:"teamno" json:"teamno"`
	Year       int    `orm:"year" json:"year"`
	Month      int    `orm:"month" json:"month"`
	TeamCount  int    `orm:"team_count" json:"team_count"`
	TeamIds    string `orm:"team_ids" json:"team_ids"`
	Status     int    `orm:"status" json:"status"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Paytime    int    `orm:"paytime" json:"paytime"`
}

func (*Ewei_shop_author_team) TableName() string {
	return "ewei_shop_author_team"
}

type Ewei_shop_commission_clickcount struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	FromOpenid string `orm:"from_openid" json:"from_openid"`
	Clicktime  int    `orm:"clicktime" json:"clicktime"`
}

func (*Ewei_shop_commission_clickcount) TableName() string {
	return "ewei_shop_commission_clickcount"
}

type Ewei_shop_pc_slide struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Type         int    `orm:"type" json:"type"`
	Advname      string `orm:"advname" json:"advname"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Backcolor    string `orm:"backcolor" json:"backcolor"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Shopid       int    `orm:"shopid" json:"shopid"`
}

func (*Ewei_shop_pc_slide) TableName() string {
	return "ewei_shop_pc_slide"
}

type Rule struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Module       string `orm:"module" json:"module"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
	Containtype  string `orm:"containtype" json:"containtype"`
	ReplyType    int    `orm:"reply_type" json:"reply_type"`
}

func (*Rule) TableName() string {
	return "rule"
}

type Users_bind struct {
	Id            int    `orm:"id" json:"id"`
	Uid           int    `orm:"uid" json:"uid"`
	BindSign      string `orm:"bind_sign" json:"bind_sign"`
	ThirdType     int    `orm:"third_type" json:"third_type"`
	ThirdNickname string `orm:"third_nickname" json:"third_nickname"`
}

func (*Users_bind) TableName() string {
	return "users_bind"
}

type Wechat_attachment struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Acid            int    `orm:"acid" json:"acid"`
	Uid             int    `orm:"uid" json:"uid"`
	Filename        string `orm:"filename" json:"filename"`
	Attachment      string `orm:"attachment" json:"attachment"`
	MediaId         string `orm:"media_id" json:"media_id"`
	Width           int    `orm:"width" json:"width"`
	Height          int    `orm:"height" json:"height"`
	Type            string `orm:"type" json:"type"`
	Model           string `orm:"model" json:"model"`
	Tag             string `orm:"tag" json:"tag"`
	Createtime      int    `orm:"createtime" json:"createtime"`
	ModuleUploadDir string `orm:"module_upload_dir" json:"module_upload_dir"`
}

func (*Wechat_attachment) TableName() string {
	return "wechat_attachment"
}

type Coupon_store struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Couponid string `orm:"couponid" json:"couponid"`
	Storeid  int    `orm:"storeid" json:"storeid"`
}

func (*Coupon_store) TableName() string {
	return "coupon_store"
}

type Ewei_shop_cashier_operator struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Cashierid    int    `orm:"cashierid" json:"cashierid"`
	Title        string `orm:"title" json:"title"`
	Manageopenid string `orm:"manageopenid" json:"manageopenid"`
	Username     string `orm:"username" json:"username"`
	Password     string `orm:"password" json:"password"`
	Salt         string `orm:"salt" json:"salt"`
	Perm         string `orm:"perm" json:"perm"`
	Createtime   int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_cashier_operator) TableName() string {
	return "ewei_shop_cashier_operator"
}

type Ewei_shop_exchange_cart struct {
	Id          int     `orm:"id" json:"id"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Openid      string  `orm:"openid" json:"openid"`
	Goodsid     int     `orm:"goodsid" json:"goodsid"`
	Total       int     `orm:"total" json:"total"`
	Marketprice float64 `orm:"marketprice" json:"marketprice"`
	Optionid    int     `orm:"optionid" json:"optionid"`
	Selected    int     `orm:"selected" json:"selected"`
	Deleted     int     `orm:"deleted" json:"deleted"`
	Merchid     int     `orm:"merchid" json:"merchid"`
	Title       string  `orm:"title" json:"title"`
	Groupid     int     `orm:"groupid" json:"groupid"`
	Serial      string  `orm:"serial" json:"serial"`
}

func (*Ewei_shop_exchange_cart) TableName() string {
	return "ewei_shop_exchange_cart"
}

type Ewei_shop_live_setting struct {
	Id             int    `orm:"id" json:"id"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Ismember       int    `orm:"ismember" json:"ismember"`
	ShareTitle     string `orm:"share_title" json:"share_title"`
	ShareIcon      string `orm:"share_icon" json:"share_icon"`
	ShareDesc      string `orm:"share_desc" json:"share_desc"`
	ShareUrl       string `orm:"share_url" json:"share_url"`
	Livenoticetime int    `orm:"livenoticetime" json:"livenoticetime"`
}

func (*Ewei_shop_live_setting) TableName() string {
	return "ewei_shop_live_setting"
}

type Ewei_shop_merch_nav struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Navname      string `orm:"navname" json:"navname"`
	Icon         string `orm:"icon" json:"icon"`
	Url          string `orm:"url" json:"url"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_nav) TableName() string {
	return "ewei_shop_merch_nav"
}

type Ewei_shop_perm_role struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Rolename string `orm:"rolename" json:"rolename"`
	Status   int    `orm:"status" json:"status"`
	Perms    string `orm:"perms" json:"perms"`
	Deleted  int    `orm:"deleted" json:"deleted"`
	Perms2   string `orm:"perms2" json:"perms2"`
}

func (*Ewei_shop_perm_role) TableName() string {
	return "ewei_shop_perm_role"
}

type Ewei_shop_postera_log struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Openid       string  `orm:"openid" json:"openid"`
	Posterid     int     `orm:"posterid" json:"posterid"`
	FromOpenid   string  `orm:"from_openid" json:"from_openid"`
	Subcredit    int     `orm:"subcredit" json:"subcredit"`
	Submoney     float64 `orm:"submoney" json:"submoney"`
	Reccredit    int     `orm:"reccredit" json:"reccredit"`
	Recmoney     float64 `orm:"recmoney" json:"recmoney"`
	Createtime   int     `orm:"createtime" json:"createtime"`
	Reccouponid  int     `orm:"reccouponid" json:"reccouponid"`
	Reccouponnum int     `orm:"reccouponnum" json:"reccouponnum"`
	Subcouponid  int     `orm:"subcouponid" json:"subcouponid"`
	Subcouponnum int     `orm:"subcouponnum" json:"subcouponnum"`
}

func (*Ewei_shop_postera_log) TableName() string {
	return "ewei_shop_postera_log"
}

type Ewei_shop_qa_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
}

func (*Ewei_shop_qa_category) TableName() string {
	return "ewei_shop_qa_category"
}

type Ewei_shop_system_banner struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Url          string `orm:"url" json:"url"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
	Background   string `orm:"background" json:"background"`
}

func (*Ewei_shop_system_banner) TableName() string {
	return "ewei_shop_system_banner"
}

type Site_article struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Rid          int    `orm:"rid" json:"rid"`
	Kid          int    `orm:"kid" json:"kid"`
	Iscommend    int    `orm:"iscommend" json:"iscommend"`
	Ishot        int    `orm:"ishot" json:"ishot"`
	Pcate        int    `orm:"pcate" json:"pcate"`
	Ccate        int    `orm:"ccate" json:"ccate"`
	Template     string `orm:"template" json:"template"`
	Title        string `orm:"title" json:"title"`
	Description  string `orm:"description" json:"description"`
	Content      string `orm:"content" json:"content"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Incontent    int    `orm:"incontent" json:"incontent"`
	Source       string `orm:"source" json:"source"`
	Author       string `orm:"author" json:"author"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Linkurl      string `orm:"linkurl" json:"linkurl"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Edittime     int    `orm:"edittime" json:"edittime"`
	Click        int    `orm:"click" json:"click"`
	Type         string `orm:"type" json:"type"`
	Credit       string `orm:"credit" json:"credit"`
}

func (*Site_article) TableName() string {
	return "site_article"
}

type Users_invitation struct {
	Id         int    `orm:"id" json:"id"`
	Code       string `orm:"code" json:"code"`
	Fromuid    int    `orm:"fromuid" json:"fromuid"`
	Inviteuid  int    `orm:"inviteuid" json:"inviteuid"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Users_invitation) TableName() string {
	return "users_invitation"
}

type Ewei_shop_live_favorite struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Roomid     int    `orm:"roomid" json:"roomid"`
	Openid     string `orm:"openid" json:"openid"`
	Deleted    int    `orm:"deleted" json:"deleted"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_live_favorite) TableName() string {
	return "ewei_shop_live_favorite"
}

type Ewei_shop_pc_adv struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Advname string `orm:"advname" json:"advname"`
	Title   string `orm:"title" json:"title"`
	Src     string `orm:"src" json:"src"`
	Alt     string `orm:"alt" json:"alt"`
	Enabled int    `orm:"enabled" json:"enabled"`
	Link    string `orm:"link" json:"link"`
	Width   int    `orm:"width" json:"width"`
	Height  int    `orm:"height" json:"height"`
}

func (*Ewei_shop_pc_adv) TableName() string {
	return "ewei_shop_pc_adv"
}

type Ewei_shop_task_list struct {
	Status         int    `orm:"status" json:"status"`
	Displayorder   int    `orm:"displayorder" json:"displayorder"`
	Id             int    `orm:"id" json:"id"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Title          string `orm:"title" json:"title"`
	Image          string `orm:"image" json:"image"`
	Type           string `orm:"type" json:"type"`
	Starttime      string `orm:"starttime" json:"starttime"`
	Endtime        string `orm:"endtime" json:"endtime"`
	Demand         int    `orm:"demand" json:"demand"`
	Requiregoods   string `orm:"requiregoods" json:"requiregoods"`
	Picktype       int    `orm:"picktype" json:"picktype"`
	StopType       int    `orm:"stop_type" json:"stop_type"`
	StopLimit      int    `orm:"stop_limit" json:"stop_limit"`
	StopTime       string `orm:"stop_time" json:"stop_time"`
	StopCycle      int    `orm:"stop_cycle" json:"stop_cycle"`
	RepeatType     int    `orm:"repeat_type" json:"repeat_type"`
	RepeatInterval int    `orm:"repeat_interval" json:"repeat_interval"`
	RepeatCycle    int    `orm:"repeat_cycle" json:"repeat_cycle"`
	Reward         string `orm:"reward" json:"reward"`
	Followreward   string `orm:"followreward" json:"followreward"`
	GoodsLimit     int    `orm:"goods_limit" json:"goods_limit"`
	Notice         string `orm:"notice" json:"notice"`
	DesignData     string `orm:"design_data" json:"design_data"`
	DesignBg       string `orm:"design_bg" json:"design_bg"`
	NativeData     string `orm:"native_data" json:"native_data"`
	NativeData2    string `orm:"native_data2" json:"native_data2"`
	NativeData3    string `orm:"native_data3" json:"native_data3"`
	Reward2        string `orm:"reward2" json:"reward2"`
	Reward3        string `orm:"reward3" json:"reward3"`
	Level2         int    `orm:"level2" json:"level2"`
	Level3         int    `orm:"level3" json:"level3"`
	MemberGroup    string `orm:"member_group" json:"member_group"`
	AutoPick       int    `orm:"auto_pick" json:"auto_pick"`
	KeywordPick    string `orm:"keyword_pick" json:"keyword_pick"`
	Verb           string `orm:"verb" json:"verb"`
	Unit           string `orm:"unit" json:"unit"`
	MemberLevel    int    `orm:"member_level" json:"member_level"`
}

func (*Ewei_shop_task_list) TableName() string {
	return "ewei_shop_task_list"
}

type Ewei_shop_wxapp_tmessage struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Name            string `orm:"name" json:"name"`
	Templateid      string `orm:"templateid" json:"templateid"`
	Datas           string `orm:"datas" json:"datas"`
	EmphasisKeyword int    `orm:"emphasis_keyword" json:"emphasis_keyword"`
	Status          int    `orm:"status" json:"status"`
}

func (*Ewei_shop_wxapp_tmessage) TableName() string {
	return "ewei_shop_wxapp_tmessage"
}

type Activity_clerks struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Uid      int    `orm:"uid" json:"uid"`
	Storeid  int    `orm:"storeid" json:"storeid"`
	Name     string `orm:"name" json:"name"`
	Password string `orm:"password" json:"password"`
	Mobile   string `orm:"mobile" json:"mobile"`
	Openid   string `orm:"openid" json:"openid"`
	Nickname string `orm:"nickname" json:"nickname"`
}

func (*Activity_clerks) TableName() string {
	return "activity_clerks"
}

type Ewei_shop_city_express struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Merchid      int     `orm:"merchid" json:"merchid"`
	StartFee     float64 `orm:"start_fee" json:"start_fee"`
	StartKm      int     `orm:"start_km" json:"start_km"`
	PreKm        int     `orm:"pre_km" json:"pre_km"`
	PreKmFee     float64 `orm:"pre_km_fee" json:"pre_km_fee"`
	FixedKm      int     `orm:"fixed_km" json:"fixed_km"`
	FixedFee     float64 `orm:"fixed_fee" json:"fixed_fee"`
	ReceiveGoods int     `orm:"receive_goods" json:"receive_goods"`
	Lng          string  `orm:"lng" json:"lng"`
	Lat          string  `orm:"lat" json:"lat"`
	Range        int     `orm:"range" json:"range"`
	Zoom         int     `orm:"zoom" json:"zoom"`
	ExpressType  int     `orm:"express_type" json:"express_type"`
	Config       string  `orm:"config" json:"config"`
	Tel1         string  `orm:"tel1" json:"tel1"`
	Tel2         string  `orm:"tel2" json:"tel2"`
	IsSum        int     `orm:"is_sum" json:"is_sum"`
	IsDispatch   int     `orm:"is_dispatch" json:"is_dispatch"`
	Enabled      int     `orm:"enabled" json:"enabled"`
	GeoKey       string  `orm:"geo_key" json:"geo_key"`
}

func (*Ewei_shop_city_express) TableName() string {
	return "ewei_shop_city_express"
}

type Ewei_shop_exchange_record struct {
	Id         int     `orm:"id" json:"id"`
	Key        string  `orm:"key" json:"key"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Goods      string  `orm:"goods" json:"goods"`
	Orderid    string  `orm:"orderid" json:"orderid"`
	Time       int     `orm:"time" json:"time"`
	Openid     string  `orm:"openid" json:"openid"`
	Mode       int     `orm:"mode" json:"mode"`
	Balance    float64 `orm:"balance" json:"balance"`
	Red        float64 `orm:"red" json:"red"`
	Coupon     string  `orm:"coupon" json:"coupon"`
	Score      int     `orm:"score" json:"score"`
	Nickname   string  `orm:"nickname" json:"nickname"`
	Groupid    int     `orm:"groupid" json:"groupid"`
	Title      string  `orm:"title" json:"title"`
	Serial     string  `orm:"serial" json:"serial"`
	Ordersn    string  `orm:"ordersn" json:"ordersn"`
	GoodsTitle string  `orm:"goods_title" json:"goods_title"`
}

func (*Ewei_shop_exchange_record) TableName() string {
	return "ewei_shop_exchange_record"
}

type Ewei_shop_member_cart struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Openid        string  `orm:"openid" json:"openid"`
	Goodsid       int     `orm:"goodsid" json:"goodsid"`
	Total         int     `orm:"total" json:"total"`
	Marketprice   float64 `orm:"marketprice" json:"marketprice"`
	Deleted       int     `orm:"deleted" json:"deleted"`
	Optionid      int     `orm:"optionid" json:"optionid"`
	Createtime    int     `orm:"createtime" json:"createtime"`
	Diyformdataid int     `orm:"diyformdataid" json:"diyformdataid"`
	Diyformdata   string  `orm:"diyformdata" json:"diyformdata"`
	Diyformfields string  `orm:"diyformfields" json:"diyformfields"`
	Diyformid     int     `orm:"diyformid" json:"diyformid"`
	Selected      int     `orm:"selected" json:"selected"`
	Selectedadd   int     `orm:"selectedadd" json:"selectedadd"`
	Merchid       int     `orm:"merchid" json:"merchid"`
	Isnewstore    int     `orm:"isnewstore" json:"isnewstore"`
}

func (*Ewei_shop_member_cart) TableName() string {
	return "ewei_shop_member_cart"
}

type Ewei_shop_sale_coupon_data struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Openid   string `orm:"openid" json:"openid"`
	Couponid int    `orm:"couponid" json:"couponid"`
	Gettime  int    `orm:"gettime" json:"gettime"`
	Gettype  int    `orm:"gettype" json:"gettype"`
	Usedtime int    `orm:"usedtime" json:"usedtime"`
	Orderid  int    `orm:"orderid" json:"orderid"`
}

func (*Ewei_shop_sale_coupon_data) TableName() string {
	return "ewei_shop_sale_coupon_data"
}

type Ewei_shop_sns_member struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Level      int    `orm:"level" json:"level"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Credit     int    `orm:"credit" json:"credit"`
	Sign       string `orm:"sign" json:"sign"`
	Isblack    int    `orm:"isblack" json:"isblack"`
	Notupgrade int    `orm:"notupgrade" json:"notupgrade"`
}

func (*Ewei_shop_sns_member) TableName() string {
	return "ewei_shop_sns_member"
}

type Video_reply struct {
	Id          int    `orm:"id" json:"id"`
	Rid         int    `orm:"rid" json:"rid"`
	Title       string `orm:"title" json:"title"`
	Description string `orm:"description" json:"description"`
	Mediaid     string `orm:"mediaid" json:"mediaid"`
	Createtime  int    `orm:"createtime" json:"createtime"`
}

func (*Video_reply) TableName() string {
	return "video_reply"
}

type Core_paylog struct {
	Plid        int     `orm:"plid" json:"plid"`
	Type        string  `orm:"type" json:"type"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Acid        int     `orm:"acid" json:"acid"`
	Openid      string  `orm:"openid" json:"openid"`
	Uniontid    string  `orm:"uniontid" json:"uniontid"`
	Tid         string  `orm:"tid" json:"tid"`
	Fee         float64 `orm:"fee" json:"fee"`
	Status      int     `orm:"status" json:"status"`
	Module      string  `orm:"module" json:"module"`
	Tag         string  `orm:"tag" json:"tag"`
	IsUsecard   int     `orm:"is_usecard" json:"is_usecard"`
	CardType    int     `orm:"card_type" json:"card_type"`
	CardId      string  `orm:"card_id" json:"card_id"`
	CardFee     float64 `orm:"card_fee" json:"card_fee"`
	EncryptCode string  `orm:"encrypt_code" json:"encrypt_code"`
}

func (*Core_paylog) TableName() string {
	return "core_paylog"
}

type Ewei_shop_bargain_account struct {
	Id          int    `orm:"id" json:"id"`
	MallName    string `orm:"mall_name" json:"mall_name"`
	Banner      string `orm:"banner" json:"banner"`
	MallTitle   string `orm:"mall_title" json:"mall_title"`
	MallContent string `orm:"mall_content" json:"mall_content"`
	MallLogo    string `orm:"mall_logo" json:"mall_logo"`
	Message     int    `orm:"message" json:"message"`
	Partin      int    `orm:"partin" json:"partin"`
	Rule        string `orm:"rule" json:"rule"`
	EndMessage  int    `orm:"end_message" json:"end_message"`
	FollowSwi   int    `orm:"follow_swi" json:"follow_swi"`
	Sharestyle  int    `orm:"sharestyle" json:"sharestyle"`
}

func (*Ewei_shop_bargain_account) TableName() string {
	return "ewei_shop_bargain_account"
}

type Ewei_shop_creditshop_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Advimg       string `orm:"advimg" json:"advimg"`
	Advurl       string `orm:"advurl" json:"advurl"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
}

func (*Ewei_shop_creditshop_category) TableName() string {
	return "ewei_shop_creditshop_category"
}

type Ewei_shop_exhelper_esheet struct {
	Id      int    `orm:"id" json:"id"`
	Name    string `orm:"name" json:"name"`
	Express string `orm:"express" json:"express"`
	Code    string `orm:"code" json:"code"`
	Datas   string `orm:"datas" json:"datas"`
}

func (*Ewei_shop_exhelper_esheet) TableName() string {
	return "ewei_shop_exhelper_esheet"
}

type Ewei_shop_exhelper_sys struct {
	Id        int    `orm:"id" json:"id"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Ip        string `orm:"ip" json:"ip"`
	Port      int    `orm:"port" json:"port"`
	IpCloud   string `orm:"ip_cloud" json:"ip_cloud"`
	PortCloud int    `orm:"port_cloud" json:"port_cloud"`
	IsCloud   int    `orm:"is_cloud" json:"is_cloud"`
	Merchid   int    `orm:"merchid" json:"merchid"`
	Ebusiness string `orm:"ebusiness" json:"ebusiness"`
	Apikey    string `orm:"apikey" json:"apikey"`
}

func (*Ewei_shop_exhelper_sys) TableName() string {
	return "ewei_shop_exhelper_sys"
}

type Ewei_shop_live_coupon struct {
	Id          int `orm:"id" json:"id"`
	Uniacid     int `orm:"uniacid" json:"uniacid"`
	Roomid      int `orm:"roomid" json:"roomid"`
	Couponid    int `orm:"couponid" json:"couponid"`
	Coupontotal int `orm:"coupontotal" json:"coupontotal"`
	Couponlimit int `orm:"couponlimit" json:"couponlimit"`
}

func (*Ewei_shop_live_coupon) TableName() string {
	return "ewei_shop_live_coupon"
}

type Ewei_shop_merch_perm_log struct {
	Id         int    `orm:"id" json:"id"`
	Uid        int    `orm:"uid" json:"uid"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Name       string `orm:"name" json:"name"`
	Type       string `orm:"type" json:"type"`
	Op         string `orm:"op" json:"op"`
	Ip         string `orm:"ip" json:"ip"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Merchid    int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_perm_log) TableName() string {
	return "ewei_shop_merch_perm_log"
}

type Ewei_shop_verifygoods_log struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Verifygoodsid int    `orm:"verifygoodsid" json:"verifygoodsid"`
	Salerid       int    `orm:"salerid" json:"salerid"`
	Storeid       int    `orm:"storeid" json:"storeid"`
	Verifynum     int    `orm:"verifynum" json:"verifynum"`
	Verifydate    int    `orm:"verifydate" json:"verifydate"`
	Remarks       string `orm:"remarks" json:"remarks"`
}

func (*Ewei_shop_verifygoods_log) TableName() string {
	return "ewei_shop_verifygoods_log"
}

type Ewei_shop_member_group struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Groupname   string `orm:"groupname" json:"groupname"`
	Description string `orm:"description" json:"description"`
}

func (*Ewei_shop_member_group) TableName() string {
	return "ewei_shop_member_group"
}

type Ewei_shop_member_printer struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Title      string `orm:"title" json:"title"`
	Type       int    `orm:"type" json:"type"`
	PrintData  string `orm:"print_data" json:"print_data"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Merchid    int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_member_printer) TableName() string {
	return "ewei_shop_member_printer"
}

type Ewei_shop_merch_billo struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Billid     int     `orm:"billid" json:"billid"`
	Orderid    int     `orm:"orderid" json:"orderid"`
	Ordermoney float64 `orm:"ordermoney" json:"ordermoney"`
}

func (*Ewei_shop_merch_billo) TableName() string {
	return "ewei_shop_merch_billo"
}

type Ewei_shop_nav struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Navname      string `orm:"navname" json:"navname"`
	Icon         string `orm:"icon" json:"icon"`
	Url          string `orm:"url" json:"url"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
	Iswxapp      int    `orm:"iswxapp" json:"iswxapp"`
}

func (*Ewei_shop_nav) TableName() string {
	return "ewei_shop_nav"
}

type Ewei_shop_system_article struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Author       string `orm:"author" json:"author"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Content      string `orm:"content" json:"content"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Cate         int    `orm:"cate" json:"cate"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_article) TableName() string {
	return "ewei_shop_system_article"
}

type Ewei_shop_verifygoods struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Openid          string `orm:"openid" json:"openid"`
	Orderid         int    `orm:"orderid" json:"orderid"`
	Ordergoodsid    int    `orm:"ordergoodsid" json:"ordergoodsid"`
	Storeid         int    `orm:"storeid" json:"storeid"`
	Starttime       int    `orm:"starttime" json:"starttime"`
	Limitdays       int    `orm:"limitdays" json:"limitdays"`
	Limitnum        int    `orm:"limitnum" json:"limitnum"`
	Used            int    `orm:"used" json:"used"`
	Verifycode      string `orm:"verifycode" json:"verifycode"`
	Codeinvalidtime int    `orm:"codeinvalidtime" json:"codeinvalidtime"`
	Invalid         int    `orm:"invalid" json:"invalid"`
	Getcard         int    `orm:"getcard" json:"getcard"`
	Activecard      int    `orm:"activecard" json:"activecard"`
	Cardcode        string `orm:"cardcode" json:"cardcode"`
	Limittype       int    `orm:"limittype" json:"limittype"`
	Limitdate       int    `orm:"limitdate" json:"limitdate"`
}

func (*Ewei_shop_verifygoods) TableName() string {
	return "ewei_shop_verifygoods"
}

type Ewei_shop_wxapp_startadv struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Status       int    `orm:"status" json:"status"`
	Name         string `orm:"name" json:"name"`
	Data         string `orm:"data" json:"data"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Lastedittime int    `orm:"lastedittime" json:"lastedittime"`
}

func (*Ewei_shop_wxapp_startadv) TableName() string {
	return "ewei_shop_wxapp_startadv"
}

type Ewei_shop_abonus_billp struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Billid       int     `orm:"billid" json:"billid"`
	Openid       string  `orm:"openid" json:"openid"`
	Payno        string  `orm:"payno" json:"payno"`
	Paytype      int     `orm:"paytype" json:"paytype"`
	Bonus1       float64 `orm:"bonus1" json:"bonus1"`
	Bonus2       float64 `orm:"bonus2" json:"bonus2"`
	Bonus3       float64 `orm:"bonus3" json:"bonus3"`
	Money1       float64 `orm:"money1" json:"money1"`
	Realmoney1   float64 `orm:"realmoney1" json:"realmoney1"`
	Paymoney1    float64 `orm:"paymoney1" json:"paymoney1"`
	Money2       float64 `orm:"money2" json:"money2"`
	Realmoney2   float64 `orm:"realmoney2" json:"realmoney2"`
	Paymoney2    float64 `orm:"paymoney2" json:"paymoney2"`
	Money3       float64 `orm:"money3" json:"money3"`
	Realmoney3   float64 `orm:"realmoney3" json:"realmoney3"`
	Paymoney3    float64 `orm:"paymoney3" json:"paymoney3"`
	Chargemoney1 float64 `orm:"chargemoney1" json:"chargemoney1"`
	Chargemoney2 float64 `orm:"chargemoney2" json:"chargemoney2"`
	Chargemoney3 float64 `orm:"chargemoney3" json:"chargemoney3"`
	Charge       float64 `orm:"charge" json:"charge"`
	Status       int     `orm:"status" json:"status"`
	Reason       string  `orm:"reason" json:"reason"`
	Paytime      int     `orm:"paytime" json:"paytime"`
}

func (*Ewei_shop_abonus_billp) TableName() string {
	return "ewei_shop_abonus_billp"
}

type Ewei_shop_seckill_category struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
}

func (*Ewei_shop_seckill_category) TableName() string {
	return "ewei_shop_seckill_category"
}

type Ewei_shop_sns_category struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Advimg       string `orm:"advimg" json:"advimg"`
	Advurl       string `orm:"advurl" json:"advurl"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
}

func (*Ewei_shop_sns_category) TableName() string {
	return "ewei_shop_sns_category"
}

type Ewei_shop_system_company_article struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Author       string `orm:"author" json:"author"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Content      string `orm:"content" json:"content"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Cate         int    `orm:"cate" json:"cate"`
	Status       int    `orm:"status" json:"status"`
}

func (*Ewei_shop_system_company_article) TableName() string {
	return "ewei_shop_system_company_article"
}

type Site_multi struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Title    string `orm:"title" json:"title"`
	Styleid  int    `orm:"styleid" json:"styleid"`
	SiteInfo string `orm:"site_info" json:"site_info"`
	Status   int    `orm:"status" json:"status"`
	Bindhost string `orm:"bindhost" json:"bindhost"`
}

func (*Site_multi) TableName() string {
	return "site_multi"
}

type Site_page struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Multiid     int    `orm:"multiid" json:"multiid"`
	Title       string `orm:"title" json:"title"`
	Description string `orm:"description" json:"description"`
	Params      string `orm:"params" json:"params"`
	Html        string `orm:"html" json:"html"`
	Multipage   string `orm:"multipage" json:"multipage"`
	Type        int    `orm:"type" json:"type"`
	Status      int    `orm:"status" json:"status"`
	Createtime  int    `orm:"createtime" json:"createtime"`
	Goodnum     int    `orm:"goodnum" json:"goodnum"`
}

func (*Site_page) TableName() string {
	return "site_page"
}

type Ewei_shop_author_team_pay struct {
	Id       int     `orm:"id" json:"id"`
	Uniacid  int     `orm:"uniacid" json:"uniacid"`
	Teamid   int     `orm:"teamid" json:"teamid"`
	Mid      int     `orm:"mid" json:"mid"`
	Payno    string  `orm:"payno" json:"payno"`
	Money    float64 `orm:"money" json:"money"`
	Paymoney float64 `orm:"paymoney" json:"paymoney"`
	Paytime  int     `orm:"paytime" json:"paytime"`
}

func (*Ewei_shop_author_team_pay) TableName() string {
	return "ewei_shop_author_team_pay"
}

type Ewei_shop_globonus_bill struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Billno          string  `orm:"billno" json:"billno"`
	Paytype         int     `orm:"paytype" json:"paytype"`
	Year            int     `orm:"year" json:"year"`
	Month           int     `orm:"month" json:"month"`
	Week            int     `orm:"week" json:"week"`
	Ordercount      int     `orm:"ordercount" json:"ordercount"`
	Ordermoney      float64 `orm:"ordermoney" json:"ordermoney"`
	Bonusmoney      float64 `orm:"bonusmoney" json:"bonusmoney"`
	BonusmoneySend  float64 `orm:"bonusmoney_send" json:"bonusmoney_send"`
	BonusmoneyPay   float64 `orm:"bonusmoney_pay" json:"bonusmoney_pay"`
	Paytime         int     `orm:"paytime" json:"paytime"`
	Partnercount    int     `orm:"partnercount" json:"partnercount"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Status          int     `orm:"status" json:"status"`
	Starttime       int     `orm:"starttime" json:"starttime"`
	Endtime         int     `orm:"endtime" json:"endtime"`
	Confirmtime     int     `orm:"confirmtime" json:"confirmtime"`
	Bonusordermoney float64 `orm:"bonusordermoney" json:"bonusordermoney"`
	Bonusrate       float64 `orm:"bonusrate" json:"bonusrate"`
}

func (*Ewei_shop_globonus_bill) TableName() string {
	return "ewei_shop_globonus_bill"
}

type Mc_members struct {
	Uid             int     `orm:"uid" json:"uid"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Mobile          string  `orm:"mobile" json:"mobile"`
	Email           string  `orm:"email" json:"email"`
	Password        string  `orm:"password" json:"password"`
	Salt            string  `orm:"salt" json:"salt"`
	Groupid         int     `orm:"groupid" json:"groupid"`
	Credit1         float64 `orm:"credit1" json:"credit1"`
	Credit2         float64 `orm:"credit2" json:"credit2"`
	Credit3         float64 `orm:"credit3" json:"credit3"`
	Credit4         float64 `orm:"credit4" json:"credit4"`
	Credit5         float64 `orm:"credit5" json:"credit5"`
	Credit6         float64 `orm:"credit6" json:"credit6"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Realname        string  `orm:"realname" json:"realname"`
	Nickname        string  `orm:"nickname" json:"nickname"`
	Avatar          string  `orm:"avatar" json:"avatar"`
	Qq              string  `orm:"qq" json:"qq"`
	Vip             int     `orm:"vip" json:"vip"`
	Gender          int     `orm:"gender" json:"gender"`
	Birthyear       int     `orm:"birthyear" json:"birthyear"`
	Birthmonth      int     `orm:"birthmonth" json:"birthmonth"`
	Birthday        int     `orm:"birthday" json:"birthday"`
	Constellation   string  `orm:"constellation" json:"constellation"`
	Zodiac          string  `orm:"zodiac" json:"zodiac"`
	Telephone       string  `orm:"telephone" json:"telephone"`
	Idcard          string  `orm:"idcard" json:"idcard"`
	Studentid       string  `orm:"studentid" json:"studentid"`
	Grade           string  `orm:"grade" json:"grade"`
	Address         string  `orm:"address" json:"address"`
	Zipcode         string  `orm:"zipcode" json:"zipcode"`
	Nationality     string  `orm:"nationality" json:"nationality"`
	Resideprovince  string  `orm:"resideprovince" json:"resideprovince"`
	Residecity      string  `orm:"residecity" json:"residecity"`
	Residedist      string  `orm:"residedist" json:"residedist"`
	Graduateschool  string  `orm:"graduateschool" json:"graduateschool"`
	Company         string  `orm:"company" json:"company"`
	Education       string  `orm:"education" json:"education"`
	Occupation      string  `orm:"occupation" json:"occupation"`
	Position        string  `orm:"position" json:"position"`
	Revenue         string  `orm:"revenue" json:"revenue"`
	Affectivestatus string  `orm:"affectivestatus" json:"affectivestatus"`
	Lookingfor      string  `orm:"lookingfor" json:"lookingfor"`
	Bloodtype       string  `orm:"bloodtype" json:"bloodtype"`
	Height          string  `orm:"height" json:"height"`
	Weight          string  `orm:"weight" json:"weight"`
	Alipay          string  `orm:"alipay" json:"alipay"`
	Msn             string  `orm:"msn" json:"msn"`
	Taobao          string  `orm:"taobao" json:"taobao"`
	Site            string  `orm:"site" json:"site"`
	Bio             string  `orm:"bio" json:"bio"`
	Interest        string  `orm:"interest" json:"interest"`
}

func (*Mc_members) TableName() string {
	return "mc_members"
}

type Users_profile struct {
	Id                 int    `orm:"id" json:"id"`
	Uid                int    `orm:"uid" json:"uid"`
	Createtime         int    `orm:"createtime" json:"createtime"`
	Edittime           int    `orm:"edittime" json:"edittime"`
	Realname           string `orm:"realname" json:"realname"`
	Nickname           string `orm:"nickname" json:"nickname"`
	Avatar             string `orm:"avatar" json:"avatar"`
	Qq                 string `orm:"qq" json:"qq"`
	Mobile             string `orm:"mobile" json:"mobile"`
	Fakeid             string `orm:"fakeid" json:"fakeid"`
	Vip                int    `orm:"vip" json:"vip"`
	Gender             int    `orm:"gender" json:"gender"`
	Birthyear          int    `orm:"birthyear" json:"birthyear"`
	Birthmonth         int    `orm:"birthmonth" json:"birthmonth"`
	Birthday           int    `orm:"birthday" json:"birthday"`
	Constellation      string `orm:"constellation" json:"constellation"`
	Zodiac             string `orm:"zodiac" json:"zodiac"`
	Telephone          string `orm:"telephone" json:"telephone"`
	Idcard             string `orm:"idcard" json:"idcard"`
	Studentid          string `orm:"studentid" json:"studentid"`
	Grade              string `orm:"grade" json:"grade"`
	Address            string `orm:"address" json:"address"`
	Zipcode            string `orm:"zipcode" json:"zipcode"`
	Nationality        string `orm:"nationality" json:"nationality"`
	Resideprovince     string `orm:"resideprovince" json:"resideprovince"`
	Residecity         string `orm:"residecity" json:"residecity"`
	Residedist         string `orm:"residedist" json:"residedist"`
	Graduateschool     string `orm:"graduateschool" json:"graduateschool"`
	Company            string `orm:"company" json:"company"`
	Education          string `orm:"education" json:"education"`
	Occupation         string `orm:"occupation" json:"occupation"`
	Position           string `orm:"position" json:"position"`
	Revenue            string `orm:"revenue" json:"revenue"`
	Affectivestatus    string `orm:"affectivestatus" json:"affectivestatus"`
	Lookingfor         string `orm:"lookingfor" json:"lookingfor"`
	Bloodtype          string `orm:"bloodtype" json:"bloodtype"`
	Height             string `orm:"height" json:"height"`
	Weight             string `orm:"weight" json:"weight"`
	Alipay             string `orm:"alipay" json:"alipay"`
	Msn                string `orm:"msn" json:"msn"`
	Email              string `orm:"email" json:"email"`
	Taobao             string `orm:"taobao" json:"taobao"`
	Site               string `orm:"site" json:"site"`
	Bio                string `orm:"bio" json:"bio"`
	Interest           string `orm:"interest" json:"interest"`
	Workerid           string `orm:"workerid" json:"workerid"`
	IsSendMobileStatus int    `orm:"is_send_mobile_status" json:"is_send_mobile_status"`
	SendExpireStatus   int    `orm:"send_expire_status" json:"send_expire_status"`
}

func (*Users_profile) TableName() string {
	return "users_profile"
}

type Ewei_shop_gift struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Title        string  `orm:"title" json:"title"`
	Thumb        string  `orm:"thumb" json:"thumb"`
	Activity     int     `orm:"activity" json:"activity"`
	Orderprice   float64 `orm:"orderprice" json:"orderprice"`
	Goodsid      string  `orm:"goodsid" json:"goodsid"`
	Giftgoodsid  string  `orm:"giftgoodsid" json:"giftgoodsid"`
	Starttime    int     `orm:"starttime" json:"starttime"`
	Endtime      int     `orm:"endtime" json:"endtime"`
	Status       int     `orm:"status" json:"status"`
	Displayorder int     `orm:"displayorder" json:"displayorder"`
	ShareTitle   string  `orm:"share_title" json:"share_title"`
	ShareIcon    string  `orm:"share_icon" json:"share_icon"`
	ShareDesc    string  `orm:"share_desc" json:"share_desc"`
}

func (*Ewei_shop_gift) TableName() string {
	return "ewei_shop_gift"
}

type Ewei_shop_lottery_log struct {
	LogId       int    `orm:"log_id" json:"log_id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	LotteryId   int    `orm:"lottery_id" json:"lottery_id"`
	JoinUser    string `orm:"join_user" json:"join_user"`
	LotteryData string `orm:"lottery_data" json:"lottery_data"`
	IsReward    int    `orm:"is_reward" json:"is_reward"`
	Addtime     int    `orm:"addtime" json:"addtime"`
}

func (*Ewei_shop_lottery_log) TableName() string {
	return "ewei_shop_lottery_log"
}

type Ewei_shop_qa_question struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Cate         int    `orm:"cate" json:"cate"`
	Title        string `orm:"title" json:"title"`
	Keywords     string `orm:"keywords" json:"keywords"`
	Content      string `orm:"content" json:"content"`
	Status       int    `orm:"status" json:"status"`
	Isrecommand  int    `orm:"isrecommand" json:"isrecommand"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Lastedittime int    `orm:"lastedittime" json:"lastedittime"`
}

func (*Ewei_shop_qa_question) TableName() string {
	return "ewei_shop_qa_question"
}

type Ewei_shop_task_join struct {
	JoinId        int    `orm:"join_id" json:"join_id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	JoinUser      string `orm:"join_user" json:"join_user"`
	TaskId        int    `orm:"task_id" json:"task_id"`
	TaskType      int    `orm:"task_type" json:"task_type"`
	Needcount     int    `orm:"needcount" json:"needcount"`
	Completecount int    `orm:"completecount" json:"completecount"`
	RewardData    string `orm:"reward_data" json:"reward_data"`
	IsReward      int    `orm:"is_reward" json:"is_reward"`
	Failtime      int    `orm:"failtime" json:"failtime"`
	Addtime       int    `orm:"addtime" json:"addtime"`
}

func (*Ewei_shop_task_join) TableName() string {
	return "ewei_shop_task_join"
}

type Site_store_goods struct {
	Id           int     `orm:"id" json:"id"`
	Type         int     `orm:"type" json:"type"`
	Title        string  `orm:"title" json:"title"`
	Module       string  `orm:"module" json:"module"`
	AccountNum   int     `orm:"account_num" json:"account_num"`
	WxappNum     int     `orm:"wxapp_num" json:"wxapp_num"`
	Price        float64 `orm:"price" json:"price"`
	Unit         string  `orm:"unit" json:"unit"`
	Slide        string  `orm:"slide" json:"slide"`
	CategoryId   int     `orm:"category_id" json:"category_id"`
	TitleInitial string  `orm:"title_initial" json:"title_initial"`
	Status       int     `orm:"status" json:"status"`
	Createtime   int     `orm:"createtime" json:"createtime"`
	Synopsis     string  `orm:"synopsis" json:"synopsis"`
	Description  string  `orm:"description" json:"description"`
	ModuleGroup  int     `orm:"module_group" json:"module_group"`
	ApiNum       int     `orm:"api_num" json:"api_num"`
}

func (*Site_store_goods) TableName() string {
	return "site_store_goods"
}

type Coupon_modules struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Acid     int    `orm:"acid" json:"acid"`
	Couponid int    `orm:"couponid" json:"couponid"`
	Module   string `orm:"module" json:"module"`
}

func (*Coupon_modules) TableName() string {
	return "coupon_modules"
}

type Ewei_shop_article_comment struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Articleid       int    `orm:"articleid" json:"articleid"`
	Openid          string `orm:"openid" json:"openid"`
	Nickname        string `orm:"nickname" json:"nickname"`
	Headimgurl      string `orm:"headimgurl" json:"headimgurl"`
	Content         string `orm:"content" json:"content"`
	ReplyCreatetime int    `orm:"reply_createtime" json:"reply_createtime"`
	Createtime      int    `orm:"createtime" json:"createtime"`
	Deleted         int    `orm:"deleted" json:"deleted"`
	Status          int    `orm:"status" json:"status"`
	ReplyContent    string `orm:"reply_content" json:"reply_content"`
}

func (*Ewei_shop_article_comment) TableName() string {
	return "ewei_shop_article_comment"
}

type Ewei_shop_commission_apply struct {
	Id             int     `orm:"id" json:"id"`
	Uniacid        int     `orm:"uniacid" json:"uniacid"`
	Applyno        string  `orm:"applyno" json:"applyno"`
	Mid            int     `orm:"mid" json:"mid"`
	Type           int     `orm:"type" json:"type"`
	Orderids       string  `orm:"orderids" json:"orderids"`
	Commission     float64 `orm:"commission" json:"commission"`
	CommissionPay  float64 `orm:"commission_pay" json:"commission_pay"`
	Content        string  `orm:"content" json:"content"`
	Status         int     `orm:"status" json:"status"`
	Applytime      int     `orm:"applytime" json:"applytime"`
	Checktime      int     `orm:"checktime" json:"checktime"`
	Paytime        int     `orm:"paytime" json:"paytime"`
	Invalidtime    int     `orm:"invalidtime" json:"invalidtime"`
	Refusetime     int     `orm:"refusetime" json:"refusetime"`
	Realmoney      float64 `orm:"realmoney" json:"realmoney"`
	Charge         float64 `orm:"charge" json:"charge"`
	Deductionmoney float64 `orm:"deductionmoney" json:"deductionmoney"`
	Beginmoney     float64 `orm:"beginmoney" json:"beginmoney"`
	Endmoney       float64 `orm:"endmoney" json:"endmoney"`
	Alipay         string  `orm:"alipay" json:"alipay"`
	Bankname       string  `orm:"bankname" json:"bankname"`
	Bankcard       string  `orm:"bankcard" json:"bankcard"`
	Realname       string  `orm:"realname" json:"realname"`
	Repurchase     float64 `orm:"repurchase" json:"repurchase"`
	Alipay1        string  `orm:"alipay1" json:"alipay1"`
	Bankname1      string  `orm:"bankname1" json:"bankname1"`
	Bankcard1      string  `orm:"bankcard1" json:"bankcard1"`
	Sendmoney      float64 `orm:"sendmoney" json:"sendmoney"`
	Senddata       string  `orm:"senddata" json:"senddata"`
}

func (*Ewei_shop_commission_apply) TableName() string {
	return "ewei_shop_commission_apply"
}

type Ewei_shop_form struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Isrequire    int    `orm:"isrequire" json:"isrequire"`
	Key          string `orm:"key" json:"key"`
	Title        string `orm:"title" json:"title"`
	Type         string `orm:"type" json:"type"`
	Values       string `orm:"values" json:"values"`
	Cate         int    `orm:"cate" json:"cate"`
}

func (*Ewei_shop_form) TableName() string {
	return "ewei_shop_form"
}

type Ewei_shop_live_goods struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Goodsid      int     `orm:"goodsid" json:"goodsid"`
	Liveid       int     `orm:"liveid" json:"liveid"`
	Liveprice    float64 `orm:"liveprice" json:"liveprice"`
	Minliveprice float64 `orm:"minliveprice" json:"minliveprice"`
	Maxliveprice float64 `orm:"maxliveprice" json:"maxliveprice"`
}

func (*Ewei_shop_live_goods) TableName() string {
	return "ewei_shop_live_goods"
}

type Ewei_shop_designer struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Pagename   string `orm:"pagename" json:"pagename"`
	Pagetype   int    `orm:"pagetype" json:"pagetype"`
	Pageinfo   string `orm:"pageinfo" json:"pageinfo"`
	Createtime string `orm:"createtime" json:"createtime"`
	Keyword    string `orm:"keyword" json:"keyword"`
	Savetime   string `orm:"savetime" json:"savetime"`
	Setdefault int    `orm:"setdefault" json:"setdefault"`
	Datas      string `orm:"datas" json:"datas"`
}

func (*Ewei_shop_designer) TableName() string {
	return "ewei_shop_designer"
}

type Ewei_shop_merch_account struct {
	Id        int    `orm:"id" json:"id"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Openid    string `orm:"openid" json:"openid"`
	Merchid   int    `orm:"merchid" json:"merchid"`
	Username  string `orm:"username" json:"username"`
	Pwd       string `orm:"pwd" json:"pwd"`
	Salt      string `orm:"salt" json:"salt"`
	Status    int    `orm:"status" json:"status"`
	Perms     string `orm:"perms" json:"perms"`
	Isfounder int    `orm:"isfounder" json:"isfounder"`
	Lastip    string `orm:"lastip" json:"lastip"`
	Lastvisit string `orm:"lastvisit" json:"lastvisit"`
	Roleid    int    `orm:"roleid" json:"roleid"`
}

func (*Ewei_shop_merch_account) TableName() string {
	return "ewei_shop_merch_account"
}

type Ewei_shop_sns_complaincate struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Name         string `orm:"name" json:"name"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_sns_complaincate) TableName() string {
	return "ewei_shop_sns_complaincate"
}

type Site_templates struct {
	Id          int    `orm:"id" json:"id"`
	Name        string `orm:"name" json:"name"`
	Title       string `orm:"title" json:"title"`
	Version     string `orm:"version" json:"version"`
	Description string `orm:"description" json:"description"`
	Author      string `orm:"author" json:"author"`
	Url         string `orm:"url" json:"url"`
	Type        string `orm:"type" json:"type"`
	Sections    int    `orm:"sections" json:"sections"`
}

func (*Site_templates) TableName() string {
	return "site_templates"
}

type Ewei_shop_article_category struct {
	Id           int    `orm:"id" json:"id"`
	CategoryName string `orm:"category_name" json:"category_name"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Isshow       int    `orm:"isshow" json:"isshow"`
}

func (*Ewei_shop_article_category) TableName() string {
	return "ewei_shop_article_category"
}

type Ewei_shop_member_message_template_type struct {
	Id           int    `orm:"id" json:"id"`
	Name         string `orm:"name" json:"name"`
	Typecode     string `orm:"typecode" json:"typecode"`
	Templatecode string `orm:"templatecode" json:"templatecode"`
	Templateid   string `orm:"templateid" json:"templateid"`
	Templatename string `orm:"templatename" json:"templatename"`
	Content      string `orm:"content" json:"content"`
	Showtotaladd int    `orm:"showtotaladd" json:"showtotaladd"`
	Typegroup    string `orm:"typegroup" json:"typegroup"`
	Groupname    string `orm:"groupname" json:"groupname"`
}

func (*Ewei_shop_member_message_template_type) TableName() string {
	return "ewei_shop_member_message_template_type"
}

type Ewei_shop_merch_user struct {
	Id            int     `orm:"id" json:"id"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Regid         int     `orm:"regid" json:"regid"`
	Openid        string  `orm:"openid" json:"openid"`
	Groupid       int     `orm:"groupid" json:"groupid"`
	Merchno       string  `orm:"merchno" json:"merchno"`
	Merchname     string  `orm:"merchname" json:"merchname"`
	Salecate      string  `orm:"salecate" json:"salecate"`
	Desc          string  `orm:"desc" json:"desc"`
	Realname      string  `orm:"realname" json:"realname"`
	Mobile        string  `orm:"mobile" json:"mobile"`
	Status        int     `orm:"status" json:"status"`
	Accounttime   int     `orm:"accounttime" json:"accounttime"`
	Diyformdata   string  `orm:"diyformdata" json:"diyformdata"`
	Diyformfields string  `orm:"diyformfields" json:"diyformfields"`
	Applytime     int     `orm:"applytime" json:"applytime"`
	Accounttotal  int     `orm:"accounttotal" json:"accounttotal"`
	Remark        string  `orm:"remark" json:"remark"`
	Jointime      int     `orm:"jointime" json:"jointime"`
	Accountid     int     `orm:"accountid" json:"accountid"`
	Sets          string  `orm:"sets" json:"sets"`
	Logo          string  `orm:"logo" json:"logo"`
	Payopenid     string  `orm:"payopenid" json:"payopenid"`
	Payrate       float64 `orm:"payrate" json:"payrate"`
	Isrecommand   int     `orm:"isrecommand" json:"isrecommand"`
	Cateid        int     `orm:"cateid" json:"cateid"`
	Address       string  `orm:"address" json:"address"`
	Tel           string  `orm:"tel" json:"tel"`
	Lat           string  `orm:"lat" json:"lat"`
	Lng           string  `orm:"lng" json:"lng"`
	Pluginset     string  `orm:"pluginset" json:"pluginset"`
	Uname         string  `orm:"uname" json:"uname"`
	Upass         string  `orm:"upass" json:"upass"`
	Maxgoods      int     `orm:"maxgoods" json:"maxgoods"`
}

func (*Ewei_shop_merch_user) TableName() string {
	return "ewei_shop_merch_user"
}

type Ewei_shop_sale_coupon struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Name       string  `orm:"name" json:"name"`
	Type       int     `orm:"type" json:"type"`
	Ckey       float64 `orm:"ckey" json:"ckey"`
	Cvalue     float64 `orm:"cvalue" json:"cvalue"`
	Nums       int     `orm:"nums" json:"nums"`
	Createtime int     `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_sale_coupon) TableName() string {
	return "ewei_shop_sale_coupon"
}

type Modules_rank struct {
	Id         int    `orm:"id" json:"id"`
	ModuleName string `orm:"module_name" json:"module_name"`
	Uid        int    `orm:"uid" json:"uid"`
	Rank       int    `orm:"rank" json:"rank"`
}

func (*Modules_rank) TableName() string {
	return "modules_rank"
}

type Core_resource struct {
	Mid      int    `orm:"mid" json:"mid"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	MediaId  string `orm:"media_id" json:"media_id"`
	Trunk    int    `orm:"trunk" json:"trunk"`
	Type     string `orm:"type" json:"type"`
	Dateline int    `orm:"dateline" json:"dateline"`
}

func (*Core_resource) TableName() string {
	return "core_resource"
}

type Ewei_shop_carrier struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Realname     string `orm:"realname" json:"realname"`
	Mobile       string `orm:"mobile" json:"mobile"`
	Address      string `orm:"address" json:"address"`
	Deleted      int    `orm:"deleted" json:"deleted"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_carrier) TableName() string {
	return "ewei_shop_carrier"
}

type Ewei_shop_coupon_log struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Logno        string `orm:"logno" json:"logno"`
	Openid       string `orm:"openid" json:"openid"`
	Couponid     int    `orm:"couponid" json:"couponid"`
	Status       int    `orm:"status" json:"status"`
	Paystatus    int    `orm:"paystatus" json:"paystatus"`
	Creditstatus int    `orm:"creditstatus" json:"creditstatus"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Paytype      int    `orm:"paytype" json:"paytype"`
	Getfrom      int    `orm:"getfrom" json:"getfrom"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_coupon_log) TableName() string {
	return "ewei_shop_coupon_log"
}

type Ewei_shop_groups_verify struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Orderid    int    `orm:"orderid" json:"orderid"`
	Verifycode string `orm:"verifycode" json:"verifycode"`
	Storeid    int    `orm:"storeid" json:"storeid"`
	Verifier   string `orm:"verifier" json:"verifier"`
	Isverify   int    `orm:"isverify" json:"isverify"`
	Verifytime int    `orm:"verifytime" json:"verifytime"`
}

func (*Ewei_shop_groups_verify) TableName() string {
	return "ewei_shop_groups_verify"
}

type Ewei_shop_merch_reg struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Openid        string `orm:"openid" json:"openid"`
	Merchname     string `orm:"merchname" json:"merchname"`
	Salecate      string `orm:"salecate" json:"salecate"`
	Desc          string `orm:"desc" json:"desc"`
	Realname      string `orm:"realname" json:"realname"`
	Mobile        string `orm:"mobile" json:"mobile"`
	Status        int    `orm:"status" json:"status"`
	Diyformdata   string `orm:"diyformdata" json:"diyformdata"`
	Diyformfields string `orm:"diyformfields" json:"diyformfields"`
	Applytime     int    `orm:"applytime" json:"applytime"`
	Reason        string `orm:"reason" json:"reason"`
	Uname         string `orm:"uname" json:"uname"`
	Upass         string `orm:"upass" json:"upass"`
}

func (*Ewei_shop_merch_reg) TableName() string {
	return "ewei_shop_merch_reg"
}

type Ewei_shop_package struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Title        string  `orm:"title" json:"title"`
	Price        float64 `orm:"price" json:"price"`
	Freight      float64 `orm:"freight" json:"freight"`
	Thumb        string  `orm:"thumb" json:"thumb"`
	Starttime    int     `orm:"starttime" json:"starttime"`
	Endtime      int     `orm:"endtime" json:"endtime"`
	Goodsid      string  `orm:"goodsid" json:"goodsid"`
	Cash         int     `orm:"cash" json:"cash"`
	ShareTitle   string  `orm:"share_title" json:"share_title"`
	ShareIcon    string  `orm:"share_icon" json:"share_icon"`
	ShareDesc    string  `orm:"share_desc" json:"share_desc"`
	Status       int     `orm:"status" json:"status"`
	Deleted      int     `orm:"deleted" json:"deleted"`
	Displayorder int     `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_package) TableName() string {
	return "ewei_shop_package"
}

type Ewei_shop_perm_user struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Uid        int    `orm:"uid" json:"uid"`
	Username   string `orm:"username" json:"username"`
	Password   string `orm:"password" json:"password"`
	Roleid     int    `orm:"roleid" json:"roleid"`
	Status     int    `orm:"status" json:"status"`
	Perms      string `orm:"perms" json:"perms"`
	Deleted    int    `orm:"deleted" json:"deleted"`
	Realname   string `orm:"realname" json:"realname"`
	Mobile     string `orm:"mobile" json:"mobile"`
	Perms2     string `orm:"perms2" json:"perms2"`
	Openid     string `orm:"openid" json:"openid"`
	OpenidWa   string `orm:"openid_wa" json:"openid_wa"`
	MemberNick string `orm:"member_nick" json:"member_nick"`
}

func (*Ewei_shop_perm_user) TableName() string {
	return "ewei_shop_perm_user"
}

type Ewei_shop_task struct {
	Id          int     `orm:"id" json:"id"`
	Uniacid     int     `orm:"uniacid" json:"uniacid"`
	Title       string  `orm:"title" json:"title"`
	Type        int     `orm:"type" json:"type"`
	Starttime   int     `orm:"starttime" json:"starttime"`
	Endtime     int     `orm:"endtime" json:"endtime"`
	Dotime      int     `orm:"dotime" json:"dotime"`
	Donetime    int     `orm:"donetime" json:"donetime"`
	Timelimit   float64 `orm:"timelimit" json:"timelimit"`
	Keyword     string  `orm:"keyword" json:"keyword"`
	Status      int     `orm:"status" json:"status"`
	Explain     string  `orm:"explain" json:"explain"`
	RequireData string  `orm:"require_data" json:"require_data"`
	RewardData  string  `orm:"reward_data" json:"reward_data"`
	Period      int     `orm:"period" json:"period"`
	Repeat      int     `orm:"repeat" json:"repeat"`
	Maxtimes    int     `orm:"maxtimes" json:"maxtimes"`
	Everyhours  float64 `orm:"everyhours" json:"everyhours"`
	Logo        string  `orm:"logo" json:"logo"`
}

func (*Ewei_shop_task) TableName() string {
	return "ewei_shop_task"
}

type Site_store_create_account struct {
	Id      int `orm:"id" json:"id"`
	Uid     int `orm:"uid" json:"uid"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Type    int `orm:"type" json:"type"`
}

func (*Site_store_create_account) TableName() string {
	return "site_store_create_account"
}

type Article_unread_notice struct {
	Id       int `orm:"id" json:"id"`
	NoticeId int `orm:"notice_id" json:"notice_id"`
	Uid      int `orm:"uid" json:"uid"`
	IsNew    int `orm:"is_new" json:"is_new"`
}

func (*Article_unread_notice) TableName() string {
	return "article_unread_notice"
}

type Ewei_shop_live struct {
	Id                  int     `orm:"id" json:"id"`
	Uniacid             int     `orm:"uniacid" json:"uniacid"`
	Merchid             int     `orm:"merchid" json:"merchid"`
	Title               string  `orm:"title" json:"title"`
	Livetype            int     `orm:"livetype" json:"livetype"`
	Liveidentity        string  `orm:"liveidentity" json:"liveidentity"`
	Screen              int     `orm:"screen" json:"screen"`
	Goodsid             string  `orm:"goodsid" json:"goodsid"`
	Category            int     `orm:"category" json:"category"`
	Url                 string  `orm:"url" json:"url"`
	Thumb               string  `orm:"thumb" json:"thumb"`
	Hot                 int     `orm:"hot" json:"hot"`
	Recommend           int     `orm:"recommend" json:"recommend"`
	Living              int     `orm:"living" json:"living"`
	Status              int     `orm:"status" json:"status"`
	Displayorder        int     `orm:"displayorder" json:"displayorder"`
	Livetime            int     `orm:"livetime" json:"livetime"`
	Lastlivetime        int     `orm:"lastlivetime" json:"lastlivetime"`
	Createtime          int     `orm:"createtime" json:"createtime"`
	Introduce           string  `orm:"introduce" json:"introduce"`
	Packetmoney         float64 `orm:"packetmoney" json:"packetmoney"`
	Packettotal         int     `orm:"packettotal" json:"packettotal"`
	Packetprice         float64 `orm:"packetprice" json:"packetprice"`
	Packetdes           string  `orm:"packetdes" json:"packetdes"`
	Couponid            string  `orm:"couponid" json:"couponid"`
	ShareTitle          string  `orm:"share_title" json:"share_title"`
	ShareIcon           string  `orm:"share_icon" json:"share_icon"`
	ShareDesc           string  `orm:"share_desc" json:"share_desc"`
	ShareUrl            string  `orm:"share_url" json:"share_url"`
	Subscribe           int     `orm:"subscribe" json:"subscribe"`
	Subscribenotice     int     `orm:"subscribenotice" json:"subscribenotice"`
	Visit               int     `orm:"visit" json:"visit"`
	Video               string  `orm:"video" json:"video"`
	Covertype           int     `orm:"covertype" json:"covertype"`
	Cover               string  `orm:"cover" json:"cover"`
	Iscoupon            int     `orm:"iscoupon" json:"iscoupon"`
	Nestable            string  `orm:"nestable" json:"nestable"`
	Tabs                string  `orm:"tabs" json:"tabs"`
	InvitationId        int     `orm:"invitation_id" json:"invitation_id"`
	Showlevels          string  `orm:"showlevels" json:"showlevels"`
	Showgroups          string  `orm:"showgroups" json:"showgroups"`
	Showcommission      string  `orm:"showcommission" json:"showcommission"`
	JurisdictionUrl     string  `orm:"jurisdiction_url" json:"jurisdiction_url"`
	JurisdictionurlShow int     `orm:"jurisdictionurl_show" json:"jurisdictionurl_show"`
	Notice              string  `orm:"notice" json:"notice"`
	NoticeUrl           string  `orm:"notice_url" json:"notice_url"`
	Followqrcode        string  `orm:"followqrcode" json:"followqrcode"`
	CouponNum           int     `orm:"coupon_num" json:"coupon_num"`
}

func (*Ewei_shop_live) TableName() string {
	return "ewei_shop_live"
}

type Ewei_shop_member struct {
	Id                  int     `orm:"id" json:"id"`
	Uniacid             int     `orm:"uniacid" json:"uniacid"`
	Uid                 int     `orm:"uid" json:"uid"`
	Groupid             int     `orm:"groupid" json:"groupid"`
	Level               int     `orm:"level" json:"level"`
	Agentid             int     `orm:"agentid" json:"agentid"`
	Openid              string  `orm:"openid" json:"openid"`
	Realname            string  `orm:"realname" json:"realname"`
	Mobile              string  `orm:"mobile" json:"mobile"`
	Pwd                 string  `orm:"pwd" json:"pwd"`
	Weixin              string  `orm:"weixin" json:"weixin"`
	Content             string  `orm:"content" json:"content"`
	Createtime          int     `orm:"createtime" json:"createtime"`
	Agenttime           int     `orm:"agenttime" json:"agenttime"`
	Status              int     `orm:"status" json:"status"`
	Isagent             int     `orm:"isagent" json:"isagent"`
	Clickcount          int     `orm:"clickcount" json:"clickcount"`
	Agentlevel          int     `orm:"agentlevel" json:"agentlevel"`
	Noticeset           string  `orm:"noticeset" json:"noticeset"`
	Nickname            string  `orm:"nickname" json:"nickname"`
	Credit1             float64 `orm:"credit1" json:"credit1"`
	Credit2             float64 `orm:"credit2" json:"credit2"`
	Birthyear           string  `orm:"birthyear" json:"birthyear"`
	Birthmonth          string  `orm:"birthmonth" json:"birthmonth"`
	Birthday            string  `orm:"birthday" json:"birthday"`
	Gender              int     `orm:"gender" json:"gender"`
	Avatar              string  `orm:"avatar" json:"avatar"`
	Province            string  `orm:"province" json:"province"`
	City                string  `orm:"city" json:"city"`
	Area                string  `orm:"area" json:"area"`
	Childtime           int     `orm:"childtime" json:"childtime"`
	Inviter             int     `orm:"inviter" json:"inviter"`
	Agentnotupgrade     int     `orm:"agentnotupgrade" json:"agentnotupgrade"`
	Agentselectgoods    int     `orm:"agentselectgoods" json:"agentselectgoods"`
	Agentblack          int     `orm:"agentblack" json:"agentblack"`
	Fixagentid          int     `orm:"fixagentid" json:"fixagentid"`
	Diymemberid         int     `orm:"diymemberid" json:"diymemberid"`
	Diymemberfields     string  `orm:"diymemberfields" json:"diymemberfields"`
	Diymemberdata       string  `orm:"diymemberdata" json:"diymemberdata"`
	Diymemberdataid     int     `orm:"diymemberdataid" json:"diymemberdataid"`
	Diycommissionid     int     `orm:"diycommissionid" json:"diycommissionid"`
	Diycommissionfields string  `orm:"diycommissionfields" json:"diycommissionfields"`
	Diycommissiondata   string  `orm:"diycommissiondata" json:"diycommissiondata"`
	Diycommissiondataid int     `orm:"diycommissiondataid" json:"diycommissiondataid"`
	Isblack             int     `orm:"isblack" json:"isblack"`
	Username            string  `orm:"username" json:"username"`
	CommissionTotal     float64 `orm:"commission_total" json:"commission_total"`
	Endtime2            int     `orm:"endtime2" json:"endtime2"`
	Ispartner           int     `orm:"ispartner" json:"ispartner"`
	Partnertime         int     `orm:"partnertime" json:"partnertime"`
	Partnerstatus       int     `orm:"partnerstatus" json:"partnerstatus"`
	Partnerblack        int     `orm:"partnerblack" json:"partnerblack"`
	Partnerlevel        int     `orm:"partnerlevel" json:"partnerlevel"`
	Partnernotupgrade   int     `orm:"partnernotupgrade" json:"partnernotupgrade"`
	Diyglobonusid       int     `orm:"diyglobonusid" json:"diyglobonusid"`
	Diyglobonusdata     string  `orm:"diyglobonusdata" json:"diyglobonusdata"`
	Diyglobonusfields   string  `orm:"diyglobonusfields" json:"diyglobonusfields"`
	Isaagent            int     `orm:"isaagent" json:"isaagent"`
	Aagentlevel         int     `orm:"aagentlevel" json:"aagentlevel"`
	Aagenttime          int     `orm:"aagenttime" json:"aagenttime"`
	Aagentstatus        int     `orm:"aagentstatus" json:"aagentstatus"`
	Aagentblack         int     `orm:"aagentblack" json:"aagentblack"`
	Aagentnotupgrade    int     `orm:"aagentnotupgrade" json:"aagentnotupgrade"`
	Aagenttype          int     `orm:"aagenttype" json:"aagenttype"`
	Aagentprovinces     string  `orm:"aagentprovinces" json:"aagentprovinces"`
	Aagentcitys         string  `orm:"aagentcitys" json:"aagentcitys"`
	Aagentareas         string  `orm:"aagentareas" json:"aagentareas"`
	Diyaagentid         int     `orm:"diyaagentid" json:"diyaagentid"`
	Diyaagentdata       string  `orm:"diyaagentdata" json:"diyaagentdata"`
	Diyaagentfields     string  `orm:"diyaagentfields" json:"diyaagentfields"`
	Salt                string  `orm:"salt" json:"salt"`
	Mobileverify        int     `orm:"mobileverify" json:"mobileverify"`
	Mobileuser          int     `orm:"mobileuser" json:"mobileuser"`
	CarrierMobile       string  `orm:"carrier_mobile" json:"carrier_mobile"`
	Isauthor            int     `orm:"isauthor" json:"isauthor"`
	Authortime          int     `orm:"authortime" json:"authortime"`
	Authorstatus        int     `orm:"authorstatus" json:"authorstatus"`
	Authorblack         int     `orm:"authorblack" json:"authorblack"`
	Authorlevel         int     `orm:"authorlevel" json:"authorlevel"`
	Authornotupgrade    int     `orm:"authornotupgrade" json:"authornotupgrade"`
	Diyauthorid         int     `orm:"diyauthorid" json:"diyauthorid"`
	Diyauthordata       string  `orm:"diyauthordata" json:"diyauthordata"`
	Diyauthorfields     string  `orm:"diyauthorfields" json:"diyauthorfields"`
	Authorid            int     `orm:"authorid" json:"authorid"`
	Comefrom            string  `orm:"comefrom" json:"comefrom"`
	OpenidQq            string  `orm:"openid_qq" json:"openid_qq"`
	OpenidWx            string  `orm:"openid_wx" json:"openid_wx"`
	Diymaxcredit        int     `orm:"diymaxcredit" json:"diymaxcredit"`
	Maxcredit           int     `orm:"maxcredit" json:"maxcredit"`
	Datavalue           string  `orm:"datavalue" json:"datavalue"`
	OpenidWa            string  `orm:"openid_wa" json:"openid_wa"`
	NicknameWechat      string  `orm:"nickname_wechat" json:"nickname_wechat"`
	AvatarWechat        string  `orm:"avatar_wechat" json:"avatar_wechat"`
	Updateaddress       int     `orm:"updateaddress" json:"updateaddress"`
	Membercardid        string  `orm:"membercardid" json:"membercardid"`
	Membercardcode      string  `orm:"membercardcode" json:"membercardcode"`
	Membershipnumber    string  `orm:"membershipnumber" json:"membershipnumber"`
	Membercardactive    int     `orm:"membercardactive" json:"membercardactive"`
	Commission          float64 `orm:"commission" json:"commission"`
	CommissionPay       float64 `orm:"commission_pay" json:"commission_pay"`
	Idnumber            string  `orm:"idnumber" json:"idnumber"`
	Wxcardupdatetime    int     `orm:"wxcardupdatetime" json:"wxcardupdatetime"`
}

func (*Ewei_shop_member) TableName() string {
	return "ewei_shop_member"
}

type Ewei_shop_order_peerpay struct {
	Id               int     `orm:"id" json:"id"`
	Uniacid          int     `orm:"uniacid" json:"uniacid"`
	Orderid          int     `orm:"orderid" json:"orderid"`
	PeerpayType      int     `orm:"peerpay_type" json:"peerpay_type"`
	PeerpayPrice     float64 `orm:"peerpay_price" json:"peerpay_price"`
	PeerpayMaxprice  float64 `orm:"peerpay_maxprice" json:"peerpay_maxprice"`
	PeerpayRealprice float64 `orm:"peerpay_realprice" json:"peerpay_realprice"`
	PeerpaySelfpay   float64 `orm:"peerpay_selfpay" json:"peerpay_selfpay"`
	PeerpayMessage   string  `orm:"peerpay_message" json:"peerpay_message"`
	Status           int     `orm:"status" json:"status"`
	Createtime       int     `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_order_peerpay) TableName() string {
	return "ewei_shop_order_peerpay"
}

type Ewei_shop_order_print struct {
	Id      int `orm:"id" json:"id"`
	Status  int `orm:"status" json:"status"`
	Sid     int `orm:"sid" json:"sid"`
	Foid    int `orm:"foid" json:"foid"`
	Oid     int `orm:"oid" json:"oid"`
	Pid     int `orm:"pid" json:"pid"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Addtime int `orm:"addtime" json:"addtime"`
}

func (*Ewei_shop_order_print) TableName() string {
	return "ewei_shop_order_print"
}

type Ewei_shop_system_case struct {
	Id           int    `orm:"id" json:"id"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Qr           string `orm:"qr" json:"qr"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Status       int    `orm:"status" json:"status"`
	Cate         int    `orm:"cate" json:"cate"`
	Description  string `orm:"description" json:"description"`
}

func (*Ewei_shop_system_case) TableName() string {
	return "ewei_shop_system_case"
}

type Ewei_shop_task_extension_join struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Title        string `orm:"title" json:"title"`
	Uid          int    `orm:"uid" json:"uid"`
	Taskid       int    `orm:"taskid" json:"taskid"`
	Openid       string `orm:"openid" json:"openid"`
	RequireData  string `orm:"require_data" json:"require_data"`
	ProgressData string `orm:"progress_data" json:"progress_data"`
	RewardData   string `orm:"reward_data" json:"reward_data"`
	Completetime int    `orm:"completetime" json:"completetime"`
	Pickuptime   int    `orm:"pickuptime" json:"pickuptime"`
	Endtime      int    `orm:"endtime" json:"endtime"`
	Dotime       int    `orm:"dotime" json:"dotime"`
	Rewarded     string `orm:"rewarded" json:"rewarded"`
	Logo         string `orm:"logo" json:"logo"`
}

func (*Ewei_shop_task_extension_join) TableName() string {
	return "ewei_shop_task_extension_join"
}

type Uni_group struct {
	Id        int    `orm:"id" json:"id"`
	OwnerUid  int    `orm:"owner_uid" json:"owner_uid"`
	Name      string `orm:"name" json:"name"`
	Modules   string `orm:"modules" json:"modules"`
	Templates string `orm:"templates" json:"templates"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
}

func (*Uni_group) TableName() string {
	return "uni_group"
}

type Userapi_reply struct {
	Id          int    `orm:"id" json:"id"`
	Rid         int    `orm:"rid" json:"rid"`
	Description string `orm:"description" json:"description"`
	Apiurl      string `orm:"apiurl" json:"apiurl"`
	Token       string `orm:"token" json:"token"`
	DefaultText string `orm:"default_text" json:"default_text"`
	Cachetime   int    `orm:"cachetime" json:"cachetime"`
}

func (*Userapi_reply) TableName() string {
	return "userapi_reply"
}

type Coupon struct {
	Id                   int    `orm:"id" json:"id"`
	Uniacid              int    `orm:"uniacid" json:"uniacid"`
	Acid                 int    `orm:"acid" json:"acid"`
	CardId               string `orm:"card_id" json:"card_id"`
	Type                 string `orm:"type" json:"type"`
	LogoUrl              string `orm:"logo_url" json:"logo_url"`
	CodeType             int    `orm:"code_type" json:"code_type"`
	BrandName            string `orm:"brand_name" json:"brand_name"`
	Title                string `orm:"title" json:"title"`
	SubTitle             string `orm:"sub_title" json:"sub_title"`
	Color                string `orm:"color" json:"color"`
	Notice               string `orm:"notice" json:"notice"`
	Description          string `orm:"description" json:"description"`
	DateInfo             string `orm:"date_info" json:"date_info"`
	Quantity             int    `orm:"quantity" json:"quantity"`
	UseCustomCode        int    `orm:"use_custom_code" json:"use_custom_code"`
	BindOpenid           int    `orm:"bind_openid" json:"bind_openid"`
	CanShare             int    `orm:"can_share" json:"can_share"`
	CanGiveFriend        int    `orm:"can_give_friend" json:"can_give_friend"`
	GetLimit             int    `orm:"get_limit" json:"get_limit"`
	ServicePhone         string `orm:"service_phone" json:"service_phone"`
	Extra                string `orm:"extra" json:"extra"`
	Status               int    `orm:"status" json:"status"`
	IsDisplay            int    `orm:"is_display" json:"is_display"`
	IsSelfconsume        int    `orm:"is_selfconsume" json:"is_selfconsume"`
	PromotionUrlName     string `orm:"promotion_url_name" json:"promotion_url_name"`
	PromotionUrl         string `orm:"promotion_url" json:"promotion_url"`
	PromotionUrlSubTitle string `orm:"promotion_url_sub_title" json:"promotion_url_sub_title"`
	Source               int    `orm:"source" json:"source"`
	Dosage               int    `orm:"dosage" json:"dosage"`
}

func (*Coupon) TableName() string {
	return "coupon"
}

type Ewei_shop_coupon_data struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Couponid   int    `orm:"couponid" json:"couponid"`
	Gettype    int    `orm:"gettype" json:"gettype"`
	Used       int    `orm:"used" json:"used"`
	Usetime    int    `orm:"usetime" json:"usetime"`
	Gettime    int    `orm:"gettime" json:"gettime"`
	Senduid    int    `orm:"senduid" json:"senduid"`
	Ordersn    string `orm:"ordersn" json:"ordersn"`
	Back       int    `orm:"back" json:"back"`
	Backtime   int    `orm:"backtime" json:"backtime"`
	Merchid    int    `orm:"merchid" json:"merchid"`
	Isnew      int    `orm:"isnew" json:"isnew"`
	Nocount    int    `orm:"nocount" json:"nocount"`
	Shareident string `orm:"shareident" json:"shareident"`
	Textkey    int    `orm:"textkey" json:"textkey"`
}

func (*Ewei_shop_coupon_data) TableName() string {
	return "ewei_shop_coupon_data"
}

type Ewei_shop_creditshop_log struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Logno           string  `orm:"logno" json:"logno"`
	Eno             string  `orm:"eno" json:"eno"`
	Openid          string  `orm:"openid" json:"openid"`
	Goodsid         int     `orm:"goodsid" json:"goodsid"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Status          int     `orm:"status" json:"status"`
	Paystatus       int     `orm:"paystatus" json:"paystatus"`
	Paytype         int     `orm:"paytype" json:"paytype"`
	Dispatchstatus  int     `orm:"dispatchstatus" json:"dispatchstatus"`
	Creditpay       int     `orm:"creditpay" json:"creditpay"`
	Addressid       int     `orm:"addressid" json:"addressid"`
	Dispatchno      string  `orm:"dispatchno" json:"dispatchno"`
	Usetime         int     `orm:"usetime" json:"usetime"`
	Express         string  `orm:"express" json:"express"`
	Expresssn       string  `orm:"expresssn" json:"expresssn"`
	Expresscom      string  `orm:"expresscom" json:"expresscom"`
	Verifyopenid    string  `orm:"verifyopenid" json:"verifyopenid"`
	Storeid         int     `orm:"storeid" json:"storeid"`
	Realname        string  `orm:"realname" json:"realname"`
	Mobile          string  `orm:"mobile" json:"mobile"`
	Couponid        int     `orm:"couponid" json:"couponid"`
	Dupdate1        int     `orm:"dupdate1" json:"dupdate1"`
	Transid         string  `orm:"transid" json:"transid"`
	Dispatchtransid string  `orm:"dispatchtransid" json:"dispatchtransid"`
	Address         string  `orm:"address" json:"address"`
	Optionid        int     `orm:"optionid" json:"optionid"`
	TimeSend        int     `orm:"time_send" json:"time_send"`
	TimeFinish      int     `orm:"time_finish" json:"time_finish"`
	Iscomment       int     `orm:"iscomment" json:"iscomment"`
	Dispatchtime    int     `orm:"dispatchtime" json:"dispatchtime"`
	Verifynum       int     `orm:"verifynum" json:"verifynum"`
	Verifytime      int     `orm:"verifytime" json:"verifytime"`
	Merchid         int     `orm:"merchid" json:"merchid"`
	Remarksaler     string  `orm:"remarksaler" json:"remarksaler"`
	Dispatch        float64 `orm:"dispatch" json:"dispatch"`
	Money           float64 `orm:"money" json:"money"`
	Credit          int     `orm:"credit" json:"credit"`
	GoodsNum        int     `orm:"goods_num" json:"goods_num"`
}

func (*Ewei_shop_creditshop_log) TableName() string {
	return "ewei_shop_creditshop_log"
}

type Ewei_shop_groups_set struct {
	Id                int     `orm:"id" json:"id"`
	Uniacid           string  `orm:"uniacid" json:"uniacid"`
	Groups            int     `orm:"groups" json:"groups"`
	Followurl         string  `orm:"followurl" json:"followurl"`
	Groupsurl         string  `orm:"groupsurl" json:"groupsurl"`
	ShareTitle        string  `orm:"share_title" json:"share_title"`
	ShareIcon         string  `orm:"share_icon" json:"share_icon"`
	ShareDesc         string  `orm:"share_desc" json:"share_desc"`
	GroupsDescription string  `orm:"groups_description" json:"groups_description"`
	Description       int     `orm:"description" json:"description"`
	Followqrcode      string  `orm:"followqrcode" json:"followqrcode"`
	ShareUrl          string  `orm:"share_url" json:"share_url"`
	Creditdeduct      int     `orm:"creditdeduct" json:"creditdeduct"`
	Groupsdeduct      int     `orm:"groupsdeduct" json:"groupsdeduct"`
	Credit            int     `orm:"credit" json:"credit"`
	Groupsmoney       float64 `orm:"groupsmoney" json:"groupsmoney"`
	Refund            int     `orm:"refund" json:"refund"`
	Refundday         int     `orm:"refundday" json:"refundday"`
	Goodsid           string  `orm:"goodsid" json:"goodsid"`
	Rules             string  `orm:"rules" json:"rules"`
	Receive           int     `orm:"receive" json:"receive"`
	Discount          int     `orm:"discount" json:"discount"`
	Headstype         int     `orm:"headstype" json:"headstype"`
	Headsmoney        float64 `orm:"headsmoney" json:"headsmoney"`
	Headsdiscount     int     `orm:"headsdiscount" json:"headsdiscount"`
}

func (*Ewei_shop_groups_set) TableName() string {
	return "ewei_shop_groups_set"
}

type Ewei_shop_order_peerpay_payinfo struct {
	Id           int     `orm:"id" json:"id"`
	Pid          int     `orm:"pid" json:"pid"`
	Uid          int     `orm:"uid" json:"uid"`
	Uname        string  `orm:"uname" json:"uname"`
	Usay         string  `orm:"usay" json:"usay"`
	Price        float64 `orm:"price" json:"price"`
	Createtime   int     `orm:"createtime" json:"createtime"`
	Headimg      string  `orm:"headimg" json:"headimg"`
	Refundstatus int     `orm:"refundstatus" json:"refundstatus"`
	Refundprice  float64 `orm:"refundprice" json:"refundprice"`
	Openid       string  `orm:"openid" json:"openid"`
	Tid          string  `orm:"tid" json:"tid"`
}

func (*Ewei_shop_order_peerpay_payinfo) TableName() string {
	return "ewei_shop_order_peerpay_payinfo"
}

type Ewei_shop_order_refund struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Orderid         int     `orm:"orderid" json:"orderid"`
	Refundno        string  `orm:"refundno" json:"refundno"`
	Price           string  `orm:"price" json:"price"`
	Reason          string  `orm:"reason" json:"reason"`
	Images          string  `orm:"images" json:"images"`
	Content         string  `orm:"content" json:"content"`
	Createtime      int     `orm:"createtime" json:"createtime"`
	Status          int     `orm:"status" json:"status"`
	Reply           string  `orm:"reply" json:"reply"`
	Refundtype      int     `orm:"refundtype" json:"refundtype"`
	Orderprice      float64 `orm:"orderprice" json:"orderprice"`
	Applyprice      float64 `orm:"applyprice" json:"applyprice"`
	Imgs            string  `orm:"imgs" json:"imgs"`
	Rtype           int     `orm:"rtype" json:"rtype"`
	Refundaddress   string  `orm:"refundaddress" json:"refundaddress"`
	Message         string  `orm:"message" json:"message"`
	Express         string  `orm:"express" json:"express"`
	Expresscom      string  `orm:"expresscom" json:"expresscom"`
	Expresssn       string  `orm:"expresssn" json:"expresssn"`
	Operatetime     int     `orm:"operatetime" json:"operatetime"`
	Sendtime        int     `orm:"sendtime" json:"sendtime"`
	Returntime      int     `orm:"returntime" json:"returntime"`
	Refundtime      int     `orm:"refundtime" json:"refundtime"`
	Rexpress        string  `orm:"rexpress" json:"rexpress"`
	Rexpresscom     string  `orm:"rexpresscom" json:"rexpresscom"`
	Rexpresssn      string  `orm:"rexpresssn" json:"rexpresssn"`
	Refundaddressid int     `orm:"refundaddressid" json:"refundaddressid"`
	Endtime         int     `orm:"endtime" json:"endtime"`
	Realprice       float64 `orm:"realprice" json:"realprice"`
	Merchid         int     `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_order_refund) TableName() string {
	return "ewei_shop_order_refund"
}

type Ewei_shop_sysset struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Sets    string `orm:"sets" json:"sets"`
	Plugins string `orm:"plugins" json:"plugins"`
	Sec     string `orm:"sec" json:"sec"`
}

func (*Ewei_shop_sysset) TableName() string {
	return "ewei_shop_sysset"
}

type Users struct {
	Uid            int    `orm:"uid" json:"uid"`
	OwnerUid       int    `orm:"owner_uid" json:"owner_uid"`
	Groupid        int    `orm:"groupid" json:"groupid"`
	FounderGroupid int    `orm:"founder_groupid" json:"founder_groupid"`
	Username       string `orm:"username" json:"username"`
	Password       string `orm:"password" json:"password"`
	Salt           string `orm:"salt" json:"salt"`
	Type           int    `orm:"type" json:"type"`
	Status         int    `orm:"status" json:"status"`
	Joindate       int    `orm:"joindate" json:"joindate"`
	Joinip         string `orm:"joinip" json:"joinip"`
	Lastvisit      int    `orm:"lastvisit" json:"lastvisit"`
	Lastip         string `orm:"lastip" json:"lastip"`
	Remark         string `orm:"remark" json:"remark"`
	Starttime      int    `orm:"starttime" json:"starttime"`
	Endtime        int    `orm:"endtime" json:"endtime"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	Tid            int    `orm:"tid" json:"tid"`
	Schoolid       int    `orm:"schoolid" json:"schoolid"`
	RegisterType   int    `orm:"register_type" json:"register_type"`
	Openid         string `orm:"openid" json:"openid"`
}

func (*Users) TableName() string {
	return "users"
}

type Account struct {
	Acid      int    `orm:"acid" json:"acid"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Hash      string `orm:"hash" json:"hash"`
	Type      int    `orm:"type" json:"type"`
	Isconnect int    `orm:"isconnect" json:"isconnect"`
	Isdeleted int    `orm:"isdeleted" json:"isdeleted"`
	Endtime   int    `orm:"endtime" json:"endtime"`
}

func (*Account) TableName() string {
	return "account"
}

type Ewei_shop_abonus_billo struct {
	Id         int     `orm:"id" json:"id"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Billid     int     `orm:"billid" json:"billid"`
	Orderid    int     `orm:"orderid" json:"orderid"`
	Ordermoney float64 `orm:"ordermoney" json:"ordermoney"`
}

func (*Ewei_shop_abonus_billo) TableName() string {
	return "ewei_shop_abonus_billo"
}

type Ewei_shop_bargain_goods struct {
	Id          int     `orm:"id" json:"id"`
	AccountId   int     `orm:"account_id" json:"account_id"`
	GoodsId     string  `orm:"goods_id" json:"goods_id"`
	EndPrice    float64 `orm:"end_price" json:"end_price"`
	StartTime   string  `orm:"start_time" json:"start_time"`
	EndTime     string  `orm:"end_time" json:"end_time"`
	Status      int     `orm:"status" json:"status"`
	Type        int     `orm:"type" json:"type"`
	UserSet     string  `orm:"user_set" json:"user_set"`
	Rule        string  `orm:"rule" json:"rule"`
	ActTimes    int     `orm:"act_times" json:"act_times"`
	Mode        int     `orm:"mode" json:"mode"`
	TotalTime   int     `orm:"total_time" json:"total_time"`
	EachTime    int     `orm:"each_time" json:"each_time"`
	TimeLimit   int     `orm:"time_limit" json:"time_limit"`
	Probability string  `orm:"probability" json:"probability"`
	Custom      string  `orm:"custom" json:"custom"`
	Maximum     int     `orm:"maximum" json:"maximum"`
	Initiate    int     `orm:"initiate" json:"initiate"`
	Myself      int     `orm:"myself" json:"myself"`
}

func (*Ewei_shop_bargain_goods) TableName() string {
	return "ewei_shop_bargain_goods"
}

type Ewei_shop_creditshop_spec_item struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Specid       int    `orm:"specid" json:"specid"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Show         int    `orm:"show" json:"show"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	ValueId      string `orm:"valueId" json:"valueId"`
	Virtual      int    `orm:"virtual" json:"virtual"`
}

func (*Ewei_shop_creditshop_spec_item) TableName() string {
	return "ewei_shop_creditshop_spec_item"
}

type Ewei_shop_member_group_log struct {
	LogId   int    `orm:"log_id" json:"log_id"`
	Mid     int    `orm:"mid" json:"mid"`
	Openid  string `orm:"openid" json:"openid"`
	GroupId int    `orm:"group_id" json:"group_id"`
	AddTime string `orm:"add_time" json:"add_time"`
	Content string `orm:"content" json:"content"`
}

func (*Ewei_shop_member_group_log) TableName() string {
	return "ewei_shop_member_group_log"
}

type Ewei_shop_member_printer_template struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Title      string `orm:"title" json:"title"`
	Type       int    `orm:"type" json:"type"`
	PrintTitle string `orm:"print_title" json:"print_title"`
	PrintStyle string `orm:"print_style" json:"print_style"`
	PrintData  string `orm:"print_data" json:"print_data"`
	Code       string `orm:"code" json:"code"`
	Qrcode     string `orm:"qrcode" json:"qrcode"`
	Createtime int    `orm:"createtime" json:"createtime"`
	Merchid    int    `orm:"merchid" json:"merchid"`
	Goodssn    int    `orm:"goodssn" json:"goodssn"`
	Productsn  int    `orm:"productsn" json:"productsn"`
}

func (*Ewei_shop_member_printer_template) TableName() string {
	return "ewei_shop_member_printer_template"
}

type Ewei_shop_multi_shop struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Uid           int    `orm:"uid" json:"uid"`
	Name          string `orm:"name" json:"name"`
	Company       string `orm:"company" json:"company"`
	Sales         string `orm:"sales" json:"sales"`
	Starttime     int    `orm:"starttime" json:"starttime"`
	Endtime       int    `orm:"endtime" json:"endtime"`
	Applytime     int    `orm:"applytime" json:"applytime"`
	Jointime      int    `orm:"jointime" json:"jointime"`
	Status        int    `orm:"status" json:"status"`
	Refusecontent string `orm:"refusecontent" json:"refusecontent"`
}

func (*Ewei_shop_multi_shop) TableName() string {
	return "ewei_shop_multi_shop"
}

type Ewei_shop_sms_set struct {
	Id                 int     `orm:"id" json:"id"`
	Uniacid            int     `orm:"uniacid" json:"uniacid"`
	Juhe               int     `orm:"juhe" json:"juhe"`
	JuheKey            string  `orm:"juhe_key" json:"juhe_key"`
	Dayu               int     `orm:"dayu" json:"dayu"`
	DayuKey            string  `orm:"dayu_key" json:"dayu_key"`
	DayuSecret         string  `orm:"dayu_secret" json:"dayu_secret"`
	Emay               int     `orm:"emay" json:"emay"`
	EmayUrl            string  `orm:"emay_url" json:"emay_url"`
	EmaySn             string  `orm:"emay_sn" json:"emay_sn"`
	EmayPw             string  `orm:"emay_pw" json:"emay_pw"`
	EmaySk             string  `orm:"emay_sk" json:"emay_sk"`
	EmayPhost          string  `orm:"emay_phost" json:"emay_phost"`
	EmayPport          int     `orm:"emay_pport" json:"emay_pport"`
	EmayPuser          string  `orm:"emay_puser" json:"emay_puser"`
	EmayPpw            string  `orm:"emay_ppw" json:"emay_ppw"`
	EmayOut            int     `orm:"emay_out" json:"emay_out"`
	EmayOutresp        int     `orm:"emay_outresp" json:"emay_outresp"`
	EmayWarn           float64 `orm:"emay_warn" json:"emay_warn"`
	EmayMobile         string  `orm:"emay_mobile" json:"emay_mobile"`
	EmayWarnTime       int     `orm:"emay_warn_time" json:"emay_warn_time"`
	Aliyun             int     `orm:"aliyun" json:"aliyun"`
	AliyunAppcode      string  `orm:"aliyun_appcode" json:"aliyun_appcode"`
	AliyunNew          int     `orm:"aliyun_new" json:"aliyun_new"`
	AliyunNewKeyid     string  `orm:"aliyun_new_keyid" json:"aliyun_new_keyid"`
	AliyunNewKeysecret string  `orm:"aliyun_new_keysecret" json:"aliyun_new_keysecret"`
}

func (*Ewei_shop_sms_set) TableName() string {
	return "ewei_shop_sms_set"
}

type Ewei_shop_system_plugingrant_order struct {
	Id         int     `orm:"id" json:"id"`
	Logno      string  `orm:"logno" json:"logno"`
	Code       string  `orm:"code" json:"code"`
	Uniacid    int     `orm:"uniacid" json:"uniacid"`
	Username   string  `orm:"username" json:"username"`
	Pluginid   string  `orm:"pluginid" json:"pluginid"`
	Price      float64 `orm:"price" json:"price"`
	Month      int     `orm:"month" json:"month"`
	Createtime int     `orm:"createtime" json:"createtime"`
	Paystatus  int     `orm:"paystatus" json:"paystatus"`
	Paytime    int     `orm:"paytime" json:"paytime"`
	Paytype    int     `orm:"paytype" json:"paytype"`
}

func (*Ewei_shop_system_plugingrant_order) TableName() string {
	return "ewei_shop_system_plugingrant_order"
}

type Ewei_shop_task_poster struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Days         int    `orm:"days" json:"days"`
	Title        string `orm:"title" json:"title"`
	Bg           string `orm:"bg" json:"bg"`
	Data         string `orm:"data" json:"data"`
	Keyword      string `orm:"keyword" json:"keyword"`
	Resptype     int    `orm:"resptype" json:"resptype"`
	Resptext     string `orm:"resptext" json:"resptext"`
	Resptitle    string `orm:"resptitle" json:"resptitle"`
	Respthumb    string `orm:"respthumb" json:"respthumb"`
	Respdesc     string `orm:"respdesc" json:"respdesc"`
	Respurl      string `orm:"respurl" json:"respurl"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Waittext     string `orm:"waittext" json:"waittext"`
	Oktext       string `orm:"oktext" json:"oktext"`
	Scantext     string `orm:"scantext" json:"scantext"`
	Beagent      int    `orm:"beagent" json:"beagent"`
	Bedown       int    `orm:"bedown" json:"bedown"`
	Timestart    int    `orm:"timestart" json:"timestart"`
	Timeend      int    `orm:"timeend" json:"timeend"`
	IsRepeat     int    `orm:"is_repeat" json:"is_repeat"`
	Getposter    string `orm:"getposter" json:"getposter"`
	Status       int    `orm:"status" json:"status"`
	Starttext    string `orm:"starttext" json:"starttext"`
	Endtext      string `orm:"endtext" json:"endtext"`
	RewardData   string `orm:"reward_data" json:"reward_data"`
	Needcount    int    `orm:"needcount" json:"needcount"`
	IsDelete     int    `orm:"is_delete" json:"is_delete"`
	PosterType   int    `orm:"poster_type" json:"poster_type"`
	RewardDays   int    `orm:"reward_days" json:"reward_days"`
	Titleicon    string `orm:"titleicon" json:"titleicon"`
	PosterBanner string `orm:"poster_banner" json:"poster_banner"`
	IsGoods      int    `orm:"is_goods" json:"is_goods"`
	Autoposter   int    `orm:"autoposter" json:"autoposter"`
}

func (*Ewei_shop_task_poster) TableName() string {
	return "ewei_shop_task_poster"
}

type Ewei_shop_creditshop_verify struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Logid      int    `orm:"logid" json:"logid"`
	Verifycode string `orm:"verifycode" json:"verifycode"`
	Storeid    int    `orm:"storeid" json:"storeid"`
	Verifier   string `orm:"verifier" json:"verifier"`
	Isverify   int    `orm:"isverify" json:"isverify"`
	Verifytime int    `orm:"verifytime" json:"verifytime"`
	Merchid    int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_creditshop_verify) TableName() string {
	return "ewei_shop_creditshop_verify"
}

type Ewei_shop_diyform_category struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
}

func (*Ewei_shop_diyform_category) TableName() string {
	return "ewei_shop_diyform_category"
}

type Ewei_shop_goods_group struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Name     string `orm:"name" json:"name"`
	Goodsids string `orm:"goodsids" json:"goodsids"`
	Enabled  int    `orm:"enabled" json:"enabled"`
	Merchid  int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_goods_group) TableName() string {
	return "ewei_shop_goods_group"
}

type Ewei_shop_merch_banner struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Bannername   string `orm:"bannername" json:"bannername"`
	Link         string `orm:"link" json:"link"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Merchid      int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_merch_banner) TableName() string {
	return "ewei_shop_merch_banner"
}

type Ewei_shop_merch_group struct {
	Id                 int    `orm:"id" json:"id"`
	Uniacid            int    `orm:"uniacid" json:"uniacid"`
	Groupname          string `orm:"groupname" json:"groupname"`
	Createtime         int    `orm:"createtime" json:"createtime"`
	Status             int    `orm:"status" json:"status"`
	Isdefault          int    `orm:"isdefault" json:"isdefault"`
	Goodschecked       int    `orm:"goodschecked" json:"goodschecked"`
	Commissionchecked  int    `orm:"commissionchecked" json:"commissionchecked"`
	Changepricechecked int    `orm:"changepricechecked" json:"changepricechecked"`
	Finishchecked      int    `orm:"finishchecked" json:"finishchecked"`
}

func (*Ewei_shop_merch_group) TableName() string {
	return "ewei_shop_merch_group"
}

type Ewei_shop_sign_set struct {
	Id                 int    `orm:"id" json:"id"`
	Uniacid            int    `orm:"uniacid" json:"uniacid"`
	Iscenter           int    `orm:"iscenter" json:"iscenter"`
	Iscreditshop       int    `orm:"iscreditshop" json:"iscreditshop"`
	Keyword            string `orm:"keyword" json:"keyword"`
	Title              string `orm:"title" json:"title"`
	Thumb              string `orm:"thumb" json:"thumb"`
	Desc               string `orm:"desc" json:"desc"`
	Isopen             int    `orm:"isopen" json:"isopen"`
	Signold            int    `orm:"signold" json:"signold"`
	SignoldPrice       int    `orm:"signold_price" json:"signold_price"`
	SignoldType        int    `orm:"signold_type" json:"signold_type"`
	Textsign           string `orm:"textsign" json:"textsign"`
	Textsignold        string `orm:"textsignold" json:"textsignold"`
	Textsigned         string `orm:"textsigned" json:"textsigned"`
	Textsignforget     string `orm:"textsignforget" json:"textsignforget"`
	Maincolor          string `orm:"maincolor" json:"maincolor"`
	Cycle              int    `orm:"cycle" json:"cycle"`
	RewardDefaultFirst int    `orm:"reward_default_first" json:"reward_default_first"`
	RewardDefaultDay   int    `orm:"reward_default_day" json:"reward_default_day"`
	RewordOrder        string `orm:"reword_order" json:"reword_order"`
	RewordSum          string `orm:"reword_sum" json:"reword_sum"`
	RewordSpecial      string `orm:"reword_special" json:"reword_special"`
	SignRule           string `orm:"sign_rule" json:"sign_rule"`
	Share              int    `orm:"share" json:"share"`
}

func (*Ewei_shop_sign_set) TableName() string {
	return "ewei_shop_sign_set"
}

type Ewei_shop_task_type struct {
	Id          int    `orm:"id" json:"id"`
	TypeKey     string `orm:"type_key" json:"type_key"`
	TypeName    string `orm:"type_name" json:"type_name"`
	Description string `orm:"description" json:"description"`
	Verb        string `orm:"verb" json:"verb"`
	Numeric     int    `orm:"numeric" json:"numeric"`
	Unit        string `orm:"unit" json:"unit"`
	Goods       int    `orm:"goods" json:"goods"`
	Theme       string `orm:"theme" json:"theme"`
	Once        int    `orm:"once" json:"once"`
}

func (*Ewei_shop_task_type) TableName() string {
	return "ewei_shop_task_type"
}

type Users_group struct {
	Id            int    `orm:"id" json:"id"`
	OwnerUid      int    `orm:"owner_uid" json:"owner_uid"`
	Name          string `orm:"name" json:"name"`
	Package       string `orm:"package" json:"package"`
	Maxaccount    int    `orm:"maxaccount" json:"maxaccount"`
	Maxsubaccount int    `orm:"maxsubaccount" json:"maxsubaccount"`
	Timelimit     int    `orm:"timelimit" json:"timelimit"`
	Maxwxapp      int    `orm:"maxwxapp" json:"maxwxapp"`
	Maxwebapp     int    `orm:"maxwebapp" json:"maxwebapp"`
}

func (*Users_group) TableName() string {
	return "users_group"
}

type Ewei_shop_creditshop_comment struct {
	Id                 int    `orm:"id" json:"id"`
	Uniacid            int    `orm:"uniacid" json:"uniacid"`
	Logid              int    `orm:"logid" json:"logid"`
	Logno              string `orm:"logno" json:"logno"`
	Goodsid            int    `orm:"goodsid" json:"goodsid"`
	Openid             string `orm:"openid" json:"openid"`
	Nickname           string `orm:"nickname" json:"nickname"`
	Headimg            string `orm:"headimg" json:"headimg"`
	Level              int    `orm:"level" json:"level"`
	Content            string `orm:"content" json:"content"`
	Images             string `orm:"images" json:"images"`
	Time               int    `orm:"time" json:"time"`
	ReplyContent       string `orm:"reply_content" json:"reply_content"`
	ReplyImages        string `orm:"reply_images" json:"reply_images"`
	ReplyTime          int    `orm:"reply_time" json:"reply_time"`
	AppendContent      string `orm:"append_content" json:"append_content"`
	AppendImages       string `orm:"append_images" json:"append_images"`
	AppendTime         int    `orm:"append_time" json:"append_time"`
	AppendReplyContent string `orm:"append_reply_content" json:"append_reply_content"`
	AppendReplyImages  string `orm:"append_reply_images" json:"append_reply_images"`
	AppendReplyTime    int    `orm:"append_reply_time" json:"append_reply_time"`
	Istop              int    `orm:"istop" json:"istop"`
	Checked            int    `orm:"checked" json:"checked"`
	AppendChecked      int    `orm:"append_checked" json:"append_checked"`
	Virtual            int    `orm:"virtual" json:"virtual"`
	Deleted            int    `orm:"deleted" json:"deleted"`
	Merchid            int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_creditshop_comment) TableName() string {
	return "ewei_shop_creditshop_comment"
}

type Ewei_shop_goods_labelstyle struct {
	Id      int `orm:"id" json:"id"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Style   int `orm:"style" json:"style"`
}

func (*Ewei_shop_goods_labelstyle) TableName() string {
	return "ewei_shop_goods_labelstyle"
}

type Ewei_shop_invitation_qr struct {
	Id           int    `orm:"id" json:"id"`
	Acid         int    `orm:"acid" json:"acid"`
	Openid       string `orm:"openid" json:"openid"`
	Invitationid int    `orm:"invitationid" json:"invitationid"`
	Roomid       int    `orm:"roomid" json:"roomid"`
	Sceneid      int    `orm:"sceneid" json:"sceneid"`
	Ticket       string `orm:"ticket" json:"ticket"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Expire       int    `orm:"expire" json:"expire"`
	Qrimg        string `orm:"qrimg" json:"qrimg"`
}

func (*Ewei_shop_invitation_qr) TableName() string {
	return "ewei_shop_invitation_qr"
}

type Ewei_shop_member_address struct {
	Id              int    `orm:"id" json:"id"`
	Uniacid         int    `orm:"uniacid" json:"uniacid"`
	Openid          string `orm:"openid" json:"openid"`
	Realname        string `orm:"realname" json:"realname"`
	Mobile          string `orm:"mobile" json:"mobile"`
	Province        string `orm:"province" json:"province"`
	City            string `orm:"city" json:"city"`
	Area            string `orm:"area" json:"area"`
	Address         string `orm:"address" json:"address"`
	Isdefault       int    `orm:"isdefault" json:"isdefault"`
	Zipcode         string `orm:"zipcode" json:"zipcode"`
	Deleted         int    `orm:"deleted" json:"deleted"`
	Street          string `orm:"street" json:"street"`
	Datavalue       string `orm:"datavalue" json:"datavalue"`
	Streetdatavalue string `orm:"streetdatavalue" json:"streetdatavalue"`
	Lng             string `orm:"lng" json:"lng"`
	Lat             string `orm:"lat" json:"lat"`
}

func (*Ewei_shop_member_address) TableName() string {
	return "ewei_shop_member_address"
}

type Ewei_shop_mc_merchant struct {
	Id             int    `orm:"id" json:"id"`
	Uniacid        int    `orm:"uniacid" json:"uniacid"`
	MerchantNo     string `orm:"merchant_no" json:"merchant_no"`
	Username       string `orm:"username" json:"username"`
	Password       string `orm:"password" json:"password"`
	Salt           string `orm:"salt" json:"salt"`
	ContactName    string `orm:"contact_name" json:"contact_name"`
	ContactMobile  string `orm:"contact_mobile" json:"contact_mobile"`
	ContactAddress string `orm:"contact_address" json:"contact_address"`
	Type           int    `orm:"type" json:"type"`
	Status         int    `orm:"status" json:"status"`
	Createtime     int    `orm:"createtime" json:"createtime"`
	Validitytime   int    `orm:"validitytime" json:"validitytime"`
	Industry       string `orm:"industry" json:"industry"`
	Remark         string `orm:"remark" json:"remark"`
}

func (*Ewei_shop_mc_merchant) TableName() string {
	return "ewei_shop_mc_merchant"
}

type Ewei_shop_commission_rank struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Type    int    `orm:"type" json:"type"`
	Num     int    `orm:"num" json:"num"`
	Status  int    `orm:"status" json:"status"`
	Title   string `orm:"title" json:"title"`
	Content string `orm:"content" json:"content"`
}

func (*Ewei_shop_commission_rank) TableName() string {
	return "ewei_shop_commission_rank"
}

type Ewei_shop_exchange_setting struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Freeze     int    `orm:"freeze" json:"freeze"`
	Mistake    int    `orm:"mistake" json:"mistake"`
	Grouplimit int    `orm:"grouplimit" json:"grouplimit"`
	Alllimit   int    `orm:"alllimit" json:"alllimit"`
	NoQrimg    int    `orm:"no_qrimg" json:"no_qrimg"`
	Rule       string `orm:"rule" json:"rule"`
}

func (*Ewei_shop_exchange_setting) TableName() string {
	return "ewei_shop_exchange_setting"
}

type Ewei_shop_groups_order struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Openid       string  `orm:"openid" json:"openid"`
	Orderno      string  `orm:"orderno" json:"orderno"`
	Groupnum     int     `orm:"groupnum" json:"groupnum"`
	Paytime      int     `orm:"paytime" json:"paytime"`
	Price        float64 `orm:"price" json:"price"`
	Freight      float64 `orm:"freight" json:"freight"`
	Status       int     `orm:"status" json:"status"`
	PayType      string  `orm:"pay_type" json:"pay_type"`
	Goodid       int     `orm:"goodid" json:"goodid"`
	Teamid       int     `orm:"teamid" json:"teamid"`
	IsTeam       int     `orm:"is_team" json:"is_team"`
	Heads        int     `orm:"heads" json:"heads"`
	Starttime    int     `orm:"starttime" json:"starttime"`
	Endtime      int     `orm:"endtime" json:"endtime"`
	Createtime   int     `orm:"createtime" json:"createtime"`
	Success      int     `orm:"success" json:"success"`
	Delete       int     `orm:"delete" json:"delete"`
	Credit       int     `orm:"credit" json:"credit"`
	Creditmoney  float64 `orm:"creditmoney" json:"creditmoney"`
	Dispatchid   int     `orm:"dispatchid" json:"dispatchid"`
	Addressid    int     `orm:"addressid" json:"addressid"`
	Address      string  `orm:"address" json:"address"`
	Discount     float64 `orm:"discount" json:"discount"`
	Canceltime   int     `orm:"canceltime" json:"canceltime"`
	Finishtime   int     `orm:"finishtime" json:"finishtime"`
	Refundid     int     `orm:"refundid" json:"refundid"`
	Refundstate  int     `orm:"refundstate" json:"refundstate"`
	Refundtime   int     `orm:"refundtime" json:"refundtime"`
	Express      string  `orm:"express" json:"express"`
	Expresscom   string  `orm:"expresscom" json:"expresscom"`
	Expresssn    string  `orm:"expresssn" json:"expresssn"`
	Sendtime     int     `orm:"sendtime" json:"sendtime"`
	Remark       string  `orm:"remark" json:"remark"`
	Remarkclose  string  `orm:"remarkclose" json:"remarkclose"`
	Remarksend   string  `orm:"remarksend" json:"remarksend"`
	Message      string  `orm:"message" json:"message"`
	Deleted      int     `orm:"deleted" json:"deleted"`
	Realname     string  `orm:"realname" json:"realname"`
	Mobile       string  `orm:"mobile" json:"mobile"`
	Isverify     int     `orm:"isverify" json:"isverify"`
	Verifytype   int     `orm:"verifytype" json:"verifytype"`
	Verifycode   string  `orm:"verifycode" json:"verifycode"`
	Verifynum    int     `orm:"verifynum" json:"verifynum"`
	Printstate   int     `orm:"printstate" json:"printstate"`
	Printstate2  int     `orm:"printstate2" json:"printstate2"`
	Apppay       int     `orm:"apppay" json:"apppay"`
	Isborrow     int     `orm:"isborrow" json:"isborrow"`
	Borrowopenid string  `orm:"borrowopenid" json:"borrowopenid"`
}

func (*Ewei_shop_groups_order) TableName() string {
	return "ewei_shop_groups_order"
}

type Mc_chats_record struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Acid       int    `orm:"acid" json:"acid"`
	Flag       int    `orm:"flag" json:"flag"`
	Openid     string `orm:"openid" json:"openid"`
	Msgtype    string `orm:"msgtype" json:"msgtype"`
	Content    string `orm:"content" json:"content"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Mc_chats_record) TableName() string {
	return "mc_chats_record"
}

type Mc_mass_record struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Acid          int    `orm:"acid" json:"acid"`
	Groupname     string `orm:"groupname" json:"groupname"`
	Fansnum       int    `orm:"fansnum" json:"fansnum"`
	Msgtype       string `orm:"msgtype" json:"msgtype"`
	Content       string `orm:"content" json:"content"`
	Group         int    `orm:"group" json:"group"`
	AttachId      int    `orm:"attach_id" json:"attach_id"`
	MediaId       string `orm:"media_id" json:"media_id"`
	Type          int    `orm:"type" json:"type"`
	Status        int    `orm:"status" json:"status"`
	CronId        int    `orm:"cron_id" json:"cron_id"`
	Sendtime      int    `orm:"sendtime" json:"sendtime"`
	Finalsendtime int    `orm:"finalsendtime" json:"finalsendtime"`
	Createtime    int    `orm:"createtime" json:"createtime"`
}

func (*Mc_mass_record) TableName() string {
	return "mc_mass_record"
}

type Site_styles_vars struct {
	Id          int    `orm:"id" json:"id"`
	Uniacid     int    `orm:"uniacid" json:"uniacid"`
	Templateid  int    `orm:"templateid" json:"templateid"`
	Styleid     int    `orm:"styleid" json:"styleid"`
	Variable    string `orm:"variable" json:"variable"`
	Content     string `orm:"content" json:"content"`
	Description string `orm:"description" json:"description"`
}

func (*Site_styles_vars) TableName() string {
	return "site_styles_vars"
}

type Ewei_shop_groups_goods struct {
	Id            int     `orm:"id" json:"id"`
	Displayorder  int     `orm:"displayorder" json:"displayorder"`
	Uniacid       int     `orm:"uniacid" json:"uniacid"`
	Title         string  `orm:"title" json:"title"`
	Category      int     `orm:"category" json:"category"`
	Stock         int     `orm:"stock" json:"stock"`
	Price         float64 `orm:"price" json:"price"`
	Groupsprice   float64 `orm:"groupsprice" json:"groupsprice"`
	Singleprice   float64 `orm:"singleprice" json:"singleprice"`
	Goodsnum      int     `orm:"goodsnum" json:"goodsnum"`
	Units         string  `orm:"units" json:"units"`
	Freight       float64 `orm:"freight" json:"freight"`
	Endtime       int     `orm:"endtime" json:"endtime"`
	Groupnum      int     `orm:"groupnum" json:"groupnum"`
	Sales         int     `orm:"sales" json:"sales"`
	Thumb         string  `orm:"thumb" json:"thumb"`
	Description   string  `orm:"description" json:"description"`
	Content       string  `orm:"content" json:"content"`
	Createtime    int     `orm:"createtime" json:"createtime"`
	Status        int     `orm:"status" json:"status"`
	Ishot         int     `orm:"ishot" json:"ishot"`
	Deleted       int     `orm:"deleted" json:"deleted"`
	Goodsid       int     `orm:"goodsid" json:"goodsid"`
	Followneed    int     `orm:"followneed" json:"followneed"`
	Followtext    string  `orm:"followtext" json:"followtext"`
	ShareTitle    string  `orm:"share_title" json:"share_title"`
	ShareIcon     string  `orm:"share_icon" json:"share_icon"`
	ShareDesc     string  `orm:"share_desc" json:"share_desc"`
	Goodssn       string  `orm:"goodssn" json:"goodssn"`
	Productsn     string  `orm:"productsn" json:"productsn"`
	Showstock     int     `orm:"showstock" json:"showstock"`
	Purchaselimit int     `orm:"purchaselimit" json:"purchaselimit"`
	Single        int     `orm:"single" json:"single"`
	Dispatchtype  int     `orm:"dispatchtype" json:"dispatchtype"`
	Dispatchid    int     `orm:"dispatchid" json:"dispatchid"`
	Isindex       int     `orm:"isindex" json:"isindex"`
	Followurl     string  `orm:"followurl" json:"followurl"`
	Deduct        float64 `orm:"deduct" json:"deduct"`
	Rights        int     `orm:"rights" json:"rights"`
	ThumbUrl      string  `orm:"thumb_url" json:"thumb_url"`
	Gid           int     `orm:"gid" json:"gid"`
	Discount      int     `orm:"discount" json:"discount"`
	Headstype     int     `orm:"headstype" json:"headstype"`
	Headsmoney    float64 `orm:"headsmoney" json:"headsmoney"`
	Headsdiscount int     `orm:"headsdiscount" json:"headsdiscount"`
	Isdiscount    int     `orm:"isdiscount" json:"isdiscount"`
	Isverify      int     `orm:"isverify" json:"isverify"`
	Verifytype    int     `orm:"verifytype" json:"verifytype"`
	Verifynum     int     `orm:"verifynum" json:"verifynum"`
	Storeids      string  `orm:"storeids" json:"storeids"`
	Merchid       int     `orm:"merchid" json:"merchid"`
	Shorttitle    string  `orm:"shorttitle" json:"shorttitle"`
	Teamnum       int     `orm:"teamnum" json:"teamnum"`
}

func (*Ewei_shop_groups_goods) TableName() string {
	return "ewei_shop_groups_goods"
}

type Ewei_shop_system_plugingrant_log struct {
	Id           int    `orm:"id" json:"id"`
	Logno        string `orm:"logno" json:"logno"`
	Code         string `orm:"code" json:"code"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Pluginid     int    `orm:"pluginid" json:"pluginid"`
	Identity     string `orm:"identity" json:"identity"`
	Type         string `orm:"type" json:"type"`
	Month        int    `orm:"month" json:"month"`
	Permendtime  int    `orm:"permendtime" json:"permendtime"`
	Permlasttime int    `orm:"permlasttime" json:"permlasttime"`
	Isperm       int    `orm:"isperm" json:"isperm"`
	Createtime   int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_system_plugingrant_log) TableName() string {
	return "ewei_shop_system_plugingrant_log"
}

type Ewei_shop_task_qr struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Openid   string `orm:"openid" json:"openid"`
	Recordid int    `orm:"recordid" json:"recordid"`
	Sceneid  string `orm:"sceneid" json:"sceneid"`
	Mediaid  string `orm:"mediaid" json:"mediaid"`
	Ticket   string `orm:"ticket" json:"ticket"`
}

func (*Ewei_shop_task_qr) TableName() string {
	return "ewei_shop_task_qr"
}

type Ewei_shop_article_share struct {
	Id        int     `orm:"id" json:"id"`
	Aid       int     `orm:"aid" json:"aid"`
	ShareUser int     `orm:"share_user" json:"share_user"`
	ClickUser int     `orm:"click_user" json:"click_user"`
	ClickDate string  `orm:"click_date" json:"click_date"`
	AddCredit int     `orm:"add_credit" json:"add_credit"`
	AddMoney  float64 `orm:"add_money" json:"add_money"`
	Uniacid   int     `orm:"uniacid" json:"uniacid"`
}

func (*Ewei_shop_article_share) TableName() string {
	return "ewei_shop_article_share"
}

type Ewei_shop_globonus_level struct {
	Id              int     `orm:"id" json:"id"`
	Uniacid         int     `orm:"uniacid" json:"uniacid"`
	Levelname       string  `orm:"levelname" json:"levelname"`
	Bonus           float64 `orm:"bonus" json:"bonus"`
	Ordermoney      float64 `orm:"ordermoney" json:"ordermoney"`
	Ordercount      int     `orm:"ordercount" json:"ordercount"`
	Commissionmoney float64 `orm:"commissionmoney" json:"commissionmoney"`
	Bonusmoney      float64 `orm:"bonusmoney" json:"bonusmoney"`
	Downcount       int     `orm:"downcount" json:"downcount"`
}

func (*Ewei_shop_globonus_level) TableName() string {
	return "ewei_shop_globonus_level"
}

type Ewei_shop_live_view struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Openid  string `orm:"openid" json:"openid"`
	Roomid  int    `orm:"roomid" json:"roomid"`
	Viewing int    `orm:"viewing" json:"viewing"`
}

func (*Ewei_shop_live_view) TableName() string {
	return "ewei_shop_live_view"
}

type Ewei_shop_package_goods struct {
	Id           int     `orm:"id" json:"id"`
	Uniacid      int     `orm:"uniacid" json:"uniacid"`
	Pid          int     `orm:"pid" json:"pid"`
	Goodsid      int     `orm:"goodsid" json:"goodsid"`
	Title        string  `orm:"title" json:"title"`
	Thumb        string  `orm:"thumb" json:"thumb"`
	Price        float64 `orm:"price" json:"price"`
	Option       string  `orm:"option" json:"option"`
	Goodssn      string  `orm:"goodssn" json:"goodssn"`
	Productsn    string  `orm:"productsn" json:"productsn"`
	Hasoption    int     `orm:"hasoption" json:"hasoption"`
	Marketprice  float64 `orm:"marketprice" json:"marketprice"`
	Packageprice float64 `orm:"packageprice" json:"packageprice"`
	Commission1  float64 `orm:"commission1" json:"commission1"`
	Commission2  float64 `orm:"commission2" json:"commission2"`
	Commission3  float64 `orm:"commission3" json:"commission3"`
}

func (*Ewei_shop_package_goods) TableName() string {
	return "ewei_shop_package_goods"
}

type Ewei_shop_plugin struct {
	Id           int    `orm:"id" json:"id"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
	Identity     string `orm:"identity" json:"identity"`
	Category     string `orm:"category" json:"category"`
	Name         string `orm:"name" json:"name"`
	Version      string `orm:"version" json:"version"`
	Author       string `orm:"author" json:"author"`
	Status       int    `orm:"status" json:"status"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Desc         string `orm:"desc" json:"desc"`
	Iscom        int    `orm:"iscom" json:"iscom"`
	Deprecated   int    `orm:"deprecated" json:"deprecated"`
	Isv2         int    `orm:"isv2" json:"isv2"`
}

func (*Ewei_shop_plugin) TableName() string {
	return "ewei_shop_plugin"
}

type Ewei_shop_sns_board struct {
	Id                    int    `orm:"id" json:"id"`
	Uniacid               int    `orm:"uniacid" json:"uniacid"`
	Cid                   int    `orm:"cid" json:"cid"`
	Title                 string `orm:"title" json:"title"`
	Logo                  string `orm:"logo" json:"logo"`
	Desc                  string `orm:"desc" json:"desc"`
	Displayorder          int    `orm:"displayorder" json:"displayorder"`
	Enabled               int    `orm:"enabled" json:"enabled"`
	Showgroups            string `orm:"showgroups" json:"showgroups"`
	Showlevels            string `orm:"showlevels" json:"showlevels"`
	Postgroups            string `orm:"postgroups" json:"postgroups"`
	Postlevels            string `orm:"postlevels" json:"postlevels"`
	Showagentlevels       string `orm:"showagentlevels" json:"showagentlevels"`
	Postagentlevels       string `orm:"postagentlevels" json:"postagentlevels"`
	Postcredit            int    `orm:"postcredit" json:"postcredit"`
	Replycredit           int    `orm:"replycredit" json:"replycredit"`
	Bestcredit            int    `orm:"bestcredit" json:"bestcredit"`
	Bestboardcredit       int    `orm:"bestboardcredit" json:"bestboardcredit"`
	Notagent              int    `orm:"notagent" json:"notagent"`
	Notagentpost          int    `orm:"notagentpost" json:"notagentpost"`
	Topcredit             int    `orm:"topcredit" json:"topcredit"`
	Topboardcredit        int    `orm:"topboardcredit" json:"topboardcredit"`
	Status                int    `orm:"status" json:"status"`
	Noimage               int    `orm:"noimage" json:"noimage"`
	Novoice               int    `orm:"novoice" json:"novoice"`
	Needfollow            int    `orm:"needfollow" json:"needfollow"`
	Needpostfollow        int    `orm:"needpostfollow" json:"needpostfollow"`
	ShareTitle            string `orm:"share_title" json:"share_title"`
	ShareIcon             string `orm:"share_icon" json:"share_icon"`
	ShareDesc             string `orm:"share_desc" json:"share_desc"`
	Keyword               string `orm:"keyword" json:"keyword"`
	Isrecommand           int    `orm:"isrecommand" json:"isrecommand"`
	Banner                string `orm:"banner" json:"banner"`
	Needcheck             int    `orm:"needcheck" json:"needcheck"`
	Needcheckmanager      int    `orm:"needcheckmanager" json:"needcheckmanager"`
	Needcheckreply        int    `orm:"needcheckreply" json:"needcheckreply"`
	Needcheckreplymanager int    `orm:"needcheckreplymanager" json:"needcheckreplymanager"`
	Showsnslevels         string `orm:"showsnslevels" json:"showsnslevels"`
	Postsnslevels         string `orm:"postsnslevels" json:"postsnslevels"`
	Showpartnerlevels     string `orm:"showpartnerlevels" json:"showpartnerlevels"`
	Postpartnerlevels     string `orm:"postpartnerlevels" json:"postpartnerlevels"`
	Notpartner            int    `orm:"notpartner" json:"notpartner"`
	Notpartnerpost        int    `orm:"notpartnerpost" json:"notpartnerpost"`
}

func (*Ewei_shop_sns_board) TableName() string {
	return "ewei_shop_sns_board"
}

type Ewei_shop_wxapp_bind struct {
	Id      int `orm:"id" json:"id"`
	Uniacid int `orm:"uniacid" json:"uniacid"`
	Wxapp   int `orm:"wxapp" json:"wxapp"`
}

func (*Ewei_shop_wxapp_bind) TableName() string {
	return "ewei_shop_wxapp_bind"
}

type Ewei_shop_coupon_guess struct {
	Id       int    `orm:"id" json:"id"`
	Uniacid  int    `orm:"uniacid" json:"uniacid"`
	Couponid int    `orm:"couponid" json:"couponid"`
	Openid   string `orm:"openid" json:"openid"`
	Times    int    `orm:"times" json:"times"`
	Pwdkey   string `orm:"pwdkey" json:"pwdkey"`
	Ok       int    `orm:"ok" json:"ok"`
	Merchid  int    `orm:"merchid" json:"merchid"`
}

func (*Ewei_shop_coupon_guess) TableName() string {
	return "ewei_shop_coupon_guess"
}

type Ewei_shop_goods_option struct {
	Id               int     `orm:"id" json:"id"`
	Uniacid          int     `orm:"uniacid" json:"uniacid"`
	Goodsid          int     `orm:"goodsid" json:"goodsid"`
	Title            string  `orm:"title" json:"title"`
	Thumb            string  `orm:"thumb" json:"thumb"`
	Productprice     float64 `orm:"productprice" json:"productprice"`
	Marketprice      float64 `orm:"marketprice" json:"marketprice"`
	Costprice        float64 `orm:"costprice" json:"costprice"`
	Stock            int     `orm:"stock" json:"stock"`
	Weight           float64 `orm:"weight" json:"weight"`
	Displayorder     int     `orm:"displayorder" json:"displayorder"`
	Specs            string  `orm:"specs" json:"specs"`
	SkuId            string  `orm:"skuId" json:"skuId"`
	Goodssn          string  `orm:"goodssn" json:"goodssn"`
	Productsn        string  `orm:"productsn" json:"productsn"`
	Virtual          int     `orm:"virtual" json:"virtual"`
	ExchangeStock    int     `orm:"exchange_stock" json:"exchange_stock"`
	ExchangePostage  float64 `orm:"exchange_postage" json:"exchange_postage"`
	Presellprice     float64 `orm:"presellprice" json:"presellprice"`
	Day              int     `orm:"day" json:"day"`
	Allfullbackprice float64 `orm:"allfullbackprice" json:"allfullbackprice"`
	Fullbackprice    float64 `orm:"fullbackprice" json:"fullbackprice"`
	Allfullbackratio float64 `orm:"allfullbackratio" json:"allfullbackratio"`
	Fullbackratio    float64 `orm:"fullbackratio" json:"fullbackratio"`
	Isfullback       int     `orm:"isfullback" json:"isfullback"`
	Islive           int     `orm:"islive" json:"islive"`
	Liveprice        float64 `orm:"liveprice" json:"liveprice"`
}

func (*Ewei_shop_goods_option) TableName() string {
	return "ewei_shop_goods_option"
}

type Ewei_shop_perm_plugin struct {
	Id      int    `orm:"id" json:"id"`
	Acid    int    `orm:"acid" json:"acid"`
	Uid     int    `orm:"uid" json:"uid"`
	Type    int    `orm:"type" json:"type"`
	Plugins string `orm:"plugins" json:"plugins"`
	Coms    string `orm:"coms" json:"coms"`
	Datas   string `orm:"datas" json:"datas"`
}

func (*Ewei_shop_perm_plugin) TableName() string {
	return "ewei_shop_perm_plugin"
}

type Ewei_shop_wxapp_poster struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Type         int    `orm:"type" json:"type"`
	Title        string `orm:"title" json:"title"`
	Thumb        string `orm:"thumb" json:"thumb"`
	Bgimg        string `orm:"bgimg" json:"bgimg"`
	Data         string `orm:"data" json:"data"`
	Createtime   int    `orm:"createtime" json:"createtime"`
	Lastedittime int    `orm:"lastedittime" json:"lastedittime"`
	Status       int    `orm:"status" json:"status"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Ewei_shop_wxapp_poster) TableName() string {
	return "ewei_shop_wxapp_poster"
}

type Mc_oauth_fans struct {
	Id          int    `orm:"id" json:"id"`
	OauthOpenid string `orm:"oauth_openid" json:"oauth_openid"`
	Acid        int    `orm:"acid" json:"acid"`
	Uid         int    `orm:"uid" json:"uid"`
	Openid      string `orm:"openid" json:"openid"`
}

func (*Mc_oauth_fans) TableName() string {
	return "mc_oauth_fans"
}

type Wxcard_reply struct {
	Id        int    `orm:"id" json:"id"`
	Rid       int    `orm:"rid" json:"rid"`
	Title     string `orm:"title" json:"title"`
	CardId    string `orm:"card_id" json:"card_id"`
	Cid       int    `orm:"cid" json:"cid"`
	BrandName string `orm:"brand_name" json:"brand_name"`
	LogoUrl   string `orm:"logo_url" json:"logo_url"`
	Success   string `orm:"success" json:"success"`
	Error     string `orm:"error" json:"error"`
}

func (*Wxcard_reply) TableName() string {
	return "wxcard_reply"
}

type Ewei_shop_cashier_pay_log_goods struct {
	Id        int     `orm:"id" json:"id"`
	Cashierid int     `orm:"cashierid" json:"cashierid"`
	Logid     int     `orm:"logid" json:"logid"`
	Goodsid   int     `orm:"goodsid" json:"goodsid"`
	Price     float64 `orm:"price" json:"price"`
	Total     int     `orm:"total" json:"total"`
}

func (*Ewei_shop_cashier_pay_log_goods) TableName() string {
	return "ewei_shop_cashier_pay_log_goods"
}

type Ewei_shop_customer_robot struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Cate       int    `orm:"cate" json:"cate"`
	Keywords   string `orm:"keywords" json:"keywords"`
	Title      string `orm:"title" json:"title"`
	Content    string `orm:"content" json:"content"`
	Url        string `orm:"url" json:"url"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_customer_robot) TableName() string {
	return "ewei_shop_customer_robot"
}

type Ewei_shop_exhelper_senduser struct {
	Id            int    `orm:"id" json:"id"`
	Uniacid       int    `orm:"uniacid" json:"uniacid"`
	Sendername    string `orm:"sendername" json:"sendername"`
	Sendertel     string `orm:"sendertel" json:"sendertel"`
	Sendersign    string `orm:"sendersign" json:"sendersign"`
	Sendercode    int    `orm:"sendercode" json:"sendercode"`
	Senderaddress string `orm:"senderaddress" json:"senderaddress"`
	Sendercity    string `orm:"sendercity" json:"sendercity"`
	Isdefault     int    `orm:"isdefault" json:"isdefault"`
	Merchid       int    `orm:"merchid" json:"merchid"`
	Province      string `orm:"province" json:"province"`
	City          string `orm:"city" json:"city"`
	Area          string `orm:"area" json:"area"`
}

func (*Ewei_shop_exhelper_senduser) TableName() string {
	return "ewei_shop_exhelper_senduser"
}

type Ewei_shop_feedback struct {
	Id         int    `orm:"id" json:"id"`
	Uniacid    int    `orm:"uniacid" json:"uniacid"`
	Openid     string `orm:"openid" json:"openid"`
	Type       int    `orm:"type" json:"type"`
	Status     int    `orm:"status" json:"status"`
	Feedbackid string `orm:"feedbackid" json:"feedbackid"`
	Transid    string `orm:"transid" json:"transid"`
	Reason     string `orm:"reason" json:"reason"`
	Solution   string `orm:"solution" json:"solution"`
	Remark     string `orm:"remark" json:"remark"`
	Createtime int    `orm:"createtime" json:"createtime"`
}

func (*Ewei_shop_feedback) TableName() string {
	return "ewei_shop_feedback"
}

type Ewei_shop_form_category struct {
	Id      int    `orm:"id" json:"id"`
	Uniacid int    `orm:"uniacid" json:"uniacid"`
	Name    string `orm:"name" json:"name"`
}

func (*Ewei_shop_form_category) TableName() string {
	return "ewei_shop_form_category"
}

type Mc_groups struct {
	Groupid   int    `orm:"groupid" json:"groupid"`
	Uniacid   int    `orm:"uniacid" json:"uniacid"`
	Title     string `orm:"title" json:"title"`
	Credit    int    `orm:"credit" json:"credit"`
	Isdefault int    `orm:"isdefault" json:"isdefault"`
	Orderlist int    `orm:"orderlist" json:"orderlist"`
}

func (*Mc_groups) TableName() string {
	return "mc_groups"
}

type Uni_account_modules struct {
	Id           int    `orm:"id" json:"id"`
	Uniacid      int    `orm:"uniacid" json:"uniacid"`
	Module       string `orm:"module" json:"module"`
	Enabled      int    `orm:"enabled" json:"enabled"`
	Settings     string `orm:"settings" json:"settings"`
	Shortcut     int    `orm:"shortcut" json:"shortcut"`
	Displayorder int    `orm:"displayorder" json:"displayorder"`
}

func (*Uni_account_modules) TableName() string {
	return "uni_account_modules"
}
