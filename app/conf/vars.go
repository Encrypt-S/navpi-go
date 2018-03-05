package conf

// AppConf holds the app config
// NavConf           string   `json:"navconf"`
// RunningNavVersion string   `json:"runningNavVersion"`
// AllowedIps        []string `json:"allowedIps"`
// UIPassword        string   `json:"uiPassword"`
var AppConf AppConfig

// NavConf holds the rpc config
// RPCUser     string `json:"rpcUser"`
// RPCPassword string `json:"rpcPassword"`
var NavConf NavConfig

// ServerConf holds the server config
// ManagerAPIPort   int64
// DaemonAPIPort    int64
// SetupAPIPort     int64
// LatestReleaseAPI string
// ReleaseAPI       string
// DaemonHeartbeat  int64
var ServerConf ServerConfig
