# Table: kolide_person_open_issue

Lists the unresolved and unexempted issues created when a device owned by a specific person fails a check; some checks, when they fail, will produce multiple issues, each with a unique primary_key_value.

You will need to provide a valid `person_id` for all queries to this table.

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
  kolide_person_open_issue
where
  person_id = '12345';
```

### List all device-blocking issues

```sql
select
  title,
  device_id,
  detected_at,
  value
from
  kolide_person_open_issue
where
  person_id = '12345'
  and
  blocks_device_at is not null;
```
