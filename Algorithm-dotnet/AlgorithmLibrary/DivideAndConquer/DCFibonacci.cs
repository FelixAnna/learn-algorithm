using System.Numerics;
using Entity;

namespace AlgorithmLibrary.DivideAndConquer
{
    public class DcFibonacci : ICalculator<BigInteger>
    {
        public BigInteger Calculate(BigInteger n)
        {
            var basicArray = new[] { new BigInteger[] { 1, 1 }, new BigInteger[] { 1, 0 } };
            var matrix = new BigIntegerMatrix(basicArray);
            var result = new MatrixNPowerCalculator().Power(matrix, n);
            return result[1][0];
        }
    }
}
