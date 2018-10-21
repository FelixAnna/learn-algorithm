using System;
using System.Collections.Generic;
using System.Text;

namespace AlgorithmLibrary
{
    interface INPower<T> where T : struct
    {
        T Power(T value, int nPower);
    }
}
