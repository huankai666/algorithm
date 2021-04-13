import java.util.Arrays;
import java.util.Scanner;

public class BubbleSort {
    public static void main(String[] args) {
        String[] strArr = new Scanner(System.in).nextLine().split(" ");
        int length = strArr.length;
        int[] intArr = new int[length];
        for (int i = 0; i < length; i++) {
            intArr[i] = Integer.parseInt(strArr[i]);
        }
        bubbleSort(intArr);
        System.out.println(Arrays.toString(intArr));
    }

    public static void bubbleSort(int[] arr) {
        if (arr == null || arr.length < 2) return;
        int length = arr.length;
        for (int i = 0; i < length - 1; i++) {
            for (int j = 0; j < length - 1 - i; j++) {
                if (arr[j] > arr[j + 1]) {  // 升序
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }
    }
}
