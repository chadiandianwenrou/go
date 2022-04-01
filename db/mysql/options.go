package mysql

// Options options
type Options struct {
	Hostname     string `json:"hostname,omitempty" yaml:"hostname"`
	Port         int    `json:"port,omitempty" yaml:"port"`
	Username     string `json:"username,omitempty" yaml:"username"`
	Password     string `json:"password,omitempty" yaml:"password"`
	Database     string `json:"database,omitempty" yaml:"database"`
	PingAttempts int    `json:"pingAttempts,omitempty" yaml:"pingAttempts"`
	MaxLifetime  int    `json:"maxLifetime,omitempty" yaml:"maxLifetime"` // second
	MaxConns     int    `json:"maxConns,omitempty" yaml:"maxConns"`
}

// Option Options function
type Option func(*Options)

// NewOptions new Options
func NewOptions(opts ...Option) Options {
	opt := Options{
		Hostname:     "127.0.0.1",
		Port:         3306,
		Username:     "root",
		Password:     "root",
		Database:     "test",
		PingAttempts: 3,
		MaxLifetime:  10 * 60, // 10minutes
		MaxConns:     10,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Hostname set hostname
func Hostname(hostname string) Option {
	return func(o *Options) {
		o.Hostname = hostname
	}
}

// Port set port
func Port(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

// Username set username
func Username(username string) Option {
	return func(o *Options) {
		o.Username = username
	}
}

// Password set password
func Password(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}

// Database set database
func Database(database string) Option {
	return func(o *Options) {
		o.Database = database
	}
}

// PingAttempts set PingAttempts
func PingAttempts(attempts int) Option {
	return func(o *Options) {
		o.PingAttempts = attempts
	}
}
