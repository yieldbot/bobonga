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
	"github.com/yieldbot/dhuran"
	"io/ioutil"
	// "github.com/shirou/gopsutil/process"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Get the pid for the supplied process
func get_pid(app string) string {
	pid_exp := regexp.MustCompile("[0-9]+")
	go_pid := os.Getpid()
	a_pid := ""
	fmt.Printf("golang binary pid: %v\n", go_pid)

	ps_aef := exec.Command("ps", "-aef")

	out, err := ps_aef.Output()
	if err != nil {
		dhuran.Check(err)
	}

	ps_aef.Start()

	lines := strings.Split(string(out), "\n")
	for i := range lines {
		if strings.Contains(lines[i], app) && !strings.Contains(lines[i], strconv.Itoa(go_pid)) {
			a_pid = pid_exp.FindString(lines[i])
			break
		}
	}
	return a_pid
}

// Calculate if the value is over a threshold
func determine_threshold(limit int, threshold int, num_fd int) bool {
	alarm := true
	t_limit := (float64(threshold) / float64(100)) * float64(limit)

	if float64(num_fd) > float64(t_limit) {
		alarm = true
	} else {
		alarm = false
	}
	return alarm
}

// Get the current number of open file handles for the process
func get_file_handles(pid string) (int, int, int) {
	var _s, _h string
	var s, h int
	limit_exp := regexp.MustCompile("[0-9]+")
	filename := `/proc/` + pid + `/limits`
	fd_loc := "/proc/" + pid + "/fd"
	num_fd := 0

	limits, err := ioutil.ReadFile(filename)
	if err != nil {
		dhuran.Check(err)
	}

	lines := strings.Split(string(limits), "\n")
	for i := range lines {
		if strings.Contains(lines[i], "open files") {
			limits := limit_exp.FindAllString(lines[i], 2)
			_s = limits[0]
			_h = limits[1]

			s, err = strconv.Atoi(_s)
			if err != nil {
				// YELLOW handle error
				fmt.Println(err)
				os.Exit(2)
			}
			h, err = strconv.Atoi(_h)
			if err != nil {
				// YELLOW handle error
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}

	files, _ := ioutil.ReadDir(fd_loc)
	for _ = range files {
		num_fd = num_fd + 1
	}
	return s, h, num_fd
}

func main() {

	// set commandline flags
	AppPtr := flag.String("app", "sbin/init", "the process name")
	WarnPtr := flag.Int("warn", 75, "the alert warning threshold percentage")
	CritPtr := flag.Int("crit", 75, "the alert critical threshold percentage")

	flag.Parse()
	app := *AppPtr
	warn_threshold := *WarnPtr
	crit_threshold := *CritPtr

	var app_pid string
	var s_limit, h_limit, open_fd int

	if app != "" {
		// YELLOW check for a null or nil string
		// need to check for the process better, regex?
		app_pid = get_pid(app)
		s_limit, h_limit, open_fd = get_file_handles(app_pid)
		if determine_threshold(h_limit, crit_threshold, open_fd) {
			fmt.Printf("%v is over %v percent of the the open file handles hard limit of %v\n", app, crit_threshold, h_limit)
			os.Exit(2)
		} else if determine_threshold(s_limit, warn_threshold, open_fd) {
			fmt.Printf("%v is over %v percent of the open file handles soft limit of %v\n", app, warn_threshold, s_limit)
			os.Exit(1)
		} else {
			// YELLOW need to set some other conditions here in case this fails
			fmt.Printf("warning threshold: %v percent, critical threshold: %v percent\n", warn_threshold, crit_threshold)
			fmt.Printf("this is the number of open files at the specific point in time: %v\n", open_fd)
			fmt.Printf("app pid is: %v\n", app_pid)
			fmt.Printf("This is the soft limit: %v\n", s_limit)
			fmt.Printf("This is the hard limit: %v\n", h_limit)
			os.Exit(0)
		}
	} else {
		fmt.Printf("Please enter a process name to check")
		os.Exit(100)
	}
}
