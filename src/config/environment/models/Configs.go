package models

type Configs struct {
	App struct {
		ENV         string `mapstructure:"env"`
		Port        string `mapstructure:"port"`
		ContextPath string `mapstructure:"context_path"`
	} `mapstructure:"app"`
	Gin struct {
		Mode     string `mapstructure:"mode"`
		RootPath string `mapstructure:"root_path"`
		Version  string `mapstructure:"version"`
		Configs  struct {
			TrustedProxies string `mapstructure:"trusted_proxies"`
		} `mapstructure:"configs"`
	} `mapstructure:"gin"`
}
