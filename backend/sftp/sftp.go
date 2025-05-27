// SFTP Management Package - TODO: Should this be generalized into FTP and SFTP?
package sftp

import (
	"context"
	"fmt"
)

type SFTP struct {
	ctx context.Context
}

func NewSFTP(ctx context.Context) *SFTP {
	return &SFTP{ctx: ctx}
}

func (s *SFTP) SayHello(name string) string {
	return fmt.Sprintf("Hello %s, welcome to SFTP!", name)
}
