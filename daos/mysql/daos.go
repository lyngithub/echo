package mysql

var (
	Adminlog = NewDaoAdminlog()
	User     = NewDaoUser()
	Menu     = NewDaoMenu()
	Role     = NewDaoRole()
	Ver      = NewDaoVer()
)
