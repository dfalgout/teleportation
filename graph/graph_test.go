package graph

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// Test names for nodes
// Typos aren't very fun to debug..
const (
	arlington    = "Arlington"
	atlanta      = "Atlanta"
	baltimore    = "Baltimore"
	losAngeles   = "Los Angeles"
	newYork      = "New York"
	oakland      = "Oakland"
	philadelphia = "Philadelphia"
	portland     = "Portland"
	sanFrancisco = "San Francisco"
	seattle      = "Seattle"
	washington   = "Washington"
)

var (
	fullGraph *Graph
)

func TestMain(m *testing.M) {
	fullGraph = initGraph()
	os.Exit(m.Run())
}

func initGraph() *Graph {
	graph := NewGraph()

	graphData := make(map[string][]string)

	// Washington - Baltimore
	// Washington - Atlanta
	// Baltimore - Philadelphia
	// Philadelphia - New York
	// Los Angeles - San Francisco
	// San Francisco - Oakland
	// Los Angeles - Oakland
	// Seattle - New York
	// Seattle - Baltimore

	graphData[washington] = []string{
		baltimore,
		atlanta,
	}

	graphData[baltimore] = []string{
		philadelphia,
	}

	graphData[philadelphia] = []string{
		newYork,
	}

	graphData[losAngeles] = []string{
		sanFrancisco,
		oakland,
	}

	graphData[sanFrancisco] = []string{
		oakland,
	}

	graphData[seattle] = []string{
		newYork,
		baltimore,
	}

	for source, dstList := range graphData {
		for _, dst := range dstList {
			graph.AddTarget(source, dst)
		}
	}

	// printGraph(graph)

	return graph
}

// For debugging puposes
func printGraph(g *Graph) {
	fmt.Println(">>>>>>>> START GRAPH <<<<<<<<")
	for _, v := range g.vertices {
		fmt.Printf("%s\n", v.GetName())
		for _, e := range v.GetEdges() {
			fmt.Printf("\t-> %s\n", e)
		}
	}

	fmt.Println(">>>>>>>> END GRAPH <<<<<<<<")
}

func contains(lhs []string, rhs string) bool {
	for _, v := range lhs {
		if v == rhs {
			return true
		}
	}
	return false
}

func TestAddTarget(t *testing.T) {
	graph := NewGraph()

	source1 := seattle
	targets1 := []string{
		portland,
		sanFrancisco,
		arlington,
	}

	source2 := portland
	targets2 := []string{
		seattle, //Test no dups in bidirectional adding
		sanFrancisco,
	}

	// Set success criteria
	shouldBe := make(graphMap)
	seattleNode := NewNode(seattle)
	portlandNode := NewNode(portland)
	sanFranciscoNode := NewNode(sanFrancisco)
	arlingtonNode := NewNode(arlington)

	seattleTargets := []*node{
		portlandNode,
		sanFranciscoNode,
		arlingtonNode,
	}
	portlandTargets := []*node{
		seattleNode,
		sanFranciscoNode,
	}
	sanFranciscoTargets := []*node{
		seattleNode,
		portlandNode,
	}
	arlingtonTargets := []*node{
		seattleNode,
	}

	seattleNode.AddEdges(seattleTargets)
	portlandNode.AddEdges(portlandTargets)
	sanFranciscoNode.AddEdges(sanFranciscoTargets)
	arlingtonNode.AddEdges(arlingtonTargets)

	shouldBe[seattle] = seattleNode
	shouldBe[portland] = portlandNode
	shouldBe[sanFrancisco] = sanFranciscoNode
	shouldBe[arlington] = arlingtonNode

	graph.AddNode(source1)
	for _, t := range targets1 {
		graph.AddTarget(source1, t)
	}

	graph.AddNode(source2)
	for _, t := range targets2 {
		graph.AddTarget(source2, t)
	}

	// printGraph(graph)

	for k, v := range graph.vertices {
		// List of *nodes should be the same
		if !reflect.DeepEqual(v.GetEdges(), shouldBe[k].GetEdges()) {
			t.Fail()
		}
	}
}

func TestAddNode(t *testing.T) {
	graph := NewGraph()

	testNode1 := portland
	testNode2 := arlington

	graph.AddNode(testNode1)
	graph.AddNode(testNode2)

	node1, err := graph.GetNodeByName(testNode1)
	if err != nil || node1.GetName() != testNode1 {
		t.Log(err)
		t.Fail()
	}

	node2, err := graph.GetNodeByName(testNode2)
	if err != nil || node2.GetName() != testNode2 {
		t.Log(err)
		t.Fail()
	}
}

