package cmd

type IdentityConfig struct {
	Cert       string `mapstructure:"cert"`
	PrivateKey string `mapstructure:"privateKey"`
	MspID      string `mapstructure:"mspId"`
}

type ConnectionConfig struct {
	TLSCert string `mapstructure:"tlsCert"`
	TLS     bool   `mapstructure:"tls"`
}

type Config struct {
	Identity   IdentityConfig   `mapstructure:"identity"`
	Connection ConnectionConfig `mapstructure:"connection"`
}

var config Config
