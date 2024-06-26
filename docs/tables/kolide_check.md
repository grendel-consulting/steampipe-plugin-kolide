# Table: kolide_check

Lists the checks that Kolide runs on a device on a regular cadence, which are tests that typically produce a passing or failing result.

## Examples

### Basic info

```sql
select
  id,
  name,
  topics,
  compatible_platforms,
  targeted_groups,
  blocking_group_names,
  blocking_enabled
from
  kolide_check;
```

### List all the checks relating to a specific operating system

```sql
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
where
  compatible_platforms @> '["darwin"]';
```

### List all the topics that Kolide breaks checks down into

```sql
select
  distinct topic
from
  kolide_check,
  jsonb_array_elements_text(topics) as topic;
```

### List all the checks relating to a specific topic

```sql
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
where
  topics @> '["remote-services"]';
```
