package systemSetting

var (
	UserMenus    map[string]struct{}
	IpWhitelists map[int64]map[string]struct{}
	LoginAdmin   map[int64]string
)

func init() {
	UserMenus = make(map[string]struct{})
	IpWhitelists = make(map[int64]map[string]struct{})
	LoginAdmin = make(map[int64]string)
}
