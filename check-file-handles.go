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
  "strconv"
	"os"
  "os/exec"
//	"time"
)

// Get the pid for the supplied process
func get_pid(app string) string {
  go_pid := strconv.Itoa(os.Getpid())
  fmt.Printf(go_pid)
  ps_aef := exec.Command("ps", "-aef")
  grep_find := exec.Command("grep", app)
  grep_exclude := exec.Command("grep", "-v", go_pid)


  outPipe, err := ps_aef.StdoutPipe()
  if err != nil {
    dhuran.Check(err)
  }

  ps_aef.Start()
  grep_find.Stdin = outPipe

  outGrep, err := grep_find.StdoutPipe()
  if err != nil {
    dhuran.Check(err)
  }
fmt.Printf("one\n")
  // fmt.Printf("%v",outGrep)
  grep_exclude.Stdin = outGrep
fmt.Printf("two\n")
  out, err := grep_exclude.Output()
  fmt.Printf("three\n")
  // fmt.Printf(out)
  if err != nil {

    fmt.Printf("%v\n", err)
    // dhuran.Check(err)
  }
  fmt.Printf("four\n")
  defer outPipe.Close()
  defer outGrep.Close()

fmt.Printf(go_pid)

  return string(out)
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
	// PidPtr := flag.String("pid", "1", "the pid for the process you wish to check")
	AppPtr := flag.String("app", "sbin/init", "the process name")
	WarnPtr := flag.Int("warn", 75, "the alert warning threshold percentage")
	CritPtr := flag.Int("crit", 75, "the alert critical threshold percentage")

	flag.Parse()
	// PidPtr := *PidPtr
	app := *AppPtr
	warn_threshold := *WarnPtr
	crit_threshold := *CritPtr

	// I don't want to call these if they are not needed
	// sensu_event := new(dracky.Sensu_Event)
	// user_event := new(dracky.User_Event)
	//t_format := *timePtr

	// sensu_env := dracky.Set_sensu_env()

	//   if t_format != "" {
	//     // get the format of the time
	//     es_index = es_index + t_format
	//   }

	if app != "" {
    pid := get_pid(app)
    fmt.Printf(pid)
  } else {
    fmt.Printf("Please enter a process name to check")
		os.Exit(100)
    }

    fmt.Printf(out)
    fmt.Printf("%v , %v\n", warn_threshold, crit_threshold)



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
