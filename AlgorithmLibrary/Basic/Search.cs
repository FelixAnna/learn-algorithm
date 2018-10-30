using AlgorithmLibrary.DivideAndConquer;
using System;
using System.Collections.Generic;
using System.Linq;

namespace AlgorithmLibrary.Basic
{
    public class Search<T> : ISearch<T> where T : IComparable<T>
    {
        public T Max(IEnumerable<T> list, int rank)
        {
            return Min(list, list.Count() - rank);
        }

        public T Min(IEnumerable<T> list, int rank)
        {
            if (rank <= 0 || rank >= list.Count())
            {
                throw new IndexOutOfRangeException($"{rank} is not a valid index.");
            }
            var sortedList = new QuickSort<T>().Sort(list.ToArray());
            return sortedList[rank - 1];
        }
    }
}
