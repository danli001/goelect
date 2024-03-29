package config

// Config represents the elect config
type Config struct {
	// ConnectTimeout represents the timeout duration for a rpc connection
	ConnectTimeout uint `json:"connect_timeout,omitempty"`
	// Peers contain information about all nodes in the cluster.
	Peers []NodeConfig `json:"peers" json:"peers,omitempty"`
}

type NodeConfig struct {
	// ID of node
	ID string `json:"id"`
	// Address of node, used for establishing connections
	Address string `json:"address"`
	// NoVote represents whether the node participates in voting or not
	NoVote bool `json:"no_vote"`
	// Tags represent additional label information of the node
	Tags map[string]string `json:"tags"`
}
