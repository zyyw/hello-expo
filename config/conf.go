package config

const Env = "env"

type ServerConf struct {
	HttpPort int
}

type LogConf struct {
	Path  string
	Level string
}

type Config struct {
	Server ServerConf
	Log    LogConf
}
