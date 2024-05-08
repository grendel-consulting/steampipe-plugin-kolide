select
  title,
  detected_at,
  blocks_device_at,
  resolved_at,
  exempted
from
  kolide_device_open_issue
where
  device_id = '1553'
order by
  detected_at asc;
