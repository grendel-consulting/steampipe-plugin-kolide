# Table: kolide_k2_deprovisioned_person

Lists anyone who has been removed from Kolide via SCIM.

## Examples

### Basic info

```sql
select
  name,
  email
from
  kolide_k2_deprovisioned_person;
```

### List all deprovisioned users who still have a device

```sql
select
  name,
  email
from
  kolide_k2_deprovisioned_person
where 
  has_registered_device = 'true'
```
