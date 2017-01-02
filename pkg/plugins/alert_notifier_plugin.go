package plugins

import (
	"encoding/json"
	"fmt"
	"path"
	"plugin"

	"github.com/grafana/grafana/pkg/services/alerting"
)

type AlertNotifierPlugin struct {
	PluginBase
}

func (fp *AlertNotifierPlugin) init() {
}

func (p *AlertNotifierPlugin) Load(decoder *json.Decoder, pluginDir string) error {
	if err := decoder.Decode(&p); err != nil {
		return err
	}

	fullPath := path.Join(pluginDir, p.Id+".so")
	plug, err := plugin.Open(fullPath)
	if err != nil {
		return err
	}

	pluginInit, err := plug.Lookup("New")

	if err != nil {
		return err
	}

	casted, ok := pluginInit.(func() error)
	if !ok {
		return fmt.Errorf("Could not cast plugin constructor")
	}

	casted()

	notifier, err := casted()
	if err != nil {
		return err
	}

	alerting.RegisterNotifier(p.Id, notifier)

	if err != nil {
		return err
	}

	if err := p.registerPlugin(pluginDir); err != nil {
		return err
	}
	AlertNotifierPlugins[p.Id] = p

	return nil
}
