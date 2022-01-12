using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AlgorithmLibrary.DynamicProgramming
{
    public class StepProblem
    {
        private readonly Dictionary<int, long> waysDict;
        private readonly int[] seeds;

        public StepProblem()
        {
            seeds = Enumerable.Range(1, 2).ToArray();
            waysDict = new Dictionary<int, long>();
            waysDict[0] = 0;
            waysDict[1] = 1;
        }

        public long FindWays(int steps)
        {
            if(waysDict.ContainsKey(steps))
            {
                return waysDict[steps];
            }

            for (int i=0; i<= steps; i++)
            {
                if (waysDict.ContainsKey(i))
                {
                    continue;
                }
                //way[n] = way[n-1] + way [n-2] + way[n-3]
                //0: defalut-0
                //1: 0 - 1
                //2: 0, 1 - 2
                //3: 0, 1, 2 - 2 + 1 + 1
                //4: 1,2,3 - 1,2,4  1111, 121,211,112,13,31,22
                var results = seeds.Select(sd=>i-sd)
                    .Where(subWay => subWay>=0)
                    .Distinct()
                    .Sum(subWay =>
                {
                    if(subWay == 0)
                    {
                        return 1;
                    }

                    return waysDict[subWay];
                });

                waysDict[i] = results;
            }

            return waysDict[steps];
        }
    }
}
