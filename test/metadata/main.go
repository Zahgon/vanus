package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("etcd.metadata")
	if err != nil {
		panic("failed to open metadata file, " + err.Error())
	}
	reader := bufio.NewReader(f)
	eventbus := map[string]map[string]interface{}{}
	subscriptions := map[string]map[string]interface{}{}
	for {
		k, err := readline(reader)
		if err == io.EOF {
			break
		}
		key := string(k)
		value, err := readline(reader)
		if err != nil {
			panic("corrupted data, " + string(key))
		}
		m := map[string]interface{}{}
		if strings.HasPrefix(key, "/vanus/vanus/internal/resource/eventbus") {
			if err = json.Unmarshal(value, &m); err != nil {
				panic(fmt.Sprintf("failed to unmarshall data, key: %s, valus: %s", key, string(value)))
			}
			eventbus[key] = m
		} else if strings.HasPrefix(key, "/vanus/trigger/subscriptions") {
			if err = json.Unmarshal(value, &m); err != nil {
				panic(fmt.Sprintf("failed to unmarshall data, key: %s, valus: %s", key, string(value)))
			}
			subscriptions[key] = m
		}
	}
	_ = f.Close()
	f, err = os.Create("vanus-restore.sh")
	if err != nil {
		panic("failed to create script file, " + err.Error())
	}
	generateEventbusCreateScripts(f, eventbus)
	generateSubscriptions(f, subscriptions)
	_ = f.Close()
}

func generateEventbusCreateScripts(w io.Writer, m map[string]map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func generateSubscriptions(w io.Writer, m map[string]map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func readline(reader *bufio.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
