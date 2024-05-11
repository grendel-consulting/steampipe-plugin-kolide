# Table: kolide_registration_request

Lists the registration requests made when attempting to register a device with Kolide. These can be approved or denied by admins.

## Examples

### Basic info

```sql
select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id
from
  kolide_registration_request;
```

### List all unresolved requests

```sql
select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id
from
  kolide_registration_request
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
  device_id
from
  kolide_registration_request
where
  requested_at > date_trunc('day', current_date) - interval '4 weeks';
```
