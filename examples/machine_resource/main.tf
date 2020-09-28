terraform {
  required_providers {
    danube = {
      source = "fishman/danube"
      # version = "0.5.2"
    }
  }
}

provider "danube" {
}

resource "danube_machine" "test" {
  hostname = "mytestvm3.lan"
  vcpus    = 1
  ram      = 256

  tags = [
    "juju-controller-uuid:b45e6a5b-bebb-40cb-8266-666b44bc986c",
    "juju-machine-id:default-machine-11",
    "juju-model-uuid:e84e5d9b-1ef5-4d71-80cb-1eec77775b9e",
    "juju-units-deployed:etcd/0 kubernetes-master/0"
  ]

  mdata = {
    test  = "data"
    test2 = "data"
  }
}

output "tags" {
  value = danube_machine.test.tags
}
