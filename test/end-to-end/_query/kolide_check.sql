select
  id,
  name,
  topics,
  compatible_platforms,
  targeted_groups,
  blocking_group_names,
  blocking_enabled
from
  kolide_check
order by
  id asc;
