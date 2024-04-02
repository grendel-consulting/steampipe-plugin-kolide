# Table: kolide_k2_check

Lists the checks that Kolide runs on a device on a regular cadence, which are tests that typically produces a passing or failing result.

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
  kolide_k2_check;
```

### List all of the checks relating to a specific operating system

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
  kolide_k2_check 
where 
  compatible_platforms @> '["darwin"]';
```

### List all of the topics that Kolide breaks checks down into

```sql
select 
  distinct topic
from 
  kolide_k2_check,
  jsonb_array_elements_text(topics) as topic;
```

### List all of the checks relating to a specific topic

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
  kolide_k2_check 
where 
  topics @> '["remote-services"]';
```