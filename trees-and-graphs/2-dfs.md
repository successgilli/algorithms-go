# Binary Trees - DFS
How do we traverse a binary tree ?

We traverse the tree just like we do a Linked list by using right and left pointers.

The difference is that with Linked lists, we usually do it iteratively, while with Binary trees, we usually do it recursively.

There are two main types of tree traversal; BFS and DFS (`preorder`, `inorder`, and `postoder`).

## Depth-first search - DFS
In DFS, we prioritize depth by traversing as far down the tree in one direction before going in the other direction.

To be able to track back our position, we use recursion but this can also be implemented using a stack.

E.G:
```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    dfs(node.left)
    dfs(node.right)

    return
}
```

### Pre-order Traversal
Here, we first visit the rot, then the left subtree and finally, the right subtree.

Suitable for running an operation on all nodes.

```go
func preOrderDFS(node *TreeNode) {
    if node == nil {
        return
    }

    fmt.Println(node.val)
    preOrderDFS(node.left)
    preOrderDFS(node.right)

    return
}
```

### In-order Traversal
Here, we first check the left subtree, visit the root, then move to the right Subtree.

In-order Traversal (Left → Root → Right)
- Visits left subtree, processes root, then right subtree
- On a BST: produces sorted ascending output
- On a BT: simply traverses in Left → Root → Right order
  with no guaranteed sorting
- Time complexity: O(n)

```go
func inOrderDFS(node *TreeNode) {
    if node == nil {
        return
    }

    inOrderDFS(node.left)
    fmt.Println(node.val)
    inOrderDFS(node.right)

    return
}
```

### Post-order Traversal
Here, we first check the left subtree, visit the right subtree, then lastly, the root

Post-order Traversal (Left → Right → Root)

```go
func postOrderDFS(node *TreeNode) {
    if node == nil {
        return
    }

    postOrderDFS(node.left)
    postOrderDFS(node.right)
    fmt.Println(node.val)

    return
}
```
