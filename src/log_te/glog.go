package log_te

import (
	"flag"
	"github.com/golang/glog"
)

func UserGlog() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("This is a info log")
	glog.Warning("This is a warning log")
	glog.Error("This is a error log")
	//glog.Fatal("This is a fatal log")
}
