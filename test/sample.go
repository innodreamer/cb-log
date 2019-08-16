package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/cloud-barista/cb-log"
)

var cblogger *logrus.Logger

type DBCONN struct {
	name string
}

func init() {
	// cblog is a global variable.
	cblogger = cblog.GetLogger("CB-SPIDER")
}

func main() {

	for {
		cblogger.Info("start.........")

		err := createUser1("newUser")
		if err != nil {
			cblogger.Error(err)
		}

		err = createUser2("newUser")
		if err != nil {
			cblogger.Error(err)
		}

		cblogger.Info("end.........")

		time.Sleep(time.Second*2)
		fmt.Print("\n")
	}
}

func createUser1(newUser string) error {
	cblogger.Info("start creating user.")

	var db *DBCONN
	db = new(DBCONN)
	if db == nil {
		cblogger.Error("DBMS Session is closed!!")
		//panic(fmt.Errorf("panic by ..... !"))
	}
	
	isExist, err := checkUser(newUser)
	if isExist {
		return err
	}

	cblogger.Info("finish creating user.")
	return nil
}

func createUser2(newUser string) error {
        cblogger.Info("start creating user.")

        var db *DBCONN
//      db := new(DBCONN)
        if db == nil {
                cblogger.Error("DBMS Session is closed!!")
                //panic(fmt.Errorf("panic by ..... !"))
        }

        isExist, err := checkUser(newUser)
        if isExist {
                return err
        }

        cblogger.Info("finish creating user.")
        return nil
}


func checkUser(user string) (bool, error) {
	return false, fmt.Errorf("%s: already existed User!!", user)
}

