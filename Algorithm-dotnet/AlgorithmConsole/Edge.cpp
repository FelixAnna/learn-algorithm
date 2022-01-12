#include "pch.h"
#include "Edge.h"

Edge::Edge(Vertex* start, Vertex* end, double distance)
{
	this->Start = start;
	this->End = end;
	this->Distance = distance;
	this->MinDistance = -1;
}


Edge::~Edge()
{
}
