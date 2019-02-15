#include "pch.h"
#include "Graph.h"


Graph::Graph()
{
	this->Vertexes = {};
	this->Edges = {};
}


Graph::~Graph()
{
}

void Graph::AddVertex(Vertex* vertex)
{
	this->Vertexes.push_back(vertex);
}

Edge* Graph::AddLinkedEdge(Edge* newEdge)
{
	for (Edge* e : Edges)
	{
		if (e->Point->Identifier == newEdge->Point->Identifier)
		{
			e->Point = newEdge->Point;
			return e;
		}
	}

	Edges.push_back(newEdge);
	return newEdge;
}
