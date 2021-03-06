package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    	string
	LogPath     	string
	WSAddr      	string
	CertFile    	string
	KeyFile     	string

	CLTCPAddr     	string
	CLMaxConnNum  	int

	GSTCPAddr     	string
	GSMaxConnNum  	int

	ConsolePort 	int
	ProfilePath 	string

	DBAddr			string
	ServerNum		int
	ServerRunMaps	[]int
}

func init() {
	data, err := ioutil.ReadFile("conf/center.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
