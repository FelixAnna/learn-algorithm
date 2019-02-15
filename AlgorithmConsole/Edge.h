#include "Vertex.h"
#pragma once
class Edge
{
public:
	Edge(Vertex* start, Vertex* end, double distance);
	~Edge();

	Vertex* Start;
	Vertex* End;
	double Distance;

	double MinDistance;
	Vertex* MinVertex;
};

