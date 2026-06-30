package main

type StatusResponse struct {
	VPN         bool   `json:"vpn"`
	Internet    bool   `json:"internet"`
	VPS         bool   `json:"vps"`
	ExternalIP  string `json:"external_ip"`
	CPUUsage    string `json:"cpu_usage"`
	RAMUsage    string `json:"ram_usage"`
	Uptime      string `json:"uptime"`
	Temperature string `json:"temperature"`
	Version     string `json:"version"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type VLESSServer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
	UUID        string `json:"uuid"`
	SNI         string `json:"sni"`
	PublicKey   string `json:"public_key"`
	ShortID     string `json:"short_id"`
	Fingerprint string `json:"fingerprint"`
	Flow        string `json:"flow"`
	Enabled     bool   `json:"enabled"`
	Ping        int64  `json:"ping"`
}

type BackupInfo struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Size      int64  `json:"size"`
}

type RoutingRule struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Action  string `json:"action"`
}

type PanelConfig struct {
	AutoFailover bool          `json:"auto_failover"`
	Servers       []VLESSServer `json:"servers"`
	Routing       []RoutingRule `json:"routing"`
}
