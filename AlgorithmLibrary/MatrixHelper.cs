using System;

namespace AlgorithmLibrary
{
    public class MatrixHelper
    {
        public static int[][] Multiply(int[][] matFirst, int[][] matSecond)
        {
            if (matFirst[0].Length != matSecond.Length)
            {
                throw new ArgumentException("The width of the first matrix must equal to the height of the second matrix.");
            }

            var height = matSecond[0].Length;
            var result = new int[matFirst.Length][];
            for (var i = 0; i < matFirst.Length; i++)
            {
                result[i] = new int[height];
                for (var j = 0; j < matSecond[0].Length; j++)
                {
                    var sum = 0;
                    for (var k = 0; k < matFirst[0].Length; k++)
                    {
                        sum += matFirst[i][k] * matSecond[k][j];
                    }
                    result[i][j] = sum;
                }
            }

            return result;
        }
    }
}
