#include "pch.h"
#include "Edge.h"


Edge::Edge()
{
	Distance = 0.0;
}

Edge::Edge(Vertex * point, double distance)
{
	this->Point = point;
	this->Distance = distance;
}


Edge::~Edge()
{
}

Edge* Edge::AppendNext(Vertex * point, double distance)
{
	this->Next = new Edge(point, distance);
	return this->Next;
}
