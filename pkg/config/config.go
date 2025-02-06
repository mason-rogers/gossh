package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Host struct {
	Name     string `mapstructure:"name"`
	Hostname string `mapstructure:"hostname"`
	User     string `mapstructure:"user"`
	Port     int    `mapstructure:"port"`
	KeyFile  string `mapstructure:"keyfile,omitempty"`
	JumpHost string `mapstructure:"jumphost,omitempty"`
}

type Group struct {
	Name   string  `mapstructure:"name"`
	Hosts  []Host  `mapstructure:"hosts,omitempty"`
	Groups []Group `mapstructure:"groups,omitempty"`
}

type Config struct {
	JumpHosts []Host  `mapstructure:"jumphosts"`
	Groups    []Group `mapstructure:"groups"`
}

var config Config

func Load() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName(".gossh")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")

	// Set defaults
	viper.SetDefault("jumphosts", []Host{})
	viper.SetDefault("groups", []Group{})

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			panic(err)
		}

		// Config file not found - create default
		configPath := filepath.Join(home, ".gossh.yaml")
		if err := viper.SafeWriteConfigAs(configPath); err != nil {
			panic(err)
		}
	}

	// Unmarshal into struct
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
}

func Get() Config {
	return config
}
