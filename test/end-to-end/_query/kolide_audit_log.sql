select
  timestamp,
  description,
  actor_name
from
  kolide_audit_log
order by
  timestamp asc;
