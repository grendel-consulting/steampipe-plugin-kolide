select
  title,
  person_id
from
  kolide_person_open_issue
where
  person_id = '1607'
order by
  detected_at asc;
