table "categories" {
  schema = schema.ecom
  column "id" {
    null           = false
    type           = int
    auto_increment = true
  }
  column "name" {
    null    = false
    type    = varchar(128)
    charset = "utf8mb3"
    collate = "utf8mb3_general_ci"
  }
  column "created_at" {
    null    = true
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = true
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  index "name" {
    unique  = true
    columns = [column.name]
  }
}
table "item_images" {
  schema = schema.ecom
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "display_name" {
    null = true
    type = varchar(128)
  }
  column "image_url" {
    null = false
    type = varchar(255)
  }
  column "item_id" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "item_images_ibfk_1" {
    columns     = [column.item_id]
    ref_columns = [table.items.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "item_id" {
    columns = [column.item_id]
  }
}
table "items" {
  schema = schema.ecom
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "name" {
    null    = true
    type    = varchar(128)
    charset = "utf8mb3"
    collate = "utf8mb3_general_ci"
  }
  column "category_id" {
    null = true
    type = int
  }
  column "short_description" {
    null    = true
    type    = varchar(255)
    charset = "utf8mb3"
    collate = "utf8mb3_general_ci"
  }
  column "original_price" {
    null = false
    type = float
  }
  column "created_at" {
    null    = true
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = true
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "items_ibfk_1" {
    columns     = [column.category_id]
    ref_columns = [table.categories.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "category_id" {
    columns = [column.category_id]
  }
}
schema "ecom" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
