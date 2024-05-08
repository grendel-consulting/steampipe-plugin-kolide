select
  title,
  detected_at,
  blocks_device_at,
  resolved_at,
  exempted
from
  kolide_person_open_issue
where
  person_id = '12345'
order by
  detected_at asc;
