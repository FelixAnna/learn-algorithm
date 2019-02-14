#include <string>
#include <iostream>
#include <string.h>
using namespace std;

#pragma once
class LCS
{
private:
	string first;
	string second;
	int rowLength;
	int colLength;
	int** mappings;
public:
	LCS(string first, string second);
	~LCS();

	int FindLCS_Length();
	string GetOneLCS();
};

