using System;
using System.Collections.Generic;
using System.Text;

namespace AlgorithmLibrary
{
    public interface ISort<T> where T : IComparable<T>
    {
        IList<T> Sort(T[] array);
    }
}
