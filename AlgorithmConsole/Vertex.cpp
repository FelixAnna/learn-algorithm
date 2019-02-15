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
	// TODO: 在此处添加实现代码.
}


void Vertex::Disable()
{
	this->Active = false;
	// TODO: 在此处添加实现代码.
}
