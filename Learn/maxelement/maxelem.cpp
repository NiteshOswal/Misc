/**
 * CPP alternative to find max element from an array
 */
#include <iostream>
using namespace std;

int max_element_position(int a[], int size) {
    int i, positionOfMaxElement = 0;
    if(size < 1) return -1;
    for(i = 0; i < size; ++i)
        if(a[i] > a[positionOfMaxElement])
            positionOfMaxElement = i;
    return positionOfMaxElement;
}

int main() {
    int a[10], size = 0, i, pos;
    cin >> size;
    for(i = 0; i < size; ++i) {
        cin >> a[i];
    }
    pos = max_element_position(a, size);
    if(pos == -1) return 1;
    cout << pos << "\n" << a[pos];
    return 0;
}