# Table: kolide_k2_audit_logs

Lists the tracked events occurring in the Kolide web console

## Examples

### Basic info

```sql
select
  timestamp,
  description,
  actor_name
from
  kolide_k2_audit_logs;
```
