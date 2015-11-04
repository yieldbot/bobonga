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
//	"encoding/json"
	"flag"
	"fmt"
	"github.com/yieldbot/dhuran"
//	"github.com/yieldbot/dracky"
//	"github.com/olivere/elastic"
	// "io/ioutil"
  // "github.com/shirou/gopsutil/process"
  "strconv"
  "regexp"
  "strings"
	"os"
  "os/exec"
//	"time"
)

// Get the pid for the supplied process
func get_pid(app string) string {
  pid_exp := regexp.MustCompile("[0-9]+")
  go_pid := os.Getpid()
  var a_pid string
  fmt.Printf("golang binary pid: %v\n", go_pid)

  ps_aef := exec.Command("ps", "-aef")

  out, err := ps_aef.Output()
  if err != nil {
    dhuran.Check(err)
  }

  ps_aef.Start()

  lines := strings.Split(string(out), "\n")
  for i := range lines {
    if (strings.Contains(lines[i], app) && ! strings.Contains(lines[i], strconv.Itoa(go_pid))){
      // fmt.Printf("%v\n", lines[i])
      a_pid = pid_exp.FindString(lines[i])
      return a_pid
      // matty := strings.Split(lines[i]," ")
    // for j := range matty {
    //   fmt.Printf("%v\n", matty[j])
    // }
    //     return lines[i]
    }
  }
  return ""
}

// Calculate if the value is over a threshold
// func determine_threshold(val int, threshold int) bool {
//
// }
//
// // Get the current number of open file handles for the process
// func get_file_handles(pid string) int {
//
// }

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

	if app != "" {
    // YELLOW check for a null or nil string
    // need to check for the process better, regex?
    app_pid = get_pid(app)

  } else {
    fmt.Printf("Please enter a process name to check")
		os.Exit(100)
    }

    fmt.Printf("warning threshold: %v, critical threshold: %v\n", warn_threshold, crit_threshold)
    fmt.Printf("app pid is: %v\n", app_pid)




	// 	if err != nil {
	// 		dhuran.Check(err)
	// 	}
	// 	err = json.Unmarshal(user_input, &user_event)
	// 	if err != nil {
	// 		dhuran.Check(err)
	// 	}
	// 	es_type = "user"
	// } else if (rd_stdin == false) && (input_file == "") {
	// 	fmt.Printf("Please enter a file to read from")
	// 	os.Exit(1)
	// } else {
	// 	sensu_event = sensu_event.Acquire_sensu_event()
	// }
  //
	// // Create a client
	// client, err := elastic.NewClient(
	// 	elastic.SetURL("http://" + es_host + ":" + es_port),
	// )
	// if err != nil {
	// 	dhuran.Check(err)
	// }
  //
	// // Check to see if the index exists and if not create it
	// if client.IndexExists == nil { // need to test to make sure this does what I want
	// 	_, err = client.CreateIndex(es_index).Do()
	// 	if err != nil {
	// 		dhuran.Check(err)
	// 	}
	// }
  //
	// // Create an Elasticsearch document. The document type will define the mapping used for the document.
	// doc := make(map[string]string)
	// var doc_id string
	// switch es_type {
	// case "sensu":
	// 	doc_id = dracky.Event_name(sensu_event.Client.Name, sensu_event.Check.Name)
	// 	doc["monitored_instance"] = sensu_event.Acquire_monitored_instance()
	// 	doc["sensu_client"] = sensu_event.Client.Name
	// 	doc["incident_timestamp"] = time.Unix(sensu_event.Check.Issued, 0).Format(time.RFC822Z)
	// 	doc["check_name"] = dracky.Create_check_name(sensu_event.Check.Name)
	// 	doc["check_state"] = dracky.Define_status(sensu_event.Check.Status)
	// 	doc["sensu_env"] = dracky.Define_sensu_env(sensu_env.Sensu.Environment)
	// 	doc["instance_address"] = sensu_event.Client.Address
	// 	doc["check_state_duration"] = dracky.Define_check_state_duration()
	// case "user":
	// 	doc["product"] = user_event.Product
	// 	doc["data"] = user_event.Data
	// 	doc["timestamp"] = time.Unix(sensu_event.Check.Issued, 0).Format(time.RFC822Z) // dracky.Set_time(user_event.Timestamp)
	// default:
	// 	fmt.Printf("Type is not correctly set")
	// 	os.Exit(2)
	// }
  //
	// // Add a document to the Elasticsearch index
	// _, err = client.Index().
	// 	Index(es_index).
	// 	Type(es_type).
	// 	Id(doc_id).
	// 	BodyJson(doc).
	// 	Do()
	// if err != nil {
	// 	dhuran.Check(err)
	// }
  //
	// // Log a successful document push to stdout. I don't add the id here as some id's are fixed but
	// // the user has the ability to autogenerate an id if they don't want to provide one.
	// fmt.Printf("Record added to ES\n")
}
