#pragma once
class dynamicPrograming
{
private:
	int prices[10];
	int bst[10000];
	void Initial();
	int getPrice(int input);
public:
	dynamicPrograming();
	~dynamicPrograming();

	int GetBestSolution(int input);
};

