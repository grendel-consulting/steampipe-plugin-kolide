select
  id,
  query,
  status,
  created_at,
  updated_at
from
  kolide_live_query_campaign
order by
  created_at desc;
