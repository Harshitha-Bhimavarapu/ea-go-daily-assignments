package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Job struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

var byteData []byte
var err error

func main() {
	readData()
	var jobs []Job
	json.Unmarshal(byteData, &jobs)

	var jobChannel = make(chan Result)

	go MarkJobStatus(jobs, jobChannel)

	var jobStatus []Result
	for eachJob := range jobChannel {
		jobStatus = append(jobStatus, eachJob)
	}
	WriteDataToFile(jobStatus)
}

func readData() {
	JsonFile, err := os.Open("jobs.json")

	if err != nil {
		fmt.Println(err)
	}

	defer JsonFile.Close()

	byteData, err = ioutil.ReadAll(JsonFile)
}

func MarkJobStatus(jobs []Job, channel chan Result) {
	for i := 0; i < len(jobs); i++ {
		channel <- Result{Id: jobs[i].Id, Status: "SUCCESS"}
	}
	close(channel)
}

func WriteDataToFile(jobStatus []Result) {
	file, _ := json.MarshalIndent(jobStatus, "", " ")
	_ = ioutil.WriteFile("status.json", file, 0644)
}
