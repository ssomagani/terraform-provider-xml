resource "xml_file" "example" {
    filename = "/tmp/test.xml"
    elements = {
        "root.second-level.third-level.@attribute" = "value"
        "root.second-level.@attribute" = "value"
        "root.2nd-level.third-level.@name" = "some-value"
    }
}
