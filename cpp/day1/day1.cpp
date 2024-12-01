#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>

int main() {
    std::ifstream ifs("input.txt");
    int left;
    int right;
    std::vector<int> left_side;
    std::vector<int> right_side;
    while (ifs >> left >> right) {
        left_side.push_back(left);
        right_side.push_back(right);
    }
    std::sort(left_side.begin(), left_side.end());
    std::sort(right_side.begin(), right_side.end());
    int sum = 0;
    /*
    Part 1:
    for (int i = 0; i < left_side.size(); i++) {
        sum += abs(right_side[i] - left_side[i]);
    }
    */
    for (int i = 0; i < left_side.size(); i++) {
        int partial = 0;
        for (int j = 0; j < right_side.size(); j++) {
            if (left_side[i] == right_side[j]) {
                partial++;
            }
        }
        sum += partial * left_side[i];
    }
    std::cout << sum << std::endl;
    return 0;
}