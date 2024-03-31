# Table: kolide_k2_admin_user

Lists the users with access to the Kolide dashboard. Depending on your organization's restrictions, they are able to view and manage checks, devices, users, and other settings.

## Examples

### Basic info

```sql
select
  first_name,
  last_name,
  email,
  access,
  created_at
from
  kolide_k2_admin_user;
```

### List all Kolide super admins

```sql
select 
  email
from
  kolide_k2_admin_user
where
  access = 'full'
```