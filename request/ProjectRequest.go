package request

type Project struct {
	Name     	 string `form:"name" json:"name" xml:"name"`
	CallbackUrl  string `form:"callback_url" json:"callback_url" xml:"callback_url"`
	VPNTunnel    bool   `form:"vpn_tunnel" json:"vpn_tunnel" xml:"vpn_tunnel"`
}