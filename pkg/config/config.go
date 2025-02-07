package config

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
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
	Debug     bool    `mapstructure:"debug"`
	JumpHosts []Host  `mapstructure:"jumphosts"`
	Groups    []Group `mapstructure:"groups"`
}

var config *Config

func Load() {
	redBold := color.New(color.Bold, color.FgRed)

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName(".gossh")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			redBold.Fprintf(os.Stderr, "No config file found. Please create one at %s/.gossh.yaml\n", home)
			os.Exit(0)
		}

		var configParseError viper.ConfigParseError
		if errors.As(err, &configParseError) {
			redBold.Fprintln(os.Stderr, "Config file loaded contains invalid YAML")
			redBold.Fprintf(os.Stderr, "Error: %s\n", configParseError.Error())
			os.Exit(0)
		}

		panic(err)
	}

	// Unmarshal into struct
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	// Validate config
	validationErrors := Get().Validate()
	if len(validationErrors) > 0 {
		redBold.Fprintf(os.Stderr, "• %d errors found in configuration\n\n", len(validationErrors))

		for _, validationErr := range validationErrors {
			fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("⨯"), validationErr)
		}

		os.Exit(0)
	}
}

func Get() Config {
	return *config
}
