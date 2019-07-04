package iplist

type WhiteList struct {
	ipSet []string
	ipSetAdapter StoreAdapter
}

func NewWhiteList(adapter StoreAdapter) *WhiteList {
	ret := WhiteList{}

	ret.ipSet = adapter.GetIpSet()
	ret.ipSetAdapter = adapter

	return &ret
}

func (w *WhiteList) Allow(ip string) bool {
	for _, v := range w.ipSet {
		if v == ip {
			return true
		}
	}

	return false
}

func (w *WhiteList) AddIp(ip string) bool {
	w.ipSet = append(w.ipSet, ip)
	return w.ipSetAdapter.AddIp(ip)
}
