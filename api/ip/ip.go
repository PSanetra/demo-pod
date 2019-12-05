package ip

type Ip struct {
	// IP which sent this request (maybe the ip of a proxy)
	ClientIp string `json:"client_ip"`
	// clientIp or IP from the X-Forwarded-For Header
	OriginalIp string `json:"original_ip"`
}
