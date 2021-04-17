package config

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// ConfigName containes the name of the configuration file
	ConfigName = "config"
	// ConfigType containes the type of the configuration file
	ConfigType = "yaml"
	// ConfigFolder containes the name of the folder that the
	// config file is located
	ConfigFolder = ".ditctl"
)

var (
	// ConfigDirPath is the DIR path that containes the config file
	ConfigDirPath string
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		cobra.CheckErr(err)
	}
	ConfigDirPath = filepath.Join(home, ConfigFolder)

	viper.SetDefault("credentials.username", "")
	viper.SetDefault("credentials.password", "")
	viper.SetDefault("context", "")
	viper.SetDefault("state.reminders", []Reminder{})
	viper.SetDefault("state.semesters", []Semester{})
}

func CreateEmpty(path string) (err error) {
	if err = os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// change file permissions for security
	if err = f.Chmod(0600); err != nil {
		fmt.Println("3")
		return
	}

	// create an empty defailt config
	if err = viper.WriteConfig(); err != nil {
		return
	}

	return
}
