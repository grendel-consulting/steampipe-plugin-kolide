# Table: kolide_person_registered_device

Lists the registered devices belonging to a specified person.

You will need to provide a valid `person_id` for all queries to this table.

## Examples

### Basic info

```sql
select
  name,
  hardware_model,
  serial
from
  kolide_person_registered_device
where
  person_id = '12345';
```
