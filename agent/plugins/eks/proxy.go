package eks

import (
	"fmt"
	"net/http"

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
		for _, c := range pluginInput.RunCommand {
			log.Infof("Eks plugin command received %s", c)
		}

		resp, err := http.Get("http://localhost:8000")
		log.Infof("Eks plugin query succeeded")
		if err != nil || resp.StatusCode > 300 {
			output.SetExitCode(1)
			output.SetStatus(pluginutil.GetStatus(1, cancelFlag))
			status := output.GetStatus()
			if status != contracts.ResultStatusCancelled &&
				status != contracts.ResultStatusTimedOut &&
				status != contracts.ResultStatusSuccessAndReboot {
				output.MarkAsFailed(fmt.Errorf("failed to run commands: %v", err))
			}
			return
		}
		output.SetExitCode(0)
		output.SetStatus(pluginutil.GetStatus(0, cancelFlag))
	}
}
