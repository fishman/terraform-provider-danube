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

data "danube_machine" "test" {
  hostname = "testvm.lan"
}

output "alias" {
  value = data.danube_machine.test.alias
}
