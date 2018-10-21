using System;
using System.Collections.Generic;
using System.Text;

namespace AlgorithmLibrary
{
    interface ICalculate<T> where T : struct
    {
        T Calculate(T value);
    }
}
