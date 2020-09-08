package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const gowildVersion = "0.0.1"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PrintVersion prints CLI Information
func PrintVersion() {
	fmt.Printf("gowild - https://github.com/havenbarnes/gowild\nv%v\n", gowildVersion)
}

// PrintHelp prints an optional error message and usage information
func PrintHelp(cmd string, invalidCommand bool) {
	fmt.Print("gowild - https://github.com/havenbarnes/gowild\n\n")
	if invalidCommand {
		fmt.Printf("Invalid command: %s\n\n", cmd)
	}
	fmt.Print(`Usage:
  gowild record         begin recording bash commands
  gowild stop           end recording session and create shell script
  gowild help           Print Help (this message) and exit
  gowild version        Print version information and exit`)
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
	return "/usr/local/bin/.gowild"
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
		if strings.Contains(cmd, "gowild record") {
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

// ExecRecord begins a recording session by writing a flag file.
func ExecRecord() {
	if isRecording() {
		fmt.Println("Existing recording session found. Starting new recording session from here. Run 'gowild stop' to end recording")
	} else {
		fmt.Println("Now recording commands... run 'gowild stop' to end recording.\nTo start over, just run 'gowild record' again.")
	}
	setRecording(true)
}

// ExecStop ends a recording session by accessing bash history, writing an output file, and deleting the flag file.
func ExecStop() {
	if !isRecording() {
		fmt.Println("No recording session found. Run 'gowild record' to begin recording commands.")
		return
	}

	writeFile()
	deleteConfig()
}

func writeFile() {
	filename := getOutputFilename()
	path := filepath.Join(getCwd(), filename)
	f, err := os.Create(path)
	check(err)
	_, err = f.WriteString("#!/bin/bash\n" + getBashHistory())
	check(err)
	defer f.Close()

	err = os.Chmod(path, 0777)
	check(err)
	fmt.Printf("Recorded commands written to shell script %s\n", filepath.Join(getCwd(), filename))
}

func getOutputFilename() string {
	defaultName := "gowild.sh"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What should the output file be named? [%s]: ", defaultName)
	scanner.Scan()
	filename := scanner.Text()
	check(scanner.Err())

	if strings.TrimSpace(filename) == "" {
		filename = defaultName
	}
	if filename[len(filename)-3:] == ".sh" {
		return filename
	}
	filename = filename + ".sh"
	return filename
}

func testToggle() string {
	if isRecording() {
		return "stop"
	}
	return "record"
}

func main() {
	var cmd string
	if len(os.Args) < 2 {
		PrintHelp("No command specified.", true)
		return
	} else if len(os.Args) > 2 {
		PrintHelp("Too many arguments provided.", true)
		return
	}
	cmd = os.Args[1]

	if cmd == "record" {
		ExecRecord()
	} else if cmd == "stop" {
		ExecStop()
	} else if cmd == "help" {
		PrintHelp("", false)
	} else if cmd == "version" {
		PrintVersion()
	} else {
		PrintHelp(cmd, true)
	}
}
