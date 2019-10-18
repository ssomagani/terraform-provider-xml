package xml

import (
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/ssomagani/mxj"
    "io/ioutil"
    "strings"
)

const ATTR_PREFIX = "@"
const FIELD_SEP = ":"
const INDENT = "    "
const PRETTY_PREFIX = ""

func init() {
    mxj.SetAttrPrefix(ATTR_PREFIX)
}

func resourceDeployment() *schema.Resource {
        return &schema.Resource{
                Create: resourceDeploymentCreate,
                Read:   resourceDeploymentRead,
                Update: resourceDeploymentUpdate,
                Delete: resourceDeploymentDelete,

                Schema: map[string]*schema.Schema{
                    "filename": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                    "elements": &schema.Schema{
                        Required: true,
                        Type: schema.TypeMap,
                        Elem: &schema.Schema{
                            Type: schema.TypeString,
                        },
                    },
                },
        }
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
    filename := d.Get("filename").(string)
    xmlMap, err := parseXML(filename)
    if err != nil {
        xmlMap = mxj.Map(make(map[string]interface{}, 0))
    }
    
    var elements map[string]interface{} = d.Get("elements").(map[string]interface{})
    for key, val := range elements {
        update(&xmlMap, key, val.(string))
    }
    
    xmlStr, _ := xmlMap.XmlIndent(PRETTY_PREFIX, INDENT)
    ioutil.WriteFile(filename, xmlStr, 0644)
    d.SetId(filename)
    return resourceDeploymentRead(d, m)
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, m interface{}) error {
    filename := d.Get("filename").(string)
    xmlMap, err := parseXML(filename)
    if err != nil {
        return err
    }
    
    var elements map[string]interface{} = d.Get("elements").(map[string]interface{})
    for key, val := range elements {
        update(&xmlMap, key, val.(string))
    }
    
    xmlStr, _ := xmlMap.XmlIndent(PRETTY_PREFIX, INDENT)
    ioutil.WriteFile(filename, xmlStr, 0644)
    d.SetId(filename)
    return resourceDeploymentRead(d, m)
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}

func update(xmlMap *mxj.Map, fullPath string, value string) error {
    tokens := strings.Split(fullPath, ATTR_PREFIX)
    path_with_dot_end := tokens[0]
    path_rune := []rune(path_with_dot_end)
    path := string(path_rune[0:len(path_rune)-1])
    
    attr := tokens[1]
    val := ATTR_PREFIX+attr+FIELD_SEP+value
    
    _, err := xmlMap.UpdateValuesForPath(val, path)
    if err != nil {
        return err
    }
    return nil
}

func parseXML(filename string) (mxj.Map, error) {
    xmlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    xmlMap, _ := mxj.NewMapXml([]byte(xmlFile))
    return xmlMap, nil
}
