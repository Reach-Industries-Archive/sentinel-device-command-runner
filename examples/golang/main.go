package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Sentinel struct {
	authKey string
	baseUrl string
}

type Command struct {
	Command string `json:"command"`
	Upload  string `json:"upload"`
}

type Output struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func NewSentinel(authKey string) *Sentinel {
	return &Sentinel{
		authKey: authKey,
		baseUrl: os.Getenv("SENTINEL_URL"),
	}
}

func (s *Sentinel) pollCommands(deviceId string) ([]Command, error) {
	commandsUrl := fmt.Sprintf("%s/device/commands", s.baseUrl)

	req, err := http.NewRequest("GET", commandsUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", s.authKey)
	req.Header.Add("DeviceId", deviceId)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commands []Command
	err = json.NewDecoder(resp.Body).Decode(&commands)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func (s *Sentinel) uploadOutput(url string, stdout []byte, stderr []byte) error {
	output := Output{
		Stdout: string(stdout),
		Stderr: string(stderr),
	}

	payload, err := json.Marshal(output)
	if err != nil {
		return err
	}

	_, err = http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	return nil
}

func execWithTimeout(command string, timeout time.Duration) ([]byte, []byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", command)
	stdout, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		return nil, nil, ctx.Err()
	}

	if err != nil {
		return nil, nil, err
	}

	return stdout, nil, nil
}

func main() {
	authKey := os.Getenv("SENTINEL_AUTH_KEY")
	deviceId := os.Getenv("SENTINEL_DEVICE_ID")

	sentinel := NewSentinel(authKey)

	for {
		commands, err := sentinel.pollCommands(deviceId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, command := range commands {
			fmt.Println("Executing command: ", command.Command)
			stdout, stderr, err := execWithTimeout(command.Command, 10*time.Second)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Uploading output...")
			err = sentinel.uploadOutput(command.Upload, stdout, stderr)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("No commands left to execute.")
		time.Sleep(60 * time.Second)
	}
}
