#include "Edge.h"
#include <list>
#include <iostream>

using namespace std;

#pragma once
class Graph
{
public:
	Graph();
	~Graph();

	list<Vertex*> Vertexes;
	list<Edge*> Edges;

	void AddVertex(Vertex* vertex);
	Edge* AddLinkedEdge(Edge* newEdge);
};

