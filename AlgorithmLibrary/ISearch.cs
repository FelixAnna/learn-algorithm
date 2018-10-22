using System.Collections.Generic;

namespace AlgorithmLibrary
{
    public interface ISearch<T>
    {
        T Min(IEnumerable<T> list);

        T Max(IEnumerable<T> list);
    }
}
