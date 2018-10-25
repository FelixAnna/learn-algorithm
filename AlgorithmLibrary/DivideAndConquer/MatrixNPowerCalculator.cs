namespace AlgorithmLibrary.DivideAndConquer
{
    using Entity;
    using System.Numerics;

    public class MatrixNPowerCalculator
    {
        public IntegerMatrix Power(IntegerMatrix value, int n)
        {
            if (n == 0)
            {
                return value;
            }

            if (n % 2 == 0)
            {
                var halfValue = Power(value, n / 2);
                return halfValue * halfValue;
            }
            else
            {
                var halfValue = Power(value, n / 2);
                return halfValue * halfValue * value;
            }
        }

        public BigIntegerMatrix Power(BigIntegerMatrix value, BigInteger n)
        {
            if (n <= 1)
            {
                return value;
            }

            if (n % 2 == 0)
            {
                var halfValue = Power(value, n / 2);
                return halfValue * halfValue;
            }
            else
            {
                var halfValue = Power(value, n / 2);
                return halfValue * halfValue * value;
            }
        }
    }
}
