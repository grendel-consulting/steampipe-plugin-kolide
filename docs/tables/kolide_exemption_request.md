# Table: kolide_exemption_request

Lists the exemption requests made when a user desires to permanently ignore a specific issue or set of issues on a single device. These can be approved or denied by admins.

## Examples

### Basic info

```sql
select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id,
  issues
from
  kolide_exemption_request;
```

### List all unresolved requests

```sql
select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id,
  issues
from
  kolide_exemption_request
where
  status != 'approved';
```

### List all recent requests

```sql
select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id,
  issues
from
  kolide_exemption_request
where
  requested_at > date_trunc('day', current_date) - interval '4 weeks';
```
