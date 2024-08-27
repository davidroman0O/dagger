<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# dagger

```go
// import "github.com/autom8ter/dagger/v3"
import "github.com/davidroman/dagger"
```

> I'm sorry for the fork autom8ter, I don't want the graphviz dependency, it is printing unwanted lines and i don't have the time RIGHT NOW to find why and fix it

Package dagger is a collection of generic, concurrency safe datastructures including a Directed Acyclic Graph and others. Datastructures are implemented using generics in Go 1.18.

Supported Datastructures:

DAG: thread safe directed acyclic graph

Queue: unbounded thread safe fifo queue

Stack: unbounded thread safe lifo stack

BoundedQueue: bounded thread safe fifo queue with a fixed capacity

PriorityQueue: thread safe priority queue

HashMap: thread safe hashmap

Set: thread safe set

ChannelGroup: thread safe group of channels for broadcasting 1 value to N channels

MultiContext: thread safe context for coordinating the cancellation of multiple contexts

Borrower: thread safe object ownership manager

## Index

- [func UniqueID\(prefix string\) string](<#UniqueID>)
- [type BoundedQueue](<#BoundedQueue>)
  - [func NewBoundedQueue\[T any\]\(maxSize int\) \*BoundedQueue\[T\]](<#NewBoundedQueue>)
  - [func \(q \*BoundedQueue\[T\]\) Close\(\)](<#BoundedQueue[T].Close>)
  - [func \(q \*BoundedQueue\[T\]\) Len\(\) int](<#BoundedQueue[T].Len>)
  - [func \(q \*BoundedQueue\[T\]\) Pop\(\) \(T, bool\)](<#BoundedQueue[T].Pop>)
  - [func \(q \*BoundedQueue\[T\]\) PopContext\(ctx context.Context\) \(T, bool\)](<#BoundedQueue[T].PopContext>)
  - [func \(q \*BoundedQueue\[T\]\) Push\(val T\) bool](<#BoundedQueue[T].Push>)
  - [func \(q \*BoundedQueue\[T\]\) PushContext\(ctx context.Context, val T\) bool](<#BoundedQueue[T].PushContext>)
  - [func \(q \*BoundedQueue\[T\]\) Range\(fn func\(element T\) bool\)](<#BoundedQueue[T].Range>)
  - [func \(q \*BoundedQueue\[T\]\) RangeContext\(ctx context.Context, fn func\(element T\) bool\)](<#BoundedQueue[T].RangeContext>)
- [type DAG](<#DAG>)
  - [func NewDAG\[T Node\]\(opts ...DagOpt\) \(\*DAG\[T\], error\)](<#NewDAG>)
  - [func \(g \*DAG\[T\]\) Acyclic\(\) bool](<#DAG[T].Acyclic>)
  - [func \(g \*DAG\[T\]\) BFS\(ctx context.Context, reverse bool, start \*GraphNode\[T\], search GraphSearchFunc\[T\]\) error](<#DAG[T].BFS>)
  - [func \(g \*DAG\[T\]\) DFS\(ctx context.Context, reverse bool, start \*GraphNode\[T\], fn GraphSearchFunc\[T\]\) error](<#DAG[T].DFS>)
  - [func \(g \*DAG\[T\]\) GetEdge\(id string\) \(\*GraphEdge\[T\], bool\)](<#DAG[T].GetEdge>)
  - [func \(g \*DAG\[T\]\) GetEdges\(\) \[\]\*GraphEdge\[T\]](<#DAG[T].GetEdges>)
  - [func \(g \*DAG\[T\]\) GetNode\(id string\) \(\*GraphNode\[T\], bool\)](<#DAG[T].GetNode>)
  - [func \(g \*DAG\[T\]\) GetNodes\(\) \[\]\*GraphNode\[T\]](<#DAG[T].GetNodes>)
  - [func \(g \*DAG\[T\]\) GraphViz\(\) \(image.Image, error\)](<#DAG[T].GraphViz>)
  - [func \(g \*DAG\[T\]\) HasEdge\(id string\) bool](<#DAG[T].HasEdge>)
  - [func \(g \*DAG\[T\]\) HasNode\(id string\) bool](<#DAG[T].HasNode>)
  - [func \(g \*DAG\[T\]\) RangeEdges\(fn func\(e \*GraphEdge\[T\]\) bool\)](<#DAG[T].RangeEdges>)
  - [func \(g \*DAG\[T\]\) RangeNodes\(fn func\(n \*GraphNode\[T\]\) bool\)](<#DAG[T].RangeNodes>)
  - [func \(g \*DAG\[T\]\) SetNode\(node Node\) \*GraphNode\[T\]](<#DAG[T].SetNode>)
  - [func \(g \*DAG\[T\]\) Size\(\) \(int, int\)](<#DAG[T].Size>)
  - [func \(g \*DAG\[T\]\) TopologicalSort\(reverse bool\) \(\[\]\*GraphNode\[T\], error\)](<#DAG[T].TopologicalSort>)
- [type DagOpt](<#DagOpt>)
  - [func WithVizualization\(\) DagOpt](<#WithVizualization>)
- [type GraphEdge](<#GraphEdge>)
  - [func \(n \*GraphEdge\[T\]\) From\(\) \*GraphNode\[T\]](<#GraphEdge[T].From>)
  - [func \(n \*GraphEdge\[T\]\) ID\(\) string](<#GraphEdge[T].ID>)
  - [func \(n \*GraphEdge\[T\]\) Metadata\(\) map\[string\]string](<#GraphEdge[T].Metadata>)
  - [func \(n \*GraphEdge\[T\]\) Relationship\(\) string](<#GraphEdge[T].Relationship>)
  - [func \(n \*GraphEdge\[T\]\) SetMetadata\(metadata map\[string\]string\)](<#GraphEdge[T].SetMetadata>)
  - [func \(n \*GraphEdge\[T\]\) To\(\) \*GraphNode\[T\]](<#GraphEdge[T].To>)
- [type GraphNode](<#GraphNode>)
  - [func \(n \*GraphNode\[T\]\) Ancestors\(fn func\(node \*GraphNode\[T\]\) bool\)](<#GraphNode[T].Ancestors>)
  - [func \(n \*GraphNode\[T\]\) BFS\(ctx context.Context, reverse bool, fn GraphSearchFunc\[T\]\) error](<#GraphNode[T].BFS>)
  - [func \(n \*GraphNode\[T\]\) DFS\(ctx context.Context, reverse bool, fn GraphSearchFunc\[T\]\) error](<#GraphNode[T].DFS>)
  - [func \(n \*GraphNode\[T\]\) Descendants\(fn func\(node \*GraphNode\[T\]\) bool\)](<#GraphNode[T].Descendants>)
  - [func \(n \*GraphNode\[T\]\) EdgesFrom\(relationship string, fn func\(e \*GraphEdge\[T\]\) bool\)](<#GraphNode[T].EdgesFrom>)
  - [func \(n \*GraphNode\[T\]\) EdgesTo\(relationship string, fn func\(e \*GraphEdge\[T\]\) bool\)](<#GraphNode[T].EdgesTo>)
  - [func \(n \*GraphNode\[T\]\) Graph\(\) \*DAG\[T\]](<#GraphNode[T].Graph>)
  - [func \(n \*GraphNode\[T\]\) IsConnectedTo\(node \*GraphNode\[T\]\) bool](<#GraphNode[T].IsConnectedTo>)
  - [func \(n \*GraphNode\[T\]\) Remove\(\) error](<#GraphNode[T].Remove>)
  - [func \(n \*GraphNode\[T\]\) RemoveEdge\(edgeID string\)](<#GraphNode[T].RemoveEdge>)
  - [func \(n \*GraphNode\[T\]\) SetEdge\(relationship string, toNode Node, metadata map\[string\]string\) \(\*GraphEdge\[T\], error\)](<#GraphNode[T].SetEdge>)
- [type GraphSearchFunc](<#GraphSearchFunc>)
- [type HashMap](<#HashMap>)
  - [func NewHashMap\[K comparable, V any\]\(\) \*HashMap\[K, V\]](<#NewHashMap>)
  - [func \(n \*HashMap\[K, V\]\) Clear\(\)](<#HashMap[K, V].Clear>)
  - [func \(n \*HashMap\[K, V\]\) Delete\(key K\)](<#HashMap[K, V].Delete>)
  - [func \(n \*HashMap\[K, V\]\) Exists\(key K\) bool](<#HashMap[K, V].Exists>)
  - [func \(n \*HashMap\[K, V\]\) Filter\(f func\(key K, value V\) bool\) \*HashMap\[K, V\]](<#HashMap[K, V].Filter>)
  - [func \(n \*HashMap\[K, V\]\) Get\(key K\) \(V, bool\)](<#HashMap[K, V].Get>)
  - [func \(n \*HashMap\[K, V\]\) Keys\(\) \[\]K](<#HashMap[K, V].Keys>)
  - [func \(n \*HashMap\[K, V\]\) Len\(\) int](<#HashMap[K, V].Len>)
  - [func \(n \*HashMap\[K, V\]\) Map\(\) map\[K\]V](<#HashMap[K, V].Map>)
  - [func \(n \*HashMap\[K, V\]\) Range\(f func\(key K, value V\) bool\)](<#HashMap[K, V].Range>)
  - [func \(n \*HashMap\[K, V\]\) Set\(key K, value V\)](<#HashMap[K, V].Set>)
  - [func \(n \*HashMap\[K, V\]\) Values\(\) \[\]V](<#HashMap[K, V].Values>)
- [type Node](<#Node>)
- [type PriorityQueue](<#PriorityQueue>)
  - [func NewPriorityQueue\[T any\]\(\) \*PriorityQueue\[T\]](<#NewPriorityQueue>)
  - [func \(q \*PriorityQueue\[T\]\) Len\(\) int](<#PriorityQueue[T].Len>)
  - [func \(q \*PriorityQueue\[T\]\) Peek\(\) \(T, bool\)](<#PriorityQueue[T].Peek>)
  - [func \(q \*PriorityQueue\[T\]\) Pop\(\) \(T, bool\)](<#PriorityQueue[T].Pop>)
  - [func \(q \*PriorityQueue\[T\]\) Push\(item T, weight float64\)](<#PriorityQueue[T].Push>)
  - [func \(q \*PriorityQueue\[T\]\) UpdatePriority\(value T, priority float64\)](<#PriorityQueue[T].UpdatePriority>)
- [type Queue](<#Queue>)
  - [func NewQueue\[T any\]\(\) \*Queue\[T\]](<#NewQueue>)
  - [func \(s \*Queue\[T\]\) Len\(\) int](<#Queue[T].Len>)
  - [func \(s \*Queue\[T\]\) Peek\(\) \(T, bool\)](<#Queue[T].Peek>)
  - [func \(s \*Queue\[T\]\) Pop\(\) \(T, bool\)](<#Queue[T].Pop>)
  - [func \(s \*Queue\[T\]\) Push\(f T\)](<#Queue[T].Push>)
  - [func \(q \*Queue\[T\]\) Range\(fn func\(element T\) bool\)](<#Queue[T].Range>)
  - [func \(q \*Queue\[T\]\) RangeUntil\(fn func\(element T\) bool, done chan struct\{\}\)](<#Queue[T].RangeUntil>)
- [type Set](<#Set>)
  - [func NewSet\[T comparable\]\(\) \*Set\[T\]](<#NewSet>)
  - [func \(s \*Set\[T\]\) Add\(val T\)](<#Set[T].Add>)
  - [func \(s \*Set\[T\]\) Contains\(val T\) bool](<#Set[T].Contains>)
  - [func \(s \*Set\[T\]\) Len\(\) int](<#Set[T].Len>)
  - [func \(s \*Set\[T\]\) Range\(fn func\(element T\) bool\)](<#Set[T].Range>)
  - [func \(s \*Set\[T\]\) Remove\(val T\)](<#Set[T].Remove>)
  - [func \(s \*Set\[T\]\) Sort\(lessFunc func\(i T, j T\) bool\) \[\]T](<#Set[T].Sort>)
  - [func \(s \*Set\[T\]\) Values\(\) \[\]T](<#Set[T].Values>)
- [type Stack](<#Stack>)
  - [func NewStack\[T any\]\(\) \*Stack\[T\]](<#NewStack>)
  - [func \(s \*Stack\[T\]\) Clear\(\)](<#Stack[T].Clear>)
  - [func \(s \*Stack\[T\]\) Len\(\) int](<#Stack[T].Len>)
  - [func \(s \*Stack\[T\]\) Peek\(\) \(T, bool\)](<#Stack[T].Peek>)
  - [func \(s \*Stack\[T\]\) Pop\(\) \(T, bool\)](<#Stack[T].Pop>)
  - [func \(s \*Stack\[T\]\) Push\(f T\)](<#Stack[T].Push>)
  - [func \(s \*Stack\[T\]\) Range\(fn func\(element T\) bool\)](<#Stack[T].Range>)
  - [func \(s \*Stack\[T\]\) RangeUntil\(fn func\(element T\) bool, done chan struct\{\}\)](<#Stack[T].RangeUntil>)
  - [func \(s \*Stack\[T\]\) Sort\(lessFunc func\(i T, j T\) bool\) \[\]T](<#Stack[T].Sort>)
  - [func \(s \*Stack\[T\]\) Values\(\) \[\]T](<#Stack[T].Values>)


<a name="UniqueID"></a>
## func [UniqueID](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L50>)

```go
func UniqueID(prefix string) string
```

UniqueID returns a unique identifier with the given prefix

<a name="BoundedQueue"></a>
## type [BoundedQueue](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L922-L925>)

BoundedQueue is a basic FIFO BoundedQueue based on a buffered channel

```go
type BoundedQueue[T any] struct {
    // contains filtered or unexported fields
}
```

<a name="NewBoundedQueue"></a>
### func [NewBoundedQueue](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L929>)

```go
func NewBoundedQueue[T any](maxSize int) *BoundedQueue[T]
```

NewBoundedQueue returns a new BoundedQueue with the given max size. When the max size is reached, the queue will block until a value is removed. If maxSize is 0, the queue will always block until a value is removed. The BoundedQueue is concurrent\-safe.

<a name="BoundedQueue[T].Close"></a>
### func \(\*BoundedQueue\[T\]\) [Close](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L961>)

```go
func (q *BoundedQueue[T]) Close()
```

Close closes the BoundedQueue channel.

<a name="BoundedQueue[T].Len"></a>
### func \(\*BoundedQueue\[T\]\) [Len](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L991>)

```go
func (q *BoundedQueue[T]) Len() int
```

Len returns the number of elements in the BoundedQueue.

<a name="BoundedQueue[T].Pop"></a>
### func \(\*BoundedQueue\[T\]\) [Pop](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L980>)

```go
func (q *BoundedQueue[T]) Pop() (T, bool)
```

Pop removes and returns an element from the beginning of the BoundedQueue.

<a name="BoundedQueue[T].PopContext"></a>
### func \(\*BoundedQueue\[T\]\) [PopContext](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L986>)

```go
func (q *BoundedQueue[T]) PopContext(ctx context.Context) (T, bool)
```

PopContext removes and returns an element from the beginning of the BoundedQueue. If no element is available, it will block until an element is available or the context is cancelled.

<a name="BoundedQueue[T].Push"></a>
### func \(\*BoundedQueue\[T\]\) [Push](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L969>)

```go
func (q *BoundedQueue[T]) Push(val T) bool
```

Push adds an element to the end of the BoundedQueue and returns a channel that will block until the element is added. If the queue is full, it will block until an element is removed.

<a name="BoundedQueue[T].PushContext"></a>
### func \(\*BoundedQueue\[T\]\) [PushContext](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L975>)

```go
func (q *BoundedQueue[T]) PushContext(ctx context.Context, val T) bool
```

PushContext adds an element to the end of the BoundedQueue and returns a channel that will block until the element is added. If the queue is full, it will block until an element is removed or the context is cancelled.

<a name="BoundedQueue[T].Range"></a>
### func \(\*BoundedQueue\[T\]\) [Range](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L934>)

```go
func (q *BoundedQueue[T]) Range(fn func(element T) bool)
```

Range executes a provided function once for each BoundedQueue element until it returns false.

<a name="BoundedQueue[T].RangeContext"></a>
### func \(\*BoundedQueue\[T\]\) [RangeContext](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L948>)

```go
func (q *BoundedQueue[T]) RangeContext(ctx context.Context, fn func(element T) bool)
```

RangeContext executes a provided function once for each BoundedQueue element until it returns false or a value is sent to the done channel. Use this function when you want to continuously process items from the queue until a done signal is received.

<a name="DAG"></a>
## type [DAG](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L322-L329>)

DAG is a concurrency safe, mutable, in\-memory directed graph

```go
type DAG[T Node] struct {
    // contains filtered or unexported fields
}
```

<a name="NewDAG"></a>
### func [NewDAG](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L346>)

```go
func NewDAG[T Node](opts ...DagOpt) (*DAG[T], error)
```

NewDAG creates a new Directed Acyclic Graph instance

<a name="DAG[T].Acyclic"></a>
### func \(\*DAG\[T\]\) [Acyclic](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L646>)

```go
func (g *DAG[T]) Acyclic() bool
```

Acyclic returns true if the graph contains no cycles.

<a name="DAG[T].BFS"></a>
### func \(\*DAG\[T\]\) [BFS](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L462>)

```go
func (g *DAG[T]) BFS(ctx context.Context, reverse bool, start *GraphNode[T], search GraphSearchFunc[T]) error
```

BFS executes a depth first search on the graph starting from the current node. The reverse parameter determines whether the search is reversed or not. The fn parameter is a function that is called on each node in the graph. If the function returns false, the search is stopped.

<a name="DAG[T].DFS"></a>
### func \(\*DAG\[T\]\) [DFS](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L489>)

```go
func (g *DAG[T]) DFS(ctx context.Context, reverse bool, start *GraphNode[T], fn GraphSearchFunc[T]) error
```

DFS executes a depth first search on the graph starting from the current node. The reverse parameter determines whether the search is reversed or not. The fn parameter is a function that is called on each node in the graph. If the function returns false, the search is stopped.

<a name="DAG[T].GetEdge"></a>
### func \(\*DAG\[T\]\) [GetEdge](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L437>)

```go
func (g *DAG[T]) GetEdge(id string) (*GraphEdge[T], bool)
```

GetEdge returns the edge with the given id

<a name="DAG[T].GetEdges"></a>
### func \(\*DAG\[T\]\) [GetEdges](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L427>)

```go
func (g *DAG[T]) GetEdges() []*GraphEdge[T]
```

GetEdges returns all edges in the graph

<a name="DAG[T].GetNode"></a>
### func \(\*DAG\[T\]\) [GetNode](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L406>)

```go
func (g *DAG[T]) GetNode(id string) (*GraphNode[T], bool)
```

GetNode returns the node with the given id

<a name="DAG[T].GetNodes"></a>
### func \(\*DAG\[T\]\) [GetNodes](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L417>)

```go
func (g *DAG[T]) GetNodes() []*GraphNode[T]
```

GetNodes returns all nodes in the graph

<a name="DAG[T].GraphViz"></a>
### func \(\*DAG\[T\]\) [GraphViz](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L731>)

```go
func (g *DAG[T]) GraphViz() (image.Image, error)
```

GraphViz returns a graphviz image

<a name="DAG[T].HasEdge"></a>
### func \(\*DAG\[T\]\) [HasEdge](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L400>)

```go
func (g *DAG[T]) HasEdge(id string) bool
```

HasEdge returns true if the edge with the given id exists in the graph

<a name="DAG[T].HasNode"></a>
### func \(\*DAG\[T\]\) [HasNode](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L394>)

```go
func (g *DAG[T]) HasNode(id string) bool
```

HasNode returns true if the node with the given id exists in the graph

<a name="DAG[T].RangeEdges"></a>
### func \(\*DAG\[T\]\) [RangeEdges](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L443>)

```go
func (g *DAG[T]) RangeEdges(fn func(e *GraphEdge[T]) bool)
```

RangeEdges iterates over all edges in the graph

<a name="DAG[T].RangeNodes"></a>
### func \(\*DAG\[T\]\) [RangeNodes](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L450>)

```go
func (g *DAG[T]) RangeNodes(fn func(n *GraphNode[T]) bool)
```

RangeNodes iterates over all nodes in the graph

<a name="DAG[T].SetNode"></a>
### func \(\*DAG\[T\]\) [SetNode](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L366>)

```go
func (g *DAG[T]) SetNode(node Node) *GraphNode[T]
```

SetNode sets a node in the graph \- it will use the node's ID as the key and overwrite any existing node with the same ID

<a name="DAG[T].Size"></a>
### func \(\*DAG\[T\]\) [Size](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L412>)

```go
func (g *DAG[T]) Size() (int, int)
```

Size returns the number of nodes and edges in the graph

<a name="DAG[T].TopologicalSort"></a>
### func \(\*DAG\[T\]\) [TopologicalSort](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L682>)

```go
func (g *DAG[T]) TopologicalSort(reverse bool) ([]*GraphNode[T], error)
```



<a name="DagOpt"></a>
## type [DagOpt](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L336>)

DagOpt is an option for configuring a DAG

```go
type DagOpt func(*dagOpts)
```

<a name="WithVizualization"></a>
### func [WithVizualization](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L339>)

```go
func WithVizualization() DagOpt
```

WithVizualization enables graphviz visualization on the DAG

<a name="GraphEdge"></a>
## type [GraphEdge](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L63-L75>)

GraphEdge is a relationship between two nodes

```go
type GraphEdge[T Node] struct {
    // contains filtered or unexported fields
}
```

<a name="GraphEdge[T].From"></a>
### func \(\*GraphEdge\[T\]\) [From](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L88>)

```go
func (n *GraphEdge[T]) From() *GraphNode[T]
```

From returns the from node of the edge

<a name="GraphEdge[T].ID"></a>
### func \(\*GraphEdge\[T\]\) [ID](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L78>)

```go
func (n *GraphEdge[T]) ID() string
```

ID returns the unique identifier of the node

<a name="GraphEdge[T].Metadata"></a>
### func \(\*GraphEdge\[T\]\) [Metadata](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L83>)

```go
func (n *GraphEdge[T]) Metadata() map[string]string
```

Metadata returns the metadata of the node

<a name="GraphEdge[T].Relationship"></a>
### func \(\*GraphEdge\[T\]\) [Relationship](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L98>)

```go
func (n *GraphEdge[T]) Relationship() string
```

Relationship returns the relationship between the two nodes

<a name="GraphEdge[T].SetMetadata"></a>
### func \(\*GraphEdge\[T\]\) [SetMetadata](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L103>)

```go
func (n *GraphEdge[T]) SetMetadata(metadata map[string]string)
```

SetMetadata sets the metadata of the node

<a name="GraphEdge[T].To"></a>
### func \(\*GraphEdge\[T\]\) [To](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L93>)

```go
func (n *GraphEdge[T]) To() *GraphNode[T]
```

To returns the to node of the edge

<a name="GraphNode"></a>
## type [GraphNode](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L120-L126>)

GraphNode is a node in the graph. It can be connected to other nodes via edges.

```go
type GraphNode[T Node] struct {
    Node
    // contains filtered or unexported fields
}
```

<a name="GraphNode[T].Ancestors"></a>
### func \(\*GraphNode\[T\]\) [Ancestors](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L258>)

```go
func (n *GraphNode[T]) Ancestors(fn func(node *GraphNode[T]) bool)
```

Ancestors returns the ancestors of the current node

<a name="GraphNode[T].BFS"></a>
### func \(\*GraphNode\[T\]\) [BFS](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L134>)

```go
func (n *GraphNode[T]) BFS(ctx context.Context, reverse bool, fn GraphSearchFunc[T]) error
```

BFS performs a breadth\-first search on the graph starting from the current node

<a name="GraphNode[T].DFS"></a>
### func \(\*GraphNode\[T\]\) [DFS](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L129>)

```go
func (n *GraphNode[T]) DFS(ctx context.Context, reverse bool, fn GraphSearchFunc[T]) error
```

DFS performs a depth\-first search on the graph starting from the current node

<a name="GraphNode[T].Descendants"></a>
### func \(\*GraphNode\[T\]\) [Descendants](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L280>)

```go
func (n *GraphNode[T]) Descendants(fn func(node *GraphNode[T]) bool)
```

Descendants returns the descendants of the current node

<a name="GraphNode[T].EdgesFrom"></a>
### func \(\*GraphNode\[T\]\) [EdgesFrom](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L140>)

```go
func (n *GraphNode[T]) EdgesFrom(relationship string, fn func(e *GraphEdge[T]) bool)
```

EdgesFrom iterates over the edges from the current node to other nodes with the given relationship. If the relationship is empty, all relationships will be iterated over.

<a name="GraphNode[T].EdgesTo"></a>
### func \(\*GraphNode\[T\]\) [EdgesTo](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L153>)

```go
func (n *GraphNode[T]) EdgesTo(relationship string, fn func(e *GraphEdge[T]) bool)
```

EdgesTo iterates over the edges from other nodes to the current node with the given relationship. If the relationship is empty, all relationships will be iterated over.

<a name="GraphNode[T].Graph"></a>
### func \(\*GraphNode\[T\]\) [Graph](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L253>)

```go
func (n *GraphNode[T]) Graph() *DAG[T]
```

DirectedGraph returns the graph the node belongs to

<a name="GraphNode[T].IsConnectedTo"></a>
### func \(\*GraphNode\[T\]\) [IsConnectedTo](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L302>)

```go
func (n *GraphNode[T]) IsConnectedTo(node *GraphNode[T]) bool
```

IsConnectedTo returns true if the current node is connected to the given node in any direction

<a name="GraphNode[T].Remove"></a>
### func \(\*GraphNode\[T\]\) [Remove](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L236>)

```go
func (n *GraphNode[T]) Remove() error
```

Remove removes the current node from the graph

<a name="GraphNode[T].RemoveEdge"></a>
### func \(\*GraphNode\[T\]\) [RemoveEdge](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L216>)

```go
func (n *GraphNode[T]) RemoveEdge(edgeID string)
```

RemoveEdge removes an edge from the current node by edgeID

<a name="GraphNode[T].SetEdge"></a>
### func \(\*GraphNode\[T\]\) [SetEdge](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L168>)

```go
func (n *GraphNode[T]) SetEdge(relationship string, toNode Node, metadata map[string]string) (*GraphEdge[T], error)
```

SetEdge sets an edge from the current node to the node with the given nodeID. If the nodeID does not exist, an error is returned. If the edgeID is empty, a unique id will be generated. If the metadata is nil, an empty map will be used.

<a name="GraphSearchFunc"></a>
## type [GraphSearchFunc](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L457>)

GraphSearchFunc is a function that is called on each node in the graph during a search

```go
type GraphSearchFunc[T Node] func(ctx context.Context, relationship string, node *GraphNode[T]) bool
```

<a name="HashMap"></a>
## type [HashMap](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L752-L754>)

HashMap is a thread safe map

```go
type HashMap[K comparable, V any] struct {
    // contains filtered or unexported fields
}
```

<a name="NewHashMap"></a>
### func [NewHashMap](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L745>)

```go
func NewHashMap[K comparable, V any]() *HashMap[K, V]
```

NewHashMap creates a new generic hash map

<a name="HashMap[K, V].Clear"></a>
### func \(\*HashMap\[K, V\]\) [Clear](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L792>)

```go
func (n *HashMap[K, V]) Clear()
```

Clear clears the map

<a name="HashMap[K, V].Delete"></a>
### func \(\*HashMap\[K, V\]\) [Delete](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L781>)

```go
func (n *HashMap[K, V]) Delete(key K)
```

Delete deletes the key from the map

<a name="HashMap[K, V].Exists"></a>
### func \(\*HashMap\[K, V\]\) [Exists](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L786>)

```go
func (n *HashMap[K, V]) Exists(key K) bool
```

Exists returns true if the key exists in the map

<a name="HashMap[K, V].Filter"></a>
### func \(\*HashMap\[K, V\]\) [Filter](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L827>)

```go
func (n *HashMap[K, V]) Filter(f func(key K, value V) bool) *HashMap[K, V]
```

Filter returns a new hashmap with the values that return true from the function

<a name="HashMap[K, V].Get"></a>
### func \(\*HashMap\[K, V\]\) [Get](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L767>)

```go
func (n *HashMap[K, V]) Get(key K) (V, bool)
```

Get gets the value from the key

<a name="HashMap[K, V].Keys"></a>
### func \(\*HashMap\[K, V\]\) [Keys](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L800>)

```go
func (n *HashMap[K, V]) Keys() []K
```

Keys returns a copy of the keys in the map as a slice

<a name="HashMap[K, V].Len"></a>
### func \(\*HashMap\[K, V\]\) [Len](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L757>)

```go
func (n *HashMap[K, V]) Len() int
```

Len returns the length of the map

<a name="HashMap[K, V].Map"></a>
### func \(\*HashMap\[K, V\]\) [Map](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L839>)

```go
func (n *HashMap[K, V]) Map() map[K]V
```

Map returns a copy of the hashmap as a map\[string\]T

<a name="HashMap[K, V].Range"></a>
### func \(\*HashMap\[K, V\]\) [Range](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L820>)

```go
func (n *HashMap[K, V]) Range(f func(key K, value V) bool)
```

Range ranges over the map with a function until false is returned

<a name="HashMap[K, V].Set"></a>
### func \(\*HashMap\[K, V\]\) [Set](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L776>)

```go
func (n *HashMap[K, V]) Set(key K, value V)
```

Set sets the key to the value

<a name="HashMap[K, V].Values"></a>
### func \(\*HashMap\[K, V\]\) [Values](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L810>)

```go
func (n *HashMap[K, V]) Values() []V
```

Values returns a copy of the values in the map as a slice

<a name="Node"></a>
## type [Node](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L110-L117>)

Node is a node in the graph. It can be connected to other nodes via edges.

```go
type Node interface {
    // ID returns the unique identifier of the node
    ID() string
    // Metadata returns the metadata of the node
    Metadata() map[string]string
    // SetMetadata sets the metadata of the node
    SetMetadata(metadata map[string]string)
}
```

<a name="PriorityQueue"></a>
## type [PriorityQueue](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L855-L858>)

PriorityQueue is a thread safe priority queue

```go
type PriorityQueue[T any] struct {
    // contains filtered or unexported fields
}
```

<a name="NewPriorityQueue"></a>
### func [NewPriorityQueue](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L861>)

```go
func NewPriorityQueue[T any]() *PriorityQueue[T]
```

NewPriorityQueue creates a new priority queue

<a name="PriorityQueue[T].Len"></a>
### func \(\*PriorityQueue\[T\]\) [Len](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L883>)

```go
func (q *PriorityQueue[T]) Len() int
```

Len returns the length of the queue

<a name="PriorityQueue[T].Peek"></a>
### func \(\*PriorityQueue\[T\]\) [Peek](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L912>)

```go
func (q *PriorityQueue[T]) Peek() (T, bool)
```

Peek returns the next item in the queue without removing it

<a name="PriorityQueue[T].Pop"></a>
### func \(\*PriorityQueue\[T\]\) [Pop](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L900>)

```go
func (q *PriorityQueue[T]) Pop() (T, bool)
```

Pop pops an item off the queue

<a name="PriorityQueue[T].Push"></a>
### func \(\*PriorityQueue\[T\]\) [Push](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L890>)

```go
func (q *PriorityQueue[T]) Push(item T, weight float64)
```

Push pushes an item onto the queue

<a name="PriorityQueue[T].UpdatePriority"></a>
### func \(\*PriorityQueue\[T\]\) [UpdatePriority](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L868>)

```go
func (q *PriorityQueue[T]) UpdatePriority(value T, priority float64)
```



<a name="Queue"></a>
## type [Queue](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L996-L999>)

Queue is a thread safe non\-blocking queue

```go
type Queue[T any] struct {
    // contains filtered or unexported fields
}
```

<a name="NewQueue"></a>
### func [NewQueue](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1002>)

```go
func NewQueue[T any]() *Queue[T]
```

NewQueue returns a new Queue

<a name="Queue[T].Len"></a>
### func \(\*Queue\[T\]\) [Len](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1029>)

```go
func (s *Queue[T]) Len() int
```

Len returns the length of the queue

<a name="Queue[T].Peek"></a>
### func \(\*Queue\[T\]\) [Peek](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1036>)

```go
func (s *Queue[T]) Peek() (T, bool)
```

Peek returns the next item in the queue without removing it

<a name="Queue[T].Pop"></a>
### func \(\*Queue\[T\]\) [Pop](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1015>)

```go
func (s *Queue[T]) Pop() (T, bool)
```

Pop and return top element of Queue. Return false if Queue is empty.

<a name="Queue[T].Push"></a>
### func \(\*Queue\[T\]\) [Push](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1008>)

```go
func (s *Queue[T]) Push(f T)
```

Push a new value onto the Queue

<a name="Queue[T].Range"></a>
### func \(\*Queue\[T\]\) [Range](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1046>)

```go
func (q *Queue[T]) Range(fn func(element T) bool)
```

Range executes a provided function once for each Queue element until it returns false or the Queue is empty.

<a name="Queue[T].RangeUntil"></a>
### func \(\*Queue\[T\]\) [RangeUntil](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1060>)

```go
func (q *Queue[T]) RangeUntil(fn func(element T) bool, done chan struct{})
```

RangeUntil executes a provided function once for each Queue element until it returns false or a value is sent on the done channel. Use this function when you want to continuously process items from the queue until a done signal is received.

<a name="Set"></a>
## type [Set](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1186-L1189>)

Set is a basic thread\-safe Set implementation.

```go
type Set[T comparable] struct {
    // contains filtered or unexported fields
}
```

<a name="NewSet"></a>
### func [NewSet](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1192>)

```go
func NewSet[T comparable]() *Set[T]
```

NewSet returns a new Set with the given initial size.

<a name="Set[T].Add"></a>
### func \(\*Set\[T\]\) [Add](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1198>)

```go
func (s *Set[T]) Add(val T)
```

Add adds an element to the Set.

<a name="Set[T].Contains"></a>
### func \(\*Set\[T\]\) [Contains](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1212>)

```go
func (s *Set[T]) Contains(val T) bool
```

Contains returns true if the Set contains the element.

<a name="Set[T].Len"></a>
### func \(\*Set\[T\]\) [Len](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1231>)

```go
func (s *Set[T]) Len() int
```

Len returns the number of elements in the Set.

<a name="Set[T].Range"></a>
### func \(\*Set\[T\]\) [Range](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1220>)

```go
func (s *Set[T]) Range(fn func(element T) bool)
```

Range executes a provided function once for each Set element until it returns false.

<a name="Set[T].Remove"></a>
### func \(\*Set\[T\]\) [Remove](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1205>)

```go
func (s *Set[T]) Remove(val T)
```

Remove removes an element from the Set.

<a name="Set[T].Sort"></a>
### func \(\*Set\[T\]\) [Sort](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1249>)

```go
func (s *Set[T]) Sort(lessFunc func(i T, j T) bool) []T
```

Sort returns the values of the set as an array sorted by the provided less function

<a name="Set[T].Values"></a>
### func \(\*Set\[T\]\) [Values](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1238>)

```go
func (s *Set[T]) Values() []T
```

Values returns the values of the set as an array

<a name="Stack"></a>
## type [Stack](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1084-L1087>)

Stack is a basic LIFO Stack

```go
type Stack[T any] struct {
    // contains filtered or unexported fields
}
```

<a name="NewStack"></a>
### func [NewStack](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1078>)

```go
func NewStack[T any]() *Stack[T]
```

NewStack returns a new Stack instance

<a name="Stack[T].Clear"></a>
### func \(\*Stack\[T\]\) [Clear](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1130>)

```go
func (s *Stack[T]) Clear()
```

Clear removes all elements from the Stack

<a name="Stack[T].Len"></a>
### func \(\*Stack\[T\]\) [Len](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1166>)

```go
func (s *Stack[T]) Len() int
```

Len returns the number of elements in the Stack.

<a name="Stack[T].Peek"></a>
### func \(\*Stack\[T\]\) [Peek](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1173>)

```go
func (s *Stack[T]) Peek() (T, bool)
```

Peek returns the top element of the Stack without removing it. Return false if Stack is empty.

<a name="Stack[T].Pop"></a>
### func \(\*Stack\[T\]\) [Pop](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1116>)

```go
func (s *Stack[T]) Pop() (T, bool)
```

Pop removes and return top element of Stack. Return false if Stack is empty.

<a name="Stack[T].Push"></a>
### func \(\*Stack\[T\]\) [Push](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1109>)

```go
func (s *Stack[T]) Push(f T)
```

Push a new value onto the Stack \(LIFO\)

<a name="Stack[T].Range"></a>
### func \(\*Stack\[T\]\) [Range](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1137>)

```go
func (s *Stack[T]) Range(fn func(element T) bool)
```

Range executes a provided function once for each Stack element until it returns false.

<a name="Stack[T].RangeUntil"></a>
### func \(\*Stack\[T\]\) [RangeUntil](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1091>)

```go
func (s *Stack[T]) RangeUntil(fn func(element T) bool, done chan struct{})
```

RangeUntil executes a provided function once after calling Pop on the stack until the function returns false or a value is sent on the done channel. Use this function when you want to continuously process items from the stack until a done signal is received.

<a name="Stack[T].Sort"></a>
### func \(\*Stack\[T\]\) [Sort](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1157>)

```go
func (s *Stack[T]) Sort(lessFunc func(i T, j T) bool) []T
```

Sort returns the values of the stack as an array sorted by the provided less function

<a name="Stack[T].Values"></a>
### func \(\*Stack\[T\]\) [Values](<https://github.com/autom8ter/dagger/blob/main/dagger.go#L1150>)

```go
func (s *Stack[T]) Values() []T
```

Values returns the values of the stack as an array

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
