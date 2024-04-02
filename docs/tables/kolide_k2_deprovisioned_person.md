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

### List all deprovisioned people who still have a device

```sql
select
  name,
  email,
  last_authenticated_at
from
  kolide_k2_deprovisioned_person
where 
  has_registered_device = 'true';
```

### list all deprovisioned people who have logged in within the past week

```sql
select
  name,
  email,
  last_authenticated_at
from
  kolide_k2_deprovisioned_person
where 
  last_authenticated_at > date_trunc('day', current_date) - interval '1 week';
```

### List all deprovisioned people who were never created

Depending on which Kolide product you use, the deprovisioned people may in fact be regular Kolide K2 users who have been been used for zero-trust device authentication

```sql
select
  name,
  email
from
  kolide_k2_deprovisioned_person
where 
  created_at is null;
```