using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AlgorithmLibrary.DynamicProgramming
{
    public class CutRodProblem
    {
        private Dictionary<int, float> basicUnitPrice;
        private Dictionary<int, float> bestTotalPrice;
        private Dictionary<int, int> bestCutStatergy;

        public CutRodProblem()
        {
            basicUnitPrice = new();
            basicUnitPrice.Add(1, 2);
            basicUnitPrice.Add(2, 3);
            basicUnitPrice.Add(3, 7);
            basicUnitPrice.Add(4, 9);
            basicUnitPrice.Add(5, 12);
            basicUnitPrice.Add(6, 15);
            basicUnitPrice.Add(7, 17);
            basicUnitPrice.Add(8, 18);
            basicUnitPrice.Add(9, 19);

            bestTotalPrice = new();
            bestTotalPrice.Add(0, 0);

            bestCutStatergy = new();
        }

        public void FindBestCutStatergy(int totalLength)
        {
            float bestPrice = 0;
            if (!bestTotalPrice.ContainsKey(totalLength))
            {
                for (int i = 0; i <= totalLength; i++)
                {
                    if (bestTotalPrice.ContainsKey(i))
                    {
                        continue;
                    }

                    var bestPriceCurrent = basicUnitPrice.Where(x => i >= x.Key).Select(x => new { x.Key, Price = bestTotalPrice[i - x.Key] + x.Value }).OrderByDescending(x=>x.Price).FirstOrDefault();

                    bestTotalPrice.Add(i, bestPriceCurrent.Price);
                    bestCutStatergy.Add(i, bestPriceCurrent.Key);
                }
            }
            bestPrice = bestTotalPrice[totalLength];

            var remainingLength = totalLength;
            do
            {
                var cutPoint = bestCutStatergy[remainingLength];
                Console.Write(cutPoint);
                Console.Write(",");
                remainingLength -= cutPoint;
            } while (remainingLength > 0);
            return;
        }
    }
}
