using System;
using System.Collections.Generic;
using System.Text;

namespace AlgorithmLibrary
{
    public class MatrixHelper
    {
        public static int[][] Muptiple(int[][] matFirst, int[][] matSecond)
        {
            if (matFirst[0].Length != matSecond.Length)
            {
                throw new ArgumentException("The width of the first matrix must equal to the height of the second matrix.");
            }

            var height = matSecond[0].Length;
            var resut = new int[matFirst.Length][];
            for (int i = 0; i < matFirst.Length; i++)
            {
                resut[i] = new int[height];
                for (int j = 0; j < matSecond[0].Length; j++)
                {
                    var sum = 0;
                    for (int k = 0; k < matFirst[0].Length; k++)
                    {
                        sum += matFirst[i][k] * matSecond[k][j];
                    }
                    resut[i][j] = sum;
                }
            }

            return resut;
        }
    }
}
