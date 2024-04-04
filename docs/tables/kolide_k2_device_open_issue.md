# Table: kolide_k2_device_open_issue

Lists the unresolved and unexempted issues created when a specific device fails a check; some checks, when they fail, will produce multiple Issues, each with a unique primary_key_value.

You will need to provide a valid `device_id` for all queries to this table.

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
  kolide_k2_device_open_issue
where
  device_id = '1553';
```

### List all device-blocking issues

```sql
select
  title,
  detected_at,
  value
from
  kolide_k2_device_open_issue
where
  device_id = '1553'
  and
  blocks_device_at is not null;
```

### Diagnose specific issues with battery health

Batteries in most modern laptops have a recharging "cycle count", after which the battery is considered to be fully consumed

```sql
select
  d.name,
  d.hardware_model,
  d.serial,
  o.value->'cycle_count' as battery_cycles,
  o.value->'health' as battery_health,
  o.value->'max_capacity' as battery_max
from
  kolide_k2_device_open_issue o,
  kolide_k2_device d
where
  o.check_id = '15804'
  and
  o.device_id = d.id
  and
  o.device_id = '1553';
 ```
