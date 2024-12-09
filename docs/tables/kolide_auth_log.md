# Table: kolide_auth_log

Lists the authentication attempts occurring when a user tries to sign in to an App protected by Kolide Device Trust.

## Examples

### Basic info

```sql
select
  timestamp,
  person_name,
  initial_status,
  result
from
  kolide_auth_log;
```

### List all attempts from the past day

```sql
select
  timestamp,
  person_name,
  initial_status,
  result
from
  kolide_auth_log
where
  timestamp > date_trunc('day', current_date) - interval '1 day';
```

### List all failed attempts performed by a specific user

```sql
select
  timestamp,
  initial_status,
  result
from
  kolide_auth_log
where
  person_name = 'Dennis Nedry'
  and
  result = 'Fail';
```
