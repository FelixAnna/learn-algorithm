using AlgorithmLibrary;
using AlgorithmLibrary.Basic;
using AlgorithmLibrary.DivideAndConquer;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;

namespace ConsoleApp
{
    using Entity;

    class Program
    {
        static void Main(string[] args)
        {
            int arrayLength = 8;
            var array = new int[arrayLength];
            for (int i = 0; i < arrayLength; i++)
            {
                array[i] = new Random().Next(arrayLength * 10);
            }
            //Print(array);
            var st = new Stopwatch();
            //st.Start();
            //var result = new MergedSort<int>().Sort(array);
            //st.Stop();
            //Console.WriteLine(st.ElapsedMilliseconds);
            //Print(result);

            //st.Restart();
            //result = new NormalSort<int>().Sort(array);
            //st.Stop();
            //Console.WriteLine(st.ElapsedMilliseconds);
            //Print(result);

            //st.Restart();
            //result = new QuickSort<int>().Sort(array);
            //st.Stop();
            //Console.WriteLine(st.ElapsedMilliseconds);
            //Print(result);

            st.Restart();
            var result = new LinearSort().Sort(array);
            st.Stop();
            Console.WriteLine(st.ElapsedMilliseconds);
            Print(result);

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

            int deepth = 20;
            st.Restart();
            value = new Fibonacci().Calculate(deepth);
            st.Stop();
            Console.WriteLine(value);
            Console.WriteLine(st.ElapsedMilliseconds);

            st.Restart();
            value = new DcFibonacci().Calculate(deepth);
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


            var matrix2 = new BigIntegerMatrix(new System.Numerics.BigInteger[arrayLength][]);
            for (var i = 0; i < arrayLength; i++)
            {
                matrix2[i] = new System.Numerics.BigInteger[arrayLength];
                for (var j = 0; j < arrayLength; j++)
                {
                    matrix2[i][j] = new Random(i*j).Next(100);
                }
            }

            st.Restart();
            var resultMat = matrix2 * matrix2;
            st.Stop();
            Print(resultMat.AsEnumerable());
            Console.WriteLine(st.ElapsedMilliseconds);

            st.Restart();
            resultMat = BigIntegerMatrix.StrassenMultiply(matrix2, matrix2);
            st.Stop();
            Print(resultMat.AsEnumerable());
            Console.WriteLine(st.ElapsedMilliseconds);

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
