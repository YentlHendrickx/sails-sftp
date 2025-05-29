package types

type ProtocolType string

const (
	ProtocolSFTP ProtocolType = "sftp"
	ProtocolFTP  ProtocolType = "ftp"

	// TODO: Long term, i want full SSH support, not just SFTP over SSH - so we can have terminal access directly from the app
	// ProtocolSSH ProtocolType = "ssh"
)

func (p ProtocolType) String() string {
	return string(p)
}
