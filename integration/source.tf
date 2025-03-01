resource "materialize_source_load_generator" "load_generator" {
  name                = "load_gen"
  schema_name         = materialize_schema.schema.name
  database_name       = materialize_database.database.name
  size                = "1"
  load_generator_type = "COUNTER"

  counter_options {
    tick_interval = "500ms"
  }
}

resource "materialize_source_load_generator" "load_generator_cluster" {
  name                = "load_gen_cluster"
  schema_name         = materialize_schema.schema.name
  database_name       = materialize_database.database.name
  cluster_name        = materialize_cluster.cluster_source.name
  load_generator_type = "COUNTER"

  counter_options {
    tick_interval = "500ms"
  }
}

resource "materialize_source_postgres" "example_source_postgres" {
  name = "source_postgres"
  size = "2"
  postgres_connection {
    name          = materialize_connection_postgres.postgres_connection.name
    schema_name   = materialize_connection_postgres.postgres_connection.schema_name
    database_name = materialize_connection_postgres.postgres_connection.database_name
  }
  publication = "mz_source"
  table {
    name  = "table1"
    alias = "s1_table1"
  }
  table {
    name  = "table2"
    alias = "s2_table1"
  }
}

resource "materialize_source_kafka" "example_source_kafka_format_text" {
  name = "source_kafka_text"
  size = "2"
  kafka_connection {
    name          = materialize_connection_kafka.kafka_connection.name
    schema_name   = materialize_connection_kafka.kafka_connection.schema_name
    database_name = materialize_connection_kafka.kafka_connection.database_name
  }
  topic = "topic1"
  key_format {
    text = true
  }
  value_format {
    text = true
  }
}

resource "materialize_source_kafka" "example_source_kafka_format_avro" {
  name = "source_kafka_avro"
  size = "2"
  kafka_connection {
    name          = materialize_connection_kafka.kafka_connection.name
    schema_name   = materialize_connection_kafka.kafka_connection.schema_name
    database_name = materialize_connection_kafka.kafka_connection.database_name
  }
  format {
    avro {
      schema_registry_connection {
        name          = materialize_connection_confluent_schema_registry.schema_registry.name
        schema_name   = materialize_connection_confluent_schema_registry.schema_registry.schema_name
        database_name = materialize_connection_confluent_schema_registry.schema_registry.database_name
      }
    }
  }
  envelope {
    none = true
  }
  topic      = "topic1"
  depends_on = [materialize_sink_kafka.sink_kafka]
}

output "qualified_load_generator" {
  value = materialize_source_load_generator.load_generator.qualified_sql_name
}

data "materialize_source" "all" {}
