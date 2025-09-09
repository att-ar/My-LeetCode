#include <iostream>
#include <vector>
#include <queue>
#include <memory> // For std::unique_ptr
#include "helper.h"

TreeNode *buildTree(const std::vector<int> &values)
{
    if (values.empty())
    {
        return nullptr;
    }

    TreeNode *root = new TreeNode(values[0]);
    std::queue<TreeNode *> q;
    q.push(root);

    size_t i = 1;
    while (!q.empty() && i < values.size())
    {
        TreeNode *current = q.front();
        q.pop();

        if (i < values.size() && values[i] != -1)
        { // Assuming -1 represents null
            current->left = new TreeNode(values[i]);
            q.push(current->left);
        }
        i++;

        if (i < values.size() && values[i] != -1)
        {
            current->right = new TreeNode(values[i]);
            q.push(current->right);
        }
        i++;
    }
    return root;
}

void printTree(TreeNode *root)
{
    if (!root)
    {
        std::cout << "Tree is empty." << std::endl;
        return;
    }

    std::queue<TreeNode *> q;
    q.push(root);

    while (!q.empty())
    {
        int levelSize = static_cast<int>(q.size());
        for (int i = 0; i < levelSize; ++i)
        {
            TreeNode *current = q.front();
            q.pop();

            if (current)
            {
                std::cout << current->val << " ";
                q.push(current->left);
                q.push(current->right);
            }
            else
            {
                std::cout << "null ";
            }
        }
        std::cout << std::endl;
    }
}

void deleteTree(TreeNode *root)
{
    if (root)
    {
        deleteTree(root->left);
        deleteTree(root->right);
        delete root;
    }
}