func TestGetTargets(t *testing.T) {
	testSource1 := washington
	testSource2 := newYork

	// Create nodes
	washingtonNode := NewNode(washington)
	newYorkNode := NewNode(newYork)
	baltimoreNode := NewNode(baltimore)
	atlantaNode := NewNode(atlanta)
	philadelphiaNode := NewNode(philadelphia)
	seattleNode := NewNode(seattle)

	// Create what paths the nodes should have
	washingtonTargets := []*node{
		baltimoreNode,
		atlantaNode,
	}

	newYorkTargets := []*node{
		philadelphiaNode,
		seattleNode,
	}

	shouldBeForWashington := []string{
		baltimore,
		atlanta,
	}

	shouldBeForNewYork := []string{
		philadelphia,
		seattle,
	}

	// Add the nodes
	washingtonNode.AddEdges(washingtonTargets)
	newYorkNode.AddEdges(newYorkTargets)

	targets1, err := fullGraph.GetTargets(testSource1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for _, v := range targets1 {
		if !contains(shouldBeForWashington, v.GetName()) {
			t.Log("Targets don't match.")
			t.Fail()
		}
	}

	targets2, err := fullGraph.GetTargets(testSource2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for _, v := range targets2 {
		if !contains(shouldBeForNewYork, v.GetName()) {
			t.Log("Targets don't match.")
			t.Fail()
		}
	}
}

func TestGetTargetsErrorNoNodes(t *testing.T) {
	graph := NewGraph()

	testItem := "Doesn't Exist"

	node, err := graph.GetTargets(testItem)
	if err.Error() != fmt.Sprintf("node with name %s doesn't exist", testItem) {
		t.Logf("Error should NOT be nil, but equals %s\n", err)
		t.Fail()
	}

	if node != nil {
		t.Logf("Node should be nil but equals %s\n", node)
		t.Fail()
	}
}

func TestTeleportToBFS(t *testing.T) {
	graph := initGraph()

	// printGraph(graph)

	// This should return true
	foundSource := seattle
	foundDest := washington

	found := graph.TeleportToBFS(foundSource, foundDest)

	if !found {
		t.Log("Path should exist")
		t.Fail()
	}

	// This should return false
	notFoundSource := losAngeles
	notFoundDest := baltimore

	notFound := graph.TeleportToBFS(notFoundSource, notFoundDest)

	if notFound {
		t.Log("Path should NOT exist")
		t.Fail()
	}

	// This should return false
	selfSource := seattle
	selfDest := seattle

	selfFound := graph.TeleportToBFS(selfSource, selfDest)

	if selfFound {
		t.Log("Path to self should NOT exist")
		t.Fail()
	}

	noSourceFound := "foo"
	destFound := seattle

	nodeDontExist := graph.TeleportToBFS(noSourceFound, destFound)

	if nodeDontExist {
		t.Log("Source Node doesn't exist, there should NOT be a path here")
		t.Fail()
	}

	source := seattle
	noDestFound := "bar"

	destDontExist := graph.TeleportToBFS(source, noDestFound)

	if destDontExist {
		t.Log("Dest node doesn't exist, there should NOT be a path here")
		t.Fail()
	}
}

func TestGetEdgesAtDepth(t *testing.T) {
	graph := initGraph()

	source := seattle

	shouldBeDepth1 := []string{
		newYork,
		baltimore,
	}

	shouldBeDepth2 := []string{
		newYork,
		philadelphia,
		baltimore,
		washington,
	}

	shouldBeFailedNoSource := []string{}

	// printGraph(graph)

	results1 := graph.GetEdgesAtDepth(source, 1)
	if !reflect.DeepEqual(results1, shouldBeDepth1) {
		t.Logf("results: %s, expected: %s\n", results1, shouldBeDepth1)
		t.Fail()
	}

	results2 := graph.GetEdgesAtDepth(source, 2)
	if !reflect.DeepEqual(results2, shouldBeDepth2) {
		t.Logf("results: %s, expected: %s\n", results2, shouldBeDepth2)
		t.Fail()
	}

	results3 := graph.GetEdgesAtDepth("foo", 2)
	if !reflect.DeepEqual(results3, shouldBeFailedNoSource) {
		t.Logf("results: %s, expected: %s\n", results3, shouldBeFailedNoSource)
		t.Fail()
	}
}

func TestDetectCycles(t *testing.T) {
	graph := initGraph()

	cycleFound := graph.DetectCycles(oakland)
	if !cycleFound {
		t.Logf("%s should have cyclces", oakland)
		t.Fail()
	}

	cycleFound = graph.DetectCycles(washington)
	if !cycleFound {
		t.Logf("%s should have cyclces", washington)
		t.Fail()
	}

	nodeDoesntExistNotFound := graph.DetectCycles("foo")
	if nodeDoesntExistNotFound {
		t.Logf("%s should not have cycles, because the node doesn't exist", "foo")
		t.Fail()
	}
}

func TestCheckUniqueCycle(t *testing.T) {
	graph := initGraph()

	uniqueCycleFound := graph.CheckUniqueCycle(oakland)
	if !uniqueCycleFound {
		t.Logf("%s should have a unique cycle", oakland)
		t.Fail()
	}

	uniqueCycleNotFound := graph.CheckUniqueCycle(washington)
	if uniqueCycleNotFound {
		t.Logf("%s should not have a unique cycle", washington)
		t.Fail()
	}

	nodeDoesntExistNotFound := graph.CheckUniqueCycle("foo")
	if nodeDoesntExistNotFound {
		t.Logf("%s should not have a unique cycle, because the node doesn't exist", "foo")
		t.Fail()
	}
}
