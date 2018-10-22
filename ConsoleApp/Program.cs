using AlgorithmLibrary;
using AlgorithmLibrary.Basic;
using AlgorithmLibrary.DivideAndConquer;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;

namespace ConsoleApp
{
    class Program
    {
        static void Main(string[] args)
        {
            int arrayLength = 1000;
            var array = new int[arrayLength];
            for (int i = 0; i < arrayLength; i++)
            {
                array[i] = new Random(i).Next(arrayLength * 10);
            }
            //Print(array);
            var st = new Stopwatch();
            st.Start();
            var result = new MergedSort<int>().Sort(array);
            st.Stop();
            Console.WriteLine(st.ElapsedMilliseconds);
            //Print(result);

            st.Restart();
            result = new NormalSort<int>().Sort(array);
            st.Stop();
            Console.WriteLine(st.ElapsedMilliseconds);
            //Print(result);

            st.Restart();
            var value = new NPowerCalculator().Power(5, 999);
            st.Stop();
            Console.WriteLine(value);
            Console.WriteLine(st.ElapsedMilliseconds);

            st.Restart();
            value = new DcNPowerCalculator().Power(5, 999);
            st.Stop();
            Console.WriteLine(value);
            Console.WriteLine(st.ElapsedMilliseconds);

            var matrix = new[] { new[] { 1, 1 }, new[] { 1, 0 } };
            var mat = MatrixHelper.Multiply(matrix, matrix);
            mat = MatrixHelper.Multiply(mat, matrix);
            mat = MatrixHelper.Multiply(mat, matrix);
            mat = MatrixHelper.Multiply(mat, matrix);
            mat = MatrixHelper.Multiply(mat, matrix);
            Print(mat);
            Console.ReadLine();
        }

        private static void Print<T>(IEnumerable<T> array)
        {
            foreach (var item in array)
            {
                if (item is IEnumerable enumerable)
                {
                    foreach (var subItem in enumerable)
                    {
                        Console.Write(subItem);
                        Console.Write(",");
                    }

                    Console.WriteLine();
                }
                else
                {
                    Console.Write(item);
                    Console.Write(",");
                }
            }

            Console.WriteLine();
        }
    }
}
