#include "pch.h"
#include "Vertex.h"


Vertex::Vertex()
{
	Active = true;
}


Vertex::Vertex(string identifier): Vertex()
{
	this->Identifier = identifier;
	this->Name = identifier;
}

Vertex::Vertex(string identifier, string name) : Vertex()
{
	this->Identifier = identifier;
	this->Name = name;
}

Vertex::~Vertex()
{
}


void Vertex::Enable()
{
	this->Active = true;
	// TODO: �ڴ˴����ʵ�ִ���.
}


void Vertex::Disable()
{
	this->Active = false;
	// TODO: �ڴ˴����ʵ�ִ���.
}
