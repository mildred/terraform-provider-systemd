package systemd

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"log_level": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "info",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"systemd_unit": resourceSystemdUnit(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

type providerConfiguration struct {
	Logger hclog.Logger
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level: hclog.LevelFromString(data.Get("log_level").(string)),
	})
	configuration := &providerConfiguration{
		Logger: logger,
	}
	return configuration, nil
}
