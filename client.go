package lgpclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// LogData represents a key-value map of logging data.
// Note that the only data accepted is of type string.
type LogData map[string]string

const (
	responseOK     = "OK"
	responseERROR  = "ERROR"
	errorSeperator = "|"
)

var serverAddr, logAddr string

// Use sets the address to use
// The address is in the form of IP/FQDN:port
func Use(addrToUse string) {
	serverAddr = fmt.Sprintf("http://%s/", addrToUse)
	logAddr = serverAddr + "log"
}

// Log accepts a logData, then it logs and if it fails it returns an error
// table - The table to log to.
// logData - Is the actual data to log, in key-value
func Log(table string, logData *LogData) error {
	logQuery := url.Values{}
	for k, v := range *logData {
		logQuery.Add(k, v)
	}
	res, err := http.PostForm(fmt.Sprintf("%s/%s", logAddr, table), logQuery)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	textResponse := string(body)

	if textResponse == responseOK {
		return nil
	}

	if strings.Index(textResponse, responseERROR) == 0 {
		return errors.New(strings.SplitN(textResponse, errorSeperator, 1)[0])
	}

	return fmt.Errorf("Unkown error %s. Status code %d", textResponse, res.StatusCode)

}
