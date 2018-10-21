using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace AlgorithmLibrary.Basic
{
    public class NormalSort<T> : ISort<T> where T : IComparable<T>
    {
        public IList<T> Sort(T[] array)
        {
            if (array == null || array.Length <= 1)
            {
                return array?.ToList();
            }

            for (int i = 0; i < array.Length; i++)
            {
                for (int j = i + 1; j < array.Length; j++)
                {
                    if (array[i].CompareTo(array[j]) > 0)
                    {
                        var temp = array[i];
                        array[i] = array[j];
                        array[j] = temp;
                    }
                }
            }

            return array.ToList();
        }
    }
}
