It is possible to the send data the of a Substreams to a subgraph, thus creating a Substreams-powered Subgraph.

## Which Option To Use
There are two ways of making Substreams sink to a subgraph, it is really a matter of where you put your logic, in the subgraph or the Substreams.

- [Substreams Triggers](./triggers.md): Consume from any Substreams module by importing the Protobuf model through a subgraph handler and write all your transformations using AssemblyScript. This method creates the subgraph entities directly in the subgraph.
- [Substreams Graph-Out](./graph-out.md): By writing the subgraph entities into Substreams, you can consume the module's output directly into `graph-node`. You will create a [graph_out module](./graph-out.md) that emits an [EntityChanges](https://github.com/streamingfast/substreams-sink-entity-changes/blob/develop/proto/sf/substreams/sink/entity/v1/entity.proto#L11) Protobuf.
The subgraph will read the `EntityChanges` object and consume the data directly.
 
Having more of your logic in Substreams benefits from a parallelized model and a cursor to [never miss data](../../../references/reliability-guarantees.md), whereas triggers will be linearly consumed in `graph-node`.
 

<figure><img src="../../../.gitbook/assets/consume/service-subgraph.png" width="100%" /></figure>
