# Kolide plugin for Steampipe
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8741/badge)](https://www.bestpractices.dev/projects/8741)[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/grendel-consulting/steampipe-plugin-kolide/badge)](https://securityscorecards.dev/viewer/?uri=github.com/grendel-consulting/steampipe-plugin-kolide)[![Go Report Card](https://goreportcard.com/badge/github.com/grendel-consulting/steampipe-plugin-kolide)](https://goreportcard.com/report/github.com/grendel-consulting/steampipe-plugin-kolide)

Use SQL to query Devices, People, Checks, Issues and more across your [Kolide](https://www.kolide.com/) fleet. Built atop the [Kolide API](https://www.kolide.com/docs/developers/api).

## Quick Start

Install the plugin with [Steampipe](https://steampipe.io)

```zsh
steampipe plugin install grendel-consulting/kolide
```

Create your Kolide API token and config your connection in `~/.steampipe/config/kolide.spc`

Run a query:

```sql
select name,hardware_model,serial from kolide_device;
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```zsh
git clone https://github.com/turbot/steampipe-plugin-aws.git
cd steampipe-plugin-aws
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make install
```

Configure the plugin:

```zsh
cp config/* ~/.steampipe/config
code ~/.steampipe/config/kolide.spc
```

Try it!

```zsh
steampipe query
> .inspect kolide
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND 4.0](https://creativecommons.org/licenses/by-nc-nd/4.0/) (docs) licenses. Please see our [code of conduct](.github/CODE_OF_CONDUCT.md). We look forward to collaborating with you and encourage your wider involvement and contribution in their community!

## Get Involved

**[Join #steampipe on the Turbot Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Kolide Plugin](https://github.com/grendel-consulting/steampipe-plugin-kolide/labels/help%20wanted)
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
