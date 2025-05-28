package protocols

import (
	"context"
	"sails-sftp/backend/protocols/sftp"
	"sails-sftp/backend/types"
	"sails-sftp/backend/utils/logging"

	"github.com/google/uuid"
)

type ConnectionManager struct {
	ctx         context.Context
	Logger      *logging.Logger
	Connections map[string]types.BaseClient // Use base client so we can have both file transfer + regular SSH in the future
}

func NewManager(ctx context.Context, logger *logging.Logger) *ConnectionManager {
	return &ConnectionManager{
		ctx:    ctx,
		Logger: logger,
		// No connections yet
	}
}

func (c *ConnectionManager) Connect() *types.ConnectResult {
	// TODO:
	// This should be loaded from sqlite, the ID is only done once, each connection later has a UUID
	// So: ID => Identifies profile in sqlite, UUID => Identifies connection instance in memory
	connectionID := uuid.New().String()
	c.Logger.Log(logging.LevelInfo, "Creating new connection with ID: "+connectionID)

	// Create new connection config
	connectionConfig := sftp.SFTPConfig{
		ConnectionConfig: types.ConnectionConfig{
			ID:       connectionID,
			Name:     "New SFTP Connection",
			Host:     "localhost", // TODO: Replace with actual host
			Port:     22,          // Default SFTP port
			Username: "sftp",      // TODO: Replace with actual username
			Password: "test",      // TODO: Replace with actual password, handle securely
		},
	}

	// Create new client based on type
	client := sftp.NewSFTPClient(c.ctx, connectionConfig, c.Logger)
	error := client.Connect(c.ctx)
	if error != nil {
		c.Logger.Log(logging.LevelError, "Failed to connect to SFTP server: "+error.Error())
		client.Disconnect(c.ctx) // Ensure we clean up the client if connection fails
		return &types.ConnectResult{
			IsConnected: false,
			Type:        types.ProtocolSFTP,
			Data:        error.Error(),
		}
	}

	/** @var client types.FileTransferClient */
	fileInfo, error := client.ListFiles(c.ctx, "/")

	if error != nil {
		c.Logger.Log(logging.LevelError, "Failed to list files on SFTP server: "+error.Error())
		client.Disconnect(c.ctx) // Ensure we clean up the client if listing fails
	}

	return &types.ConnectResult{
		IsConnected: true,
		Type:        types.ProtocolSFTP,
		Data:        fileInfo,
	}
}
