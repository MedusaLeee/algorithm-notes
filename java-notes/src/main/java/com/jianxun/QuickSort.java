package com.jianxun;

import java.util.Arrays;

public class QuickSort {
    private static void quickSort(int[] arr, int start, int end) {
        if (start >= end) {
            return;
        }
        int i, j, index;
        i = start;
        j = end;
        index = arr[i];

        while (i < j) {
            while (i < j && arr[j] >= index) {
                j--;
            }
            while (i < j && arr[i] < index) {
                i++;
            }
            if (i < j) {
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
        if (arr[i] < arr[start]) {
            int temp = arr[i];
            arr[i] = arr[start];
            arr[start] = temp;
        }

        quickSort(arr, start, i - 1);
        quickSort(arr, i + 1, end);

    }

    public static void main(String[] args) {
        int[] a = {1, 2, 4, 5, 7, 4, 5, 3, 9, 0};
        quickSort(a, 0, a.length - 1);
        System.out.println(Arrays.toString(a));
    }
}
