#include "helper.h"
#include <iostream>

/* 1123. Lowest Common Ancestor of Deepest Leaves
100% time, 39% memory
Lowkey spent like 10 minutes just staring at the graph and dry runing solutions

I'm thinking a DFS that returns the max depth reached by the children.
Whenever those max depths match, you found the lowest ancestor for the leaves of that subtree
Return the node alongside the max depths.
If the depths do not match then return the lowest ancestor of the node with the bigger max depth.
This goes all the way until you reach the root node.
*/
struct lcaDL
{
    int depth;
    TreeNode *lca;
};
lcaDL lcaDLdfs(TreeNode *node, int depth)
{
    lcaDL result{depth, nullptr};
    if (!node)
    {
        return result;
    }

    // do the DFS (post order)
    lcaDL left{lcaDLdfs(node->left, depth + 1)};
    lcaDL right{lcaDLdfs(node->right, depth + 1)};

    // update result struct if needed
    result.depth = std::max(left.depth, right.depth);
    if (left.depth == right.depth)
    {
        result.lca = node;
        return result;
    }
    else if (left.depth > right.depth)
    {
        result.lca = left.lca;
        return result;
    }
    else
    {
        result.lca = right.lca;
        return result;
    }
}
TreeNode *lcaDeepestLeaves(TreeNode *root)
{
    lcaDL result{lcaDLdfs(root, 0)};
    return result.lca;
}
/* Someone's better solution
void dfs(TreeNode* root, int depth, int &max_depth){
    if(root == NULL) return;
    max_depth = max(max_depth, depth);
    dfs(root->left, depth+1, max_depth);
    dfs(root->right, depth+1, max_depth);
}
TreeNode* lcaDeepestLeaves(TreeNode* root, int depth, int max_depth){
    if(root == NULL) return NULL;
    if(depth == max_depth) return root;
    TreeNode* left  = lcaDeepestLeaves(root->left, depth+1, max_depth);
    TreeNode* right = lcaDeepestLeaves(root->right,depth+1, max_depth);
    if(left != NULL & right != NULL) return root;
    return left != NULL ? left : right;
}
TreeNode* lcaDeepestLeaves(TreeNode* root)
{
    int max_depth = 0;
    dfs(root,0, max_depth);
    return lcaDeepestLeaves(root,0,max_depth);
} */

int main()
{
    std::cout << "\n\nLowest Common Ancestor of Deepest Leaves\n";
    std::vector<int> treeValues = {3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}; // -1 represents null
    TreeNode *root = buildTree(treeValues);
    TreeNode *lca = lcaDeepestLeaves(root);
    std::cout << lca->val << "\n";
    deleteTree(root);
    lca = nullptr; // after deleteTree(root) is called, lca is a dangling ptr

    return 0;
}
