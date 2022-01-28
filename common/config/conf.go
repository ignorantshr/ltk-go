package config

type Server struct {
	Host string
	Port int
}

type RocketMQ struct {
	Host string
	Port int
}

type Config struct {
	Server   Server
	RocketMQ RocketMQ
}
