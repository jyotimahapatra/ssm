package endpoint

import (
	"fmt"

	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/executers"
	"github.com/aws/amazon-ssm-agent/agent/framework/processor/executer/iohandler"
	"github.com/aws/amazon-ssm-agent/agent/jsonutil"
	"github.com/aws/amazon-ssm-agent/agent/plugins/pluginutil"
	"github.com/aws/amazon-ssm-agent/agent/plugins/runscript"
	messageContracts "github.com/aws/amazon-ssm-agent/agent/runcommand/contracts"
	"github.com/aws/amazon-ssm-agent/agent/task"
)

type endpointPlugin struct {
	context         context.T
	CommandExecuter executers.T
}

func NewEndpointPlugin(context context.T) (*endpointPlugin, error) {
	endpointPlugin := endpointPlugin{
		context:         context,
		CommandExecuter: executers.EndpointExecutor{},
	}

	return &endpointPlugin, nil
}

func (p *endpointPlugin) Execute(config contracts.Configuration, cancelFlag task.CancelFlag, output iohandler.IOHandler) {
	log := p.context.Log()
	log.Infof("endpointPlugin started with configuration %v", config)

	runCommandID, err := messageContracts.GetCommandID(config.MessageId)

	if err != nil {
		log.Warnf("Error extracting RunCommandID from config %v: Error: %v", config, err)
		runCommandID = ""
	}

	log.Infof("endpointPlugin started command %s", runCommandID)

	if cancelFlag.ShutDown() {
		output.MarkAsShutdown()
	} else if cancelFlag.Canceled() {
		output.MarkAsCancelled()
	} else {
		var pluginInput runscript.RunScriptPluginInput
		rawPluginInput := config.Properties
		err := jsonutil.Remarshal(rawPluginInput, &pluginInput)
		if err != nil {
			errorString := fmt.Errorf("invalid format in plugin properties %v;\nerror %v", rawPluginInput, err)
			output.MarkAsFailed(errorString)
			return
		}
		if !pluginutil.ValidatePluginId(pluginInput.ID) {
			pluginInput.ID = ""
		}
		executionTimeout := pluginutil.ValidateExecutionTimeout(log, pluginInput.TimeoutSeconds)
		if len(pluginInput.RunCommand) > 1 {
			errorString := fmt.Errorf("invalid count of commands: %v", pluginInput.RunCommand)
			output.MarkAsFailed(errorString)
		}
		exitCode, err := p.CommandExecuter.NewExecute(p.context, "", output.GetStdoutWriter(), output.GetStderrWriter(), cancelFlag, executionTimeout, pluginInput.RunCommand[0], []string{}, pluginInput.Environment)
		output.SetExitCode(exitCode)
		output.SetStatus(pluginutil.GetStatus(exitCode, cancelFlag))

		if err != nil {
			status := output.GetStatus()
			if status != contracts.ResultStatusCancelled &&
				status != contracts.ResultStatusTimedOut &&
				status != contracts.ResultStatusSuccessAndReboot {
				output.MarkAsFailed(fmt.Errorf("failed to run commands: %v", err))
			}
		}
	}
}
