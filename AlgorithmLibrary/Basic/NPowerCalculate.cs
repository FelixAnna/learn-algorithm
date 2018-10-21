using System;
using System.Collections.Generic;
using System.Numerics;
using System.Text;

namespace AlgorithmLibrary.Basic
{
    public class NPowerCalculate : INPower<BigInteger>
    {
        public BigInteger Power(BigInteger value, int nPower)
        {
            if (nPower == 0)
            {
                return 1;
            }

            return value * Power(value, nPower - 1);
        }
    }
}
