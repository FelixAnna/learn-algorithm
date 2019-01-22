using System;
using System.Collections.Generic;
using System.Linq;

namespace AlgorithmLibrary.Basic
{
    public class BinaryTree<T> where T : IComparable<T>
    {
        public BinaryTree<T> Left { get; set; }
        public BinaryTree<T> Right { get; set; }
        public T Value { get; set; }
        private BinaryTree()
        {

        }

        public BinaryTree(params T[] data)
        {
            var tree = BuidTree(data);

            this.Value = tree.Value;
            this.Left = tree.Left;
            this.Right = tree.Right;
        }

        private BinaryTree<T> BuidTree(IEnumerable<T> data)
        {
            var tree = new BinaryTree<T>();
            if (!data.Any())
            {
                return null;
            }
            else if (data.Count() == 1)
            {
                tree.Left = null;
                tree.Value = data.First();
                tree.Right = null;
            }
            else
            {
                var middleIndex = data.Count() / 2;

                tree.Left = BuidTree(data.Take(middleIndex));
                tree.Value = data.ElementAt(middleIndex);
                tree.Right = BuidTree(data.Skip(middleIndex + 1));
            }

            return tree;
        }

        public IEnumerable<T> Visit()
        {
            var result = new List<T>();
            if (this.Value != null)
            {
                var leftValues = this.Left?.Visit();
                if (leftValues != null && leftValues.Any())
                {
                    result.AddRange(leftValues);
                }

                result.Add(this.Value);

                var rightValues = this.Right?.Visit();
                if (rightValues != null && rightValues.Any())
                {
                    result.AddRange(rightValues);
                }
            }

            return result;
        }

        public BinaryTree<T> Find(T value)
        {
            if (this.Value == null)
            {
                return null;
            }

            var compare = this.Value.CompareTo(value);
            if (compare == 0)
            {
                return this;
            }
            else if (compare > 0)
            {
                return this.Left?.Find(value);
            }
            else
            {
                return this.Right?.Find(value);
            }
        }

        public void Insert(T value)
        {
            var currentNode = this;
            var parentNode = currentNode = this;
            do
            {
                var compare = currentNode.Value.CompareTo(value);
                if (compare > 0)
                {
                    parentNode = currentNode;
                    currentNode = currentNode.Left;
                }
                else if (compare < 0)
                {
                    parentNode = currentNode;
                    currentNode = currentNode.Right;
                }
                else
                {
                    break;
                }
            }
            while (currentNode != null);

            if (parentNode.Value.CompareTo(value) > 0)
            {
                var temp = parentNode.Left;
                var newNode = new BinaryTree<T>(value);
                parentNode.Left = newNode;
                newNode.Left = temp;
            }
            else
            {
                var temp = parentNode.Right;
                var newNode = new BinaryTree<T>(value);
                parentNode.Right = newNode;
                newNode.Right = temp;
            }
        }
    }
}
