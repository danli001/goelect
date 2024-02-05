package config

type Config struct {
	ConnectTimeout uint         `json:"connect_timeout,omitempty"`
	Peers          []NodeConfig `json:"peers" json:"peers,omitempty"`
}

type NodeConfig struct {
	ID      string            `json:"id"`
	Address string            `json:"address"`
	NoVote  bool              `json:"no_vote"`
	Tags    map[string]string `json:"tags"`
}
