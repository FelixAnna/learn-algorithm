namespace AlgorithmLibrary
{
    public interface INPower<T> where T : struct
    {
        T Power(T value, int n);
    }
}
