using System.Collections.Generic;

namespace AlgorithmLibrary.DivideAndConquer
{
    public class LinearSort : ISort<int>
    {
        public IList<int> Sort(int[] array)
        {
            return Sort(array, 0, 10000);
        }

        private IList<int> Sort(int[] array, int start, int end)
        {
            var countT = new int[end - start + 1];

            foreach (var val in array)
            {
                countT[val - start] = countT[val - start] + 1;
            }

            for (var i = 1; i < countT.Length; i++)
            {
                countT[i] = countT[i] + countT[i - 1];
            }

            var result = new int[array.Length];
            for (var i = array.Length - 1; i >= 0; i--)
            {
                var pos = countT[array[i]];
                result[pos - 1] = array[i];
                countT[array[i]] = pos--;
            }

            return result;
        }
    }
}
