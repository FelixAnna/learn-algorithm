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
	prices[8] = 20;
	prices[9] = 30;


	for (int j = 0; j < 10000; j++) {
		bst[j] = 0;
	}
}

dynamicPrograming::dynamicPrograming()
{
	Initial();
}

int dynamicPrograming::GetBestSolution(int input)
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
			auto new_solution = getPrice(j) + GetBestSolution(i - j);
			if (new_solution > solution) {
				solution = new_solution;
				cut = j;
			}
		}

		std::cout << cut << std::endl;
		bst[i] = solution;
	}

	return bst[input];
}

int dynamicPrograming::getPrice(int input) {
	if (input >= 10) {
		return 0;
	}

	return prices[input];
}

dynamicPrograming::~dynamicPrograming()
{
}
