using System;
using System.Collections.Generic;

namespace AlgorithmLibrary
{
    public interface ISearch<T> where T : IComparable<T>
    {
        /// <summary>
        /// Find the rank-th minnum value from the array
        /// </summary>
        /// <param name="list"></param>
        /// <param name="rank"></param>
        /// <returns></returns>
        T Min(IEnumerable<T> list, int rank = 1);

        /// <summary>
        /// Find the rank-th maxnum value from the array
        /// </summary>
        /// <param name="list"></param>
        /// <param name="rank"></param>
        /// <returns></returns>
        T Max(IEnumerable<T> list, int rank = 1);
    }
}
