// SFTP Management Package - TODO: Should this be generalized into FTP and SFTP?
package sftp

import (
	"context"
	"fmt"
	"sails-sftp/backend/utils/logging"
)

type SFTP struct {
	ctx    context.Context
	logger *logging.Logger
}

func NewSFTP(ctx context.Context, logger *logging.Logger) *SFTP {
	return &SFTP{ctx: ctx, logger: logger}
}

func (s *SFTP) SayHello(name string) string {
	s.logger.Log(logging.LevelInfo, fmt.Sprintf("Saying hello to %s", name))
	return fmt.Sprintf("Hello %s, welcome to SFTP!", name)
}
