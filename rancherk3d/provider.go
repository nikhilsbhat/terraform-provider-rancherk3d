// ----------------------------------------------------------------------------
//
//     ***     TERRAGEN GENERATED CODE    ***    TERRAGEN GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file was auto generated by Terragen.
//     This autogenerated code has to be enhanced further to make it fully working terraform-provider.
//
//     Get more information on how terragen works.
//     https://github.com/nikhilsbhat/terragen
//
// ----------------------------------------------------------------------------
package rancherk3d

import (
	"fmt"
	"terraform-provider-rancherk3d/k3d"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"kubernetes_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("KUBERNETES_VERSION", "1.20.2-k3s1"),
				Description: "version of kubernetes cluster that has to be created (tag of k3s to be passed)",
			},
			"k3d_api_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("K3D_API_VERSION", "k3d.io/v1alpha2"),
				Description: "api version of k3d to be used while creation of cluster",
			},
			"kind": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     false,
				DefaultFunc:  schema.EnvDefaultFunc("K3D_KIND", "Simple"),
				Description:  "kind of config file that you want to use",
				ValidateFunc: ValidateKindFunc,
			},
			"runtime": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("K3D_RUNTIME", "docker"),
				Description: "runtime in which cluster has to be created, defaults to docker",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			//"rancherk3d_create_cluster":  resourceRANCHERK3D(),
			//"rancherk3d_create_registry": resourceRegistry(),
			"rancherk3d_load_image": resourceImage(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"null_data_source": dataSourceRANCHERK3D(),
		},

		ConfigureContextFunc: k3d.GetK3dConfig,
	}
}

func ValidateKindFunc(v interface{}, k string) (warnings []string, errors []error) {
	if v.(string) != "Simple" {
		return nil, []error{fmt.Errorf("kind '%s' is unsupported only supported value is Simple", k), fmt.Errorf("for more info refer 'https://k3d.io/usage/configfile/'")}
	}
	return
}
