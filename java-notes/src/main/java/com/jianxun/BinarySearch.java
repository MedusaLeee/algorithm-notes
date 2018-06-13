package com.jianxun;

public class BinarySearch {
    public static int binarySearch(int[] s, int k) {
        int low = 0;
        int high = s.length - 1;
        while (low < high) {
            int mid = (low + high) / 2;
            int guess = s[mid];
            if (guess == k) {
                return mid;
            }
            if (guess > k) {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        int result = binarySearch(arr, 1);
        System.out.println(result);
        result = binarySearch(arr, 3);
        System.out.println(result);
        result = binarySearch(arr, 6);
        System.out.println(result);
    }
}
