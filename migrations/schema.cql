CREATE KEYSPACE IF NOT EXISTS go_api
  WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'replication_factor': 3
};


CREATE TABLE IF NOT EXISTS go_api.users (
  email VARCHAR NOT NULL,
  name VARCHAR,
  hashed_pass text,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY(email)
);
