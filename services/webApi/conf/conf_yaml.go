package conf

import (
	"github.com/ItsWewin/superfactory/basicmatter"
	"github.com/ItsWewin/superfactory/logger"
)

var Conf = new(Config)

type Config struct {
	Server       *Server       `yaml:"server"`
	Mysql        MysqlConf     `yaml:"mysql"`
	Redis        Redis         `yaml:"redis"`
	OAuth        *OAuth        `yaml:"oauth"`
	GrpcClient   *GrpcClients  `yaml:"grpc_client"`
	SessionStore *SessionStore `yaml:"session_store"`
}

type SessionStore struct {
	Redis *RedisConf `yaml:"redis"`
}

type Server struct {
	Name       string      `yaml:"name"`
	HttpServer *HttpServer `yaml:"http_server"`
	GrpcServer *GrpcServer `yaml:"grpc_server"`
}

type HttpServer struct {
	Addr string `yaml:"addr"`
}

type GrpcClients struct {
	Callisto   *GrpcClient `yaml:"callisto"`
	UserServer *GrpcClient `yaml:"ganymede"`
	Charon     *GrpcClient `yaml:"charon"`
}

type GrpcClient struct {
	Addr string `json:"addr"`
}

type GrpcServer struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type MysqlConf struct {
	Europadb *Mysql `yaml:"europadb"`
}

type Mysql struct {
	DB           string `yaml:"db"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `json:"port"`
	Protocol     string `yaml:"protocol"`
	MaxLifeTime  int64  `yaml:"max_life_time"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type Redis struct {
	RedisO RedisConf `yaml:"redis_0"`
}

type RedisConf struct {
	Addr     string `yaml:"addr"`
	DBNo     int    `yaml:"db_no"`
	Password string `yaml:"password"`
}

type OAuth struct {
	Github *GitHubOAuth `yaml:"github"`
}

type GitHubOAuth struct {
	Provider       string `yaml:"provider"`
	ClientID       string `yaml:"client_id"`
	ClientSecret   string `yaml:"client_secrets"`
	RedirectUrl    string `yaml:"redirect_url"`
	AccessTokenUrl string `yaml:"access_token_url"`
	UserProfileUrl string `yaml:"user_profile_url"`
}

func (c *Config) MaterType() string {
	return basicmatter.MasterConfigBasicYaml
}

func (c *Config) RunHttpServer() bool {
	return true
}

func (c *Config) HttpAddr() string {
	return c.Server.HttpServer.Addr
}

func (c *Config) Print() {
	logger.Infof("config: %s", logger.ToJson(c))
}
