package eks

import (
	"bytes"
	ctx "context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/framework/processor/executer/iohandler"
	"github.com/aws/amazon-ssm-agent/agent/jsonutil"
	"github.com/aws/amazon-ssm-agent/agent/plugins/pluginutil"
	"github.com/aws/amazon-ssm-agent/agent/plugins/runscript"
	messageContracts "github.com/aws/amazon-ssm-agent/agent/runcommand/contracts"
	"github.com/aws/amazon-ssm-agent/agent/task"
)

// runShellPlugin is the type for the RunShellScript plugin and embeds Plugin struct.
type runEksPlugin struct {
	context context.T
}

// NewRunShellPlugin returns a new instance of the SHPlugin.
func NewEksPlugin(context context.T) (*runEksPlugin, error) {
	eksplugin := runEksPlugin{
		context: context,
	}

	return &eksplugin, nil
}

func (p *runEksPlugin) Execute(config contracts.Configuration, cancelFlag task.CancelFlag, output iohandler.IOHandler) {
	log := p.context.Log()
	log.Infof("Eks plugin started with configuration %v", config)

	runCommandID, err := messageContracts.GetCommandID(config.MessageId)
	if err != nil {
		log.Warnf("Error extracting RunCommandID from config %v: Error: %v", config, err)
		runCommandID = ""
	}
	log.Infof("Eks plugin command received commandId: %s", runCommandID)

	if cancelFlag.ShutDown() {
		output.MarkAsShutdown()
	} else if cancelFlag.Canceled() {
		output.MarkAsCancelled()
	} else {
		var pluginInput runscript.RunScriptPluginInput
		err := jsonutil.Remarshal(config.Properties, &pluginInput)
		if err != nil {
			errorString := fmt.Errorf("invalid format in plugin properties %v;\nerror %v", config.Properties, err)
			output.MarkAsFailed(errorString)
			return
		}

		failedCode := 0
		failedMessage := ""
		for _, c := range pluginInput.RunCommand {
			log.Infof("Eks plugin command received %s", c)
			str, err := base64.StdEncoding.DecodeString(c)
			if err != nil {
				log.Infof("Eks plugin base64 decode failed %v")
				failedCode = 1
				failedMessage = err.Error()
				break
			}
			log.Infof("Eks plugin decoded base64 %s", str)
			var p Payload
			if err := json.Unmarshal([]byte(str), &p); err != nil {
				log.Infof("Eks plugin json unmarshal failed: %v", err)
				failedCode = 1
				failedMessage = err.Error()
				break
			}
			var jsonStr = fmt.Sprintf(`{"piezoDocument":"%s"}`, p.Payload)
			fmt.Printf("Sending to %s: %s", p.Command, p.Payload)
			t := http.DefaultTransport.(*http.Transport).Clone()
			t.DialContext = func(ctx ctx.Context, network, addr string) (net.Conn, error) {
				return net.Dial("tcp", "172.16.84.119:13734")
			}
			var httpClient = &http.Client{
				Timeout:   time.Second * 10,
				Transport: t,
			}
			fmt.Println(jsonStr)
			resp, err := httpClient.Post("http://172.16.84.119/execute", "application/json", bytes.NewBuffer([]byte(jsonStr)))
			if err != nil {
				failedCode = -1
				if resp != nil {
					failedCode = resp.StatusCode
					log.Infof("Eks plugin query response: %v", resp)
				}

				failedMessage = err.Error()
				log.Infof("Eks plugin query failed: %v", err)
				break
			}
			if resp.StatusCode > 300 {
				failedCode = resp.StatusCode
				body := make([]byte, 0)
				resp.Body.Read(body)
				failedMessage = fmt.Sprintf("command %s, failed with response: %s, %d", c, body, failedCode)
				log.Infof("Eks plugin query failed")
				break
			}
			log.Infof("Eks plugin query succeeded")
		}

		if failedCode != 0 {
			output.SetExitCode(failedCode)
			output.SetStatus(pluginutil.GetStatus(failedCode, cancelFlag))
			status := output.GetStatus()
			if status != contracts.ResultStatusCancelled &&
				status != contracts.ResultStatusTimedOut &&
				status != contracts.ResultStatusSuccessAndReboot {
				output.MarkAsFailed(fmt.Errorf(failedMessage))
			}
			return
		}
		output.SetExitCode(0)
		output.SetStatus(pluginutil.GetStatus(0, cancelFlag))
	}
}

type Payload struct {
	Command string `json:"command,omitempty"`
	Payload string `json:"payload,omitempty"`
}
