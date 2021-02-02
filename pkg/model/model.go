package model

type ServerCfg struct {
	Port     int
	Username string
	Password string
}

type Volume struct {
	Group         string `json:"group"`
	Size          uint64 `json:"size"`
	Unit          string `json:"unit"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	ThinProvision bool   `json:"thin"`
	ThinPool      string `json:"pool"`
}

type Lun struct {
	TargetIQN string   `json:"targetIQN"`
	Volume    *Volume  `json:"volume"`
	AclIpList []string `json:"aclList"`
}

type Target struct {
	TargetId  string `json:"-"`
	TargetIQN string `json:"targetIQN"`
}

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
