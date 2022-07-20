---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_secret_sshkey Resource - terraform-provider-harness"
subcategory: ""
description: |-
  Resource for creating an ssh key type secret.
---

# harness_platform_secret_sshkey (Resource)

Resource for creating an ssh key type secret.

## Example Usage

```terraform
resource "harness_platform_secret_sshkey" "key_tab_file_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  kerberos {
    tgt_key_tab_file_path_spec {
      key_path = "key_path"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "KeyTabFilePath"
  }
}

resource "harness_platform_secret_sshkey" " tgt_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  kerberos {
    tgt_password_spec {
      password = "password"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "Password"
  }
}

resource "harness_platform_secret_sshkey" "sshkey_reference" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    sshkey_reference_credential {
      user_name            = "user_name"
      key                  = "key"
      encrypted_passphrase = "encrypted_passphrase"
    }
    credential_type = "KeyReference"
  }
}

resource "harness_platform_secret_sshkey" " sshkey_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    sshkey_path_credential {
      user_name            = "user_name"
      key_path             = "key_path"
      encrypted_passphrase = "encrypted_passphrase"
    }
    credential_type = "KeyPath"
  }
}

resource "harness_platform_secret_sshkey" "ssh_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    ssh_password_credential {
      user_name = "user_name"
      password  = "password"
    }
    credential_type = "Password"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `kerberos` (Block List, Max: 1) Kerberos authentication scheme (see [below for nested schema](#nestedblock--kerberos))
- `org_id` (String) Unique identifier of the organization.
- `port` (Number) SSH port
- `project_id` (String) Unique identifier of the project.
- `ssh` (Block List, Max: 1) Kerberos authentication scheme (see [below for nested schema](#nestedblock--ssh))
- `tags` (Set of String) Tags to associate with the resource. Tags should be in the form `name:value`.

<a id="nestedblock--kerberos"></a>
### Nested Schema for `kerberos`

Required:

- `principal` (String) Username to use for authentication.
- `realm` (String) Reference to a secret containing the password to use for authentication.

Optional:

- `tgt_generation_method` (String) Method to generate tgt
- `tgt_key_tab_file_path_spec` (Block List, Max: 1) Authenticate to App Dynamics using username and password. (see [below for nested schema](#nestedblock--kerberos--tgt_key_tab_file_path_spec))
- `tgt_password_spec` (Block List, Max: 1) Authenticate to App Dynamics using username and password. (see [below for nested schema](#nestedblock--kerberos--tgt_password_spec))

<a id="nestedblock--kerberos--tgt_key_tab_file_path_spec"></a>
### Nested Schema for `kerberos.tgt_key_tab_file_path_spec`

Optional:

- `key_path` (String) key path


<a id="nestedblock--kerberos--tgt_password_spec"></a>
### Nested Schema for `kerberos.tgt_password_spec`

Optional:

- `password` (String) password



<a id="nestedblock--ssh"></a>
### Nested Schema for `ssh`

Required:

- `credential_type` (String) This specifies SSH credential type as Password, KeyPath or KeyReference

Optional:

- `ssh_password_credential` (Block List, Max: 1) SSH credential of type keyReference (see [below for nested schema](#nestedblock--ssh--ssh_password_credential))
- `sshkey_path_credential` (Block List, Max: 1) SSH credential of type keyPath (see [below for nested schema](#nestedblock--ssh--sshkey_path_credential))
- `sshkey_reference_credential` (Block List, Max: 1) SSH credential of type keyReference (see [below for nested schema](#nestedblock--ssh--sshkey_reference_credential))

<a id="nestedblock--ssh--ssh_password_credential"></a>
### Nested Schema for `ssh.ssh_password_credential`

Required:

- `password` (String) SSH Password.
- `user_name` (String) SSH Username.


<a id="nestedblock--ssh--sshkey_path_credential"></a>
### Nested Schema for `ssh.sshkey_path_credential`

Required:

- `key_path` (String) Path of the key file.
- `user_name` (String) SSH Username.

Optional:

- `encrypted_passphrase` (String) Encrypted Passphrase


<a id="nestedblock--ssh--sshkey_reference_credential"></a>
### Nested Schema for `ssh.sshkey_reference_credential`

Required:

- `user_name` (String) SSH Username.

Optional:

- `encrypted_passphrase` (String) Encrypted Passphrase
- `key` (String) SSH key.

## Import

Import is supported using the following syntax:

```shell
# Import using secret sshkey id
terraform import harness_platform_secret_sshkey.example <secret_sshkey_id>
```