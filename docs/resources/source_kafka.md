---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "materialize_source_kafka Resource - terraform-provider-materialize"
subcategory: ""
description: |-
  A Kafka source describes a Kafka cluster you want Materialize to read data from.
---

# materialize_source_kafka (Resource)

A Kafka source describes a Kafka cluster you want Materialize to read data from.

## Example Usage

```terraform
resource "materialize_source_kafka" "example_source_kafka" {
  name        = "source_kafka"
  schema_name = "schema"
  size        = "3xsmall"
  kafka_connection {
    name          = "kafka_connection"
    database_name = "database"
    schema_name   = "schema"
  }
  format {
    avro {
      schema_registry_connection {
        name          = "csr_connection"
        database_name = "database"
        schema_name   = "schema"
      }
    }
  }
  envelope {
    none = true
  }
}

# CREATE SOURCE kafka_metadata
#   FROM KAFKA CONNECTION "database"."schema"."kafka_connection" (TOPIC 'data')
#   FORMAT AVRO USING CONFLUENT SCHEMA REGISTRY CONNECTION "database"."schema"."csr_connection"
#   ENVELOPE NONE
#   WITH (SIZE = '3xsmall');
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `kafka_connection` (Block List, Min: 1, Max: 1) The Kafka connection to use in the source. (see [below for nested schema](#nestedblock--kafka_connection))
- `name` (String) The identifier for the source.
- `topic` (String) The Kafka topic you want to subscribe to.

### Optional

- `cluster_name` (String) The cluster to maintain this source. If not specified, the size option must be specified.
- `database_name` (String) The identifier for the source database. Defaults to `MZ_DATABASE` environment variable if set or `materialize` if environment variable is not set.
- `envelope` (Block List, Max: 1) How Materialize should interpret records (e.g. append-only, upsert).. (see [below for nested schema](#nestedblock--envelope))
- `format` (Block List, Max: 1) How to decode raw bytes from different formats into data structures Materialize can understand at runtime. (see [below for nested schema](#nestedblock--format))
- `include_headers` (Boolean) Include message headers.
- `include_key` (Boolean) Include a column containing the Kafka message key. If the key is encoded using a format that includes schemas, the column will take its name from the schema. For unnamed formats (e.g. TEXT), the column will be named "key".
- `include_offset` (Boolean) Include an offset column containing the Kafka message offset.
- `include_partition` (Boolean) Include a partition column containing the Kafka message partition
- `include_timestamp` (Boolean) Include a timestamp column containing the Kafka message timestamp.
- `key_format` (Block List, Max: 1) Set the key format explicitly. (see [below for nested schema](#nestedblock--key_format))
- `primary_key` (List of String) Declare a set of columns as a primary key.
- `schema_name` (String) The identifier for the source schema. Defaults to `public`.
- `size` (String) The size of the source.
- `start_offset` (List of Number) Read partitions from the specified offset.
- `start_timestamp` (Number) Use the specified value to set "START OFFSET" based on the Kafka timestamp.
- `value_format` (Block List, Max: 1) Set the value format explicitly. (see [below for nested schema](#nestedblock--value_format))

### Read-Only

- `id` (String) The ID of this resource.
- `qualified_sql_name` (String) The fully qualified name of the source.

<a id="nestedblock--kafka_connection"></a>
### Nested Schema for `kafka_connection`

Required:

- `name` (String) The kafka_connection name.

Optional:

- `database_name` (String) The kafka_connection database name.
- `schema_name` (String) The kafka_connection schema name.


<a id="nestedblock--envelope"></a>
### Nested Schema for `envelope`

Optional:

- `debezium` (Boolean) Use the Debezium envelope, which uses a diff envelope to handle CRUD operations.
- `none` (Boolean) Use an append-only envelope. This means that records will only be appended and cannot be updated or deleted.
- `upsert` (Boolean) Use the upsert envelope, which uses message keys to handle CRUD operations.


<a id="nestedblock--format"></a>
### Nested Schema for `format`

Optional:

- `avro` (Block List, Max: 1) Avro format. (see [below for nested schema](#nestedblock--format--avro))
- `csv` (Block List, Max: 2) CSV format. (see [below for nested schema](#nestedblock--format--csv))
- `json` (Boolean) JSON format.
- `protobuf` (Block List, Max: 1) Protobuf format. (see [below for nested schema](#nestedblock--format--protobuf))
- `text` (Boolean) Text format.

<a id="nestedblock--format--avro"></a>
### Nested Schema for `format.avro`

Required:

- `schema_registry_connection` (Block List, Min: 1, Max: 1) The name of a schema registry connection. (see [below for nested schema](#nestedblock--format--avro--schema_registry_connection))

Optional:

- `key_strategy` (String) How Materialize will define the Avro schema reader key strategy.
- `value_strategy` (String) How Materialize will define the Avro schema reader value strategy.

<a id="nestedblock--format--avro--schema_registry_connection"></a>
### Nested Schema for `format.avro.schema_registry_connection`

Required:

- `name` (String) The schema_registry_connection name.

Optional:

- `database_name` (String) The schema_registry_connection database name.
- `schema_name` (String) The schema_registry_connection schema name.



<a id="nestedblock--format--csv"></a>
### Nested Schema for `format.csv`

Optional:

- `column` (Number) The columns to use for the source.
- `delimited_by` (String) The delimiter to use for the source.
- `header` (List of String) The number of columns and the name of each column using the header row.


<a id="nestedblock--format--protobuf"></a>
### Nested Schema for `format.protobuf`

Required:

- `message` (String) The name of the Protobuf message to use for the source.
- `schema_registry_connection` (Block List, Min: 1, Max: 1) The name of a schema registry connection. (see [below for nested schema](#nestedblock--format--protobuf--schema_registry_connection))

<a id="nestedblock--format--protobuf--schema_registry_connection"></a>
### Nested Schema for `format.protobuf.schema_registry_connection`

Required:

- `name` (String) The schema_registry_connection name.

Optional:

- `database_name` (String) The schema_registry_connection database name.
- `schema_name` (String) The schema_registry_connection schema name.




<a id="nestedblock--key_format"></a>
### Nested Schema for `key_format`

Optional:

- `avro` (Block List, Max: 1) Avro format. (see [below for nested schema](#nestedblock--key_format--avro))
- `csv` (Block List, Max: 2) CSV format. (see [below for nested schema](#nestedblock--key_format--csv))
- `json` (Boolean) JSON format.
- `protobuf` (Block List, Max: 1) Protobuf format. (see [below for nested schema](#nestedblock--key_format--protobuf))
- `text` (Boolean) Text format.

<a id="nestedblock--key_format--avro"></a>
### Nested Schema for `key_format.avro`

Required:

- `schema_registry_connection` (Block List, Min: 1, Max: 1) The name of a schema registry connection. (see [below for nested schema](#nestedblock--key_format--avro--schema_registry_connection))

Optional:

- `key_strategy` (String) How Materialize will define the Avro schema reader key strategy.
- `value_strategy` (String) How Materialize will define the Avro schema reader value strategy.

<a id="nestedblock--key_format--avro--schema_registry_connection"></a>
### Nested Schema for `key_format.avro.schema_registry_connection`

Required:

- `name` (String) The schema_registry_connection name.

Optional:

- `database_name` (String) The schema_registry_connection database name.
- `schema_name` (String) The schema_registry_connection schema name.



<a id="nestedblock--key_format--csv"></a>
### Nested Schema for `key_format.csv`

Optional:

- `column` (Number) The columns to use for the source.
- `delimited_by` (String) The delimiter to use for the source.
- `header` (List of String) The number of columns and the name of each column using the header row.


<a id="nestedblock--key_format--protobuf"></a>
### Nested Schema for `key_format.protobuf`

Required:

- `message` (String) The name of the Protobuf message to use for the source.
- `schema_registry_connection` (Block List, Min: 1, Max: 1) The name of a schema registry connection. (see [below for nested schema](#nestedblock--key_format--protobuf--schema_registry_connection))

<a id="nestedblock--key_format--protobuf--schema_registry_connection"></a>
### Nested Schema for `key_format.protobuf.schema_registry_connection`

Required:

- `name` (String) The schema_registry_connection name.

Optional:

- `database_name` (String) The schema_registry_connection database name.
- `schema_name` (String) The schema_registry_connection schema name.




<a id="nestedblock--value_format"></a>
### Nested Schema for `value_format`

Optional:

- `avro` (Block List, Max: 1) Avro format. (see [below for nested schema](#nestedblock--value_format--avro))
- `csv` (Block List, Max: 2) CSV format. (see [below for nested schema](#nestedblock--value_format--csv))
- `json` (Boolean) JSON format.
- `protobuf` (Block List, Max: 1) Protobuf format. (see [below for nested schema](#nestedblock--value_format--protobuf))
- `text` (Boolean) Text format.

<a id="nestedblock--value_format--avro"></a>
### Nested Schema for `value_format.avro`

Required:

- `schema_registry_connection` (Block List, Min: 1, Max: 1) The name of a schema registry connection. (see [below for nested schema](#nestedblock--value_format--avro--schema_registry_connection))

Optional:

- `key_strategy` (String) How Materialize will define the Avro schema reader key strategy.
- `value_strategy` (String) How Materialize will define the Avro schema reader value strategy.

<a id="nestedblock--value_format--avro--schema_registry_connection"></a>
### Nested Schema for `value_format.avro.schema_registry_connection`

Required:

- `name` (String) The schema_registry_connection name.

Optional:

- `database_name` (String) The schema_registry_connection database name.
- `schema_name` (String) The schema_registry_connection schema name.



<a id="nestedblock--value_format--csv"></a>
### Nested Schema for `value_format.csv`

Optional:

- `column` (Number) The columns to use for the source.
- `delimited_by` (String) The delimiter to use for the source.
- `header` (List of String) The number of columns and the name of each column using the header row.


<a id="nestedblock--value_format--protobuf"></a>
### Nested Schema for `value_format.protobuf`

Required:

- `message` (String) The name of the Protobuf message to use for the source.
- `schema_registry_connection` (Block List, Min: 1, Max: 1) The name of a schema registry connection. (see [below for nested schema](#nestedblock--value_format--protobuf--schema_registry_connection))

<a id="nestedblock--value_format--protobuf--schema_registry_connection"></a>
### Nested Schema for `value_format.protobuf.schema_registry_connection`

Required:

- `name` (String) The schema_registry_connection name.

Optional:

- `database_name` (String) The schema_registry_connection database name.
- `schema_name` (String) The schema_registry_connection schema name.

## Import

Import is supported using the following syntax:

```shell
# Sources can be imported using the source id:
terraform import materialize_source_kafka.example_source_kafka <source_id>
```
