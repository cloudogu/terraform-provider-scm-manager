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
  # can be used for authentication if the import url contains a private repository
  import_username = "testuser"
  import_password = "testpw"
}
```

## Update documentation for Terraform

There is the official tool `tfplugindocs` from Terraform, which generates a unified documentation for the providers.
This should be updated before each release if basic things have changed at the provider.

**Note**: Running the tool will remove the docs folder and recreate it. It is therefore useful to remove not generated docs before -> generate docs -> add not generated docs again.

```bash
tfplugindocs
```

### What is the Cloudogu EcoSystem?

The Cloudogu EcoSystem is an open platform, which lets you choose how and where your team creates great software. Each service or tool is delivered as a Dogu, a Docker container. Each Dogu can easily be integrated in your environment just by pulling it from our registry. We have a growing number of ready-to-use Dogus, e.g. SCM-Manager, Jenkins, Nexus, SonarQube, Redmine and many more. Every Dogu can be tailored to your specific needs. Take advantage of a central authentication service, a dynamic navigation, that lets you easily switch between the web UIs and a smart configuration magic, which automatically detects and responds to dependencies between Dogus. The Cloudogu EcoSystem is open source and it runs either on-premises or in the cloud. The Cloudogu EcoSystem is developed by Cloudogu GmbH under [MIT License](https://cloudogu.com/license.html).

### How to get in touch?
Want to talk to the Cloudogu team? Need help or support? There are several ways to get in touch with us:

* [Website](https://cloudogu.com)
* [myCloudogu-Forum](https://forum.cloudogu.com/topic/34?ctx=1)
* [Email hello@cloudogu.com](mailto:hello@cloudogu.com)

---
&copy; 2021 Cloudogu GmbH - MADE WITH :heart:&nbsp;FOR DEV ADDICTS. [Legal notice / Impressum](https://cloudogu.com/imprint.html)