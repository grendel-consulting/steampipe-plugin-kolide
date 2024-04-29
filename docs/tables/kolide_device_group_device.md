# Table: kolide_device_group_device

Lists the member devices within a device group.

You will need to provide a valid `device_group_id` for all queries to this table.

## Examples

### Basic info

```sql
select
  name,
  hardware_model,
  serial
from
  kolide_device_group_device
where
  device_group_id = '12345';
```

### List the devices in this device group whose users have been notified of resolvable issues

```sql
select
  name,
  hardware_model,
  serial
from
  kolide_device_group_device
where
  device_group_id = '12345'
  and
  auth_state = 'Notified';
```
