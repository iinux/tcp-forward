package iplist

type StoreAdapter interface {
	GetIpSet() []string
	AddIp(ip string) bool
}

