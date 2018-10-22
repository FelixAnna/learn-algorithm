using System.Numerics;

namespace AlgorithmLibrary.Basic
{
    public class Fibonacci : ICalculator<BigInteger>
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
