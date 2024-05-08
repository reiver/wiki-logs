package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/reiver/go-path"
)

const dotWikiLogsFileName string = ".wiki-logs_output"

var input string
var output string
var verbose bool

func parseFlags() {
	flag.StringVar(&input, "input", "log", "input; the path to the log directory. ex: --input=logs or --output=../../over/here or --output=path/to/log")
	flag.StringVar(&output, "output", "", "output; the path and file-name of the outputted logs file. ex: --output=logs.wiki or --output=my-file.wiki or --output=path/to/logs.wiki — not that you can also specify the output with a "+dotWikiLogsFileName+" file")
	flag.BoolVar(&verbose, "v", false, "verbose")

	flag.Parse()
}

func main() {

	parseFlags()

	{
		if "" == output {
			output = outputFromDotWikiLogs()
		}
		if "" == output {
			fmt.Fprintf(os.Stderr, "ERROR: wiki-logs output file-name not specified\n")
			os.Exit(1)
			return
		}
	}

	if verbose {
		fmt.Printf("Input: %s\n", input)
		fmt.Println("(The \"Input\" is the \"log/\" directory.)")
	}

	if verbose {
		fmt.Printf("Output: %s\n", output)
		fmt.Println("(The \"Output\" is the \"logs.wiki\" file.)")
	}

	var years []fs.DirEntry
	{
		var err error

		var name string = input

		years, err = os.ReadDir(name)
		if nil != err {
			
panic(err)
		}
	}
	if verbose {
		fmt.Printf("Num-Years: %d\n", len(years))
		fmt.Println("(The number of directories and files int the \"Input\" \"log/\" directory.)")
	}

	var logs []string
	for _, year := range years {
		if verbose {
			fmt.Printf("\tYear: %q (is-dir=%t)\n", year.Name(), year.IsDir())
		}
		if !year.IsDir() {
	/////////////// CONTINUE
			continue
		}

		var months []fs.DirEntry
		{
			var err error

			var name string = path.Join(input, year.Name())

			months, err = os.ReadDir(name)
			if nil != err {
				
panic(err)
			}

			for _, month := range months {
				if verbose {
					fmt.Printf("\t\tMonth: %q (is-dir=%t)\n", month.Name(), month.IsDir())
				}
				if !month.IsDir() {
			/////////////// CONTINUE
					continue
				}

				var days []fs.DirEntry
				{
					var err error

					var name string = path.Join(input, year.Name(), month.Name())

					days, err = os.ReadDir(name)
					if nil != err {
						
panic(err)
					}
				}

				for _, day := range days {
					if verbose {
						fmt.Printf("\t\t\tDay: %q (is-dir=%t)\n", day.Name(), day.IsDir())
					}
					if !day.IsDir() {
				/////////////// CONTINUE
						continue
					}

					var timestampfiles []fs.DirEntry
					{
						var err error

						var name string = path.Join(input, year.Name(), month.Name(), day.Name())

						timestampfiles, err = os.ReadDir(name)
						if nil != err {
							
							panic(err)
						}
					}

					for _, timestampfile := range timestampfiles {
						if verbose {
							fmt.Printf("\t\t\tTimestamp: %q (is-dir=%t)\n", timestampfile.Name(), timestampfile.IsDir())
						}
						if timestampfile.IsDir() {
					/////////////// CONTINUE
							continue
						}

						var ext string = path.Ext(timestampfile.Name())

						if ".wiki" != ext {
					/////////////// CONTINUE
							continue
						}

						var timestamp string = timestampfile.Name()
						timestamp = timestamp[:len(timestamp)-len(ext)]

						var logpath string = path.Join("log", year.Name(), month.Name(), day.Name(), timestamp)

						if verbose {
							fmt.Printf("\t\t\tLog-Path: %q\n", logpath)
						}

						logs = append(logs, logpath)
					}
				}
			}
		}
	}

	slices.SortFunc(logs, compare)

	var out *os.File
	{
		var err error

		out, err = os.Create(output)
		if nil != err {
			
panic(err)
		}
	}

	{
		_, err := io.WriteString(out, "wiki/1\n\n§ Logs\n\n")
		if nil != err {
			
panic(err)
		}
	}

	for _, log := range logs {
		if verbose {
			fmt.Printf("Log: %q\n", log)
		}

		{
			_, err := fmt.Fprintf(out, "• [[%s]]\n", log)
			if nil != err {
				
panic(err)
			}
		}
	}
}

func compare(a string, b string) int {
	if a == b {
		return 0
	}

	var topA string = path.Top(a)
	var topB string = path.Top(b)

	if topA == topB {
		var length int = len(topA)

		var newA string = a[length:]
		var newB string = b[length:]

		return compare(newA, newB)
	}

	{
		var numA uint64
		var numB uint64

		var err error

		numA, err = strconv.ParseUint(topA, 10, 64)
		if nil != err {
			
panic(err)
		}

		numB, err = strconv.ParseUint(topB, 10, 64)
		if nil != err {
			
panic(err)
		}

		switch {
		case numA < numB:
			return 1
		case numA > numB:
			return -1
		default:
			return 0
		}
	}

	return strings.Compare(a,b)
}

func outputFromDotWikiLogs() string {
	p, err := os.ReadFile(dotWikiLogsFileName)
	if nil != err {
		return ""
	}

	p = bytes.TrimRight(p, "\n\r\u0085\u2028\u2029")

	return string(p)
}
