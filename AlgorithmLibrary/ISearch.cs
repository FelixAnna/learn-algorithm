using System;
using System.Collections.Generic;
using System.Text;

namespace AlgorithmLibrary
{
    public interface ISearch<T>
    {
        T Min(IEnumerable<T> list);
        T Max(IEnumerable<T> list);
    }
}
