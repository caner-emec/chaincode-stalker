package types

type IdentityConfig struct {
	Cert       string `mapstructure:"cert"`
	PrivateKey string `mapstructure:"privateKey"`
	MspID      string `mapstructure:"mspId"`
}

type ConnectionConfig struct {
	TLSCert      string `mapstructure:"tlsCert"`
	PeerEndpoint string `mapstructure:"peerEndpoint"`
	GatewayPeer  string `mapstructure:"gatewayPeer"`
	TLS          bool   `mapstructure:"tls"`
}

type Config struct {
	Identity   IdentityConfig   `mapstructure:"identity"`
	Connection ConnectionConfig `mapstructure:"connection"`
}
