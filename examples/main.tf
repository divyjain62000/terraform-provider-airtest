terraform {
  required_providers {
    airtest = {
      source  = "airtest/airtest"
      version = "1.0.0"
    }
  }
}

provider "airtest" {
  api_token = "test123"
  api_url   = "https://api.example.com"
}

resource "airtest_user" "test" {
  name  = "John Doe"
  email = "john@example.com"
}