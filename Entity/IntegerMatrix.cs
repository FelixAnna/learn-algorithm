using System;
using System.Collections.Generic;
using System.Text;

namespace Entity
{
    public class IntegerMatrix : Matrix<int>
    {
        public IntegerMatrix(int[][] mat)
            : base(mat)
        {
        }

        public static IntegerMatrix operator *(IntegerMatrix first, IntegerMatrix second)
        {
            return first.Multiply(second);
        }

        public static IntegerMatrix operator +(IntegerMatrix first, IntegerMatrix second)
        {
            return first.Add(second);
        }

        public static IntegerMatrix operator -(IntegerMatrix first, IntegerMatrix second)
        {
            return first.Reduce(second);
        }

        public static implicit operator IntegerMatrix(int[][] value)
        {
            var mat = new IntegerMatrix(value);
            return mat;
        }

        protected override int Add(int ba, int target)
        {
            return ba + target;
        }

        protected override int Reduce(int ba, int target)
        {
            return ba - target;
        }

        protected override int Multiply(int first, int second)
        {
            return first * second;
        }
    }
}
