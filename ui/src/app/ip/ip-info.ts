export interface IpInfo {
  pod_ip_list: string[];
  // IP which sent this request (maybe the ip of a proxy)
  client_ip: string;
  // clientIp or IP from the X-Forwarded-For Header
  original_ip: string;
}
