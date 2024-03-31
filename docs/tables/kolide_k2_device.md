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
  will_block_at asc
```

### Count the number of devices by hardware model

For example, to monitor and drive hardware refresh and obselescence

```sql
select 
  hardware_model, 
  count(hardware_model) 
from 
  kolide_k2_device 
group by 
  hardware_model
```

### Count the number of devices by operating system

For example, to monitor and drive patching

```sql
select 
  operating_system, 
  count(operating_system)
from 
  kolide_k2_device 
group by 
  operating_system
```