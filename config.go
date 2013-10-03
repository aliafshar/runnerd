
package runnerd

type ProcessConfig struct {
	Name string
	Command string
	Pwd string
	Umask int
	Priority int
}

type Config struct {
	Ps []*ProcessConfig
}
