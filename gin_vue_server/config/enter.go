package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Log    Log    `yaml:"log"`
	System System `yaml:"system"`
}
