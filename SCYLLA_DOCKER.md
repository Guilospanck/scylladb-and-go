# Scylla DB using Docker
Some best practices using Scylla with Docker.

## Getting performance out of Docker Container
The Docker image defaults to a mode where Scylla's architectural optimizations are not enabled. With command-line settings we can introduce some incremental changes that boost Scylla performance on Docker even more.

### Configuring resource limits
By default Scylla uses all CPUs and memory by default. We can configure some commands to limit te memory and cpu used by Scylla.

The recommended way to run Scylla instances on the same physical hardware is by statically partitioning all resources. For example, using `--cpuset` option to assign cores 0 and 1 to one instance, and 2 and 3 to another.

In scenarios in which static-partitioning is not desired (like mostly-idle cluster without hard latency requirements), the `--overprisioned` is recommended.

> Note that specifying --cpuset will automatically disable --overprovisioned

- `--seeds SEED`: used to seed node with other nodes in a cluster.
- `--smp COUNT`: used to limit the Scylla node to a COUNT number of CPUs 
- `--memory AMOUNT`: used to limit the Scylla node to use up to an AMOUNT of memory. The AMOUNT value supports both M unit for megabytes and G unit for gigabytes 
- `--overprovisioned ENABLE`: enables certain optimizations for Scylla to run efficiently in an overprovisioned environment. 1 enables it and 0 disables it.
- `--cpuset CPUSET`: lets user define in which CPUs the node will run. Examples: single CPU(`--cpuset 1`); a range(`--cpuset 1-3`); a list(`--cpuset 1,2,3`) or a combination of the last two options (`--cpuset 1-3,5`);

Examples running the above commands with `docker`:
```bash
docker run --name some-scylla -d scylladb/scylla --seeds 192.168.0.100,192.168.0.200
docker run --name some-scylla -d scylladb/scylla --smp 2
docker run --name some-scylla -d scylladb/scylla --memory 4G
docker run --name some-scylla -d scylladb/scylla --overprovisioned 1
docker run --name some-scylla -d scylladb/scylla --cpuset 0-2,4
```

In order to run using a `docker-compose file`, you can use the `command` option:
```yaml
image: scylladb/scylla:4.1.0
restart: always
command: --seeds=scylla-node1,scylla-node2 --smp 1 --memory 750M --overprovisioned 1 --api-address 0.0.0.0
```

### More about [`seed`](https://docs.scylladb.com/kb/seed-nodes/)
A Scylla seed node is a regular Scylla node with two extra roles:
- It allows nodes to discover the cluster ring topology on startup (when joining the cluster).
  - What are the IPs of the nodes in the cluster?
  - Which token ranges are available?
  - Which nodes will own which tokens when a new node joins the cluster?
- It assists with `gossip` convergence. Gossiping with other nodes ensures that any update to the cluster is propagated across the cluster. This includes alerting when a node goes dowm, comes back or is removed from the cluster.

<b>Tips for creating Scylla seed nodes:</b>
- The first node in a new cluster needs to be a seed node.
- Ensure that all nodes in the cluster have the same seed nodes listed in each node’s scylla.yaml.
- To maintain resiliency of the cluster, it is recommended to have more than one seed node in the cluster.
- If you have more than one seed in a DC with multiple racks (or availability zones), make sure to put your seeds in different racks.
- You must have at least one node that is not a seed node. You cannot create a cluster where all nodes are seed nodes.
- You should have more than one seed node.

<b>How Many Seed Nodes Do I Need?</b>
Use the following guidelines:
- If your DC has `more` than 6 nodes in it, you need 3 seed nodes per DC.
- If your DC has `less` than 6 nodes in it, you need 2 seed nodes per DC.
-----

## Checking server with Nodetool
```bash
# status
docker exec -it some-scylla nodetool status

# describe
docker exec -it some-scylla nodetool describecluster
```

## Using cqlsh (CQL shell)
```bash
docker exec -it some-scylla cqlsh
```

## Restarting Scylla from within the running node
```bash
docker exec -it some-scylla supervisorctl restart scylla
```

## Checking current Scylla version on the node
```bash
docker exec -it some-scylla scylla --version
```

## Scylla cluster using Docker Compose
Simple file to get a cluster with 3 nodes running:
```yaml
version: '3'

services:
  some-scylla:
    image: scylladb/scylla
    container_name: some-scylla

  some-scylla2:
    image: scylladb/scylla
    container_name: some-scylla2
    command: --seeds=some-scylla

  some-scylla3:
    image: scylladb/scylla
    container_name: some-scylla3
    command: --seeds=some-scylla
```

### scylla.yaml
You can use a `scylla.yaml` file in order to define some characteristics of how scylla will work.
File example (see `./scylla.yaml`).

You can map it into the docker-compose file by using named volumes:
```yaml
volumes:
  - "./scylla.yaml:/etc/scylla/scylla.yaml"
```

### Rack, DC properties
You can use a simple text file to map the volume and tell Scylla about the Rack and DC's (datacenters) name.

`cassandra-rackdc.properties` file:
```txt
#
# cassandra-rackdc.properties
#
# The lines may include white spaces at the beginning and the end.
# The rack and data center names may also include white spaces.
# All trailing and leading white spaces will be trimmed.
#  
dc=DC1
rack=Rack1
# prefer_local=<false | true>
# dc_suffix=<Data Center name suffix, used by EC2SnitchXXX snitches>
#
```

To map it into the docker-compose file, do:
```yaml
# {host}:{container}
volumes:
  - "./cassandra-rackdc.properties:/etc/scylla/cassandra-rackdc.properties"
```

### Data file
You can create a data file to create your keyspaces, table, inserts and so on, in order to make the process easier when getting containers up.

> Remember that under REPLICATION {} you must pass the Datacenter name that you defined either in your .yaml file or using cassandra-rackdc.properties file.

Firstly, create a simple .txt file containing your .CQL commands. Example:
```txt
// initialConfig.txt

CREATE KEYSPACE catalog WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

USE catalog;

CREATE TABLE mutant_data (
first_name text,
last_name text, 
address text, 
picture_location text,
PRIMARY KEY((first_name, last_name))
);

INSERT INTO mutant_data ("first_name","last_name","address","picture_location") VALUES ('Bob','Loblaw','1313 Mockingbird Lane', 'http://www.facebook.com/bobloblaw') ;
INSERT INTO mutant_data ("first_name","last_name","address","picture_location") VALUES ('Bob','Zemuda','1202 Coffman Lane', 'http://www.facebook.com/bzemuda') ;
INSERT INTO mutant_data ("first_name","last_name","address","picture_location") VALUES ('Jim','Jeffries','1211 Hollywood Lane', 'http://www.facebook.com/jeffries') ;

```

Then you can map it into your docker-compose file as a volume:
```yaml
volumes:
  - "./initialConfig.txt:/initialConfig.txt"
```

And, once your container is running, you can pass that into your cqlsh:
```bash
docker exec scylla-node1 cqlsh -f /initialConfig.txt
```