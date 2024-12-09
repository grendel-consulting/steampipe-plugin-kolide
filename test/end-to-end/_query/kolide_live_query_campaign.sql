select
  id,
  name,
  published,
  tables_used,
  created_at
from
  kolide_live_query_campaign
order by
  created_at desc;
