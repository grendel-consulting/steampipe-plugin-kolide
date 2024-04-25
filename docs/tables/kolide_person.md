# Table: kolide_person

Lists the people within your organisation, who may not necessarily have access to the Kolide dashboard (see kolide_admin_user). Devices may be registered to a person

## Examples

### Basic info

```sql
select
  id,
  name,
  email,
  has_registered_device
from
  kolide_person;
```

### Find me which people have not been authorised recently

```sql
select
  name,
  last_authenticated_at
from
  kolide_person
where
  last_authenticated_at <> date_trunc('day', current_date) - interval '4 weeks';
```

### Find me which people have no devices

```sql
select
  name
from
  kolide_person
where
  has_registered_device = false
```
