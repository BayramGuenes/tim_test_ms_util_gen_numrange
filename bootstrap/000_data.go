package bootstrap

type ApplConfStruct struct {
	Dummy      string
	DB_UserPwd string
}

var ApplConf ApplConfStruct

var (
	ThisCommunnicationProtokoll string
	ServiceLocationIDSysReg     string
	ServiceDB                   string
	PortDB                      string
	ThisServerAddr              string
	ThisPort                    string
	PortPresentation            string
	ThisConfLocation            string
)
