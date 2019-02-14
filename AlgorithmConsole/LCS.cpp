#include "pch.h"
#include "LCS.h"
#include <iostream>
using namespace std;

LCS::LCS(string first, string second)
{
	LCS::first = first;
	LCS::second = second;
}


LCS::~LCS()
{
}

int LCS::FindLCS_Length()
{
	rowLength = first.length() + 1;
	colLength = second.length() + 1;

	mappings = new int*[rowLength];

	for (int i = 0; i < rowLength; i++) {
		mappings[i] = new int[colLength];
		mappings[i][0] = 0;
	}

	for (int i = 0; i < colLength; i++) {
		mappings[0][i] = 0;
	}

	for (int i = 1; i < rowLength; i++)
	{
		for (int j = 1; j < colLength; j++)
		{
			if (first[i - 1] == second[j - 1]) {
				mappings[i][j] = mappings[i - 1][j - 1] + 1;
			}
			else {
				mappings[i][j] = mappings[i][j - 1] > mappings[i - 1][j] ? mappings[i][j - 1] : mappings[i - 1][j];
			}
		}
	}

	return mappings[rowLength - 1][colLength - 1];
}

string LCS::GetOneLCS()
{
	const int resultLength = mappings[rowLength - 1][colLength - 1];
	char* results = new char[resultLength];
	int maxLength = colLength - 1;
	int index = resultLength - 1;
	for (int i = rowLength - 1; i > 0; i--)
	{
		for (int j = maxLength; j > 0; j--)
		{
			if (mappings[i][j] == mappings[i - 1][j - 1] + 1) {
				if (mappings[i][j] == mappings[i][j - 1]) {
					continue;
				}
				else if (mappings[i][j] == mappings[i - 1][j])
				{
					break;
				}
				else 
				{
					results[index--] = first[i - 1];
					maxLength = j - 1;
					break;
				}
			}
		}
	}

	return results;
}
