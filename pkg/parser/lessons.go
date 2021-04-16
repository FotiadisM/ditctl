package parser

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/FotiadisM/ditctl/pkg/config"
)

func FetchLessons() (sems []config.Semester, err error) {
	res, err := http.Get("https://www.di.uoa.gr/studies/undergraduate/courses")
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	doc.Find(".view-content .table-striped").Each(func(i int, s *goquery.Selection) {
		sems = append(sems, parseSem(i, s))
	})

	return
}

func parseSem(semNumber int, s *goquery.Selection) config.Semester {
	sem := config.Semester{Number: semNumber}

	s.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		sem.Lessons = append(sem.Lessons, parseLess(s))
	})

	return sem
}

func parseLess(s *goquery.Selection) config.Lesson {
	l := config.Lesson{}

	// some fields have extra white spaces that need to be removed
	s.Find("td").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			l.Name = s.Children().Text()
		case 1:
			l.Code = strings.Join(strings.Fields(s.Text()), "")
		case 2:
			l.Ects = strings.Join(strings.Fields(s.Text()), "")
		case 3:
			l.Necessity = strings.Join(strings.Fields(s.Text()), "")
		case 4:
			l.S1 = strings.Join(strings.Fields(s.Text()), "")
		case 5:
			l.S2 = strings.Join(strings.Fields(s.Text()), "")
		case 6:
			l.S3 = strings.Join(strings.Fields(s.Text()), "")
		case 7:
			l.S4 = strings.Join(strings.Fields(s.Text()), "")
		case 8:
			l.S5 = strings.Join(strings.Fields(s.Text()), "")
		case 9:
			l.S6 = strings.Join(strings.Fields(s.Text()), "")
		}
	})

	return l
}
