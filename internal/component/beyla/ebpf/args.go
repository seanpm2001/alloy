package beyla

import (
	"github.com/grafana/alloy/internal/component/discovery"
	"github.com/grafana/alloy/internal/component/otelcol"
)

// Arguments configures the Beyla component.
type Arguments struct {
	Port           string                     `alloy:"open_port,attr,optional"`
	ExecutableName string                     `alloy:"executable_name,attr,optional"`
	Routes         Routes                     `alloy:"routes,block,optional"`
	Attributes     Attributes                 `alloy:"attributes,block,optional"`
	Discovery      Discovery                  `alloy:"discovery,block,optional"`
	Output         *otelcol.ConsumerArguments `alloy:"output,block,optional"`
}

type Exports struct {
	Targets []discovery.Target `alloy:"targets,attr"`
}

type Routes struct {
	Unmatch        string   `alloy:"unmatched,attr,optional"`
	Patterns       []string `alloy:"patterns,attr,optional"`
	IgnorePatterns []string `alloy:"ignored_patterns,attr,optional"`
	IgnoredEvents  string   `alloy:"ignore_mode,attr,optional"`
}

type Attributes struct {
	Kubernetes KubernetesDecorator `alloy:"kubernetes,block"`
}

type KubernetesDecorator struct {
	Enable string `alloy:"enable,attr"`
}

type Services []Service

type Service struct {
	Name      string `alloy:"name,attr,optional"`
	Namespace string `alloy:"namespace,attr,optional"`
	OpenPorts string `alloy:"open_ports,attr,optional"`
	Path      string `alloy:"exe_path,attr,optional"`
}

type Discovery struct {
	Services Services `alloy:"services,block"`
}
