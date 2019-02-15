#include <iostream>
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

	/*
	Longest common sequence
	*/
	int FindLCS_Length();

	/*
	Get one of the common sequence
	*/
	string GetOneLCS();

	/*
	Longest common substring
	*/
	int FindLCSubstring_Length();

	/*
	Get one of the common substring
	*/
	string GetOneLCSubstring(int length);
};

