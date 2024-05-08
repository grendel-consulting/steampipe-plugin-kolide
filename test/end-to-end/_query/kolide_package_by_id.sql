select
  id,
  built_at,
  version,
  url
from
  kolide_package
where
  id = 'darwin-launchd-pkg';
