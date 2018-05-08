package models

// Config - config info
type Config struct {
	AppVersion                  string `json:"AppVersion"`
	DockerEndPoint              string `json:"DockerEndPoint`
	DockerAPIVersion            string `json:"DockerAPIVersion`
	DockerRegistryAddress       string `json:"-"`
	EnableBuildImage            bool   `json:"-"`
	DockerComposePath           string `json:"DockerComposePath"`
	DockerComposePackageMaxSize int64  `json:"DockerComposePackageMaxSize"`
	DockerAgentIPAddr           string `json:"DockerAgentIPAddr"`
	DockerClusterEnabled        bool   `json:"DockerClusterEnabled"`
	DockerClusterURIs           string `json:"DockerClusterURIs"`
	DockerClusterName           string `json:"DockerClusterName"`
	DockerClusterHeartBeat      string `json:"DockerClusterHeartBeat"`
	DockerClusterTTL            string `json:"DockerClusterTTL"`
	DockerClusterPortsRange     string `json:"DockerClusterPortsRange"`
	LogLevel                    int    `json:"LogLevel"`
}
