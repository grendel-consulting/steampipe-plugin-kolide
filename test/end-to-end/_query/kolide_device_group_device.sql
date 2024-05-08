select
  name,
  hardware_model,
  serial
from
  kolide_device_group_device
where
  device_group_id = '12345'
order by
  name desc;
