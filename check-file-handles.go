// Get the number of open files for a process and compare that against /proc/<pid>/limits and alert if
// over the given threshold.
//
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package main

import (
	"flag"
	"fmt"
	"github.com/yieldbot/dracky"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Get the pid for the desired process
func getPid(app string) string {
	pidExp := regexp.MustCompile("[0-9]+")
	termExp := regexp.MustCompile(`pts/[0-9]`)
	appPid := ""

	/// the pid for the binary
	goPid := os.Getpid()
	if dracky.Debug {
		fmt.Printf("golang binary pid: %v\n", goPid)
	}

	psAEF := exec.Command("ps", "-aef")

	out, err := psAEF.Output()
	if err != nil {
		dracky.Check(err)
	}

	psAEF.Start()

	lines := strings.Split(string(out), "\n")

	if !dracky.JavaApp {
		for i := range lines {
			if !strings.Contains(lines[i], strconv.Itoa(goPid)) && !termExp.MatchString(lines[i]) {
				words := strings.Split(lines[i], " ")
				for j := range words {
					if app == words[j] {
						appPid = pidExp.FindString(lines[i])
					}
				}
			}
		}
	} else {
		for i := range lines {
			if strings.Contains(lines[i], app) && !strings.Contains(lines[i], strconv.Itoa(goPid)) && !termExp.MatchString(lines[i]) {
				appPid = pidExp.FindString(lines[i])

			}
		}
	}
	if appPid == "" {
		fmt.Printf("No process with the name " + app + " exists.\n")
		fmt.Printf("If unsure consult the documentation for examples and requirements\n")
		os.Exit(dracky.CONFIG_ERROR)
	}
	return appPid
}

// Calculate if the value is over a threshold
func determineThreshold(limit float64, threshold float64, numFD float64) bool {
	alarm := true
	tLimit := threshold / float64(100) * limit

	if numFD > float64(tLimit) {
		alarm = true
	} else {
		alarm = false
	}
	return alarm
}

// Get the current number of open file handles for the process
func getFileHandles(pid string) (float64, float64, float64) {
	var _s, _h string
	var s, h float64
	limitExp := regexp.MustCompile("[0-9]+")
	filename := `/proc/` + pid + `/limits`
	fdLoc := "/proc/" + pid + "/fd"
	numFD := 0.0

	limits, err := ioutil.ReadFile(filename)
	if err != nil {
		dracky.Check(err)
	}

	lines := strings.Split(string(limits), "\n")
	for i := range lines {
		if strings.Contains(lines[i], "open files") {
			limits := limitExp.FindAllString(lines[i], 2)
			_s = limits[0]
			_h = limits[1]

			s, err = strconv.ParseFloat(_s, 64)
			if err != nil {
				// YELLOW handle error
				fmt.Println(err)
				os.Exit(2)
			}
			h, err = strconv.ParseFloat(_h, 64)
			if err != nil {
				// YELLOW handle error
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}

	files, _ := ioutil.ReadDir(fdLoc)
	for _ = range files {
		numFD++
	}
	if numFD == 0.0 {
		fmt.Printf("There are no open file descriptors for the process, did you use sudo?\n")
		fmt.Printf("If unsure of the use, consult the documentation for examples and requirements\n")
		os.Exit(dracky.PERMISSION_ERROR)
	}
	return s, h, numFD
}

func main() {

	// set commandline flags
	AppPtr := flag.String("app", "sbin/init", "the process name")
	WarnPtr := flag.Int("warn", 75, "the alert warning threshold percentage")
	CritPtr := flag.Int("crit", 75, "the alert critical threshold percentage")
	DebugPtr := flag.Bool("debug", false, "print debugging info instead of an alert")
	JavaAppPtr := flag.Bool("java", false, "java apps process detection is different")

	flag.Parse()
	app := *AppPtr
	warnThreshold := *WarnPtr
	critThreshold := *CritPtr
	dracky.Debug = *DebugPtr
	dracky.JavaApp = *JavaAppPtr

	var appPid string
	var sLimit, hLimit, openFd float64

	if app != "" {
		appPid = getPid(app)
		sLimit, hLimit, openFd = getFileHandles(appPid)
		if dracky.Debug {
			fmt.Printf("warning threshold: %v percent, critical threshold: %v percent\n", warnThreshold, critThreshold)
			fmt.Printf("this is the number of open files at the specific point in time: %v\n", openFd)
			fmt.Printf("app pid is: %v\n", appPid)
			fmt.Printf("This is the soft limit: %v\n", sLimit)
			fmt.Printf("This is the hard limit: %v\n", hLimit)
			os.Exit(0)
		}
		if determineThreshold(hLimit, float64(critThreshold), openFd) {
			fmt.Printf("%v is over %v percent of the the open file handles hard limit of %v\n", app, critThreshold, hLimit)
			os.Exit(2)
		} else if determineThreshold(sLimit, float64(warnThreshold), openFd) {
			fmt.Printf("%v is over %v percent of the open file handles soft limit of %v\n", app, warnThreshold, sLimit)
			os.Exit(1)
		} else {
			fmt.Printf("There was an error calculating the thresholds. Check to make sure everything got convert to a float64.\n")
			fmt.Printf("If unsure of the use, consult the documentation for examples and requirements\n")
			os.Exit(dracky.RUNTIME_ERROR)
		}
	} else {
		fmt.Printf("Please enter a process name to check. \n")
		fmt.Printf("If unsure consult the documentation for examples and requirements\n")
		os.Exit(dracky.CONFIG_ERROR)
	}
}
