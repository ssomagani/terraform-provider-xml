# Terraform XML File Provider
This is the repository for the Terraform XML File Provider. This provider will allow you to create XML files from your Terraform templates. This is intended to help deployment of software that depend on XML configuration files.

For general information about Terraform, visit the [official website][3] and the
[GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform

# Using the Provider
The current version of this provider requires Terraform v0.12 or higher to run.

You need to run `terraform init` to fetch the provider before deploying

The basic purpose of this provider is to create XML files from Terraform templates. Elements can be created or added to existing XML files by simply specifying elements and values in a key=value format.

## Example
```hcl
resource "xml_file" "example" {
    filename = "/tmp/test.xml"
    elements = {
        "root.second-level.third-level.@attribute" = "value"
        "root.second-level.@attribute" = "value"
        "root.doosra-level.third-level.@name" = "some-value"
    }
}
```

## Installation
#### Build from Source

_Requires Go to be installed on the system._

```
$ go get github.com/ssomagani/terraform-provider-xml
$ cd $GOPATH/src/github.com/ssomagani/terraform-provider-xml
$ go build
```
#### Installing 3rd Party Plugins

See [Terraform documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for installing 3rd party plugins.

## Limitations
* Text elements NOT supported yet
* Wildcards NOT supported yet
* Deletion of elements on update NOT supported yet
