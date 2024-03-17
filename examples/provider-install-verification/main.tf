terraform {
  required_providers {
    alicloud-fc-url = {
      source = "registry.terraform.io/rhysmdnz/alicloud-fc-url"
    }
    alicloud = {
      source  = "aliyun/alicloud"
      version = "1.218.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "2.4.2"
    }
  }
}

provider "alicloud" {
  region = "cn-shanghai"
}


provider "alicloud-fc-url" {
  region = "cn-shanghai"
}

data "alicloud-fc-url_trigger_url" "example" {
  service_name  = alicloud_fc_service.foo.name
  function_name = alicloud_fc_function.foo.name
  trigger_name  = alicloud_fc_trigger.foo.name
}

output "fc-trigger" {
  value = data.alicloud-fc-url_trigger_url.example.url_internet
}


data "archive_file" "zip" {
  type        = "zip"
  source_file = "hello.py"
  output_path = "hello.zip"
}

data "alicloud_account" "current" {
}

resource "alicloud_fc_service" "foo" {
  name            = "my-fc-service"
  description     = "created by terraform"
  internet_access = false
}

resource "alicloud_fc_function" "foo" {
  service     = alicloud_fc_service.foo.name
  name        = "hello-world"
  description = "created by terraform"
  filename    = "./hello.zip"
  memory_size = "512"
  runtime     = "python3.10"
  handler     = "hello.handler"
}

resource "alicloud_fc_trigger" "foo" {
  service  = alicloud_fc_service.foo.name
  function = alicloud_fc_function.foo.name
  name     = "trigger-for-fc"
  type     = "http"

  config = <<EOF
    {
      "authType": "anonymous",
      "disableURLInternet": false,
      "methods": ["GET"]
    }
  
EOF
}

