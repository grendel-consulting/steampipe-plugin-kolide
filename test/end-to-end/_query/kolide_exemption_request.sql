select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id,
  issues
from
  kolide_exemption_request
order by
  id asc;
