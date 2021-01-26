terraform {
  required_providers {
    scm = {
      source = "cloudogu.com/tf/scm"
    }
  }
}

/*
provider "scm" {
 url = "http://localhost:8080/scm"
 username = "scmadmin"
 password = "scmadmin"
}
*/

provider "scm" {
  url = "https://192.168.56.2/scm"
  username = "admin"
  password = "admin123"
  skip_cert_verify = true
}

resource "scm_repository" "testrepo" {
  namespace = "scmadmin"
  name = "testrepo"
  type = "git"
  description = "this is a test repository"
  import_url = "https://github.com/cloudogu/spring-petclinic"
}