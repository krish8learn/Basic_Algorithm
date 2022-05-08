package Binary_search.JAVA;
import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class BS{
    public static int BinarySearch(int[] list, int targetValue){
        int min = 0, max = list.length-1;
        while(max>=min){
            int guess = min + (max-min)/2;
            if (list[guess]==targetValue){
                return guess;
            }else if (list[guess]>targetValue){
                max = guess -1;
            }else if (list[guess]<targetValue){
                min = guess +1;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        System.out.println(BS.BinarySearch(new int[]{22, 44, 66, 88}, 88));
    }

    @Test 
    public void TestBinarySearch(){
        assertEquals(3, BS.BinarySearch(new int[]{22, 44, 66, 88}, 88));
    }
}