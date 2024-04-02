package conf

type Config struct {
	Server   `mapstructure:"server" json:"server" yaml:"server"`
	Database `mapstructure:"database" json:"database" yaml:"database"`
}

type Server struct {
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
}

type Database struct {
	DbName   string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
}
