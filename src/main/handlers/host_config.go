

package handlers





type HostConfig struct {
     AudioRoot     string  `json:"audio_path"`
     AudioHost     string  `json:"audio_host"`
     WebHost       string  `json:"host_root"`
}

var HostConf * HostConfig

