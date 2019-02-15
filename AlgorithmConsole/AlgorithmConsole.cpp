// AlgorithmConsole.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include "pch.h"
#include "LCS.h"
#include "Graph.h"
#include "dynamicPrograming.h"
#include <iostream>
#include <string>
#include <list>
using namespace std;

void TestDP(dynamicPrograming dp, int value);
void TestDynamicPrograming();
void TestGraph();

int main()
{
	auto dp = dynamicPrograming::dynamicPrograming();
	TestDP(dp, 9999);
	TestDP(dp, 37);
	TestDP(dp, 9);
	TestDP(dp, 3);

	TestDynamicPrograming();
	TestGraph();
	std::cout << "Hello World!\n";
}

void TestDP(dynamicPrograming dp, int value) {
	auto sol = dp.GetBestValue(value);
	std::cout << "Best solution for " << value << " is:" << sol << std::endl;

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

void TestDynamicPrograming()
{
	auto lcs = LCS::LCS("etrsdefabdc", "eraqbec");
	auto length = lcs.FindLCS_Length();
	string results = lcs.GetOneLCS();
	cout << results << endl;

	auto lcsub = LCS::LCS("1abcbc2ertyuid", "34eabcbc4ertyuid5k");
	length = lcsub.FindLCSubstring_Length();
	results = lcsub.GetOneLCSubstring(length);
	cout << results << endl;
}

void TestGraph() {
	auto graph = Graph::Graph();

	//Vertex
	auto a1 = new Vertex("a1", "A one road");
	auto a2 = new Vertex("a2", "A two road");
	auto a3 = new Vertex("a3", "A three road");
	auto a4 = new Vertex("a4", "A four road");
	auto a5 = new Vertex("a5", "A five road");

	auto b1 = new Vertex("b1", "B one road");
	auto b3 = new Vertex("b3", "B three road");
	auto b4 = new Vertex("b4", "B four road");

	auto c1 = new Vertex("c1", "C one road");
	auto c4 = new Vertex("c4", "C four road");

	auto a1e = new Edge(a1, 0); a1e->AppendNext(a2, 3);
	auto a2e = new Edge(a2, 0); a2e->AppendNext(a1, 3)->AppendNext(a3, 3);
	auto a3e = new Edge(a3, 0); a3e->AppendNext(a2, 3)->AppendNext(a4, 3)->AppendNext(b1, 4)->AppendNext(b3, 4);
	auto a4e = new Edge(a4, 0); a4e->AppendNext(a3, 3)->AppendNext(a5, 3)->AppendNext(c1, 5)->AppendNext(b3, 5);
	auto a5e = new Edge(a5, 0); a5e->AppendNext(a4, 3);

	auto b1e = new Edge(b1, 0); b1e->AppendNext(a3, 4);
	auto b3e = new Edge(b3, 0); b3e->AppendNext(a3, 4)->AppendNext(b4, 4)->AppendNext(a4, 5)->AppendNext(c4, 5);
	auto b4e = new Edge(b4, 0); b4e->AppendNext(b3, 4);

	auto c1e = new Edge(c1, 0); c1e->AppendNext(a4, 5);
	auto c4e = new Edge(c4, 0); c4e->AppendNext(b3, 5);

	//Add Vertexes
	graph.AddVertex(a1);
	graph.AddVertex(a2);
	graph.AddVertex(a3);
	graph.AddVertex(a4);
	graph.AddVertex(a5);
	graph.AddVertex(b1);
	graph.AddVertex(b3);
	graph.AddVertex(b4);
	graph.AddVertex(c1);
	graph.AddVertex(c4);

	//Add Edges
	graph.AddLinkedEdge(a1e);
	graph.AddLinkedEdge(a2e);
	graph.AddLinkedEdge(a3e);
	graph.AddLinkedEdge(a4e);
	graph.AddLinkedEdge(a5e);

	graph.AddLinkedEdge(b1e);
	graph.AddLinkedEdge(b3e);
	graph.AddLinkedEdge(b4e);

	graph.AddLinkedEdge(c1e);
	graph.AddLinkedEdge(c4e);
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
