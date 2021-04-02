package config

import (
	"os"
	"path/filepath"

	"github.com/FotiadisM/ditctl/pkg/reminder"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// ConfigName containes the name of the configuration file
	ConfigName = "config"
	// ConfigType containes the type of the configuration file
	ConfigType = "yaml"
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
	ConfigDirPath = home + "/.ditctl/"

	viper.SetDefault("credential.username", "")
	viper.SetDefault("credential.password", "")
	viper.SetDefault("context", "")
	viper.SetDefault("state.reminders", []reminder.Reminder{})
}

type Config struct {
}

func CreateEmpty(path string) {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		cobra.CheckErr(err)
	}

	f, err := os.Create(path + ConfigName)
	if err != nil {
		cobra.CheckErr(err)
	}
	defer f.Close()

	// change file permissions for security
	if err := f.Chmod(0600); err != nil {
		cobra.CheckErr(err)
	}

	if err := viper.WriteConfig(); err != nil {
		cobra.CheckErr(err)
	}
}

func GetReminders() (rs []reminder.Reminder) {
	viper.UnmarshalKey("state.reminders", &rs)
	return rs
}

func AddReminder(r reminder.Reminder) error {
	var rs []reminder.Reminder
	viper.UnmarshalKey("state.reminders", &rs)
	rs = append(rs, r)
	viper.Set("state.reminders", rs)
	return viper.WriteConfig()
}

func SetReminders(rs []reminder.Reminder) error {
	viper.Set("state.reminders", rs)

	return viper.WriteConfig()
}
