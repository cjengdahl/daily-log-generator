package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"time"
)

// default is $HOME/daily-logs
const rootDirEnvVarName = "DAILY_LOG_DIRECTORY"

func main() {

	daysForward := flag.Int("o", 0, "number of days offset to create log for")
	flag.Usage = func() {
		fmt.Println("Usage: dlg [-o <days offset>]")
	}
	flag.Parse()
	currentTime := time.Now().AddDate(0, 0, *daysForward)
	year := currentTime.Format("2006")
	month := currentTime.Format("01")
	day := currentTime.Format("02")

	// create path
	rootDirPath := os.Getenv(rootDirEnvVarName)
	if rootDirPath == "" {
		usr, err := user.Current()
		if err != nil {
			printErrorAndExit(fmt.Errorf("failed to get current user: %s", err.Error()))
		}
		rootDirPath = path.Join(usr.HomeDir, "daily-logs")
	}

	// create directories (intermediate as required)
	dirPath := path.Join(rootDirPath, year, month)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		printErrorAndExit(fmt.Errorf("failed to create directories: %s", err.Error()))
	}

	filePath := path.Join(dirPath, day+".md")

	// create file, if it doesn't exist
	file, err := os.OpenFile(
		filePath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		printErrorAndExit(fmt.Errorf("failed to create file: %s", err.Error()))
	}
	defer file.Close()

	dayWritten := currentTime.Format("Monday")
	monthWritten := currentTime.Format("January")
	dayWithSuffix := ordinal(currentTime.Format("2"))

	// write file header (formatted date)
	_, err = file.WriteString(
		fmt.Sprintf("# %s, %s %s, %s\n",
			dayWritten, monthWritten, dayWithSuffix, year))
	if err != nil {
		printErrorAndExit(fmt.Errorf("error writing to file: %s", err.Error()))
	}

}

// ordinal returns the numstr + ordinal suffix for the number
func ordinal(numStr string) string {

	n, err := strconv.Atoi(numStr)
	if err != nil || n < 1 || n > 31 {
		printErrorAndExit(fmt.Errorf("invalid ordinal input: %s", numStr))
	}

	suffix := "th"
	if n%100 != 11 && n%100 != 12 && n%100 != 13 {
		switch n % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}

	return numStr + suffix

}

func printErrorAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
