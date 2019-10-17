package main

import (
        "github.com/ssomagani/terraform-provider-xml/xml"
        "github.com/hashicorp/terraform-plugin-sdk/plugin"
        "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return xml.Provider()
                },
        })
}

