using System;
using System.Collections.Generic;

namespace AlgorithmLibrary.DivideAndConquer
{
    using System.Linq;

    public class QuickSort<T> : ISort<T> where T : IComparable<T>
    {
        private IList<T> Sort(T[] array, int start, int end)
        {
            if (start >= end)
            {
                return array?.ToList();
            }

            //get random pivot
            var pivotIndex = new Random().Next(start, end);
            var pivot = array[pivotIndex];

            //swap pivot to the first
            Swap(array, start, pivotIndex);

            var i = start;
            for (var j = start + 1; j <= end; j++)
            {
                if (array[j].CompareTo(pivot) <= 0)
                {
                    //swap j and the element after startindex i;
                    Swap(array, i+1, j);
                    i++;
                }
            }

            if (i != start)
            {
                //swap
                Swap(array, i, start);

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
		
		public void Swap(T[] array, int i, int j){
			var temp = array[i];
			array[i] = array[j];
			array[j] = temp;
		}
    }
}
