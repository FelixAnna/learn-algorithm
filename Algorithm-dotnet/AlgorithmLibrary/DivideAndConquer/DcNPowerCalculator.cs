using System.Numerics;

namespace AlgorithmLibrary.DivideAndConquer
{
    public class DcNPowerCalculator : INPower<BigInteger>
    {
        public BigInteger Power(BigInteger value, int n)
        {
            if (n == 0)
            {
                return 1;
            }

            if (n % 2 == 0)
            {
                var halfValue = Power(value: value, n: n / 2);
                return halfValue * halfValue;
            }
            else
            {
                var halfValue = Power(value: value, n: n / 2);
                return halfValue * halfValue * value;
            }
        }
    }
}
