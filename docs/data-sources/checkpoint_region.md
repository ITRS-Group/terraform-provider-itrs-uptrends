---
page_title: "itrs-uptrends_checkpoint_region Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Select specific checkpoint regions by name to obtain their IDs via the Uptrends API.
---

# itrs-uptrends_checkpoint_region (Data Source)

Use this data source whenever you need checkpoint region IDs for other resources.

## Example Usage

```terraform
data "itrs-uptrends_checkpoint_region" "europe" {
  regions = ["Europe"]
}

data "itrs-uptrends_checkpoint_region" "all_regions" {
}
```

## Schema

### Optional
- `regions` (List of String) Region names to filter; omit to return all.

### Read-Only
- `selected_regions_ids` (List of Number) IDs of the matching regions.
- `regions_data` (List of Object) Region metadata:
  - `id` (Integer) Region ID.
  - `name` (String) Region name.
