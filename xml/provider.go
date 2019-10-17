package xml

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
    provider := &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                        "xml_file": resourceDeployment(),
                },
        }
	return provider
}
