# Table: kolide_k2_issue

Lists the issues created when a device fails a check; some checks, when they fail, will produce multiple Issues, each with a unique primary_key_value.

## Examples

### Basic info

```sql
select
  title,
  detected_at,
  blocks_device_at,
  resolved_at,
  exempted
from
  kolide_k2_issue;
```

### List all unresolved issues

```sql
select
  title,
  detected_at,
  value
from
  kolide_k2_issue
where
  resolved_at is null;
```

### List all device-blocking issues

```sql
select
  title,
  detected_at,
  value
from
  kolide_k2_issue
where
  blocks_device_at is not null;
```

### List all ignored issues

```sql
select
  title,
  detected_at,
  value
from
  kolide_k2_issue
where
  exempted = true;
```

### List devices with open issues

```sql
select
  device_id,
  count(device_id) as count
from
  kolide_k2_issue
where
  resolved_at is null
group by
  device_id;
```
