using System;
using System.Collections.Generic;
using System.Linq;

namespace AlgorithmLibrary.DivideAndConquer
{
    /// <summary>
    /// Merge Sort
    /// 1. Divide in to 2 (recursive)
    /// 2. Sort
    /// 3. Merge
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public class MergedSort<T> : ISort<T> where T : IComparable<T>
    {
        public IList<T> Sort(T[] array)
        {
            if (array == null || array.Length <= 1)
            {
                return array?.ToList();
            }

            var firstArray = Sort(array.Take(array.Length / 2).ToArray());
            var secondArray = Sort(array.Skip(array.Length / 2).ToArray());

            return Merge(firstArray, secondArray);
        }

        private IList<T> Merge(IList<T> firstArray, IList<T> secondArray)
        {
            if (firstArray == null || firstArray.Count == 0)
            {
                return secondArray;
            }

            if (secondArray == null || secondArray.Count == 0)
            {
                return firstArray;
            }

            T[] combinedArray = new T[firstArray.Count + secondArray.Count];
            int firstIndex, secondIndex, combinedIndex;
            firstIndex = secondIndex = combinedIndex = 0;

            do
            {
                if (firstIndex == firstArray.Count)
                {
                    combinedArray[combinedIndex++] = secondArray[secondIndex++];
                }
                else if (secondIndex == secondArray.Count)
                {
                    combinedArray[combinedIndex++] = firstArray[firstIndex++];
                }
                else
                {
                    var first = firstArray[firstIndex];
                    var last = secondArray[secondIndex];

                    if (first.CompareTo(last) > 0)
                    {
                        secondIndex++;
                        combinedArray[combinedIndex++] = last;
                    }
                    else
                    {
                        firstIndex++;
                        combinedArray[combinedIndex++] = first;
                    }
                }
            }
            while (firstIndex < firstArray.Count || secondIndex < secondArray.Count);

            return combinedArray;
        }
    }
}
