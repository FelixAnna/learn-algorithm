using System;
using System.Collections.Generic;

namespace AlgorithmLibrary.DivideAndConquer
{
    using System.Linq;

    public class QuickSort<T> : ISort<T> where T : IComparable<T>
    {
        public IList<T> Sort(T[] array, int start, int end)
        {
            if (start >= end)
            {
                return array?.ToList();
            }

            var index = new Random(start).Next(start, end);
            var pivot = array[index];

            var i = start;
            for (var j = start; j <= end; j++)
            {
                if (array[j].CompareTo(pivot) < 0)
                {
                    var temp = array[i];
                    array[i] = array[j];
                    array[j] = temp;
                    i++;
                }
            }

            if (i != start)
            {
                var temp2 = array[i];
                array[i] = pivot;
                array[index] = temp2;

                Sort(array, start, i - 1);
                Sort(array, i + 1, end);
            }
            else
            {
                Sort(array, i + 1, end);
            }

            return array;
        }

        public IList<T> Sort(T[] array)
        {
            return Sort(array, 0, array.Length - 1);
        }
    }
}
