#ifndef TREES_HELPER_H
#define TREES_HELPER_H
#include <vector>

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

TreeNode *buildTree(const std::vector<int> &values);
void printTree(TreeNode *root);
void deleteTree(TreeNode *root);
#endif
