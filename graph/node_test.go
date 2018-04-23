package graph

import (
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	nodeName := portland

	n := NewNode(nodeName)

	if n.GetName() != nodeName {
		t.Fail()
		t.Logf("Node name %s != %s", n.GetName(), nodeName)
	}
}

func TestAddEdge(t *testing.T) {
	source := portland
	target1 := arlington
	target2 := seattle

	sourceShouldBe := []*node{
		NewNode(target1),
		NewNode(target2),
	}

	targetsShouldBe := []*node{
		NewNode(source),
	}

	sourceNode := NewNode(source)
	targetNode1 := NewNode(target1)
	targetNode2 := NewNode(target2)

	sourceNode.AddEdge(targetNode1)
	sourceNode.AddEdge(targetNode2)

	sourceNodes := sourceNode.GetEdges()

	// Check Bidirectional paths
	targetNodes1 := targetNode1.GetEdges()
	targetNodes2 := targetNode2.GetEdges()

	fmt.Printf("SourceNodes: %s\n", sourceNodes)
	fmt.Printf("TargetNodes1: %s\n", targetNodes1)
	fmt.Printf("TargetNodes2: %s\n", targetNodes2)

	for i, n := range sourceNodes {
		if n.GetName() != sourceShouldBe[i].GetName() {
			t.Logf("%s != %s\n", n.GetName(), sourceShouldBe[i].GetName())
			t.Fail()
		}
	}

	for i, n := range targetNodes1 {
		if n.GetName() != targetsShouldBe[i].GetName() {
			t.Logf("%s != %s\n", n.GetName(), targetsShouldBe[i].GetName())
			t.Fail()
		}
	}

	for i, n := range targetNodes2 {
		if n.GetName() != targetsShouldBe[i].GetName() {
			t.Logf("%s != %s\n", n.GetName(), targetsShouldBe[i].GetName())
			t.Fail()
		}
	}
}

func TestIsLeaf(t *testing.T) {
	washingtonNode := NewNode(washington)
	baltimoreNode := NewNode(baltimore)
	philadelphiaNode := NewNode(philadelphia)
	atlantaNode := NewNode(atlanta)

	washingtonNode.AddEdge(baltimoreNode)
	washingtonNode.AddEdge(atlantaNode)

	baltimoreNode.AddEdge(philadelphiaNode)

	if !atlantaNode.IsLeaf() {
		t.Logf("%s should be a leaf node", atlantaNode.GetName())
		t.Fail()
	}

	if washingtonNode.IsLeaf() {
		t.Log("%s should not be a leaf node", washingtonNode.GetName())
		t.Fail()
	}
}
