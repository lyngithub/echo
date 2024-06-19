package vo

type UserMeVo struct {
	Id          int64  `json:"id"`
	ParentId    int64  `json:"parent_id"`
	LoginName   string `json:"login_name"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	LoginIp     string `json:"login_ip"`
	LoginDate   string `json:"login_date"`
}

type UserParams struct {
	Id          int64   `json:"id"`
	LoginName   string  `json:"login_name"`
	UserName    string  `json:"user_name"`
	Email       string  `json:"email"`
	Phonenumber string  `json:"phonenumber"`
	Status      string  `json:"status"`
	Password    string  `json:"password"`
	Remark      string  `json:"remark"`
	Roles       []int64 `json:"roles"`
}

type UserVo struct {
	Id           int64      `json:"id"`
	ParentId     int64      `json:"parent_id"`
	LoginName    string     `json:"login_name"`
	UserName     string     `json:"user_name"`
	IsGoogleCode string     `json:"is_google_code"`
	Email        string     `json:"email"`
	Phonenumber  string     `json:"phonenumber"`
	Status       string     `json:"status"`
	LoginIp      string     `json:"login_ip"`
	LoginDate    string     `json:"login_date"`
	Remark       string     `json:"remark"`
	Created      string     `json:"created"`
	Updated      string     `json:"updated"`
	Roles        []int64    `json:"roles"`
	RolesMap     []*RolesVo `json:"roles_map"`
}

type RolesVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RoleParams struct {
	Id       int64   `json:"id"`
	ParentId int64   `json:"parent_id"`
	RoleName string  `json:"role_name"`
	Status   string  `json:"status"`
	Remark   string  `json:"remark"`
	Menus    []int64 `json:"menus"`
}

type RoleVo struct {
	Id             int64   `json:"id"`
	ParentId       int64   `json:"parent_id"`
	PrefixRoleName string  `json:"prefix_role_name"`
	RoleName       string  `json:"role_name"`
	Status         string  `json:"status"`
	Remark         string  `json:"remark"`
	Created        string  `json:"created"`
	Updated        string  `json:"updated"`
	Menus          []int64 `json:"menus"`
}

type MenuParams struct {
	Id             int64  `json:"id"`
	ParentId       int64  `json:"parent_id"`
	MenuName       string `json:"menu_name"`
	Weights        int    `json:"weights"`
	Method         string `json:"method"`
	Url            string `json:"url"`
	Pages          string `json:"pages"`
	MenuType       string `json:"menu_type"`
	Classification string `json:"classification"`
	Visible        string `json:"visible"`
	Icon           string `json:"icon"`
	Remark         string `json:"remark"`
}

type MenuVo struct {
	Id             int64     `json:"id"`
	ParentId       int64     `json:"parent_id"`
	MenuName       string    `json:"menu_name"`
	Weights        int       `json:"weights"`
	Method         string    `json:"method"`
	Url            string    `json:"url"`
	Pages          string    `json:"pages"`
	MenuType       string    `json:"menu_type"`
	Classification string    `json:"classification"`
	Visible        string    `json:"visible"`
	Icon           string    `json:"icon"`
	Remark         string    `json:"remark"`
	Created        string    `json:"created"`
	Updated        string    `json:"updated"`
	Children       []*MenuVo `json:"children"`
}

type AdminlogVo struct {
	Id        int64  `json:"id"`
	Method    string `json:"method"`
	Url       string `json:"url"`
	Params    string `json:"params"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	CreateBy  int64  `json:"create_by"`
	Created   string `json:"created"`
}
