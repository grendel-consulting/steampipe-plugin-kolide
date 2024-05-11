select
  id,
  status,
  requested_at,
  requester_message,
  requester_id,
  device_id
from
  kolide_registration_request
order by
  id asc;
