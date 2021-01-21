terraform {
  required_providers {
    scm = {
      source = "cloudogu.com/tf/scm"
    }
  }
}

provider "scm" {
  username = "scmadmin"
  password = "scmadmin"
}


resource "scm_repository" "testrepo" {
  namespace = "scmadmin"
  name = "testrepo"
  type = "git"
}