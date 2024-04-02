# Table: kolide_k2_device

Lists the devices that have been enrolled in Kolide.

## Examples

### Basic info

```sql
select
  name,
  hardware_model,
  serial
from
  kolide_k2_device;
```

### List all devices that are in a failed or failing auth state

```sql
select
  name,
  registered_owner_identifier,
  will_block_at
from
  kolide_k2_device
where
  auth_state != 'Good'
order by
  will_block_at asc;
```