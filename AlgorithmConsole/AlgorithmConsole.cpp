// AlgorithmConsole.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include "pch.h"
#include "LCS.h"
#include "dynamicPrograming.h"
#include <iostream>
using namespace std;

void TestDP(dynamicPrograming dp, int value);

int main()
{
	auto dp=dynamicPrograming::dynamicPrograming();
	TestDP(dp, 9999);
	TestDP(dp, 37);
	TestDP(dp, 9);
	TestDP(dp, 3);

	auto lcs = LCS::LCS("etrsdefabdc", "eraqbec");
	auto length=lcs.FindLCS_Length();
	auto results = lcs.GetOneLCS();
    std::cout << "Hello World!\n"; 
}

void TestDP(dynamicPrograming dp, int value) {
	auto sol = dp.GetBestValue(value);
	std::cout << "Best solution for "<<value<<" is:" << sol << std::endl;

	int* solution = dp.GetBestSolution(value);
	for (int i = 0; i < CalculateLevel; i++)
	{
		if (*(solution + i) <= 0) {
			break;
		}

		std::cout << *(solution + i) << " ";
	}
	std::cout << "Done!\n";
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门提示: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件
