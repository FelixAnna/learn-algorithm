using System.Numerics;

namespace AlgorithmLibrary.Basic
{
    public class Fibonacci : ICalculator<BigInteger>
    {
        public BigInteger Calculate(BigInteger n)
        {
            if (n <= 2)
            {
                return 1;
            }

            return Calculate(n: n - 1) + Calculate(n: n - 2);
        }

        public BigInteger CalculateUp(BigInteger n)
        {
            if (n <= 1)
            {
                return 1;
            }
            BigInteger[] values = new BigInteger[(int)n];
            values[0] = 1;
            values[1] = 1;

            for (var index = 2; index < n; index++)
            {
                values[index] = values[index - 1] + values[index - 2];
            }

            return values[(int)n - 1];
        }
    }
}
