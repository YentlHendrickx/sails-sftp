package sftp

type Client struct {
	Host       string
	Port       int
	Username   string
	Password   string // should be handled securely -> Do we worry about hashing this?
	PrivateKey string // Path to private key file, if used
	Passphrase string // Passphrase for private key, if used
	Timeout    int    // Timeout in seconds for connection attempts
}
