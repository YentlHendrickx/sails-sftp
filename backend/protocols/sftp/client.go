package sftp

import (
	"context"
	"fmt"
	// "io"
	// "os"
	"sails-sftp/backend/types"
	"sails-sftp/backend/utils/logging"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	config      SFTPConfig
	logger      *logging.Logger
	sshClient   *ssh.Client
	sftpClient  *sftp.Client
	isConnected bool
	ctx         context.Context
}

func NewSFTPClient(ctx context.Context, config SFTPConfig, logger *logging.Logger) *Client {
	return &Client{
		ctx:    ctx,
		config: config,
		logger: logger,
	}
}

func (c *Client) GetType() types.ProtocolType {
	return types.ProtocolSFTP
}

func (c *Client) GetConfig() types.ConnectionConfig {
	return c.config.ConnectionConfig
}

func (c *Client) IsConnected() bool {
	return c.isConnected
}

func (c *Client) Connect(ctx context.Context) error {
	if c.IsConnected() {
		c.logger.Log(logging.LevelWarn, fmt.Sprintf("Already connected to SFTP server %s:%d", c.config.Host, c.config.Port))
		return fmt.Errorf("Already connected to SFTP server %s:%d", c.config.Host, c.config.Port)
	}

	c.logger.Log(logging.LevelInfo, fmt.Sprintf("Connecting to SFTP server %s:%d", c.config.Host, c.config.Port))

	sshConfg := &ssh.ClientConfig{
		User: c.config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.config.Password),
			// TODO: Add support for public key authentication
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: Use a proper host key verification in production
	}

	addr := fmt.Sprintf("%s:%d", c.config.Host, c.config.Port)

	// Establish SSH connection
	client, err := ssh.Dial("tcp", addr, sshConfg)

	if err != nil {
		c.logger.Log(logging.LevelError, fmt.Sprintf("Failed to connect to SFTP server %s:%d: %v", c.config.Host, c.config.Port, err))
		return fmt.Errorf("failed to connect to SFTP server %s:%d: %w", c.config.Host, c.config.Port, err)
	}

	c.sshClient = client

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		client.Close()
		c.logger.Log(logging.LevelError, fmt.Sprintf("Failed to create SFTP client: %v", err))
		return fmt.Errorf("failed to create SFTP client: %w", err)
	}

	c.sftpClient = sftpClient
	c.isConnected = true
	c.logger.Log(logging.LevelInfo, fmt.Sprintf("Successfully connected to SFTP server %s:%d", c.config.Host, c.config.Port))
	return nil
}

func (c *Client) Disconnect(ctx context.Context) error {
	if !c.IsConnected() {
		c.logger.Log(logging.LevelWarn, fmt.Sprintf("Not connected to SFTP server %s:%d", c.config.Host, c.config.Port))
		return fmt.Errorf("not connected to SFTP server %s:%d", c.config.Host, c.config.Port)
	}

	c.logger.Log(logging.LevelInfo, fmt.Sprintf("Disconnecting from SFTP server %s:%d", c.config.Host, c.config.Port))

	var gotError bool
	if c.sftpClient != nil {
		if err := c.sftpClient.Close(); err != nil {
			c.logger.Log(logging.LevelError, fmt.Sprintf("Failed to close SFTP client: %v", err))
			c.logger.Log(logging.LevelInfo, fmt.Sprintf("Continueing with cleanup despite error..."))
			gotError = true
		}

		c.sftpClient = nil
	}

	if c.sshClient != nil {
		if err := c.sshClient.Close(); err != nil {
			c.logger.Log(logging.LevelError, fmt.Sprintf("Failed to close SSH client: %v", err))
			c.logger.Log(logging.LevelInfo, fmt.Sprintf("Continueing with cleanup despite error..."))
			gotError = true
		}

		c.sshClient = nil
	}

	c.isConnected = false
	if gotError {
		c.logger.Log(logging.LevelError, fmt.Sprintf("Disconnected from SFTP server %s:%d with errors", c.config.Host, c.config.Port))
		return fmt.Errorf("disconnected from SFTP server %s:%d with errors", c.config.Host, c.config.Port)
	}

	c.logger.Log(logging.LevelInfo, fmt.Sprintf("Successfully disconnected from SFTP server %s:%d", c.config.Host, c.config.Port))
	return nil
}

func (c *Client) ListFiles(ctx context.Context, path string) ([]types.FileInfo, error) {
	if !c.IsConnected() {
		return nil, fmt.Errorf("not connected to SFTP server %s:%d", c.config.Host, c.config.Port)
	}

	c.logger.Log(logging.LevelInfo, fmt.Sprintf("Listing files in directory: %s", path))

	files, err := c.sftpClient.ReadDir(path)
	if err != nil {
		c.logger.Log(logging.LevelError, fmt.Sprintf("Failed to list files in directory %s: %v", path, err))
		return nil, fmt.Errorf("failed to list files in directory %s: %w", path, err)
	}

	fileInfos := make([]types.FileInfo, len(files))
	for i, file := range files {
		fileInfos[i] = types.FileInfo{
			Name:  file.Name(),
			Size:  file.Size(),
			Mode:  file.Mode().String(),
			Mtime: file.ModTime().Unix(),
			IsDir: file.IsDir(),
		}
	}

	return fileInfos, nil
}
