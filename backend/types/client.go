package types

import (
	"context"
)

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
