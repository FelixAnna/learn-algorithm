using System;
using System.Collections.Generic;
using System.Linq;

namespace AlgorithmLibrary.DivideAndConquer
{
    public class DCSearch<T> : ISearch<T> where T : IComparable<T>
    {
        public T Max(IEnumerable<T> list, int rank)
        {
            return Min(list, list.Count() - rank);
        }

        public T Min(IEnumerable<T> list, int rank = 1)
        {
            return QuickSearch(list.ToArray(), 0, list.Count() - 1, rank);
        }

        private T QuickSearch(T[] array, int startIndex, int endIndex, int rank)
        {
            if (rank - 1 < startIndex || rank - 1 > endIndex)
            {
                return default(T);
            }

			//get random pivot
            int indexOfPivot = new Random().Next(endIndex - startIndex) + startIndex;
            
			//swap element of start and pivot
			T temp = array[startIndex];
            array[startIndex] = array[indexOfPivot];
            array[indexOfPivot] = temp;

            int i = startIndex;
            for (var j = i + 1; j <= endIndex; j++)
            {
                if (array[j].CompareTo(array[startIndex]) <= 0)
                {
                    //swap j and the element after startindex i;
                    Swap(array, i+1, j);
                    i++;
                }
            }

            if (i != startIndex)
            {
                //swap startindex and last swapped index (latest i)
                Swap(array, i, startIndex);

                if (i == rank - 1)
                {
                    return array[i];
                }
                else if (i > rank - 1)
                {
                    return QuickSearch(array, startIndex, i - 1, rank);
                }
                else
                {
                    return QuickSearch(array, i + 1, endIndex, rank);
                }
            }
            else
            {
                if (i == rank - 1)
                {
                    return array[i];
                }
                else
                {
                    return QuickSearch(array, i + 1, endIndex, rank);
                }
            }
        }
		
		
		private void Swap(T[] array, int i, int j){
			var temp = array[i];
			array[i] = array[j];
			array[j] = temp;
		}
    }
}
