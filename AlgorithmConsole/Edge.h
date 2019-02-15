#include "Vertex.h"
#pragma once
class Edge
{
public:
	Edge();
	Edge(Vertex* point, double distance);
	~Edge();

	double Distance;
	Vertex* Point;
	Edge* Next;

	Edge* AppendNext(Vertex* point, double distance);
};

