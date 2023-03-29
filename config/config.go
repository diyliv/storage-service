package config

import "github.com/spf13/viper"

type Config struct {
	TimeScaleDb TimeScaleDb `yaml:"TimeScaleDb"`
	Kafka       Kafka       `yaml:"Kafka"`
}

type TimeScaleDb struct {
	Host            string `yaml:"Host"`
	Port            string `yaml:"Port"`
	Login           string `yaml:"Login"`
	Password        string `yaml:"Password"`
	DB              string `yaml:"DB"`
	ConnMaxLifeTime int    `yaml:"ConnMaxLifeTime"`
	MaxOpenConn     int    `yaml:"MaxOpenConn"`
	MaxIdleConn     int    `yaml:"MaxIdleConn"`
	SSLMode         string `yaml:"SSLMode"`
}

type Kafka struct {
	Brokers   []string `yaml:"Brokers"`
	Topic     string   `yaml:"Topic"`
	Partition int      `yaml:"Partition"`
	GroupID   string   `yaml:"GroupID"`
	ReadFrom  string   `yaml:"ReadFrom"`
}

func ReadConfig(cfgName, cfgType, cfgPath string) *Config {
	var cfg Config
	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
