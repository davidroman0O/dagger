package dagger_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/davidroman0O/dagger"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("DAGGER_DEBUG", "true")
}

type User struct {
	id       string
	Name     string
	metadata map[string]string
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Metadata() map[string]string {
	return u.metadata
}

func (u *User) SetMetadata(metadata map[string]string) {
	if u.metadata == nil {
		u.metadata = map[string]string{}
	}
	for k, v := range metadata {
		u.metadata[k] = v
	}
}

var Jane = &User{
	id:   "1",
	Name: "Jane",
}

var John = &User{
	id:   "2",
	Name: "John",
}

var Jake = &User{
	id:   "3",
	Name: "Jake",
}

var users = []*User{
	Jane,
	John,
	Jake,
}

func randUser() *User {
	return &User{
		id:   strconv.Itoa(int(time.Now().UnixNano())),
		Name: "Jane - " + strconv.Itoa(int(time.Now().UnixNano())),
	}
}

func TestGraph(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	t.Run("set node", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		for _, user := range users {
			node := graph.SetNode(user)
			assert.NotNil(t, node)
		}
	})
	t.Run("set edge", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		jane := graph.SetNode(Jane)
		john := graph.SetNode(John)
		edge, err := jane.SetEdge("knows", john, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, edge)
	})
	t.Run("set edge with node", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		jane := graph.SetNode(Jane)
		john := graph.SetNode(John)
		edge, err := jane.SetEdge("knows", john, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, edge)
	})
	t.Run("set edge with node then get", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		jane := graph.SetNode(Jane)
		john := graph.SetNode(John)
		edge, err := jane.SetEdge("knows", john, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, edge)
		edge, ok := graph.GetEdge(edge.ID())
		assert.True(t, ok)
		assert.NotNil(t, edge)
	})
	t.Run("check edges from/to", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		jane := graph.SetNode(Jane)
		john := graph.SetNode(John)
		edge, err := jane.SetEdge("knows", john, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, edge)
		jane.EdgesFrom("", func(e *dagger.GraphEdge[*User]) bool {
			assert.Equal(t, e.ID(), edge.ID())
			return true
		})
		jane.EdgesTo("", func(e *dagger.GraphEdge[*User]) bool {
			assert.NotEqual(t, e.ID(), edge.ID())
			return true
		})
		john.EdgesTo("", func(e *dagger.GraphEdge[*User]) bool {
			assert.Equal(t, e.ID(), edge.ID())
			return true
		})
		john.EdgesFrom("", func(e *dagger.GraphEdge[*User]) bool {
			assert.NotEqual(t, e.ID(), edge.ID())
			return true
		})
	})
	t.Run("remove node", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		jane := graph.SetNode(Jane)
		john := graph.SetNode(John)
		edge, err := jane.SetEdge("knows", john, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, edge)
		assert.NoError(t, jane.Remove())
		_, ok := graph.GetNode(jane.ID())
		assert.False(t, ok)
		_, ok = graph.GetEdge(edge.ID())
		assert.False(t, ok)
	})
	t.Run("remove edge", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		jane := graph.SetNode(Jane)
		john := graph.SetNode(John)
		edge, err := jane.SetEdge("knows", john, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, edge)
		jane.RemoveEdge(edge.ID())
		_, ok := graph.GetEdge(edge.ID())
		assert.False(t, ok)

	})

	t.Run("breadthFirstSearch", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		lastNode := graph.SetNode(Jane)
		for i := 0; i < 100; i++ {
			node := graph.SetNode(randUser())
			edge, err := lastNode.SetEdge("knows", node, map[string]string{
				"weight": fmt.Sprintf("%d", i),
			})
			assert.Nil(t, err)
			assert.NotNil(t, edge)
		}
		nc, ec := graph.Size()
		t.Logf("nodes=%v edges=%v last node: %s", nc, ec, lastNode.ID())
		nodes := make([]*dagger.GraphNode[*User], 0)
		assert.NoError(t, graph.BFS(ctx, false, lastNode, func(ctx context.Context, relationship string, node *dagger.GraphNode[*User]) bool {
			nodes = append(nodes, node)
			return true
		}))

		assert.Equal(t, 100, len(nodes))
	})
	t.Run("depthFirstSearch", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		lastNode := graph.SetNode(Jane)
		for i := 0; i < 100; i++ {
			node := graph.SetNode(randUser())
			edge, err := lastNode.SetEdge("knows", node, map[string]string{
				"weight": fmt.Sprintf("%d", i),
			})
			assert.Nil(t, err)
			assert.NotNil(t, edge)
		}
		nc, ec := graph.Size()
		t.Logf("nodes=%v edges=%v last node: %s", nc, ec, lastNode.ID())
		nodes := make([]*dagger.GraphNode[*User], 0)

		assert.NoError(t, graph.DFS(ctx, false, lastNode, func(ctx context.Context, relationship string, node *dagger.GraphNode[*User]) bool {
			nodes = append(nodes, node)
			return true
		}))
		assert.Equal(t, 100, len(nodes))
	})
	t.Run("acyclic", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		lastNode := graph.SetNode(Jane)
		for i := 0; i < 100; i++ {
			node := graph.SetNode(randUser())
			edge, err := lastNode.SetEdge("knows", node, map[string]string{
				"weight": fmt.Sprintf("%d", i),
			})
			assert.Nil(t, err)
			assert.NotNil(t, edge)
		}
		nc, ec := graph.Size()
		t.Logf("nodes=%v edges=%v last node: %s", nc, ec, lastNode.ID())
		assert.True(t, graph.Acyclic())
	})
	t.Run("topological reverse sort", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		lastNode := graph.SetNode(Jane)
		for i := 0; i < 100; i++ {
			node := graph.SetNode(randUser())
			edge, err := lastNode.SetEdge("knows", node, map[string]string{
				"weight": fmt.Sprintf("%d", i),
			})
			assert.Nil(t, err)
			assert.NotNil(t, edge)
			lastNode = node
		}
		nc, ec := graph.Size()
		assert.True(t, graph.Acyclic())
		t.Logf("nodes=%v edges=%v last node: %s", nc, ec, lastNode.ID())
		nodes, err := graph.TopologicalSort(true)
		assert.NoError(t, err)
		var n *dagger.GraphNode[*User]
		for _, node := range nodes {
			if n == nil {
				n = node
				continue
			}
			split := strings.Split(n.ID(), "-")
			lastNodeID, _ := strconv.Atoi(split[len(split)-1])
			split = strings.Split(node.ID(), "-")
			nodeID, _ := strconv.Atoi(split[len(split)-1])
			assert.LessOrEqual(t, lastNodeID, nodeID)
			n = node
		}
		assert.Equal(t, 101, len(nodes))
	})
	t.Run("topological sort", func(t *testing.T) {
		graph, err := dagger.NewDAG[*User]()
		assert.NoError(t, err)
		lastNode := graph.SetNode(Jane)
		for i := 0; i < 100; i++ {
			node := graph.SetNode(randUser())
			edge, err := lastNode.SetEdge("knows", node, map[string]string{
				"weight": fmt.Sprintf("%d", i),
			})
			assert.Nil(t, err)
			assert.NotNil(t, edge)
			lastNode = node
		}
		nc, ec := graph.Size()
		assert.True(t, graph.Acyclic())
		t.Logf("nodes=%v edges=%v last node: %s", nc, ec, lastNode.ID())
		nodes, err := graph.TopologicalSort(false)
		assert.NoError(t, err)
		var n *dagger.GraphNode[*User]
		for _, node := range nodes {
			if n == nil {
				n = node
				continue
			}
			split := strings.Split(n.ID(), "-")
			lastNodeID, _ := strconv.Atoi(split[len(split)-1])
			split = strings.Split(node.ID(), "-")
			nodeID, _ := strconv.Atoi(split[len(split)-1])
			assert.GreaterOrEqual(t, lastNodeID, nodeID)
			n = node
		}
		assert.Equal(t, 101, len(nodes))
	})
	//t.Run("strongly knows", func(t *testing.T) {
	//	graph, err := dagger.NewDAG[*User]()
	//	{
	//		lastNode := graph.SetNode(Jane)
	//		for i := 0; i < 50; i++ {
	//			node := graph.SetNode(randUser())
	//			edge, err := lastNode.SetEdge("knows", node, map[string]string{
	//				"weight": fmt.Sprintf("%d", i),
	//			})
	//			assert.Nil(t, err)
	//			assert.NotNil(t, edge)
	//			lastNode = node
	//		}
	//	}
	//	{
	//		lastNode := graph.SetNode(dagger.UniqueID("jane"))
	//		for i := 51; i < 100; i++ {
	//			node := graph.SetNode(randUser())
	//			edge, err := lastNode.SetEdge("knows", node, map[string]string{
	//				"weight": fmt.Sprintf("%d", i),
	//			})
	//			assert.Nil(t, err)
	//			assert.NotNil(t, edge)
	//			lastNode = node
	//		}
	//	}
	//	assert.True(t, graph.Acyclic())
	//	components, err := graph.StronglyConnected()
	//	assert.NoError(t, err)
	//	assert.Equal(t, 2, len(components))
	//	bits, _ := json.MarshalIndent(components, "", "  ")
	//	t.Logf("strongly knows components: %v", string(bits))
	//})
}

