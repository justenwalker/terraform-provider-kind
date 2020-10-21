package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	stdlog "log"
	"os"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/log"
	"strings"
)

const (
	ProviderTypeDocker = "docker"
	ProviderTypePodman = "podman"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kind_cluster": resourceCluster(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"kind_cluster": dataCluster(),
		},
		Schema: map[string]*schema.Schema{
			"provider": {
				Type:             schema.TypeString,
				Description:      "The provider used to run the containers. Can be either `docker` or `podman`",
				Optional:         true,
				Default:          ProviderTypeDocker,
				ValidateDiagFunc: validateProvider,
			},
			"kubeconfig": {
				Type:        schema.TypeString,
				Description: "Sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config",
				Optional:    true,
			},
			"verbosity": {
				Type:        schema.TypeInt,
				Description: "Sets the logging verbosity. larger number means more logs.",
				Default:     -1,
				Optional:    true,
			},
		},
		ConfigureContextFunc: configureProviderMeta,
	}
}

func configureProviderMeta(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var meta Meta
	meta.ProviderName = data.Get("provider").(string)
	if meta.ProviderName == "" {
		meta.ProviderName = ProviderTypeDocker
	}
	if v, ok := data.GetOk("kubeconfig"); ok {
		meta.KubeConfigPath = v.(string)
	}
	var verbosity log.Level
	if v, ok := data.GetOk("verbosity"); ok {
		verbosity = log.Level(v.(int))
	}
	var opt cluster.ProviderOption
	switch meta.ProviderName {
	case "", ProviderTypeDocker:
		opt = cluster.ProviderWithDocker()
	case ProviderTypePodman:
		opt = cluster.ProviderWithPodman()
	default:
		return nil, diag.FromErr(fmt.Errorf("unknown provider name %q", meta.ProviderName))
	}
	meta.Provider = cluster.NewProvider(opt, cluster.ProviderWithLogger(logger{
		Logger:   stdlog.New(os.Stderr, "<KIND> ", stdlog.LstdFlags),
		MaxLevel: verbosity,
	}))
	return &meta, nil
}

var validProviders = []string{
	ProviderTypeDocker,
	ProviderTypePodman,
}

func validateProvider(v interface{}, path cty.Path) (diags diag.Diagnostics) {
	provider := v.(string)
	for _, p := range validProviders {
		if p == provider {
			return
		}
	}
	diags = append(diags, diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       "invalid provider name",
		Detail:        fmt.Sprintf("provider name was not valid. Must be one of: %s", strings.Join(validProviders, ",")),
		AttributePath: path,
	})
	return
}
