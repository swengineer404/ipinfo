package ipinfo

type IPInfo struct {
	IsBogon      bool     `json:"is_bogon"`
	IsDatacenter bool     `json:"is_datacenter"`
	IsTOR        bool     `json:"is_tor"`
	IsProxy      bool     `json:"is_proxy"`
	IsVPN        bool     `json:"is_vpn"`
	IsAbuser     bool     `json:"is_abuser"`
	ASN          ASN      `json:"asn"`
	Location     Location `json:"location"`
}

type ASN struct {
	ASN    int    `json:"asn"`
	Route  string `json:"route"`
	Org    string `json:"org"`
	Domain string `json:"domain"`
}

type Location struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}
