package types

import (
	"context"
)

type FileInfo struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Mode  string `json:"mode"`  // File permissions
	Mtime int64  `json:"mtime"` // Last modified time in Unix timestamp
	IsDir bool   `json:"is_dir"`
}

type ConnectionConfig struct {
	ID       string `json:"id"`   // id for connection profiles
	Name     string `json:"name"` // Name of profile
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"` // TODO: should be handled securely
	// TODO: UUID? should uuid be passed here or is the 'config' only gonna be used for the profile?
}

type ProtocolType string

const (
	ProtocolSFTP ProtocolType = "sftp"
	ProtocolFTP  ProtocolType = "ftp"

	// TODO: Long term, i want full SSH support, not just SFTP over SSH - so we can have terminal access directly from the app
	// ProtocolSSH ProtocolType = "ssh"
)

type ConnectResult struct {
	IsConnected bool
	Type        ProtocolType
	Data        any

	// TODO: Implemen this back to frontend, store in pinia store so we can identify the connection
	ConnectionUUID string // UUID of the connection, used to identify the connection in the manager
}

// BaseClient is a common interface for all protocol clients
type BaseClient interface {
	Connect(ctx context.Context) error
	// We need context here to allow for cancellation and timeouts
	Disconnect(ctx context.Context) error

	// NOTE: Think more about this - should we return an error if not connected or should error be = 'not connected'?
	IsConnected() bool
	GetType() ProtocolType
	GetConfig() ConnectionConfig
}

// FileTransfer functions wrapped in a common interface
type FileTransferClient interface {
	BaseClient

	ListFiles(ctx context.Context, path string) ([]FileInfo, error)              // ls
	UploadFile(ctx context.Context, localPath string, remotePath string) error   // put
	DownloadFile(ctx context.Context, remotePath string, localPath string) error // get
	DeleteFile(ctx context.Context, remotePath string) error                     // rm
	CreateDirectory(ctx context.Context, path string) error                      // mkdir
	DeleteDirectory(ctx context.Context, path string) error                      // rmdir
	Rename(ctx context.Context, oldPath string, newPath string) error            // mv
}
