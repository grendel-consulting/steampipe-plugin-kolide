# Table: kolide_live_query_campaign

Lists the collection of ad hoc SQL queries set up in your Kolide fleet

## Examples

### Basic info

```sql
select
  id,
  name,
  published,
  tables_used,
  created_at
from
  kolide_live_query_campaign;
```

### Lists all recently created and running campaigns

```sql
select
  id,
  name,
  published,
  tables_used,
  created_at
from
  kolide_live_query_campaign
where
  created_at > date_trunc('day', current_date) - interval '7 days'
  and
  published = true;
```
