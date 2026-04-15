---
page_title: "itrs-uptrends_operatorgroup Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Fetch an operator group by GUID or description when you need to reference it in other resources.
---

# itrs-uptrends_operatorgroup (Data Source)

Lookup an operator group to reuse its GUID or description without hard-coding values.

## Example Usage

```terraform
data "itrs-uptrends_operatorgroup" "by_description" {
  description = "Support Team"
}

data "itrs-uptrends_operatorgroup" "by_id" {
  id = "f5a3f2a3-1234-5678-9abc-def012345678"
}
```

## Schema

### Optional
- `id` (String) Operator group GUID. Provide this or `description`.
- `description` (String) Operator group description. Provide this or `id`. If the description is not unique it is going to give an error.
