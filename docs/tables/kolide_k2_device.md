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

### List all Apple devices not on the latest major version of macOS

```sql
select
  name,
  registered_owner_identifier,
  serial
from 
  kolide_k2_device 
where 
  device_type = 'Mac' 
  and 
  operating_system not like '%Sonoma%';
```

### List all Apple devices not on the latest minor version of macOS Sonoma

```sql
select
  name,
  registered_owner_identifier,
  serial
from 
  kolide_k2_device 
where 
  device_type = 'Mac' 
  and 
  operating_system not like '%14.4%'
```

### List all Apple devices considered vintage

Provided as an example; we recommend reviewing the official [list of discontinued, vintage and obsolete products](https://support.apple.com/en-us/102772).

```sql
select
  name,
  registered_owner_identifier,
  serial,
  hardware_model
from
  kolide_k2_device
where
  device_type = 'Mac'
  and
  ( hardware_model like '%2015%' or hardware_model like '%2016%' or hardware_model like '%2017%' )
```
