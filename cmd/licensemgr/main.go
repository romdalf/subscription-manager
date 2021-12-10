package main

import (
	// "fmt"
	// "context"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	// "os/signal"
	// "syscall"
	// "strconv"
	// "time"

	"k8s.io/klog/v2"
)

var (
	listenAddr    = flag.String("listen-addr", "http://localhost:1337", "API endpoint address")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	klog.Info("Subscription Manager service is starting (version: ", "[dev]", ")")

	// if *logFormatJSON {
	// 	klog.SetLogger(json.JSONLogger)
	// }
	// ctx := withShutdownSignal(context.Background())

	// klog.InfoS("Starting Subscription Manager service", "version", version.BuildVersion)

	// // future func getID
	klog.Info("Getting data from subscription API endpoint...")
	// res, error := http.get(*listenAddr)
	apiEndPoint := *listenAddr + "/api/subscriptions"

	response, error := http.Get(apiEndPoint)
	if error != nil {
		klog.Errorln(error)
		os.Exit(1)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		klog.Info("API endpoint response is: ", string(data))
	}

}
