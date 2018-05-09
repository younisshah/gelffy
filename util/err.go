package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

func CheckError(err error) {
	if err != nil {
		logrus.Println("Error: " + err.Error())
		os.Exit(-1)
	}
}
