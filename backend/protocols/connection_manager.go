package protocols

import (
	"context"
	"sails-sftp/backend/protocols/sftp"
	"sails-sftp/backend/utils/logging"

	"github.com/google/uuid"
)

func NewManager(ctx context.Context, logger *logging.Logger) *ConnectionManager {
	return &ConnectionManager{
		ctx:    ctx,
		Logger: logger,
		// No connections yet
	}
}

func (c *ConnectionManager) Connect() *ConnectResult {
	// Create new UUID for connection
	connectionID := uuid.New().String()
	c.Logger.Log(logging.LevelInfo, "Creating new connection with ID: "+connectionID)

	// Create new connection config
	connectionConfig := ConnectionConfig{
		ID:       connectionID,
		Name:     "New SFTP Connection",
		Host:     "localhost", // TODO: Replace with actual host
		Port:     22,          // Default SFTP port
		Username: "user",      // TODO: Replace with actual username
		Password: "password",  // TODO: Replace with actual password, handle securely
	}

	// Create new client based on type
	client := sftp.NewSFTPClient(c.ctx, connectionConfig, c.Logger)

	// If error, report error

	return &ConnectResult{
		IsConnected: true,
		Type:        ProtocolSFTP,
		Data:        nil,
	}
}
