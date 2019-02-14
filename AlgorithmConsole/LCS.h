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

	/*
	Longest common sequence
	*/
	int FindLCS_Length();

	/*
	Get one of the common sequence
	*/
	string GetOneLCS();

	/*
	Longest common sequence
	*/
	int FindLCSubstring_Length();

	/*
	Get one of the common sequence
	*/
	string GetOneLCSubstring(int length);
};

