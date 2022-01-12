#include "LinkedListEdge.h"
#include "Edge.h"
#include <list>
#include <iostream>

using namespace std;

#pragma once
class Graph
{
private:
	bool CheckVisited(list<Edge*> visited, Vertex* startNode, Vertex* endNode);
	list<Edge*>* FindSP(Vertex* startNode, Vertex* endNode, list<Edge*>* visited);
public:
	Graph();
	~Graph();

	list<Vertex*> Vertexes;
	list<LinkedListEdge*> Edges;

	void AddVertex(Vertex* vertex);
	LinkedListEdge* AddLinkedEdge(LinkedListEdge* newEdge);

	Vertex* GetVertex(string name);
	LinkedListEdge* GetLinkedEdge(Vertex* name);

	void FindSP(string start, string end);
};

