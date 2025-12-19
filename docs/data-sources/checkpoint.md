---
page_title: "itrs-uptrends_checkpoint Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Select one or more checkpoints by name to return their IDs and metadata via the Uptrends API.
---

# itrs-uptrends_checkpoint (Data Source)

Use this data source when you need checkpoint IDs or attributes (IP addresses, HA flags, etc.) without creating them yourself.

## Example Usage

```terraform
data "itrs-uptrends_checkpoint" "example" {
  checkpoints = ["Amsterdam", "Cairo"]
}

data "itrs-uptrends_checkpoint" "all_checkpoints" {
}
```

## Schema

### Optional
- `checkpoints` (List of String) List of checkpoint names to filter. Leave empty to select all.

### Read-Only
- `selected_checkpoints_ids` (List of Number) IDs of the checkpoints matching the requested names (or all when none specified).
- `checkpoints_data` (List of Object) Metadata objects describing each selected checkpoint:
  - `id` (Integer) Checkpoint ID.
  - `type` (String) Checkpoint type.
  - `attributes` (Object) Checkpoint attributes:
    - `checkpoint_name` (String)
    - `code` (String)
    - `ipv4_addresses` (List of String)
    - `ipv6_addresses` (List of String)
    - `is_primary_checkpoint` (Boolean)
    - `supports_ipv6` (Boolean)
    - `has_high_availability` (Boolean)
  - `links` (Object) Contains `self` (String) link.
