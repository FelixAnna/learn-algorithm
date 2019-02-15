#include "Vertex.h"
#pragma once
class LinkedListEdge
{
public:
	LinkedListEdge();
	LinkedListEdge(Vertex* point, double distance);
	~LinkedListEdge();

	double Distance;
	Vertex* Node;
	LinkedListEdge* Next;

	LinkedListEdge* AppendNext(Vertex* point, double distance);
};

