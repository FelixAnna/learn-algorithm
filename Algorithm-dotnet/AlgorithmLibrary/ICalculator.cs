namespace AlgorithmLibrary
{
    public interface ICalculator<T> where T : struct
    {
        T Calculate(T value);
    }
}
