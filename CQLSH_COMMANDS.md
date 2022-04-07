# CQL and CQLsh commands
Once your containers are up and running, get into one of the nodes by running:
```bash
sudo docker exec -it scylla-node1 cqlsh
```
This will enter into the container node and get into the `cqlsh` space.

Once there, you can run some commands:

### Describes all keyspaces that exist
```bash
DESC keyspaces
```
### Describes a keyspace named `catalog`
```bash
DESC catalog
```
### Describes a table `mutant_data` inside the `catalog` keyspace
```bash
DESC catalog.mutant_data
```
### Queries all data inside a table called `mutant_data` from a keyspace named `catalog`
```bash
SELECT * FROM catalog.mutant_data;
```
And of course you can run basically any other [CQL command].

[CQL command]: (https://docs.datastax.com/en/cql-oss/3.x/cql/cql_reference/cqlSelect.html)