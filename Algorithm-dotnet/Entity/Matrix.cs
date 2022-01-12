using System;

namespace Entity
{
    public abstract class Matrix<T> where T : struct
    {
        private readonly T[][] matrix;

        protected Matrix(T[][] mat)
        {
            matrix = mat;
        }

        public int Length => matrix.Length;

        public T[] this[int index]
        {
            get => matrix[index];
            set => matrix[index] = value;
        }

        public T this[int row, int col]
        {
            get => matrix[row][col];
            set => matrix[row][col] = value;
        }

        public T[][] Multiply(Matrix<T> target)
        {
            if (this[0].Length != target.Length)
            {
                throw new ArgumentException("The width of the first matrix must equal to the height of the second matrix.");
            }

            int width = target[0].Length;
            var result = new T[this.Length][];
            for (var i = 0; i < this.Length; i++)
            {
                result[i] = new T[width];
                for (var j = 0; j < target[0].Length; j++)
                {
                    var sum = default(T);
                    for (var k = 0; k < this[0].Length; k++)
                    {
                        sum = Add(sum, Multiply(this[i][k], target[k][j]));
                    }

                    result[i][j] = sum;
                }
            }

            return result;
        }

        public T[][] Add(Matrix<T> target)
        {
            if (this.Length != target.Length || this[0].Length != target[0].Length)
            {
                throw new ArgumentException("The two matrix must have the same size.");
            }

            var width = this[0].Length;
            var result = new T[this.Length][];
            for (var i = 0; i < this.Length; i++)
            {
                result[i] = new T[width];
                for (var j = 0; j < this[0].Length; j++)
                {
                    result[i][j] = Add(this[i][j], target[i][j]);
                }
            }

            return result;
        }

        public T[][] Reduce(Matrix<T> target)
        {
            if (this.Length != target.Length || this[0].Length != target[0].Length)
            {
                throw new ArgumentException("The two matrix must have the same size.");
            }

            var width = this[0].Length;
            var result = new T[this.Length][];
            for (var i = 0; i < this.Length; i++)
            {
                result[i] = new T[width];
                for (var j = 0; j < this[0].Length; j++)
                {
                    result[i][j] = Reduce(this[i][j], target[i][j]);
                }
            }

            return result;
        }

        public T[][] AsEnumerable()
        {
            return matrix;
        }

        protected abstract T Multiply(T first, T second);

        protected abstract T Add(T ba, T target);

        protected abstract T Reduce(T ba, T target);
    }
}
