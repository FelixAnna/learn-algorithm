#include "pch.h"
#include "dynamicPrograming.h"
#include <iostream>

void dynamicPrograming::Initial()
{
	/*for (int i = 0; i < 10; i++) {
		prices[i] = i * i;
	}*/

	prices[0] = 0;
	prices[1] = 2;
	prices[2] = 3;
	prices[3] = 7;
	prices[4] = 9;
	prices[5] = 12;
	prices[6] = 15;
	prices[7] = 17;
	prices[8] = 18;
	prices[9] = 19;

	for (int j = 0; j < CalculateLevel; j++) {
		bst[j] = 0;
		bstStrategy[j] = 0;
	}
}

dynamicPrograming::dynamicPrograming()
{
	Initial();
}

int dynamicPrograming::GetBestValue(int input)
{
	if (input == 0) {
		return 0;
	}

	if (bst[input] > 0) {
		return bst[input];
	}

	auto cut = 0;
	for (int i = 1; i <= input; i++) {

		auto solution = 0;
		for (int j = 1; j <= i; j++) {
			auto new_solution = getPrice(j) + GetBestValue(i - j);
			if (new_solution > solution) {
				solution = new_solution;
				cut = j;
			}
		}

		bstStrategy[i] = cut;
		bst[i] = solution;
	}

	return bst[input];
}

int* dynamicPrograming::GetBestSolution(int input) {
	int solution[10000];
	int * result = solution;
	int remaining = input;
	for (int i = 0; i < CalculateLevel; i++) {
		if (bstStrategy[remaining] < 0) {
			break;
		}

		* result = bstStrategy[remaining];
		remaining -= bstStrategy[remaining];
		result++;
	}

	return solution;
}
int dynamicPrograming::getPrice(int input) {
	if (input >= PriceLevel) {
		return 0;
	}

	return prices[input];
}

dynamicPrograming::~dynamicPrograming()
{
}
