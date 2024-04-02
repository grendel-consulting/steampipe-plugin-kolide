# Table: kolide_k2_package

Lists the published installation packages for the Kolide Launcher agent for each major operating system.

## Examples

### Basic info

```sql
select
  id,
  built_at,
  version,
  url
from
  kolide_k2_package;
```

### Find me which installers have had recent releases

```sql
select 
  id, 
  version 
from 
  kolide_k2_package 
where 
  built_at > date_trunc('day', current_date) - interval '4 weeks';
```

### Find me the installer url for macOS

```sql
select 
  url 
from
  kolide_k2_package 
where 
  id like 'darwin%'
```