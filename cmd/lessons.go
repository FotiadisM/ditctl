/*
Copyright © 2021 Fotiadis Michail <fotiadis.michalis20@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/FotiadisM/ditctl/pkg/config"
	"github.com/FotiadisM/ditctl/pkg/parser"
)

// lessonsCmd represents the lessons command
var lessonsCmd = &cobra.Command{
	Use:   "lessons [lessonID]",
	Short: "Retrieve informations about University lessons",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var sems []config.Semester
		sems = config.GetSemesters()

		if len(sems) == 0 {
			var err error
			sems, err = refreshCache()
			if err != nil {
				cobra.CheckErr(err)
			}
		}

		printNormal(sems)
	},
}

func init() {
	rootCmd.AddCommand(lessonsCmd)

	lessonsCmd.Flags().BoolP("verbose", "v", false, "verbose output")
	lessonsCmd.Flags().BoolP("refresh", "r", false, "fetch lesson from the internet and update the cache")
	lessonsCmd.Flags().IntP("semester", "s", -1, "only retrieve lessons of the given semester")
	lessonsCmd.Flags().Lookup("semester").DefValue = "all"
	lessonsCmd.Flags().StringP("output", "o", "table", "output mode (table, yaml, json)")
}

func refreshCache() (sems []config.Semester, err error) {
	sems, err = parser.FetchLessons()
	if err != nil {
		return
	}

	if err = config.SetSemesters(sems); err != nil {
		return
	}

	return
}

func printNormal(sems []config.Semester) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Render()
}

func printVerbose(sems []config.Semester) {

}

func printYAML(sems []config.Semester) {

}

func printJSON(sems []config.Semester) {

}
