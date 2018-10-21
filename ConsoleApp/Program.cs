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
                array[i] = (int)(new Random(i).Next(arrayLength * 10));
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
            var value = new NPowerCalculate().Power(5, 999);
            st.Stop();
            Console.WriteLine(value);
            Console.WriteLine(st.ElapsedMilliseconds);

            st.Restart();
            value = new DCNPowerCalculate().Power(5, 999);
            st.Stop();
            Console.WriteLine(value);
            Console.WriteLine(st.ElapsedMilliseconds);

            var matrix = new int[][] { new int[] { 1, 1 }, new int[] { 1, 0 } };
            var mat = MatrixHelper.Muptiple(matrix, matrix);
            mat= MatrixHelper.Muptiple(mat, matrix);
            mat = MatrixHelper.Muptiple(mat, matrix);
            mat = MatrixHelper.Muptiple(mat, matrix);
            mat = MatrixHelper.Muptiple(mat, matrix);
            Print(mat);
            Console.ReadLine();
        }

        static void Print<T>(IEnumerable<T> array)
        {
            foreach (var item in array)
            {
                if (item is IEnumerable)
                {
                    foreach (var subItem in (item as IEnumerable))
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
