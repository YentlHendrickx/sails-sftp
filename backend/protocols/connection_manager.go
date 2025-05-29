package protocols

import (
	"context"
	"fmt"
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
		ctx:         ctx,
		Logger:      logger,
		Connections: make(map[string]types.BaseClient),
	}
}

func (c *ConnectionManager) Connect(config *types.ConnectionInput) (*types.ConnectResult, error) {
	// TODO:
	// This should be loaded from sqlite, the ID is only done once, each connection later has a UUID
	// So: ID => Identifies profile in sqlite, UUID => Identifies connection instance in memory
	connectionUUID := uuid.New().String()
	c.Logger.Log(logging.LevelInfo, "Creating new connection with UUID: "+connectionUUID)

	connectionConfig := &types.ConnectionConfig{
		UUID:     connectionUUID,
		Name:     config.Host, // Use host as name for simplicity, can be changed later
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		Password: config.Password, // TODO: Handle password securely, maybe use a vault or encryption
	}

	var client types.BaseClient
	switch config.Type {
	case types.ProtocolSFTP:
		sftpConfig := &sftp.SFTPConfig{
			ConnectionConfig: *connectionConfig,
		}

		client = sftp.NewSFTPClient(c.ctx, *sftpConfig, c.Logger) // Initialize with nil config, will be set later
	default:
		c.Logger.Log(logging.LevelError, "Unsupported protocol type: "+config.Type.String())
		return nil, &types.SailsError{
			Operation: "Connect",
			Reason:    fmt.Sprintf("Unsupported protocol type: %s", config.Type.String()),
			Code:      400,
		}
	}

	error := client.Connect(c.ctx)

	if error != nil {
		c.Logger.Log(logging.LevelError, "Failed to connect to "+config.Type.String()+" server: "+error.Error())
		client.Disconnect(c.ctx) // Ensure we clean up the client if connection fails
		return nil, &types.SailsError{
			Operation: "Connect",
			Reason:    fmt.Sprintf("Failed to connect to %s server: %s", config.Type.String(), error.Error()),
			Code:      500,
		}
	}

	// Make sure c.Connections is initialized
	c.Connections[connectionUUID] = client // Store the client in the connections map

	return &types.ConnectResult{
		IsConnected: true,
		Type:        types.ProtocolSFTP,
		Config:      c.Connections[connectionUUID].GetConfig(),
	}, nil
}

func (c *ConnectionManager) Disconnect(uuid string) (bool, error) {
	client, exists := c.Connections[uuid]
	if !exists {
		c.Logger.Log(logging.LevelError, "Connection with UUID "+uuid+" does not exist")
		return false, &types.SailsError{
			Operation: "Disconnect",
			Reason:    fmt.Sprintf("Connection with UUID %s does not exist", uuid),
			Code:      404,
		}
	}

	err := client.Disconnect(c.ctx)
	if err != nil {
		c.Logger.Log(logging.LevelError, "Failed to disconnect from connection with UUID "+uuid+": "+err.Error())
		return false, &types.SailsError{
			Operation: "Disconnect",
			Reason:    fmt.Sprintf("Failed to disconnect from connection with UUID %s: %s", uuid, err.Error()),
			Code:      500,
		}
	}

	delete(c.Connections, uuid) // Remove the client from the connections map
	c.Logger.Log(logging.LevelInfo, "Successfully disconnected from connection with UUID "+uuid)

	return true, nil
}
