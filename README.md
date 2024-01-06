# tfswitch

<div align="center"><img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/smallerlogo.png" alt="drawing" width="120" height="130"></div>

**tfswitch** lets you switch between different versions of [Terraform] or [OpenTofu]. If you do not have a particular version installed, `tfswitch` will download the version you desire. The installation is minimal and easy. Once installed, simply select the version you require from the dropdown and start using Terraform or OpenTofu.

For requested versions 1.5.x and older, `tfswitch` will always install [Terraform]. If you request version 1.6.0 or newer, `tfswitch` will install [Terraform] by default, however, you can specify that you prefer [OpenTofu] instead.

The "TF" in the name can stand for _Terraform_ or _Tofu_ — your choice!

> [!NOTE]
> This is a **hard fork** of [warrensbox/terraform-switcher](https://github.com/warrensbox/terraform-switcher) (which _appeared_ to be abandoned based on the commit history), and adds support for OpenTofu as an alternative for Terraform 1.6 and newer.

## Installation

`tfswitch` is available for macOS and Linux-based operating systems.

### macOS via Homebrew

Installation for macOS is the easiest with [Homebrew].

### Linux via curl

Installation for Linux operation systems.

```bash
curl -L https://raw.githubusercontent.com/skyzyx/tfswitch/release/install.sh | bash
```

### Download from GitHub

Alternatively, you can [install binaries from our releases](https://github.com/skyzyx/tfswitch/releases).

## Usage

### Select a version (interactive)

You can switch between different versions by typing the command `tfswitch` on your terminal.

1. Select the version you require by using the up and down arrow.
1. Press **Enter** to select the desired version.

> [!NOTE]
> The recently **selected** versions are presented at the top of the dropdown.

### Supply version on command line (non-interactive)

Install and/or switch to Terraform v1.5.0 _specifically_.

```bash
tfswitch 1.5.0
```

### View all versions including alpha, beta, and RCs (non-interactive)

View all available versions.

```bash
tfswitch --list-all
```

### Use environment variables (non-interactive)

Set the `TF_VERSION` environment variable to your desired Terraform/OpenTofu version.

```bash
# Will automatically switch to Terraform 1.5.0
export TF_VERSION=1.5.0
tfswitch
```

### Install very latest version (non-interactive)

```bash
tfswitch --latest
```

### Install latest stable version for a major.minor line

Install and/or switch to v1.5.7 — the latest version of the 1.5.x line.

```bash
tfswitch --latest-stable 1.5
```

### Install latest pre-release version for a major.minor line

Install and/or switch to v1.7.0-rc1 — the latest pre-release of the 1.7.x line (as of 2024-01-05).

```bash
tfswitch --latest-pre 1.7
```

### Show very latest version

View the version number of the latest release without installing/switching-to it.

```bash
tfswitch --show-latest
```

### Show latest stable version for a major.minor line

View the version number of the latest version of the 1.5.x line without installing/switching-to it.

```bash
tfswitch --show-latest-stable 1.5
```

### Show latest pre-release version for a major.minor line

View the version number of the latest pre-release of the 1.7.x line without installing/switching-to it (as of 2024-01-05).

```bash
tfswitch --show-latest-pre 1.7
```

### Use `version.tf`/`versions.tf` file

If a `*.tf` file with the `terraform` → `required_version` constraint is included in the current directory, it should automatically install and/or switch-to that version. For example, the following should automatically switch to the very latest version because there's no upper-bound on the constraint.

```hcl
terraform {
  required_version = ">= 0.12.9"

  required_providers {
    aws        = ">= 2.52.0"
    kubernetes = ">= 1.11.1"
  }
}
```

### Use `.tfswitch.toml` file

This is for non-admin users with limited privileges on their computer. It is similar to using a `.tfswitchrc` file, but you can specify a custom binary path for your installation.

1. Create a custom binary path. For example:

    ```bash
    mkdir -p ~/bin
    ```

1. Add the path to your `$PATH` environment variable via your Bash or Zsh profile.

    ```bash
    export PATH="$PATH:$HOME/bin"
    ```

1. Pass the `--bin` parameter with your custom path to install. For example:

    ```bash
    tfswitch --bin $HOME/bin/terraform 1.5.0
    ```

1. Alternatively, you can create a `.tfswitch.toml` file in your current directory (`./`) OR in your home directory (`~/.tfswitch.toml`). The TOML file in the current directory has a higher precedence than TOML file in the home directory.

    ```toml
    bin = "$HOME/bin/terraform"
    version = "0.11.3"
    ```

1. Run `tfswitch` and it should automatically install the required version in the specified binary path.

> [!NOTE]
> For Linux users that do not have write permission to `/usr/local/bin/`, `tfswitch` will attempt to install to `$HOME/bin/`. Run `export PATH="$PATH:$HOME/bin"` to append this location to your `$PATH`.

### Use `.tfswitchrc`/`.terraform-version` file

Create a `.tfswitchrc` or `.terraform-version` file containing the desired version. For example, if you wanted to install Terraform v1.5.0, you would use:

```bash
echo "1.5.0" > .tfswitchrc
```

If both `.tfswitchrc` and `.terraform-version` are provided, `.tfswitchrc` will be preferred.

### Use `terragrunt.hcl` file

If a `terragrunt.hcl` file with the Terraform constraint is included in the current directory, `tfswitch` will automatically install and/or switch-to that version. For example, the following should automatically switch to the latest 1.5 version:

```toml
terragrunt_version_constraint = ">= 0.26, < 0.27"
terraform_version_constraint  = ">= 1.5, < 1.6"
# ...
```

### Get the version from a subdirectory

```bash
tfswitch --chdir ./terraform
```

### Use custom mirror

To install from a remote mirror other than the default (<https://releases.hashicorp.com/terraform>). Use the `--mirror` parameter.

```bash
tfswitch --mirror https://example.jfrog.io/artifactory/hashicorp
```

### Set a default version for CI pipeline

When running inside of a CI pipeline, you may want a default or fallback version to prevent the pipeline from hanging. This example will install version `1.5.0` when no other version specifications could be detected.

```bash
tfswitch --default 1.5.0
```

## Automation

### Automatically switch with Bash

Add the following to the end of your `~/.bashrc` file. You can adapt this to look for `.tfswitchrc`, `.tfswitch.toml`, `.terraform-version`, or `versions.tf`.

```bash
cdtfswitch() {
  builtin cd "$@";
  cdir=$PWD;
  if [ -e "$cdir/.tfswitchrc" ]; then
    tfswitch
  fi
}
alias cd='cdtfswitch'
```

### Automatically switch with Zsh

Add the following to the end of your `~/.zshrc` file. You can adapt this to look for `.tfswitchrc`, `.tfswitch.toml`, `.terraform-version`, or `versions.tf`.

```zsh
load-tfswitch() {
  local tfswitchrc_path=".tfswitchrc"

  if [ -f "$tfswitchrc_path" ]; then
    tfswitch
  fi
}
add-zsh-hook chpwd load-tfswitch
```

> [!NOTE]
> If you see an error like `command not found: add-zsh-hook`, then you might be on an older version of Zsh (see below), or you simply need to load `add-zsh-hook` by adding `autoload -U add-zsh-hook` to your `~/.zshrc` file.

#### Older versions of Zsh

```zsh
cd() {
  builtin cd "$@";
  cdir=$PWD;
  if [ -e "$cdir/.tfswitchrc" ]; then
    tfswitch
  fi
}
```

### Automatically switch with Fish shell

Add the following to the end of your `~/.config/fish/config.fish` file. You can adapt this to look for `.tfswitchrc`, `.tfswitch.toml`, `.terraform-version`, or `versions.tf`.

```fish
function switch_terraform --on-event fish_postexec
  string match --regex '^cd\s' "$argv" > /dev/null
  set --local is_command_cd $status

  if test $is_command_cd -eq 0
    if count *.tf > /dev/null
      grep -c "required_version" *.tf > /dev/null
      set --local tf_contains_version $status

      if test $tf_contains_version -eq 0
      end
    end
  end
end
```

## Order of precedence

| Order | Method               |
|-------|----------------------|
| 1     | `.tfswitch.toml`     |
| 2     | `.tfswitchrc`        |
| 3     | `.terraform-version` |
| 4     | Environment variable |

  [Homebrew]: https://brew.sh
  [OpenTofu]: https://opentofu.org
  [Terraform]: https://www.terraform.io
