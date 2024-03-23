---
organization: Grendel Consulting
category: ["asset management"]
brand_color: "#7450F6"
display_name: Kolide K2
name: kolide
description: Kolide gives you accurate, valuable and complete fleet visibility across Mac, Windows and Linux endpoints
og_description: Query Kolide with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/grendel-consulting/kolide-social-graphic.png"
icon_url: "/images/plugins/grendel-consulting/kolide.svg"
---

# Kolide + Steampipe

[Kolide]() gives you accurate, valueable and complete fleet visibility across Mac, Windows and Linux endpoints.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

This is an unofficial plugin, leveraging the public [Kolide K2 API](https://www.kolide.com/docs/developers/api) through the Steampipe engine. Prospective users are encouraged to undergo their usual due diligence in using third-party software.

List all devices monitored by Kolide

```sql
select
  id,
  serial, 
  name
from
  kolide_k2_device
```
```
+------+------------+---------+
| id   | serial     | name    |
+------+------------+---------+
| 1553 | X02YZ1ZYZX | ikebana |
+------+------------+---------+
```

## Documentation

- [Table Definitions and Examples](/plugins/grendel-consulting/kolide/tables)

## Get Started

### Installation

Download and install the latest Kolide plugin:

```zsh
steampipe plugin install grendel-consulting/kolide
```

### Credentials

| Item        | Description                                                                                                                                                                   |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Kolide requires an [API token](https://www.kolide.com/docs/developers/api#authentication), which can be created by an administrator with "Full Access" permissions.           |
| Permissions | By default, your Kolide API token will only have access to read-only endpoints; this should be sufficient for using this plugin.                                              |
| Radius      | Each connection represents a single Kolide account. You will need to manage token rotation and revocation, as Kolide API tokens have no in-built expiry.                      |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/kolide.spc`)<br />2. Credentials specified in environment variables, e.g., `KOLIDE_API_TOKEN`. |

### Configuration

Installing the latest Kolide plugin will create a config file (`~/.steampipe/config/kolide.spc`) with a single connection named `kolide`:

Configure your account details in `~/.steampipe/config/kolide.spc`:

```hcl
connection "kolide" {
  plugin = "grendel-consulting/kolide"

  # Your Kolide K2 API key. Required.
  # Get your API key from Kolide, instructions here: https://www.kolide.com/docs/developers/api#creating-an-api-key.
  # Alternately you set with the `KOLIDE_API_TOKEN` environment variable.
  # api_key = "k2sk_v1_thisIsOurExampleKey"
}
```

Alternatively, and **only if the `api_token` is omitted** in the connections, you can use the  Kolide environment variable to obtain credentials only if api_token is not specified in the connection:

```zsh
export KOLIDE_K2_TOKEN=k2sk_v1_thisIsOurExampleKey
```

## Multiple Connections

You may create multiple Kolide connecions to aggregate queries across multiple Kolide fleets if, for example, you're managing devices on multiple client organisations. You can read up in more detail under [Multi-Account Connections](https://steampipe.io/docs/managing/connections#using-aggregators)

## Get Involved

* Open source: https://github.com/grendel-consulting/steampipe-plugin-kolide