package upload

import (
	"os"
	"fmt"
	"bytes"
	"net/http"
	"github.com/sirupsen/logrus"
	"sync"
)

// Gelf log buffer
const Kb = 1024

// StartUpload bulk uploads the parsed Gelf log entries to Loggly
func StartUpload(token string, logChan chan []byte) {

	logglyUrl := os.Getenv("LOGGLY_URL")
	url := fmt.Sprintf(logglyUrl, token)

	buffer := new(bytes.Buffer)

	for {

		logBytes := <-logChan

		buffer.Write(logBytes)
		buffer.WriteString("\n")

		if buffer.Len() > Kb {
			// upload to loggly
			uploadBulk(url, buffer)
			buffer.Reset()
		}
	}
}

// uploadBulk does an HTTP POST to Loggly
func uploadBulk(url string, buffer *bytes.Buffer) {

	req, err := http.NewRequest("POST", url,
		bytes.NewReader(buffer.Bytes()))

	if err != nil {
		logrus.Errorln("Error creating bulk upload HTTP request: " + err.Error())
		return
	}

	var client *http.Client
	(&sync.Once{}).Do(func() {
		client = &http.Client{}
	})

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		logrus.Errorln("failed to upload to Loggly: " + err.Error())
		return
	}
	logrus.Debugf("sent sent %v bytes to Loggly\n", buffer.Len())
}
