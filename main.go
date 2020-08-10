package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const activeRecordingFlag = "gowild_recording_flag"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printHelp(cmd string) {
	fmt.Printf(`gowild - https://github.com/havenbarnes/gowild

Invalid command: %v

Usage:
  gowild start          begin recording bash commands
  gowild stop           end recording session and create shell script
  gowild help           Print Help (this message) and exit
  gowild version        Print version information and exit
	`, cmd)
}

var cwd string

func getCwd() string {
	if cwd == "" {
		path, err := os.Getwd()
		check(err)
		cwd = path
	}
	return cwd
}

func getConfigPath() string {
	return filepath.Join(getCwd(), ".gowild")
}

func deleteConfig() {
	os.Remove(getConfigPath())
}

func getHistoryPath() string {
	home := os.Getenv("HOME")
	if _, err := os.Stat(filepath.Join(home, ".zsh_history")); os.IsNotExist(err) {
		return filepath.Join(home, ".bash_history")
	}
	return filepath.Join(home, ".zsh_history")
}

func reverse(strings []string) []string {
	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}

func getBashHistory() string {
	dat, err := ioutil.ReadFile(getHistoryPath())
	check(err)
	fullHistory := strings.Split(string(dat), "\n")
	historyLen := len(fullHistory)

	var lastIndex int
	for i := range fullHistory {
		index := historyLen - 1 - i
		cmd := fullHistory[index]
		fullHistory[index] = strings.Join(strings.Split(fullHistory[index], ";")[1:], "")
		if strings.Contains(cmd, "go run main.go record") || strings.Contains(cmd, "gowild record") {
			lastIndex = index
			break
		}
	}

	recordedHistory := fullHistory[lastIndex+1 : historyLen-2]
	return strings.Join(recordedHistory, "\n")
}

func isRecording() bool {
	dat, err := ioutil.ReadFile(getConfigPath())
	if err != nil || dat == nil {
		return false
	}
	return string(dat) == "1"
}

func setRecording(recording bool) {
	var value string
	if recording {
		value = "1"
	} else {
		value = "0"
	}

	f, err := os.Create(filepath.Join(getConfigPath()))
	check(err)
	_, err = f.WriteString(value)
	check(err)
	defer f.Close()
}

func execRecord() {
	if isRecording() {
		fmt.Println("Existing recording session found. Starting new recording session from here. Run 'gowild stop' to end recording")
	} else {
		fmt.Println("Now recording commands... run 'gowild stop' to end recording")
	}
	setRecording(true)
}

func execStop() {
	if !isRecording() {
		fmt.Println("No recording session found. Run 'gowild record' to begin recording commands.")
		return
	}

	writeFile()
	deleteConfig()
	fmt.Println("Recorded commands written to shell script output.sh")
}

func writeFile() {
	path := filepath.Join(getCwd(), "output.sh")
	f, err := os.Create(path)
	check(err)
	_, err = f.WriteString("#!/bin/bash\n" + getBashHistory())
	check(err)
	defer f.Close()

	err = os.Chmod(path, 0777)
	check(err)
}

func testToggle() string {
	if isRecording() {
		return "stop"
	}
	return "record"
}

func main() {
	var cmd string
	if len(os.Args) != 2 {
		printHelp("No command specified.")
		return
	} else {
		cmd = os.Args[1]
	}

	if cmd == "record" {
		execRecord()
	} else if cmd == "stop" {
		execStop()
	} else {
		printHelp(cmd)
	}
}
