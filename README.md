# Kolide plugin for Steampipe
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8741/badge)](https://www.bestpractices.dev/projects/8741)

Use SQL to query Devices, People, Checks, Issues and more across your [Kolide](https://www.kolide.com/) fleet. Built atop the [Kolide K2 API](https://www.kolide.com/docs/developers/api).

## Quick Start

Install the plugin with [Steampipe](https://steampipe.io)

```zsh
steampipe plugin install grendel-consulting/kolide
```

Create your Kolide API token and config your connection in `~/.steampipe/config/kolide.spc`

```zsh
steampipe query "select name,hardware_model,serial from kolide_k2_device;"
```

## Development

Build the plugin and install it in your `.steampipe` directory:

```zsh
make install
```

Copy the default config file:

```zsh
make reconfigure
```
