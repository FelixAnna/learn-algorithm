using System;
using System.Collections.Generic;
using System.Linq;

namespace AlgorithmLibrary.Basic
{
    public class BinaryTree<T> where T : IComparable<T>
    {
        public BinaryTree<T> Left { get; set; }
        public BinaryTree<T> Right { get; set; }
        public BinaryTree<T> Parent { get; set; }
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
                tree.Parent = null;
            }
            else
            {
                var middleIndex = data.Count() / 2;

                tree.Left = BuidTree(data.Take(middleIndex));
                tree.Value = data.ElementAt(middleIndex);
                tree.Right = BuidTree(data.Skip(middleIndex + 1));

                if (tree.Left != null)
                {
                    tree.Left.Parent = tree;
                }

                if (tree.Right != null)
                {
                    tree.Right.Parent = tree;
                }
            }

            return tree;
        }

        public IEnumerable<T> Visit()
        {
            var result = new List<T>();

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

            return result;
        }

        public BinaryTree<T> Find(T value)
        {
            var compare = this.Value?.CompareTo(value);
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
                var newNode = new BinaryTree<T>(value)
                {
                    Parent = parentNode
                };
                parentNode.Left = newNode;
                newNode.Left = temp;
                if (temp != null)
                {
                    temp.Parent = newNode;
                }
            }
            else
            {
                var temp = parentNode.Right;
                var newNode = new BinaryTree<T>(value)
                {
                    Parent = parentNode
                };
                parentNode.Right = newNode;
                newNode.Right = temp;
                if (temp != null)
                {
                    temp.Parent = newNode;
                }
            }
        }

        public void Delete(T value)
        {
            var currentNode = this;
            do
            {
                var compareResult = currentNode.Value.CompareTo(value);
                if (compareResult > 0)
                {
                    currentNode = currentNode.Left;
                }
                else if (compareResult < 0)
                {
                    currentNode = currentNode.Right;
                }
                else
                {
                    break;
                }
            } while (currentNode != null && currentNode.Value != null);

            if (currentNode != null)
            {
                if (currentNode.Left == null && currentNode.Right == null)
                {
                    if (currentNode.Parent?.Left == currentNode)
                    {
                        currentNode.Parent.Left = null;
                    }
                    else if (currentNode.Parent?.Right == currentNode)
                    {
                        currentNode.Parent.Right = null;
                    }
                    else if (currentNode.Parent == null)  //only one node
                    {
                        currentNode.Value = default(T);
                    }
                }
                else if (currentNode.Left == null)
                {
                    currentNode.Right.Parent = currentNode.Parent;

                    if (currentNode.Parent?.Left == currentNode)
                    {
                        currentNode.Parent.Left = currentNode.Right;
                    }
                    else if (currentNode.Parent?.Right == currentNode)
                    {
                        currentNode.Parent.Right = currentNode.Right;
                    }
                    else if (currentNode.Parent == null)  //only one node
                    {
                        currentNode.Value = currentNode.Right.Value;
                        currentNode.Left = currentNode.Right.Left;
                        currentNode.Right = currentNode.Right.Right;
                    }
                }
                else if (currentNode.Right == null)
                {
                    currentNode.Left.Parent = currentNode.Parent;

                    if (currentNode.Parent?.Left == currentNode)
                    {
                        currentNode.Parent.Left = currentNode.Left;
                    }
                    else if (currentNode.Parent?.Right == currentNode)
                    {
                        currentNode.Parent.Right = currentNode.Left;
                    }
                    else if (currentNode.Parent == null)  //only one node
                    {
                        currentNode.Value = currentNode.Left.Value;
                        currentNode.Left = currentNode.Left.Left;
                        currentNode.Right = currentNode.Left.Right;
                    }
                }
                else
                {
                    var leftest = GetLeftest(currentNode.Right);
                    currentNode.Value = leftest.Value;
                    if(leftest.Parent?.Left?.Value?.CompareTo(leftest.Value) == 0)
                    {
                        leftest.Parent.Left = leftest.Right;
                    }
                    else
                    {
                        leftest.Parent.Right = leftest.Right;
                    }

                    if (leftest.Right != null)
                    {
                        leftest.Right.Parent = leftest.Parent;
                    }
                }
            }
        }

        public BinaryTree<T> GetLeftest(BinaryTree<T> binaryTree)
        {
            var result = binaryTree;
            while (result.Left != null)
            {
                result = result.Left;
            }

            return result;
        }
    }
}
