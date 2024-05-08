select
  id,
  name,
  email,
  has_registered_device
from
  kolide_person
order by
  name desc;
