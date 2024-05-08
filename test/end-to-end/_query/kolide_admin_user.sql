select
  first_name,
  last_name,
  email,
  access,
  created_at
from
  kolide_admin_user
order by
  email desc;
