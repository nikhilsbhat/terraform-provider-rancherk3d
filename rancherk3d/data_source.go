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
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRANCHERK3D() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRANCHERK3DRead,

		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceRANCHERK3DRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Your code goes here
	return nil
}
