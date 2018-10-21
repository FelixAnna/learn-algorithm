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

            var arrayOne = Sort(array.Take(array.Length / 2).ToArray());
            var arrayTwo = Sort(array.Skip(array.Length / 2).ToArray());

            return Merge(arrayOne, arrayTwo);
        }

        private IList<T> Merge(IList<T> firstArray, IList<T> lastArray)
        {
            if (firstArray == null || firstArray.Count == 0)
            {
                return lastArray;
            }

            if (lastArray == null || lastArray.Count == 0)
            {
                return firstArray;
            }

            T[] result = new T[firstArray.Count + lastArray.Count];
            int i, j, k;
            i = j = k = 0;

            do
            {
                if (i == firstArray.Count)
                {
                    result[k++] = lastArray[j++];
                }
                else if (j == lastArray.Count)
                {
                    result[k++] = firstArray[i++];
                }
                else
                {
                    var first = firstArray[i];
                    var last = lastArray[j];

                    if (first.CompareTo(last) > 0)
                    {
                        j++;
                        result[k++] = last;
                    }
                    else
                    {
                        i++;
                        result[k++] = first;
                    }
                }
            } while (i < firstArray.Count || j < lastArray.Count);

            return result;
        }
    }
}
