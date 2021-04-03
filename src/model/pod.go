package model

type PodModel struct {
	PodName  string
	PodImage string
	PodNode  string
}

func MockPodList() []*PodModel {
	return []*PodModel{
		{PodName: "one", PodImage: "nginx:1.18", PodNode: "node1"},
		{PodName: "two", PodImage: "nginx:1.18", PodNode: "node1"},
		{PodName: "three", PodImage: "nginx:1.18", PodNode: "node1"},
	}
}
