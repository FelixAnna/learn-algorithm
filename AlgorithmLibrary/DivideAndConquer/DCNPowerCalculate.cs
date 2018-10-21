using System.Numerics;

namespace AlgorithmLibrary.DivideAndConquer
{
    public class DCNPowerCalculate : INPower<BigInteger>
    {
        public BigInteger Power(BigInteger value, int nPower)
        {
            if (nPower == 0)
            {
                return 1;
            }

            if (nPower % 2 == 0)
            {
                var halfValue = Power(value, nPower / 2);
                return halfValue * halfValue;
            }
            else
            {
                var halfValue = Power(value, nPower / 2);
                return halfValue * halfValue * value;
            }
        }
    }
}
