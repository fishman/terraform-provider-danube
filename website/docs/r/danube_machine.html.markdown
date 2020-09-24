---
layout: "danube"
page_title: "Danube: danube_machine"
sidebar_current: "docs-danube-resource"
description: |-
  Sample resource in the Terraform provider danube.
---

# danube_machine resource

Sample resource in the Terraform provider danube.

## Example Usage

```hcl
resource "danube_machine" "example" {
  machine = "foo"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Machine name.

