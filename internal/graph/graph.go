package graph

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	jsoniter "github.com/json-iterator/go"
)

var graphNodes = []opts.GraphNode{
	{Name: "Node1"},
	{Name: "Node2"},
	{Name: "Node3"},
	{Name: "Node4"},
	{Name: "Node5"},
	{Name: "Node6"},
	{Name: "Node7"},
	{Name: "Node8"},
}

func genLinks() []opts.GraphLink {
	links := make([]opts.GraphLink, 0)
	for i := 0; i < len(graphNodes); i++ {
		for j := 0; j < len(graphNodes); j++ {
			links = append(links, opts.GraphLink{Source: graphNodes[i].Name, Target: graphNodes[j].Name})
		}
	}
	return links
}

func graphBase() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic graph example"}),
	)
	graph.AddSeries("graph", graphNodes, genLinks(),
		charts.WithGraphChartOpts(
			opts.GraphChart{Force: &opts.GraphForce{Repulsion: 8000}},
		),
	)
	return graph
}

func graphCircle() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Circular layout"}),
	)

	graph.AddSeries("graph", graphNodes, genLinks()).
		SetSeriesOptions(
			charts.WithGraphChartOpts(
				opts.GraphChart{
					Force:  &opts.GraphForce{Repulsion: 8000},
					Layout: "circular",
				}),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "right"}),
		)
	return graph
}

func graphNpmDep(graphJson string) *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "npm dependencies demo",
		}))

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	f, err := ioutil.ReadFile(graphJson)
	if err != nil {
		panic(err)
	}
	
	type Data struct {
		Nodes []opts.GraphNode
		Links []opts.GraphLink
	}

	var data Data
	if err := jsoniter.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}

	graph.AddSeries("graph", data.Nodes, data.Links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout:             "force",
					Force:  &opts.GraphForce{Repulsion: 8000},

				Roam:               true,
				FocusNodeAdjacency: true,
			}),
			charts.WithEmphasisOpts(opts.Emphasis{
				Label: &opts.Label{
					Show:     true,
					Color:    "black",
					Position: "left",
				},
			}),
			charts.WithLineStyleOpts(opts.LineStyle{
				Curveness: 0.3,
			}),
		)
	return graph
}

type GraphExamples struct{
	NpmFile string
}

func (ge GraphExamples) Examples() {
	fmt.Printf("Using chart file: %s\n",ge.NpmFile)
	page := components.NewPage()
	page.AddCharts(
		// graphBase(),
		// graphCircle(),
		graphNpmDep(ge.NpmFile),
	)

	f, err := os.Create("web/html/graph.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
