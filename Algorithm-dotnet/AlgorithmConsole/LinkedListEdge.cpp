#include "pch.h"
#include "LinkedListEdge.h"


LinkedListEdge::LinkedListEdge()
{
	Distance = 0.0;
}

LinkedListEdge::LinkedListEdge(Vertex * node, double distance)
{
	this->Node = node;
	this->Distance = distance;
}


LinkedListEdge::~LinkedListEdge()
{
}

LinkedListEdge* LinkedListEdge::AppendNext(Vertex * node, double distance)
{
	this->Next = new LinkedListEdge(node, distance);
	return this->Next;
}
