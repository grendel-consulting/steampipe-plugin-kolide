select
  timestamp,
  person_name,
  initial_status,
  result
from
  kolide_auth_log
where
  id = '1234';
