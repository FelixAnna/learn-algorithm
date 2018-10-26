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

            //get random pivot
            var index = new Random().Next(start, end);
            var pivot = array[index];

            //swap pivot to the first
            var temp = array[start];
            array[start] = pivot;
            array[index] = temp;

            var i = start;
            for (var j = start + 1; j <= end; j++)
            {
                if (array[j].CompareTo(pivot) <= 0)
                {
                    temp = array[i + 1];
                    array[i + 1] = array[j];
                    array[j] = temp;
                    i++;
                }
            }

            if (i != start)
            {
                temp = array[i];
                array[i] = array[start];
                array[start] = temp;

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
