package main
    
func main() {

}

type TreeNode struct {
    val int
    left *TreeNode
    right *TreeNode
}

func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    dfs(node.left)
    dfs(node.right)

    return
}