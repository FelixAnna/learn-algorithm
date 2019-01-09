using System;
using System.Collections.Generic;
using System.Text;

namespace Entity
{
    public class IntegerMatrix : Matrix<int>
    {
        public IntegerMatrix(int[][] mat)
            : base(mat: mat)
        {
        }

        public static IntegerMatrix operator *(IntegerMatrix first, IntegerMatrix second)
        {
            return first.Multiply(target: second);
        }

        public static IntegerMatrix operator +(IntegerMatrix first, IntegerMatrix second)
        {
            return first.Add(target: second);
        }

        public static IntegerMatrix operator -(IntegerMatrix first, IntegerMatrix second)
        {
            return first.Reduce(target: second);
        }

        public static implicit operator IntegerMatrix(int[][] value)
        {
            var mat = new IntegerMatrix(mat: value);
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
