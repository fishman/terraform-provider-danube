---
layout: "danube"
page_title: "Danube: danube_machine"
sidebar_current: "docs-danube-data-source"
description: |-
  Sample data source in the Terraform provider danube.
---

# machine_data_source

Sample data source in the Terraform provider danube.

## Example Usage

```hcl
data "danube_machine" "example" {
  name = "foo"
}
```

## Attributes Reference

* `name` - Sample attribute.
