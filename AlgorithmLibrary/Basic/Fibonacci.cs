using System;
using System.Collections.Generic;
using System.Numerics;
using System.Text;

namespace AlgorithmLibrary.Basic
{
    public class Fibonacci : ICalculate<BigInteger>
    {
        public BigInteger Calculate(BigInteger n)
        {
            if (n <= 1)
            {
                return 1;
            }

            return Calculate(n - 1) + Calculate(n - 2);
        }
    }
}
