package config

import "github.com/spf13/viper"

type Lesson struct {
	// ID is the url id needed to fetch more informatio about the lesson
	// ex https://www.di.uoa.gr/studies/undergraduate/31 where ID = 31
	ID        string
	Name      string
	Code      string
	Ects      string
	Necessity string
	S1        string
	S2        string
	S3        string
	S4        string
	S5        string
	S6        string
}

type Semester struct {
	Number  int
	Lessons []Lesson
}

func GetSemesters() (s []Semester) {
	viper.UnmarshalKey("state.semesters", &s)
	return s
}

func SetSemesters(s []Semester) error {
	viper.Set("state.semesters", s)

	return viper.WriteConfig()
}
