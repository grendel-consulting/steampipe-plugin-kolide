select
  name,
  hardware_model,
  serial
from
  kolide_person_registered_device
where
  person_id = '12345'
order by
  name desc;
