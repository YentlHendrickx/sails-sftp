package types

// Struct used to pass connection input from frontend to backend
type ConnectionInput struct {
	Host     string       `json:"host"`
	Type     ProtocolType `json:"type"`
	Port     int          `json:"port"`
	Username string       `json:"username"`
	Password string       `json:"password"` // TODO: should be handled securely
}

type ConnectionConfig struct {
	Name     string `json:"name"` // Name of profile
	UUID     string `json:"uuid"` // UUID for the connection instance, used to identify the connection in memory
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"` // TODO: should be handled securely
}

type ConnectResult struct {
	IsConnected bool             `json:"is_connected"`
	Type        ProtocolType     `json:"type"`   // Protocol type, used to identify the connection type
	Config      ConnectionConfig `json:"config"` // Data returned from the connection, can be nil if not connected
}
