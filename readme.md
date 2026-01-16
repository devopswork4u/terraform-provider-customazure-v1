# Terraform Provider for Microsoft Graph

## Usage

```hcl
terraform {
  required_providers {
    msgraph = {
      source  = "rahulsharma/msgraph"
      version = "0.1.0"
    }
  }
}

provider "msgraph" {
  tenant_id     = var.tenant_id
  client_id     = var.client_id
  client_secret = var.client_secret
}
