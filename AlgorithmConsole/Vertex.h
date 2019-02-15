#include <iostream>
using namespace std;

#pragma once
class Vertex
{
public:
	Vertex();
	Vertex(string identifier);
	Vertex(string identifier, string name);
	~Vertex();

	string Identifier;
	string Name;
	bool Active;

	void Enable();
	void Disable();
};

