#define PriceLevel 10
#define CalculateLevel 10000
#pragma once
class dynamicPrograming
{
private:
	int prices[PriceLevel];
	int bstStrategy[CalculateLevel];
	int bst[CalculateLevel];
	void Initial();
	int getPrice(int input);
public:
	dynamicPrograming();
	~dynamicPrograming();

	int GetBestValue(int input);
	int* GetBestSolution(int input);
};

