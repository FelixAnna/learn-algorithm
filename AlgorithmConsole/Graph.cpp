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

LinkedListEdge* Graph::AddLinkedEdge(LinkedListEdge* newEdge)
{
	for (LinkedListEdge* e : Edges)
	{
		if (e->Node->Identifier == newEdge->Node->Identifier)
		{
			e->Node = newEdge->Node;
			return e;
		}
	}

	Edges.push_back(newEdge);
	return newEdge;
}

Vertex * Graph::GetVertex(string name)
{
	for (Vertex* v : Vertexes)
	{
		if (v->Identifier == name || v->Name == name)
		{
			return v;
		}
	}
	return nullptr;
}

LinkedListEdge * Graph::GetLinkedEdge(Vertex * node)
{
	for (LinkedListEdge* e : Edges)
	{
		if (e->Node->Identifier == node->Identifier)
		{
			return e;
		}
	}

	return nullptr;
}

void Graph::FindSP(string start, string end) {
	auto startNode = this->GetVertex(start);
	auto endNode = this->GetVertex(end);
	auto visited = new list<Edge*>();

	visited = FindSP(startNode, endNode, visited);

	for (Edge* e : *visited)
	{
		auto startMinDist = 0;
		auto start = e->Start->Identifier;
		for (Edge* es : *visited)
		{
			if (es->End->Identifier == start) {
				startMinDist = es->MinDistance > 0 ? es->MinDistance : 0;
				break;
			}
		}

		e->MinDistance = startMinDist + e->Distance;
	}

	for (Edge* e : *visited)
	{
		auto minDist = 0;
		auto end = e->End->Identifier;
		for (Edge* es : *visited)
		{
			if (es->End->Identifier == end) {
				minDist = minDist == 0 || es->MinDistance < minDist ? es->MinDistance : minDist;
			}
		}

		for (Edge* es : *visited)
		{
			if (es->End->Identifier == end) {
				es->MinDistance = minDist;
			}
		}
	}
}

list<Edge*>* Graph::FindSP(Vertex* startNode, Vertex* endNode, list<Edge*>* visited) {
	auto edges = this->GetLinkedEdge(startNode);
	auto startEdge = edges;
	auto nextEdge = edges->Next;
	while (nextEdge != nullptr) {
		auto distance = nextEdge->Distance;
		if (nextEdge->Node->Identifier == endNode->Identifier) {
			visited->push_back(new Edge(startEdge->Node, nextEdge->Node, distance));
		}

		if (!CheckVisited(*visited, startEdge->Node, nextEdge->Node)) {
			visited->push_back(new Edge(startEdge->Node, nextEdge->Node, distance));
			FindSP(nextEdge->Node, endNode, visited);
		}

		nextEdge = nextEdge->Next;
	}

	return visited;
}

bool Graph::CheckVisited(list<Edge*> visited, Vertex* startNode, Vertex* endNode) {
	for (Edge* v : visited)
	{
		if (v->Start->Identifier == startNode->Identifier && v->End->Identifier == endNode->Identifier)
		{
			return true;
		}

		if (v->Start->Identifier == endNode->Identifier && v->End->Identifier == startNode->Identifier)
		{
			return true;
		}
	}

	return false;
}
