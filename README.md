# Terraform Provider for SCM-Manager

## Build provider

Run the following command to build the provider

```shell
make compile
```

For a new release use:

```shell
make clean package signature
```

## Development/ Testing

Either use a prebuild version of the provider or build it locally and put it in the `~/.terraform.d/plugins` directory.

For local development the following command can be used:

```shell
make install-local
```

### Test sample configuration

To run the example configuration you need:

- the [Terraform 0.14+ CLI](https://learn.hashicorp.com/tutorials/terraform/install-cli) installed locally
- [Docker](https://www.docker.com/products/docker-desktop) and [Docker Compose](https://docs.docker.com/compose/install/)

Start a local instance of the scm-manager:

```shell
docker-compose up -d
```

After installing the provider, navigate to the `examples` directory.

```shell
cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```

## Provider Configuration

To use the scm provider terraform needs know its location which can be defined in the following way:

```tf
terraform {
  required_providers {
    scm = {
      source = "cloudogu.com/tf/scm"
    }
  }
}
```

A complete version of the provider configuration:

```tf
provider "scm" {
  url = "http://localhost:8080/scm"
  skip_cert_verify = "false"
  username = "scmadmin"
  password = "scmadmin"
}
```

## Resources
The scm provider enables terraform to create the following resources:
### scm_repository
```tf
resource "scm_repository" "testrepo" {
  namespace = "scmadmin"
  name = "testrepo"
  type = "git"
  description = "this is a test repository"
  contact = "scmadmin@test.test"
  # can be used to populate the new repository with the content of another repository
  import_url = "https://github.com/cloudogu/spring-petclinic"  
}
```