func TestBoundedQueue(t *testing.T) {
	q := dagger.NewBoundedQueue[string](100)
	for i := 0; i < 100; i++ {
		q.Push(string(fmt.Sprintf("node-%d", i)))
	}
	assert.Equal(t, q.Len(), 100)
	for i := 0; i < 100; i++ {
		v, ok := q.Pop()
		assert.NotNil(t, v)
		assert.True(t, ok)
	}
	assert.Equal(t, q.Len(), 0)
}

func TestStack(t *testing.T) {
	s := dagger.NewStack[string]()
	for i := 0; i < 100; i++ {
		s.Push(string(fmt.Sprintf("node-%d", i)))
	}
	assert.Equal(t, s.Len(), 100)
	for i := 0; i < 100; i++ {
		v, ok := s.Pop()
		assert.NotNil(t, v)
		assert.True(t, ok)
	}
	assert.Equal(t, s.Len(), 0)
}

func TestSet(t *testing.T) {
	s := dagger.NewSet[string]()
	for i := 0; i < 100; i++ {
		s.Add(string(fmt.Sprintf("node-%d", i)))
	}
	assert.Equal(t, s.Len(), 100)
	for i := 0; i < 100; i++ {
		ok := s.Contains(string(fmt.Sprintf("node-%d", i)))
		assert.True(t, ok)
	}
	assert.Equal(t, s.Len(), 100)
	for i := 0; i < 100; i++ {
		s.Remove(string(fmt.Sprintf("node-%d", i)))
	}
	assert.Equal(t, s.Len(), 0)
}

func TestNewPriorityQueue(t *testing.T) {
	q := dagger.NewPriorityQueue[string]()
	for i := 0; i < 100; i++ {
		q.Push(string(fmt.Sprintf("node-%d", i)), float64(i))
	}
	assert.Equal(t, q.Len(), 100)
	first, _ := q.Peek()
	assert.EqualValues(t, first, string("node-0"))
	for i := 0; i < 100; i++ {
		v, ok := q.Pop()
		assert.NotNil(t, v)
		assert.True(t, ok)
	}
	assert.Equal(t, q.Len(), 0)
}

func TestQueue(t *testing.T) {
	q := dagger.NewQueue[string]()
	for i := 0; i < 100; i++ {
		q.Push(string(fmt.Sprintf("node-%d", i)))
	}
	assert.Equal(t, q.Len(), 100)
	for i := 0; i < 100; i++ {
		v, ok := q.Pop()
		assert.NotNil(t, v)
		assert.True(t, ok)
	}
	assert.Equal(t, q.Len(), 0)
}
