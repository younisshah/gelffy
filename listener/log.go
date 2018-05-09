package listener

import (
	"net"
	"github.com/sirupsen/logrus"
	"encoding/json"
	"github.com/pkg/errors"
)

// ListenForLogs listens for the Gelf logs from the UDP server
// parse them to get the 'level' and 'msg' fields and sends them to the channel.
func ListenForLogs(conn *net.UDPConn, logChan chan []byte) {

	buffer := make([]byte, 10240) // 10 Kb buffer
	log := make(map[string]interface{})

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			logrus.Errorln("failed to read UDP packet into buffer: %v\n", err.Error())
			continue
		}

		err = json.Unmarshal(buffer[0:n], &log)
		if err != nil {
			logrus.Errorln("failed to unmarshal log entry to JSON")
			logrus.Errorln("Error: " + err.Error())
			continue
		}

		shortMessageLogBytes, err := transform(log)

		if err == nil {
			logChan <- shortMessageLogBytes
		}
	}
}

// transform parses the gelf log entries to a byte slice containing 'level' and 'msg' log entries
func transform(log map[string]interface{}) ([]byte, error) {

	if shortMsg, ok := log["short_message"]; ok {
		smStr := shortMsg.(string)

		smLogEntry := make(map[string]interface{})

		err := json.Unmarshal([]byte(smStr), &smLogEntry)

		if err != nil {
			logrus.Printf("failed to parse short_message: %v\n", err.Error())
			return nil, errors.New("failed to parse short_message property")
		}

		if smLogEntry != nil {
			log["msg"] = smLogEntry["msg"].(string)
			log["level"] = smLogEntry["level"].(string)
			delete(log, "short_message")
		} else {
			logrus.Println("failed to parse log entry short_message: " + smStr)
			return nil, errors.New("failed to parse log entry short_message")
		}
		logrus.Println(log["msg"].(string))

		return json.Marshal(log)
	} else {
		logrus.Errorln("short_message field missing from log entry")
		return nil, errors.New("short_message field missing from log entry")
	}
}
