package main

func ExamplePrintVersion() {
	PrintVersion()
	// Output: gowild - https://github.com/havenbarnes/gowild
	//v0.0.1
}

func ExamplePrintHelp() {
	PrintHelp("foo", false)
	// Output: gowild - https://github.com/havenbarnes/gowild
	//
	// Usage:
	//   gowild record         begin recording bash commands
	//   gowild stop           end recording session and create shell script
	//   gowild help           Print Help (this message) and exit
	//   gowild version        Print version information and exit
}

func ExamplePrintHelp_invalidCommand() {
	PrintHelp("foo", true)
	// Output: gowild - https://github.com/havenbarnes/gowild
	//
	// Invalid command: foo
	//
	// Usage:
	//   gowild record         begin recording bash commands
	//   gowild stop           end recording session and create shell script
	//   gowild help           Print Help (this message) and exit
	//   gowild version        Print version information and exit
}
