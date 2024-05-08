select
  title,
  detected_at,
  blocks_device_at,
  resolved_at,
  exempted
from
  kolide_issue
order by
  detected_at asc;
