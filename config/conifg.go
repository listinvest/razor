package config

type ManagerConfig struct {
	Address        string
	MaxConnections int
	Trackers       []string
}
type PeerConfig struct {
	InfoHash    []byte
	PeerID      []byte
	PieceCount  int
	IdleTimeout int
	BlockSize   uint32
	PieceSize   uint32
}
