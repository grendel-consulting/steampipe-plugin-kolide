# Table: kolide_audit_log

Lists the tracked events occurring in the Kolide web console.

## Examples

### Basic info

```sql
select
  timestamp,
  description,
  actor_name
from
  kolide_audit_log;
```

### List all events from the past day

```sql
select
  timestamp,
  description,
  actor_name
from
  kolide_audit_log
where
  timestamp > date_trunc('day', current_date) - interval '1 day';
```

### List all events performed by a specific user

```sql
select
  timestamp,
  description,
from
  kolide_audit_log
where
  actor_name = 'Dennis Nedry';
```
