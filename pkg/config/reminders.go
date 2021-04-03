package config

import (
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Reminder defines the fields of a reminder
type Reminder struct {
	ID          int
	Time        string
	Description string
}

// NewReminder creates a new reminder from the user input
func NewReminder(givenDay string, givenTime string, givenDiscription []string) (r Reminder, err error) {
	t, err := time.Parse("2/1/2006 15:04", givenDay+" "+givenTime)
	if err != nil {
		return r, err
	}

	// TODO add a more robust id generation technique that takes into
	// account current ids.
	r.ID = rand.New(rand.NewSource(time.Now().Unix())).Intn(1000)
	r.Time = t.Format(time.RFC3339)
	r.Description = strings.Join(givenDiscription, " ")

	return r, err
}

func GetReminders() (rs []Reminder) {
	viper.UnmarshalKey("state.reminders", &rs)
	return rs
}

func AddReminder(r Reminder) error {
	var rs []Reminder
	viper.UnmarshalKey("state.reminders", &rs)
	rs = append(rs, r)
	viper.Set("state.reminders", rs)
	return viper.WriteConfig()
}

func SetReminders(rs []Reminder) error {
	viper.Set("state.reminders", rs)

	return viper.WriteConfig()
}
