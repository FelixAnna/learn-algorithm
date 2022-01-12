using System.Numerics;

namespace AlgorithmLibrary.Basic
{
    public class NPowerCalculator : INPower<BigInteger>
    {
        public BigInteger Power(BigInteger value, int n)
        {
            if (n == 0)
            {
                return 1;
            }

            return value * Power(value: value, n: n - 1);
        }
    }
}
